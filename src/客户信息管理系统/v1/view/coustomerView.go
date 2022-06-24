package main

import (
	"fmt"
	"客户信息管理系统/v1/model"
	"客户信息管理系统/v1/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

func (this *customerView) exit() {
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
func (this *customerView) delete() {
	fmt.Println("-------------------删除客户----------------------")
	fmt.Println("请选择待删除客户的编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(y/n):")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("--------------------删除完成-------------")
		} else {
			fmt.Println("--------------------删除失败-------------")
		}
	}
}
func (this *customerView) add() {
	fmt.Println("---------------添加客户-------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("-------------添加成功---------")
	} else {

		fmt.Println("-------------添加失败---------")
	}
}

func (this *customerView) list() {

	customers := this.customerService.List()
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i := 1; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n------------客户列表完成------------------\n")
}

func (this *customerView) mainMenu() {

	for {
		fmt.Println("---------------客户信息管理软件----------------")
		fmt.Println("                      1:添加客户")
		fmt.Println("                      2:修改客户")
		fmt.Println("                      3:删除客户")
		fmt.Println("                      4:客户列表")
		fmt.Println("                      5:退    出")
		fmt.Print("请现在(1-5):")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {

	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.mainMenu()
}
