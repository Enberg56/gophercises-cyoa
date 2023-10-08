package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	cyoa "github.com/Enberg56/gophercises-cyoa"
)

func main() {
	JsonFileName := flag.String("file", "gopher.json", "A Json file that consist containing CYOA")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *JsonFileName)

	f, err := os.Open(*JsonFileName)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

	// var jsonFileContent []story

	// err := json.Unmarshal(JsonFileName, &animals)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Printf("%+v", animals)
}
