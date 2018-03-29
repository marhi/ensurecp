ensurecp
========
Recursive file copying with checksum checking and JSON log output.

# Example

```
package main

import (
	"github.com/marhi/ensurecp"
	"log"
	"fmt"
)

func main() {
	ensurecp.SetLogging(true)
	err := ensurecp.RCopy("/home/marhi/Work", "/tmp/gotest2")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(ensurecp.ExportLog())
}
```
