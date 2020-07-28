package common

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func ReadCSV(path string) (*[][]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}
	return &csvLines, nil
}

func ParseURL(uri string) *url.URL {
	url, err := url.Parse(uri)

	if err != nil {
		log.Fatalf("could not parse url %q: %v", uri, err)
	}

	if url.Scheme == "" {
		url.Scheme = "https"
	}

	return url
}