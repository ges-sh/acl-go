package acl_test

import (
	"testing"

	acl "github.com/ges-sh/acl-go"
)

const (
	Guest = 0
	User  = 1
	Admin = 2
)

const (
	Read uint64 = 1 << iota
	Write
)

const (
	Article int = iota
	Video
)

func TestACL(t *testing.T) {
	acl := acl.ACL{}

	acl.AddRole(Guest)
	acl.AddPerms(Guest, Article, Read)
	acl.AddPerms(Guest, Video, Read)

	acl.AddRole(User, Guest)
	acl.AddPerms(User, Video, Write)

	acl.AddRole(Admin, User)
	acl.AddPerms(Admin, Article, Write)

	if acl.Can(Guest, Article, Write) {
		t.Errorf("Guest shouldn't be able to write articles")
		return
	}
}
