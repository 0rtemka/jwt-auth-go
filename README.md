Это приложение является частью сервиса аутентификации 
пользователей с использованием Access(JWT) и Refresh токенов.

# Технологии
- Go
- MongoDB
- JWT

# Запуск
Для запуска потребуется:
1. СУБД MongoDB, развернутая локально или удаленно
2. Установить актуальные параметры для базы данных (адрес, на котором работает Mongo и название базы данных) в файле **configs/config.toml** 
3. Импоритировать данные из файла **data/test.users.json** в базу данных 
4. Выполнить команду `go run ./cmd/main.go`

# Взаимодействие 
Предполагается, что приложение запущено локально на порту 8080, поэтому 
все примеры будут использовать локальный адрес `localhost:8080`.

## Аутентификация

1. `POST -> http://localhost:8080/auth/sign-in/{user_id}` - получение Access и Refresh токена 
по GUID пользователя. При этом access токен отправляется в теле ответа, а refresh устанавливается в cookie 'refresh_token'.

2. `POST -> http://localhost:8080/auth/refresh/{user_id}` - выполнение операции refresh на пару access, refresh токенов.
   Refresh токен должен быть установлен в cookie 'refresh_token', а Access токен должен быть указан в Authorization Header в следующем формате:
   `Bearer {token}`

### Примеры 

`curl -XPOST 'http://localhost:8080/auth/sign-in/64e2632eeced0aeb9b550162'`

```json
{
"access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI3NzE4MTcsImlhdCI6MTY5Mjc3MDkxNywidXNlcl9pZCI6IjY0ZTI2MzJlZWNlZDBhZWI5YjU1MDE2MiJ9.GLFduKYS9rxVgth0stMJH6Q03q6tF0jTNeb3GPw5qKPPQVUIZNiiN1z6ocdKSK1w_U3A_mpTo3g-Hx8YI4v5Aw"
}
```

`curl -XPOST -H 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI3NzE4MTcsImlhdCI6MTY5Mjc3MDkxNywidXNlcl9pZCI6IjY0ZTI2MzJlZWNlZDBhZWI5YjU1MDE2MiJ9.GLFduKYS9rxVgth0stMJH6Q03q6tF0jTNeb3GPw5qKPPQVUIZNiiN1z6ocdKSK1w_U3A_mpTo3g-Hx8YI4v5Aw' 'http://localhost:8080/auth/refresh/64e2632eeced0aeb9b550162'`

```json
{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI3NzE4NzYsImlhdCI6MTY5Mjc3MDk3NiwidXNlcl9pZCI6Ik9iamVjdElEKFwiNjRlMjYzMmVlY2VkMGFlYjliNTUwMTYyXCIpIn0.qD5LR7WYA5CV7cwJJNgtC25YHLCLJO94gpyfe2t2qJGKoNUXfRwp5xs5MYW4LWkAEBykYpYaN7P1qldGJQ8Kzw"
}
```

## Пользователи 

`GET -> http://localhost:8080/users` - получение всех пользователей из базы данных.

### Примеры 

`curl -XGET 'http://localhost:8080/users'`

```json
[
    {
        "id": "64e257abeced0aeb9b550160",
        "email": "petya@mail.com",
        "name": "Petya"
    },
    {
        "id": "64e26295eced0aeb9b550161",
        "email": "vasya@mail.com",
        "name": "Vasya"
    },
    {
        "id": "64e2632eeced0aeb9b550162",
        "email": "kolya@mail.com",
        "name": "Kolya"
    }
]
```

## Проверка

`GET -> http://localhost:8080/` - данный адрес предназначен для проверки работы acсess токена. Если токен корректный, то в ответе будет сообщение 'Secured Route',
а иначе - сообщение об ошибке со статусом 401 Unauthorized.
