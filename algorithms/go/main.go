package main

import (
	"fmt"
	clear_digits "leetcode/clear-digits"
	count_increasing_quadruplets "leetcode/count-increasing-quadruplets"
	find_the_maximum_length_of_a_good_subsequence_i "leetcode/find-the-maximum-length-of-a-good-subsequence-i"
	happy_students "leetcode/happy-students"
	length_of_the_longest_alphabetical_continuous_substring "leetcode/length-of-the-longest-alphabetical-continuous-substring"
	longest_palindromic_substring "leetcode/longest-palindromic-substring"
	longest_substring_without_repeating_characters "leetcode/longest-substring-without-repeating-characters"
	maximize_the_confusion_of_an_exam "leetcode/maximize-the-confusion-of-an-exam"
	maximum_strength_of_a_group "leetcode/maximum-strength-of-a-group"
	minimum_array_end "leetcode/minimum-array-end"
	reverse_integer "leetcode/reverse-integer"
	seat_reservation_manager "leetcode/seat-reservation-manager"
	the_latest_time_to_catch_a_bus "leetcode/the-latest-time-to-catch-a-bus"
	time_needed_to_buy_tickets "leetcode/time-needed-to-buy-tickets"
	two_sum "leetcode/two-sum"
	zigzag_conversion "leetcode/zigzag-conversion"
	"math"
)

func main() {
	test1845()
}

func test1845() {
	seatManager := seat_reservation_manager.Constructor(5)
	fmt.Println(seatManager.Reserve())
	fmt.Println(seatManager.Reserve())
	seatManager.Unreserve(2)
	fmt.Println(seatManager.Reserve())
	fmt.Println(seatManager.Reserve())
	fmt.Println(seatManager.Reserve())
	fmt.Println(seatManager.Reserve())
	seatManager.Unreserve(5)
}

func test2073() {
	fmt.Println(time_needed_to_buy_tickets.TimeRequiredToBuy([]int{2, 3, 2}, 2))
	fmt.Println(time_needed_to_buy_tickets.TimeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}

func test2414() {
	fmt.Println(length_of_the_longest_alphabetical_continuous_substring.LongestContinuousSubstring("abacaba"))
	fmt.Println(length_of_the_longest_alphabetical_continuous_substring.LongestContinuousSubstring("abcde"))
	fmt.Println(length_of_the_longest_alphabetical_continuous_substring.LongestContinuousSubstring("a"))
	fmt.Println(length_of_the_longest_alphabetical_continuous_substring.LongestContinuousSubstring("aaaaabcdeaaaaa"))
	fmt.Println(length_of_the_longest_alphabetical_continuous_substring.LongestContinuousSubstring("abcdefghijklmnopqrstuvwxyza"))
}

func test2332() {
	fmt.Println(the_latest_time_to_catch_a_bus.LatestTimeCatchTheBus([]int{10, 20}, []int{2, 17, 18, 19}, 2))
	fmt.Println(the_latest_time_to_catch_a_bus.LatestTimeCatchTheBus([]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2))
	fmt.Println(the_latest_time_to_catch_a_bus.LatestTimeCatchTheBus([]int{11}, []int{9, 11}, 1))
	fmt.Println(the_latest_time_to_catch_a_bus.LatestTimeCatchTheBus([]int{11}, []int{11, 13}, 2))
}
func test2552() {
	fmt.Println(count_increasing_quadruplets.CountQuadruplets([]int{1, 3, 2, 4, 5}))
	fmt.Println(count_increasing_quadruplets.CountQuadruplets([]int{1, 2, 3, 4}))
}

func test3176() {
	fmt.Println(find_the_maximum_length_of_a_good_subsequence_i.MaximumLength([]int{1, 2, 1, 1, 3}, 2))
	fmt.Println(find_the_maximum_length_of_a_good_subsequence_i.MaximumLength([]int{1, 2, 3, 4, 5, 1}, 0))
	fmt.Println(find_the_maximum_length_of_a_good_subsequence_i.MaximumLength([]int{1}, 0))
	fmt.Println(find_the_maximum_length_of_a_good_subsequence_i.MaximumLength([]int{1}, 1))
}

func test3174() {
	fmt.Println(clear_digits.ClearDigits("abc"))
	fmt.Println(clear_digits.ClearDigits("cb34"))
	fmt.Println(clear_digits.ClearDigits("a1"))
	fmt.Println(clear_digits.ClearDigits("a1b2"))
	fmt.Println(clear_digits.ClearDigits("a1bc2"))
	fmt.Println(clear_digits.ClearDigits("a1bc23"))
}

func test7() {
	fmt.Println(reverse_integer.Reverse(120))
	fmt.Println(reverse_integer.Reverse(-123))
	fmt.Println(reverse_integer.Reverse(2147483621))
	fmt.Println(reverse_integer.Reverse(-2147483621))
	fmt.Println(reverse_integer.Reverse(7463847412))
	fmt.Println(reverse_integer.Reverse(8463847412))
	fmt.Println(reverse_integer.Reverse(-8463847412))
	fmt.Println(reverse_integer.Reverse(-9463847412))
	fmt.Println(reverse_integer.Reverse(-8463847413))
	fmt.Println(math.MaxInt32, reverse_integer.Reverse(math.MaxInt32))
	fmt.Println(math.MinInt32, reverse_integer.Reverse(math.MinInt32))
}

func test2680() {
	fmt.Println(happy_students.CountWays([]int{0}))
	fmt.Println(happy_students.CountWays([]int{0, 1}))
	fmt.Println(happy_students.CountWays([]int{1, 1}))
	fmt.Println(happy_students.CountWays([]int{0, 1, 2}))
	fmt.Println(happy_students.CountWays([]int{0, 3, 2, 2}))
	fmt.Println(happy_students.CountWays([]int{6, 0, 3, 3, 6, 7, 2, 7}))

}

func test6() {
	fmt.Println(zigzag_conversion.Convert("PAYPALISHIRING", 3))
	fmt.Println(zigzag_conversion.Convert("PAYPALISHIRING", 4))
}

func test2708() {
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{3, -1, -5, 2, 5, -9}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{-4, -5, -4}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{-1}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{-2}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{0}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{-2, 0}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{0, 0}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{-9, 1}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{9, 0}))
	fmt.Println(maximum_strength_of_a_group.MaxStrength([]int{-9, -1, 0, 2}))
}

func test2024() {
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("T", 1))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("F", 1))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("TT", 1))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("TT", 2))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("FF", 1))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("FF", 2))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("TF", 1))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("TF", 2))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("TFFT", 1))
	fmt.Println(maximize_the_confusion_of_an_exam.MaxConsecutiveAnswers("TFFT", 2))
}

func test5() {
	fmt.Println(longest_palindromic_substring.LongestPalindrome("a"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("ab"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("aba"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("abba"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("aa"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("aaa"))
	fmt.Println(longest_palindromic_substring.LongestPalindrome("aaaa"))
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
