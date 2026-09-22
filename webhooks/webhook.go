package webhook

import "net/http"

type MessageEvent struct {
	Event     string      `json:"event"`
	Timestamp string      `json:"timestamp"`
	Data      MessageData `json:"data"`
}

type MessageData struct {
	MessageID      string `json:"messageId"`
	ConversationID string `json:"conversationId"`
	SenderID       string `json:"senderId"`
	ReceiverID     string `json:"receiverId"`
	Content        string `json:"content"`
}

func messageWebHook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var event MessageEvent

}
