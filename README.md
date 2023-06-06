# How to Bazel

My goal is to learn Bazel to understand its value proposition, encouraged/required patterns, requirements[^runway], limitations and trade-offs.

### Backlog

- run goimports and gofmt as a check and as a task?
  - what's the pattern for targets that change files in the project?
  - `bazel run @io_bazel_rules_go//go -- fmt .` ?
  - https://bazel.build/reference/be/shell ?
  - https://stackoverflow.com/questions/53163137/how-to-properly-invoke-executable-from-usr-bin-via-starlark
  - https://chromium.googlesource.com/external/github.com/bazelbuild/buildtools/+/HEAD/buildifier/README.md
- "how to write rules": crash course
  - https://jayconrod.com/posts/106/writing-bazel-rules--simple-binary-rule
  - https://www.jayconrod.com/posts/111/writing-bazel-rules--platforms-and-toolchains
  - https://www.stevenengelhardt.com/2020/11/19/practical-bazel-wrapping-run-targets-to-provide-additional-context/
- editor support a.k.a can Bazel also provide IDE tooling?
  - https://github.com/bazelbuild/rules_go/wiki/Editor-setup
  - https://jayconrod.com/posts/125/go-editor-support-in-bazel-workspaces
  - https://github.com/bazelbuild/rules_go/wiki/Editor-and-tool-integration
- learning the basic rules
  - e.g. `genrule` (https://bazel.build/reference/be/general#genrule)
  - https://github.com/bazelbuild/bazel-skylib/blob/main/docs/run_binary_doc.md
- How to manage Bazel itself? (https://github.com/bazelbuild/bazelisk vs a Bazel "dev" container)
  - is not a container required for running Bazel RE anyways?
- How (and when) to use Bzlmod?
- Bazel for infrastructure resources: does it fly?
  - https://github.com/tmc/rules_helm and https://github.com/jmileson/rules_terraform LoL

### Takeaways (so far)

- There are many quality of life improvements for both "tools maintainers" and "feature developers".
- There are many caveats and limitations and corner cases when compared to the native toolchain of each language. Even a seasoned developer on a given language platform will have to learn the Bazel ways.

[^runway]: What do I need to adopt (certain features) of Bazel? How tall shall I be?