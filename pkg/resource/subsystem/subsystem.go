package subsystem

type System interface {
	Apply() error
	AddPid(pid int) error
	Destroy() error
	Name() string // 子系统的名称，比如memory,cpu等
}