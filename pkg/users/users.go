package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/cache"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/config"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/errors"
)

type GithubUser struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Company     string `json:"company"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"public_repos"`
}

func (user GithubUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

type GithubUsersService struct {
	Cache cache.Cache
}

func NewUsersService() *GithubUsersService {
	return &GithubUsersService{
		Cache: cache.NewRedisCache(
			config.GetConfig().RedisUrl(),
			"",
			0,
			time.Minute*time.Duration(config.GetConfig().Cache.ExpiryMins),
		),
	}
}

type GithubUserArray []GithubUser

func (users GithubUserArray) MarshalBinary() ([]byte, error) {
	return json.Marshal(users)
}

func (users GithubUserArray) Len() int {
	return len(users)
}

func (users GithubUserArray) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(users[i].Name)
	jRune, _ := utf8.DecodeRuneInString(users[j].Name)
	return int32(iRune) < int32(jRune)
}

func (users GithubUserArray) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func (srv GithubUsersService) GetUsers(ctx context.Context, usernames ...string) (GithubUserArray, error) {
	if len(usernames) > 10 {
		return nil, errors.MaxLenghtError(10, len(usernames))
	}

	var users GithubUserArray

	data := srv.Cache.Get(strings.Join(usernames, ","))
	if data != nil {
		if err := json.Unmarshal([]byte(*data), &users); err == nil {
			return users, nil
		}
	}

	var wg sync.WaitGroup
	getUser := func(wg *sync.WaitGroup, username string) {
		defer wg.Done()
		user, err := srv.getUser(ctx, username)
		if err != nil {
			errors.Send(err)
			return
		}
		users = append(users, *user)
	}

	for _, username := range usernames {
		wg.Add(1)
		go getUser(&wg, username)
	}
	wg.Wait()

	sort.Sort(users)
	go srv.Cache.Set(strings.Join(usernames, ","), users)
	return users, nil
}

func (srv GithubUsersService) getUser(ctx context.Context, username string) (*GithubUser, error) {
	if username == "" {
		return nil, fmt.Errorf("empty username")
	}
	data := srv.Cache.Get(username)
	var user GithubUser
	if data != nil {
		if err := json.Unmarshal([]byte(*data), &user); err == nil {
			return &user, nil
		}
		//in case of marshall error, attempt to get data from github api
	}
	resp, err := http.Get(fmt.Sprintf("%s/%s", config.GetConfig().Github.GetUserAPI, username))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.HttpNotFoundError
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return nil, err
	}

	go srv.Cache.Set(username, user)
	return &user, nil
}
