package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/conradludgate/tree"
)

func main() {
	tree.HandleJSON(nil)
	tree.HandleTXT(nil)

	data, err := tree.GenerateTree("game", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Raw format:")
	fmt.Printf("%v\n", data)

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Encoded as JSON:")
	fmt.Println(string(b))

	fmt.Println()
	fmt.Println("Get entities/monster/lore:")
	lore, err := tree.Get(data, "entities/monster/lore")
	if err != nil {
		log.Println("Lore does not exist")
	} else {
		fmt.Println(lore)
	}
}
