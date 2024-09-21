# Docker+Nginx 🚀
<a href="https://gofiber.io">
  <picture>
    <source height="125" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
    <img height="125" alt="Fiber" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
  </picture>
</a>


## Features

- **Docker and Nginx** : Deploy in docker using 5 replicas and load balancer with Nginx
- **Logger**: The application includes a request logger for monitoring HTTP requests.

## Endpoints

| Name         | Rute     | Parameters | State     | Protected | Method |
|--------------|----------| ---------- | --------- | --------- |--------|
| Hello        | /hello   | No         | Completed | No        | GET    |

## Getting Started

To get a local copy up and running, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Build the Docker image with docker compose
4. Run the Docker compose composition
  ```bash
  docker compose up --build
  ```
5. Access the application at `http://localhost:8080/hello`.

