package githubprofilessvc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type GithubUser struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Company     string `json:"company"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"public_repos"`
}

type GithubUsersService struct{}

func (srv GithubUsersService) GetUsers(ctx context.Context, usernames ...string) ([]GithubUser, error) {
	var users []GithubUser
	var wg sync.WaitGroup
	for _, username := range usernames {
		go func(wg *sync.WaitGroup, username string) {
			defer wg.Done()
			user, err := srv.GetUser(ctx, username)
			if err == nil {
				users = append(users, user)
			}
		}(&wg, username)
	}
	wg.Wait()
	return users, nil
}

func (srv GithubUsersService) GetUser(ctx context.Context, username string) (GithubUser, error) {
	var user GithubUser
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", username))
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
