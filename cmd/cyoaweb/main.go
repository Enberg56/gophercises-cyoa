package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/Enberg56/gophercises-cyoa"
)

func main() {
	JsonFileName := flag.String("file", "../../gopher.json", "A Json file that consist containing CYOA")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *JsonFileName)

	f, err := os.Open(*JsonFileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonDecoder(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
