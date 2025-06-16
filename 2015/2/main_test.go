package main

import (
	"reflect"
	"testing"
)

func TestGetPresentWraper(t *testing.T) {
	t.Run(
		"it should return 52 when 2x3x4 is provided",
		func(t *testing.T) {
			present := Present{
				lenght: 2,
				width:  3,
				height: 4,
			}

			got := GetPresentWraper(present)
			want := PresentWrap{lenght: present.lenght, width: present.width, heigth: present.height, squareFeet: 52}

			assertPresentWrap(t, want, got)
		},
	)

	t.Run(
		"it should return 42 when 1x1x10 is provided",
		func(t *testing.T) {
			present := Present{
				lenght: 1,
				width: 1,
				height: 10,
			}

			got := GetPresentWraper(present)
			want := PresentWrap{lenght: present.lenght, width: present.width, heigth: present.height, squareFeet: 42}

			assertPresentWrap(t, want, got)
		},
	)
}

func TestGetExtraPaper(t *testing.T) {
	t.Run("it should return 6 when 2x3x4 is provided", func(t *testing.T) {

		present := Present{lenght: 2, width: 3, height: 4}
		presentWrap := PresentWrap{lenght: present.lenght, width: present.width, heigth: present.height, squareFeet: 52}

		got := presentWrap.GetExtraPaper()
		want := 6

		assertEqual(t, want, got)
	})

	t.Run(
		"it should return 1 when 1x1x10 is provided",
		func(t *testing.T) {
			present := Present{
				lenght: 1,
				width:  1,
				height: 10,
			}

			presentWrap := PresentWrap{lenght: present.lenght, width: present.width, heigth: present.height, squareFeet: 42}

			got := presentWrap.GetExtraPaper()
			want := 1

			assertEqual(t, want, got)
		},
	)
}

func assertPresentWrap(t testing.TB, want, got PresentWrap) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func assertEqual(t testing.TB, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
