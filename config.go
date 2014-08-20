package main

// START OMIT
type Config struct {
	// MaxCount sets the maximum number of iterations. If zero,
	// MaxCountScale is used.
	MaxCount int
	// MaxCountScale is a non-negative scale factor applied to the default
	// maximum. If zero, the default is unchanged.
	MaxCountScale float64
	// If non-nil, rand is a source of random numbers. Otherwise a default
	// pseudo-random source will be used.
	Rand *rand.Rand
	// If non-nil, the Values function generates a slice of arbitrary
	// reflect.Values that are congruent with the arguments to the function
	// being tested. Otherwise, the top-level Values function is used
	// to generate them.
	Values func([]reflect.Value, *rand.Rand)
}

// END OMIT
