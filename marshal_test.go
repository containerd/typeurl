package typeurl

import (
	"fmt"
	"reflect"
	"testing"
)

type TestType struct {
	ID string
}

func init() {
	Register(&TestType{}, "typeurl.Type")
}

func TestMarshalEvent(t *testing.T) {
	for _, testcase := range []struct {
		event interface{}
		url   string
	}{
		{
			event: &TestType{ID: "Test"},
			url:   "typeurl.Type",
		},
	} {
		t.Run(fmt.Sprintf("%T", testcase.event), func(t *testing.T) {
			a, err := MarshalAny(testcase.event)
			if err != nil {
				t.Fatal(err)
			}
			if a.TypeUrl != testcase.url {
				t.Fatalf("unexpected url: %q != %q", a.TypeUrl, testcase.url)
			}

			v, err := UnmarshalAny(a)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(v, testcase.event) {
				t.Fatalf("round trip failed %v != %v", v, testcase.event)
			}
		})
	}
}

func BenchmarkMarshalEvent(b *testing.B) {
	ev := &TestType{}
	expected, err := MarshalAny(ev)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		a, err := MarshalAny(ev)
		if err != nil {
			b.Fatal(err)
		}
		if a.TypeUrl != expected.TypeUrl {
			b.Fatalf("incorrect type url: %v != %v", a, expected)
		}
	}
}
