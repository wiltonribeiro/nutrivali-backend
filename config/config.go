package config

import (
	"context"
	"github.com/BurntSushi/toml"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

type Config struct {
	Server   string
	Database string
}

var client *mongo.Client = nil

func read(config *Config) error {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return  err
	}
	return nil
}

func (c *Config) Connect() (*mongo.Database, error){
	err := read(c)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, c.Server)
	var database = client.Database(c.Database)
	return database, err
}

func (c *Config) Disconnect() {
	if client != nil {
		_ = client.Disconnect(context.Background())
	}
}