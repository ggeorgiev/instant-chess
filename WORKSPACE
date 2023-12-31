workspace(name = "com_github_netskope_ep_barn")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.20.7")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies()

http_archive(
    name = "bazel_gomock",
    sha256 = "2da16771642ce7f75a8d620a1029b83ee29b206c6665bb8c92f003b427e35dbf",
    strip_prefix = "bazel_gomock-4f2ee840432b1a08ccc46ee4f2c1f5a2bad8fade",
    url = "https://github.com/jmhodges/bazel_gomock/archive/4f2ee840432b1a08ccc46ee4f2c1f5a2bad8fade.tar.gz",
)

http_archive(
    name = "aspect_bazel_lib",
    sha256 = "c858cc637db5370f6fd752478d1153955b4b4cbec7ffe95eb4a47a48499a79c3",
    strip_prefix = "bazel-lib-2.0.3",
    url = "https://github.com/aspect-build/bazel-lib/releases/download/v2.0.3/bazel-lib-v2.0.3.tar.gz",
)

load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies", "aspect_bazel_lib_register_toolchains")

aspect_bazel_lib_dependencies()

aspect_bazel_lib_register_toolchains()