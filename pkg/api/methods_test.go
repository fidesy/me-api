package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	meclient = New(&Config{})
	ctx = context.Background()
)

func TestGetCollections(t *testing.T) {
	result, err := meclient.GetCollections(ctx, 0, 20)
	assert.Nil(t, err)
	assert.Equal(t, 20, len(result))
	assert.NotNil(t, result[0])
}

func TestGetListings(t *testing.T) {
	result, err := meclient.GetListings(ctx, "primates", 0, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(result))
	assert.NotNil(t, result[0])
}

func TestGetActivities(t *testing.T) {
	result, err := meclient.GetActivities(ctx, "primates", 0, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(result))
	assert.NotNil(t, result[0])
}

func TestGetStats(t *testing.T) {
	result, err := meclient.GetStats(ctx, "primates")
	assert.Nil(t, err)
	assert.Equal(t, result.Symbol, "primates")
	assert.NotEqual(t, 0, result.FloorPrice)
}

func TestGetLaunchpadCollections(t *testing.T) {
	result, err := meclient.GetLaunchpadCollections(ctx, 0, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(result))
	assert.NotNil(t, result[0])
}