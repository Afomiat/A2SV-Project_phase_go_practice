**[API Documentation ](https://documenter.getpostman.com/view/36737395/2sA3rzKD3H)**
# Task Manager API Documentation 

## Overview: 
This API allows users and admins to manage tasks. All endpoints require authentication via a JWT token. Users have restricted access to their own data, while admins have unrestricted access to all data.

## Authentication and Authorization

## User Roles: 

User: Can only access and manipulate their own tasks.
Admin: Can access and manage all tasks.

## JWT Token: 

Required for all endpoints.
Must be included in the Authorization header as a Bearer token.

## Endpoints

### POST /register

- **Description**: Registers a new user with a specified role (either "user" or "admin").
- **Required Fields**: Username, Password, Role (user/admin)
- **Authorization**: None

#### GET /tasks

- **Description**:  Retrieves a list of tasks.
For Users: Retrieves tasks owned by the authenticated user.
For Admins: Retrieves all tasks.

- **Authorization**: Bearer token

### GET /tasks/

- **Description**:  Retrieves a specific task by its ID.
For Users: Only retrieves the task if it is owned by the authenticated user.
For Admins: Can retrieve any task by ID

- **Authorization**:  Bearer token

### POST /tasks

- **Description**: Adds a new task.
For Users: The task is automatically assigned to the authenticated user.
For Admins: The task can be assigned to any user.

- **Authorization**:  Bearer token

### PUT /tasks/

- **Description**: Updates an existing task by its ID.
For Users: Can only update the task if it is owned by the authenticated user.
For Admins: Can update any task.


- **Authorization**: Bearer token

### DELETE /tasks/

- **Description**: Deletes a task by its ID.
For Users: Can only delete the task if it is owned by the authenticated user.
For Admins: Can delete any task.

- **Authorization**: Bearer token

