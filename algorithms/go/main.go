package main

import (
	"fmt"
	"leetcode/two_sum"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	fmt.Println(two_sum.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(two_sum.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(two_sum.TwoSum([]int{3, 3}, 6))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
