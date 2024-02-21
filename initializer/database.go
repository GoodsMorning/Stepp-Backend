package initializer

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	credential, err := LoadCredential()
	if err != nil {
		return nil, err
	}

	connStr := "dbname=" + credential.DbName +
		" user=" + credential.DbUsername +
		" password=" + credential.DbPassword +
		" host=" + credential.DbIpAddress +
		" port=" + credential.DbPort

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error can't connect to POSTGRES :", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the Postgres Database")
	return db, nil
}
