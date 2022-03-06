# github.com/mj-hagonoy/githubprofilessvc

A small Golang service that accepts a list of github usernames (max 10 names) and returns the following basic information:
- name
- login
- company
- number of followers
- number of public repos

## REST API
### Get github users
Returns public information of github users
```
GET /api/v1/github/users?username=<list of users>
```
where:

- `username` : is a comma separated list of strings

Default Response
```
Status: 200 OK
Headers:
    - Content-type: "application/json"
Body: 
[
    {
        "name": "Jen Hagonoy",
        "login": "mj-hagonoy",
        "company": "",
        "followers": 1,
        "public_repos": 6
    }
]
```

Empty Response
```
Status: 204 No Content
```

Error Response: Limit reached
```
Status: 400 Bad Request

Body: 
{
    "error": "input error: expected 10, got 17"
}
```

## Dependencies
- [`docker`, `docker-compose`](https://www.docker.com/)
- [`redis`](https://hub.docker.com/_/redis) - caching layer

## Installation
Development workspace
```
git clone https://github.com/mj-hagonoy/githubprofilessvc.git
cd githubprofilessvc
docker-compose up -d
```

## Configuration
See [config.yaml](./config.yaml)
```
host: "localhost"
port: ":8080"
github:
  get-user-api: "https://api.github.com/users"
cache:
  type: "redis"
  host:  "localhost"
  port: ":6379"
  expiry-mins: 2
```

### Environment variables:
- `REDIS_URL` : redis url to be used, default value "localhost:6379"

## External API
- [`GET https://api.github.com/users/{username}`](https://docs.github.com/en/rest/reference/users#get-a-user)

## External packages
- [`gopkg.in/yaml.v2`](https://github.com/go-yaml/yaml/tree/v2.4.0) 
- [`github.com/go-redis/redis/v8`](https://github.com/go-redis/redis)
- [`github.com/stretchr/testify`](https://github.com/stretchr/testify)


## License
See [`Apache License v2.0`](./LICENSE)