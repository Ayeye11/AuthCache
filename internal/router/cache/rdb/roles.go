package rdb

import (
	"context"
	"fmt"
	"time"

	"github.com/Ayeye11/AuthCache/internal/common/types"
	pb "github.com/Ayeye11/AuthCache/internal/router/cache/proto/gen"
	"google.golang.org/protobuf/proto"
)

func (c *cache) SaveRole(role *types.Role, perms []*types.Permission) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if role == nil || perms == nil {
		return fmt.Errorf("missing value")
	}

	permsPM := make([]*pb.Permission, 0, len(perms))
	for _, v := range perms {

		permsPM = append(permsPM, &pb.Permission{
			Category: v.Category,
			Action:   v.Action,
		})
	}

	rolePM := &pb.Role{
		ID:    int64(role.ID),
		Name:  role.Name,
		Perms: permsPM,
	}

	pb, err := proto.Marshal(rolePM)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, fmt.Sprintf("role:%d", rolePM.ID), pb, c.ttl).Err()
}

func (c *cache) GetRole(roleID int) (*types.Role, []*types.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := c.client.Get(ctx, fmt.Sprintf("role:%d", roleID)).Result()
	if err != nil {
		return nil, nil, err
	}

	rolePM := &pb.Role{}
	if err := proto.Unmarshal([]byte(res), rolePM); err != nil {
		return nil, nil, err
	}

	perms := make([]*types.Permission, 0, len(rolePM.Perms))
	for _, v := range rolePM.Perms {
		perms = append(perms, &types.Permission{
			Category: v.Category,
			Action:   v.Action,
		})
	}

	return &types.Role{ID: uint(rolePM.ID), Name: rolePM.Name}, perms, nil
}
