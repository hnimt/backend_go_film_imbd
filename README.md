# Microservice IMBD Film

## Project Architecture

![alt text](architecture.png)

---

```
├── cmd
│   ├── client
│   │   └── crawl_biz.go
│   └── server
│       ├── auth
│       │   └── auth.go
│       ├── biz
│       │   └── biz.go
│       ├── bookmark
│       │   └── bookmark.go
│       └── gate
│           └── gate.go
├── common
│   ├── entity
│   │   ├── bookmark.go
│   │   ├── film.go
│   │   └── user.go
│   ├── helper
│   │   └── validate.go
│   ├── middleware
│   │   └── jwt.go
│   ├── model
│   │   ├── req
│   │   │   └── req_bm.go
│   │   ├── res
│   │   │   └── auth.go
│   │   └── res.go
│   ├── repo
│   │   ├── bookmark_repo.go
│   │   ├── film_repo.go
│   │   └── user_repo.go
│   └── security
│       ├── claims.go
│       ├── jwt.go
│       └── pass.go
├── config
│   ├── cache
│   │   └── redis_cache.go
│   ├── db
│   │   └── postgres.go
│   └── grpc
│       ├── client.go
│       └── client_config.go
├── go.mod
├── go.sum
└── services
    ├── auth
    │   ├── handler
    │   │   └── auth_handler.go
    │   └── pb
    │       ├── auth.pb.go
    │       └── auth.proto
    ├── biz
    │   ├── handler
    │   │   ├── crawl_handler.go
    │   │   └── film_handler.go
    │   └── pb
    │       ├── pb_crawl
    │       │   ├── crawl.pb.go
    │       │   └── crawl.proto
    │       └── pb_film
    │           ├── film.pb.go
    │           └── film.proto
    ├── bookmark
    │   ├── handler
    │   │   └── bookmark_handler.go
    │   └── pb
    │       ├── bookmark.pb.go
    │       └── bookmark.proto
    ├── crawler
    │   └── handler
    │       └── hander_crawl.go
    └── gate
        └── handler
            ├── auth_handler.go
            ├── bookmark_handler.go
            └── film_handler.go
```
