package main

import (
	"GoDemo/internal/app"
	"GoDemo/internal/plog"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const LowestPort = 2_000
const HighestPort = 10_000
const ConfigFilePath = ".\\config\\"
const ConfigFileName = "excludeFromGit"
const ConfigFileType = "yaml"
const DatabaseStringName = "DATABASE_ACCESS_STRING"

func main() {
	portNumber := flag.Int("port", LowestPort, "The port number to listen on")
	flag.Parse()

	viper.AddConfigPath(ConfigFilePath)
	viper.SetConfigType(ConfigFileType)
	viper.SetConfigName(ConfigFileName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	db, err := openDB(viper.GetString(DatabaseStringName))
	if err != nil {
		log.Fatalf("Could not open the database.  Sql.Open err: %v", err)
	}
	defer db.Close()

	application := app.NewApp()
	application.Database = db

	foundWorkingSocket := false
	for !foundWorkingSocket {
		application.Logger.Info("Starting server...", plog.KV{Key: "port", Value: *portNumber})
		err := http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), application.Routes())
		if err != nil {
			application.Logger.Error("Failed attempt to connect...", plog.KV{Key: "port", Value: *portNumber})
			*portNumber = *portNumber + 1
		} else if *portNumber > HighestPort {
			application.Logger.Error("Failed to connect to any port.  Stopping.")
			return
		} else {
			foundWorkingSocket = true
			application.Logger.Info("Successfully ran server.", plog.KV{Key: "port", Value: *portNumber})
		}
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
