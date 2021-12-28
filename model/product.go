package model

import (
	"github.com/bxcodec/faker/v3"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code        string
	Price       uint
	Description string
}

func PrepareProducts(n int) ([]interface{}, error) {
	var ps []interface{}
	for i := 0; i < n; i++ {
		p := Product{}
		err := faker.FakeData(&p)
		if err != nil {
			return nil, err
		}

		p.ID = 0
		p.CreatedAt = time.Time{}
		p.UpdatedAt = time.Time{}
		p.DeletedAt = nil
		ps = append(ps, &p)
	}
	return ps, nil
}
