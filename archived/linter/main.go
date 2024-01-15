package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Info struct {
	Name                string
	URL                 string
	Description         string
	Startdate           string
	Enddate             string
	Eventstatus         string
	Eventattendancemode string
	Location            Location
}

type Location struct {
	Type string
	URL  string
}

func cleanHTMLReactScriptTag(s string) string {
	s = strings.TrimPrefix(s, `<script data-react-helmet="true" type="application/ld+json">`)
	s = strings.TrimSuffix(s, `</script>`)
	return s
}

func GetMeetupsURLs(body []byte) ([]string, error) {
	var urls []string
	str2 := strings.SplitAfter(string(body), `</script>`)
	for _, s := range str2 {
		if strings.Contains(s, `"url"`) {
			m := make(map[string]interface{})
			meetupInfo := cleanHTMLReactScriptTag(s)
			json.Unmarshal([]byte(meetupInfo), &m)
			urls = append(urls, m["url"].(string))
		}
	}
	return urls, nil
}

type Client struct {
	client *http.Client
	proURL string
}

func Setup() *Client {
	c := Client{
		client: &http.Client{
			Timeout: time.Second * 300,
		},
		proURL: "https://www.meetup.com/pro/forge-utah/",
	}
	return &c
}

func (c *Client) GetProPage() ([]byte, error) {
	if c.client == nil {
		return []byte{}, errors.New("http.Client not initialized")
	}
	return c.GetWebPage(c.proURL)
}

func (c *Client) GetWebPage(url string) ([]byte, error) {
	if c.client == nil {
		return []byte{}, errors.New("http.Client not initialized")
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot create request %w", err)
	}
	request.AddCookie(&http.Cookie{
		Name:  "name",
		Value: "value",
	})
	request.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return []byte{}, fmt.Errorf("failure in Do request:\n %w ---\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Cannot parse response body: %w", err)
	}
	return body, nil
}

func (c *Client) GetMeetupInfo(url string) (Info, error) {
	body, err := c.GetWebPage(url)
	if err != nil {
		return Info{}, fmt.Errorf("Cannot get meetup Info: %w", err)
	}
	var parsedMeetup Info
	str2 := strings.SplitAfter(string(body), `</script>`)
	for _, s := range str2 {
		if strings.Contains(s, `"url"`) {
			infos := cleanHTMLReactScriptTag(s)
			json.Unmarshal([]byte(infos), &parsedMeetup)
			if url == parsedMeetup.URL {
				break
			}
			fmt.Println(parsedMeetup.URL)
			continue
		}
	}
	return parsedMeetup, nil
}

func runCrawler(w http.ResponseWriter, r *http.Request) {
	meetupClt := Setup()
	body, err := meetupClt.GetProPage()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	urls, err := GetMeetupsURLs(body)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	var infos []Info
	for _, url := range urls {
		info, err := meetupClt.GetMeetupInfo(url)
		fmt.Fprintln(w, err)
		infos = append(infos, info)
	}
	fmt.Fprintf(w, "%f meetups found\n", len(infos))
	fmt.Fprintf(w, "Saved meetups in firestore\n")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	c, _ := os.ReadDir("./")
	fmt.Fprintln(w, "Listing subdir/parent")
	for _, entry := range c {
		fmt.Fprintln(w, " ", entry.Name(), entry.IsDir())
	}
	w.Write([]byte("we are live"))
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/crawl", runCrawler)
	http.HandleFunc("/health", healthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
