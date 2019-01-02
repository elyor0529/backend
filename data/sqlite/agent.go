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
	"database/sql"

	"github.com/parle-io/backend/data"
)

type Agents struct{}

func (a *Agents) GetByEmail(email string) (*data.Agent, error) {
	row := db.QueryRow("SELECT * FROM agents WHERE Email=?", email)

	agent := &data.Agent{}
	if err := a.read(row, agent); err != nil {
		return nil, err
	}
	return agent, nil
}

func (a *Agents) read(row *sql.Row, v *data.Agent) error {
	return row.Scan(&v.ID, &v.Email, &v.Password, &v.Created)
}
