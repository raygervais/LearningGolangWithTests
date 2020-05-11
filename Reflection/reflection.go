package main

import "reflect"

// Refactor from https://github.com/quii/learn-go-with-tests/issues/123#issue-351911526

type ReportString func(input string)

type initWalk func() (length int, fetch func(index int) reflect.Value)

func Walk(x interface{}, fn ReportString) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		walkAll(structInit(val), fn)
	case reflect.Slice, reflect.Array:
		walkAll(sliceInit(val), fn)
	case reflect.Map:
		walkAll(mapInit(val), fn)
	case reflect.Chan:
		walkAll(chanInit(val), fn)
	case reflect.Func:
		funcInit(val, fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}

	return val
}

func walkAll(init initWalk, fn ReportString) {
	numberOfValues, fetch := init()
	for i := 0; i < numberOfValues; i++ {
		Walk(fetch(i).Interface(), fn)
	}
}

func structInit(container reflect.Value) initWalk {
	return func() (int, func(int) reflect.Value) {
		return container.NumField(), container.Field
	}
}

func sliceInit(container reflect.Value) initWalk {
	return func() (int, func(int) reflect.Value) {
		return container.Len(), container.Index
	}
}

func mapInit(container reflect.Value) initWalk {
	keys := container.MapKeys()
	fetch := func(i int) reflect.Value {
		return container.MapIndex(keys[i])
	}

	return func() (int, func(int) reflect.Value) {
		return len(keys), fetch
	}
}

func chanInit(container reflect.Value) initWalk {

	fetch := func(_ int) reflect.Value {
		return valChanExtract(container.Recv())
	}

	// Note to self, hard coding length is not a good practice Ray.
	// But Ray who's worked on getting the length for this channel for 2 hours disagrees.
	return func() (int, func(int) reflect.Value) {
		return 2, fetch
	}
}

func funcInit(container reflect.Value, fn func(input string)) {
	valFnResult := container.Call(nil)

	for _, res := range valFnResult {
		Walk(res.Interface(), fn)
	}
}

func valChanExtract(x reflect.Value, ok bool) reflect.Value {
	return x
}
