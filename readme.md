go mod init github.com/lenarbatdalov/learn

# tag
```
tag v1 - добавил структуру, сервис и контроллер
tag v2 - добавил logger, basic-auth и провел dump(gindump) + тестовый запрос(tester.go)
tag v3 - data binding and validation
tag v4 - HTML, Templates and Multi-Route Grouping
tag v5 - docker and jwt
tag v6 - gorm and CRUD
tag v7 - swagger added
```

# swagger
```
go mod tidy
go mod vendor
chmod +x swagger.sh
./swagger.sh
go run main.go
http://localhost:8080/swagger/index.html
```

# basic auth
```
login:      pragmatic
password:   reviews
```

# jwt
```
получаю токен
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'username=pragmatic' \
--data-urlencode 'password=reviews'

добавляю контент
curl --location --request POST 'http://localhost:8080/api/videos' \
--header 'Authorization: Bearer ...' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Cool",
    "description": "description",
    "url": "https://www.youtube.com/embed/p3maH9G_DLM",
    "author": {
        "firstname": "firstname",
        "lastname": "lastname",
        "age": 30,
        "email": "test@mail.ru"
    }
}'
```

# json data
```
{
    "title": "Cool title",
    "description": "description 1",
    "url": "https://www.youtube.com/embed/sDJLQMZzzM4",
    "author": {
        "firstname": "firstname",
        "lastname": "lastname",
        "age": 1,
        "email": "test@mail.ru"
    }
}
```
