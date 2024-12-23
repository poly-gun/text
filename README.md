# `text` - String Utilities Package

The purpose of `text` is to provide string utilities, but with special considerations relating to web-service(s),
cloud providers, and external APIs.

Interfacing with packages such as AWS and Stripe GO SDKs often require working with pointer inputs and outputs. Such
cases can result in undesired, unexpected behavior. When writing API services, tracing issues can become difficult without
warnings or general logging.

That's where the `text` package's [variadic options](./options.go) are perhaps the most notable; configurations
can be set that will log unexpected nil or zeroth input arguments.

## Documentation

Official `godoc` documentation (with examples) can be found at the [Package Registry](https://pkg.go.dev/github.com/poly-gun/text).

## Usage

###### Add Package Dependency

```bash
go get -u github.com/poly-gun/text
```

###### Import & Implement

`main.go`

```go
package main

import (
    "github.com/poly-gun/text"
)

func main() {
    const example = "v-value"

    ptr := text.Pointer(example) // initialize ptr as a reference to example

    // --> optionally construct text.Options for logging
    ptr = text.Pointer(example, text.Variadic(o text.Options) {
        o.Log = true // output a warning if string is empty
    })

    ...
}
```

- Please refer to the [code examples](./example_test.go) for additional usage and implementation details.
- See https://pkg.go.dev/github.com/poly-gun/text for additional documentation.

## Contributions

See the [**Contributing Guide**](./CONTRIBUTING.md) for additional details on getting started.
