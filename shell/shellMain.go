package shell

import (
	"Janus/sqlconnect"
	"database/sql"
	"github.com/fatih/color"
	"gopkg.in/abiosoft/ishell.v2"
)

var Db *sql.DB

func Start() {
	shell := ishell.New()

	// display welcome info.
	shell.Println(`Welcome to Janus! 
Janus is an all-in-one CLI-based SQL helper. type 'intro' to see what Janus can do!`)

	shell.AddCmd(&ishell.Cmd{
		Name: "query:run",
		Help: "Query the database",
		Func: func(c *ishell.Context) { Query(c) } ,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "query:run-interval",
		Help: "Query the database at set intervals",
		Func: func(c *ishell.Context) { QueryInterval(c) } ,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "query:set",
		Help: "Set the query json file to use (Including location)",
		Func: func(c *ishell.Context) { SetFile(c) } ,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "query:list",
		Help: "List all available queries in JSON file (If using custom json file, make sure to set it first with query:set)",
		Func: func(c *ishell.Context) { ListQueries(c) } ,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "connect:ping",
		Help: "Pings the database to ensure you're still connected",
		Func: func(c *ishell.Context) {
			err := janussql.PingDb(c, Db)
			if err != nil { c.Println(color.RedString("Could not ping database: ") + color.RedString(err.Error())) }
			},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "connect:mysql",
		Help: "Connects you to a mysql database",
		Func: func(c *ishell.Context) { ConnectMysql(c) } ,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "intro",
		Help: "Explanation of what Janus can do for you.",
		Func: func(c *ishell.Context) { intro(c) } ,
	})

	// run shell
	shell.SetPrompt("Janus> ")
	shell.Run()
}

func intro(c *ishell.Context) {
	c.Println(`
Janus is an all-in-one CLI based SQL helper! It allows you to run pre-defined SQL queries from a json file and more! 
Janus was made for hard to modify legacy apps, and people in more managerial positions with less coding prowess. 
Essentially, Janus allows anyone on your team to access data from your database without needing to write queries!

Janus is also capable of starting and stopping intervaled SQL queries. Essentially cron jobs that are run and 
controlled by Janus. This is especially useful for volatile legacy software that's hard to modify. 
A good example are legacy Java applications that generate SQl queries through a series of complex method calls. 
I've worked on legacy Java apps in the past where plopping in an SQL string was a no-go. 

Instead, you can use Janus to run the SQl queries.

If you're using a modern API where SQl queries and cron jobs aren't a huge issue, Janus probably isn't for you. 
Especially if the non-coders in your team don't need constant rapid access to specific data from your prod DB at will.

Janus was written entirely in GO by Lorenzo Torelli. Feel free to use and modify at will.
`)
}
