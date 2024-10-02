# Datetime Server - Fatma Ebrahim

This repository implements a basic HTTP server that returns the current date and time using the HTTP and Gin frameworks. It tests and containerizes the application using Docker and Docker Compose.

## Installation

To install the server package, run the following command:

```shell
git clone https://github.com/codescalersinternships/Datetime-Client-Fatma-Ebrahim.git
```

To install the needed dependencies:

```shell
go mod download
```

## Usage

1. Run both HTTP and Gin servers using Docker Compose with the following command:
   ```shell
   docker compose up
   ```

2. Run both HTTP and Gin servers using the Makefile with targets `format`, `lint`, `build`, `docker`, or use the target `all` to run all of them:
   ```shell
   make all
   ```

Now you have the datetime server running on your machine using either the Gin or HTTP frameworks!
