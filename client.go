package main

import (
	"bufio"
	"fmt"
	"log"
	rpc "net/rpc"
	"os"
	"strings"
)

// Common data structures
type UserMessage struct {
	UserName string
	Content  string
}

type ChatHistory struct {
	Messages []string
}

type Args struct {
	A int
	B int
}

func main() {
	client, err := rpc.Dial("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	if !scanner.Scan() {
		log.Fatal("Failed to read name")
	}
	userName := strings.TrimSpace(scanner.Text())
	if userName == "" {
		userName = "Anonymous"
	}

	fmt.Printf("\nWelcome %s to the Chatroom!\n", userName)
	fmt.Println("Type 'exit' to quit, 'history' to see chat history")
	fmt.Println("==================================================")

	for {
		fmt.Print("Enter message: ")
		if !scanner.Scan() {
			break
		}

		message := strings.TrimSpace(scanner.Text())
		
		if strings.ToLower(message) == "exit" {
			fmt.Println("Goodbye!", userName)
			break
		}
		
		if strings.ToLower(message) == "history" {
			var history ChatHistory
			err := client.Call("Listener.GetHistory", true, &history)
			if err != nil {
				log.Println("Error fetching history:", err)
				continue
			}
			
			fmt.Println("\n=== Chat History ===")
			if len(history.Messages) == 0 {
				fmt.Println("No messages yet")
			} else {
				for i, msg := range history.Messages {
					fmt.Printf("%d. %s\n", i+1, msg)
				}
			}
			fmt.Println("====================\n")
			continue
		}

		if message == "" {
			continue
		}

		userMsg := &UserMessage{
			UserName: userName,
			Content:  message,
		}
		var history ChatHistory
		
		err = client.Call("Listener.GetLine", userMsg, &history)
		if err != nil {
			log.Println("Error sending message:", err)
			// Try to reconnect
			client, err = rpc.Dial("tcp", "0.0.0.0:42586")
			if err != nil {
				log.Fatal("Failed to reconnect to server:", err)
			}
			continue
		}

		fmt.Println("\n=== Updated Chat ===")
		if len(history.Messages) == 0 {
			fmt.Println("No messages yet")
		} else {
			for i, msg := range history.Messages {
				fmt.Printf("%d. %s\n", i+1, msg)
			}
		}
		fmt.Println("====================\n")
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
}