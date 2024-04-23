package gocontextmanager

import (
	"sync"
)

const maxMessagesPerContext = 25

type Message struct {
	content        string
	role           string
	userID         string
	userGlobalName string
	userName       string
}

type context struct {
	messages []Message
}

type ContextManager struct {
	contexts sync.Map
}

func NewContextManager() *ContextManager {
	return &ContextManager{}
}

func (cm *ContextManager) AddContext(id string, content string, userRole string, userID string, globalName string, name string) {
	value, ok := cm.contexts.Load(id)
	if !ok {
		cm.contexts.Store(id, context{messages: []Message{{
			content:        content,
			role:           userRole,
			userID:         userID,
			userGlobalName: globalName,
			userName:       name,
		}}})
		return
	}

	existingContext, _ := value.(context)
	if len(existingContext.messages) >= maxMessagesPerContext {
		existingContext.messages = existingContext.messages[1:]
	}
	existingContext.messages = append(existingContext.messages, Message{
		content:        content,
		role:           userRole,
		userID:         userID,
		userGlobalName: globalName,
		userName:       name,
	})
	cm.contexts.Store(id, existingContext)
}

func (cm *ContextManager) GetContext(id string) []Message {
	value, ok := cm.contexts.Load(id)
	if !ok {
		return nil
	}
	return value.(context).messages
}
