package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Tiago"},
			[]string{"Tiago"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Tiago", "São Paulo"},
			[]string{"Tiago", "São Paulo"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Suian", 34},
			[]string{"Suian"},
		},
		{
			"nested fields",
			Person{
				"Tiago",
				Profile{42, "São Paulo"},
			},
			[]string{"Tiago", "São Paulo"},
		},
		{
			"pointers to things",
			&Person{
				"Tiago",
				Profile{42, "São Paulo"},
			},
			[]string{"Tiago", "São Paulo"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
