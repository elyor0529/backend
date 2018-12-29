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
	"database/sql"
	"time"
)

type Agent struct {
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:-`
	Created  time.Time `json:"created"`
}

type Agents struct{}

func (a *Agents) GetByEmail(email string) (*Agent, error) {
	row := db.QueryRow("SELECT * FROM agents WHERE Email=?", email)

	agent := &Agent{}
	if err := a.read(row, agent); err != nil {
		return nil, err
	}
	return agent, nil
}

func (a *Agents) read(row *sql.Row, v *Agent) error {
	return row.Scan(&v.ID, &v.Email, &v.Password, &v.Created)
}
