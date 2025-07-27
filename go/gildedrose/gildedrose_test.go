package gildedrose_test

import (
	"testing"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestSulfuras(t *testing.T) {
	var testCases = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", 10, 80},
	}

	gildedrose.UpdateQuality(testCases)

	for _, testCase := range testCases {
		if (testCase.Quality != 80){
			t.Errorf("Sulfuras Quality: Expected %d but got %d ", 80, testCase.Quality)
		}
	}
}
