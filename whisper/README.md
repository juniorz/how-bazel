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

### Building with Bazel

#### Bazel data model

> Bazel builds software from source code organized in a directory tree called a workspace. Source files in the workspace are organized in a nested hierarchy of packages, where each package is a directory that contains a set of related source files and one `BUILD` file.
> Source: https://bazel.build/concepts/build-ref

1. Workspace has many repositories
1. Workspace has many packages
1. Package has many targets, defined on a `BUILD` file
1. Target has many dependencies
1. Package has many files, but each file belongs to at most one package

#### Building with `rules_go`

Bazel's plugin mechanism allows to extend the rules available in a workspace via "rulesets". In this way, language-specific toolchains and their operation can live outside of the Bazel project.

To build `golang` packages one must:

1. Add `rules_go` to the `WORKSPACE`
1. Configure a `golang` toolchain
1. Use rules to configure targets for each package

At this point, you will be able to `build`, `test` and `run` targets:

```
bazel build @//whisper:whisper
bazel test @//whisper:whisper_test --test_summary=detailed --test_output=all
bazel run @//whisper:whisper
```

To execute a command, Bazel will download the required plugins, the Golang SDK ("distribution"), [configure the toolchain based on the target platform](https://github.com/bazelbuild/rules_go/blob/master/go/toolchains.rst), and finally execute a specific action for each target.

This model simplifies A LOT many tasks, such as Golang version upgrades, supporting multiple host/target platforms (e.g. darwin/arm64, linux/amd64), and even providing a consistent build environment for both development and continuous integration.

Note that every file that is used to generate a target [MUST be explicitly added as a dependency](https://bazel.build/concepts/dependencies#actual-and-declared-dependencies).

#### Labels

Each target is uniquely identified by a **label**. `@//whisper:whisper` is a fully qualified label for the (file?) target `whisper` in the `whisper` package in the main repository (the current workspace). Labels can be relative (within a package).

There's [many ways to specify target patterns](https://bazel.build/run/build#specifying-build-targets), and certain labels are conveniently equivalent:

```
bazel run @//whisper:whisper
bazel run //whisper:whisper
bazel run //whisper
```

NOTE: `//whisper` can also (confusingly) refer to a package name when used in `package_group`, but can NEVER refer to all targets in the same package when used in `BUILD`.
