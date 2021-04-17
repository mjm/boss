package cfg

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

const DefaultProjectFile = "boss.hcl"

func ReadFile(filename string) (*Cfg, error) {
	p := hclparse.NewParser()
	f, diags := p.ParseHCLFile(filename)
	if diags.HasErrors() {
		return nil, diags
	}

	var cfg Cfg
	if diags := gohcl.DecodeBody(f.Body, nil, &cfg); diags.HasErrors() {
		return nil, diags
	}
	return &cfg, nil
}
