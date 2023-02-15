package handler

import (
	"context"
	"github.com/golineshop/product/common"
	"github.com/golineshop/product/domain/model"
	"github.com/golineshop/product/domain/service"
	proto "github.com/golineshop/product/proto"
)

type ProductController struct {
	ProductService service.IProductService
}

// AddProduct 添加商品
func (p *ProductController) AddProduct(ctx context.Context, request *proto.ProductInfo, response *proto.ResponseProduct) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	productID, err := p.ProductService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	response.ProductId = productID
	return nil
}

// FindProductByID 根据ID查找商品
func (p *ProductController) FindProductByID(ctx context.Context, request *proto.RequestID, response *proto.ProductInfo) error {
	productData, err := p.ProductService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, response); err != nil {
		return err
	}
	return nil
}

// UpdateProduct 商品更新
func (p *ProductController) UpdateProduct(ctx context.Context, request *proto.ProductInfo, response *proto.Response) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	err := p.ProductService.UpdateProduct(productAdd)
	if err != nil {
		return err
	}
	response.Msg = "更新成功"
	return nil
}

// DeleteProductByID 根据ID删除对应商品
func (p *ProductController) DeleteProductByID(ctx context.Context, request *proto.RequestID, response *proto.Response) error {
	if err := p.ProductService.DeleteProduct(request.ProductId); err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

// FindAllProduct 查找所有商品
func (p *ProductController) FindAllProduct(ctx context.Context, request *proto.RequestAll, response *proto.AllProduct) error {
	productAll, err := p.ProductService.FindAllProduct()
	if err != nil {
		return err
	}

	for _, v := range productAll {
		productInfo := &proto.ProductInfo{}
		err := common.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}
