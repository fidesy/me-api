package api

import (
	"context"
	"fmt"

	models "github.com/fidesy/me-api/pkg/models"
)

var (
	baseURL        = "https://api-mainnet.magiceden.dev/v2"
	collectionsURL = baseURL + "/collections?offset=%d&limit=%d"
	listingsURL    = baseURL + "/collections/%s/listings?offset=%d&limit=%d"
	activitiesURL  = baseURL + "/collections/%s/activities?offset=%d&limit=%d"
	statsURL       = baseURL + "/collections/%s/stats"
	launchpadURL   = baseURL + "/launchpad/collections?offset=%d&limit=%d"
)

func (c *Client) GetCollections(ctx context.Context, offset, limit int) ([]*models.Collection, error) {
	var collections []*models.Collection
	err := c.GetResponse(ctx, fmt.Sprintf(collectionsURL, offset, limit), &collections)
	return collections, err
}

func (c *Client) GetListings(ctx context.Context, symbol string, offset, limit int) ([]*models.Listing, error) {
	var listings []*models.Listing
	err := c.GetResponse(ctx, fmt.Sprintf(listingsURL, symbol, offset, limit), &listings)
	return listings, err
}

func (c *Client) GetActivities(ctx context.Context, symbol string, offset, limit int) ([]*models.Activity, error) {
	var activities []*models.Activity
	err := c.GetResponse(ctx, fmt.Sprintf(activitiesURL, symbol, offset, limit), &activities)
	return activities, err
}

func (c *Client) GetStats(ctx context.Context, symbol string) (*models.Stats, error) {
	var stats *models.Stats
	err := c.GetResponse(ctx, fmt.Sprintf(statsURL, symbol), &stats)
	return stats, err
}

func (c *Client) GetLaunchpadCollections(ctx context.Context, offset, limit int) ([]*models.LaunchpadCollection, error) {
	var launchpadCollections []*models.LaunchpadCollection
	err := c.GetResponse(ctx, fmt.Sprintf(launchpadURL, offset, limit), &launchpadCollections)
	return launchpadCollections, err
}
