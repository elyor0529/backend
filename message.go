package main

type commands string

const (
	// CommandAuth used for users and agents to authenticate
	CommandAuth commands = "auth"
	// CommandListConversation for listing conversations
	CommandListConversation commands = "lc"
	// CommandNewConversation when creating a new conversation
	CommandNewConversation commands = "nc"
)

// Message represents all message sent and received via websocket
type Message struct {
	Token             string   `json:"token"`
	ConversationToken string   `json:"conversationId"`
	Command           commands `json:"cmd"`
	Data              string   `json:"data"`
}
