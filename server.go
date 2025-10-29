package main

import (
	"fmt"
	"log"
	"net"
	rpc "net/rpc"
	"sync"
	"time"
)

type Listener int

var (
	messages []string
	mutex    sync.RWMutex
)


type Args struct {
	A int
	B int
}

type Message struct {
	Content string
	Sender  string
	Time    time.Time
}

type ChatHistory struct {
	Messages []string
}

type UserMessage struct {
	UserName string
	Content  string
}


func (l *Listener) GetLine(userMsg *UserMessage, history *ChatHistory) error {
	mutex.Lock()
	formattedMessage := fmt.Sprintf("[%s] %s: %s", time.Now().Format("15:04:05"), userMsg.UserName, userMsg.Content)
	messages = append(messages, formattedMessage)
	fmt.Printf("New message from %s: %s\n", userMsg.UserName, userMsg.Content)
	mutex.Unlock()
	
	mutex.RLock()
	history.Messages = make([]string, len(messages))
	copy(history.Messages, messages)
	mutex.RUnlock()
	
	return nil
}

func (l *Listener) GetHistory(empty bool, history *ChatHistory) error {
	mutex.RLock()
	defer mutex.RUnlock()
	
	history.Messages = make([]string, len(messages))
	copy(history.Messages, messages)
	return nil
}

func (l *Listener) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	fmt.Printf("Arith: %d+%d done on the server\n", args.A, args.B)
	return nil
}

func main() {
	messages = make([]string, 0)
	
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	fmt.Println("Chat server started on port 42586...")
	fmt.Println("Waiting for clients to connect...")
	rpc.Accept(inbound)
}