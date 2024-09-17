package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"log"
	"strings"
)

func CreateProduct(product model.Product) int {
	product.Id = 0
	isValid := validate(product)
	if !isValid {
		log.Println("Invalid near to create product service!")
		return 0
	}
	return dao.CreateProduct(product)
}

func UpdateProduct(product model.Product) bool {
	isValid := !isExist(product.Name)
	if !isValid {
		log.Println("Invalid near to update product service!")
		return false
	}
	return dao.UpdateProduct(product)
}

func DeleteProductById(id int) bool {
	isDeleted := dao.DeleteById(id)
	if !isDeleted {
		log.Printf("Delete product by id %d error!", id)
	}
	return isDeleted
}
func DeleteProductByName(name string) bool {
	isDeleted := dao.DeleteByName(name)
	if !isDeleted {
		log.Printf("Delete product by name %s error!", name)
	}
	return isDeleted
}

func GetProductById(id int) model.Product {
	return dao.SelectById(id)
}

func GetProductByName(name string) model.Product {
	return dao.SelectByName(name)
}

func GetAllProducts() []model.Product {
	products := dao.SelectAll()
	if products == nil {
		log.Println("No products found!")
	}
	return products
}

func validate(product model.Product) bool {
	return !isExist(product.Name) &&
		strings.TrimSpace(product.Name) != ""
}

func IsExist(name string) bool {
	return isExist(name)
}

func isExist(name string) bool {
	products := dao.SelectAll()
	for _, product := range products {
		if product.Name == name {
			return true
		}
	}
	return false
}

func isExistId(id uint64) bool {
	products := dao.SelectAll()
	for _, product := range products {
		if product.Id == id {
			return true
		}
	}
	return false
}
