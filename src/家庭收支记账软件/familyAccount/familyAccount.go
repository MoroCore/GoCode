package main

import (
	"fmt"
)

type familyAccount struct {
	key      string
	balance  float64
	money    float64
	note     string
	flag     bool
	details  string
	loop     bool
	username string
	password string
}

func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说	明",
		username: "moro",
		password: "123456",
	}
}

func (this *familyAccount) showDetails() {
	fmt.Println("---------------------当前收支明细记录----------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧!")
	}
}
func (this *familyAccount) income() {
	fmt.Println("登记收入")
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入的说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t %v \t %v \t %v", this.balance, this.money, this.note)
	this.flag = true
}
func (this *familyAccount) pay() {
	fmt.Println("登记支出")
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出的说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t %v \t%v \t %v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *familyAccount) exit() {
	fmt.Println("你确定有退出吗？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误,请重新输入y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}
func (this *familyAccount) Transfer() {
	fmt.Println("登记转账信息")
	fmt.Println("本次转账金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次转账的说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t %v \t%v \t %v", this.balance, this.money, this.note)
	this.flag = true
}
func (this *familyAccount) login() bool {
	username := ""
	password := ""
	fmt.Println("请输入你的账户和密码")
	fmt.Scanln(&username)
	fmt.Scanln(&password)
	if this.username == username && this.password == password {
		return true
	} else {
		fmt.Println(username, password)
		return false
	}
}
func (this *familyAccount) MainMenu() {

	login := this.login()
	if !login {
		return
	}
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------------")
		fmt.Println("                      1:收支明细")
		fmt.Println("                      2:登记收入")
		fmt.Println("                      3:登记支出")
		fmt.Println("                      4:转账")
		fmt.Println("                      5:退出软件")
		fmt.Print("请选择(1-5): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":

		case "5":
			this.exit()
		default:
			fmt.Println("请输入正确的选项......")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}
func main() {

	familyAccount := NewFamilyAccount()
	familyAccount.MainMenu()
}
