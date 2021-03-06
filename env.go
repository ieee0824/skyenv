package skyenv

import (
	"errors"
	"os"

	"github.com/google/skylark"
)

func init() {
	skylark.Universe["getenv"] = skylark.NewBuiltin("getnev", getenv)
	skylark.Universe["setenv"] = skylark.NewBuiltin("setnev", setenv)
	skylark.Universe["unsetenv"] = skylark.NewBuiltin("unsetnev", unsetenv)
}

func getenv(thread *skylark.Thread, fn *skylark.Builtin, args skylark.Tuple, kwargs []skylark.Tuple) (skylark.Value, error) {
	if len(args) != 1 {
		return skylark.None, errors.New("not enough arguments in call to getenv")
	}

	key, ok := skylark.AsString(args[0])
	if !ok {
		return skylark.None, errors.New("type string in argument to getenv")
	}

	env := os.Getenv(key)
	return skylark.String(env), nil
}

func setenv(thread *skylark.Thread, fn *skylark.Builtin, args skylark.Tuple, kwargs []skylark.Tuple) (skylark.Value, error) {
	if len(args) != 2 {
		return skylark.None, errors.New("not enough arguments in call to setenv")
	}

	key, ok := skylark.AsString(args[0])
	if !ok {
		return skylark.None, errors.New("type string in argument to setenv")
	}

	val, ok := skylark.AsString(args[1])
	if !ok {
		return skylark.None, errors.New("type string in argument to setenv")
	}

	if err := os.Setenv(key, val); err != nil {
		return skylark.None, err
	}
	return skylark.None, nil
}

func unsetenv(thread *skylark.Thread, fn *skylark.Builtin, args skylark.Tuple, kwargs []skylark.Tuple) (skylark.Value, error) {
	if len(args) == 0 {
		return skylark.None, errors.New("not enough arguments in call to unsetenv")
	} else if 2 <= len(args) {
		return skylark.None, errors.New("too many arguments in call to unsetenv")
	}

	key, ok := skylark.AsString(args[0])
	if !ok {
		return skylark.None, errors.New("type string in argument to unsetenv")
	}

	if err := os.Unsetenv(key); err != nil {
		return skylark.None, err
	}
	return skylark.None, nil
}
