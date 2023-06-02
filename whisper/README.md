## Building Golang with Bazel

Dependencies on the host

```
brew install bazel
brew install golang
brew install colima docker
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

### Bulding a ~~Docker~~ OCI image

[`rules_oci`](https://github.com/bazel-contrib/rules_oci) is the current plugin to build container images, and the previous [`rules_docker`](https://github.com/bazelbuild/rules_docker#status) is in maintenance mode.

Here's where things get a little complicated in terms of the architecture. By default, Bazel uses single-platform builds: the host (where Bazel runs), execution (where the build tools run), and the target (where the final output reside/runs) are the same.

That means `bazel build` will produce a `darwin/arm64` binary when run on a [Mac computer with Apple silicon](https://support.apple.com/en-us/HT211814), a `darwin/amd64` on any Intel-based Mac computer, and a `linux/amd64` on any other x86_64 Linux computer.

Remember to match the platform of your container and your application binary:

```
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64 @//whisper:tarball
colima start --arch aarch64 --cpu 4 --memory 2
docker load --input $(bazel cquery --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64 --output=files @//whisper:tarball)
docker run --rm juniorz/how-bazel/whisper:dev
colima stop
```

and

```
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 @//whisper:tarball
colima start --profile intel --arch x86_64 --cpu 4 --memory 2
docker load --input $(bazel cquery --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 --output=files @//whisper:tarball)
docker run --platform linux/amd64 --rm juniorz/how-bazel/whisper:dev
colima stop --profile intel
```

Note that this model also simplifies many tasks, such as base image management (e.g. upgrades, consistency across packages) and building "cross platform" images (e.g. build `amd64` images from `darwin/arm64` hosts). Also note that you can't generate the tarball for [both target platforms at once](https://github.com/bazelbuild/bazel/issues/6044) because the output-dir is based on the "host" platform.

#### Restricting targets to specific platforms

Container are supported by a wide range of operating systems, as long as they are not macOS :P
For this reason, it is reasonable to restrict the target that generate container images to target platforms in which containers are supported. This can easily be implemented via `target_compatible_with` along with [platform constraints](https://bazel.build/extending/platforms#skipping-incompatible-targets).
