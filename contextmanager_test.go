package gocontextmanager

import (
	"reflect"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContextManagerTypes(t *testing.T) {
	cm := NewContextManager()
	assert.NotNil(t, cm, "ContextManager instance should not be nil")
	assert.Equal(t, reflect.TypeOf(&ContextManager{}), reflect.TypeOf(cm), "cm should be of type *ContextManager")
	assert.IsType(t, sync.Map{}, cm.contexts, "cm.contexts should be of type sync.Map")
}

func TestAddFewerThan25Items(t *testing.T) {
	cm := NewContextManager()
	userRole := "user"
	userID := "123"
	globalName := "GlobalUser"
	name := "John Doe"
	for i := 1; i < 24; i++ {
		cm.AddContext("testID", "Message "+strconv.Itoa(i), userRole, userID, globalName, name)
	}
	got := cm.GetContext("testID")
	assert.Equal(t, 23, len(got), "Expected 23 messages")
	assert.Equal(t, "Message 1", got[0].content, "First message should be 'Message 1'")
	assert.Equal(t, "Message 10", got[9].content, "Tenth message should be 'Message 10'")
}

func TestAddExactly25Items(t *testing.T) {
	cm := NewContextManager()
	userRole := "user"
	userID := "123"
	globalName := "GlobalUser"
	name := "John Doe"
	for i := 1; i <= 25; i++ {
		cm.AddContext("testID", "Message "+strconv.Itoa(i), userRole, userID, globalName, name)
	}
	got := cm.GetContext("testID")
	assert.Equal(t, 25, len(got), "Expected 25 messages")
	assert.Equal(t, "Message 1", got[0].content, "First message should be 'Message 1'")
	assert.Equal(t, "Message 10", got[9].content, "Tenth message should be 'Message 10'")
	assert.Equal(t, "Message 25", got[24].content, "Twenty-fifth message should be 'Message 25'")
}

func TestAddMoreThan25Items(t *testing.T) {
	cm := NewContextManager()
	userRole := "user"
	userID := "123"
	globalName := "GlobalUser"
	name := "John Doe"
	for i := 1; i <= 26; i++ {
		cm.AddContext("testID", "Message "+strconv.Itoa(i), userRole, userID, globalName, name)
	}
	got := cm.GetContext("testID")
	assert.Equal(t, 25, len(got), "Expected 25 messages after adding 26 items")
	assert.Equal(t, "Message 2", got[0].content, "First message should be 'Message 2' after adding 26 items")
	assert.Equal(t, "Message 11", got[9].content, "Tenth message should be 'Message 11'")
	assert.Equal(t, "Message 26", got[24].content, "Twenty-sixth message should be 'Message 26'")
}
