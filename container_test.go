package initcontainer


import (
	"testing"

	"math/rand"
	"time"
)


// TestNew makes sure that the initcontainer.New() func returns a non-nil value.
func TestNew(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )

	// Create new container.
	numLevels := uint(1 + randomness.Intn(40))

	initContainer := New(numLevels)

	// Confirm that what we get back is not nil.
	if nil == initContainer {
		t.Errorf("Excepted New() to return a container, but instead it returned %v", initContainer)
	}
}

// TestRegisterFuncAndInit tests out stuff in a way somewhat more similar to real life usage.
func TestRegisterFuncAndInit(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )

	// Do this test a number of times.
	numTestRepeats := 1 + rand.Intn(10)

	for testRepeatNumber := 0; testRepeatNumber < numTestRepeats ; testRepeatNumber++ {

		// Figure out how many levels we want.
		numLevels := 5 + uint(randomness.Intn(20))

		// Figure out the size of each level.
		levelSizes := make([]uint, numLevels)

		for i:=uint(0); i<numLevels; i++ {

			levelSizes[i] = uint(randomness.Intn(25))
		}

		// Note that all items in the slice are initialized
		// to false.
		//
		// This slice is used in the tests that come after
		// this.
		//
		// If those tests are successful, then every item
		// in this slice should be turned to true.
		slice := make([][]bool, numLevels)
		for i:=uint(0); i<numLevels; i++ {
			slice[i] = make([]bool, levelSizes[i])
		}

		// Create a new container.
		initContainer := New(numLevels)

		// Put stuff into the container.
		for levelNumber := uint(0); levelNumber < numLevels; levelNumber++ {

			levelSize := levelSizes[levelNumber]

			for levelIndex := uint(0); levelIndex < levelSize; levelIndex++ {

				fn := func(i uint, ii uint) (func()) {

					fn1 := func() {

						slice[i][ii] = true

					}

					return fn1

				}(levelNumber, levelIndex)



				initContainer.RegisterFunc(fn, levelNumber)

			}

		}

		// Init all the stuff in the container.
		initContainer.Init()

		// Confirm that all the funcs put in the container were actually executed.
		for levelNumber, subSlice := range slice {

			for levelIndex, value := range subSlice {

				if true != value {
					t.Errorf("For level number %d and level index %d, expected value to be true, but actually was %t", levelNumber, levelIndex, value)
				}
			}
		}

	}
}

// TestRegisterAndInit tests out stuff in a way somewhat more similar to real life usage.
func TestRegisterAndInit(t *testing.T) {

	// Initialize.
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )

	// Do this test a number of times.
	numTestRepeats := 1 + rand.Intn(10)

	for testRepeatNumber := 0; testRepeatNumber < numTestRepeats ; testRepeatNumber++ {

		// Figure out how many levels we want.
		numLevels := 5 + uint(randomness.Intn(20))

		// Figure out the size of each level.
		levelSizes := make([]uint, numLevels)

		for i:=uint(0); i<numLevels; i++ {

			levelSizes[i] = uint(randomness.Intn(25))
		}

		// Note that all items in the slice are initialized
		// to false.
		//
		// This slice is used in the tests that come after
		// this.
		//
		// If those tests are successful, then every item
		// in this slice should be turned to true.
		slice := make([][]bool, numLevels)
		for i:=uint(0); i<numLevels; i++ {
			slice[i] = make([]bool, levelSizes[i])
		}

		// Create a new container.
		initContainer := New(numLevels)

		// Put stuff into the container.
		for levelNumber := uint(0); levelNumber < numLevels; levelNumber++ {

			levelSize := levelSizes[levelNumber]

			for levelIndex := uint(0); levelIndex < levelSize; levelIndex++ {

				fn := func(i uint, ii uint) (func()) {

					fn1 := func() {

						slice[i][ii] = true

					}

					return fn1

				}(levelNumber, levelIndex)


				initializer := newInitializerFromFunc(fn)
				initContainer.Register(initializer, levelNumber)

			}

		}

		// Init all the stuff in the container.
		initContainer.Init()

		// Confirm that all the funcs put in the container were actually executed.
		for levelNumber, subSlice := range slice {

			for levelIndex, value := range subSlice {

				if true != value {
					t.Errorf("For level number %d and level index %d, expected value to be true, but actually was %t", levelNumber, levelIndex, value)
				}
			}
		}

	}
}
