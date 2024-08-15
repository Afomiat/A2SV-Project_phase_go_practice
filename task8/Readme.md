# Clean Architecture RESTful API with MongoDB and JWT Authentication

This repository contains a backend RESTful API built using Clean Architecture principles in Go. The API is designed with scalability and maintainability in mind, utilizing the Gin web framework for HTTP routing and MongoDB as the database. JWT (JSON Web Tokens) is used for secure user authentication and authorization.

## Features

- **User Management**: Register, log in, and manage user profiles with role-based access control.
- **Task Management**: Create, read, update, and delete tasks, with both personal and admin views.
- **Security**: Authentication and authorization implemented via middleware, ensuring secure access to resources.
- **Scalability**: Clear separation of concerns across different layers, promoting scalability and ease of maintenance.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.16+**: [Install Go](https://golang.org/doc/install)
- **MongoDB**: [Install MongoDB](https://docs.mongodb.com/manual/installation/)

## Installation

Follow these steps to set up and run the project locally:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/teklumt/A2SV-Backend-Tasks-2024.git
    cd A2SV-Backend-Tasks-2024/Task7-%20Clean%20Architecture
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Set up environment variables**:

   - Create a `.env` file in the root directory.
   - Add the following environment variables:

    ```
    MONGO_URI=your_mongodb_uri
    JWT_SECRET=your_jwt_secret_key
    ```

4. **Run the application**:

    ```bash
    go run main.go
    ```

## API Endpoints

### User Operations:

- **GET /users**: Retrieve all users (Admin only).
- **GET /users/:id**: Retrieve a user by ID.
- **GET /users/me**: Retrieve the authenticated user's profile.
- **DELETE /users/:id**: Delete a user by ID.

### Task Operations:

- **POST /tasks**: Create a new task.
- **GET /tasks**: Retrieve all tasks (Admin only).
- **GET /tasks/:id**: Retrieve a task by ID.
- **GET /tasks/me**: Retrieve tasks assigned to the authenticated user.
- **DELETE /tasks/:id**: Delete a task by ID.
- **PUT /tasks/:id**: Update a task by ID.

## Testing

To run tests for this project:

```bash
go test ./...
