package gildedrose

import (
	"testing"
)

// Generalized helper function to check for errors in Quality or SellIn
func testErrorGildedRose(t *testing.T, expected, got int, field string) {
	if got != expected {
		t.Errorf("Expected %s to be %d but got %d", field, expected, got)
	}
}

func TestNormalItemDecreasesQualityAndSellIn(t *testing.T) {
	item := Item{Name: NORMAL_ITEM, SellIn: 10, Quality: 20}
	item.UpdateQuality()

	testErrorGildedRose(t, 19, item.Quality, "Quality")
	testErrorGildedRose(t, 9, item.SellIn, "SellIn")
}

func TestNormalItemQualityDoesNotGoBelowZero(t *testing.T) {
	item := Item{Name: NORMAL_ITEM, SellIn: 1, Quality: 0}
	item.UpdateQuality()

	testErrorGildedRose(t, 0, item.Quality, "Quality")
	testErrorGildedRose(t, 0, item.SellIn, "SellIn")
}

func TestNormalItemDegradesTwiceAsFastAfterSellIn(t *testing.T) {
	item := Item{Name: NORMAL_ITEM, SellIn: 0, Quality: 10}
	item.UpdateQuality()

	testErrorGildedRose(t, 8, item.Quality, "Quality")
	testErrorGildedRose(t, -1, item.SellIn, "SellIn")
}

func TestNormalItemWithMinQualityAfterSellIn(t *testing.T) {
	item := Item{Name: NORMAL_ITEM, SellIn: 0, Quality: 1}
	item.UpdateQuality()

	testErrorGildedRose(t, 0, item.Quality, "Quality")
	testErrorGildedRose(t, -1, item.SellIn, "SellIn")
}

func TestAgedBrieIncreasesQualityBeforeSellIn(t *testing.T) {
	item := Item{Name: AGED_BRIE_ITEM, SellIn: 5, Quality: 10}
	item.UpdateQuality()

	testErrorGildedRose(t, 11, item.Quality, "Quality")
	testErrorGildedRose(t, 4, item.SellIn, "SellIn")
}

func TestAgedBrieIncreasesQualityAfterSellIn(t *testing.T) {
	item := Item{Name: AGED_BRIE_ITEM, SellIn: 0, Quality: 10}
	item.UpdateQuality()

	testErrorGildedRose(t, 12, item.Quality, "Quality")
	testErrorGildedRose(t, -1, item.SellIn, "SellIn")
}

func TestAgedBrieQualityMaxThreshold(t *testing.T) {
	item := Item{Name: AGED_BRIE_ITEM, SellIn: -1, Quality: 49}
	item.UpdateQuality()

	testErrorGildedRose(t, 50, item.Quality, "Quality")
	testErrorGildedRose(t, -2, item.SellIn, "SellIn")
}

func TestAgedBrieQualityMaxThreshold2(t *testing.T) {
	item := Item{Name: AGED_BRIE_ITEM, SellIn: 2, Quality: 50}
	item.UpdateQuality()

	testErrorGildedRose(t, 50, item.Quality, "Quality")
	testErrorGildedRose(t, 1, item.SellIn, "SellIn")
}

func TestBackstagePassesIncreaseQuality(t *testing.T) {
	item := Item{Name: BACKSTAGE_PASSES_ITEM, SellIn: 15, Quality: 20}
	item.UpdateQuality()

	testErrorGildedRose(t, 21, item.Quality, "Quality")
	testErrorGildedRose(t, 14, item.SellIn, "SellIn")
}

func TestBackstagePassesIncreaseBy2WhenSellIn10OrLess(t *testing.T) {
	item := Item{Name: BACKSTAGE_PASSES_ITEM, SellIn: 10, Quality: 20}
	item.UpdateQuality()

	testErrorGildedRose(t, 22, item.Quality, "Quality")
	testErrorGildedRose(t, 9, item.SellIn, "SellIn")
}

func TestBackstagePassesIncreaseBy3WhenSellIn5OrLess(t *testing.T) {
	item := Item{Name: BACKSTAGE_PASSES_ITEM, SellIn: 5, Quality: 20}
	item.UpdateQuality()

	testErrorGildedRose(t, 23, item.Quality, "Quality")
	testErrorGildedRose(t, 4, item.SellIn, "SellIn")
}

func TestBackstagePassesDropToZeroAfterConcert(t *testing.T) {
	item := Item{Name: BACKSTAGE_PASSES_ITEM, SellIn: 0, Quality: 20}
	item.UpdateQuality()

	testErrorGildedRose(t, 0, item.Quality, "Quality")
	testErrorGildedRose(t, -1, item.SellIn, "SellIn")
}

func TestBackstagePassesMaxThreshold(t *testing.T) {
	item := Item{Name: BACKSTAGE_PASSES_ITEM, SellIn: 3, Quality: 50}
	item.UpdateQuality()

	testErrorGildedRose(t, 50, item.Quality, "Quality")
	testErrorGildedRose(t, 2, item.SellIn, "SellIn")
}

func TestSulfurasNeverChanges(t *testing.T) {
	item := Item{Name: SULFURAS_ITEM, SellIn: 0, Quality: 80}
	item.UpdateQuality()

	testErrorGildedRose(t, 80, item.Quality, "Quality")
	testErrorGildedRose(t, 0, item.SellIn, "SellIn")
}

func TestConjuredItemsDegradeTwiceAsFast(t *testing.T) {
	item := Item{Name: CONJURED_ITEM, SellIn: 5, Quality: 10}
	item.UpdateQuality()

	testErrorGildedRose(t, 8, item.Quality, "Quality")
	testErrorGildedRose(t, 4, item.SellIn, "SellIn")
}

func TestConjuredItemsDegradeFourTimesAfterSellIn(t *testing.T) {
	item := Item{Name: CONJURED_ITEM, SellIn: 0, Quality: 10}
	item.UpdateQuality()

	testErrorGildedRose(t, 6, item.Quality, "Quality")
	testErrorGildedRose(t, -1, item.SellIn, "SellIn")
}

func TestConjuredItemsQualityDoesNotGoBelowZero(t *testing.T) {
	item := Item{Name: CONJURED_ITEM, SellIn: 5, Quality: 1}
	item.UpdateQuality()

	testErrorGildedRose(t, 0, item.Quality, "Quality")
	testErrorGildedRose(t, 4, item.SellIn, "SellIn")
}
