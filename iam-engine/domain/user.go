package domain

type member struct {
	id string
}

type group struct {
	id      string
	members []member
}

type permission struct {
	id   string
	name string
}

type role struct {
	id          string
	name        string
	permissions []permission
}

type policy struct {
	id      string
	members []member
	groups  []group
	roles   []role
}

type resource struct {
	id string
}

/*
	Members (User/Group)
	Role (Read-only, Full-access)
	Policy (1 and 2)
	Resource (User/Group)
*/
