package initcontainer


// internalInitializerFromFunc turns a func() (that one would want to use
// for initialization) into something that fits an Initializer interface
// (so that it can be passed to Container.Register()).
type internalInitializerFromFunc struct {
	fn (func())
}

// newInitializerFromFunc return a new internalInitializerFromFunc.
func newInitializerFromFunc(fn func()) Initializer {

	initializer := internalInitializerFromFunc{
		fn:fn,
	}

	return &initializer
}

// Init is necessary to make *internalInitializerFromFunc fit the
// Initializer interface.
func (initializer *internalInitializerFromFunc) Init() {
	initializer.fn()
}
