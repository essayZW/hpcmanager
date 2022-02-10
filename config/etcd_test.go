package config

import (
	"strconv"
	"sync"
	"testing"

	"github.com/essayZW/hpcmanager"
)

func TestEtcdDynamicConfig(t *testing.T) {
	hpcmanager.LoadCommonArgs()

	wt := new(sync.WaitGroup)
	var example3 float64 = 10086
	var example4 bool
	var example5 string = "initTest"

	examples := []struct {
		Value   interface{}
		Path    string
		Handler ValueChange
		Error   bool
	}{
		{
			// test nil ptr
			Value:   nil,
			Path:    "test1",
			Handler: nil,
			Error:   true,
		},
		{
			// test non-ptr
			Value:   1,
			Path:    "test2",
			Handler: nil,
			Error:   true,
		},
		{
			// test number data, type float64
			Value: &example3,
			Path:  "test3",
			Handler: func(newV interface{}) {
				t.Logf("test3 change %v", newV)
				wt.Done()
			},
			Error: false,
		},
		{
			// test bool data, type bool
			Value: &example4,
			Path:  "test4",
			Handler: func(newV interface{}) {
				t.Logf("test4 change %v", newV)
				wt.Done()
			},
			Error: false,
		},
		{
			// test string data, type string
			Value: &example5,
			Path:  "test5",
			Handler: func(newV interface{}) {
				t.Logf("test5 change %v", newV)
				wt.Done()
			},
			Error: false,
		},
	}
	etcdConfig, err := NewEtcd()
	if err != nil {
		t.Fatal(err)
	}

	for index, example := range examples {
		wt.Add(1)
		t.Run("TestEtcdDynamicConfig"+strconv.Itoa(index), func(t *testing.T) {
			err := etcdConfig.Registry(example.Path, example.Value, example.Handler)
			if err != nil {
				if !example.Error {
					t.Error(err)
				}
				wt.Done()
			}
		})

	}
	wt.Wait()
}
