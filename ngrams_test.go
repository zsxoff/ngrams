package ngrams

import (
	"reflect"
	"testing"
)

type testpair struct {
	line   string
	ngrams []string
}

var testsUnigrams = []testpair{
	{"h", []string{"h"}},
	{"he", []string{"h", "e"}},
	{"hello", []string{"h", "e", "l", "l", "o"}},
	{"hello ", []string{"h", "e", "l", "l", "o", " "}},
}

var testsBigrams = []testpair{
	{"h", []string{}},
	{"he", []string{"he"}},
	{"hello", []string{"he", "el", "ll", "lo"}},
	{"hello ", []string{"he", "el", "ll", "lo", "o "}},
}

var testsTrigrams = []testpair{
	{"h", []string{}},
	{"he", []string{}},
	{"hello", []string{"hel", "ell", "llo"}},
	{"hello ", []string{"hel", "ell", "llo", "lo "}},
}

func TestUnigrams(t *testing.T) {
	for _, pair := range testsUnigrams {
		result := Unigrams(pair.line)

		if !reflect.DeepEqual(result, pair.ngrams) {
			t.Error("For", pair.line, "expected", pair.ngrams, "got", result)
		}
	}
}

func TestBigrams(t *testing.T) {
	for _, pair := range testsBigrams {
		result := Bigrams(pair.line)

		if !reflect.DeepEqual(result, pair.ngrams) {
			t.Error("For", pair.line, "expected", pair.ngrams, "got", result)
		}
	}
}

func TestTrigrams(t *testing.T) {
	for _, pair := range testsTrigrams {
		result := Trigrams(pair.line)

		if !reflect.DeepEqual(result, pair.ngrams) {
			t.Error("For", pair.line, "expected", pair.ngrams, "got", result)
		}
	}
}
