// 1845. 座位预约管理系统
// https://leetcode.cn/problems/seat-reservation-manager

package seat_reservation_manager

type SeatManager struct {
	array []int
}

func Constructor(n int) SeatManager {
	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	return SeatManager{array: a}
}

func (this *SeatManager) Reserve() int {
	a := this.array
	n := len(a) - 1
	x := a[0]
	if n != 0 {
		a[0] = a[n]
		// down
		i := 0
		for {
			j1 := i*2 + 1
			if j1 >= n || j1 < 0 {
				break
			}
			j := j1
			if j2 := j1 + 1; j2 < n && a[j2] < a[j1] {
				j = j2
			}
			if a[i] <= a[j] {
				break
			}
			a[i], a[j] = a[j], a[i]
			i = j
		}
	}
	this.array = a[0:n]
	return x
}

func (this *SeatManager) Unreserve(seatNumber int) {
	a := this.array
	a = append(a, seatNumber)
	n := len(a)
	j := n - 1
	for {
		i := (j - 1) / 2
		if i == j || a[i] <= a[j] {
			break
		}
		a[i], a[j] = a[j], a[i]
		j = i
	}
	this.array = a
}
