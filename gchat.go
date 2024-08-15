package gchat_wh

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-vs/gchat-wh/messages"
	"net/http"
)

type GChat struct {
	config Config
	client *http.Client
}

type Config struct {
	SpaceID string
	Key     string
	Token   string
}

type Ops func(g *GChat)

func WithClient(c *http.Client) Ops {
	return func(g *GChat) {
		g.client = c
	}
}

func New(config Config, ops ...Ops) *GChat {
	g := &GChat{
		config: config,
		client: &http.Client{},
	}
	for _, op := range ops {
		op(g)
	}
	return g
}

func (g *GChat) Send(msg *messages.Message) (*messages.Message, error) {
	reqBody, err := jsonStream(msg)
	if err != nil {
		return nil, err
	}
	args := make(map[string]string)
	if msg.Thread != nil && msg.Thread.ThreadKey != "" {
		args["messageReplyOption"] = "REPLY_MESSAGE_FALLBACK_TO_NEW_THREAD"
	}
	url := buildURL(g.config.SpaceID, g.config.Key, g.config.Token, args)
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// try 400
		var respBody ErrorResponse
		if resp.StatusCode == http.StatusBadRequest {
			if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
				return nil, err
			}
			return nil, errors.New(respBody.Error.Message)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var respBody messages.Message
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
