package app

type state struct {
	Created int
	Ready   int
	Running int
}

type env struct {
	Default string
	Local   string
	Dev     string
	Test    string
	Mock    string
}

var State = &state{
	Created: 0,
	Ready:   1,
	Running: 2,
}

var Env = &env{
	Default: "",
	Local:   "local",
	Dev:     "dev",
	Test:    "test",
	Mock:    "mock",
}