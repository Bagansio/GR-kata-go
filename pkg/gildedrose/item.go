package gildedrose

import (
	"strings"
)

// Item types
const (
	AGED_BRIE_ITEM        = "Aged Brie"
	BACKSTAGE_PASSES_ITEM = "Backstage passes"
	SULFURAS_ITEM         = "Sulfuras"
	CONJURED_ITEM         = "Conjured"
	NORMAL_ITEM           = "Normal"
)

// Thresholds
const (
	MIN_QUALITY = 0
	MAX_QUALITY = 50

	EXPIRATION_SELL_IN = 0
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) GetItemType() string {
	if strings.Contains(i.Name, AGED_BRIE_ITEM) {
		return AGED_BRIE_ITEM
	}
	if strings.Contains(i.Name, SULFURAS_ITEM) {
		return SULFURAS_ITEM
	}
	if strings.Contains(i.Name, BACKSTAGE_PASSES_ITEM) {
		return BACKSTAGE_PASSES_ITEM
	}
	if strings.Contains(i.Name, CONJURED_ITEM) {
		return CONJURED_ITEM
	}
	return NORMAL_ITEM
}

func (item *Item) UpdateQuality() {
	itemType := item.GetItemType()
	switch itemType {
	case AGED_BRIE_ITEM:
		item.AgedBrieUpdateQuality()
	case BACKSTAGE_PASSES_ITEM:
		item.BackstagePassesUpdateQuality()
	case SULFURAS_ITEM:
		item.SulfurasUpdateQuality()
	case CONJURED_ITEM:
		item.ConjuredUpdateQuality()
	default:
		item.NormalUpdateQuality()
	}
}

func (item *Item) AgedBrieUpdateQuality() {

	item.Quality++

	if item.SellIn <= EXPIRATION_SELL_IN {
		item.Quality++
	}

	if item.Quality > MAX_QUALITY {
		item.Quality = MAX_QUALITY
	}

	item.SellIn--
}

func (item *Item) BackstagePassesUpdateQuality() {

	if item.SellIn <= EXPIRATION_SELL_IN {
		item.Quality = MIN_QUALITY
	} else if item.SellIn <= 5 {
		item.Quality += 3
	} else if item.SellIn <= 10 {
		item.Quality += 2
	} else {
		item.Quality++
	}

	if item.Quality > MAX_QUALITY {
		item.Quality = MAX_QUALITY
	}

	item.SellIn--
}

func (item *Item) SulfurasUpdateQuality() {
	// No need to update anything
}

func (item *Item) ConjuredUpdateQuality() {

	qualityToSubs := 2
	if item.SellIn <= EXPIRATION_SELL_IN {
		qualityToSubs = 4
	}

	item.Quality -= qualityToSubs

	if item.Quality < MIN_QUALITY {
		item.Quality = MIN_QUALITY
	}

	item.SellIn--
}

/*
There are different approach of this depending of efficiency needs or maintainability
*/
func (item *Item) NormalUpdateQuality() {

	qualityToSubs := 1
	if item.SellIn <= EXPIRATION_SELL_IN {
		qualityToSubs = 2
	}

	item.Quality -= qualityToSubs

	if item.Quality < MIN_QUALITY {
		item.Quality = MIN_QUALITY
	}

	item.SellIn--
}
