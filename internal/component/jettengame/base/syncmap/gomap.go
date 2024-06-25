package syncmap

import "sync"

type GoMap struct {
	key interface{}
	val interface{}
	sync.Map
}

func (m GoMap) Make(key, value interface{}) {
	m.key = key
	m.val = value
}
