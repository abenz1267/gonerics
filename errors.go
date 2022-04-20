package gonerics

import (
	"fmt"
	"log"
	"runtime"
)

func TryResult[T any](result T, err error) T {
	doPanic(err)

	return result
}

func Try(err error) {
	doPanic(err)
}

// This allocated way more because of the call to runtime.Caller.
func doPanic(err error) {
	if err != nil {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			panic(fmt.Errorf("%s:%d: %w", file, no, err))
		}

		panic(err)
	}
}

func RecoverHandler(handlers ...func(r interface{})) {
	if r := recover(); r != nil {
		for _, v := range handlers {
			v(r)
		}
	}
}

func RecoverPrint() {
	if r := recover(); r != nil {
		log.Println(r)
	}
}

func RecoverPanic() {
	if r := recover(); r != nil {
		panic(r)
	}
}

func RecoverFatal() {
	if r := recover(); r != nil {
		log.Fatal(r)
	}
}
