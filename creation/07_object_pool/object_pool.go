package objectpool

import (
	"fmt"
	"strconv"
)

type Proxy struct {
	Account  string
	Password string
}

func (p *Proxy) Show() {
	fmt.Printf("Proxy account %s, password %s\n", p.Account, p.Password)
}

type ProxyPool chan *Proxy

func NewProxyPool(total int) ProxyPool {
	p := make(ProxyPool, total)
	for i := 0; i < total; i++ {
		p <- &Proxy{
			Account:  strconv.Itoa(i),
			Password: strconv.Itoa(i),
		}
	}
	return p
}
