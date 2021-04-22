package cfg

type Cfg struct {
	Project *Project `hcl:"project,block"`
}

type Project struct {
	Name  string  `hcl:"name,label"`
	Tasks []*Task `hcl:"task,block"`
}

type Task struct {
	Name             string `hcl:"name,label"`
	Command          string `hcl:"cmd,attr"`
	WorkingDir       string `hcl:"working_dir,optional"`
	DisableAutostart bool   `hcl:"disable_autostart,optional"`
}
