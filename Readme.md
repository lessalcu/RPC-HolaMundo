# RPC Hello World

This is a simple project that implements an RPC (Remote Procedure Call) service in Go, allowing users to submit a name via a web form and receive a custom greeting. The project also automatically generates Swagger documentation.

## Features

- **RPC (Remote Procedure Call):** Implements an RPC server that receives a name and returns a custom greeting.

- **Swagger Documentation Generation:** Uses the Swag library to automatically generate API documentation.

- **Web Server:** A web page that allows users to enter their name and receive a greeting.

- **Dockerized:** The project is fully dockerized, making it easy to run in any environment.

## Requirements

- Docker
- Go 1.18 or higher

## Installation

To build the Docker image and run the application, follow these steps:

1. **Clone the repository:**

```bash
git clone https://github.com/lessalcu/RPC-HolaMundo.git
cd RPC-HolaMundo
```

2. **Build the Docker image:**

Make sure you are in the root directory of the project and run the following command to build the Docker image:

```bash
docker build -t hello-world-rpc .
```

3. **Run the container:**

After building the image, run the following command to start the container:

```bash
docker run -p 8080:8080 -p 1234:1234 hello-world-rpc
```

This exposes two ports:
- **8080**: To access the web application (form).
- **1234**: For RPC communication.

4. **Access Swagger documentation:**

Once the container is running, you can access the Swagger API documentation at the following link:

```
http://localhost:8080/swagger/index.html
```

5. **Interact with the application:**

Open your browser and go to:

```
http://localhost:8080
```

From there you can enter your name in the form and receive a personalized greeting.

## Docker Hub Commands

To run the application using Docker Hub:

1. **Start the container from Docker Hub:**

```bash
docker pull lssalas/hello-world-rpc
docker run -p 8080:8080 -p 1234:1234 lssalas/hello-world-rpc
```

## Powered by

- Go 1.18+
- Swag (for Swagger documentation)
- Docker