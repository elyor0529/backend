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

package sqlite

import (
	"testing"
)

func TestAgentCanLogin(t *testing.T) {
	agents := &Agents{}
	me, err := agents.GetByEmail("admin@changeme.com")
	if err != nil {
		t.Fatal(err)
	} else if me.ID == 0 || me.Email != "admin@changeme.com" {
		t.Errorf("expected ID to be > 0 got %d and email to be admin@changeme.com got %s", me.ID, me.Email)
	}
}
