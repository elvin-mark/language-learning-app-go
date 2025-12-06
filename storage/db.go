package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb(filePath string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", filePath)
	if err != nil {
		return
	}

	// enable foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return
	}

	return
}

// func main() {

// 	// 2. Create a table
// 	createTable := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		age INTEGER
// 	);`
// 	_, err = db.Exec(createTable)
// 	if err != nil {
// 		log.Fatal("Create table error:", err)
// 	}

// 	// 3. Insert data
// 	insertStmt := `INSERT INTO users (name, age) VALUES (?, ?)`
// 	result, err := db.Exec(insertStmt, "Alice", 25)
// 	if err != nil {
// 		log.Fatal("Insert error:", err)
// 	}

// 	id, _ := result.LastInsertId()
// 	fmt.Println("Inserted user with ID:", id)

// 	// 4. Query a single row
// 	var name string
// 	var age int
// 	row := db.QueryRow(`SELECT name, age FROM users WHERE id = ?`, id)
// 	err = row.Scan(&name, &age)
// 	if err != nil {
// 		log.Fatal("QueryRow error:", err)
// 	}
// 	fmt.Printf("User[%d] => Name: %s, Age: %d\n", id, name, age)

// 	// 5. Query multiple rows
// 	rows, err := db.Query(`SELECT id, name, age FROM users`)
// 	if err != nil {
// 		log.Fatal("Query error:", err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("\nAll users:")
// 	for rows.Next() {
// 		var uid int
// 		var uname string
// 		var uage int
// 		rows.Scan(&uid, &uname, &uage)
// 		fmt.Printf("ID: %d | Name: %s | Age: %d\n", uid, uname, uage)
// 	}

// 	// 6. Update a row
// 	_, err = db.Exec(`UPDATE users SET age = ? WHERE id = ?`, 30, id)
// 	if err != nil {
// 		log.Fatal("Update error:", err)
// 	}
// 	fmt.Println("\nUpdated user age to 30")

// 	// 7. Delete a row
// 	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, id)
// 	if err != nil {
// 		log.Fatal("Delete error:", err)
// 	}
// 	fmt.Println("Deleted user with ID:", id)

// 	fmt.Println("\nDone ✔")
// }
