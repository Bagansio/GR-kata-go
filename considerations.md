# Considerations
- Execution print (item printing format, requires knowing the item data and its position and some noise) example:&{+5 Dexterity Vest 10 20}
- Using "magic numbers" instead constants
- A method with multiple functions
- project structure not a golang standard 
- UpdateQuality function is confuse (updates item qualities) -> Have a struct function
- Test is poor and not works\
- a lot of redundancy (quality < 50)
- confusing code

#

- Constants could be a package named item types and then rename as SULFURAS instead SULFURAS_ITEM_NAME and use like: 

```go
import (
    it `item_types`
)

it.SULFURAS
```
