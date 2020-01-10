package install

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
)

// Install is a struct to keep tract of the install
type Install struct {
	version  string
	pwd      string
	filename string
	filepath string
}

// New is a function that constructs a new instance of Install
//
// Arguments:
//     version (string): The version of Go to install
//
// Returns:
//     (*Install): A pointer to the new instance of Install
func New(version string) *Install {
	var I Install

	I.version = version

	return &I
}

// Run is a function that runs all operations to perform an install
//
// Arguments:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (I *Install) Run() error {
	I.getPWD()
	I.calculateFilename()
	I.calculateFilePath()

	if err := I.download(); err != nil {
		return err
	}
	return nil
}

func (I *Install) getPWD() {
	var err error

	I.pwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

func (I *Install) calculateFilename() {
	I.filename = fmt.Sprintf("%s.%s-%s.tar.gz", I.version, runtime.GOOS, runtime.GOARCH)
}

func (I *Install) calculateFilePath() {
	I.filepath = path.Join(I.pwd, I.filename)
}

func (I *Install) download() error {
	var err error

	log.Printf("Downloading %s\n", I.version)
	defer log.Println("Done")

	resp, err := http.Get(fmt.Sprintf("https://dl.google.com/go/%s", I.filename))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(I.filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return err
}
