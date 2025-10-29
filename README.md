# 💬 Simple Chatroom RPC System

https://github.com/user-attachments/assets/4ad46cdb-31a9-48d8-a375-567a5f6dc746


A **distributed chatroom system** built using **Go's Remote Procedure Call (RPC)** technology.  
This system allows multiple users to connect to a central server and exchange messages in real-time.

---

## 📋 General Idea

The chatroom consists of a central **server** and multiple **clients** that communicate using RPC.

- Users can send and receive messages instantly.  
- The server stores all messages with timestamps and usernames.  
- Clients can view chat history and see new messages as they arrive.

---

## 🏗️ Architecture

### 🖥️ Server — The Central Message Hub
- Stores all chat messages  
- Manages client connections  
- Broadcasts messages to all connected users  

### 💻 Client — User Interface Application
- Connects to the server  
- Sends and receives messages  
- Displays chat history  

---

## 🔄 How It Works

1. **Server starts first** — Listens for incoming connections on port `42586`  
2. **Client connects** — User enters their name and connects to the server  
3. **Send messages** — User types messages that are sent to the server  
4. **Server stores messages** — Messages saved with timestamps and usernames  
5. **View history** — Users can see all previous messages  
6. **Multiple clients** — Many users can connect simultaneously  

---

## 💡 Key Features

✅ **User Identification:** Each message shows who sent it  
🕒 **Timestamps:** All messages include the time sent  
📜 **Chat History:** Complete message history available  
⚡ **Real-time Updates:** See new messages instantly  
🔁 **Error Handling:** Automatic reconnection if server disconnects  

### 🧭 Simple Commands


---

## 🛠️ Technical Implementation

### 🧩 RPC Methods Used
- `GetLine()` — Send a new message to the server  
- `GetHistory()` — Retrieve all chat messages  

### 📦 Data Structures
- **UserMessage** — Contains username + message content  
- **ChatHistory** — Stores list of all formatted messages  

---

## 🚀 How to Run

### 1 Start the Server
```bash
go run server.go
```


### 2 Start the Client
```bash
go run client.go
```

### 3 Enter your username and start chatting! 
