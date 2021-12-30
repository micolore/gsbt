package main

import "fmt"

type Account interface {
	Query(id string) int
	Update(id string, value int)
}

type AccountImpl struct {
	Id    string
	Name  string
	Value int
}

func (a *AccountImpl) Query(_ string) int {
	fmt.Println("AccountImpl.Query")
	return 100
}

func (a *AccountImpl) Update(_ string, _ int) {
	fmt.Println("AccountImpl.Update")
}

var New = func(id, name string, value int) Account {
	return &AccountImpl{id, name, value}
}

type Proxy struct {
	Account Account
}

func (p *Proxy) Query(id string) int {
	fmt.Println("Proxy.Query begin------->")
	value := p.Account.Query(id)
	fmt.Println("Proxy.Query end<---------")
	return value
}

func (p *Proxy) Update(id string, value int) {
	fmt.Println("Proxy.Update begin------->")
	p.Account.Update(id, value)
	fmt.Println("Proxy.Update end<---------")
}


func init() {
	New = func(id, name string, value int) Account {
		a := &AccountImpl{id, name, value}
		p := &Proxy{a}
		return p
	}
}

// 静态代理
func main() {
	id := "100111"
	a := New(id, "ZhangSan", 100)
	a.Query(id)
	a.Update(id, 500)
}
