package service

import (
	"客户信息管理系统/v1/model"
)

type CustomerService struct {
	custers     []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {

	customerService := &CustomerService{}
	customerService.customerNum = 1
	costomer := model.NewCustomer(1, "张三", "男", 20, "112", "648960069@qq.com")
	customerService.custers = append(customerService.custers, costomer)
	return customerService
}
func (this *CustomerService) List() []model.Customer {
	return this.custers
}
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.custers = append(this.custers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {

	index := this.FindId(id)

	if index == -1 {
		return false
	}
	this.custers = append(this.custers[:index], this.custers[index+1:]...)
	return true
}
func (this *CustomerService) FindId(id int) int {
	index := -1

	for i := 0; i < len(this.custers); i++ {
		index = i
	}
	return index
}
