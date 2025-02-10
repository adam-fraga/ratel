package db

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	er "github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
	"golang.org/x/term"
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

		if err := createDbContainer(&dbConfig); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => func: InitDbDevelopmentContainer()",
				Msg:    "Failed to create database container, error: " + err.Error(),
				Err:    err,
			}
		}

	} else if dbProvider == "sqlite" {
		if err := createSqliteLocalDb(); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => Func: InitDbDevelopmentContainer()",
				Msg:    "Failed to create database container ",
				Err:    nil,
			}
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
				Msg:    "Failed to launch postgresql container, error: " + err.Error(),
				Err:    err,
			}
		}
	case "mongo":
		dbConf.DbPort = "27017"
		confirmConfig(dbConf)
		if err := runMongoDockerCmd(dbConf); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => Func: createDbContainer()",
				Msg:    "Failed to launch mongo db container, error: " + err.Error(),
				Err:    err,
			}
		}
	case "mariadb":
		dbConf.DbPort = "3306"
		confirmConfig(dbConf)
		if err := runMariadbDockerCmd(dbConf); err != nil {
			return &er.DBError{
				Origin: "File: handlers/db/createDbContainer.go => Func: createDbContainer()",
				Msg:    "Failed to launch maria db container, error: " + err.Error(),
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
func createSqliteLocalDb() error {
	fmt.Println("Creating a SQLite local database (TODO)")
	return nil
}

func promptDbConfig(dbConfig *DbUserConfig) *DbUserConfig {

	switch dbConfig.DbProvider {
	case "postgres":
		dbConfig.DbPort = "5432"
	case "mongo":
		dbConfig.DbPort = "27017"
	case "sqlite":
		dbConfig.DbPort = ""
	case "mariadb":
		dbConfig.DbPort = "3306"
	default:
		ut.PrintErrorMsg("DB Provider " + dbConfig.DbProvider + "is not supported")
	}

	os.Stdin.WriteString("Please enter the database user: ")
	fmt.Scanln(&dbConfig.DbUser)

	os.Stdin.WriteString("Please enter the database name: ")
	fmt.Scanln(&dbConfig.DbName)

	os.Stdin.WriteString("Please enter the database password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		ut.PrintErrorMsg("error reading password, error: " + err.Error())
		return promptDbConfig(dbConfig)
	}
	fmt.Println()

	dbConfig.DbPassword = string(password)

	os.Stdin.WriteString("Please confirm password: ")
	passwordConfirm, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		ut.PrintErrorMsg("Error reading password confirmation")
		return promptDbConfig(dbConfig)
	}
	fmt.Println() // Add a newline after password confirmation

	if dbConfig.DbPassword != string(passwordConfirm) {
		ut.PrintErrorMsg("Your informations are incorrect")
		return promptDbConfig(dbConfig)
	}

	return dbConfig
}

func confirmConfig(dbConfig *DbUserConfig) {
	ut.PrintInfoMsg(fmt.Sprintf("\nDatabase configuration:\n\nDB Provider: %s\nDB Port: %s\nDB User: %s\nDB Name: %s\n",
		dbConfig.DbProvider, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbName))

	os.Stdin.WriteString("Is the configuration correct? (Y/N): ")
	var confirm string

	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) != "y" {
		ut.PrintInfoMsg("You need to confirm with \"y\"")
		promptDbConfig(dbConfig)
	}
}
