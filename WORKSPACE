load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Downloads rules_go@0.39.0 as an `http_archive`
# From: https://github.com/bazelbuild/rules_go/releases
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6b65cb7917b4d1709f9410ffe00ecf3e160edf674b78c54a894471320862184f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.39.0/rules_go-v0.39.0.zip",
    ],
)

# Loads `rules_go` 
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

# https://go.dev/doc/devel/release
# TODO: upgrade to 1.19.7 because of https://github.com/golang/go/commit/639b67ed114151c0d786aa26e7faeab942400703
go_register_toolchains(version = "1.19")
