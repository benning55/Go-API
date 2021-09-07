package tronics

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProductsById(c echo.Context) error {
	var product map[int]string
	for _, p := range products {
		for k := range p {
			fmt.Println(c.Param("id"))
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := v.Struct(reqBody); err != nil {
		return err
	}

	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			fmt.Println(c.Param("id"))
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := v.Struct(reqBody); err != nil {
		return err
	}

	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		for k := range p {
			fmt.Println(c.Param("id"))
			if pID == k {
				product = p
				index = i
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	products = splice(products, index)
	return c.JSON(http.StatusOK, product)
}
