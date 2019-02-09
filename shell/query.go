// Copyright © 2018 Lorenzo TorelliLorenzo@tortonmind.com
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

package shell

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"

	ishell "gopkg.in/abiosoft/ishell.v2"
)

var file = "queries.json"

// queryCmd represents the query command
func Query(c *ishell.Context) {

	if Db == nil {
		c.Println(color.RedString("\nCannot execute a query when database has not been initialized. Please use 'connect' first.\n"))
		return
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		c.Println(color.RedString("Could not connect to database. Make sure you connected through the 'connect' command when starting, or through the cli."))
		c.Println(pingErr.Error())
		return
	}

	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	c.Print(color.GreenString("Query name: "))
	query := c.ReadLine()

	queryMap := checkQuery(c, query)

	if queryMap == nil {
		return
	}

	c.Println(color.GreenString("\nEXECUTING QUERY..."))

	rows, err := Db.Query(queryMap[query]["query"])
	if err != nil {
		c.Println(err.Error())
		return
	}
	defer rows.Close()

	data := make([]map[string]string, 0)

	processErr := processRows(rows, &data, c)
	if processErr != nil {
		return
	}

	for i, _ := range data {
		for key, val := range data[i] {
			c.Printf(color.MagentaString("%s ")+"-"+color.BlueString(" %s")+" | ", key, val)
		}
		c.Println("")
	}
	c.Println("")
}

func QueryInterval(c *ishell.Context) {

	c.Print("Query to run: ")
	query := c.ReadLine()

	queryMap := checkQuery(c, query)

	if queryMap == nil {
		return
	}

}

func checkQuery(c *ishell.Context, query string) map[string]map[string]string {

	queryMap := ParseJson(file, c)
	if queryMap == nil {
		return nil
	}

	if _, ok := queryMap[query]; !ok {
		c.Printf(color.RedString("\n%s does not exist in json map!\n"), query)
		c.Print("Would you like to see available queries? (y/n) : ")

		if c.ReadLine() == "y" {
			ListQueries(c)
		}
		return nil
	}

	return queryMap

}

func processRows(rows *sql.Rows, data *[]map[string]string, c *ishell.Context) error {

	cols, err := rows.Columns()
	if err != nil {
		c.Println(err.Error())
		return nil
	}

	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))

		rowData := make(map[string]string)

		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		err := rows.Scan(columnPointers...)
		if err != nil {
			c.Println(color.RedString("Could not scan rows - " + color.RedString(err.Error())))
			return err
		}

		for i, colName := range cols {
			rowData[colName] = columns[i]
		}
		*data = append(*data, rowData)
	}

	return nil

}

func ListQueries(c *ishell.Context) {
	queryMap := ParseJson(file, c)
	if queryMap == nil {
		return
	}

	c.Println(color.GreenString("\nAVAILABLE QUERIES: "))
	for key, val := range queryMap {
		c.Printf(color.CyanString("%s - %s\n"), key, val["description"])
	}
	c.Println("")
}

func SetFile(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	c.Print(color.CyanString("\nFile name: "))

	tmpFile := c.ReadLine()

	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		c.Println(color.RedString("File does not exist! Please try again!\n"))
		return
	}
	if !strings.Contains(strings.ToLower(tmpFile), ".json") {
		c.Println(color.YellowString(`Warning: File does not appear to be a json file! Please ensure that the file contains the proper json structure.
Note: to ensure that your json has been properly imported, please run `) + color.BlueString("'query:list'\n"))
	}

	file = tmpFile
	c.Println(color.GreenString("File set successfully!\n"))
}

func ParseJson(openFile string, c *ishell.Context) map[string]map[string]string {

	queryMap := make(map[string]map[string]string)

	jsonFile, err := os.Open(openFile)
	if err != nil {
		c.Printf(color.RedString("Could not open file %s. Are you sure the it's the right file/location?\n"), openFile)
		return nil
	}
	defer jsonFile.Close()

	jsonByte, _ := ioutil.ReadAll(jsonFile)
	jsonErr := json.Unmarshal([]byte(jsonByte), &queryMap)
	if jsonErr != nil {
		c.Println(color.RedString("Could not unmarshal json byte data - " + color.RedString(jsonErr.Error())))
		return nil
	}
	return queryMap
}
