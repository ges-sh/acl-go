package acl_test

import (
	"fmt"
	"testing"

	acl "github.com/ges-sh/acl-go/v2"
)

const (
	User = iota
	Admin
)

// Permissions
const (
	Create = 1 << iota
	Delete
	// ...
)

const (
	Article = iota
	Video
)

func TestACL(t *testing.T) {
	acl := acl.ACL{}

	user := acl.AddRole(User)
	user.AddPerms(Article, Create)

	fmt.Println(acl.Can(User, Article, Create)) // true
}
