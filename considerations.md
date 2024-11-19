# Considerations
- Execution print (item printing format, requires knowing the item data and its position and some noise) example:&{+5 Dexterity Vest 10 20} 
- Using "magic numbers" instead constants 
- A method with multiple functions 
- project structure not a golang standard  
- UpdateQuality function is confuse (updates item qualities) -> Have a struct function 
- Test is poor and not works\ 
- a lot of redundancy (quality < 50) 
- confusing code 

# Alternative solution strategy:

- Constants could be a package named item types and then rename as SULFURAS instead SULFURAS_ITEM_NAME and use like: 

```go
import (
    it `item_types`
)

it.SULFURAS
```

- To avoid redundancy and being more maintainble,  avoided as it is a simple case and it would add some extra logic that currently will be just "noise"

```go
// First solution:
type Item struct {
	Name            string
	SellIn, Quality int
}

type ItemType interface {
	UpdateQuality()
}

type AgedBrie struct {
	Item
}

type Sulfuras struct {
	Item
}

func (item *AgedBrie) UpdateQuality() {
    ....
}

func (item *Sulfuras) UpdateQuality() {}

...

func (item *Item) GetItemType() ItemType {
	switch {
	case strings.Contains(item.Name, AGED_BRIE_ITEM):
		return &AgedBrie{*item}
	case strings.Contains(item.Name, SULFURAS_ITEM):
		return &Sulfuras{*item}
	...
	}
}

func (item *Item) UpdateQuality() {
	// Get the item type dynamically based on the name or properties
	itemType := item.GetItemType()

	// Call the UpdateQuality method specific to the item type
	itemType.UpdateQuality()
}
```