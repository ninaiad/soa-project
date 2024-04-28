# Main service REST API

Основной сервис социальной сети, реализующий аутентификацию пользователей и работу с их личными данными.

Описание методов в спецификации Open API находятся в файле [`openapi.yml`](https://github.com/ninaiad/toy-social/blob/main/main_service/openapi.yml).

## Структура программы

Серверная часть веб-приложения разделена на три основных части: 
1. Handler
2. Service
3. Database

### Handler

Уровень Handler обрабатывает HTTP запросы.

#### /internal/handler/auth.go
Обработчики HTTP запросов, связанные с аутентификацией.

#### /internal/handler/posts.go
Обработчики handlers HTTP запросов, связанные с созданием и просмотром постов пользователей (запросы перенаправляются в gRPC сервер Posts Service).

### Service

Уровень реализующий логику обработки данных пользователей, аутентификацию.

#### /internal/service/auth.go
При создании новой учетной записи или авторизации в уже существующую, в случае успешного обновления базы данных, для клиента генерируется Bearer token по стандарту JWT (JSON Web Token), используя который в последующих запросах, клиент получает доступ к просмотру и обновлению своих данных. 

### Database

На данном уровне программы, происходит работа с СУБД: подключение к ней (а именно к отдельному приложению PostgreSQL, обернутому в Docker контейнер), инициализация базы данных с помощью файлов миграций, а также обновление данных в ней.

## Примеры запросов

#### Создание пользователя
```sh
curl --location 'http://localhost:8000/sign-up' \
--header 'Content-Type: application/json' \
--data '{
"username": "bluefinch11",
"password": "vEryStr0ngP4ssw0Rd1ndEEd"
}'
```

#### Обновление личных данных
```sh
curl --location --request PUT 'http://localhost:8000/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyODA3MDIsImlhdCI6MTcxNDIzNzUwMiwidXNlcl9pZCI6MX0.N3Ya20bKxFRba459_B7pItsL1wdESqZtkR3F3GWWft4' \
--data '{
"name" : "Blue"
}'
```

#### Создание поста
```sh
curl --location 'http://localhost:8000/post' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyODA3MDIsImlhdCI6MTcxNDIzNzUwMiwidXNlcl9pZCI6MX0.N3Ya20bKxFRba459_B7pItsL1wdESqZtkR3F3GWWft4' \
--data '{
"text" : "first post"
}'
```

#### Просмотр поста
```sh
curl --location 'http://localhost:8000/post?id=7' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyODA3MDIsImlhdCI6MTcxNDIzNzUwMiwidXNlcl9pZCI6MX0.N3Ya20bKxFRba459_B7pItsL1wdESqZtkR3F3GWWft4'
```

#### Просмотр нескольких постов
```sh
curl --location 'http://localhost:8000/posts?page_num=1&page_size=5' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQyODA3MDIsImlhdCI6MTcxNDIzNzUwMiwidXNlcl9pZCI6MX0.N3Ya20bKxFRba459_B7pItsL1wdESqZtkR3F3GWWft4'
```