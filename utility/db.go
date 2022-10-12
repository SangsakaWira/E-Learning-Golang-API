package utility

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func DatabaseInit() {

	db_url := goDotEnvVariable("URL")
	conn, err := sql.Open("pgx", db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	testConnection(conn)
	defer conn.Close()
}

func testConnection(conn *sql.DB) error {
	rows, _ := conn.Query("select * from pelanggan")

	for rows.Next() {
		var id int
		var nama string
		var alamat string
		var email string
		var dompet int
		err := rows.Scan(&id, &nama, &alamat, &email, &dompet)
		if err != nil {
			return err
		}
		fmt.Printf("%d. %s %s %s %d\n", id, nama, alamat, email, dompet)
	}

	return rows.Err()
}
