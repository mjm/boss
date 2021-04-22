package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "toolchain",
    strip_srcs = True,
    version = "1.16.3",
)

go_module(
    name = "com_github_urfave_cli_v2",
    module = "github.com/urfave/cli/v2",
    version = "v2.3.0",
    deps = [
        ":com_github_cpuguy83_go-md2man_v2",
    ],
)

go_module(
    name = "com_github_cpuguy83_go-md2man_v2",
    install = [
        "md2man",
    ],
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.0-20190314233015-f79a8a8ca69d",
    deps = [
        ":com_github_russross_blackfriday_v2",
    ],
)

go_module(
    name = "com_github_russross_blackfriday_v2",
    module = "github.com/russross/blackfriday/v2",
    version = "v2.0.1",
    deps = [
        ":com_github_shurcooL_sanitized_anchor_name",
    ],
)

go_module(
    name = "com_github_shurcooL_sanitized_anchor_name",
    module = "github.com/shurcooL/sanitized_anchor_name",
    version = "v1.0.0",
)

go_module(
    name = "com_github_hashicorp_hcl_v2",
    install = [
        ".",
        "ext/customdecode",
        "gohcl",
        "hclparse",
        "hclsyntax",
        "hclwrite",
        "json",
    ],
    module = "github.com/hashicorp/hcl/v2",
    version = "v2.9.1",
    deps = [
        ":com_github_agext_levenshtein",
        ":com_github_apparentlymart_go-textseg_v13",
        ":com_github_google_go-cmp",
        ":com_github_mitchellh_go-wordwrap",
        ":com_github_spf13_pflag",
        ":com_github_zclconf_go-cty",
    ],
)

go_module(
    name = "com_github_mitchellh_go-wordwrap",
    module = "github.com/mitchellh/go-wordwrap",
    version = "v0.0.0-20150314170334-ad45545899c7",
    deps = [
        ":com_github_zclconf_go-cty",
    ],
)

go_module(
    name = "com_github_zclconf_go-cty",
    install = [
        "cty",
        "cty/convert",
        "cty/function",
        "cty/function/stdlib",
        "cty/gocty",
        "cty/json",
        "cty/set",
    ],
    module = "github.com/zclconf/go-cty",
    version = "v1.8.0",
    deps = [
        ":com_github_apparentlymart_go-textseg_v13",
        ":org_golang_x_text",
    ],
)

go_module(
    name = "org_golang_x_text",
    install = [
        "encoding/...",
        "secure/...",
        "transform",
        "unicode/...",
    ],
    module = "golang.org/x/text",
    version = "v0.3.5",
)

go_module(
    name = "com_github_agext_levenshtein",
    module = "github.com/agext/levenshtein",
    version = "v1.2.1",
)

go_module(
    name = "com_github_apparentlymart_go-textseg_v13",
    install = [
        "textseg",
    ],
    module = "github.com/apparentlymart/go-textseg/v13",
    version = "v13.0.0",
)

go_module(
    name = "com_github_google_go-cmp",
    install = [
        "cmp/...",
    ],
    module = "github.com/google/go-cmp",
    version = "v0.3.1",
)

go_module(
    name = "com_github_spf13_pflag",
    module = "github.com/spf13/pflag",
    version = "v1.0.2",
)

go_module(
    name = "org_golang_x_crypto",
    install = [
        "ssh/terminal",
    ],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20190426145343-a29dc8fdc734",
    deps = [
        ":org_golang_x_sys",
    ],
)

go_mod_download(
    name = "protobuf_download",
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
)

go_module(
    name = "com_github_golang_protobuf",
    download = ":protobuf_download",
    exported_deps = [
        ":org_golang_google_protobuf",
    ],
    install = ["..."],
    module = "github.com/golang/protobuf",
    strip = [
        "proto/proto3_proto",
        "conformance",
    ],
)

go_module(
    name = "org_golang_google_protobuf",
    install = ["..."],
    module = "google.golang.org/protobuf",
    version = "v1.25.0",
    deps = [
        ":com_github_google_go-cmp",
    ],
)

go_module(
    name = "org_golang_google_grpc",
    exported_deps = [
        ":org_golang_x_net",
        ":org_golang_google_protobuf",
    ],
    install = [
        ".",
        "attributes",
        "backoff",
        "balancer",
        "balancer/base",
        "balancer/grpclb/...",
        "balancer/roundrobin",
        "binarylog/...",
        "codes",
        "connectivity",
        "credentials",
        "credentials/oauth",
        "encoding",
        "encoding/gzip",
        "encoding/proto",
        "grpclog",
        "health",
        "health/grpc_health_v1",
        "internal/...",
        "keepalive",
        "metadata",
        "peer",
        "resolver",
        "resolver/dns",
        "resolver/passthrough",
        "serviceconfig",
        "stats",
        "status",
        "tap",
    ],
    module = "google.golang.org/grpc",
    version = "v1.34.0",
    deps = [
        ":com_github_google_go-cmp",
        ":genproto_rpc",
        ":org_golang_x_oauth2",
        ":com_github_golang_protobuf",
        ":com_github_google_uuid",
        ":org_golang_x_sys",
    ],
)

go_module(
    name = "protoc-gen-go",
    binary = True,
    download = ":protobuf_download",
    install = ["protoc-gen-go"],
    module = "github.com/golang/protobuf",
    deps = [":org_golang_google_protobuf"],
)

go_module(
    name = "protoc-gen-go-grpc",
    binary = True,
    module = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
    version = "v1.1.0",
    deps = [
        ":org_golang_google_grpc",
        ":com_github_golang_protobuf",
    ],
)

go_mod_download(
    name = "genproto_download",
    module = "google.golang.org/genproto",
    version = "v0.0.0-20210315173758-2651cd453018",
)

# genproto stuff for grpc to resolve cyclic deps
go_module(
    name = "genproto_rpc",
    download = ":genproto_download",
    install = [
        "googleapis/rpc/...",
    ],
    module = "google.golang.org/genproto",
    deps = [
        ":com_github_golang_protobuf",
    ],
)

go_module(
    name = "org_golang_x_oauth2",
    install = [
        ".",
        "google",
        "internal",
        "jws",
        "jwt",
    ],
    module = "golang.org/x/oauth2",
    version = "d668ce993890a79bda886613ee587a69dd5da7a6",
    visibility = ["//third_party/go/..."],
    deps = [
        ":com_google_cloud_go",
        ":org_golang_x_net",
    ],
)

go_module(
    name = "com_google_cloud_go",
    install = ["compute/metadata"],
    module = "cloud.google.com/go",
    version = "v0.40.0",
)

go_module(
    name = "org_golang_x_net",
    install = ["..."],
    module = "golang.org/x/net",
    version = "136a25c244d3019482a795d728110278d6ba09a4",
    deps = [
        ":org_golang_x_text",
        ":org_golang_x_crypto",
    ],
)

go_module(
    name = "com_github_google_uuid",
    module = "github.com/google/uuid",
    version = "v1.1.1",
)

go_module(
    name = "org_golang_x_sys",
    install = ["..."],
    module = "golang.org/x/sys",
    version = "765f4ea38db36397e827c4153018aa272eed7835",
)
