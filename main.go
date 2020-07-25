package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Test1(url string) error {
	// ts := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return nil
	}

	return nil
}

type ReportWebSite struct {
	TotalWebSites int   `json:"total_websites"`
	Success       int   `json:"success"`
	Failure       int   `json:"failure"`
	TotalTime     int64 `json:"total_time"`
}

func ReadCSV(path string) (*[][]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}
	return &csvLines, nil
}

func main() {

	r := ReportWebSite{}

	urlsLine, err := ReadCSV("emp.csv")
	if err != nil {
		fmt.Println("error csv")
	}

	r.TotalWebSites = len(*urlsLine)

	for _, urlLine := range *urlsLine {
		url := urlLine[0]

		ts := time.Now()
		err := Test1(url)
		tn := time.Since(ts).Nanoseconds()

		r.TotalTime = r.TotalTime + tn

		if err != nil {
			r.Failure = r.Failure + 1
		} else {
			r.Success = r.Success + 1
		}
	}

	fmt.Println("Perform website checking...")
	fmt.Println("Done!")
	fmt.Println("Checked webistes: ", r.TotalWebSites)
	fmt.Println("Successful websites: ", r.Success)
	fmt.Println("Failure websites: ", r.Failure)
	fmt.Println("Total times to finished checking website: ", r.TotalTime/1000000000)

}
