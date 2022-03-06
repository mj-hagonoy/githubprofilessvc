package users_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/config"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/users"
	"github.com/stretchr/testify/assert"
)

type testCache struct {
	store map[string]string
	mutex sync.RWMutex
}

func (cache *testCache) Set(key string, value interface{}) {
	bytesData, _ := json.Marshal(value)
	cache.mutex.Lock()
	cache.store[key] = string(bytesData)
	cache.mutex.Unlock()
}

func (cache *testCache) Get(key string) *string {
	val, exists := cache.store[key]
	if exists {
		return &val
	}
	return nil
}

func TestGetUsers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		{
			"name": "Jen Hagonoy",
			"login": "mj-hagonoy",
			"company": "",
			"followers": 1,
			"public_repos": 6
		}
		`))
	}))
	defer ts.Close()
	conf := config.NewConfig()
	conf.Github.GetUserAPI = ts.URL

	srv := users.GithubUsersService{
		Cache: &testCache{
			store: make(map[string]string),
		},
	}
	users, err := srv.GetUsers(context.Background(), "mj-hagonoy")
	assert.Nil(t, err)
	assert.Len(t, users, 1)
}
