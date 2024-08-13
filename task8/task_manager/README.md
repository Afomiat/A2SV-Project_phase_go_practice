# Task Manager Project with Clean Architecture

## Overview

This project is a Task Manager application built using Clean Architecture principles. It involves CRUD operations (Create, Read, Update, Delete) for managing users and tasks. The system supports role-based access control (RBAC) to ensure different roles (e.g., User and Admin) have appropriate permissions.

## 🚀 Getting Started

In this project, you'll build a RESTful API to manage tasks and users. The focus is on maintaining a modular, maintainable, and testable system through the separation of concerns. Below is an overview of the architecture and implementation steps.

## 🔨 Project Structure

The project follows a Clean Architecture structure with the following layers:

- **Entities**: Contains the core business logic (e.g., `User` and `Task` entities).
- **Use Cases**: Application-specific business rules (e.g., `CreateUser`, `AssignTask`).
- **Interface Adapters**: Converts data from the database to entities and handles communication (e.g., REST controllers, data mappers).
- **Frameworks and Drivers**: External services and tools integration (e.g., Database, HTTP Server).

## 📂 Folder Structure

The project's folder structure is organized as follows:


task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go
    └── user_usecases.go
🧑‍💻 CRUD Operations
The system supports CRUD operations on both Users and Tasks, with different permissions based on roles:

POST /users: Create a new user. Admins can create users with different roles (e.g., User, Admin).
GET /users/{id}: Retrieve user details. Admins can retrieve any user's details, while Users can only retrieve their own.
PUT /tasks/{id}: Update task details. Users can update their own tasks; Admins can update any task.
DELETE /users/{id}: Delete a user. Only Admins can delete users.