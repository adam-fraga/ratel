package migration

type Migration struct {
	ID         int
	name       string
	fileName   string
	Definition MigrationDefinition
}

type MigrationDefinition struct {
}

func (Migration) CreateMigration(migration Migration) {

}

func createMigrationFile() {

}

func promptMigrationDefinition() {

}

func writeMigrationFile() {

}
