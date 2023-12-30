# Lacoste Romberg D179 | Golang-Microservices
> This project is a work in progress just for fun on the weekend, new features will be added over time.

The main idea of creating this project is implementing an infrastructure for up and running distributed system with the latest technology and we will not deal mainly with business.

## The Goals of This Project

- `Rabbitmq` for `Event Driven Architecture` between our microservices with `streadway/amqp` library.
- `gRPC` for `internal communication` between our microservices with `grpc/grpc-go` library.
- `Postgres` for `database` in our microservices with `go-gorm/gorm` library.
- `go-playground/validator` for `validating input` data in the REST calls.
- `OpenTelemetry` for `distributed tracing` top of `Jaeger`.
- `OAuth2` for implementation `authentication` and `authorization` with `go-oauth2/oauth2` library.
- `Echo framework` for `RESTFul api`.
- `Swagger` with `swaggo/swag` library for api documentation.
- `uber-go/fx` library for `dependency injection`.
- `Viper` for `configuration management`.
- `logrus` as a `structured logger`.
- `Unit Testing`,`Integration Testing` and `End To End Testing` for testing level.
- `Docker-Compose` for our `deployment` mechanism.
- `OpenTelemetry` for `monitoring` top of `Prometteuse` and `Grafana`
- `MongoDB` for read side with `mongo-driver`.
- `Domain Driven Design` (DDD) to implement all `business` processes in microservices.