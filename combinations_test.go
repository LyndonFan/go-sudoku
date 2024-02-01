package main

import (
	"reflect"
	"testing"
)

// Generates all combinations of k elements from a set of n elements
func TestGenerateCombinations(t *testing.T) {
	ch := make(chan []int)
	go combinationsRange(4, 2, ch)
	combinations := make([][]int, 0)
	for combination := range ch {
		combinations = append(combinations, combination)
	}
	expected := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	if !reflect.DeepEqual(combinations, expected) {
		t.Errorf("Expected %v, but got %v", expected, combinations)
	}
}

// Generates combinations in lexicographic order
func TestLexicographicOrder(t *testing.T) {
	ch := make(chan []int)
	go combinationsRange(4, 2, ch)
	combinations := make([][]int, 0)
	for combination := range ch {
		combinations = append(combinations, combination)
	}
	expected := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	if !reflect.DeepEqual(combinations, expected) {
		t.Errorf("Expected %v, but got %v", expected, combinations)
	}
}

// Generates an empty combination when k = 0
func TestEmptyCombination(t *testing.T) {
	ch := make(chan []int)
	go combinationsRange(4, 0, ch)
	combinations := make([][]int, 0)
	for combination := range ch {
		combinations = append(combinations, combination)
	}
	expected := [][]int{{}}
	if !reflect.DeepEqual(combinations, expected) {
		t.Errorf("Expected %v, but got %v", expected, combinations)
	}
}

// Generates no combination when n = 0
func TestNoCombinationEmptyList(t *testing.T) {
	ch := make(chan []int)
	go combinationsRange(0, 2, ch)
	combinations := make([][]int, 0)
	for combination := range ch {
		combinations = append(combinations, combination)
	}
	expected := [][]int{}
	if !reflect.DeepEqual(combinations, expected) {
		t.Errorf("Expected %v, but got %v", expected, combinations)
	}
}

// Generates no combination when k > n
func TestNoCombinationTooMany(t *testing.T) {
	ch := make(chan []int)
	go combinationsRange(2, 3, ch)
	combinations := make([][]int, 0)
	for combination := range ch {
		combinations = append(combinations, combination)
	}
	expected := [][]int{}
	if !reflect.DeepEqual(combinations, expected) {
		t.Errorf("Expected %v, but got %v", expected, combinations)
	}
}

func TestCombinationDifferentList(t *testing.T) {
	ch := make(chan []int)
	xs := []int{5, -7, 9, -11}
	go combinations(xs, 2, ch)
	combinations := make([][]int, 0)
	for combination := range ch {
		combinations = append(combinations, combination)
	}
	expected := [][]int{{5, -7}, {5, 9}, {5, -11}, {-7, 9}, {-7, -11}, {9, -11}}
	if !reflect.DeepEqual(combinations, expected) {
		t.Errorf("Expected %v, but got %v", expected, combinations)
	}
}
