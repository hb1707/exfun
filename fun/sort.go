package fun

import "sort"

type MapKV struct {
	Key   string
	Value int
}
type MapKVList []MapKV

func (p MapKVList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p MapKVList) Len() int           { return len(p) }
func (p MapKVList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// SortMapByValue todo 有问题
// SortMapByValue 对map按照value进行排序
func SortMapByValue(m map[string]int) MapKVList {
	p := make(MapKVList, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapKV{k, v}
	}
	sort.Sort(p)
	return p
}
