
# Backend Golang Coding Test

A simple RESTful API in Golang for user management using MongoDB, JWT authentication, and clean architecture principles.
Includes testing, middleware, background tasks, and optionally Docker/gRPC support.

## Features

- User registration & authentication (JWT)

- CRUD for users

- MongoDB integration via official driver

- Middleware for logging

- Background goroutine for user count logging

- Unit tests with MongoDB mocking

- JWT-secured endpoints

- Optional: Docker Compose, validation


## Installation

Docker run for start MongoDB
```bash
  docker-compose up -d
```

Run the API
```bash
  go run main.go
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_HOST`

`DB_PORT`

`DB_NAME`

`PORT`
    
## API Reference

#### Register

```http
  POST /auth/register
```

| Parameter  | Type     | Description                       |
| :--------- | :------- | :-------------------------------- |
| `name`     | `string` | **Required**. User name           |
| `email`    | `string` | **Required**. User email (unique) |
| `password` | `string` | **Required**. User password       |

#### Login

```http
  POST /auth/login
```

| Parameter  | Type     | Description                       |
| :--------- | :------- | :-------------------------------- |
| `email`    | `string` | **Required**. User email (unique) |
| `password` | `string` | **Required**. User password       |

#### Get All Users (protected)
```http
  GET /users
  Authorization: Bearer <token>
```

#### Get User By ID
```http
  GET /users/:id
  Authorization: Bearer <token>
```

#### Create User
```http
  POST /users
  Authorization: Bearer <token>
```

| Parameter  | Type     | Description                       |
| :--------- | :------- | :-------------------------------- |
| `name`     | `string` | **Required**. User name           |
| `email`    | `string` | **Required**. User email (unique) |
| `password` | `string` | **Required**. User password       |

#### Update User
```http
  PUT /users/:id
  Authorization: Bearer <token>
```
| Parameter  | Type     | Description          |
| :--------- | :------- | :------------------- |
| `name`     | `string` |  User name           |
| `email`    | `string` |  User email (unique) |

#### Delete User
```http
  DELETE /users/:id
  Authorization: Bearer <token>
```



## Running Tests

Unit tests use mtest to mock MongoDB operations.

```bash
  go test ./... -v
```

Test coverage includes:

- Register/Login

- Get/Update/Delete User

- JWT validation

- MongoDB mocking with mtest
## Authors

- [@w-worw](https://www.github.com/w-worw)

