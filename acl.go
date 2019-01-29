package acl

// ACL contains all role's permissions. It is possible to have up to 64 permissions on a single object.
type ACL map[int]map[int]uint64

// AddPerms adds permissions p on object o for role r
func (a ACL) AddPerms(r int, o int, p ...uint64) {
	for i := range p {
		a[r][o] |= p[i]
	}
}

// RevokePerms removes permissions p on object o for role r
func (a ACL) RevokePerms(r int, o int, p ...uint64) {
	for i := range p {
		a[r][o] &^= p[i]
	}
}

// AddRole add new role r which inherits permissions from roles within inh.
func (a ACL) AddRole(r int, inh ...int) {
	a[r] = make(map[int]uint64)
	for i := range inh {
		for j, v := range a[inh[i]] {
			a[r][j] |= v
		}
	}
}

// Can specifies if user of role r has permission p on object o
func (a ACL) Can(r int, o int, p uint64) bool {
	return a[r][o]&p != 0
}
