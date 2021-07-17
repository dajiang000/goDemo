package demo

import "fmt"

type account struct {
	accountID string
	balance   float64
	pwd       string
}

func NewAccount(accountID string) *account {
	return &account{accountID: accountID}
}

func (a *account) SetBalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("输入正确的余额数目")
	}
}

func (a *account) GetBalance() float64 {
	return a.balance
}

func (a *account) SetPwd(pwd string) {
	if len(pwd) == 6 {
		a.pwd = pwd
	} else {
		fmt.Println("输入的密码格式不正确")
	}
}
func (a *account) GetPwd() string {
	return a.pwd
}
