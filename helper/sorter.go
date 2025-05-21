package helper

import "ancora/model"

func MergeSort(products []model.Product, ascending bool) []model.Product {
	if len(products) <= 1 {
		return products
	}

	mid := len(products) / 2
	left := MergeSort(products[:mid], ascending)
	right := MergeSort(products[mid:], ascending)

	return merge(left, right, ascending)
}

func merge(left, right []model.Product, ascending bool) []model.Product {
	result := []model.Product{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if (ascending && left[i].Price < right[j].Price) || (!ascending && left[i].Price > right[j].Price) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
