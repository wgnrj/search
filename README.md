# Search functionality #

This is a simple concurrent search package written in Go to search for files in a given directory that contain a given tag.

It contains a Search function and a struct SearchResult, that can be used to store the results and the tag and directory.

## Usage ##

Import the package with `import "github.com/wgnrj/search"` and start searching.

```go
package main

import (
    "fmt"

    "github.com/wgnrj/search"
)

func main() {
    for _, s := Search("data/", "Hello") {
        fmt.Println(s)
    }
}
```

More examples can be seen in the `search_test.go` file, that uses the built-in `testing` functionality of Go.
