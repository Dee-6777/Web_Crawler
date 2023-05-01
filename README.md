# WEB CRAWLER

### _Problem Statement_
```
Build a simple web crawler service (server and client) in a programming language of your choice. 
The server would receive requests from a client to crawl a URL and should send the response i.e sitemap back to the client. 
The Service/Server which is crawler should be limited to one domain 
```

## DESCRIPTION

A web server written in Golang which crawls a URL and generate all the links present in each page of the base url. 

## PREREQUISITE

The only prerequisite is to have `go1.19.5` installed in the local system!!


## GET STARTED 

### INSTALLATION
```
go install github.com/Dee-6777/Web_Crawler
```

### HOW TO USE THIS SERVER? 
* First Step: `Start a server at the port :8080 by running this command`
```
go run main.go 
```
* Second Step: `Paste this url in your browser and fill this form and submit. It will take a while to load an all the links will be generated on the screen`
```
http://localhost:8080/form.html
```
## UI of my CLI TOOL

Made using html for a simple and smooth user experience!

Here is the preview of my server :
### Starting page of our server `http://localhost:8080/form.html` (form layout)

### Generating all the links after the base url is crawled `http://localhost:8080/form.html`

### quit
Press ` Ctrl + C ` in the terminal to stop the server.