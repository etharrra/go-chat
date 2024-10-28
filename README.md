# go-chat

This project is a Go web server built using the [Fiber](https://gofiber.io/) framework. It serves static files, renders HTML templates, and supports WebSocket connections for real-time data communication.

## Features

-   **Static File Server**: Hosts static files like CSS, JavaScript, and images.
-   **Template Rendering**: Serves HTML pages with dynamic content using Fiber's HTML views engine.
-   **WebSocket Support**: Real-time, bidirectional communication using WebSocket for live updates and messages.
-   **Basic API Endpoint**: `/ping` endpoint for testing server connectivity.

## Project Structure

```plaintext
.
├── main.go             # Main application entry point
├── views/              # HTML template files
├── static/             # Static assets (CSS, JS, images, etc.)
└── handlers/           # Custom request handlers (e.g., for index page)
```

## Prerequisites

-   **Go**: Version 1.16 or higher
-   **Fiber**: Go web framework (imported in the project)
-   **WebSocket**: Fiber's WebSocket library for real-time communication

## Setup

1. **Clone the repository**:

    ```bash
    git clone https://github.com/etharrra/go-chat.git
    cd go-chat
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Run the server**:

    ```bash
    go run main.go
    ```

4. **Access the server**:
    - Open a browser and go to `http://127.0.0.1:8080`.
    - Test WebSocket at `http://127.0.0.1:8080/ws`.
    - Access static files at `http://127.0.0.1:8080/static`.

### Example Request

-   **GET /ping**: Returns `pong` for server testing.

## Contributors

-   [Thar Htoo](https://github.com/etharrra)

Feel free to contribute by forking the repository and submitting a pull request!
