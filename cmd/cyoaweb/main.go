package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/philip314/cyoa"
)

func main() {
	portFlag := flag.Int("port", 8080, "Port that application runs on")
	filename := flag.String("file", "gopher.json", "JSON file of story")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit("Error opening file")
	}

	story, err := cyoa.CreateStory(file)
	if err != nil {
		exit(fmt.Sprintf("Error creating story from JSON file, %s", err))
	}
	storyTemplate := template.Must(template.New("").Parse(htmlTemplateStoryPrefix))

	handler := cyoa.StoryHandler(story, cyoa.UseTemplate(storyTemplate), cyoa.UsePathFunc(storyPathParseFunc))
	mux := http.NewServeMux()
	mux.Handle("/story/", handler)

	port := fmt.Sprintf(":%d", *portFlag)

	fmt.Printf("Starting server on port%s\n", port)

	log.Fatal(http.ListenAndServe(port, mux))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Function to handle paths with /story/ prefix
func storyPathParseFunc(r *http.Request) string {
	path := r.URL.Path
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// Template with /story/ prefix on options
var htmlTemplateStoryPrefix = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8"/>
		<title>
			Choose Your Own Adventure
		</title>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>
			{{range .Story}}
				<p>{{.}}</p>
			{{end}}
			<ul>
			{{range .Options}}
				<li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
			<ul>
		</section>
	</body>

	<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align:center;
			position:relative;
		}
		.page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FCF6FC;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #797;
		}
		
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		
		li {
			padding-top: 10px;
		}
		
		a,
		a:visited {
			text-decoration: underline;
			color: #555;
		}
		
		a:active,
		a:hover {
			color: #222;
		}
		
		p {
			text-indent: 1em;
		}
	</style>
</html>`
