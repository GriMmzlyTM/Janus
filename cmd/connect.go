// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Janus/shell"
	janussql "Janus/sqlconnect"
	"fmt"

	"github.com/spf13/cobra"
)

var dbDriver string
var dbUser string
var dbPass string
var dbName string
var dbHost string
var dbString string

type pet struct {
	name    string
	owner   string
	species string
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect Janus to your database",
	Long: `Allows you to connect Janus to your database. You must either provide individual data through flags (As well as driver)
or provide a master string (--to) as well as the driver. Type Janus connect --help for more info.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn := janussql.DatabaseObject{
			Driver:       dbDriver,
			Host:         dbHost,
			User:         dbUser,
			Password:     dbPass,
			Dbname:       dbName,
			MasterString: dbString}
		dbTemp, err := janussql.Connect(conn)
		if err != nil {
			fmt.Println("COULD NOT CONNECT TO DATABASE. " + err.Error())
			return
		}

		shell.Db = dbTemp

		shell.Start()
	},
}

func init() {

	connectCmd.Flags().StringVarP(&dbDriver, "driver", "d", "", "Database driver to use when connecting")
	connectCmd.Flags().StringVarP(&dbUser, "user", "u", "", "Database username to use when connecting | sqlconnect or postgres")
	connectCmd.Flags().StringVarP(&dbPass, "password", "p", "", "Database password to use when connecting")
	connectCmd.Flags().StringVarP(&dbName, "database", "D", "", "Database to connect to")
	connectCmd.Flags().StringVarP(&dbHost, "host", "H", "", "host to connect to")
	connectCmd.Flags().StringVarP(&dbString, "to", "t", "", "Full database string to use for connection. formatted "+
		"<database_username>:<database_password>@tcp(<database_ip>)/<database_name>")
	connectCmd.MarkFlagRequired("driver")

	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
