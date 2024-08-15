package gchat_wh

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestResponse(t *testing.T) {
	data := `{
  "name": "spaces/[space]/messages/[message]",
  "sender": {
    "name": "users/[user]",
    "displayName": "[display_name]",
    "type": "BOT"
  },
  "createTime": "2024-06-23T15:21:43.152415Z",
  "text": "Hello, World!",
  "thread": {
    "name": "spaces/[space]/threads/[thread]",
    "retentionSettings": {
      "state": "PERMANENT"
    }
  },
  "space": {
    "name": "spaces/[space]",
    "type": "ROOM",
    "displayName": "[display_name]",
    "spaceThreadingState": "THREADED_MESSAGES",
    "spaceType": "SPACE",
    "spaceHistoryState": "HISTORY_ON",
    "createTime": "2024-05-24T03:15:33.655646Z"
  },
  "argumentText": "Hello, World!",
  "retentionSettings": {
    "state": "PERMANENT"
  },
  "formattedText": "Hello, World!"
}
`
	var resp ErrorResponse
	if err := json.NewDecoder(strings.NewReader(data)).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
