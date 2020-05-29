package ignore

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// CreateGitIgnore creates the .gitignore on given path
func CreateGitIgnore(params string, path string) error {
	fullPath := filepath.Join(path, ".gitignore")

	gi, err := FetchGitIgnore(params)
	if err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(gi)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// FetchGitIgnore reads the bytes from gitignore.io API
func FetchGitIgnore(params string) ([]byte, error) {
	fetchURL := fmt.Sprintf("http://gitignore.io/api/%s", params)

	resp, err := http.Get(fetchURL)
	if err != nil {
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
