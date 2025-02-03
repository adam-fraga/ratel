package db

import (
	"fmt"
	"os"
	"os/exec"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

// DbUserConfig represent the configuration of the database for the local development
type DbUserConfig struct {
	DbProvider string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// InitDbDevelopmentContainer run the database container according to the choosed provider
func InitDbDevelopmentContainer(dbProvider string) error {
	if dbProvider == "postgres" || dbProvider == "mongo" || dbProvider == "mariadb" {
		var dbConfig DbUserConfig
		dbConfig.DbProvider = dbProvider

		err := createDbContainer(&dbConfig)

		if err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => func: InitDbDevelopmentContainer()",
				Msg:    "Failed to create database container",
				Err:    err,
			}
		}
	} else if dbProvider == "sqlite" {
		createSqliteLocalDb()
	} else {
		return &er.DBError{
			Origin: "File: handlers/db/createDbContainer.go => Func: InitDbDevelopmentContainer()",
			Msg:    "Error: Unsupported database provider. Please choose one of the following supported providers: postgres, mongo, mariadb, sqlite.",
			Err:    nil,
		}
	}
	return nil
}

// createDbContainer create the database container
func createDbContainer(dbConfig *DbUserConfig) error {
	dbConf := promptDbConfig(dbConfig)

	switch dbConf.DbProvider {
	case "postgres":
		confirmConfig(dbConf)
		if err := runPostgresDockerCmd(dbConf); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => Func: createDbContainer()",
				Msg:    "Error running the command for Postgres SQL container",
				Err:    err,
			}
		}
	case "mongo":
		dbConf.DbPort = "27017"
		confirmConfig(dbConf)
		if err := runMongoDockerCmd(dbConf); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => Func: createDbContainer()",
				Msg:    "Error running the command for MongoDB container",
				Err:    err,
			}
		}
	case "mariadb":
		dbConf.DbPort = "3306"
		confirmConfig(dbConf)
		if err := runMariadbDockerCmd(dbConf); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => Func: createDbContainer()",
				Msg:    "Error running the command for Mariadb mongo container",
				Err:    err,
			}
		}
	default:
		return &er.DBError{
			Origin: "File: handlers/db/createDbContainer.go => Func: createDbContainer()",
			Msg:    "Unsupported database provider. Please choose one of the following supported providers: postgres, mongo, mariadb, sqlite.",
			Err:    nil,
		}
	}
	return nil
}

// runPostgresDockerCmd run the command to create a Postgres container
func runPostgresDockerCmd(dbConfig *DbUserConfig) error {
	ut.PrintInfoMsg("Creating a Postgres container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider+"-"+dbConfig.DbName, "-e",
		"POSTGRES_PASSWORD="+dbConfig.DbPassword,
		"-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &er.DBError{
			Origin: "File: handlers/db/createDbContainer.go => Func: runPostgresDockerCmd()",
			Msg:    "Error running the command for Postgres container",
			Err:    nil,
		}
	}
	return nil
}

// runMongoDockerCmd run the command to create a Mongo container
func runMongoDockerCmd(dbConfig *DbUserConfig) error {
	ut.PrintInfoMsg("Creating a Mongo container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &er.DBError{
			Origin: "File: handlers/db/createDbContainer.go => Func: runMongoDockerCmd()",
			Msg:    "Error running the command for Mongo container",
			Err:    nil,
		}
	}
	return nil
}

// runMariadbDockerCmd run the command to create a MariaDB container
func runMariadbDockerCmd(dbConfig *DbUserConfig) error {
	ut.PrintInfoMsg("Creating a MariaDB container")
	cmd := exec.Command("docker", "run", "--detach",
		"--name", dbConfig.DbName,
		"--env", "MARIADB_USER="+dbConfig.DbUser,
		"--env", "MARIADB_PASSWORD="+dbConfig.DbPassword,
		"--env", "MARIADB_ROOT_PASSWORD="+dbConfig.DbPassword,
		"-d", dbConfig.DbProvider+":latest")

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &er.DBError{
			Origin: "File: handlers/db/createDbContainer.go => Func: runMariadbDockerCmd()",
			Msg:    "Error running the command for MariaDB container",
			Err:    nil,
		}
	}
	return nil
}

// createSqliteLocalDb create a SQLite local database
func createSqliteLocalDb() {
	fmt.Println("Creating a SQLite local database (TODO)")
}

func promptDbConfig(dbConfig *DbUserConfig) *DbUserConfig {

	if dbConfig.DbProvider != "postgres" {
		os.Stdin.WriteString("Please enter the database user: ")
		fmt.Scanln(&dbConfig.DbUser)

		os.Stdin.WriteString("Please enter the database name: ")
		fmt.Scanln(&dbConfig.DbName)
	} else {
		dbConfig.DbPort = "5432"
		dbConfig.DbName = "postgres"
		dbConfig.DbUser = "postgres"
	}

	os.Stdin.WriteString("Please enter the database password: ")
	fmt.Scanln(&dbConfig.DbPassword)

	os.Stdin.WriteString("Please confirm password: ")
	var passwordConfirm string
	fmt.Scanln(&passwordConfirm)

	if dbConfig.DbPassword != passwordConfirm {
		ut.PrintErrorMsg("Your informations are incorrect")
		promptDbConfig(dbConfig)
	}

	return dbConfig
}

func confirmConfig(dbConfig *DbUserConfig) {
	ut.PrintInfoMsg(fmt.Sprintf("\nDatabase configuration:\n\nDB Provider: %s\nDB Port: %s\nDB User: %s\nDB Name: %s\n",
		dbConfig.DbProvider, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbName))

	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var confirm string

	fmt.Scanln(&confirm)

	if confirm != "y" {
		ut.PrintInfoMsg("You need to confirm with \"y\"")
		promptDbConfig(dbConfig)
	}
}
