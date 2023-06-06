#
# rules_go
#

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


#
# rules_pkg
#
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "rules_pkg",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
    ],
    sha256 = "8f9ee2dc10c1ae514ee599a8b42ed99fa262b757058f65ad3c384289ff70c4b8",
)
load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")
rules_pkg_dependencies()


#
# aspect_bazel_lib
#
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "aspect_bazel_lib",
    sha256 = "e3151d87910f69cf1fc88755392d7c878034a69d6499b287bcfc00b1cf9bb415",
    strip_prefix = "bazel-lib-1.32.1",
    url = "https://github.com/aspect-build/bazel-lib/releases/download/v1.32.1/bazel-lib-v1.32.1.tar.gz",
)

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")

aspect_bazel_lib_dependencies()


#
# rules_oci
#

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_oci",
    sha256 = "db57efd706f01eb3ce771468366baa1614b5b25f4cce99757e2b8d942155b8ec",
    strip_prefix = "rules_oci-1.0.0",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.0.0/rules_oci-v1.0.0.tar.gz",
)

load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

rules_oci_dependencies()

load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "LATEST_ZOT_VERSION", "oci_register_toolchains")

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
    # Uncommenting the zot toolchain will cause it to be used instead of crane for some tasks.
    # Note that it does not support docker-format images.
    # zot_version = LATEST_ZOT_VERSION,
)

# pulls base image (must happen in the workspace)
load("@rules_oci//oci:pull.bzl", "oci_pull")

# Each pulled image is made available as an (external) workspace, and can be referenced by its label.
# $ bazel query '//external:*' | grep distroless_base
# Loading: 0 packages loaded
# //external:distroless_base
# //external:distroless_base_linux_amd64
# //external:distroless_base_linux_arm64

# https://github.com/GoogleContainerTools/distroless
oci_pull(
    name = "distroless_base",
    digest = "sha256:73deaaf6a207c1a33850257ba74e0f196bc418636cada9943a03d7abea980d6d",
    image = "gcr.io/distroless/base-debian11",
    platforms = [
        "linux/amd64",
        "linux/arm64",
    ],
)
