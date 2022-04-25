# mountinfo

The purpose of this library is to read the `mountinfo` file as a whole
or as single lines in order to get information about the mount points of
a specific process as described in `man 5 proc`.


## API

- `ParseMountInfo(buffer io.Reader) ([]Mountinfo, error)`: ParseMountInfo parses the mountinfo content from an io.Reader, e.g a file
- `ParseMountInfoString(tx string) *Mountinfo`:  ParseMountInfoString transforms a mountinfo string in a struct of type Mountinfo
- `GetMountInfo(mountinfoPath string) ([]Mountinfo, error)`: GetMountInfo reads the mountinfo file and returns a slice of structs of type Mountinfo

## Examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/fntlnz/mountinfo"
)

func main() {
	minfo, err := mountinfo.GetMountInfo("/proc/self/mountinfo")
	if err != nil {
		log.Fatal("error getting mountinfo: %v", err)
	}
	fmt.Printf("Mountpoint: %s", minfo[0].MountPoint)
}
```
