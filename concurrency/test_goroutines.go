package main

import "fmt"
import "strings"

func main() {
	var testMessage = &Message{content: "This is a test message"}

	processMessage(testMessage, getMessageFormatter(3), nil)

	// create channel
	ch := make(chan string)

	go processMessage(testMessage, getMessageFormatter(1), ch)
	go processMessage(testMessage, getMessageFormatter(2), ch)
	go processMessage(testMessage, getMessageFormatter(3), ch)

	// wait for goroutines to finish and get results
	m1, m2, m3 := <-ch, <-ch, <-ch

	// output results
	fmt.Println(">>> Processed messages:")
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}

type Message struct {
	content string
}

type ParseableMessage interface {
	parseMessage(message string, formatter *func(data string) string) string
}

func (message *Message) parseMessage(formatter *func(data string) string) string {
	return (*formatter)(message.content)
}

// message parsers
var toLowerCase = func(data string) string {
	return strings.ToLower(data)
}

var toUpperCase = func(data string) string {
	return strings.ToUpper(data)
}

var defaultFormatter = func(data string) string {
	return "Returning original message content: " + data
}

// parser selector
func getMessageFormatter(messageType int8) *func(data string) string {
	switch messageType {
	case 1:
		return &toLowerCase
	case 2:
		return &toUpperCase
	default:
		return &defaultFormatter
	}
}

// go routine action
func processMessage(message *Message, formatter *func(data string) string, c chan string) {
	if c != nil {
		c <- "Processed message: " + message.parseMessage(formatter)
	}
}
