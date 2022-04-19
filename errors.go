package gonerics

import (
	"fmt"
	"runtime"
)

func TryResult[T any](result T, err error) T {
	doPanic(err)

	return result
}

func Try(err error) {
	doPanic(err)
}

func doPanic(err error) {
	if err != nil {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			panic(fmt.Errorf("%s:%d: %w", file, no, err))
		}

		panic(err)
	}
}
