package acl

type ACL interface {
	Can(int, int, byte) bool
}

type Acl map[int]map[int]byte

func (a Acl) AddPerms(r int, o int, p ...byte) {
	for i := range p {
		a[r][o] |= p[i]
	}
}

func (a Acl) RevokePerms(r int, o int, p ...byte) {
	for i := range p {
		a[r][o] &^= p[i]
	}
}

func (a Acl) AddRole(r int, inh ...int) {
	a[r] = make(map[int]byte)
	for i := range inh {
		for j, v := range a[inh[i]] {
			a[r][j] |= v
		}
	}
}

func (a Acl) Can(r int, o int, p byte) bool {
	return a[r][o]&p != 0
}
