/*
## License
This project is licensed under the MIT Licence. Refer to https://github.com/mstgnz/go-socket/blob/main/LICENSE for more information.
*/
class Socket {
    constructor() {
        this.players = [];
        this.messages = [];
        this.objSize = 0;
        this.isAnimate = false;
        this.game = document.getElementById("game")
        this.chat = document.getElementById("chat")
        this.form = document.getElementById("form")
        this.input = document.getElementById("input")
        this.connected = document.getElementById("connected")
        this.unconnected = document.getElementById("unconnected")
        this.game.addEventListener("click", this.onClickPlayer)
        this.input.addEventListener("keypress", this.onMessage)
        this.form.addEventListener("submit", (event) => {
            event.preventDefault()
        })
        this.player = {
            name: Math.random().toString(36).replace(/[^a-z]+/g, '').substring(0, 5)
        }
        this.socket = new WebSocket("ws://localhost:3000/ws")
        this.socketProcess()
    }

    socketProcess() {

        // on open
        this.socket.onopen = () => {
            this.send("new")
        }

        // on message
        this.socket.onmessage = (event) => {
            const response = JSON.parse(event.data)
            //console.log(response)
            switch (response.type) {
                case "init":
                    this.handleInit(response)
                    break
                case "new":
                    this.handleNewPlayer(response)
                    break
                case "animate":
                    this.handleAnimate(response)
                    break
                case "message":
                    this.handleMessage(response)
                    break
                case "error":
                    this.handleError(response)
                    break
                case "name":
                    this.handleName(response)
                    break
                case "disconnect":
                    this.handleDisconnect(response)
                    break
            }
        }

        // on close
        this.socket.onclose = () => {
            this.connected.style.display = "none"
            this.unconnected.style.display = "block"
        }

        // on error
        this.socket.onerror = (err) => {
            this.connected.style.display = "none"
            this.unconnected.style.display = "block"
        }
    }

    handleInit(response) {
        this.handlePlayers(response)
        this.handleMessages(response)
    }

    handlePlayers(response){
        this.players = response.players ? response.players : []
        response.players.forEach((player) => {
            if (player.name === this.player.name) {
                this.player = player;
            }
            this.addPlayerToGameArea(player)
        })
    }

    handleMessages(response) {
        this.messages = response.messages ? response.messages : []
        if(response.messages){
            response.messages.forEach(message => {
                this.addMessageToChat(`${message.name}: ${message.message}`)
            })
        }
        this.scrollTop()
    }

    handleNewPlayer(response) {
        if(this.player.name !== response.player.name){
            this.players.push(response.player)
            this.addPlayerToGameArea(response.player)
            this.addMessageToChat(`[SERVER] ${response.player.name} connected`)
            this.scrollTop()
        }
    }

    handleAnimate(response) {
        const player = this.players.find((player) => player.name === response.player.name)
        player.position.x = response.player.position.x;
        player.position.y = response.player.position.y;
        this.animateElement(player)
    }

    handleMessage(response) {
        if(response.player.name !== this.player.name){
            this.messages.push({"name": response.player.name, "message": response.message})
            this.addMessageToChat(`${response.player.name}: ${response.message}`)
            this.showBubble(response.player.name, response.message)
            this.scrollTop()
        }
    }

    handleName(response) {
        if (response.player.name !== this.player.name) {
            const player = this.players.find((player) => player.name === response.player.name)
            if(player){
                this.changeName(player, response.message)
                player.name = response.message
            }
        }
    }

    handleDisconnect(response) {
        this.players = this.players.filter((p) => p.name !== response.player.name);
        this.addMessageToChat(`[SERVER]: ${response.player.name} disconnected`)
        this.scrollTop()
    }

    handleError(response){
        this.connected.style.display = "none"
        this.unconnected.style.display = "block"
        this.unconnected.innerHTML = response.message
    }

    scrollTop() {
        setTimeout(() => {
            this.chat.scrollTop = this.chat.scrollHeight;
        })
    }

    animateElement(player) {
        const element = document.getElementById(player.name)
        if(element){
            const center = this.objSize / 2;
            element.style.left = (player.position.x - center) + "px";
            element.style.top = (player.position.y - center) + "px";
            this.isAnimate = false;
        }
    }

    send(type, message = "") {
        this.socket.send(
            JSON.stringify({
                type: type,
                message: message,
                player: this.player
            })
        )
    }

    onMessage = (event) => {
        if (event.key === "Enter" && this.input.value.length) {
            this.addMessageToChat(`${this.player.name}: ${this.input.value}`)
            this.send("message", this.input.value)
            this.showBubble(this.player.name, this.input.value)
            this.checkCommand(this.input.value)
            this.input.value = "";
        }
    }

    onClickPlayer = (event) => {
        const element = document.getElementById(this.player.name)
        if(element){
            this.objSize = element.offsetWidth;
            const center = this.objSize / 2;
            this.player.position.x = event.offsetX - center;
            this.player.position.y = event.offsetY - center;
            if (!this.isAnimate) {
                this.isAnimate = true;
                this.send("animate","")
                this.animateElement(this.player)
            }
        }
    }

    showBubble(name, message) {
        const element = document.getElementById(name)
        if(element){
            const messageElement = element.querySelector('.message')
            messageElement.style.display = 'block';
            messageElement.innerHTML = message;
            setTimeout(function () {
                messageElement.style.display = 'none';
            }, this.bubbleLifeTime(message))
        }
    }

    bubbleLifeTime(message) {
        const min = 500;
        const max = 3000;
        const msPerLetter = 40;
        let bubbleTime = min + message.length * msPerLetter;
        return bubbleTime > max ? max : bubbleTime;
    }

    checkCommand(command){
        command = command.split(" ")
        if(command[1]) {
            switch (command[0]) {
                case ":change-name":
                    const newName = command[1].substring(0, 5)
                    this.send("name", newName)
                    this.changeName(this.player, newName)
                    this.player.name = command[1];
                    break
                case ":change-bg":
                    const newBg = command[1];
                    document.body.style.backgroundImage = `url(${newBg})`
                    break
            }
        }
        if(command[0] === ":change-bg"){
            fetch("https://source.unsplash.com/random/1920x1080").then((response) => {
                if (response.ok) {
                    document.body.style.backgroundImage = `url(${response.url})`
                }
            })
        }
    }

    changeName(player, name){
        const element = document.getElementById(player.name)
        if(element){
            const nameElement = element.querySelector('.name')
            nameElement.innerHTML = name
        }
    }

    addPlayerToGameArea(player) {
        this.game.innerHTML += `<div class="circle" id="${player.name}" style="left:${player.position.x}px;top:${player.position.y}px; background-color: ${player.color}">
            <div class="relative">
                <span class="name">${player.name}</span>
                <div class="message"></div>
            </div>
        </div>`;
    }

    addMessageToChat(message) {
        this.chat.innerHTML += `<div class="item">
            <div class="content">
                <span>${message}</span>
            </div>
        </div>`;
    }
}

new Socket()