package gchat_wh

import "github.com/go-vs/gchat-wh/messages"

func Message(txt string) *messages.Message {
	return &messages.Message{
		Text: txt,
	}
}

func Reply(thread, txt string) *messages.Message {
	return &messages.Message{
		Text: txt,
		Thread: &messages.Thread{
			ThreadKey: thread,
		},
	}
}
