package product

import (
	"fmt"
	"strconv"
)


type Service struct {}

func NewService() *Service{

	return &Service{}
}

func (s *Service) List() []Product{
	return AllProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	elem, err := hasElemProduct(AllProducts, strconv.Itoa(idx))

	if err != nil {
		fmt.Println("get product not found elem")
		return &elem, err
	}

	return &elem, nil
}

func (s *Service) Delete(idx int) (*Product, error){


	elem, err := hasElemProduct(AllProducts, strconv.Itoa(idx))
	if err != nil {
		fmt.Printf("delete product not found: %s", err)
		return &Product{}, err
	}

	AllProducts = append(AllProducts[:idx], AllProducts[idx+1:]...)

	return &elem, nil
}

func (s *Service) New(name string) Product{

	AllProducts = append(AllProducts, Product{
		Title: name,
	})

	return AllProducts[len(AllProducts) - 1]
}