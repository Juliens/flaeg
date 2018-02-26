package parse

import (
	"reflect"
	"testing"
	"time"
)

func TestSliceStringsSet(t *testing.T) {
	checkMap := map[string]SliceStrings{
		"str":            {"str"},
		"str1,str2":      {"str1", "str2"},
		"str1;str2":      {"str1", "str2"},
		"str1,str2;str3": {"str1", "str2", "str3"},
	}

	for str, check := range checkMap {
		var slice SliceStrings
		if err := slice.Set(str); err != nil {
			t.Fatalf("Error :%s", err)
		}

		if !reflect.DeepEqual(slice, check) {
			t.Errorf("Expected: %s\ngot: %s", check, slice)
		}
	}
}

func TestSliceStringsSetAdd(t *testing.T) {
	slice := SliceStrings{"str1"}

	//test
	if err := slice.Set("str2,str3"); err != nil {
		t.Fatalf("Error :%s", err)
	}

	//check
	check := SliceStrings{"str1", "str2", "str3"}
	if !reflect.DeepEqual(slice, check) {
		t.Errorf("Expected: %s\ngot: %s", check, slice)
	}
}

func TestSliceStringsGet(t *testing.T) {
	slices := []SliceStrings{
		{"str"},
		{"str1", "str2"},
		{"str1", "str2", "str3"},
	}
	check := [][]string{
		{"str"},
		{"str1", "str2"},
		{"str1", "str2", "str3"},
	}

	for i, slice := range slices {
		if !reflect.DeepEqual(slice.Get(), check[i]) {
			t.Errorf("Expected: %s\ngot: %s", check[i], slice)
		}
	}
}

func TestSliceStringsString(t *testing.T) {
	slices := []SliceStrings{
		{"str"},
		{"str1", "str2"},
		{"str1", "str2", "str3"},
	}
	check := []string{
		"[str]",
		"[str1 str2]",
		"[str1 str2 str3]",
	}

	for i, slice := range slices {
		if !reflect.DeepEqual(slice.String(), check[i]) {
			t.Errorf("Expected:%s\ngot:%s", check[i], slice)
		}
	}
}

func TestSliceStringsSetValue(t *testing.T) {
	check := []SliceStrings{
		{"str"},
		{"str1", "str2"},
		{"str1", "str2", "str3"},
	}
	slices := [][]string{
		{"str"},
		{"str1", "str2"},
		{"str1", "str2", "str3"},
	}

	for i, s := range slices {
		var slice SliceStrings
		slice.SetValue(s)
		if !reflect.DeepEqual(slice, check[i]) {
			t.Errorf("Expected: %s\ngot: %s", check[i], slice)
		}
	}
}

func TestSetDuration(t *testing.T) {
	tests := []struct {
		in  string
		out time.Duration
	}{
		{
			in:  "42",
			out: 42 * time.Second,
		},
		{
			in:  "42s",
			out: 42 * time.Second,
		},
		{
			in:  "5m",
			out: 5 * time.Minute,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()

			var dur Duration
			if err := dur.Set(test.in); err != nil {
				t.Fatal(err)
			}

			if time.Duration(dur) != test.out {
				t.Errorf("got %#v, want %#v", time.Duration(dur), test.out)
			}
		})
	}
}

func TestUnmarshalTextDuration(t *testing.T) {
	var dur Duration
	if err := dur.UnmarshalText([]byte("42")); err != nil {
		t.Fatalf("got error %s", err)
	}

	if time.Duration(dur) != 42*time.Second {
		t.Errorf("got %#v, want %#v", time.Duration(dur), 42*time.Second)
	}
}

func TestMarshalTextDuration(t *testing.T) {
	var dur Duration
	dur.Set("42")
	duration, _ := dur.MarshalText()
	if string(duration) != "42s" {
		t.Errorf("got %#v, want %#v", dur, "42s")
	}
}

func TestMarshalTextDurationWithHourAndMinutes(t *testing.T) {
	var dur Duration
	dur.Set("3670")
	duration, _ := dur.MarshalText()
	if string(duration) != "1h1m10s" {
		t.Errorf("got %#v, want %#v", dur, "1h1m10s")
	}
}

func TestUnmarshalJsonDuration(t *testing.T) {
	var dur Duration
	if err := dur.UnmarshalJSON([]byte("1000000000")); err != nil {
		t.Fatalf("go error %s", err)
	}

	if time.Duration(dur) != time.Second {
		t.Errorf("got %#v, want %#v", time.Duration(dur), time.Second)
	}
}
