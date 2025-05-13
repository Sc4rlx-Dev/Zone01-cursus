package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	// "server"
	"net-cat/server"
)
const defaultPort = "8989"
const maxClients = 10

var (
	mutex sync.Mutex
)

func main() {
	// Get the port to listen on.
	port, err := getPort()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen on the specified port.
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer listener.Close()

	// Print the port number.
	fmt.Println("Listening on port:", port)

	// Run an infinite loop to handle connections.
	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept:", err)
			continue
		}

		// Handle the connection in a separate goroutine.
		go handleConnection(conn)
	}
}

// getPort retrieves the port number from the command-line arguments.
// If no port is provided, it returns the default port.
// If an invalid port is provided, it returns an error.
func getPort() (string, error) {
	args := os.Args[1:]

	// No port provided, return the default port.
	if len(args) == 0 {
		return defaultPort, nil
	}
	// One argument provided, validate if it's a valid integer port.
	if len(args) == 1 {
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return "", fmt.Errorf("[USAGE]: ./server $port")
		}
		return args[0], nil
	}

	// Invalid number of arguments, return an error.
	return "", fmt.Errorf("[USAGE]: ./server $port")
}

// handleConnection manages a single client connection to the chat server.
func handleConnection(conn net.Conn) {
    defer conn.Close()

    conn.Write([]byte("[WELCOME TO THE CHAT SERVER]\n"))

    Name, err := readName(conn)
    if err != nil {
        return
    }

    mutex.Lock()
    if len(server.Clients) >= maxClients {
        conn.Write([]byte("Server full. Try again later.\n"))
        mutex.Unlock()
        return
    }
    mutex.Unlock()

    client := &server.Client{Conn: conn, Name: Name}

    mutex.Lock()
    server.Clients[conn] = client
    mutex.Unlock()

    server.Broadcast(fmt.Sprintf("%s has joined our chat...", client.Name), "", conn)

    sendHistory(conn)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := strings.TrimSpace(scanner.Text())
        if text == "" {
            continue
        }

        if !isPrintableASCII(text) {
            conn.Write([]byte("[MESSAGE MUST CONTAIN ONLY PRINTABLE ASCII CHARACTERS]\n"))
            continue
        }

        // Clear the input line for the sender
        conn.Write([]byte("\033[2K\r"))

        // Broadcast the received message to all Clients, including the sender.
        server.Broadcast(text, client.Name, conn)
    }

    mutex.Lock()
    delete(server.Clients, conn)
    mutex.Unlock()

    server.Broadcast(fmt.Sprintf("%s has left our chat...", client.Name), "", conn)
}


func isPrintableASCII(s string) bool {
	for _, r := range s {
		if r < 32 || r > 127 {
			return false
		}
	}
	return true
}

func readName(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)

	for {
		conn.Write([]byte("[ENTER YOUR Name]: "))
		Name, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		Name = strings.TrimSpace(Name)

		if Name == "" {
			conn.Write([]byte("[Name CANNOT BE EMPTY]\n"))
			continue
		}

		if !isPrintableASCII(Name) {
			conn.Write([]byte("[Name MUST CONTAIN ONLY PRINTABLE ASCII CHARACTERS]\n"))
			continue
		}

		// Check if the Name is already taken
		mutex.Lock()
		NameTaken := false
		for _, client := range server.Clients {
			if client.Name == Name {
				NameTaken = true
				break
			}
		}
		mutex.Unlock()

		if NameTaken {
			conn.Write([]byte("[Name ALREADY TAKEN]\n"))
			continue
		}

		return Name, nil
	}
}

// sendHistory sends the chat history to a newly connected client.
func sendHistory(conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	// Iterate through all the stored messages and send them to the newly connected client.
	for _, msg := range server.Messages {
		fmt.Fprintln(conn, msg)
	}
}