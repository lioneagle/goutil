package container

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Scope[T any] struct {
	name      string
	fullPath  string
	data      T
	subScopes *OrderedMap[string, *Scope[T]]
}

func NewScope[T any](name string) *Scope[T] {
	return &Scope[T]{
		name:      name,
		fullPath:  name,
		subScopes: NewOrderedMap[string, *Scope[T]](),
	}
}

func (this *Scope[T]) GetFullPath() string {
	return this.fullPath
}

func (this *Scope[T]) GetData() T {
	return this.data
}

func (this *Scope[T]) SetData(data T) {
	this.data = data
}

func (this *Scope[T]) FindScopeByFullPath(fullPath string) *Scope[T] {
	_, path := GetPathAndNameByFullPath(fullPath)
	return this.FindScope(path...)
}

func (this *Scope[T]) FindScope(path ...string) *Scope[T] {
	scope := this

	for _, key := range path {
		var ok bool
		scope, ok = scope.subScopes.Find(key)
		if !ok {
			return nil
		}
	}
	return scope
}

func (this *Scope[T]) AddScopeByFullPath(fullPath string) *Scope[T] {
	_, path := GetPathAndNameByFullPath(fullPath)
	return this.AddScope(path...)
}

func (this *Scope[T]) AddScope(path ...string) *Scope[T] {
	scope := this
	fullPath := ""

	i := 0
	for ; i < len(path); i++ {
		fullPath = scope.fullPath
		subScope := scope.FindScope(path[i])
		if subScope == nil {
			break
		}
		scope = subScope
	}

	for ; i < len(path); i++ {
		fullPath = fmt.Sprintf("%s%s/", fullPath, path[i])
		scope = scope.addOneScope(path[i], fullPath)
	}

	return scope
}

func (this *Scope[T]) addOneScope(path, fullPath string) *Scope[T] {
	scope := NewScope[T](path)
	scope.fullPath = fullPath
	this.subScopes.Add(path, scope)
	return scope
}

func (this *Scope[T]) AddData(data T, path ...string) error {
	scope := this.FindScope(path...)
	if scope == nil {
		return errors.Errorf("Scope %s: add scope failed, path = %v", this.name, path)
	}
	scope.data = data
	return nil
}

func (this *Scope[T]) FindData(path ...string) *T {
	keyNum := len(path)
	if keyNum <= 0 {
		return nil
	}

	scope := this.FindScope(path[:keyNum-1]...)
	if scope == nil {
		return nil
	}
	return &scope.data
}

func (this *Scope[T]) ForEach(op func(fullPath string, data *T) (halt bool, err error)) error {
	if op == nil {
		return nil
	}

	halt, err := op(this.fullPath, &this.data)
	if err != nil {
		return err
	}

	if halt {
		return nil
	}

	err = this.subScopes.ForEach(func(key string, val *Scope[T]) (halt bool, err error) {
		return false, val.ForEach(op)
	})

	return err
}

func GetPathAndNameByFullPath(fullPath string) (name string, path []string) {
	path = strings.Split(fullPath, "/")
	if len(path) <= 1 || path[0] != "" {
		return "", nil
	}

	name = path[len(path)-1]
	path = path[1 : len(path)-1]
	return name, path
}
