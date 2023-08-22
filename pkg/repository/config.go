package repository

type MongoConfig struct {
	URI      string `toml:"uri"`
	Database string `toml:"database"`
}
