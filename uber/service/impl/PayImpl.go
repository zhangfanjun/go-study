package impl

type WeChat struct {
	Dollar float64
}

/**
* 采用值调用
 */
func (w WeChat) Pay(money float64) {
	w.Dollar -= money
}

/**
* 采用地址调用
 */
func (w WeChat) Show() float64 {
	return w.Dollar
}

type Alipay struct {
	Dollar float64
}

/**
* 采用地址调用
 */
func (a *Alipay) Pay(money float64) {
	a.Dollar -= money
}
