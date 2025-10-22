// 3147. 从魔法师身上吸取的最大能量
// https://leetcode.cn/problems/taking-maximum-energy-from-the-mystic-dungeon/

package taking_maximum_energy_from_the_mystic_dungeon

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	m := energy[n-1]
	s := make([]int, k)
	for i := n - 1; i >= 0; i-- {
		r := i % k
		s[r] += energy[i]
		m = max(m, s[r])
	}
	return m
}

func MaximumEnergy(energy []int, k int) int {
	return maximumEnergy(energy, k)
}
