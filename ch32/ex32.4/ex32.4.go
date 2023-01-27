package main

import "fmt"

type FlyWeightFactory struct {
	pool     []*FlyWeight
	AllocCnt int
}

func (factory *FlyWeightFactory) Create() *FlyWeight {
	var flyWeight *FlyWeight
	fmt.Printf("Pool Size => %v\n", len(factory.pool))
	if len(factory.pool) > 0 {
		flyWeight, factory.pool = factory.pool[len(factory.pool)-1], factory.pool[:len(factory.pool)-1]
		flyWeight.Reuse()
	} else {
		flyWeight = &FlyWeight{}
		factory.AllocCnt++
	}
	return flyWeight
}

func (factory *FlyWeightFactory) Dispose(flyWeight *FlyWeight) {
	flyWeight.Dispose()
	factory.pool = append(factory.pool, flyWeight)
}

func NewFlyWeightFactory(initSize int) *FlyWeightFactory {
	return &FlyWeightFactory{pool: make([]*FlyWeight, 0, initSize)}
}

type FlyWeight struct {
	SomeData   string
	isDisposed bool
}

func (f *FlyWeight) Reuse() {
	f.isDisposed = false
}

func (f *FlyWeight) Dispose() {
	f.isDisposed = true
}

func (f *FlyWeight) IsDisposed() bool {
	return f.isDisposed
}

func main() {
	factory := NewFlyWeightFactory(1000)
	flyWeights := make([]*FlyWeight, 0)
	for i := 0; i < 1000; i++ {
		flyWeight := factory.Create()
		flyWeight.SomeData = "SomeData"
		flyWeights = append(flyWeights, flyWeight)
	}
	fmt.Println(factory.pool)
}
