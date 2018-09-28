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

package data

import

// sqlite3 db driver
_ "github.com/mattn/go-sqlite3"

// Company represents a company
type Company struct {
	ID int64 `json:"id"`

	Name    string `json:"name"`
	Created int64  `json:"created"`
}

// Companies holds all data access functions related to companies
type Companies struct{}

// Identify records or update a company
func (c *Companies) Identify() {
	db.Exec()
}
