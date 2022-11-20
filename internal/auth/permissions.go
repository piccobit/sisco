package auth

type Permissions uint64

const (
	Unknown Permissions = iota
	User                = 1 << (iota - 1)
	Service
	Admin
)
