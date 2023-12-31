package core

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"sync"
)

type Application struct {
	config          Configuration
	errorLog        *log.Logger
	infoLog         *log.Logger
	srv             *http.Server
	db              *sqlx.DB
	dbaseConnected  bool
	vk              *api.VK
	counter         int
	taskCounter     int
	completeCounter int
	running         bool
	wg              sync.WaitGroup
	defVkClient     *http.Client
	proxyClient     *http.Client
	currentClient   *http.Client
}

var App Application

func InitCore() {

	println("Starting VKExplorer app...")
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("Executable path:", exePath)

	//
	LoadConfig()

	// init log system
	App.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}

func StartCore() {
	//InitTray()
	//go HandleTray()

	InitDatabase()
	//go handleDatabase()

	InitWeb()
	go HandleWeb()

	fmt.Println("Server is listening http://127.0.0.1:8080 ...")
}

func QuitCore() {
	quitDatabase()
	println("VKExplorer app finished")
}
