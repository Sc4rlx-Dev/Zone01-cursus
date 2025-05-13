package server

import (
    "fmt"
    "net"
    "sync"
    "time"
)

const (
    // ANSI escape codes for formatting
    ansiClearLine = "\033[2K"      // Clear the current line
    ansiReset     = "\033[0m"      // Reset text formatting
    ansiGreen     = "\033[32m"     // Green text for user messages
    ansiYellow    = "\033[33m"     // Yellow text for server messages
)

var (
    Clients  = make(map[net.Conn]*Client) // Exported
    Messages []string                     // Exported
    mutex    sync.Mutex
)

type Client struct {
    Conn net.Conn // Exported
    Name string   // Exported
}

func Broadcast(message, senderName string, senderConn net.Conn) {
    mutex.Lock()
    defer mutex.Unlock()

    // If the message is empty (e.g., user pressed "Enter"), clear the line and return
    if message == "" {
        fmt.Fprint(senderConn, "\r"+ansiClearLine)
        return
    }

    var formatted string
    if senderName == "" {
        // Server message with yellow text
        formatted = fmt.Sprintf("%s[%s]: %s%s", ansiYellow, time.Now().Format("2006-01-02 15:04:05"), message, ansiReset)
    } else {
        // User message with green text
        formatted = fmt.Sprintf("[%s]%s[%s]: %s%s", time.Now().Format("2006-01-02 15:04:05"), ansiGreen, senderName, message, ansiReset)
        // Store the formatted message in the Messages slice for later use.
        Messages = append(Messages, formatted)
    }

    // Send the formatted message to all connected clients
    for conn := range Clients {
        // Clear the current input line for all clients
        fmt.Fprint(conn, "\r"+ansiClearLine)

        if conn == senderConn {
            // Replace the sender's input line with the formatted message
            fmt.Fprintf(conn, "%s\n", formatted)
        } else {
            // Send the formatted message to other clients
            fmt.Fprintln(conn, formatted)

            // Restore the input prompt for other clients
            fmt.Fprint(conn, "> ")
        }
    }
}