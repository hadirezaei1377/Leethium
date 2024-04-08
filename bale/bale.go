package bale

import "fmt"

// Struct to represent user permissions
type Permissions struct {
	CanSeeMessages      bool
	CanDeleteMessages   bool
	CanEditMessages     bool
	CanKickMembers      bool
	CanMakeMembersAdmin bool
	CanAddMembers       bool
}

// Function to get user permissions based on an int8 input
func GetUserPermissions(permissionBits int8) Permissions {
	return Permissions{
		CanSeeMessages:      permissionBits&(1<<0) != 0,
		CanDeleteMessages:   permissionBits&(1<<1) != 0,
		CanEditMessages:     permissionBits&(1<<2) != 0,
		CanKickMembers:      permissionBits&(1<<3) != 0,
		CanMakeMembersAdmin: permissionBits&(1<<4) != 0,
		CanAddMembers:       permissionBits&(1<<5) != 0,
	}
}

// Function to set user permissions based on a Permissions struct input
func SetUserPermissions(permissions Permissions) int8 {
	var permissionBits int8
	if permissions.CanSeeMessages {
		permissionBits |= 1 << 0
	}
	if permissions.CanDeleteMessages {
		permissionBits |= 1 << 1
	}
	if permissions.CanEditMessages {
		permissionBits |= 1 << 2
	}
	if permissions.CanKickMembers {
		permissionBits |= 1 << 3
	}
	if permissions.CanMakeMembersAdmin {
		permissionBits |= 1 << 4
	}
	if permissions.CanAddMembers {
		permissionBits |= 1 << 5
	}
	return permissionBits
}

func main() {
	// Example usage
	userPermissions := GetUserPermissions(26) // Example input, 26 in binary is 11010
	fmt.Println(userPermissions)

	newPermissions := Permissions{
		CanSeeMessages:      true,
		CanDeleteMessages:   true,
		CanEditMessages:     false,
		CanKickMembers:      false,
		CanMakeMembersAdmin: true,
		CanAddMembers:       false,
	}
	newPermissionBits := SetUserPermissions(newPermissions)
	fmt.Println(newPermissionBits)
}
