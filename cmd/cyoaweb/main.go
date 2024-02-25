package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application")
	JsonFileName := flag.String("file", "../../gopher.json", "A Json file that consist containing CYOA")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *JsonFileName)

	f, err := os.Open(*JsonFileName)
	if err != nil {
		panic(err)
	}

	story, err := JsonDecoder(f)
	if err != nil {
		panic(err)
	}

	h := NewHandler(story)
	fmt.Printf("Starting the server on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

var defaultHandelerTmpl = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>
`

func JsonDecoder(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandelerTmpl))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
