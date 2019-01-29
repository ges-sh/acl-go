package acl

// Role represents single role in system
type Role map[int]uint64

// AddPerms adds perms to object to role.
func (r Role) AddPerms(obj int, perms ...uint64) {
	for _, perm := range perms {
		r[obj] |= perm
	}
}

// RevokePerms removes perms to object from role.
func (r Role) RevokePerms(obj int, perms ...uint64) {
	for _, perm := range perms {
		r[obj] &^= perm
	}
}

// Can specifies whether role have permission on object
func (r Role) Can(obj int, perm uint64) bool {
	return r[obj]&perm != 0
}
