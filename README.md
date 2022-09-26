# go-gin-users-crud
- REST Api using Go's Gin Gonic Framework to execute CRUD Operations on MongoDB
- Uses MVC Architecture 


**APIs**

| Name    | method | url                  |
| ------- | ------ | -------------------- |
| Get All | GET    | /v1/user/getall      |
| Get     | GET    | /v1/user/get:name    |
| Create  | POST   | /v1/user/create      |
| Update  | PATCH  | /v1/user/update      |
| Delete  | DELETE | /v1/user/delete:name |

Todo:
- Init Go Module

- Install Dependencies:
```sh
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver
go get github.com/joho/godotenv
```

- Run App
```sh
go run main.go
```


