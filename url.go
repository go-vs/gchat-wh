package gchat_wh

import "fmt"

func buildURL(spaceID, key, token string, arg map[string]string) string {
	url := fmt.Sprintf("https://chat.googleapis.com/v1/spaces/%v/messages?key=%v&token=%v", spaceID, key, token)
	for k, v := range arg {
		url += fmt.Sprintf("&%v=%v", k, v)
	}
	return url
}
