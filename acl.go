package acl

type ACL interface {
	Can(int, int, uint64) bool
	Perms(int, int) uint64
}

type Acl map[int]map[int]uint64

func (a Acl) AddPerms(r int, o int, p ...uint64) {
	for i := range p {
		a[r][o] |= p[i]
	}
}

func (a Acl) RevokePerms(r int, o int, p ...uint64) {
	for i := range p {
		a[r][o] &^= p[i]
	}
}

func (a Acl) AddRole(r int, inh ...int) {
	a[r] = make(map[int]uint64)
	for i := range inh {
		for j, v := range a[inh[i]] {
			a[r][j] |= v
		}
	}
}

func (a Acl) Can(r int, o int, p uint64) bool {
	return a[r][o]&p != 0
}

func (a Acl) Perms(r int, o int) uint64 {
	return a[r][o]
}
