# Toy social

A toy social network backend with microservice architecture.

The application is accessible via Gateway service [API]()

## Services

See architecture details in [`/docs`]()

#### Gateway
- Authenticates requests
- Creates new users & saves their data in a database (PostgreSQL)
- Redirects requests to other services via gRPC
- Sends likes & views requests to Apache Kafka

#### Posts
- Manipulates posts data saved in a database (PostgreSQL)

#### Statistics
- Aggregates likes & views statistics saved in a database (Clickhouse)


## Running

The application:
```sh
make run
```

End-to-end tests:
```sh
make test-e2e
```

Posts service integration tests: 
```sh
make test-posts
```

Unit tests can be run with `go test` whithin each of the service directories.

## Examples

#### User sign-up
```sh
curl --location 'http://localhost:8000/user/sign-up' \
--header 'Content-Type: application/json' \
--data '{
  "username": "bluefinch11",
  "password": "vEryStr0ngP4ssw0Rd1ndEEd"
}'
```

```json
{
    "token": "example_token",
    "user_id": 26
}
```

#### User data update
```sh
curl --location --request PUT 'http://localhost:8000/user/' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer example_token' \
--data '{
  "name": "Blue"
}'
```

```json
{
    "username": "bluefinch11",
    "name": "Blue",
    "surname": "",
    "birthday": "0001-01-01T00:00:00Z",
    "email": "",
    "phone": ""
}
```

#### Post creation
```sh
curl --location 'http://localhost:8000/post' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer example_token' \
--data '{
  "text": "First post!"
}'
```

```json
{
    "post_id": 57
}
```

#### View post data

Use `author_id` query parameter to view posts of users different from the one currently logged in.

```sh
curl --location --request GET 'http://localhost:8000/post?id=57&author_id=26' \
--header 'Authorization: Bearer example_token'
```

```json
{
    "id": 57,
    "text": "First post!",
    "time_updated": "2024-06-17T06:50:32Z"
}
```

#### View several posts of a user

Use `author_id` query parameter to view posts of users different from the one currently logged in.

```sh
curl --location 'http://localhost:8000/posts?author_id=1&page_num=1&page_size=5' \
--header 'Authorization: Bearer example_token'
```

```json
{
    "author_id": 1,
    "page_num": 1,
    "page_size": 2,
    "posts": [
        {
            "id": 5,
            "text": "T-O-F-F-E-E",
            "time_updated": "2024-06-15T13:55:11Z"
        },
        {
            "id": 3,
            "text": "mrs dalloway said she would buy the flowers herself",
            "time_updated": "2024-06-15T13:54:54Z"
        }
    ]
}
```

#### View post statistics

```sh
curl --location 'http://localhost:8000/post/statistics?id=2' \
--header 'Authorization: Bearer example_token'
```

```json
{
    "id": 2,
    "author_id": 1,
    "author_username": "VirginiaWoolf",
    "num_likes": 1,
    "num_views": 4
}
```

#### View top 3 users with most likes

```sh
curl --location 'http://localhost:8000/posts/statistics/users?k=3&event_type=view' \
--header 'Authorization: Bearer example_token'
```

```json
{
    "users": [
        {
            "id": 1,
            "username": "VirginiaWoolf",
            "num_likes": 3,
            "num_views": 5
        },
        {
            "id": 3,
            "username": "DFW",
            "num_likes": 3,
            "num_views": 2
        }
    ]
}
```
