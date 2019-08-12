# version

The package `version` is used to print version information for the program. The version information can be passed to the package during the build time.

```make
VERSION   ?= $(shell git describe --tags)
REVISION  ?= $(shell git rev-parse HEAD)
BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDUSER ?= $(shell id -un)
BUILDTIME ?= $(shell date '+%Y-%m-%d@%H:%M:%S')

.PHONY: build

build:
	go build -ldflags "-X github.com/ricoberger/gocommon/version.Version=${VERSION} \
		-X github.com/ricoberger/gocommon/version.Revision=${REVISION} \
		-X github.com/ricoberger/gocommon/version.Branch=${BRANCH} \
		-X github.com/ricoberger/gocommon/version.BuildUser=${BUILDUSER} \
		-X github.com/ricoberger/gocommon/version.BuildDate=${BUILDTIME}" \
		-o ./bin/myapp ./cmd/myapp;
```

The package can then be used as follows:

```go
import (
	"fmt"

	"github.com/ricoberger/gocommon/version"
)

func main() {
	v, err := version.Print("myapp")
	if err != nil {
		return
	}

	fmt.Println(v)

	fmt.Println(version.Info())
	fmt.Println(version.BuildContext())
}
```

