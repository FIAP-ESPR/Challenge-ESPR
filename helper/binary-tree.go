package helper

import (
	"ancora/model"
	"strings"
)

func Insert(root *model.Node, product model.Product) *model.Node {
	if root == nil {
		return &model.Node{Product: product}
	}
	if strings.ToLower(product.Name) < strings.ToLower(root.Product.Name) {
		root.Left = Insert(root.Left, product)
	} else {
		root.Right = Insert(root.Right, product)
	}
	return root
}

func Autocomplete(root *model.Node, prefix string, result *[]model.Product) {
	if root == nil || len(*result) >= 5 {
		return
	}

	if strings.HasPrefix(strings.ToLower(root.Product.Name), strings.ToLower(prefix)) {
		*result = append(*result, root.Product)
	}

	Autocomplete(root.Left, prefix, result)
	Autocomplete(root.Right, prefix, result)
}
