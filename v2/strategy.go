package v2

import (
	"golang.org/x/sync/errgroup"
	"gorm-perf-test/model"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Sequential(tx *gorm.DB, data []interface{}) error {
	for _, p := range data {
		pc := p
		tx.Create(pc)
	}

	return nil
}

func MultiThread(tx *gorm.DB, data []interface{}) error {
	ch := make(chan interface{}, len(data))
	for _, p := range data {
		p1 := p
		ch <- p1
	}
	close(ch)

	var wg errgroup.Group
	nThread := 20

	for i := 0; i < nThread; i++ {
		wg.Go(func() error {
			for p := range ch {
				p1 := p
				tx.Create(p1)
			}
			return nil
		})
	}

	err := wg.Wait()
	return err
}

func BatchProduct(tx *gorm.DB, data []interface{}) error {
	records := make([]*model.Product, len(data))
	for i, v := range data {
		records[i] = v.(*model.Product)
	}
	chunkSize := 100
	tx.CreateInBatches(records, chunkSize)
	return nil
}

func BatchPerson(tx *gorm.DB, data []interface{}) error {
	records := make([]*model.Person, len(data))
	for i, v := range data {
		records[i] = v.(*model.Person)
	}
	chunkSize := 10
	tx.CreateInBatches(records, chunkSize)
	return nil
}
