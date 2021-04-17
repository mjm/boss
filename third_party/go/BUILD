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
        "transform",
        "unicode/norm",
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
)
