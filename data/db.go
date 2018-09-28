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

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// this is the sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// db is a global connection pool to the database
var db *sql.DB

func init() {
	dbName := os.Getenv("PARLE_DB")
	if len(dbName) == 0 {
		dbName = "db/dev.db"
	}
	d, err := sql.Open("sqlite3", fmt.Sprintf("%s?_fk=1", dbName))
	if err != nil {
		log.Fatal("unable to open database:", err)
	}

	if err := d.Ping(); err != nil {
		log.Fatal("unable to ping database:", err)
	}
	db = d
}

// Close makes sure the database is closed
func Close() {
	db.Close()
}
