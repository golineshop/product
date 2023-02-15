package service

import (
	"github.com/golineshop/product/domain/model"
	"github.com/golineshop/product/domain/repository"
)

type IProductService interface {
	AddProduct(product *model.Product) (ID int64, err error)
	DeleteProduct(ID int64) (err error)
	UpdateProduct(product *model.Product) (err error)
	FindProductByID(ID int64) (product *model.Product, err error)
	FindAllProduct() (productList []model.Product, err error)
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{ProductRepository: productRepository}
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func (p *ProductService) AddProduct(product *model.Product) (ID int64, err error) {
	return p.ProductRepository.CreateProduct(product)
}

func (p *ProductService) DeleteProduct(ID int64) (err error) {
	return p.ProductRepository.DeleteProductByID(ID)
}

func (p *ProductService) UpdateProduct(product *model.Product) (err error) {
	return p.ProductRepository.UpdateProduct(product)
}

func (p *ProductService) FindProductByID(ID int64) (product *model.Product, err error) {
	return p.ProductRepository.FindProductByID(ID)
}

func (p *ProductService) FindAllProduct() (productList []model.Product, err error) {
	return p.ProductRepository.FindAll()
}
