package controller

import (
	"ancora/dao"
	"ancora/helper"
	"ancora/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	search := c.Query("search")
	sortOrder := c.Query("sort")
	products, _ := dao.GetProducts()

	if search != "" {
		var root *model.Node
		for _, p := range products {
			root = helper.Insert(root, p)
		}
		var results []model.Product
		helper.Autocomplete(root, search, &results)
		products = results
	}

	if sortOrder == "asc" {
		products = helper.MergeSort(products, true)
	} else if sortOrder == "desc" {
		products = helper.MergeSort(products, false)
	}

	if c.Query("k-type") == "min" {
		products = helper.GetKMinProducts(products, 5)
	} else if c.Query("k-type") == "max" {
		products = helper.GetKMaxProducts(products, 5)
	}
	context := gin.H{
		"Products": products,
	}
	c.HTML(200, "index.html", context)
}
