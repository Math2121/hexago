package cli

import (
	"fmt"

	"github.com/Math2121/hexago/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error){
	var result string = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price);
		if err!= nil {
            return result, err
        }
		result = fmt.Sprintf("Product ID #%s with the name #%s", product.GetID(), product.GetName())
	case "enable":
		product, err := service.Get(productId);
        if err!= nil {
            return result, err
        }
        product, err = service.Enable(product);
        if err!= nil {
            return result, err
        }
        result = fmt.Sprintf("Product ID #%s with the name #%s and status %s has been enabled", product.GetID(), product.GetName(), product.GetStatus())
	case "disable":
		product, err := service.Get(productId);
        if err!= nil {
            return result, err
        }
        product, err = service.Disable(product);
        if err!= nil {
            return result, err
        }
        result = fmt.Sprintf("Product ID #%s with the name #%s and status %s has been disabled", product.GetID(), product.GetName(), product.GetStatus())
	default:
		product, err := service.Get(productId);
        if err!= nil {
            return result, err
        }
		result = fmt.Sprintf("Product ID #%s with the name #%s and status %s", product.GetID(), product.GetName(), product.GetStatus())

	}

	return result, nil

}