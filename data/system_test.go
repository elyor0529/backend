/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 * You may use, distribute and modify this code under the
 * terms of the GNU Affero General Public license, as
 * published by the Free Software Foundation, either version
 * 3 of the License, or (at your option) any later version.
 *
 * You should have received a copy of the GNU Affero General
 * Public License along with this code as LICENSE file.  If not,
 * see <http://www.gnu.org/licenses/>.
 */
package data

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Open(":memory:")
	defer Close()

	sys := &System{}
	isInit, err := sys.IsInitialized()
	if err != nil {
		log.Fatal("error testing if db is initialized: ", err)
	} else if isInit {
		log.Fatal("database should not be initialized")
	}

	if err := sys.Create("../migration", "unittest"); err != nil {
		log.Fatal("error while creating initial database schema: ", err)
	}

	res := m.Run()
	os.Exit(res)
}

func TestIsInitializedAfterCreation(t *testing.T) {
	sys := &System{}
	isInit, err := sys.IsInitialized()
	if err != nil {
		t.Fatal(err)
	} else if isInit == false {
		t.Error("database responded as not initialized after initial creation")
	}
}
