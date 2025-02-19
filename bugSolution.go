```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["a"]
    if ok {
        fmt.Println("Key exists:", value)
    } else {
        fmt.Println("Key does not exist. Returning default value 0")
    }
}
```