// Package acl provides simple RBACL (Role Based Access Control List).
// It allows to create list of roles with specific permissions on given resources, protecting them from unwanted access.
package acl

// ACL contains all role's permissions. It is possible to have up to 64 permissions on a single object.
type ACL map[int]Role

// AddRole add new role r which inherits permissions from roles within inh.
func (a ACL) AddRole(r int, inh ...Role) Role {
	a[r] = make(Role)
	for _, inhRole := range inh {
		for j, v := range inhRole {
			a[r][j] |= v
		}
	}
	return a[r]
}

// Can specifies whether role have permission on object
func (a ACL) Can(role, obj int, perm uint64) bool {
	return a[role].Can(obj, perm)
}
