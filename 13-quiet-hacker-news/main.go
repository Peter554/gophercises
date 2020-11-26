package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gophercises/quiet_hn/hn"
)

const (
	batchSize = 5
)

var (
	cacheMutex   sync.Mutex
	cacheStories []item
	cacheExpires time.Time
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		stories, err := getStories(numStories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func getStories(numStories int) ([]item, error) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if cacheStories != nil && time.Now().Before(cacheExpires) && len(cacheStories) >= numStories {
		return cacheStories[:numStories], nil
	}

	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("Failed to load top stories")
	}

	stories := []item{}
	batch := 1
	for len(stories) < numStories {
		batchStories := getStoriesBatch(ids[(batch-1)*batchSize : batch*batchSize])
		stories = append(stories, batchStories...)
		batch++
	}
	stories = stories[:numStories]

	cacheStories = stories
	cacheExpires = time.Now().Add(time.Second * 10)

	return stories, nil
}

func getStoriesBatch(ids []int) []item {
	var client hn.Client

	stories := []item{}
	var wg sync.WaitGroup

	ordering := map[int]int{}

	for idx, id := range ids {
		ordering[id] = idx
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			hnItem, err := client.GetItem(id)
			if err != nil {
				return
			}
			item := parseHNItem(hnItem)
			if isStoryLink(item) {
				stories = append(stories, item)
			}
		}(id)
	}

	wg.Wait()

	sort.Slice(stories, func(i, j int) bool {
		return ordering[stories[i].ID] < ordering[stories[j].ID]
	})

	return stories
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}
