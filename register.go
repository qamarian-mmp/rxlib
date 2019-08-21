package rxlib

func NewRegister (id string, dep []string, startupFunc func (Key)) (*Register) {
	return &Register {id, dep, startupFunc}
}

// Register is a data type that could be used to register a main with rexa.
type Register struct {
	id string // The ID of the main.
	dep []string // The IDs of the main's dependencies.
	startupFunc func (Key) // The startup function of the main.
}

func (r *Register) ID () (string) {
	return r.id
}

func (r *Register) Dep () ([]string) {
	return r.dep
}

func (r *Register) StartupFunc () (func (Key)) {
	return r.startupFunc
}
