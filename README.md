myapp/
├── cmd/
│   └── main.go
├── internal/
│   ├── adapters/
│   │   ├── http/
│   │   │   ├── user_handler.go
│   │   │   ├── auth_handler.go
│   │   │   └── product_handler.go
│   │   ├── db/
│   │   │   ├── postgres/
│   │   │   │   ├── user_repository.go
│   │   │   │   ├── auth_repository.go
│   │   │   │   └── product_repository.go
│   │   │   └── redis/
│   │   │       └── cache.go
│   │   └── broker/
│   │       └── kafka/
│   │           └── event_publisher.go
│   ├── application/
│   │   ├── commands/
│   │   │   ├── user/
│   │   │   ├── auth/
│   │   │   └── product/
│   │   └── queries/
│   │       ├── user/
│   │       ├── auth/
│   │       └── product/
│   ├── domain/
│   │   ├── user.go
│   │   ├── auth.go
│   │   └── product.go
│   └── ports/
│       ├── repository/
│       │   ├── user_repository.go
│       │   ├── auth_repository.go
│       │   └── product_repository.go
│       ├── cache/
│       ├── broker/
│       └── services/
├── pkg/
│   └── middleware/
├── logs/
│   └── logstash.log
├── config/
│   └── config.go
└── README.md