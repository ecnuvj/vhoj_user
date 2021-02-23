package adapter

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
	"github.com/jinzhu/gorm"
)

func ModelRolesToRpcRoles(roles []*model.Role) []*userpb.Role {
	retRoles := make([]*userpb.Role, len(roles))
	for i, r := range roles {
		retRoles[i] = &userpb.Role{
			RoleId:   uint64(r.ID),
			RoleName: r.RoleName,
		}
	}
	return retRoles
}

func RpcRolesToModelRoles(roles []*userpb.Role) []*model.Role {
	retRoles := make([]*model.Role, len(roles))
	for i, r := range roles {
		retRoles[i] = &model.Role{
			Model:    gorm.Model{ID: uint(r.RoleId)},
			RoleName: r.RoleName,
		}
	}
	return retRoles
}
