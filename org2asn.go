package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rix4uni/org2asn/banner"
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

func getASNInfo(client *http.Client, orgName string) (ASNInfo, error) {
	// Replace spaces with '+'
    query := strings.Replace(orgName, " ", "+", -1)
	url := fmt.Sprintf("https://bgp.he.net/search?search%%5Bsearch%%5D=%s&commit=Search", query)

	resp, err := client.Get(url)
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

func readOrgNamesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var orgNames []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		orgName := strings.TrimSpace(scanner.Text())
		if orgName != "" {
			orgNames = append(orgNames, orgName)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return orgNames, nil
}

func main() {
	version := flag.Bool("version", false, "Print the version of the tool and exit.")
	silent := flag.Bool("silent", false, "Silent mode.")
	orgFlag := flag.String("org", "", "Organization name to search ASN info for.")
	listFlag := flag.String("list", "", "File with list of organization names to search ASN info for.")
	outputJSON := flag.Bool("json", false, "Output in JSON format.")
	timeout := flag.Int("timeout", 15, "Timeout for each HTTP request in seconds.")
	rateLimit := flag.Int("rate-limit", 0, "Rate limit in seconds between requests.")
	flag.Parse()

	if *version {
		banner.PrintBanner()
		banner.PrintVersion()
		return
	}

	if !*silent {
		banner.PrintBanner()
	}

	// HTTP client with timeout
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// Determine input source
	var orgNames []string
	if *orgFlag != "" {
		orgNames = append(orgNames, *orgFlag)
	} else if *listFlag != "" {
		var err error
		orgNames, err = readOrgNamesFromFile(*listFlag)
		if err != nil {
			log.Fatalf("Error reading organization list from file: %v", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			orgNames = append(orgNames, strings.TrimSpace(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
	}

	// Process each organization with rate limiting
	for _, orgName := range orgNames {
		asnInfo, err := getASNInfo(client, orgName)
		if err != nil {
			log.Printf("Error fetching ASN info for %s: %v\n", orgName, err)
			continue
		}

		if *outputJSON {
			// Print the result as formatted JSON
			jsonData, err := json.MarshalIndent(asnInfo, "", "  ")
			if err != nil {
				log.Printf("Error marshalling JSON: %v\n", err)
				continue
			}
			fmt.Println(string(jsonData))
		} else {
			// Print in plain text
			for _, detail := range asnInfo.Details {
				fmt.Printf("[%s] [%s] [%s] [%s]\n", asnInfo.Org, detail.Result, detail.Type, detail.Description)
			}
		}

		// Apply rate limiting
		time.Sleep(time.Duration(*rateLimit) * time.Second)
	}
}
