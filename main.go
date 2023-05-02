package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks" // extractlinks GO package for extracting anchor links from HTML
	//Extracts all anchor links from a HTML page into an Array of []Link
)

var ( // combine all the clients together and declared as a variable
	config = &tls.Config{
		InsecureSkipVerify: true,
	}

	transport = &http.Transport{
		TLSClientConfig: config,
	}

	netClient = &http.Client{
		Transport: transport,
	}

	queue = make(chan string)

	hasVisited = make(map[string]bool)
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func crawlUrl(href string, w http.ResponseWriter, r *http.Request) {
	hasVisited[href] = true
	fmt.Fprintf(w, "# %s\n", href)
	response, err := netClient.Get(href)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	links, err := extractlinks.All(response.Body) // extracting all the links from the html body
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, link := range links {
		absoluteURL := toFixedUrl(link.Href, href)
		go func() {
			queue <- absoluteURL // queue is used where links are added  in the back, links in the front are used to crawl
		}()
	}
}

// Ensures to crawl in a single domain
func isSameDomain(href, baseUrl string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}
	parentUri, err := url.Parse(baseUrl)
	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host {
		return false
	}

	return true
}

func toFixedUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" { // representation of url [scheme:][//[userinfo@]host][/]path[?query][#fragment]
		return base
	}
	baseUrl, err := url.Parse(base) // Parse parses a url into a URL structure.
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri) // will take host from base & path from uri. If it has its own host then it'll do base.Host + uri.Path
	return uri.String()
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	link := r.FormValue("crawl")
	fmt.Fprintf(w, "HI ! %s\n", name)
	fmt.Fprintf(w, "BASE LINK : %v\n", link)
	fmt.Fprintf(w, "Crawling......Please wait\n")

	go func() {
		queue <- link
	}()

	for href := range queue {
		if !hasVisited[href] && isSameDomain(href, link) {
			crawlUrl(href, w, r)
		}
	}
}
