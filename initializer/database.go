package initializer

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*sql.DB, *gorm.DB, error) {

	credential, err := LoadCredential()
	if err != nil {
		return nil, nil, err
	}

	connStr := "dbname=" + credential.DbName +
		" user=" + credential.DbUsername +
		" password=" + credential.DbPassword +
		" host=" + credential.DbIpAddress +
		" port=" + credential.DbPort

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error can't connect to POSTGRES :", err)
		return nil, nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	fmt.Println("Connected to the Postgres Database")
	return db, gormDB, nil
}
