package internal

import (
	"events/pkg/ecommerce"
	"log"
)

type EcommerceOrderRepository struct {
}

func (r *EcommerceOrderRepository) Save(o *ecommerce.Order) {
	log.Println("Saving order")
}
