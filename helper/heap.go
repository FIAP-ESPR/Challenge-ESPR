package helper

import (
	"ancora/model"
	"container/heap"
)

func GetKMinProducts(products []model.Product, k int) []model.Product {
	h := &model.Heap{}
	heap.Init(h)
	for _, p := range products {
		heap.Push(h, p)
	}
	result := []model.Product{}
	for i := 0; i < k && h.Len() > 0; i++ {
		result = append(result, heap.Pop(h).(model.Product))
	}
	return result
}

func GetKMaxProducts(products []model.Product, k int) []model.Product {
	h := &model.Heap{}
	heap.Init(h)
	for _, p := range products {
		heap.Push(h, p)
	}
	result := []model.Product{}
	for i := 0; i < len(products)-k && h.Len() > 0; i++ {
		heap.Pop(h)
	}
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(model.Product))
	}
	return result
}
