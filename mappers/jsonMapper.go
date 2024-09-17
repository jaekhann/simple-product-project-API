package mappers

import (
	"awesomeProject/model"
	"encoding/json"
	"log"
)

func ToJson(product model.Product) []byte {
	indent, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return indent
}

func ToJsons(products []model.Product) []byte {
	indent, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return indent
}

func FromJson(str string) model.Product {
	var p model.Product
	valid := json.Valid([]byte(str))
	if !valid {
		log.Fatalf("Invalid json: %s", str)
	}

	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
