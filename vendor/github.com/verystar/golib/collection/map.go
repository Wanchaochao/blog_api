package collection

import "sort"

type StrMap map[string]string

func (this StrMap) Keys() []string {
	rs := make([]string, 0)
	for k, _ := range this {
		rs = append(rs, k)
	}
	return rs
}

func (this StrMap) ToInterface() map[string]interface{} {
	d := make(map[string]interface{})
	for k, v := range this {
		d[k] = v
	}
	return d
}

type StrMaps []map[string]string

// [] map[string]string  => map[string] map[string]string
func (this StrMaps) IndexBy(indexKey string, keys ... string) map[string]map[string]string {
	rs := make(map[string]map[string]string)
	for _, item := range this {
		if len(keys) == 0 {
			rs[item[indexKey]] = item
		} else {
			sub_map := make(map[string]string, 0)
			for _, key := range keys {
				sub_map[key] = item[key]
			}
			rs[item[indexKey]] = sub_map
		}

	}
	return rs
}

// [] map[string]string => map[string]string
func (this StrMaps) Pluck(indexKey string, valueKey string) map[string]string {
	rs := make(map[string]string)
	for _, item := range this {
		rs[item[indexKey]] = item[valueKey]
	}
	return rs
}

// map sort
type MapSorter struct {
	Keys []string
	Vals []string
}

func NewMapSorter(m map[string]string) *MapSorter {
	ms := &MapSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}
	for k, v := range m {
		ms.Keys = append(ms.Keys, k)
		ms.Vals = append(ms.Vals, v)
	}
	return ms
}

func (ms *MapSorter) Sort() {
	sort.Sort(ms)
}

func (ms *MapSorter) Len() int           { return len(ms.Keys) }
func (ms *MapSorter) Less(i, j int) bool { return ms.Keys[i] < ms.Keys[j] }
func (ms *MapSorter) Swap(i, j int) {
	ms.Vals[i], ms.Vals[j] = ms.Vals[j], ms.Vals[i]
	ms.Keys[i], ms.Keys[j] = ms.Keys[j], ms.Keys[i]
}
