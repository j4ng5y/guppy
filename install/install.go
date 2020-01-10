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

	if err := I.install(); err != nil {
		return err
	}

	if err := I.writeRCFile(); err != nil {
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
	switch runtime.GOOS {
	case "windows":
		I.filename = fmt.Sprintf("%s.%s-%s.zip", I.version, runtime.GOOS, runtime.GOARCH)
	default:
		I.filename = fmt.Sprintf("%s.%s-%s.tar.gz", I.version, runtime.GOOS, runtime.GOARCH)
	}
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

func (I *Install) install() error {
	log.Printf("Extracting and Installing %s\n", I.version)

	switch runtime.GOOS {
	case "windows":
		if err := Unarchive(I.filepath, "C:\\"); err != nil {
			return fmt.Errorf("unable to unzip and install due to error: %v", err)
		}
		if err := os.Setenv("GOROOT", "C:\\go"); err != nil {
			return fmt.Errorf("unable to set the GOROOT environmental variable due to error: %v", err)
		}
	default:
		if err := Unarchive(I.filepath, "/usr/local"); err != nil {
			return fmt.Errorf("unable to untar and install due to error: %v", err)
		}
		if err := os.Setenv("GOROOT", "/usr/local/go"); err != nil {
			return fmt.Errorf("unable to set the GOROOT environmental variable due to error: %v", err)
		}
	}

	return nil
}

func (I *Install) writeRCFile() error {
	foundRCFiles := []string{}
	for _, i := range []string{
		"/etc/bashrc",
		"/etc/zshrc",
	} {
		_, err := os.Stat(i)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
		}
		foundRCFiles = append(foundRCFiles, i)

		f, err := os.OpenFile(i, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			return err
		}
		if _, err = f.WriteString("\nexport PATH=$PATH:/usr/local/go/bin\nexport GOROOT=/usr/local/go\n"); err != nil {
			return err
		}
	}

	if len(foundRCFiles) == 0 {
		return fmt.Errorf("no supported system level shell rc files found. You may need to add \"export PATH=$PATH:/usr/local/go/bin\" and \"export GOROOT=/usr/local/go\" to your persistent envrionment variables")
	}
	return nil
}
