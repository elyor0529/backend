/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/parle-io/backend/data/agent"
	"github.com/parle-io/backend/data/user"
	"github.com/stretchr/testify/assert"
)

func newConn() (*websocket.Conn, string, error) {
	dialer := websocket.Dialer{}

	conn, _, err := dialer.Dial(ws.url, nil)
	if err != nil {
		return nil, "", err
	}

	var result Message
	err = conn.ReadJSON(&result)
	if err != nil {
		return nil, "", err
	}

	if result.Command != CommandAuth {
		return nil, "", fmt.Errorf("expected command to be CommandAuthentication but got %v", result.Command)
	}

	token := strings.Replace(result.Data, "OK ", "", -1)
	return conn, token, err
}

func TestHubRegister(t *testing.T) {
	conn, token, err := newConn()
	assert.NoError(t, err)

	json := Message{
		Token:   token,
		Command: CommandListConversation,
		Data:    "",
	}
	err = conn.WriteJSON(json)
	assert.NoError(t, err)

	var result Message
	err = conn.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Equal(t, "[]", result.Data)
}

func TestHubAuthentication(t *testing.T) {
	conn, token, err := newConn()
	assert.NoError(t, err)

	msg := Message{
		Token:   token,
		Command: CommandAuth,
		Data:    fmt.Sprintf("%s|%s", ws.admin.Email, ws.admin.Token),
	}

	err = conn.WriteJSON(msg)
	assert.NoError(t, err)

	var result Message
	err = conn.ReadJSON(&result)
	assert.NoError(t, err)
	assert.NotContains(t, result.Data, "err:")

	var check agent.Agent
	err = json.Unmarshal([]byte(result.Data), &check)
	assert.NoError(t, err)
	assert.Equal(t, ws.admin.Email, check.Email)
}

func TestHubClientIdentification(t *testing.T) {
	visitor, token, err := newConn()
	assert.NoError(t, err)

	subMsg := user.User{}
	subMsg.BrowserAgent = "unit test browser"
	subMsg.Trackings = append(subMsg.Trackings, user.Tracking{
		Referrer: "google.com/abc",
	})

	buf, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	msg := Message{
		Token:   token,
		Command: CommandIdentify,
		Data:    string(buf),
	}

	err = visitor.WriteJSON(msg)
	assert.NoError(t, err)

	var result Message
	err = visitor.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Contains(t, result.Data, token)

	// let's try to simulate another connection
	// passing this token

	visit2, tmpToken, err := newConn()
	assert.NoError(t, err)

	msg.Token = tmpToken
	subMsg.ConnectionToken = token

	buf, err = json.Marshal(subMsg)
	assert.NoError(t, err)

	msg.Data = string(buf)

	err = visit2.WriteJSON(msg)
	assert.NoError(t, err)

	var result2 Message
	err = visit2.ReadJSON(&result2)
	assert.NoError(t, err)
	assert.Contains(t, result2.Data, token)
}

func TestHubClientListConversation(t *testing.T) {
	user, token, err := newConn()
	assert.NoError(t, err)

	msg := Message{
		Token:   token,
		Command: CommandListConversation,
		Data:    "",
	}

	err = user.WriteJSON(msg)
	assert.NoError(t, err)

	var result Message
	err = user.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Equal(t, "[]", result.Data)
}

func TestHubClientNewConversation(t *testing.T) {
	t.Skip()

	user, token, err := newConn()
	assert.NoError(t, err)

	json := Message{
		Token:   token,
		Command: CommandNewConversation,
		Data:    "Say hello to unit test!",
	}
	err = user.WriteJSON(json)
	assert.NoError(t, err)

	var result Message
	err = user.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Equal(t, "OK", result.Data)
}

func TestHubClientSendReceiveMessage(t *testing.T) {
	t.Skip()

	user, utoken, err := newConn()
	assert.NoError(t, err)

	agent, atoken, err := newConn()
	assert.NoError(t, err)

	msg := Message{
		Command: CommandAuth,
		Token:   atoken,
		Data:    os.Getenv("PARLE_AGENT_TOKEN"),
	}

	err = agent.WriteJSON(msg)
	assert.NoError(t, err)

	var result Message
	err = agent.ReadJSON(&result)
	assert.NoError(t, err)

	atoken = strings.Replace(result.Data, "OK ", "", -1)

	// new message from user
	msg.Token = utoken
	msg.Command = CommandNewConversation
	msg.Data = "Hello from unit test"

	err = user.WriteJSON(msg)
	assert.NoError(t, err)

	err = user.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Contains(t, result.Data, "OK")

	err = agent.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Equal(t, CommandNewConversation, result.Command)
	assert.Contains(t, result.Data, "OK")
}
