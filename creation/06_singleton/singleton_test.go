package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	se1 := SingletonInstanceE
	se2 := SingletonInstanceE
	if se1 != se2 {
		t.Error("Singleton instance E failed")
	}

	sl1 := GetSingletonLManager()
	sl2 := GetSingletonLManager()

	if sl1 != sl2 {
		t.Error("Singleton instace L failed")
	}
}

func TestParallelWorkerSingleton(t *testing.T) {
	var worker int = 500
	var wg sync.WaitGroup
	m := [500]*SingletonL{}
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func(x int) {
			ist := GetSingletonLManager()
			m[x] = ist
			wg.Done()
		}(i)
	}
	wg.Wait()
    for i := 1; i < worker; i++ {
		if m[i] != m[i-1] {
			t.Error("Singleton parallel failed")
		}
	}
}
