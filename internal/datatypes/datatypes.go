package datatypes

type Project struct {
	ProjectName string
	Folders     []Folder
}

type Folder struct {
	FolderName string   `json:"folderName"`
	SubFolders []Folder `json:"subFolders"`
	Files      []string `json:"files"`
}

type File struct {
	FileName    string `json:"fileName"`
	FileContent string `json:"fileContent"`
	Extension   string `json:"extension"`
}

type DbUserConfig struct {
	DbProvider string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}