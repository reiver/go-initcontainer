package initcontainer


import (
	"github.com/reiver/go-initcontainer/level"
)


type Initializer interface {
	Init()
}


type Container interface {
	// RegisterFunc registers an initialization func to be run at a specific level.
	// The lower the level, the earlier the initialization func is run.
	//
	// Here is an example:
	//
	//	const numLevels = 5
	//	initContainer := initcontainer.New(numLevels)
	//	
	//	const level = 2
	//	initContainer.RegisterFunc(myInit, level)
	RegisterFunc(func(), uint) error

	// Register registers something that fits the Initializer interface (i.e., something
	// which provides the Init() method)  to be run at a specific level.
	// The lower the level, the earlier the initialization func is run.
	//
	// Here is an example:
	//
	//	const numLevels = 5
	//	initContainer := initcontainer.New(numLevels)
	//	
	//	type MyInitializer struct {
	//		// Your stuff in here.
	//	}
	//	
	//	func (thing *MyInitializer) Init() {
	//		// some initialization happens here.
	//	}
	//	
	//	myInit := &MyInitializer{}
	//	
	//	const level = 2
	//	initContainer.RegisterFunc(myInit, level)
	Register(Initializer, uint) error

	// Init causes all the initialization funcs (registered with the Register func)
	// to be run.
	//
	// Initialization starts with level 0, then does level 1, then does level 2, etc.
	Init() error
}


type internalContainer struct {
	levels []initcontainerlevel.Level
}


// New returns a new initialization container.
//
// Each initialization container has a concept of "levels".
//
// Levels are "named" using non-negative integers. I.e.,
// 0, 1, 2, ....
//
// More specifically, the parameter to the New func 'numLevels'
// determines how many levels there are.
//
// So if, for example, 'numLeves' is 5 then the levels are:
// 0, 1, 2, 3, 4. (Notice that 5 is NOT a level in this example,
// because we count the levels starting at 0 (zero).)
func New(numLevels uint) Container {

	levels := make([]initcontainerlevel.Level, numLevels)

	for i:=uint(0); i<numLevels; i++ {
		levels[i] = initcontainerlevel.New()
	}

	container := internalContainer{
		levels: levels,
	}

	return &container
}



func (container *internalContainer) RegisterFunc(fn func(), levelNumber uint) error {
	level := container.levels[levelNumber]

	if err := level.Register(fn); nil != err {
		return err
	}

	return nil
}


func (container *internalContainer) Register(initer Initializer, levelNumber uint) error {
	fn := func(initializer Initializer) (func()) {
		fn1 := func() {
			initializer.Init()
		}

		return fn1
	}(initer)


	return container.RegisterFunc(fn, levelNumber)
}


func (container *internalContainer) Init() error {
	for _, level := range container.levels {
		if err := level.Init(); nil != err {
			return err
		}
	}

	return nil
}
