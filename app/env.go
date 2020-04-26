package app

import "reflect"

type env string

type envGroup struct {
	Default env
	Local   env
	Dev     env
	Test    env
	Prod    env
	Mock    env
}

var Env = &envGroup{
	Default: "",
	Local:   "local",
	Dev:     "dev",
	Test:    "test",
	Prod:    "prod",
	Mock:    "mock",
}

func (e *envGroup) GetPossibleValues() []string {
	var envs []string

	r := reflect.ValueOf(e)
	ri := reflect.Indirect(r)

	for i := 0; i < ri.NumField(); i++ {
		f := ri.Field(i)
		v := f.String()
		if v != "" {
			envs = append(envs, v)
		}
	}

	return envs
}
