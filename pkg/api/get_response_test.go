package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/fidesy/me-api/pkg/models"
)

var (
	client        = New(&Config{})
	collectionUrl = "https://api-mainnet.magiceden.dev/v2/collections?offset=0&limit=20"
)

func TestGetResponse(t *testing.T) {
	var collections []*models.Collection
	err := client.GetResponse(context.Background(), collectionUrl, &collections)
	assert.Nil(t, err)
	assert.NotNil(t, collections)
}
