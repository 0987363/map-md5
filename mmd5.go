package mmd5

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func MapMd5(m map[string]interface{}) string {
	st := toStruct(m)
	crc := md5.Sum(toJson(st))
	return fmt.Sprintf("%x", md5.Sum(crc[:]))
}

type Node struct {
	Key   string
	Value string
	Next  []*Node
}

func toJson(nodes []*Node) []byte {
	result, _ := json.Marshal(nodes)
	return result
}

func toStruct(m map[string]interface{}) []*Node {
	header := []*Node{}
	for k, v := range m {
		node := &Node{
			Key: k,
		}
		switch t := reflect.ValueOf(v); t.Kind() {
		case reflect.String:
			node.Value = v.(string)
		case reflect.Map:
			node.Next = toStruct(v.(map[string]interface{}))
			sortSlice(node.Next)
			//			fmt.Println("found map ", t)
		default:
			fmt.Println("unknown ", reflect.TypeOf(t))
		}
		header = append(header, node)
	}
	sortSlice(header)

	return header
}

type SortNode []*Node

func (a SortNode) Len() int      { return len(a) }
func (a SortNode) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortNode) Less(i, j int) bool {
	if strings.Compare(a[i].Key, a[j].Key) <= 0 {
		return false
	}
	return true
}

func sortSlice(nodes []*Node) {
	sort.Sort(SortNode(nodes))
}
