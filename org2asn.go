package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Result      string `json:"Result"`
	Type        string `json:"Type"`
	Description string `json:"Description"`
}

type ASNInfo struct {
	Org      string   `json:"org"`
	NumASN   int      `json:"num_ASN"`
	NumRoute int      `json:"num_Route"`
	Details  []Result `json:"details"`
}

func getASNInfo(orgName string) (ASNInfo, error) {
	url := fmt.Sprintf("https://bgp.he.net/search?search%%5Bsearch%%5D=%s&commit=Search", orgName)
	resp, err := http.Get(url)
	if err != nil {
		return ASNInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ASNInfo{}, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ASNInfo{}, err
	}

	var details []Result
	numASN := 0
	numRoute := 0

	doc.Find("div#search table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			// Skip header row
			return
		}
		columns := s.Find("td")

		result := strings.TrimSpace(columns.Eq(0).Text())
		type_ := strings.TrimSpace(columns.Eq(1).Text())
		description := strings.TrimSpace(columns.Eq(2).Text())

		if type_ != "Domain Name" && type_ != "TLD" {
			details = append(details, Result{
				Result:      result,
				Type:        type_,
				Description: description,
			})

			if type_ == "ASN" {
				numASN++
			} else if type_ == "Route" {
				numRoute++
			}
		}
	})

	return ASNInfo{
		Org:      orgName,
		NumASN:   numASN,
		NumRoute: numRoute,
		Details:  details,
	}, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		orgName := strings.TrimSpace(scanner.Text())
		asnInfo, err := getASNInfo(orgName)
		if err != nil {
			log.Printf("Error fetching ASN info for %s: %v\n", orgName, err)
			continue
		}

		// Print the result as formatted JSON
		jsonData, err := json.MarshalIndent(asnInfo, "", "  ")
		if err != nil {
			log.Printf("Error marshalling JSON: %v\n", err)
			continue
		}
		fmt.Println(string(jsonData))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
