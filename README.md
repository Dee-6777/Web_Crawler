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
First clone the repository somewhere in your $PATH. A common place would be within your $GOPATH. 

### Option 1: Install binary

Build and copy `Web_Crawler` to your $GOPATH/bin:

```
$ make install
```
### Option 2: Build from source

This command will build the Web Crawler binary, named `Web_Crawler`. This binary will be created in the root of your project folder.

```
$ make build
```

### HOW TO USE THIS SERVER? (WITHOUT DOCKERFILE)
* First Step: `Start a server at the port :8080 by running this command`
```
go run main.go 
```
OR 
```
Web_Crawler
```

* Second Step: `Paste this url in your browser and fill the form and submit. It will take a while to load and all the links will be generated on the screen`
```
http://localhost:8080/form.html
```
### HOW TO USE THIS SERVER? (WITH DOCKERFILE)
* First Step: `Build the dockerfile`
```
docker build -t <image-name>
```
* Second Step: `Run the dockerfile`
```
docker run -d -p 8080:8080 <image name>
```
* Third Step: `Paste this url in your browser and fill the form and submit. It will take a while to load and all the links will be generated on the screen`
```
http://localhost:8080/form.html
```

## UI of my CLI TOOL

Made using html for a simple and smooth user experience!

Here is the preview of my server :
### Starting page of our server 
`http://localhost:8080/form.html` (form layout)
![Screenshot from 2023-05-01 08-28-18](https://user-images.githubusercontent.com/73513838/235395539-94ebee98-dbc7-4550-a328-a9437f9b6c2c.png)

### Generating all the links after the base url is crawled 
`http://localhost:8080/form`
![Screenshot from 2023-05-01 08-29-11](https://user-images.githubusercontent.com/73513838/235395560-523d1cc4-f8a0-406a-a472-b944752599bc.png)

### quit
Press ` Ctrl + C ` in the terminal to stop the server.
