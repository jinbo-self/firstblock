package test

import "fmt"

//结构体
type DB struct {
	path string
	data map[string][]byte
}

//模拟连接
func New(path string) (*DB, error) {
	DBObject := DB{
		path: path,
		data: make(map[string][]byte),
	}
	return &DBObject, nil
}

//关闭连接
func (DBObject *DB) Close() error {
	return nil
}

//put
func (DBObject *DB) Put(key []byte, value []byte) error {
	DBObject.data[string(key)] = value
	return nil
}

//get
func (DBObject *DB) Get(key []byte) ([]byte, error) {
	if v, ok := DBObject.data[string(key)]; ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("NotFound")
	}
}

//Del
func (DBObject *DB) Del(key []byte) error {
	if _, ok := DBObject.data[string(key)]; ok {
		delete(DBObject.data, string(key))
		return nil
	} else {
		return fmt.Errorf("NotFound")
	}
}

//模拟遍历
func (DBObject *DB) Iterator() Iterator {
	return NewDefalutIteratro(DBObject.data)
}
