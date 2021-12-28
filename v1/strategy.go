package v1

import (
	"github.com/jinzhu/gorm"
	gormBulk "github.com/t-tiger/gorm-bulk-insert"
	"golang.org/x/sync/errgroup"
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

func Batch(tx *gorm.DB, data []interface{}) error {
	chunkSize := 2000
	// This does not work for complex data model, data apart from main table will be discarded.
	err := gormBulk.BulkInsert(tx, data, chunkSize)
	return err
}
