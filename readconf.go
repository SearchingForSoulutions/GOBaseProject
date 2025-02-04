package main

import (
	"database/sql"
	"fmt"
	//"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)


// Valori di default, per creare le variabili. Occhio che le voglio tutte STRING.
var (
	db       *sql.DB  // Microsoft SQL Server
	server   = "127.0.0.1"
	port     = "1433"
	user     = "user" //"user"
	password = "pswd" //"pass"
	database = "DB" //"DB"
)

// porta webrouter, default 35000
var porta string = ":35000"
var host string = "http://localhost"
var conffile = "conf/base.conf"


func ReadEnv() error {
	// Carico le variabili dal file
	//home, sep := userHomeDir()
	fmt.Println("Carico configurazione da " + conffile)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Impossibile leggere " + conffile + ": " + err.Error())
		return err
	}
	server = os.Getenv("server")
	port = os.Getenv("port")
	user = os.Getenv("user")
	password = os.Getenv("password")
	database = os.Getenv("database")
	if !(len(os.Args) > 1 && os.Args[1] == "test") {
		host = os.Getenv("host")
	}
	porta = os.Getenv("porta")
	/*
	   result := server + " " + port + " " + user + " " + password + " " + database
	   log.Println(result)
	*/
	return nil
}


func userHomeDir() (string, string) {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home, "\\"
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("HOME")
		return home, "/"
	}
	// altri OS, generico:
	return os.Getenv("HOME"), "/"
}
