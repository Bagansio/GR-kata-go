package main

import (
	g "GR-kata-go/pkg/gildedrose"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*g.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		for i := 0; i < len(items); i++ {
			fmt.Printf("%+v\n", *items[i])
		}
		fmt.Println("")
		for i := 0; i < len(items); i++ {
			items[i].UpdateQuality()
		}
	}
}