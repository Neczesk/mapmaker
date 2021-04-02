package mapmaker

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func TestImport() {
	database, _ := sql.Open("sqlite3", "./world.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tiles (tile_id INTEGER_PRIMARY_KEY, tile_type STRING NOT NULL UNIQUE")
	statement, _ = database.Prepare("INSERT INTO tiles (tile_id, tile_type) VALUES (0, 'grassland')")
	statement.Exec("bob", "dylan")
	rows, _ := database.Query("SELECT tile_id, tile_type FROM tiles")
	var id int
	var tile_type string
	for rows.Next() {
		rows.Scan(&id, &tile_type)
		fmt.Println(strconv.Itoa(id) + ": " + tile_type)
	}
}
