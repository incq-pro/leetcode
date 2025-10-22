// 2011. 执行操作后的变量值
// https://leetcode.cn/problems/final-value-of-variable-after-performing-operations

package final_value_of_variable_after_performing_operations

func finalValueAfterOperations(operations []string) int {
	v := 0
	for _, o := range operations {
		if o == "++X" || o == "X++" {
			v++
		} else {
			v--
		}
	}
	return v
}

func FinalValueAfterOperations(operations []string) int {
	return finalValueAfterOperations(operations)
}
