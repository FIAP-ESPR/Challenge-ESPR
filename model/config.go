package model

type Config struct {
	Version  string   `json:"version"`
	Ambient  Ambient  `json:"ambient"`
	Database Database `json:"database"`
	Logging  Logging  `json:"logging"`
	Template string   `json:"template"`
	Static   string   `json:"static"`
}

type Ambient struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol"`
}

type Database struct {
	Provider         string `json:"provider"`
	Host             string `json:"host"`
	Port             int    `json:"port"`
	ServiceName      string `json:"service_name"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Schema           string `json:"schema"`
	ConnectionString string `json:"connection_string"`
}

type Logging struct {
	Level     string `json:"level"`
	File      string `json:"file"`
	Mimetype  string `json:"mimetype"`
	Extension string `json:"extension"`
}
