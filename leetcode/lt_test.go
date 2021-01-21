package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"abcdefe", 6},
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"aab", 2},
		{" ", 1},
		{"anviaj", 5},
		{"abcb", 3},
		{"abba", 2},
	}
	for _, test := range tests {
		output := lengthOfLongestSubstring(test.input)
		if output != test.want {
			t.Errorf("test failed, input:%s, output:%d, want:%d",
				test.input, output, test.want)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"abcba", "abcba"},
		{"babad", "bab"},
		{"cbbd", "bb"},
	}
	for _, test := range tests {
		output := longestPalindrome(test.input)
		if output != test.want {
			t.Errorf("test failed, input:%s, output:%s, want:%s",
				test.input, output, test.want)
		}
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 8}, 1},
	}

	for _, test := range tests {
		output := maxArea(test.input)
		if output != test.want {
			t.Errorf("test failed, input:%v, output:%d, want:%d",
				test.input, output, test.want)
		}
	}
}

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-1, 2, 1, 5, -4, 2, 2, -4}, 7, 6},
	}

	for _, test := range tests {
		output := threeSumClosest(test.input1, test.input2)
		if output != test.want {
			t.Errorf("test failed, input:%v %v, output:%d, want:%d",
				test.input1, test.input2, output, test.want)
		}
	}
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func TestRemoveNthFromEnd(t *testing.T) {
	input1 := makeLinkedList([]int{1})
	input2 := 1
	want := makeLinkedList([]int{})
	output := removeNthFromEnd(input1, input2)
	if compareLinkedList(output, want) == false {
		t.Error()
	}
}

func TestGenerateParenthesis(t *testing.T) {
	input := 3
	output := generateParenthesis(input)
	if len(output) != 5 {
		t.Error()
	}
}

func TestSwapPairs(t *testing.T) {
	input := makeLinkedList([]int{1, 2, 3, 4, 5})
	want := makeLinkedList([]int{2, 1, 4, 3, 5})
	output := swapPairs(input)
	if !compareLinkedList(output, want) {
		t.Errorf("intput:%v, want:%v, output:%v", input, want, output)
	}
}

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 3, 2, 1}, []int{1, 2, 4, 1, 2, 3, 3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}},
	}
	for _, test := range tests {
		nextPermutation(test.input)
		if !compareIntegerSlice(test.input, test.want) {
			t.Error()
		}
	}
}

func TestFindPivot(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1}, 0},
		{[]int{1, 2, 3}, 0},
		{[]int{3, 1, 2}, 1},
		{[]int{2, 3, 1}, 2},
		{[]int{4, 5, 1, 2, 3}, 2},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4},
	}
	for _, test := range tests {
		output := findPivot(test.input)
		if output != test.want {
			t.Errorf("input:%+v, output:%d, want:%d",
				test.input, output, test.want)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		input3 int
		input4 int
		want   int
	}{
		{[]int{1}, 0, 0, 100, -1},
		{[]int{1}, 0, 0, 1, 0},
		{[]int{1, 2}, 0, 1, 1, 0},
		{[]int{1, 2}, 0, 1, 2, 1},
		{[]int{1, 2}, 0, 1, 100, -1},
		{[]int{1, 2, 3}, 0, 2, 2, 1},
		{[]int{3, 1, 2}, 1, 2, 2, 2},
		{[]int{2, 3, 1}, 0, 1, 2, 0},
		{[]int{4, 5, 1, 2, 3}, 2, 4, 3, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4, 6, 1, 5},
	}
	for _, test := range tests {
		output := binarySearch(test.input1, test.input2, test.input3, test.input4)
		if output != test.want {
			t.Errorf("input:%+v, output:%d, want:%d",
				test.input1, output, test.want)
		}
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   int
	}{
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 1, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{3, 1, 2}, 2, 2},
		{[]int{2, 3, 1}, 2, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	}
	for _, test := range tests {
		output := search(test.input1, test.input2)
		if output != test.want {
			t.Errorf("input:%+v, output:%d, want:%d",
				test.input1, output, test.want)
		}
	}
}

func TestFindLeft(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   int
	}{
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 1, 0},
		{[]int{2, 2, 3}, 2, 0},
		{[]int{1, 2, 2, 3}, 2, 1},
		{[]int{1, 2, 2, 3}, 3, 3},
		{[]int{1, 2, 2, 3}, 4, -1},
		{[]int{1, 2, 2, 3}, 0, -1},
	}
	for _, test := range tests {
		output := findLeft(test.input1, test.input2)
		if output != test.want {
			t.Errorf("input:%+v, output:%d, want:%d",
				test.input1, output, test.want)
		}
	}
}

func TestFindRight(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   int
	}{
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 1, 0},
		{[]int{2, 2, 3}, 2, 1},
		{[]int{1, 2, 2, 3}, 2, 2},
		{[]int{1, 2, 2, 2}, 2, 3},
		{[]int{1, 2, 2, 3}, 4, -1},
		{[]int{1, 2, 2, 3}, 0, -1},
	}
	for _, test := range tests {
		output := findRight(test.input1, test.input2)
		if output != test.want {
			t.Errorf("input:%+v, output:%d, want:%d",
				test.input1, output, test.want)
		}
	}
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   []int
	}{
		{[]int{1}, 1, []int{0, 0}},
		{[]int{1, 2}, 1, []int{0, 0}},
		{[]int{2, 2, 3}, 2, []int{0, 1}},
		{[]int{1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 2, 2}, 2, []int{1, 3}},
		{[]int{1, 2, 2, 3}, 4, []int{-1, -1}},
		{[]int{1, 2, 2, 3}, 0, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{}, 0, []int{-1, -1}},
	}
	for _, test := range tests {
		output := searchRange(test.input1, test.input2)
		if output[0] != test.want[0] || output[1] != test.want[1] {
			t.Errorf("input:%+v, output:%d, want:%d",
				test.input1, output, test.want)
		}
	}
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	if !isValidSudoku(board) {
		t.Error()
	}
}
