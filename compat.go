package check

import (
	"testing"

	"github.com/juju/tc"
)

type C = tc.C
type Checker = tc.Checker
type CheckerInfo = tc.CheckerInfo
type CommentInterface = tc.CommentInterface

var (
	IsNil        Checker = tc.IsNil
	NotNil       Checker = tc.NotNil
	Equals       Checker = tc.Equals
	DeepEquals   Checker = deepEquals()
	HasLen       Checker = tc.HasLen
	ErrorMatches Checker = tc.ErrorMatches
	Matches      Checker = tc.Matches
	Panics       Checker = tc.Panics
	PanicMatches Checker = tc.PanicMatches
	FitsTypeOf   Checker = tc.FitsTypeOf
	Implements   Checker = tc.Implements
)

func TestingT(t *testing.T) {
	t.Helper()
	tc.InternalTestingT(t)
}

func Suite(suite any) any {
	return tc.InternalSuite(suite)
}

func Not(checker Checker) Checker {
	return tc.Not(checker)
}

func Commentf(format string, args ...any) CommentInterface {
	return tc.Commentf(format, args)
}
