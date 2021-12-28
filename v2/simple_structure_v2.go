package v2

import (
	"fmt"
	"gorm-perf-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"runtime"
	"time"
)

func TestInsertProducts(data []interface{}, method func(*gorm.DB, []interface{}) error) error {
	db, err := gorm.Open(mysql.Open("root:@/gorm_perf_test?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic("failed to connect to DB.")
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return err
	}
	db.Unscoped().Delete(model.Product{}, "1=1")

	start := time.Now()
	tx := db.Begin()

	err = method(tx, data)
	if err != nil {
		return err
	}

	tx.Commit()
	elapsed := time.Since(start)

	fmt.Printf("V2,Simple: Insertion `%d` object to DB using `%s` took: `%s`\n",
		len(data),
		runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(),
		elapsed)
	return nil
}
