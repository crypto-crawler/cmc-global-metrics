package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func fetch_cmc_global_metrics() string {
	url := "https://pro-api.coinmarketcap.com/v1/global-metrics/quotes/latest"
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)

	cmc_api_key := os.Getenv("CMC_API_KEY")
	if len(cmc_api_key) == 0 {
		log.Fatalln("The CMC_API_KEY environment variable is empty")
	}
	req.Header.Set("X-CMC_PRO_API_KEY", cmc_api_key)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	
	return string(body)
}

func main() {
	text := fetch_cmc_global_metrics()

	filename := time.Now().Format("2006-01") + ".json"
	filepath := filepath.Join("data", filename)
	
	file, err := os.OpenFile(filepath,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if _, err := file.WriteString(text + "\n"); err != nil {
		log.Fatalln(err)
	}
}
