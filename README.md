# WebSocket Chat Application with Game Environment

This application is a browser-based multiplayer chat and gaming platform. Each user has a game area that covers the entire browser screen. Every user is represented by a 75x75 pixel circle with a randomly generated name inside it. Users can move their characters by clicking anywhere on the screen.

## Getting Started

Follow the steps below to start the project locally or on a web server.

### Prerequisites

- Install the Go programming language locally.
- Make sure your browser supports WebSocket.

### Installation

1. Clone this repository to a local directory:
    ```bash
    git clone https://github.com/mstgnz/go-socket
    cd go-socket
    ```

2. Run project
    ```bash
    docker build -t go-socket:latest . && docker run -d --restart=always -p 3000:3000 --name=go-socket go-socket
    ```
   OR
    ```bash
    docker compose up -d
    ```

3. Go to browser open the link http://localhost:3000

### Usage
Launch the application in your browser and join with a random username.
Click anywhere on the game area to move your character in that direction.
Type a message in the text box at the bottom and press "Enter" to send a message.

### Contributing
This project is open-source, and contributions are welcome. Feel free to contribute or provide feedback of any kind.

### License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/mstgnz/go-socket/blob/main/LICENSE) file for more details.