go_binary(
    name = "boss",
    srcs = [
        "app.go",
        "boss.go",
        "cmd_check.go",
        "cmd_start.go",
    ],
    deps = [
        "//api",
        "//cfg",
        "//server",
        "//third_party/go:com_github_urfave_cli_v2",
        "//third_party/go:org_golang_google_grpc",
    ],
)
