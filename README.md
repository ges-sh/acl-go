# acl

### Usage example

```go
	const (
		Create byte = 1 << iota
		Remove
		Rename
	)
	
	const (
		Lists int = iota
		AffCampaigns
	)

	const (
		User int = iota
		Affiliate
		Admin
	)
	
	func main() {
		var a = acl.Acl{}

		a.AddRole(User)
		a.AddPerms(User, Lists, Create, Remove, Rename)

		a.AddRole(Affiliate)
		a.AddPerms(Affiliate, AffCampaigns, Create, Remove, Rename)

		a.AddRole(Admin, User, Affiliate)

		fmt.Println(a.Can(User, Lists, Create)) // true
		fmt.Println(a.Can(User, AffCampaigns, Create)) // false
		fmt.Println(a.Can(Admin, AffCampaigns, Create)) // true
	}
```
