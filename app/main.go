package app

type state struct {
	Idled   int
	Ready   int
	Running int
}

var ApplicationState = &state{
	Idled:   0,
	Ready:   1,
	Running: 2,
}
