package _test

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

// NewForTest returns a new App for testing.
//
// https://github.com/uber-go/fx/blob/master/app_test.go#L50
func NewForTest(tb testing.TB, opts ...fx.Option) *fxtest.App {
	testOpts := []fx.Option{
		// Provide both: Logger and WithLogger so that if the test
		// WithLogger fails, we don't pollute stderr.
		fx.Logger(fxtest.NewTestPrinter(tb)),
		fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(tb) }),
	}
	opts = append(testOpts, opts...)

	return fxtest.New(tb, opts...)
}
