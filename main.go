package main

import (
	"context"

	"go.uber.org/fx"

	"fmt"

	"uber-fx-test/fop"
	"uber-fx-test/module_one"
)

type MainLoop struct {
	shutdowner fx.Shutdowner
	mo         *module_one.ModuleOne
}

func NewMainLoop(lc fx.Lifecycle, shutdowner fx.Shutdowner, mo *module_one.ModuleOne) *MainLoop {
	ml := new(MainLoop)

	ml.shutdowner = shutdowner
	ml.mo = mo

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				ml.SomeFunc()

				return nil
			},
		},
	)

	return ml
}

func (ml *MainLoop) SomeFunc() {
	fmt.Println("This is also a test.")

	ml.mo.SomeFunc()

	ml.shutdowner.Shutdown()
}

type TestStruct struct {
	first_param  string
	second_param int16
}

func NewTestStruct(options ...func(*TestStruct)) *TestStruct {
	ts := &TestStruct{}

	for _, o := range options {
		o(ts)
	}

	return ts
}

func WithFirstParam(first_param string) func(*TestStruct) {
	return func(ts *TestStruct) {
		ts.first_param = first_param
	}
}

func WithSecondParam(second_param int16) func(*TestStruct) {
	return func(ts *TestStruct) {
		ts.second_param = second_param
	}
}

func main() {
	fmt.Println("This is a test.")

	ts := NewTestStruct(
		WithFirstParam("first"),
		WithSecondParam(20),
	)

	fmt.Println(ts)

	fs := fop.NewFopStruct(
		fop.WithOptionOne("option_one"),
		fop.WithOptionTwo("option_two"),
	)

	fmt.Println(fs)

	fx.New(
		fx.Provide(module_one.NewModuleOne),
		fx.Invoke(NewMainLoop),
		fx.NopLogger,
	).Run()
}
