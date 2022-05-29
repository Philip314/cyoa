package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

func init() {
	htmlTemplate = template.Must(template.New("").Parse(defaultHtmlTemplate))
}

var htmlTemplate *template.Template

var defaultHtmlTemplate = `
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
				<li><a href="/{{.Arc}}">{{.Text}}</a></li>
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

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Construct a story object from a JSON file
func CreateStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func StoryHandler(s Story, options ...HandlerOption) http.Handler {
	storyHandler := storyHandler{s, htmlTemplate}
	for _, v := range options {
		v(&storyHandler)
	}
	return storyHandler
}

type storyHandler struct {
	story    Story
	template *template.Template
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := s.story[path]; ok {
		err := s.template.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

type HandlerOption func(*storyHandler)

func UseTemplate(template *template.Template) HandlerOption {
	return func(sh *storyHandler) {
		sh.template = template
	}
}
