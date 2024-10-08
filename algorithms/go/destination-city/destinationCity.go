// 1436. 旅行终点站
// https://leetcode.cn/problems/destination-city/

package destination_city

func destCity(paths [][]string) string {
	m := make(map[string]string, len(paths))
	for _, x := range paths {
		m[x[0]] = x[1]
	}
	s := paths[0][0]
	for {
		if d, ok := m[s]; !ok {
			return s
		} else {
			s = d
		}
	}
	return ""
}

func DestCity(paths [][]string) string {
	return destCity(paths)
}
