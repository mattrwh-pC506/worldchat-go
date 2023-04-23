package main

import (
	"github.com/google/uuid"
	"log"
	"sync"
)

var (
	storeLock sync.RWMutex
)

type Store struct {
	messages   []Message
	messageIds map[uuid.UUID]bool
}

var store *Store

func getStore() *Store {
	if store == nil {
		storeLock.Lock()
		defer storeLock.Unlock()
		store = &Store{messageIds: map[uuid.UUID]bool{}, messages: []Message{}}
	}
	return store
}

func (store *Store) addMessage(message Message) *Store {
	storeLock.Lock()
	defer storeLock.Unlock()

	storeInstance := getStore()

	log.Println(storeInstance.messageIds)
	if _, exists := storeInstance.messageIds[message.Id]; !exists {
		storeInstance.messages = append(storeInstance.messages, message)
		storeInstance.messageIds[message.Id] = true
	}

	return storeInstance
}

func (store *Store) getMessageByIndex(index int) Message {
	storeLock.RLock()
	defer storeLock.RUnlock()
	return getStore().messages[index]
}

func (store *Store) getSize() int {
	storeLock.RLock()
	defer storeLock.RUnlock()
	return len(getStore().messages)
}
