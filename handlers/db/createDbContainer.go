package db

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/adam-fraga/ratel/errors"
	dt "github.com/adam-fraga/ratel/models/datatypes"
	"github.com/adam-fraga/ratel/utils"
)

func InitDb(dbProvider string) {
	if dbProvider == "postgres" || dbProvider == "mongo" || dbProvider == "mysql" {
		var dbConfig dt.DbConfig
		dbConfig.DbProvider = dbProvider
		createDbContainer(&dbConfig)
	} else if dbProvider == "sqlite" {
		createSqliteLocalDb()
	} else {
		utils.PrintErrorMsg(fmt.Sprintf("Database provider \"%s\" not supported", dbProvider))
	}
}

func createDbContainer(dbConfig *dt.DbConfig) {
	dbConf, err := promptUserDbConfig(dbConfig)

	if err != nil {
		utils.PrintErrorMsg(err.Error())
	}

	switch dbConf.DbProvider {
	case "postgres":
		dbConf.DbPort = "5432"
		utils.PrintInfoMsg("Creating a PostgreSQL container")
		runPostgresDockerCmd(dbConf)
	case "mongo":
		dbConf.DbPort = "27017"
		utils.PrintInfoMsg("Creating a MongoDB container")
		runMongoDockerCmd(dbConf)
	case "mysql":
		dbConf.DbPort = "3306"
		utils.PrintInfoMsg("Creating a MySQL container")
		runMysqlDockerCmd(dbConf)
	default:
		utils.PrintErrorMsg(fmt.Sprintf("Database provider \"%s\" not supported", dbConf.DbProvider))
	}
}

func runPostgresDockerCmd(dbConfig *dt.DbConfig) {
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-e", "POSTGRES_PASSWORD="+dbConfig.DbPassword, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running the command" + err.Error())
	}
}

func runMongoDockerCmd(dbConfig *dt.DbConfig) {
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running the command")
	}
}

func runMysqlDockerCmd(dbConfig *dt.DbConfig) {
	cmd := exec.Command("docker", "run", "--name", dbConfig.DbProvider, "-e", "MYSQL_ROOT_PASSWORD="+dbConfig.DbPassword, "-d", "-p", dbConfig.DbPort+":"+dbConfig.DbPort, dbConfig.DbProvider)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running the command")
	}
}

func createSqliteLocalDb() {
	fmt.Println("Creating a SQLite local database")
}

func promptUserDbConfig(dbConfig *dt.DbConfig) (*dt.DbConfig, error) {

	os.Stdin.WriteString("Please enter the database user: ")
	fmt.Scanln(&dbConfig.DbUser)

	os.Stdin.WriteString("Please enter the database password: ")
	fmt.Scanln(&dbConfig.DbPassword)

	os.Stdin.WriteString("Please confirm password: ")
	var passwordConfirm string
	fmt.Scanln(&passwordConfirm)

	if dbConfig.DbPassword != passwordConfirm {
		err := &errors.ClientError{Msg: "Sorry your passwords do not match try again"}
		utils.PrintErrorMsg(err.Error())
		promptUserDbConfig(dbConfig)
	}

	os.Stdin.WriteString("Please enter the database name: ")
	fmt.Scanln(&dbConfig.DbName)

	utils.PrintInfoMsg(fmt.Sprintf("\nDatabase configuration:\n\nDB Port: %s\nDB User: %s\nDB Name: %s\n",
		dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbName))

	//Ask the user if the configuration is correct
	os.Stdin.WriteString("Is the configuration correct? (y/n): ")
	var response string

	fmt.Scanln(&response)

	if response == "n" {
		promptUserDbConfig(dbConfig)
	}

	return dbConfig, nil
}
