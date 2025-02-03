package app

import "fmt"

var funcMap = map[string]func(args ...string) (string, error){}

func RegisterFunc(name string, f func(args ...string) (string, error)) error {
	if _, ok := funcMap[name]; ok {
		return fmt.Errorf("function %s already registered", name)
	}
	funcMap[name] = f
	return nil
}

func UnregisterFunc(name string) error {
	if _, ok := funcMap[name]; !ok {
		return fmt.Errorf("function %s not registered", name)
	}
	delete(funcMap, name)
	return nil
}

// test...
func PrintArgs(args ...string) (string, error) {
	return fmt.Sprintf("Reflog register successfully: %s", args), nil
}

func init() {
	RegisterFunc("PrintArgs", PrintArgs)
}
