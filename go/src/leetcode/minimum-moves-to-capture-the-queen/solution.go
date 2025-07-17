// 3001. 捕获黑皇后需要的最少移动次数
// https://leetcode.cn/problems/minimum-moves-to-capture-the-queen/

package minimum_moves_to_capture_the_queen

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e {
		if c != a || d > max(b, f) || d < min(b, f) {
			return 1
		}
	} else if b == f {
		if d != b || c > max(a, e) || c < min(a, e) {
			return 1
		}
	} else if c+d == e+f {
		if c+d != a+b || a > max(c, e) || a < min(c, e) {
			return 1
		}
	} else if c-d == e-f {
		if a-b != c-d || a > max(c, e) || a < min(c, e) {
			return 1
		}
	}
	return 2
}

func MinMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	return minMovesToCaptureTheQueen(a, b, c, d, e, f)
}
