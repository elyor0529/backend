package main

type commands string

const (
	// CommandRegister for new session
	CommandRegister commands = "reg"
	// CommandListConversation for listing conversations
	CommandListConversation commands = "lc"
)

// Message represents all message sent and received
type Message struct {
	Token             string   `json:"token"`
	ConversationToken string   `json:"conversationId"`
	Command           commands `json:"cmd"`
	Data              string   `json:"data"`
}
