package plugin

import (
	"os/exec"
	"runtime"
)

type Runtime string
type ComponentType string

const (
	RuntimeDefault Runtime = RuntimeExec
	RuntimePython  Runtime = "python"
	RuntimeDotnet  Runtime = "dotnet"
	RuntimeNodeJS  Runtime = "nodejs"
	RuntimeExec    Runtime = "exec"
	RuntimeJava    Runtime = "java"
)

type RuntimeContext interface {
	Name() Runtime
	Extension() string
	Executable() string
	Args() []string
	Command(path string) *exec.Cmd
}

type runtimeContext struct {
	name       Runtime
	extension  string
	executable string
	args       []string
}

func (r *runtimeContext) Name() Runtime {
	return r.name
}

func (r *runtimeContext) Extension() string {
	return r.extension
}

func (r *runtimeContext) Executable() string {
	return r.executable
}

func (r *runtimeContext) Args() []string {
	return r.args
}

func (r *runtimeContext) Command(path string) *exec.Cmd {
	if r.executable == "" {
		return exec.Command(path, r.args...)
	}
	args := append(r.args, path)
	return exec.Command(r.executable, args...)
}

func NewDotnet() RuntimeContext {
	return &runtimeContext{
		name:       RuntimeDotnet,
		extension:  ".dll",
		executable: "dotnet",
		args:       []string{"run"},
	}
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
func NewExec() RuntimeContext {
	extension := ""
	if isWindows() {
		extension = ".exe"
	}
	return &runtimeContext{
		name:       RuntimeExec,
		extension:  extension,
		executable: "",
		args:       []string{},
	}
}

func NewPython() RuntimeContext {
	return &runtimeContext{
		name:       RuntimePython,
		extension:  ".py",
		executable: "python",
		args:       []string{},
	}
}

func NewNodeJS() RuntimeContext {
	return &runtimeContext{
		name:       RuntimeNodeJS,
		extension:  ".js",
		executable: "node",
		args:       []string{},
	}
}

func NewJava() RuntimeContext {
	return &runtimeContext{
		name:       RuntimeJava,
		extension:  ".jar",
		executable: "java",
		args:       []string{"-jar"},
	}
}

var runtimeContextMap = map[Runtime]RuntimeContext{
	RuntimeDotnet: NewDotnet(),
	RuntimeExec:   NewExec(),
	RuntimeJava:   NewJava(),
	RuntimeNodeJS: NewNodeJS(),
	RuntimePython: NewPython(),
}

func GetRuntimeContext(name Runtime) RuntimeContext {
	ctx, ok := runtimeContextMap[name]
	if ok {
		return ctx
	}
	return runtimeContextMap[RuntimeDefault]
}
