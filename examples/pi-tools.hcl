project "pi-tools" {
  task "presence" {
    cmd = "bazel run //detect-presence/cmd/detect-presence-srv"
  }

  task "go-links" {
    cmd = "bazel run //go-links/cmd/go-links"
  }

  task "homebase:relay" {
    cmd         = "yarn generate:relay --watch"
    working_dir = "homebase"
  }

  task "vault-proxy" {
    cmd = "ibazel run //vault-proxy/cmd/vault-proxy -- -debug -auth-path webauthn-debug -cookie-domain '' -static-dir=$PWD/vault-proxy/static"

    disable_autostart = true
  }
}
