# MyGolangProject

Golang test

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

- [Docker](https://www.docker.com/) installed on your machine.

### Installing

1. Clone the repository:

    ```bash
    git clone https://github.com/RedFC/testproject.git
    cd testproject
    ```

2. Create a `.env` file in the project root and configure the MySQL database connection:

    ```env
    DATABASE_URI=root:<password>@tcp(<host>:3306)/<database>?charset=utf8mb4&parseTime=True&loc=Local
    JWT_SECRET=super-secret-jwt-2222$2232^2323232-super-secret-jwt
    ```

3. two ways

    Build and run the Docker containers:

    ```bash
    docker-compose up --build
    ```
    Or if not working with dockers

    create a mysql instance and change the env vars as per your database name, password and host

    ```bash
    go run main.go
    ```

The application should be accessible at [http://localhost:8080](http://localhost:8080).

## Usage

- Postman collection is inside the repository