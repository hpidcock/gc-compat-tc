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
	IsNil        Checker = panicChecker{}
	NotNil       Checker = panicChecker{}
	Equals       Checker = panicChecker{}
	DeepEquals   Checker = panicChecker{}
	HasLen       Checker = panicChecker{}
	ErrorMatches Checker = panicChecker{}
	Matches      Checker = panicChecker{}
	Panics       Checker = panicChecker{}
	PanicMatches Checker = panicChecker{}
	FitsTypeOf   Checker = panicChecker{}
	Implements   Checker = panicChecker{}
)

func TestingT(t *testing.T) {
	panic("use tc.C")
}

func Suite(suite any) any {
	panic("use tc.C")
}

func Not(checker Checker) Checker {
	panic("use tc.C")
}

func Commentf(format string, args ...any) CommentInterface {
	panic("use tc.C")
}

type panicChecker struct{}

func (panicChecker) Info() *tc.CheckerInfo {
	panic("use tc.C")
}

func (panicChecker) Check(params []any, names []string) (result bool, error string) {
	panic("use tc.C")
}
