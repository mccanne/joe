# joe
JSON object extractor for Go

# example

```go
package main

import (
	"fmt"
	"github.com/mccanne/joe"
)

//
// a JSON object that could represent a thing like 'foo(123, bar, null) + 1.5'
//
const ast = `
    {"type": "Add",
     "left": {
         "type": "FunctionCall",
         "function": "foo",
         "params": [{"type": "Constant", "value": 123 },
                    {"type": "Variable", "name": "bar" },
                    {"type": "Constant", "value": null } ] },
     "right": {"type": "Constant", "value": 1.5 } }`

func main() {
	object, _ := joe.Parse([]byte(ast))
	v1, _ := object.Get("left").Get("params").Index(1).Get("name").String()
	fmt.Println(v1)
	v2, _ := object.Get("right").Get("value").Number()
	fmt.Println(v2)
	v3 := object.Get("left").Get("params").Index(2).Get("value")
	fmt.Println(v3.IsNull())
	undefined := object.Get("xxx")
	fmt.Println(undefined.IsNull())
	fmt.Println(undefined.IsUndefined())
	fmt.Println(undefined.Get("anything").IsUndefined())
	fmt.Println(undefined.Index(100).IsUndefined())
}
```
