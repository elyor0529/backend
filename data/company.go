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

// Company represents a company
type Company struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
}

// Companies holds all data access functions related to companies
type Companies interface {
	Identify() // Identify records or update a company
}
