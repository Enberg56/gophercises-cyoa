package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
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

	fmt.Printf("%+v\n", story)
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
