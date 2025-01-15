package entity

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name  string  `valid:required~Name is required`
	Price float64 `valid:required~Price is required, range(1|1000)~Price must be between 1 and 1000`
	SKU   string  `valid:required~SKU is required, matches("^[A-Z]{3}//d{5}$")~SKU must start with 3 uppercase English letters (A-Z) followed by 5 digits (0-9)`
}
