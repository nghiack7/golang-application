package xcontext

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

const (
	KeyContextID ContextKey = "context_id"
	KeyUserID    ContextKey = "user_id"
)

var AllKeys = []ContextKey{
	KeyContextID,
	KeyUserID,
}

func AttachContext(c context.Context, key ContextKey, value string) context.Context {
	return context.WithValue(c, key, value)
}

// nolint
func AttachGinContext(c *gin.Context, key ContextKey, value string) *gin.Context {
	c.Set(key.String(), value)
	return c
}
