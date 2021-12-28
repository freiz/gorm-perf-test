package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm-perf-test/model"
	v1 "gorm-perf-test/v1"
	v2 "gorm-perf-test/v2"
)

func main() {
	var err error

	products, err := model.PrepareProducts(50000)
	if err != nil {
		panic(err)
	}
	persons, err := model.PreparePersons(2000)
	if err != nil {
		panic(err)
	}

	// Simple Data Model, V1
	err = v1.TestInsertProducts(products, v1.Sequential)
	if err != nil {
		panic(err)
	}
	err = v1.TestInsertProducts(products, v1.MultiThread)
	if err != nil {
		panic(err)
	}
	err = v1.TestInsertProducts(products, v1.Batch)
	if err != nil {
		panic(err)
	}
	println()

	// Simple Data Model, V2
	err = v2.TestInsertProducts(products, v2.Sequential)
	if err != nil {
		panic(err)
	}
	err = v2.TestInsertProducts(products, v2.MultiThread)
	if err != nil {
		panic(err)
	}
	err = v2.TestInsertProducts(products, v2.BatchProduct)
	if err != nil {
		panic(err)
	}
	println()

	// Complex Data Model, V1
	err = v1.TestInsertPersons(persons, v1.Sequential)
	if err != nil {
		panic(err)
	}
	// TODO not sure why this cannot run with other tests together.
	//err = v1.TestInsertPersons(persons, v1.MultiThread)
	//if err != nil {
	//	panic(err)
	//}
	println()

	// Complex Data Model, V2
	err = v2.TestInsertPersons(persons, v2.Sequential)
	if err != nil {
		panic(err)
	}
	err = v2.TestInsertPersons(persons, v2.MultiThread)
	if err != nil {
		panic(err)
	}
	err = v2.TestInsertPersons(persons, v2.BatchPerson)
	if err != nil {
		panic(err)
	}
}
