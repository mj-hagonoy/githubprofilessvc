# github.com/mj-hagonoy/githubprofilessvc

A small Golang service that accepts a list of github usernames (max 10 names) and returns the following basic information:
- name
- login
- company
- number of followers
- number of public repos

## Table of contents
  * [Architecture](#architecture)
    + [Routing or Framework](#routing-or-framework)
    + [Caching](#caching)
  * [REST API](#rest-api)
    + [Get github users](#get-github-users)
  * [Dependencies](#dependencies)
    + [External packages](#external-packages)
  * [Installation](#installation)
  * [Configuration](#configuration)
    + [Environment variables:](#environment-variables-)
  * [External API](#external-api)
  * [Testing](#testing)
  * [Lint](#lint)
  * [Requirements Checklist](#requirements-checklist)
  * [TODO](#todo)
  * [License](#license)
  * [Links](#links)
    + [References](#references)

## Architecture
** Under construction

### Routing or Framework
Due to the size, `net/http` standard package was used.

When the need arises, we can easily integrate other routing libraries.

For the same reason, web frameworks (ie gin, echo) were not used.

### Caching
[Redis](https://redis.io/) is used for caching for it's wide support and large community.
Available in:
- [Google Cloud](https://cloud.google.com/memorystore/docs/redis)
- [Azure](https://azure.microsoft.com/en-us/services/cache/)
- [AWS](https://aws.amazon.com/elasticache/redis/)

Another library worth mentioning is [`ristretto`](https://github.com/dgraph-io/ristretto)

## REST API
### Get github users
Returns public information of github users
```
GET /api/v1/github/users?usernames=<list of users>
```
where:

- `usernames` : is a comma separated list of strings

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


### External packages
- [`gopkg.in/yaml.v2`](https://github.com/go-yaml/yaml/tree/v2.4.0) 
- [`github.com/go-redis/redis/v8`](https://github.com/go-redis/redis)
- [`github.com/stretchr/testify`](https://github.com/stretchr/testify)

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

## Testing
In root directory, run below command
```
go test ./...
```

## Lint
Run in project root directory
```
golangci-lint run ./...
```
See [golangci-lint](https://golangci-lint.run/)

## Requirements Checklist
- [x] Return the following github user information: name, login, company, number of followers, number of public repos
- [x] Max input of 10 names
- [x] Returned users are sorted alphabetically by name
- [x] If username not found, should not fail other requested usernames
- [x] Implement caching layer with 2 minutes expiry. If user's information is cached, it should NOT hit Github again
- [x] error handlings and frameworks
- [x] use regular http calls to hit github's API
- [x] use github API endpoint https://api.github.com/users/{username

## TODO
- [ ] Security (ie Authorization)
- [ ] Add unit tests and integration tests
- [ ] Deploy application (+load balancer, +api gateway, GCP or Azure)
- [ ] CI/CD
- [ ] Update logging to have (1) separate info and error loggers (2) log to external file
- [ ] Documentation (ie comments)

## License
See [`Apache License v2.0`](./LICENSE)

## Links
- [LinkedIn Profile](https://www.linkedin.com/in/jenessa-hagonoy-023b09b1/)

### References
- [Introducing Ristretto: A High-Performance Go Cache](https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/)
- [A complete Go cache library that brings you multiple ways of managing your caches](https://golangexample.com/a-complete-go-cache-library-that-brings-you-multiple-ways-of-managing-your-caches/)