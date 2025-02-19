```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, it will print 0
}
```