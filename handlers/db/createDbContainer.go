package db

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/adam-fraga/ratel/internal/errors"
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
			err := fmt.Errorf("Error trying to create DB container: " + err.Error())
			ut.PrintErrorMsg(err.Error())
			return err
		}
	} else if dbProvider == "sqlite" {
		createSqliteLocalDb()
	} else {
		err := fmt.Errorf("Database provider \"%s\" not supported", dbProvider)
		ut.PrintErrorMsg(err.Error())
		return err
	}
	return nil
}

// createDbContainer create the database container
func createDbContainer(dbConfig *DbUserConfig) error {
	dbConf, err := promptDbConfig(dbConfig)

	if err != nil {
		ut.PrintErrorMsg(err.Error())
	}

	switch dbConf.DbProvider {
	case "postgres":
		confirmConfig(dbConf)
		if err := runPostgresDockerCmd(dbConf); err != nil {
			ut.PrintErrorMsg(err.Error())
			return fmt.Errorf("Error running the command for Postgres SQL container: " + err.Error())
		}
	case "mongo":
		dbConf.DbPort = "27017"
		confirmConfig(dbConf)
		if err := runMongoDockerCmd(dbConf); err != nil {
			ut.PrintErrorMsg(err.Error())
			return fmt.Errorf("Error running the command for Mongo container: " + err.Error())
		}
	case "mariadb":
		dbConf.DbPort = "3306"
		confirmConfig(dbConf)
		if err := runMariadbDockerCmd(dbConf); err != nil {
			wrappedErr := fmt.Errorf("Error running the command for MariaDB container: %w", err)
			ut.PrintErrorMsg(wrappedErr.Error())
			return fmt.Errorf("" + err.Error())
		}
	default:
		err = fmt.Errorf("Database provider is not supported")
		ut.PrintErrorMsg(err.Error())
		return err
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
		wrappedErr := fmt.Errorf("Error running the command for Postgres container: %w", err)
		ut.PrintErrorMsg(wrappedErr.Error())
		return wrappedErr
	}
	return nil
}

// runMongoDockerCmd run the command to create a Mongo container
func runMongoDockerCmd(dbConfig *DbUserConfig) error {
	ut.PrintInfoMsg("Creating a Mongo container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		wrappedErr := fmt.Errorf("error running the command for Mongo container: %w", err)
		ut.PrintErrorMsg(wrappedErr.Error())
		return wrappedErr
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
		err = fmt.Errorf("Error running the command for MySQL container: " + err.Error())
		return err
	}
	return nil
}

// createSqliteLocalDb create a SQLite local database
func createSqliteLocalDb() {
	fmt.Println("Creating a SQLite local database")
}

func promptDbConfig(dbConfig *DbUserConfig) (*DbUserConfig, error) {

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
		err := &errors.ClientError{Msg: "Sorry your passwords do not match try again"}
		ut.PrintErrorMsg(err.Error())
		promptDbConfig(dbConfig)
	}

	return dbConfig, nil
}

func confirmConfig(dbConfig *DbUserConfig) {
	ut.PrintInfoMsg(fmt.Sprintf("\nDatabase configuration:\n\nDB Provider: %s\nDB Port: %s\nDB User: %s\nDB Name: %s\n",
		dbConfig.DbProvider, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbName))

	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var confirm string

	fmt.Scanln(&confirm)

	if confirm == "n" {
		promptDbConfig(dbConfig)
	}
}
