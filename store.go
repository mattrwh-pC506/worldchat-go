package main

import (
	"sync"
)

var (
	storeLock sync.RWMutex
)

type Store struct {
	messages []Message
	size     int
}

var store *Store

func getStore() *Store {
	if store == nil {
		storeLock.Lock()
		defer storeLock.Unlock()
		store = &Store{}
	}
	return store
}

func (store *Store) addMessage(message Message) *Store {
	storeLock.Lock()
	storeInstance := getStore()
	storeInstance.size += 1
	defer storeLock.Unlock()
	storeInstance.messages = append(storeInstance.messages, message)
	return storeInstance
}

func (store *Store) getMessageByIndex(index int) Message {
	storeLock.RLock()
	defer storeLock.RUnlock()
	return getStore().messages[index]
}
