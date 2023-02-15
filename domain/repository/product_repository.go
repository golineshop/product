package repository

import "github.com/golineshop/product/domain/model"
import "github.com/jinzhu/gorm"

type IProductRepository interface {
	InitTable() (err error)
	FindProductByID(ID int64) (product *model.Product, err error)
	CreateProduct(product *model.Product) (ID int64, err error)
	DeleteProductByID(ID int64) (err error)
	UpdateProduct(product *model.Product) (err error)
	FindAll() (productList []model.Product, err error)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

func (p *ProductRepository) InitTable() (err error) {
	return p.mysqlDb.CreateTable(&model.Product{}, &model.ProductSeo{}, &model.ProductImage{}, &model.ProductSize{}).Error
}

func (p *ProductRepository) FindProductByID(ID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, p.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product, ID).Error
}

func (p *ProductRepository) CreateProduct(product *model.Product) (ID int64, err error) {
	return product.ID, p.mysqlDb.Create(product).Error
}

func (p *ProductRepository) DeleteProductByID(ID int64) (err error) {
	// 涉及到表关联删除，是多sql操作，需要开启事务
	tx := p.mysqlDb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Unscoped().Where("id = ?", ID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("images_product_id = ?", ID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", ID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", ID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (p *ProductRepository) UpdateProduct(product *model.Product) (err error) {
	return p.mysqlDb.Model(product).Update(product).Error
}

func (p *ProductRepository) FindAll() (productList []model.Product, err error) {
	return productList, p.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productList).Error
}
