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

	assert.NoError(t, err)

	json := Message{
		Token:   "unittest",
		Command: CommandRegister,
		Data:    "",
	}
	err = conn.WriteJSON(json)
	assert.NoError(t, err)

	var data Message
	err = conn.ReadJSON(&data)
	assert.NoError(t, err)
	assert.Equal(t, json, data)
}

func TestHubClientListConversation(t *testing.T) {
	user, err := newConn()
	assert.NoError(t, err)

	json := Message{
		Token:   "unittest",
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
