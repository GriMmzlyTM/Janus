package shell

import (
	"Janus/sqlconnect"
	"github.com/fatih/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func ConnectMysql(c *ishell.Context) {

	c.Print("\n- Provide master string -\n<database_username>:<database_password>@tcp(<database_ip>)/<database_name> : ")

	conn := janussql.DatabaseObject{
		Driver: 	"mysql",
		MasterString: c.ReadLine(),
	}

	dbTemp, err := janussql.Connect(conn)
	if err != nil {
		c.Println(color.RedString("Could not connect to database: " + color.RedString(err.Error())))
		return
	}

	c.Println(color.GreenString("Successfully connected to database!"))

	Db = dbTemp

}
