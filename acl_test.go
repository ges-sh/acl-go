package acl

import (
	"testing"
)

const (
	Create byte = 1 << iota
	Remove
	Rename
	FreeCreate
	RemoveAny
)

const (
	User int = iota
	Affiliate
	Manager
	Admin
)

const (
	Lists int = iota
	AffCampaigns
	Swipes
	Campaigns
)

func TestACL(t *testing.T) {
	a := Acl{
		User: map[int]byte{
			Lists:     Create | Rename,
			Campaigns: Create | Rename,
		},
		Affiliate: map[int]byte{
			AffCampaigns: Create | Rename | Remove,
		},
		Manager: map[int]byte{
			AffCampaigns: Create | Rename | RemoveAny,
		},
		Admin: map[int]byte{
			Lists:     Create | Remove | Rename,
			Campaigns: Create | Remove | Rename,
		},
	}

	testData := []struct {
		shouldCan bool
		Role      int
		Object    int
		Action    byte
	}{
		{
			true,
			User,
			Lists,
			Create,
		},
		{
			false,
			User,
			Lists,
			Remove,
		},
		{
			false,
			User,
			Campaigns,
			Remove,
		},
		{
			true,
			Affiliate,
			AffCampaigns,
			Create,
		},
		{
			false,
			Affiliate,
			Campaigns,
			Remove,
		},
		{
			true,
			Manager,
			AffCampaigns,
			RemoveAny,
		},
		{
			true,
			Admin,
			Lists,
			Remove,
		},
		{
			true,
			Admin,
			Campaigns,
			Remove,
		},
	}

	for _, v := range testData {
		can := a.Can(v.Role, v.Object, v.Action)
		if v.shouldCan != can {
			t.Errorf("Role %v, object %v, action %v, shouldCan %v, Can %v", v.Role, v.Object, v.Action, v.shouldCan, can)
		}
	}
}
