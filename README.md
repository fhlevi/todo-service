# Todo REST API with Go

A simple Todo List REST API built with Go using SOLID principles. This project demonstrates a clean architecture pattern and includes:

- RESTful endpoints for CRUD operations
- UUID-based IDs and automatic timestamp formatting
- CORS support for React/Vite frontend
- Gorilla Mux router with `/api/v1` prefix

## 📁 Project Structure

```
todo-service/
├── config/            # Configuration setup
├── controller/        # HTTP handlers
├── model/             # Data models
├── repository/        # Data access layer
├── router/            # Router definitions
├── service/           # Business logic
├── main.go            # Entry point
├── go.mod             # Go modules
├── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20+

### Installation

```bash
git clone https://github.com/yourusername/todo-service.git
cd todo-service
go mod tidy
go run main.go
```

The server will start on `http://localhost:8000`

## 🔧 API Endpoints

| Method | Endpoint             | Description          |
|--------|----------------------|----------------------|
| GET    | `/api/v1/todos`      | Get all todos        |
| POST   | `/api/v1/todos`      | Create new todo      |
| PUT    | `/api/v1/todos`      | Update existing todo |
| DELETE | `/api/v1/todos?id=`  | Delete a todo        |

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
  "id": "uuid-generated",
  "todo": "Go to the gym",
  "date": "Monday, May 26, 2025 at 12:30 PM"
}
```