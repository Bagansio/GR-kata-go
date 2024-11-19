package test

import (
	"GR-kata-go/pkg/gildedrose"
	"testing"
)

func TestGildedRoseIntegrationAfter2Days(t *testing.T) {
	// Define the initial items as in the main function
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	// Define expected results after two days
	expectedAfterDay2 := []*gildedrose.Item{
		{"+5 Dexterity Vest", 8, 18},
		{"Aged Brie", 0, 2},
		{"Elixir of the Mongoose", 3, 5},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		{"Conjured Mana Cake", 1, 2},
	}

	// Simulate two days of updates
	for day := 0; day < 2; day++ {
		for _, item := range items {
			item.UpdateQuality()
		}
	}

	// Compare the resulting items to the expected state
	for i, item := range items {
		expected := expectedAfterDay2[i]
		if item.Name != expected.Name {
			t.Errorf("Day 2: Expected Name %s, but got %s", expected.Name, item.Name)
		}
		if item.SellIn != expected.SellIn {
			t.Errorf("Day 2: Expected SellIn %d for item %s, but got %d", expected.SellIn, item.Name, item.SellIn)
		}
		if item.Quality != expected.Quality {
			t.Errorf("Day 2: Expected Quality %d for item %s, but got %d", expected.Quality, item.Name, item.Quality)
		}
	}
}

func TestGildedRoseIntegrationAfter7Days(t *testing.T) {
	// Define the initial items as in the main function
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	}

	// Define expected results after 7 days
	expectedAfterDay7 := []*gildedrose.Item{
		{"+5 Dexterity Vest", 3, 13},
		{"Aged Brie", -5, 12},
		{"Elixir of the Mongoose", -2, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 29},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
		{"Conjured Mana Cake", -4, 0},
	}

	// Simulate updates for 7 days
	for day := 0; day < 7; day++ {
		for _, item := range items {
			item.UpdateQuality()
		}
	}

	// Validate the items against the expected results
	for i, item := range items {
		expected := expectedAfterDay7[i]
		if item.Name != expected.Name {
			t.Errorf("Day 7: Expected Name %s, but got %s", expected.Name, item.Name)
		}
		if item.SellIn != expected.SellIn {
			t.Errorf("Day 7: Expected SellIn %d for item %s, but got %d", expected.SellIn, item.Name, item.SellIn)
		}
		if item.Quality != expected.Quality {
			t.Errorf("Day 7: Expected Quality %d for item %s, but got %d", expected.Quality, item.Name, item.Quality)
		}
	}
}
