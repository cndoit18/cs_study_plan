package singleton

import "sync"

type Singleton struct {
	Value string
}

func GetInstanceFunc() func(string) *Singleton {
	var (
		s    *Singleton
		once sync.Once
	)
	return func(value string) *Singleton {
		once.Do(func() {
			s = &Singleton{Value: value}
		})
		return s
	}
}
