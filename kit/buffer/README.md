# buffer

请使用 

```go
package main

import "github.com/valyala/bytebufferpool"

func Example() {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
}

```
