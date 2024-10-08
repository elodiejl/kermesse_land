package config

import "back/models"

// 0 = organizer, 2 = admin, 4 = student , 6 = parents, 8 = stand_leader
const (
	RoleOrganizer   uint8 = 1 << iota
	RoleAdmin       uint8 = 1 << 1
	RoleStudent     uint8 = 1 << 2
	RoleParent      uint8 = 1 << 3
	RoleStandLeader uint8 = 1 << 4
)

func HasRole(user *models.User, role uint8) bool {
	return user.Roles&role != 0
}

func HasRequiredRole(userRoles, requiredRole uint8) bool {

	if requiredRole == RoleOrganizer {

		return userRoles&RoleOrganizer != 0 || userRoles&RoleAdmin != 0
	}
	if requiredRole == RoleParent {
		return userRoles&RoleParent != 0 || userRoles&RoleAdmin != 0
	}
	if requiredRole == RoleStudent {
		return userRoles&RoleStudent != 0 || userRoles&RoleAdmin != 0
	}
	if requiredRole == RoleStandLeader {
		return userRoles&RoleStandLeader != 0 || userRoles&RoleAdmin != 0
	}
	return userRoles&requiredRole != 0
}
