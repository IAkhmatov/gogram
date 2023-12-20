// Copyright (c) 2023 RoseLoverX

package utils

import (
	"reflect"
	"sync"

	"github.com/IAkhmatov/gogram/internal/encoding/tl"
)

type SyncSetInt struct {
	mutex sync.RWMutex
	m     map[int]struct{}
}

func NewSyncSetInt() *SyncSetInt {
	return &SyncSetInt{m: make(map[int]struct{})}
}

func (s *SyncSetInt) Has(key int) bool {
	s.mutex.RLock()
	_, ok := s.m[key]
	s.mutex.RUnlock()
	return ok
}

func (s *SyncSetInt) Add(key int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.m[key]
	s.m[key] = struct{}{}
	return !ok
}

func (s *SyncSetInt) Delete(key int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.m[key]
	delete(s.m, key)
	return ok
}

func (s *SyncSetInt) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m = make(map[int]struct{})
}

type SyncIntObjectChan struct {
	mutex sync.RWMutex
	m     map[int]chan tl.Object
}

func (s *SyncIntObjectChan) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m = make(map[int]chan tl.Object)
}

func NewSyncIntObjectChan() *SyncIntObjectChan {
	return &SyncIntObjectChan{m: make(map[int]chan tl.Object)}
}

func (s *SyncIntObjectChan) Has(key int) bool {
	s.mutex.RLock()
	_, ok := s.m[key]
	s.mutex.RUnlock()
	return ok
}

func (s *SyncIntObjectChan) Get(key int) (chan tl.Object, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func (s *SyncIntObjectChan) Add(key int, value chan tl.Object) {
	s.mutex.Lock()
	s.m[key] = value
	s.mutex.Unlock()
}

func (s *SyncIntObjectChan) Keys() []int {
	keys := make([]int, 0, len(s.m))
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *SyncIntObjectChan) Delete(key int) bool {
	s.mutex.Lock()
	_, ok := s.m[key]
	delete(s.m, key)
	s.mutex.Unlock()
	return ok
}

type SyncIntReflectTypes struct {
	mutex sync.RWMutex
	m     map[int][]reflect.Type
}

func (s *SyncIntReflectTypes) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m = make(map[int][]reflect.Type)
}

func NewSyncIntReflectTypes() *SyncIntReflectTypes {
	return &SyncIntReflectTypes{m: make(map[int][]reflect.Type)}
}

func (s *SyncIntReflectTypes) Has(key int) bool {
	s.mutex.RLock()
	_, ok := s.m[key]
	s.mutex.RUnlock()
	return ok
}

func (s *SyncIntReflectTypes) Get(key int) ([]reflect.Type, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func (s *SyncIntReflectTypes) Add(key int, value []reflect.Type) {
	s.mutex.Lock()
	s.m[key] = value
	s.mutex.Unlock()
}

func (s *SyncIntReflectTypes) Keys() []int {
	keys := make([]int, 0, len(s.m))
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *SyncIntReflectTypes) Delete(key int) bool {
	s.mutex.Lock()
	_, ok := s.m[key]
	delete(s.m, key)
	s.mutex.Unlock()
	return ok
}

func (s *SyncIntObjectChan) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for k, v := range s.m {
		CloseChannelWithoutPanic(v)
		delete(s.m, k)
	}
}

func CloseChannelWithoutPanic(c chan tl.Object) {
	defer func() {
		if r := recover(); r != nil {
			close(c)
		}
	}()
}
