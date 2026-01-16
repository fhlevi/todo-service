
# Go Todo API

This project is a simple Todo List REST API built with Go.

## ✨ Features

- **RESTful Endpoints**: Provides clean and predictable endpoints for CRUD operations.
- **Dockerized**: Easily run the entire application with a single `docker-compose` command.
- **CORS Support**: Pre-configured to allow requests from any origin.
- **API Documentation**: Includes a Swagger UI for easy testing and exploration of the API.

## 📁 Project Structure

```
.
├── database/
├── docs/
├── handlers/
├── models/
├── services/
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/fhlevi/todo-service.git
    cd todo-service
    ```

2.  **Run with Docker Compose:**
    ```bash
    docker-compose up --build
    ```

The API will be available at `http://localhost:3000`.

## 🔧 API Endpoints

| Method | Endpoint         | Description          |
|--------|------------------|----------------------|
| GET    | `/api/todo/`     | Get all todos        |
| GET    | `/api/todo/:id`  | Get a single todo    |
| POST   | `/api/todo/`     | Create a new todo    |
| PUT    | `/api/todo/:id`  | Update an existing todo |
| DELETE | `/api/todo/:id`  | Delete a todo        |

## 📚 API Documentation

A Swagger UI is available to visualize and interact with the API's resources.

- **Swagger UI**: `http://localhost:3000/api/docs`

## 🧪 Sample Todo Payload

### Request (POST)

```json
{
  "todo": "Go to the gym"
}
```

### Response

```json
{
  "id": 1,
  "todo": "Go to the gym",
  "date": "2026-01-13T10:00:00Z"
}
```



