package types

type Roles struct {
	Id 	int
	Name string
}

type Authority_roles struct {
	Name string
}

type Authority_permissions struct {
	Name string
}

type Permissions struct {
	Id 	int
	Role_id int
	Role_name string
	Permission_id int
	Permission_name string
}

type CreateRole struct {
	Name string `json:"name"`
}

type CreateRolePermission struct {
	Role_id string
	Permission_id string
}

type DeleteRolePermission struct {
	Role_id string
	Permission_id string
}