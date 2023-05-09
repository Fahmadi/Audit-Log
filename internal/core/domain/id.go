package domain

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

type ID int64
type ExternalID string

var lock = &sync.Mutex{}
var instance *snowflake.Node

func NewID() (ID, error) {
	// TODO: singleton or injected snowflake struct
	// var err error
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		node, err := snowflake.NewNode(1)
		if err != nil {
			return 0, err
		}
		instance = node
	}

	id64 := instance.Generate().Int64()
	return ID(id64), nil
}

func (id ID) IsZero() bool {
	return id == 0
}
