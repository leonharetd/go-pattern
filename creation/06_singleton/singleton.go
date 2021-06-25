package singleton

import (
	"sync"
)

// 饿汉式
var SingletonInstanceE *SingletonE

type SingletonE struct{}

func init() {
	SingletonInstanceE = &SingletonE{}
}

//懒汉式
var once sync.Once
var SingletonInstanceL *SingletonL

type SingletonL struct{}

type SingletonLManager struct {
	singletonExist *SingletonL
}

func GetSingletonLManager() *SingletonL {
	once.Do(func() {
		SingletonInstanceL = &SingletonL{}
	})
	return SingletonInstanceL
}
