package main

import (
	"fmt"
	"sync"

	"github.com/xiaonanln/goworld/engine/common"
	"github.com/xiaonanln/typeconv"
)

type ClientSpace struct {
	sync.Mutex

	owner *ClientBot
	ID    common.EntityID
	Kind  int

	Attrs     ClientAttrs
	destroyed bool
}

func newClientSpace(owner *ClientBot, entityid common.EntityID, data map[string]interface{}) *ClientSpace {
	space := &ClientSpace{
		owner: owner,
		ID:    entityid,
		Attrs: data,
	}
	space.Kind = int(typeconv.Int(data["_K"]))
	return space
}

func (space *ClientSpace) String() string {
	return fmt.Sprintf("ClientSpace<%d|%s>", space.Kind, space.ID)
}

func (space *ClientSpace) Destroy() {
	if space.destroyed {
		return
	}
	space.destroyed = true
}
