// Code generated by sqlc. DO NOT EDIT.

package querytest

import ()

type FooTypeUserRole string

const (
	FooTypeUserRoleAdmin FooTypeUserRole = "admin"
	FooTypeUserRoleUser  FooTypeUserRole = "user"
)

func (e *FooTypeUserRole) Scan(src interface{}) error {
	*e = FooTypeUserRole(src.([]byte))
	return nil
}

type FooUser struct {
	Role FooTypeUserRole
}