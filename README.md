# JWT Authentication App

This is a Go-based application that implements JWT (JSON Web Token) authentication.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [File Structure](#file-structure)
- [Dependencies](#dependencies)

## Introduction

The `jwt_authentication_app` is designed to provide secure authentication using JWTs. It includes user management features such as user_registration, get_user and user_login. It also has feature of refresh token which reduces redundant logins.

## Features

- User registration
- Get user
- User login
- JWT-based authentication
- Refresh JWT token

## Installation

To install and run this application, follow the steps below:

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/jwt_authentication_app.git
    cd jwt_authentication_app
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go user.go
    ```

## Usage

Once the application is running, you can interact with the API using tools like `curl` or Postman. Below are the main endpoints available.

## API Endpoints

### Register a new user

- **URL:** `http://localhost:9090/user/create`
- **Method:** `POST`
- **Body:**
    ```json
    {
      "FirstName":"your_first_name",
      "LastName":"your_last_name",
      "Email":"your_email",
      "Password":"your_password" 
    }
    ```
    
### Get user info

- **URL:** `http://localhost:9090/user/get`
- **Method:** `GET`
- **Headers:**
    `token - token_you_get_from_login_response`

### Refresh token
- **URL:** `http://localhost:9090/refresh`
- **Method:** `POST`
- **Headers:**
    `token - token_you_get_from_login_response`

## File Structure

```plaintext
jwt_authentication_app/
│
├── main.go          # Entry point of the application
├── user.go          # User model and related functions
├── go.sum           # Dependencies file
└── README.md        # Project documentation
```

## Dependencies

- github.com/golang-jwt/jwt/v4 v4.5.0
- github.com/gorilla/mux v1.8.0
- gorm.io/driver/mysql v1.3.6
- gorm.io/gorm v1.23.8
- github.com/go-sql-driver/mysql v1.6.0 // indirect
- github.com/jinzhu/inflection v1.0.0 // indirect
- github.com/jinzhu/now v1.1.5 // indirect
