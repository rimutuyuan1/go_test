package wallet

import "fmt"

type Wallet struct {
	balance int
}

//当正常调用Deposit方法时，w为调用方传递过来的副本，&w.balance和调用方的&w.balance内存地址是不同的
//这种方式为值传递
//*Wallet：指向Wallet内存地址的指针，通过指针调用，能直接修改对应内存地址中的值， 而不是改变调用方的副本的值
func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in test is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
