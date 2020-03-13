/*
Copyright 2014 SAP SE

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"testing"
)

type testIdentifier struct {
	id Identifier
	s  string
}

var testIdentifierData = []*testIdentifier{
	{"_", "_"},
	{"_A", "_A"},
	{"A#$_", "A#$_"},
	{"1", `"1"`},
	{"a", `"a"`},
	{"$", `"$"`},
	{"日本語", `"日本語"`},
	{"testTransaction", `"testTransaction"`},
	{"a.b.c", `"a.b.c"`},
	{"AAA.BBB.CCC", `"AAA.BBB.CCC"`},
}

func TestIdentifierStringer(t *testing.T) {
	for i, d := range testIdentifierData {
		if d.id.String() != d.s {
			t.Fatalf("%d id %s - expected %s", i, d.id, d.s)
		}
	}
}
