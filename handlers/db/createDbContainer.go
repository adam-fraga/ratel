package db

import (
	"fmt"
	"os"
	"os/exec"

	dt "github.com/adam-fraga/ratel/internal/datatypes"
	"github.com/adam-fraga/ratel/internal/errors"
	ut "github.com/adam-fraga/ratel/utils"
)

func InitDbDevelopmentContainer(dbProvider string) {
	if dbProvider == "postgres" || dbProvider == "mongo" || dbProvider == "mariadb" {
		var dbConfig dt.DbUserConfig
		dbConfig.DbProvider = dbProvider

		err := createDbContainer(&dbConfig)

		if err != nil {
			ut.PrintErrorMsg(err.Error())
		}
	} else if dbProvider == "sqlite" {
		createSqliteLocalDb()
	} else {
		var err = &errors.ClientError{Msg: fmt.Sprintf("Database provider \"%s\" not supported", dbProvider)}
		ut.PrintErrorMsg(err.Error())
	}
}

func createDbContainer(dbConfig *dt.DbUserConfig) error {
	dbConf, err := promptDbConfig(dbConfig)

	if err != nil {
		ut.PrintErrorMsg(err.Error())
	}

	switch dbConf.DbProvider {
	case "postgres":
		confirmConfig(dbConf)
		if err := runPostgresDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for Postgres SQL container: " + err.Error()}
		}
	case "mongo":
		dbConf.DbPort = "27017"
		confirmConfig(dbConf)
		if err := runMongoDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for Mongo container: " + err.Error()}
		}
	case "mariadb":
		dbConf.DbPort = "3306"
		confirmConfig(dbConf)
		if err := runMariadbDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for MariaDB container: " + err.Error()}
		}
	default:
		return &errors.ClientError{Msg: "Database provider not supported"}
	}
	return nil
}

func runPostgresDockerCmd(dbConfig *dt.DbUserConfig) error {
	ut.PrintInfoMsg("Creating a Postgres container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider+"-"+dbConfig.DbName, "-e",
		"POSTGRES_PASSWORD="+dbConfig.DbPassword,
		"-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for Postgres container: " + err.Error()}
	}
	return nil
}

func runMongoDockerCmd(dbConfig *dt.DbUserConfig) error {
	ut.PrintInfoMsg("Creating a Mongo container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for Mongo containero : " + err.Error()}
	}
	return nil
}

func runMariadbDockerCmd(dbConfig *dt.DbUserConfig) error {
	ut.PrintInfoMsg("Creating a MariaDB container")
	cmd := exec.Command("docker", "run", "--detach",
		"--name", dbConfig.DbName,
		"--env", "MARIADB_USER="+dbConfig.DbUser,
		"--env", "MARIADB_PASSWORD="+dbConfig.DbPassword,
		"--env", "MARIADB_ROOT_PASSWORD="+dbConfig.DbPassword,
		"-d", dbConfig.DbProvider+":latest")

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for MySQL container: " + err.Error()}
	}
	return nil
}

func createSqliteLocalDb() {
	fmt.Println("Creating a SQLite local database")
}

func promptDbConfig(dbConfig *dt.DbUserConfig) (*dt.DbUserConfig, error) {

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

func confirmConfig(dbConfig *dt.DbUserConfig) {
	ut.PrintInfoMsg(fmt.Sprintf("\nDatabase configuration:\n\nDB Provider: %s\nDB Port: %s\nDB User: %s\nDB Name: %s\n",
		dbConfig.DbProvider, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbName))

	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var confirm string

	fmt.Scanln(&confirm)

	if confirm == "n" {
		promptDbConfig(dbConfig)
	}
}
