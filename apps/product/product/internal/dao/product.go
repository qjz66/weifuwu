package dao

import "tiny_service/apps/product/product/internal/models"

func AddProduct(product *models.Product) error {
	err := models.ProductMgr(gdb).Save(product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductInfo(productId int64) (product models.Product, err error) {
	product, err = models.ProductMgr(gdb).GetFromID(uint64(productId))
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func SelectProductPage(newPage models.IPage) (results []*models.Product, total int64, err error) {
	db := models.ProductMgr(gdb).Find(&results)
	if db.Error != nil {
		return nil, 0, err
	}
	db.Count(&total)
	if len(newPage.GetOrederItemsString()) > 0 {
		db = db.Order(newPage.GetOrederItemsString())
	}
	err = db.Limit(int(newPage.GetSize())).Offset(int(newPage.Offset())).Find(&results).Error
	return results, total, nil
}
