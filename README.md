# Go Map Access Without Panic

This example demonstrates a potential pitfall in Go when dealing with maps: accessing a non-existent key does not result in a runtime panic. Instead, it returns the zero value for the map's value type.

This behavior can be both a feature and a source of unexpected behavior.  It's efficient because it avoids the overhead of checking for key existence, but it can lead to hard-to-debug issues if you're not carefully considering the possibility of missing keys.

## How to reproduce

Simply run the `bug.go` file. You will observe that attempting to access the key "a" in an uninitialized map prints 0.  This is the zero value for an integer, not an error indication.

## How to fix

The solution depends on your desired behavior.  If you want to explicitly check for key existence, use the `value, ok := map[key]` idiom:

```go
value, ok := m["a"]
if ok {
    // Key exists, use value
} else {
    // Key does not exist, handle accordingly
}
```

Alternatively, if you want to provide a default value when a key is missing, you can use the `map.Load` function, available for `sync.Map` type.

This example highlights the importance of defensive programming when working with maps in Go, ensuring you correctly handle scenarios where keys might be absent.