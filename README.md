ensurecp
========
Recursive file copying with checksum checking.

# Example

```
package main

import (
	"github.com/marhi/ensurecp"
	"log"
)

func main() {
	err := ensurecp.RCopy("/home/marhi/Work", "/tmp/gotest2")
	if err != nil {
		log.Fatalln(err)
	}
}
```
