package pools

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Setup configures the db along with pools
// number of connections and more
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/gocookbook?parseTime=true", os.Getenv("MYSQLUSERNAME"), os.Getenv("MYSQLPASSWORD")))
	if err != nil {
		return nil, err
	}

	// there will only ever be 24 open connections
	db.SetMaxOpenConns(25)

	// MaxIdleConns can never be less than max open SetMaxOpenConns
	// otherwise it'll default to that value
	db.SetMaxIdleConns(25)

	return db, nil
}
