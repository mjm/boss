package cfg

type Cfg struct {
	Project *Project `hcl:"project,block"`
}

type Project struct {
	Name string `hcl:"name,label"`
}
