package model

import (
	"github.com/bxcodec/faker/v3"
	"github.com/jinzhu/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Name        string        `faker:"name"`
	Email       string        `faker:"email"`
	PhoneNumber string        `faker:"phone_number"`
	CreditCards []*CreditCard `gorm:"ForeignKey:PersonID"`
}

type CreditCard struct {
	gorm.Model
	PersonID uint
	Number   string `faker:"cc_number"`
	Expire   time.Time
}

func PreparePersons(n int) ([]interface{}, error) {
	var ps []interface{}
	for i := 0; i < n; i++ {
		p := Person{}
		err := faker.FakeData(&p)
		if err != nil {
			return nil, err
		}

		p.ID = 0
		p.CreatedAt = time.Time{}
		p.UpdatedAt = time.Time{}
		p.DeletedAt = nil

		for i := range p.CreditCards {
			p.CreditCards[i].ID = 0
			p.CreditCards[i].CreatedAt = time.Time{}
			p.CreditCards[i].UpdatedAt = time.Time{}
			p.CreditCards[i].DeletedAt = nil
		}

		ps = append(ps, &p)
	}
	return ps, nil
}
