/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 * You may use, distribute and modify this code under the
 * terms of the GNU Affero General Public license, as
 * published by the Free Software Foundation, either version
 * 3 of theLicense, or (at your option) any later version.
 *
 * You should have received a copy of the GNU Affero General
 * Public License along with this code as LICENSE file.  If not,
 * see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
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

	return conn, result.Data, err
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

	agentToken := os.Getenv("PARLE_AGENT_TOKEN")
	json := Message{
		Token:   token,
		Command: CommandAuth,
		Data:    agentToken,
	}

	err = conn.WriteJSON(json)
	assert.NoError(t, err)

	var result Message
	err = conn.ReadJSON(&result)
	assert.NoError(t, err)
	assert.NotEqual(t, "OK: "+agentToken, result.Data)
	assert.Contains(t, result.Data, "OK")
}

func TestHubClientListConversation(t *testing.T) {
	user, token, err := newConn()
	assert.NoError(t, err)

	json := Message{
		Token:   token,
		Command: CommandListConversation,
		Data:    "",
	}

	err = user.WriteJSON(json)
	assert.NoError(t, err)

	var result Message
	err = user.ReadJSON(&result)
	assert.NoError(t, err)
	assert.Equal(t, "[]", result.Data)
}

func TestHubClientNewConversation(t *testing.T) {
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
