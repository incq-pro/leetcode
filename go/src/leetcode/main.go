package main

import (
	"fmt"
	best_time_to_buy_and_sell_stock "leetcode/best-time-to-buy-and-sell-stock"
	best_time_to_buy_and_sell_stock_ii "leetcode/best-time-to-buy-and-sell-stock-ii"
	clear_digits "leetcode/clear-digits"
	container_with_most_water "leetcode/container-with-most-water"
	count_increasing_quadruplets "leetcode/count-increasing-quadruplets"
	count_pairs_that_form_a_complete_day_i "leetcode/count-pairs-that-form-a-complete-day-i"
	count_prefixes_of_a_given_string "leetcode/count-prefixes-of-a-given-string"
	count_substrings_that_can_be_rearranged_to_contain_a_string_i "leetcode/count-substrings-that-can-be-rearranged-to-contain-a-string-i"
	count_the_hidden_sequences "leetcode/count-the-hidden-sequences"
	design_an_ordered_stream "leetcode/design-an-ordered-stream"
	design_browser_history "leetcode/design-browser-history"
	destination_city "leetcode/destination-city"
	final_array_state_after_k_multiplication_operations_i "leetcode/final-array-state-after-k-multiplication-operations-i"
	find_all_k_distant_indices_in_an_array "leetcode/find-all-k-distant-indices-in-an-array"
	find_closest_number_to_zero "leetcode/find-closest-number-to-zero"
	find_minimum_time_to_reach_last_room_i "leetcode/find-minimum-time-to-reach-last-room-i"
	find_subarray_with_bitwise_or_closest_to_k "leetcode/find-subarray-with-bitwise-or-closest-to-k"
	find_the_maximum_length_of_a_good_subsequence_i "leetcode/find-the-maximum-length-of-a-good-subsequence-i"
	find_the_maximum_length_of_valid_subsequence_i "leetcode/find-the-maximum-length-of-valid-subsequence-i"
	find_the_number_of_good_pairs_i "leetcode/find-the-number-of-good-pairs-i"
	find_the_number_of_good_pairs_ii "leetcode/find-the-number-of-good-pairs-ii"
	find_the_original_typed_string_i "leetcode/find-the-original-typed-string-i"
	gas_station "leetcode/gas-station"
	happy_students "leetcode/happy-students"
	interleaving_string "leetcode/interleaving-string"
	jump_game "leetcode/jump-game"
	jump_game_ii "leetcode/jump-game-ii"
	knight_dialer "leetcode/knight-dialer"
	length_of_the_longest_alphabetical_continuous_substring "leetcode/length-of-the-longest-alphabetical-continuous-substring"
	longest_palindromic_substring "leetcode/longest-palindromic-substring"
	longest_substring_without_repeating_characters "leetcode/longest-substring-without-repeating-characters"
	maximize_the_confusion_of_an_exam "leetcode/maximize-the-confusion-of-an-exam"
	maximum_consecutive_floors_without_special_floors "leetcode/maximum-consecutive-floors-without-special-floors"
	maximum_spending_after_buying_items "leetcode/maximum-spending-after-buying-items"
	maximum_strength_of_a_group "leetcode/maximum-strength-of-a-group"
	maximum_value_of_k_coins_from_piles "leetcode/maximum-value-of-k-coins-from-piles"
	minimum_array_end "leetcode/minimum-array-end"
	minimum_cost_to_cut_a_stick "leetcode/minimum-cost-to-cut-a-stick"
	minimum_moves_to_capture_the_queen "leetcode/minimum-moves-to-capture-the-queen"
	minimum_number_of_flips_to_make_binary_grid_palindromic_i "leetcode/minimum-number-of-flips-to-make-binary-grid-palindromic-i"
	minimum_operations_to_exceed_threshold_value_i "leetcode/minimum-operations-to-exceed-threshold-value-i"
	minimum_operations_to_exceed_threshold_value_ii "leetcode/minimum-operations-to-exceed-threshold-value-ii"
	number_of_changing_keys "leetcode/number-of-changing-keys"
	number_of_ways_to_split_array "leetcode/number-of-ways-to-split-array"
	palindrome_number "leetcode/palindrome-number"
	palindrome_partitioning "leetcode/palindrome-partitioning"
	palindrome_partitioning_ii "leetcode/palindrome-partitioning-ii"
	partition_array_such_that_maximum_difference_is_k "leetcode/partition-array-such-that-maximum-difference-is-k"
	partition_equal_subset_sum "leetcode/partition-equal-subset-sum"
	pascals_triangle_ii "leetcode/pascals-triangle-ii"
	remove_duplicate_letters "leetcode/remove-duplicate-letters"
	reverse_integer "leetcode/reverse-integer"
	seat_reservation_manager "leetcode/seat-reservation-manager"
	semi_ordered_permutation "leetcode/semi-ordered-permutation"
	shortest_distance_after_road_addition_queries_ii "leetcode/shortest-distance-after-road-addition-queries-ii"
	shortest_subarray_with_or_at_least_k_i "leetcode/shortest-subarray-with-or-at-least-k-i"
	smallest_range_i "leetcode/smallest-range-i"
	smallest_range_ii "leetcode/smallest-range-ii"
	string_to_integer_atoi "leetcode/string-to-integer-atoi"
	the_latest_time_to_catch_a_bus "leetcode/the-latest-time-to-catch-a-bus"
	time_needed_to_buy_tickets "leetcode/time-needed-to-buy-tickets"
	trapping_rain_water "leetcode/trapping-rain-water"
	"leetcode/triangle"
	two_sum "leetcode/two-sum"
	valid_word "leetcode/valid-word"
	zigzag_conversion "leetcode/zigzag-conversion"
	"math"
)

