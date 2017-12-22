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
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type wstestdata struct {
	svr *httptest.Server
	hub *Hub
	url string
}

var ws *wstestdata

func TestMain(m *testing.M) {
	hub := newHub()
	go hub.run()

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	}))
	defer svr.Close()

	ws = &wstestdata{
		svr: svr,
		hub: hub,
		url: strings.Replace(svr.URL, "http", "ws", -1),
	}

	ret := m.Run()
	os.Exit(ret)
}
