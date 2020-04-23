# Exploring Go Modules

A **module** is a collection of related Go packages. Modules are the unit of source code interchange and versioning.

The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

Modules replace the old GOPATH-based approach to specifying which source files are used in a given build.

## Goals of the Module System

1. Keep what works well with Workspaces.

- Retrieving dependencies
- Simplifing build process
- Sharing projects

2. Address the weaknesses of the Workspaces.

- Versioning and API stability
- Vendoring and reproducible builds

## History of Go Project Management

|  Versioning and API stability  | Vendoring and reproducible builds |
| :----------------------------: | :-------------------------------: |
|                                |            goven(2012)            |
| Compatibility guidelines(2013) |            godep(2013)            |
|         gopkg.in(2014)         |                                   |
|   Semantic versioning(2015)    |             gb(2015)              |
|                                |      vendor experiment(2015)      |
|           dep(2016)            |             dep(2016)             |

For further [read](https://research.swtch.com/vgo-intro)

## Overview of Modules

- A **Module** is made up of one or more related packages.
- Modules are configured via `go.mod` file.
- Modules are version controlled.
- Modules expect strict Semantic Versioning.
- Dependent libraries are kept in a cache.
- Integrity checks via checksums.

```
module github.com/{username}/{projectname}
go 1.14

require (
    github.com/gorilla/mux v1.7.3 // indirect
    github.com/gorilla/rpc v1.2.0 // indirect
    github.com/{username}/{projectname}
)

// patching is the way routing is done

replace github.com/gorilla/mux => ../mymods/github.com/fakorede/mux
exclude github.com/gorilla/rpc v1.1.0
```

## Creating and Maintaining Modules

### Creating a Module

```
go mod init github.com/fakorede/learning-golang/go-modules
```

```
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	if err := http.ListenAndServe(":3005", nil); err != nil {
		log.Fatal(err)
	}
}
```

```
go build .
.\go-modules.exe
```

### Retreiving Dependencies

Many applications rely on third-party libraries

```
go get -u github.com/gorilla/mux
```

Updates the `go.mod` file by adding:

```
...
require github.com/gorilla/mux v1.7.4 // indirect
```

Manually update imports statement with:

```
"github.com/gorilla/mux"
```

```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("Topic: " + vars["topic"]))
	})

	http.Handle("/", rtr)

	if err := http.ListenAndServe(":3005", nil); err != nil {
		log.Fatal(err)
	}
}
```

### Listing Dependencies

```
go list
go list all
go list -m all
go list -m -versions {dependency}
```

`{dependency}` ex `github.com/gorilla/mux`

### Verifying Dependencies

We can verify that all of the dependencies included in an application matches those initially retreived with the `go get` command.

`github.com/gorilla/mux v1.7.4 h1:VuZ8uybHlWmqV03+zRzdwKL4tUnIp1MAQtp1mIFE1bc=`

`h1:VuZ8uybHlWmqV03+zRzdwKL4tUnIp1MAQtp1mIFE1bc=` unique hash based on the retreived source code. Used to verify that every one using the specific version gets the same code.

```
go mod verify
```

At build time, we can be sure that everything will be built with the same dependencies the application is developed with.

### Removing Unused Dependencies

```
go mod tidy
```

Rebuilds the dependency graph of the entire application, determines what modules are no longer required and pulls them out.

### Conclusion

That covers the basic workflow if working with a module-based application. The intention of the module sysatem isn't to break anything but to leverage what we've already done well in Go and build upon that.

## Understanding Module Versioning

### Semantic Versioning

This is an example of a full version that can be used for a Go module:

```
v1.5.3-pre1
```

**v** - **Version prefix**(required)
**1** - **Major revision**(likely to break backward compatibility) guarantees that anything compatible with major version 1 should be compatible with anything else that is major version 1.
**5** - **Minor revision**(new features, doesn't break BC) taking features away or changing APIs will break backward compatibility. But adding new features will be indicated by increasing the minor revision.
**3** - **Patch version**(bug fixes, no new features, and doesn't break BC).
**pre1** - **Pre-release identifier for a new version**(if applicable) Pre-releases are used to rank releases and are compared in an alphabetical manner ex alpha < beta, pre1 < pre2. The text is arbitrary and important to establish convention within a project.

https://semver.org to learn more.

### Module Versioning Rules

#### Versioning Rules - v1 and Earlier

- No promise of backward compatibility prior to v1.0.0 - The assumption is that the library is still trying to define itself, its shape, api, features. The api and library could be changing considerably.

- Precedence determined by major, minor, then patch versions.

- The import syntax is going to look like:

```
import github.com/gorilla/mux
```

#### Versioning Rules - v2 and Beyond

- Backward compatibility should be preserved within a major version i.e consistent as minor and patch revisions are being released. If Backward Compatibility has to be broken, the major version should be incremented to indicate that.

- Each major version has a unique import path.

```
import github.com/gorilla/mux/v2
```

**Demo**

go get rsc.io/quote - defaults to v1
go get rsc.io/quote/v2 - v2

#### Versioning Rules - unversioned commits

- Still uses semver

- Relies on prerelease identifier (timestamp and commit hash)

```
require golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
```

### Module Queries

We can execrcise more control over what version of a dependency to be retreived.

- Specific version `@v1.7.2`
- Version prefix `@v1`
- Latest `@latest` (this is the default)
- Specifi commit `@c956192`
- Specific commit `@master` (latest branch)
- Comparison `@>=v1.7.2`

**Comparison Rules**

- Closest match wins
- Prereleases have lower precedence irrespective of version

**Demo**

```
go get github.com/gorilla/mux               ==>     v1.7.3
go get github.com/gorilla/mux@v1.6.1        ==>     v1.6.1
go get github.com/gorilla/mux@master        ==>     v1.6.1
go get "github.com/gorilla/mux@<v1.7.0"     ==>     v1.6.2 // closest match
```

## Advanced Module Management Tools

- go mod why : explores why a module is required.

```
go mod why github.com/gorilla/mux
```

- go mod graph : builds the dependency graph of the application.

```
go mod graph
```

- go mod edit : allows us modify our mod file automatically.

```
go mod edit -module github.com/fakorede/modules
go mod edit -go 1.12
go mod edit -require github.com/gorilla/content@v1.1.1
go mod edit -droprequire github.com/gorilla/content
go mod edit -exclude github.com/gorilla/content
go mod edit -dropexclude github.com/gorilla/content
go mod edit -replace
go mod edit -dropreplace
go mod edit -json
```

To get the list of all the `edit` commands and how they work

```
go mod help edit
```

- go mod download : used to store modules in the Go's build cache. It won't modify `go.mod`

```
go mod download
```

- go mod vendor : analyse project, finds all modules relied upon and populates them in a `vendor` directory.

### Using Vendor Directories

```
go mod vendor
go run -mod=vendor
```

### Building with the Readonly

```
go build -mod=readonly .
```
