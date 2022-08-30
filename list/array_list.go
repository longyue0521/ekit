// Copyright 2021 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

var (
	_ List[any] = &ArrayList[any]{}
)

// ArrayList 基于切片的简单封装
type ArrayList[T any] struct {
	vals []T
}

// NewArrayList 初始化一个len为0，cap为cap的ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{vals: make([]T, 0, cap)}
}

// NewArrayListOf 直接使用 ts，而不会执行复制
func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	if !a.ok(index) {
		return t, newErrIndexOutOfRange(a.Len(), index)
	}
	return a.vals[index], e
}

func (a *ArrayList[T]) ok(index int) bool {
	return 0 <= index && index < a.Len()
}

// Append 往ArrayList里追加数据
func (a *ArrayList[T]) Append(ts ...T) error {
	a.vals = append(a.vals, ts...)
	return nil
}

// Add 在ArrayList下标为index的位置插入一个元素
// 当index等于ArrayList长度等同于append
func (a *ArrayList[T]) Add(index int, t T) error {
	if index < 0 || index > len(a.vals) {
		return newErrIndexOutOfRange(a.Len(), index)
	}
	a.vals = append(a.vals[:index], append([]T{t}, a.vals[index:]...)...)
	return nil
}

// Set 设置ArrayList里index位置的值为t
func (a *ArrayList[T]) Set(index int, t T) error {
	if !a.ok(index) {
		return newErrIndexOutOfRange(a.Len(), index)
	}
	a.vals[index] = t
	return nil
}

func (a *ArrayList[T]) Delete(index int) (T, error) {
	if !a.ok(index) {
		var zero T
		return zero, newErrIndexOutOfRange(a.Len(), index)
	}
	res := a.vals[index]
	a.vals = append(a.vals[:index], a.vals[index+1:]...)
	return res, nil
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for index, value := range a.vals {
		e := fn(index, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	return append([]T{}, a.vals...)
}
