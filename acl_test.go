package acl

import "testing"

// User roles
const (
	Guest = 0
	User  = 1
	Admin = 2
)

// for test purposes
var roleLookup = map[int]string{
	Guest: "guest",
	User:  "user",
	Admin: "admin",
}

// User permissions
const (
	Read uint64 = 1 << iota
	Write
)

// for test purposes
var permLookup = map[uint64]string{
	Read:  "read",
	Write: "write",
}

// Protected resources
const (
	Article int = iota
	Video
)

// for test purposes
var resLookup = map[int]string{
	Article: "article",
	Video:   "video",
}

func TestACL(t *testing.T) {
	acl := ACL{}

	guest := acl.AddRole(Guest)
	guest.AddPerms(Article, Read)
	guest.AddPerms(Video, Read)

	user := acl.AddRole(User, guest)
	user.AddPerms(Video, Write)

	admin := acl.AddRole(Admin, user)
	admin.AddPerms(Article, Write)

	testCases := []struct {
		role     int
		resource int
		perm     uint64
		expCan   bool
	}{
		{
			role:     Guest,
			resource: Article,
			perm:     Read,
			expCan:   true,
		},
		{
			role:     User,
			resource: Video,
			perm:     Write,
			expCan:   true,
		},
		{
			role:     Guest,
			resource: Article,
			perm:     Write,
			expCan:   false,
		},
	}

	for _, tt := range testCases {
		if tt.expCan != acl.Can(tt.role, tt.resource, tt.perm) {
			if tt.expCan {
				t.Errorf("%s should be able to %s %s", roleLookup[tt.role],
					permLookup[tt.perm], resLookup[tt.resource])
				return
			}
			t.Errorf("%s shouldn't be able to %s %s", roleLookup[tt.role],
				permLookup[tt.perm], resLookup[tt.resource])
			return
		}
	}
}
