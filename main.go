package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/mmgfrcs/yugioh-analyzer/models"
	"github.com/mmgfrcs/yugioh-analyzer/parsers"
)

func main() {
	var result models.YuGiOhAPIResult
	file, err := os.ReadFile("cache.json")

	if os.IsNotExist(err) {
		log.Printf("Cache does not exist. Downloading through the API.")

		res, err := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php?name=Marincess Blue Tang")
		if err != nil {
			log.Fatalf("Error in API request: %s\n", err.Error())
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Error in Body Parsing: %s\n", err.Error())
		}
		os.WriteFile("cache.json", body, 0777)

		if err = json.Unmarshal(body, &result); err != nil {
			log.Fatalf("Error in JSON decode: %s\n", err.Error())
		}

	} else {
		log.Printf("Cache exists. Using cache.")
		if err = json.Unmarshal(file, &result); err != nil {
			log.Fatalf("Error in JSON decode: %s\n", err.Error())
		}
	}

	log.Printf("[%s] %s\n%s", result.Cards[0].Name, result.Cards[0].Type, result.Cards[0].Description)

	parser, err := participle.Build[parsers.YuGiOhDescription]()
	if err != nil {
		log.Fatalf("Error in Building Parser: %s\n", err.Error())
	}
	if err := os.WriteFile("parser.ebnf", []byte(parser.String()), 0777); err != nil {
		log.Fatalf("Error in writing parser: %s\n", err.Error())
	}
	res, err := parser.ParseString(result.Cards[0].Name, result.Cards[0].Description)
	if err != nil {
		log.Fatalf("Error in Parsing: %s\n", err.Error())
	}
	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error in JSON conversion after Parse: %s\n", err.Error())
	}
	log.Printf(string(b))

}
