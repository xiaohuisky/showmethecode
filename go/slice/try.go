package slice

import "encoding/json"

func AppendInt(n int) []int {
	list := make([]int, 10, 10)
	list = append(list, n)
	return list
}

func CapacityExpansion(n int) int {
	list := make([]int, 0, 0)
	list = append(list, n)
	for i := 1;i <= n;i++ {
		list = append(list, 0)
	}
	return cap(list)
}

type S1 struct {
	Key rune `json:"key"`
}

type S2 struct {
	Key uint `json:"key"`
}

func JsonUnmarshal() interface{}{
	s1 := new(S1)
	s1.Key = -11
	sr1, _ := json.Marshal(s1)
	s2 := new(S2)
	err := json.Unmarshal(sr1, s2)
	if err != nil {
		return err
	}
	return s2
}

func MakeSliceInit(n int) int {
	return cap(make([]int, n))
}
