package db

import (
	"fmt"
	"os"
	"os/exec"

	dt "github.com/adam-fraga/ratel/internal/datatypes"
	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
)

func InitDbDevelopmentContainer(dbProvider string) {
	if dbProvider == "postgres" || dbProvider == "mongo" || dbProvider == "mariadb" {
		var dbConfig dt.DbUserConfig
		dbConfig.DbProvider = dbProvider

		err := createDbContainer(&dbConfig)

		if err != nil {
			utils.PrintErrorMsg(err.Error())
		}
	} else if dbProvider == "sqlite" {
		createSqliteLocalDb()
	} else {
		var err = &errors.ClientError{Msg: fmt.Sprintf("Database provider \"%s\" not supported", dbProvider)}
		utils.PrintErrorMsg(err.Error())
	}
}

func createDbContainer(dbConfig *dt.DbUserConfig) error {
	dbConf, err := utils.PromptDbConfig(dbConfig)

	if err != nil {
		utils.PrintErrorMsg(err.Error())
	}

	switch dbConf.DbProvider {
	case "postgres":
		dbConf.DbPort = "5432"
		if err := runPostgresDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for Postgres SQL container: " + err.Error()}
		}
	case "mongo":
		dbConf.DbPort = "27017"
		if err := runMongoDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for Mongo container: " + err.Error()}
		}
	case "mariadb":
		dbConf.DbPort = "3306"
		if err := runMariadbDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for MariaDB container: " + err.Error()}
		}
	default:
		return &errors.ClientError{Msg: "Database provider not supported"}
	}
	return nil
}

func runPostgresDockerCmd(dbConfig *dt.DbUserConfig) error {
	utils.PrintInfoMsg("Creating a Postgres container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-e", "POSTGRES_PASSWORD="+dbConfig.DbPassword, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for Postgres container: " + err.Error()}
	}
	return nil
}

func runMongoDockerCmd(dbConfig *dt.DbUserConfig) error {
	utils.PrintInfoMsg("Creating a Mongo container")
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for Mongo containero : " + err.Error()}
	}
	return nil
}

func runMariadbDockerCmd(dbConfig *dt.DbUserConfig) error {
	utils.PrintInfoMsg("Creating a MariaDB container")
	cmd := exec.Command("docker", "run", "--detach", "--name", dbConfig.DbName, "--env", "MARIADB_ROOT_PASSWORD="+dbConfig.DbPassword, "-d", dbConfig.DbProvider+":latest")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for MySQL container: " + err.Error()}
	}
	return nil
}

func createSqliteLocalDb() {
	fmt.Println("Creating a SQLite local database")
}
