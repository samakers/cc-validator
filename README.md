# Credit Card Validator

This is a simple Go application that validates credit card numbers using the Luhn algorithm. It includes an HTTP server that accepts GET requests with a JSON payload containing a credit card number, and responds with a JSON payload indicating whether the number is valid.

## Prerequisites

- Go 1.16 or later

## Installation

1. Clone the repository:

    ```bash
    git clone [https://github.com/samakers/cc-validator]
    ```

2. Navigate to the project directory:

    ```bash
    cd cc-validator
    ```

3. Build the application:

    ```bash
    go build -o app
    ```

## Usage

1. Start the server:

    ```bash
    ./app
    ```

2. Send a GET request to the `/validate` endpoint with a JSON payload containing the credit card number. For example, using curl:

    ```bash
    curl -X GET -H "Content-Type: application/json" -d '{"num":"1234"}' http://localhost:8090/validate
    ```

    Or using Postman:

    - Set the request method to GET.
    - Enter the URL `http://localhost:8090/validate`.
    - Set the request body to raw JSON and enter `{"num":"1234"}`.

3. The server will respond with a JSON payload indicating whether the number is valid according to the Luhn algorithm. For example:

    ```json
    {
        "is_valid": false
    }
    ```
