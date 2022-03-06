package users_test

import (
	"context"
	"testing"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	srv := users.GithubUsersService{}
	users, err := srv.GetUsers(context.Background(), "mj-hagonoy")
	assert.Nil(t, err)
	assert.Len(t, users, 1)
}
