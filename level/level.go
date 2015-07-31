package initcontainerlevel


type Level interface {
	Register(func()) error
	Init() error
}


type internalLevel struct {
	fns []func()
}


func New() Level {
	const fnsInitLen = 8
	fns := make([]func(), 0, fnsInitLen)

	level := internalLevel{
		fns: fns,
	}

	return &level
}


func (lvl *internalLevel) Register(fn func()) error {
	lvl.fns = append(lvl.fns, fn)

	return nil
}


func (lvl *internalLevel) Init() error {
	for _, fn := range lvl.fns {
		fn()
	}

	return nil
}