func main() {
	test3136()
}

func test3136() {
	fmt.Println(valid_word.IsValid("234Adas"))
	fmt.Println(valid_word.IsValid("b3"))
	fmt.Println(valid_word.IsValid("a3$e"))
	fmt.Println(valid_word.IsValid("UuE6"))

}

func test3201() {
	fmt.Println(find_the_maximum_length_of_valid_subsequence_i.MaximumLength([]int{1, 2, 3, 4}))
	fmt.Println(find_the_maximum_length_of_valid_subsequence_i.MaximumLength([]int{1, 2, 1, 1, 2, 1, 2}))
	fmt.Println(find_the_maximum_length_of_valid_subsequence_i.MaximumLength([]int{1, 2, 1, 1, 2, 1, 2}))
}

func test120() {
	fmt.Println(triangle.MinimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

func test119() {
	for i := 0; i < 30; i++ {
		fmt.Println(pascals_triangle_ii.GetRow(i))
	}
}

func test3330() {
	fmt.Println(find_the_original_typed_string_i.PossibleStringCount("abbcccc"))
	fmt.Println(find_the_original_typed_string_i.PossibleStringCount("abcd"))
	fmt.Println(find_the_original_typed_string_i.PossibleStringCount("aaaa"))
}

func test2200() {
	fmt.Println(find_all_k_distant_indices_in_an_array.FindKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
	fmt.Println(find_all_k_distant_indices_in_an_array.FindKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))
}

func test2294() {
	fmt.Println(partition_array_such_that_maximum_difference_is_k.PartitionArray([]int{3, 6, 1, 2, 5}, 2))
	fmt.Println(partition_array_such_that_maximum_difference_is_k.PartitionArray([]int{1, 2, 3}, 1))
	fmt.Println(partition_array_such_that_maximum_difference_is_k.PartitionArray([]int{2, 2, 4, 5}, 0))
	fmt.Println(partition_array_such_that_maximum_difference_is_k.PartitionArray([]int{1}, 1))
	fmt.Println(partition_array_such_that_maximum_difference_is_k.PartitionArray([]int{1, 1}, 0))
}

func test3341() {
	//fmt.Println(find_minimum_time_to_reach_last_room_i.MinTimeToReach([][]int{{0, 4}, {4, 4}}))
	//fmt.Println(find_minimum_time_to_reach_last_room_i.MinTimeToReach([][]int{{0, 0, 0}, {0, 0, 0}}))
	//fmt.Println(find_minimum_time_to_reach_last_room_i.MinTimeToReach([][]int{{0, 1}, {1, 2}}))
	//fmt.Println(find_minimum_time_to_reach_last_room_i.MinTimeToReach([][]int{{0, 88, 10}, {89, 41, 69}, {63, 39, 35}}))
	fmt.Println(find_minimum_time_to_reach_last_room_i.MinTimeToReach([][]int{{94, 79, 62, 27, 69, 84}, {6, 32, 11, 82, 42, 30}}))

}

func test2145() {
	fmt.Println(count_the_hidden_sequences.NumberOfArrays([]int{1, -3, 4}, 1, 6))
	fmt.Println(count_the_hidden_sequences.NumberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5))
}

func test416() {
	fmt.Println(partition_equal_subset_sum.CanPartition([]int{1, 5, 11, 5}))
	fmt.Println(partition_equal_subset_sum.CanPartition([]int{1, 2, 3, 5}))

}

func test132() {
	fmt.Println(palindrome_partitioning_ii.MinCut("aab"))
	fmt.Println(palindrome_partitioning_ii.MinCut("a"))
	fmt.Println(palindrome_partitioning_ii.MinCut("ab"))
}

func test2255() {
	fmt.Println(count_prefixes_of_a_given_string.CountPrefixes([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))
	fmt.Println(count_prefixes_of_a_given_string.CountPrefixes([]string{"a", "a"}, "aa"))

}
func test131() {
	fmt.Println(palindrome_partitioning.Partition("aab"))
	fmt.Println(palindrome_partitioning.Partition("a"))
}

func test316() {
	println(remove_duplicate_letters.RemoveDuplicateLetters("bcabc"))
	println(remove_duplicate_letters.RemoveDuplicateLetters("cbacdcbc"))
}

func test1472() {
	obj := design_browser_history.Constructor("https://www.1.com")
	obj.Visit("https://2.com")
	println(obj.Back(1))
	println(obj.Forward(1))
	println("--")
	obj.Visit("https://3.com")
	println(obj.Back(1))
	obj.Visit("https://4.com")
	println(obj.Back(1))
	println(obj.Forward(1))

	//obj.Visit("https://qq1.com")
	//obj.Visit("https://qq2.com")
	//obj.Visit("https://qq3.com")
	//obj.Visit("https://qq4.com")
	//obj.Visit("https://qq5.com")
	//println(obj.Back(1))
	//println(obj.Back(1))
	//println(obj.Back(1))
	//obj.Visit("https://qq6.com")
	//println(obj.Back(1))
	//println(obj.Forward(1))
}

func test42() {
	println(trapping_rain_water.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	println(trapping_rain_water.Trap([]int{4, 2, 0, 3, 2, 5}))
}

func test1656() {
	o := design_an_ordered_stream.Constructor(5)
	fmt.Println(o.Insert(3, "ccccc"))
	fmt.Println(o.Insert(1, "aaaaa"))
	fmt.Println(o.Insert(2, "bbbbb"))
	fmt.Println(o.Insert(5, "eeeee"))
	fmt.Println(o.Insert(4, "ddddd"))
}

func test2218() {
	fmt.Println(maximum_value_of_k_coins_from_piles.NaxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
	fmt.Println(maximum_value_of_k_coins_from_piles.NaxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7))

}

func test134() {
	fmt.Println(gas_station.CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(gas_station.CanCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(gas_station.CanCompleteCircuit([]int{2}, []int{3}))
	fmt.Println(gas_station.CanCompleteCircuit([]int{3}, []int{2}))
}

func test2239() {
	fmt.Println(find_closest_number_to_zero.FindClosestNumber([]int{-4, -2, 1, 4, 8}))
	fmt.Println(find_closest_number_to_zero.FindClosestNumber([]int{2, -1, 1}))
	fmt.Println(find_closest_number_to_zero.FindClosestNumber([]int{2, 0, 1}))
}

func test3095() {
	//fmt.Println(shortest_subarray_with_or_at_least_k_i.MinimumSubarrayLength([]int{1, 2, 3}, 2))
	//fmt.Println(shortest_subarray_with_or_at_least_k_i.MinimumSubarrayLength([]int{2, 1, 8}, 10))
	//fmt.Println(shortest_subarray_with_or_at_least_k_i.MinimumSubarrayLength([]int{1, 2}, 0))
	fmt.Println(shortest_subarray_with_or_at_least_k_i.MinimumSubarrayLength([]int{1, 12, 2, 5}, 43))
}

func test3066() {
	fmt.Println(minimum_operations_to_exceed_threshold_value_ii.MinOperations([]int{2, 11, 10, 1, 3}, 10))
	fmt.Println(minimum_operations_to_exceed_threshold_value_ii.MinOperations([]int{1, 1, 2, 4, 9}, 20))
}

func test3065() {
	fmt.Println(minimum_operations_to_exceed_threshold_value_i.MinOperations([]int{2, 11, 10, 1, 3}, 10))
	fmt.Println(minimum_operations_to_exceed_threshold_value_i.MinOperations([]int{1, 1, 2, 4, 9}, 1))
}

func test2270() {
	fmt.Println(number_of_ways_to_split_array.WaysToSplitArray([]int{10, 4, -8, 7}))
	fmt.Println(number_of_ways_to_split_array.WaysToSplitArray([]int{2, 3, 1, 0}))

}

func test3298() {
	fmt.Println(count_substrings_that_can_be_rearranged_to_contain_a_string_i.ValidSubstringCount("ab", "abc"))
	fmt.Println(count_substrings_that_can_be_rearranged_to_contain_a_string_i.ValidSubstringCount("bcca", "abc"))
	fmt.Println(count_substrings_that_can_be_rearranged_to_contain_a_string_i.ValidSubstringCount("abcabc", "abc"))
	fmt.Println(count_substrings_that_can_be_rearranged_to_contain_a_string_i.ValidSubstringCount("abcabc", "aaabc"))
}

func test3019() {
	fmt.Println(number_of_changing_keys.CountKeyChanges("aAbBcC"))
	fmt.Println(number_of_changing_keys.CountKeyChanges("AaAaAaaA"))
}

func test2274() {
	fmt.Println(maximum_consecutive_floors_without_special_floors.MaxConsecutive(2, 9, []int{4, 6}))
	fmt.Println(maximum_consecutive_floors_without_special_floors.MaxConsecutive(6, 8, []int{7, 6, 8}))
}

func test97() {
	fmt.Println(interleaving_string.IsInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(interleaving_string.IsInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(interleaving_string.IsInterleave("", "", ""))
	fmt.Println(interleaving_string.IsInterleave("abc", "", "abc"))
	fmt.Println(interleaving_string.IsInterleave("", "abc", "abc"))
	fmt.Println(interleaving_string.IsInterleave("db", "b", "cbb"))
}

func test3264() {
	fmt.Println(final_array_state_after_k_multiplication_operations_i.GetFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
	fmt.Println(final_array_state_after_k_multiplication_operations_i.GetFinalState([]int{1, 2}, 3, 4))
}

func test122() {
	fmt.Println(best_time_to_buy_and_sell_stock_ii.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(best_time_to_buy_and_sell_stock_ii.MaxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(best_time_to_buy_and_sell_stock_ii.MaxProfit([]int{7, 6, 4, 3, 1}))
}

func test121() {
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 6, 4, 3, 1}))
}

func test45() {
	fmt.Println(jump_game_ii.Jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump_game_ii.Jump([]int{2, 3, 0, 1, 4}))
}

func test2931() {
	fmt.Println(maximum_spending_after_buying_items.MaxSpending([][]int{{8, 5, 2}, {6, 4, 1}, {9, 7, 3}}))
	fmt.Println(maximum_spending_after_buying_items.MaxSpending([][]int{{10, 8, 6, 4, 2}, {9, 7, 5, 3, 2}}))
}

func test11() {
	fmt.Println(container_with_most_water.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(container_with_most_water.MaxArea([]int{1, 1}))
}

func test2717() {
	fmt.Println(semi_ordered_permutation.SemiOrderedPermutation([]int{2, 1, 4, 3}))
	fmt.Println(semi_ordered_permutation.SemiOrderedPermutation([]int{2, 4, 1, 3}))
	fmt.Println(semi_ordered_permutation.SemiOrderedPermutation([]int{1, 3, 4, 2, 5}))
}

func test55() {
	fmt.Println(jump_game.CanJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump_game.CanJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(jump_game.CanJump([]int{1}))
	fmt.Println(jump_game.CanJump([]int{0}))
	fmt.Println(jump_game.CanJump([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
}

func test935() {
	//fmt.Println(knight_dialer.KnightDialer(1))
	//fmt.Println(knight_dialer.KnightDialer(2))
	fmt.Println(knight_dialer.KnightDialer(3131))
}

func test3001() {
	fmt.Println(minimum_moves_to_capture_the_queen.MinMovesToCaptureTheQueen(1, 1, 8, 8, 2, 3))
	fmt.Println(minimum_moves_to_capture_the_queen.MinMovesToCaptureTheQueen(5, 3, 3, 4, 5, 2))
	fmt.Println(minimum_moves_to_capture_the_queen.MinMovesToCaptureTheQueen(1, 1, 1, 2, 1, 3))
	fmt.Println(minimum_moves_to_capture_the_queen.MinMovesToCaptureTheQueen(1, 2, 1, 1, 1, 3))
	fmt.Println(minimum_moves_to_capture_the_queen.MinMovesToCaptureTheQueen(2, 2, 1, 1, 4, 4))
}

func test3244() {
	fmt.Println(shortest_distance_after_road_addition_queries_ii.ShortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
	fmt.Println(shortest_distance_after_road_addition_queries_ii.ShortestDistanceAfterQueries(4, [][]int{{0, 3}, {0, 2}}))
	fmt.Println(shortest_distance_after_road_addition_queries_ii.ShortestDistanceAfterQueries(100, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {5, 99}, {0, 99}}))
}
func test3239() {
	fmt.Println(minimum_number_of_flips_to_make_binary_grid_palindromic_i.MinFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(minimum_number_of_flips_to_make_binary_grid_palindromic_i.MinFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))
}

func test1547() {
	fmt.Println(minimum_cost_to_cut_a_stick.MinCost(7, []int{1, 3, 4, 5}))
	fmt.Println(minimum_cost_to_cut_a_stick.MinCost(9, []int{5, 6, 1, 4, 2}))
}

func test3184() {
	fmt.Println(count_pairs_that_form_a_complete_day_i.CountCompleteDayPairs([]int{12, 12, 30, 24, 24}))
	fmt.Println(count_pairs_that_form_a_complete_day_i.CountCompleteDayPairs([]int{72, 48, 24, 3}))
}

func test9080() {
	fmt.Println(smallest_range_i.SmallestRangeI([]int{1}, 0))
	fmt.Println(smallest_range_i.SmallestRangeI([]int{0, 10}, 2))
	fmt.Println(smallest_range_i.SmallestRangeI([]int{1, 3, 6}, 3))
}

func test910() {
	fmt.Println(smallest_range_ii.SmallestRangeII([]int{1}, 0))
	fmt.Println(smallest_range_ii.SmallestRangeII([]int{0, 10}, 2))
	fmt.Println(smallest_range_ii.SmallestRangeII([]int{1, 3, 6}, 3))
}

func test9() {
	fmt.Println(palindrome_number.IsPalindrome(-11))
	fmt.Println(palindrome_number.IsPalindrome(-2))
	fmt.Println(palindrome_number.IsPalindrome(-1))
	fmt.Println(palindrome_number.IsPalindrome(0))
	fmt.Println(palindrome_number.IsPalindrome(1))
	fmt.Println(palindrome_number.IsPalindrome(2))
	fmt.Println(palindrome_number.IsPalindrome(10))
	fmt.Println(palindrome_number.IsPalindrome(11))
	fmt.Println(palindrome_number.IsPalindrome(12))
	fmt.Println(palindrome_number.IsPalindrome(121))
	fmt.Println(palindrome_number.IsPalindrome(121))
}

func test3164() {
	fmt.Println(find_the_number_of_good_pairs_ii.NumberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(find_the_number_of_good_pairs_ii.NumberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3))
}

func test3162() {
	fmt.Println(find_the_number_of_good_pairs_i.NumberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(find_the_number_of_good_pairs_i.NumberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3))
}

func test8() {
	fmt.Println(string_to_integer_atoi.MyAtoi("42"))
	fmt.Println(string_to_integer_atoi.MyAtoi(" -042"))
	fmt.Println(string_to_integer_atoi.MyAtoi("1337c0d3"))
	fmt.Println(string_to_integer_atoi.MyAtoi("0-1"))
	fmt.Println(string_to_integer_atoi.MyAtoi("words and 987"))
}

func test3171() {
	fmt.Println(find_subarray_with_bitwise_or_closest_to_k.MinimumDifference([]int{1}, 1))
	fmt.Println(find_subarray_with_bitwise_or_closest_to_k.MinimumDifference([]int{1}, 2))
	fmt.Println(find_subarray_with_bitwise_or_closest_to_k.MinimumDifference([]int{1, 2}, 2))
	fmt.Println(find_subarray_with_bitwise_or_closest_to_k.MinimumDifference([]int{1, 2}, 3))
}

func test1436() {
	fmt.Println(destination_city.DestCity([][]string{{"A", "B"}, {"B", "C"}}))
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
