// 1472. 设计浏览器历史记录
// https://leetcode.cn/problems/design-browser-history/

package design_browser_history

type BrowserHistory struct {
	back, forward []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		back:    []string{homepage},
		forward: []string{},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.forward = this.forward[0:0]
	this.back = append(this.back, url)
}

func (this *BrowserHistory) Back(steps int) string {
	b := this.back
	f := this.forward
	n := len(b) - 1
	if steps > n {
		steps = n
	}
	for i := 1; i <= steps; i++ {
		f = append(f, b[len(b)-i])
	}
	b = b[0 : len(b)-steps]
	this.back = b
	this.forward = f
	return this.back[len(this.back)-1]

}

func (this *BrowserHistory) Forward(steps int) string {
	b := this.back
	f := this.forward
	n := len(f)
	if steps > n {
		steps = n
	}
	for i := 1; i <= steps; i++ {
		b = append(b, f[len(f)-i])
	}
	this.back = b
	this.forward = f[0 : len(f)-steps]
	return this.back[len(this.back)-1]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
