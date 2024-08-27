package main

import (
	"fmt"
	longest_palindromic_substring "leetcode/longest-palindromic-substring"
	longest_substring_without_repeating_characters "leetcode/longest-substring-without-repeating-characters"
	minimum_array_end "leetcode/minimum-array-end"
	two_sum "leetcode/two-sum"
)

func main() {
	test5()
}
func test5() {
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("a"))
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("ab"))
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("aba"))
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("abba"))
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("aa"))
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("aaa"))
	//fmt.Println(longest_palindromic_substring.LongestPalindrome("aaaa"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("aaaabaaaaba"))
}

func test3133() {
	fmt.Println(minimum_array_end.MinEnd(3, 4))
	fmt.Println(minimum_array_end.MinEnd(2, 7))
}

func test3() {
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("abcabcbb"))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("bbbbb"))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("pwwkew"))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring(""))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("a"))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("ab"))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("abc"))
	fmt.Println(longest_substring_without_repeating_characters.LengthOfLongestSubstring("abcd"))
}

func test2() {
	fmt.Println(two_sum.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(two_sum.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(two_sum.TwoSum([]int{3, 3}, 6))
}
