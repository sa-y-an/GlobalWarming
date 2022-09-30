package main

import "errors"

func injectApi() (int, error) {
	err := errors.New("bye")
	return 2, err
}
