package test

import "fmt"

type Iterator interface {
	//判断是否有下一个值
	Next() bool

	//遍历键值
	Key() []byte
	Value() []byte
}

//键值的结构体
type Pair struct {
	Key   []byte
	Value []byte
}

//迭代器的结构体
type DefaultIterator struct {
	data   []Pair
	index  int
	length int
}

//创建默认迭代器
func NewDefalutIteratro(data map[string][]byte) *DefaultIterator {
	DefalutIter := new(DefaultIterator)
	DefalutIter.index = -1
	DefalutIter.length = len(data)
	for k, v := range data {
		p := Pair{
			Key:   []byte(k),
			Value: v,
		}
		//遍历出的数据添加到data
		DefalutIter.data = append(DefalutIter.data, p)
	}
	return DefalutIter
}

func (DefalutIter *DefaultIterator) Next() bool {
	if DefalutIter.index < DefalutIter.length-1 {
		DefalutIter.index++
		return true
	}
	return false
}

func (DefalutIter *DefaultIterator) Key() []byte {
	if DefalutIter.index == -1 || DefalutIter.index >= DefalutIter.length {
		panic(fmt.Errorf("IndexOutOfBoundError"))
	}
	return DefalutIter.data[DefalutIter.index].Key
}

func (DefalutIter *DefaultIterator) Value() []byte {
	if DefalutIter.index == -1 || DefalutIter.index >= DefalutIter.length {
		panic(fmt.Errorf("IndexOutOfBoundError"))
	}
	return DefalutIter.data[DefalutIter.index].Value
}
