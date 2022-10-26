package celeritas

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	RootPath string
}

type PathConfig struct {
	rootPath    string
	folderNames []string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := PathConfig{rootPath: rootPath, folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"}}
	err := c.init(pathConfig)
	if err != nil {
		return err
	}

	err = checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(filepath.Join(rootPath, ".env"))
	if err != nil {
		return err
	}

	c.InfoLog, c.ErrorLog = c.initLoggers()

	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version

	return err
}

// creates the various directories / boilerplate files and folders needed to setup a project
func (c *Celeritas) init(p PathConfig) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		if err := createDirIfNotExist(filepath.Join(root, "/", path)); err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) initLoggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
