// 1656. 设计有序流
// https://leetcode.cn/problems/design-an-ordered-stream/

package design_an_ordered_stream

type OrderedStream struct {
	data []string
	ptr  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		data: make([]string, n),
		ptr:  1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey-1] = value
	var r []string
	if this.ptr == idKey {
		for _, v := range this.data[this.ptr-1:] {
			if v != "" {
				r = append(r, v)
			} else {
				break
			}
		}
		this.ptr += len(r)
	}
	return r
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.insert(idKey,value);
 */
