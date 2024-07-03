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
		{
			"slices",
			[]Profile{
				{42, "São Paulo"},
				{34, "Rio de Janeiro"},
			},
			[]string{"São Paulo", "Rio de Janeiro"},
		},
		{
			"arrays",
			[2]Profile{
				{42, "São Paulo"},
				{34, "Rio de Janeiro"},
			},
			[]string{"São Paulo", "Rio de Janeiro"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
