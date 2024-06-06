package config

var (
	// MongoDBConfig holds MongoDB connection settings
	MongoDBConfig = struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
	}{
		Host:     "localhost",
		Port:     "27017",
		Username: "my username",
		Password: "my password",
		Database: "phonebook",
	}
)
