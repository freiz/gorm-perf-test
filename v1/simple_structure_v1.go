package v1

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm-perf-test/model"
	"reflect"
	"runtime"
	"time"
)

func TestInsertProducts(data []interface{}, method func(*gorm.DB, []interface{}) error) error {
	db, err := gorm.Open("mysql", "root:@/gorm_perf_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to DB.")
	}
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)

	db.AutoMigrate(&model.Product{})
	db.Unscoped().Delete(model.Product{}, "1=1")

	start := time.Now()
	tx := db.Begin()

	err = method(tx, data)
	if err != nil {
		return err
	}

	tx.Commit()
	elapsed := time.Since(start)

	fmt.Printf("V1,Simple: Insertion `%d` object to DB using `%s` took: `%s`\n",
		len(data),
		runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(),
		elapsed)
	return nil
}
