package janussql

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	ishell "gopkg.in/abiosoft/ishell.v2"
)

type DatabaseObject struct {
	Driver       string
	Host         string
	User         string
	Password     string
	Dbname       string
	MasterString string
}

func Connect(conn DatabaseObject) (*sql.DB, error) {

	dbTemp, err := connectDatabase(conn)

	if err != nil {
		return nil, err
	}

	pingErr := dbTemp.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return dbTemp, nil

}

func PingDb(c *ishell.Context, db *sql.DB) error {

	if db == nil {
		return fmt.Errorf("Database not initialized or set. Please connect first.")
	}

	err := db.Ping()
	if err != nil {
		return err
	}

	c.Println(color.BlueString("Ping successful! You're connected to the database!"))

	return nil
}

func connectDatabase(dbObj DatabaseObject) (*sql.DB, error) {

	var dataSource string

	if dbObj.MasterString != "" {
		dataSource = dbObj.MasterString
	} else {
		dataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbObj.User, dbObj.Password, dbObj.Host, dbObj.Dbname)
	}

	db, err := sql.Open(dbObj.Driver, dataSource)
	return db, err
}
