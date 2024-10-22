# FetchProxy

## Overview

This project is a simple HTTP proxy server written in Go. It listens for incoming HTTP POST requests, fetches the content from the specified URL in the request payload, and returns the fetched content as the response.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Server

To run the server, use the following command:

```sh
go run main.go
```

By default, the server listens on localhost:8767. You can change the host and port using command-line flags:

```sh
go run main.go -host <host> -port <port>
```

#### Usage
The server exposes a single endpoint:

- ``POST /proxy``: Accepts a JSON payload with a ``url`` field and returns the content fetched from the specified URL.

##### Example Request

```http
POST http://localhost:8767/proxy
Content-Type: application/json

{
    "url": "https://google.com"
}
```

##### Example Response

```http
HTTP/1.1 200 OK
Content-Type: application/json

<content of the specified URL>
```

#### Testing

You can use the ``test.http`` file to test the server endpoints. Open the file in Visual Studio Code and use the REST Client extension to send requests.

### Acknowledgements

- Go documentation: https://golang.org/doc/
- REST Client extension for Visual Studio Code: https://marketplace.visualstudio.com/items?itemName=humao.rest-client