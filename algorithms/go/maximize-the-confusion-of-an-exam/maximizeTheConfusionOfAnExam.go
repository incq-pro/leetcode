package maximize_the_confusion_of_an_exam

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveAnswersOfA(answerKey, k, 'T'), maxConsecutiveAnswersOfA(answerKey, k, 'F'))
}

func maxConsecutiveAnswersOfA(answerKey string, k int, a byte) (ans int) {
	left := 0
	count := 0
	for right := range answerKey {
		if answerKey[right] != a {
			count++
		}
		for count > k {
			if answerKey[left] != a {
				count--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func MaxConsecutiveAnswers(answerKey string, k int) int {
	fmt.Println(answerKey, k)
	return maxConsecutiveAnswers(answerKey, k)
}
