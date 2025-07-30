package gildedrose_test

import (
	"testing"
	// "fmt"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

type testCase struct{
	testItem []*gildedrose.Item
	expectedQuality int
}

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

func TestNormalGoods(t *testing.T) {
	var testCases = []testCase{
		{testItem: []*gildedrose.Item{{"Apple", 10, 50}},
		expectedQuality: 49},
		{testItem: []*gildedrose.Item{{"Apple", 0, 50}},
		expectedQuality: 48},
		{testItem: []*gildedrose.Item{{"Apple", -1, 40}},
		expectedQuality: 38},
	}

	for _, testCase := range testCases {
		gildedrose.UpdateQuality(testCase.testItem)
		assert.Equal(t, testCase.expectedQuality, testCase.testItem[0].Quality)
	}
}
