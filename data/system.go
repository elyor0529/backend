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
	// sqlite3 db driver
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	// sqlite3 database driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	// Version represents the current version of the data package and database schema
	Version = "0.0.1"
)

// System handles everything related to database creation and schema migration
type System struct{}

// IsInitialized determines if it's the first app start
func (s *System) IsInitialized() (bool, error) {
	row := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='settings';`)
	if row == nil {
		return false, nil
	}

	var v string
	if err := row.Scan(&v); err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

// Create initializes the database and insert the app version in the settings table
func (s *System) Create(migrationPath, passwd string) error {
	b, err := ioutil.ReadFile(path.Join(migrationPath, "_init.sql"))
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(b)); err != nil {
		return err
	}

	res, err := db.Exec("INSERT INTO settings VALUES('version', ?)", Version)
	if err != nil {
		return err
	} else if ra, err := res.RowsAffected(); err != nil {
		return err
	} else if ra != 1 {
		return fmt.Errorf("row affected %d after executing the initial version insert", ra)
	}

	// create the first user
	res, err = db.Exec("INSERT INTO agents(email, password, created) VALUES('admin@changeme.com', ?, ?)", passwd, time.Now())
	if err != nil {
		return err
	} else if ra, err := res.RowsAffected(); err != nil {
		return err
	} else if ra != 1 {
		return fmt.Errorf("unable to add first agent after first initialization")
	}

	return nil
}
