package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`positive`, func(t *testing.T) {
		product := Products{
			Name:  "ravipon",
			Price: 100.50,
			SKU:   "ABC12345",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
