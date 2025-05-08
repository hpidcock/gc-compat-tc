package check

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

// Gocheck - A rich testing framework for Go
//
// Copyright (c) 2010-2013 Gustavo Niemeyer <gustavo@niemeyer.net>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// -----------------------------------------------------------------------
// DeepEquals checker.

type deepEqualsChecker struct {
	*CheckerInfo
}

// The DeepEquals checker verifies that the obtained value is deep-equal to
// the expected value.  The check will work correctly even when facing
// slices, interfaces, and values of different types (which always fail
// the test).
//
// For example:
//
//	c.Assert(value, DeepEquals, 42)
//	c.Assert(array, DeepEquals, []string{"hi", "there"})
func deepEquals() Checker {
	return &deepEqualsChecker{
		&CheckerInfo{Name: "DeepEquals", Params: []string{"obtained", "expected"}},
	}
}

func (checker *deepEqualsChecker) Check(params []interface{}, names []string) (result bool, error string) {
	result = reflect.DeepEqual(params[0], params[1])
	if !result {
		error = formatUnequal(params[0], params[1])
	}
	return
}

func diffworthy(a interface{}) bool {
	if a == nil {
		return false
	}

	t := reflect.TypeOf(a)
	switch t.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct, reflect.String, reflect.Ptr:
		return true
	}
	return false
}

// formatUnequal will dump the actual and expected values into a textual
// representation and return an error message containing a diff.
func formatUnequal(obtained interface{}, expected interface{}) string {
	// We do not do diffs for basic types because go-check already
	// shows them very cleanly.
	if !diffworthy(obtained) || !diffworthy(expected) {
		return ""
	}

	// Handle strings, short strings are ignored (go-check formats
	// them very nicely already). We do multi-line strings by
	// generating two string slices and using kr.Diff to compare
	// those (kr.Diff does not do string diffs by itself).
	aStr, aOK := obtained.(string)
	bStr, bOK := expected.(string)
	if aOK && bOK {
		l1 := strings.Split(aStr, "\n")
		l2 := strings.Split(bStr, "\n")
		// the "2" here is a bit arbitrary
		if len(l1) > 2 && len(l2) > 2 {
			diff := pretty.Diff(l1, l2)
			return fmt.Sprintf(`String difference:
%s`, formatMultiLine(strings.Join(diff, "\n"), false))
		}
		// string too short
		return ""
	}

	// generic diff
	diff := pretty.Diff(obtained, expected)
	if len(diff) == 0 {
		// No diff, this happens when e.g. just struct
		// pointers are different but the structs have
		// identical values.
		return ""
	}

	return fmt.Sprintf(`Difference:
%s`, formatMultiLine(strings.Join(diff, "\n"), false))
}

func formatMultiLine(s string, quote bool) []byte {
	b := make([]byte, 0, len(s)*2)
	i := 0
	n := len(s)
	for i < n {
		j := i + 1
		for j < n && s[j-1] != '\n' {
			j++
		}
		b = append(b, "...     "...)
		if quote {
			b = strconv.AppendQuote(b, s[i:j])
		} else {
			b = append(b, s[i:j]...)
			b = bytes.TrimSpace(b)
		}
		if quote && j < n {
			b = append(b, " +"...)
		}
		b = append(b, '\n')
		i = j
	}
	return b
}
