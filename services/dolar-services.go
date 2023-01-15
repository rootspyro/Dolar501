package services

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type DolarServices struct {
	rdClient *redis.Client
}

func NewDolarServices( client *redis.Client ) *DolarServices {
	return &DolarServices{
		rdClient: client,
	}
}

func( impl *DolarServices )GetPlatformsList() []string {
	keys := impl.rdClient.Keys(context.Background(), "*")	
	return keys.Val()
}
