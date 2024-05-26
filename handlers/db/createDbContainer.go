package db

import (
	"fmt"
	"os"
	"os/exec"

	dt "github.com/adam-fraga/ratel/datatypes"
	"github.com/adam-fraga/ratel/errors"
	"github.com/adam-fraga/ratel/utils"
)

func InitDb(dbProvider string) {
	if dbProvider == "postgres" || dbProvider == "mongo" || dbProvider == "mysql" {
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
		utils.PrintInfoMsg("Creating a PostgreSQL container")
		if err := runPostgresDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for Postgres container: " + err.Error()}
		}
	case "mongo":
		dbConf.DbPort = "27017"
		utils.PrintInfoMsg("Creating a MongoDB container")
		if err := runMongoDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for Mongo container: " + err.Error()}
		}
	case "mysql":
		dbConf.DbPort = "3306"
		utils.PrintInfoMsg("Creating a MySQL container")
		if err := runMysqlDockerCmd(dbConf); err != nil {
			return &errors.ClientError{Msg: "Error running the command for MySQL container: " + err.Error()}
		}
	default:
		return &errors.ClientError{Msg: "Database provider not supported"}
	}
	return nil
}

func runPostgresDockerCmd(dbConfig *dt.DbUserConfig) error {
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-e", "POSTGRES_PASSWORD="+dbConfig.DbPassword, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for Postgres container"}
	}
	return nil
}

func runMongoDockerCmd(dbConfig *dt.DbUserConfig) error {
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for Mongo container"}
	}
	return nil
}

func runMysqlDockerCmd(dbConfig *dt.DbUserConfig) error {
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-e", "MYSQL_ROOT_PASSWORD="+dbConfig.DbPassword, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	if err := cmd.Run(); err != nil {
		return &errors.ClientError{Msg: "Error running the command for MySQL container"}
	}
	return nil
}

func createSqliteLocalDb() {
	fmt.Println("Creating a SQLite local database")
}
