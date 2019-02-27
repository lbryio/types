# Using LBRY types in your code

## Go

```go
package main

import pb "github.com/lbryio/types/go"
import "fmt"

func main() {
    title := "Coherence"
    metadata := pb.Metadata{Title:&title}
    fmt.Printf("Let's watch %s on LBRY!\n", metadata.GetTitle())
}

```

## Python

todo

## Javascript

todo

# Compiling types

You only need to do this if you're modifying the types themselves.

- Download [the protoc binary](https://github.com/google/protobuf/releases) and put it in your path. Make sure you get the one starting with `protoc`, not `protobuf`.
- `./build.sh`