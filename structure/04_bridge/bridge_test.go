package bridge

import "testing"

func TestBridge(t *testing.T) {
    cm := CommonMessage{
		MessageChannel: &MessageEmail{},
	}
	cm.sendMessage("cm", "Bob")

	cm1 := CommonMessage{
		MessageChannel: &MessageSMS{},
	}
	cm1.sendMessage("cm", "Bob")

	um := UrgencyMessage{
		MessageChannel: &MessageSMS{},
	}
	um.sendMessage("um", "Bob")


	um1 := UrgencyMessage{
		MessageChannel: &MessageSMS{},
	}
	um1.sendMessage("um", "Bob")
}
