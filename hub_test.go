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
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func newConn() (*websocket.Conn, error) {
	dialer := websocket.Dialer{}

	conn, _, err := dialer.Dial(ws.url, nil)
	return conn, err
}

func TestHubRegister(t *testing.T) {
	conn, err := newConn()
	defer conn.Close()
	assert.NoError(t, err)

	err = conn.WriteMessage(websocket.TextMessage, []byte("test"))
	assert.NoError(t, err)

	typ, data, err := conn.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, websocket.TextMessage, typ)
	assert.Equal(t, []byte("test"), data)
}

func TestHubClientSendMessage(t *testing.T) {
	user, err := newConn()
	defer user.Close()
	assert.NoError(t, err)

	agent, err := newConn()
	defer agent.Close()
	assert.NoError(t, err)

	user.WriteMessage(websocket.TextMessage, []byte("hello"))
	typ, data, err := agent.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, websocket.TextMessage, typ)
	assert.Equal(t, []byte("hello"), data)
}
