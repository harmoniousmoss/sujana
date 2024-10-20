package scraper

import (
	"errors"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Job represents a remote job listing.
type Job struct {
	Title string `bson:"title" json:"title"`
	Link  string `bson:"link" json:"link"`
}

// ScrapeJobs scrapes jobs with retries and proxy support.
func ScrapeJobs() ([]Job, error) {
	client, err := createHTTPClient()
	if err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		jobs, err := tryScrape(client)
		if err == nil {
			return jobs, nil
		}

		log.Printf("Attempt %d failed: %v\n", i+1, err)
		time.Sleep(backoffTime(i)) // Exponential backoff
	}

	return nil, errors.New("failed to scrape jobs after retries")
}

// createHTTPClient initializes an HTTP client with proxy support and timeouts.
func createHTTPClient() (*http.Client, error) {
	proxyURLStr := os.Getenv("PROXY_URL")
	if proxyURLStr == "" {
		return nil, errors.New("proxy URL not configured")
	}

	proxyURL, err := url.Parse(proxyURLStr)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			MaxIdleConns:        10,
			MaxConnsPerHost:     10,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 30 * time.Second,
	}, nil
}

// tryScrape performs a single scrape attempt and returns the jobs or an error.
func tryScrape(client *http.Client) ([]Job, error) {
	resp, err := makeRequest(client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + resp.Status)
	}

	return parseJobs(resp.Body)
}

// makeRequest sends an HTTP request to the target URL.
func makeRequest(client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://unjobs.org/search/remote", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")
	req.Header.Set("Connection", "keep-alive")

	log.Println("Sending request to https://unjobs.org/search/remote")
	return client.Do(req)
}

// parseJobs reads the response body and extracts job listings using goquery.
func parseJobs(body io.Reader) ([]Job, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	var jobs []Job
	doc.Find(".job").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		link, exists := s.Find("a").Attr("href")
		if exists && !strings.HasPrefix(link, "http") {
			link = "https://unjobs.org" + link
		}
		jobs = append(jobs, Job{Title: title, Link: link})
	})

	log.Printf("Scraped %d jobs\n", len(jobs))
	return jobs, nil
}

// backoffTime calculates the backoff time with jitter.
func backoffTime(attempt int) time.Duration {
	base := math.Pow(2, float64(attempt))
	jitter := rand.Float64()
	return time.Duration(base+jitter) * time.Second
}
