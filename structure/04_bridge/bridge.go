package bridge

import "fmt"

type MessageChannel interface {
    send(message string, user string) 	
}

type MessageSMS struct {}
func (sms *MessageSMS) send(msg string, user string) {
	fmt.Println(msg, "to", user, "sms")
}

type MessageEmail struct {}
func (email *MessageEmail) send(msg string, user string) {
	fmt.Println(msg, "to", user, "email")
}

type Message struct {
	messageChannel MessageChannel
}

func (m *Message) sendMessage(msg string, user string) {
	m.messageChannel.send(msg, user)
}

type CommonMessage struct {
	MessageChannel MessageChannel
}

func (cm *CommonMessage) sendMessage(msg string, user string) {
	cm.MessageChannel.send(msg, user)
}

type UrgencyMessage struct {
	MessageChannel MessageChannel
}

func (um *UrgencyMessage) sendMessage(msg string, user string) {
	msg = fmt.Sprintf("urgency %s", msg)
	um.MessageChannel.send(msg, user)
}