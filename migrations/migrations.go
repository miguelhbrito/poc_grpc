package migrations

import (
	"bufio"
	"database/sql"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"fmt"
	"os"
)

func InitMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	var m *migrate.Migrate
	if runtime.GOOS == "windows" {
		fmt.Print("Windows OS detected, please enter project path(example: C:/Users/username/Documents/dev/github/):")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		m, err = migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://"+text+"poc_grpc/migrations/"),
			"postgres", driver)
		if err != nil {
			panic(err)
		}

	} else {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		m, err = migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s/migrations/", pwd),
			"postgres", driver)
		if err != nil {
			panic(err)
		}
	}

	_ = m.Down()
	_ = m.Up()

	fmt.Println("Successfully migrations applied!")
}
