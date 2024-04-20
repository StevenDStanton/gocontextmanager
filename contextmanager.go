package gocontextmanager

import (
	"sync"
)

var MAX_MESSAGES_PER_CONTEXT = 25

type ContextManager struct {
	contexts sync.Map
}

func NewContextManager() *ContextManager {
	return &ContextManager{}
}

func (cm *ContextManager) AddContext(id string, message string) {
	value, ok := cm.contexts.Load(id)
	if !ok {
		cm.contexts.Store(id, []string{message})
		return
	}

	messages, _ := value.([]string)
	if len(messages) >= MAX_MESSAGES_PER_CONTEXT {
		messages = messages[1:]
	}

	messages = append(messages, message)
	cm.contexts.Store(id, messages)
}

func (cm *ContextManager) GetContext(id string) []string {
	value, ok := cm.contexts.Load(id)
	if !ok {
		return []string{}
	}
	return value.([]string)
}
