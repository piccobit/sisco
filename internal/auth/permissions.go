package auth

type Permissions uint64

const (
	Unauthorized Permissions = iota
	User                     = 1 << (iota - 1)
	Service
	ServiceOwnerOnly
	Admin
)
