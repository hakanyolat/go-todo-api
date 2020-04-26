package app

type state int

type stateGroup struct {
	Created state
	Ready   state
	Running state
}

var State = &stateGroup{
	Created: 0,
	Ready:   1,
	Running: 2,
}
