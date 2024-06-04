package mount

type Mounter interface {
	Mount() error
	UnMount() error
}