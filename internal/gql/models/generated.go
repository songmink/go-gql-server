// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type User struct {
	ID          string         `json:"id"`
	Email       string         `json:"email"`
	AvatarURL   *string        `json:"avatarURL"`
	Name        *string        `json:"name"`
	FirstName   *string        `json:"firstName"`
	LastName    *string        `json:"lastName"`
	NickName    *string        `json:"nickName"`
	Description *string        `json:"description"`
	Location    *string        `json:"location"`
	APIkey      *string        `json:"APIkey"`
	Profiles    []*UserProfile `json:"profiles"`
	CreatedAt   *time.Time     `json:"createdAt"`
	UpdatedAt   *time.Time     `json:"updatedAt"`
}

type UserInput struct {
	Email          *string   `json:"email"`
	Password       *string   `json:"password"`
	AvatarURL      *string   `json:"avatarURL"`
	DisplayName    *string   `json:"displayName"`
	Name           *string   `json:"name"`
	FirstName      *string   `json:"firstName"`
	LastName       *string   `json:"lastName"`
	NickName       *string   `json:"nickName"`
	Description    *string   `json:"description"`
	Location       *string   `json:"location"`
	AddRoles       []*string `json:"addRoles"`
	RemRoles       []*string `json:"remRoles"`
	AddPermissions []*string `json:"addPermissions"`
	RemPermissions []*string `json:"remPermissions"`
}

type UserProfile struct {
	ID             int        `json:"id"`
	Email          string     `json:"email"`
	ExternalUserID *string    `json:"externalUserId"`
	AvatarURL      *string    `json:"avatarURL"`
	Name           *string    `json:"name"`
	FirstName      *string    `json:"firstName"`
	LastName       *string    `json:"lastName"`
	NickName       *string    `json:"nickName"`
	Description    *string    `json:"description"`
	Location       *string    `json:"location"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type Users struct {
	Count *int    `json:"count"`
	List  []*User `json:"list"`
}
