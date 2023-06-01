## Building Golang with Bazel

Dependencies on the host

```
brew install bazel
brew install golang
```

### Creating a Golang project

From https://go.dev/doc/code

```
mkdir whisper && cd whisper
# NOTE: this creates a Golang module with the Go version on my host :P
go mod init github.com/juniorz/how-bazel/whisper
<<-EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}
EOF
echo whisper > .gitignore
```

Building w/o Bazel is simple

### Building with my host's toolchain

```
go build && ./whisper
```
