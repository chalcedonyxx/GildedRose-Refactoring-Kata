package gildedrose_test

import (
	"testing"
	// "fmt"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

type testCase struct{
	testName string
	testItem []*gildedrose.Item
	expectedQuality int
}

var TestCases = []testCase{
		{
			testName: "Normal goods changes quality correctly",
			testItem: []*gildedrose.Item{{"Apple", 10, 50}},
			expectedQuality: 49,
		},
		{
			testName: "Normal goods changes quality correctly at 0 days",
			testItem: []*gildedrose.Item{{"Apple", 0, 50}},
			expectedQuality: 48,
		},
		{
			testName: "Normal goods changes quality correctly after sell by date",
			testItem: []*gildedrose.Item{{"Apple", -1, 40}},
			expectedQuality: 38,
		},
		{
			testName: "Normal goods quality cannot be less than 0",
			testItem: []*gildedrose.Item{{"Apple", 1, 0}},
			expectedQuality: 0,
		},
		{
			testName: "Normal goods quality decreases by 1 to 0 after sell by",
			testItem: []*gildedrose.Item{{"Apple", -1, 1}},
			expectedQuality: 0,
		},
		{
			testName: "Aged brie changes quality correctly",
			testItem: []*gildedrose.Item{{"Aged Brie", 10, 40}},
			expectedQuality: 41,
		},
		{
			testName: "Aged brie changes quality correctly at 0 days",
			testItem: []*gildedrose.Item{{"Aged Brie", 0, 40}},
			expectedQuality: 42,
		},
		{
			testName: "Aged brie changes quality correctly after sell by date",
			testItem: []*gildedrose.Item{{"Aged Brie", -1, 40}},
			expectedQuality: 42,
		},
		{
			testName: "Aged brie quality not greater than 50",
			testItem: []*gildedrose.Item{{"Aged Brie", 10, 50}},
			expectedQuality: 50,
		},
		{
			testName: "Aged brie quality increases by 1 to 50 after sell by",
			testItem: []*gildedrose.Item{{"Aged Brie", -1, 49}},
			expectedQuality: 50,
		},
	}

func TestSulfuras(t *testing.T) {
	var testCasesSulfuras = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", 10, 80},
	}

	gildedrose.UpdateQuality(testCasesSulfuras)

	for _, testCaseSulfuras := range testCasesSulfuras {
		if (testCaseSulfuras.Quality != 80){
			t.Errorf("Sulfuras Quality: Expected %d but got %d ", 80, testCaseSulfuras.Quality)
		}
	}
}

func TestNormalGoods(t *testing.T) {
	for _, testCase := range TestCases {
		t.Run(testCase.testName, func(t *testing.T){
			gildedrose.UpdateQuality(testCase.testItem)
			assert.Equal(t, testCase.expectedQuality, testCase.testItem[0].Quality)
		})
	}
}