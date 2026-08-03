package utils

import (
	"github.com/mohae/deepcopy"
)

func MergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	if m1 == nil {
		m1 = make(map[string]interface{})
	}
	deepCopyM1 := deepcopy.Copy(m1).(map[string]interface{})
	for k, v := range m2 {
		deepCopyM1[k] = v
	}
	return deepCopyM1
}
