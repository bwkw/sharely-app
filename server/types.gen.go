// Package calendar provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package calendar

import (
	"time"
)

// Defines values for ScheduleColor.
const (
	ScheduleColorN0 ScheduleColor = 0
	ScheduleColorN1 ScheduleColor = 1
)

// Defines values for UserSex.
const (
	UserSexN0 UserSex = 0
	UserSexN1 UserSex = 1
)

// Defines values for PostGroupsGroupIdSchedulesJSONBodyColor.
const (
	PostGroupsGroupIdSchedulesJSONBodyColorN0 PostGroupsGroupIdSchedulesJSONBodyColor = 0
	PostGroupsGroupIdSchedulesJSONBodyColorN1 PostGroupsGroupIdSchedulesJSONBodyColor = 1
)

// Defines values for PostUsersJSONBodySex.
const (
	PostUsersJSONBodySexN0 PostUsersJSONBodySex = 0
	PostUsersJSONBodySexN1 PostUsersJSONBodySex = 1
)

// Defines values for PutUsersUserIdJSONBodySex.
const (
	N0 PutUsersUserIdJSONBodySex = 0
	N1 PutUsersUserIdJSONBodySex = 1
)

// Group defines model for Group.
type Group struct {
	// Id Unique identifier for group
	Id *int `json:"id,omitempty"`

	// Name Name of the group
	Name *string `json:"name,omitempty"`
}

// Invitation defines model for Invitation.
type Invitation struct {
	// Group The group name which is invited
	Group *string `json:"group,omitempty"`

	// Receiver The user who received the invitation
	Receiver *string `json:"receiver,omitempty"`

	// Sender The user who sent the invitation
	Sender *string `json:"sender,omitempty"`
}

// Schedule defines model for Schedule.
type Schedule struct {
	// Color Color (0 for red, 1 for blue)
	Color *ScheduleColor `json:"color,omitempty"`

	// Description Detailed description of the schedule
	Description *string `json:"description,omitempty"`

	// EndDatetime End date and time of the schedule
	EndDatetime *time.Time `json:"end_datetime,omitempty"`
	Group       *Group     `json:"group,omitempty"`

	// Id Unique identifier for the schedule
	Id *int `json:"id,omitempty"`

	// StartDatetime Start date and time of the schedule
	StartDatetime *time.Time `json:"start_datetime,omitempty"`

	// Title Title of the schedule
	Title *string `json:"title,omitempty"`
	User  *User   `json:"user,omitempty"`
}

// ScheduleColor Color (0 for red, 1 for blue)
type ScheduleColor int

// User defines model for User.
type User struct {
	// Email Email of the user
	Email *string `json:"email,omitempty"`
	Group *Group  `json:"group,omitempty"`

	// Id Unique identifier for user
	Id *int `json:"id,omitempty"`

	// Image Image of the user
	Image *string `json:"image,omitempty"`

	// Name Name of the user
	Name *string `json:"name,omitempty"`

	// Password Password for the user
	Password *string `json:"password,omitempty"`

	// Sex Gender of the user (0 for male, 1 for female)
	Sex *UserSex `json:"sex,omitempty"`
}

// UserSex Gender of the user (0 for male, 1 for female)
type UserSex int

// PostGroupsJSONBody defines parameters for PostGroups.
type PostGroupsJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// PostGroupsGroupIdInvitationJSONBody defines parameters for PostGroupsGroupIdInvitation.
type PostGroupsGroupIdInvitationJSONBody struct {
	Email *string `json:"email,omitempty"`
}

// GetGroupsGroupIdSchedulesParams defines parameters for GetGroupsGroupIdSchedules.
type GetGroupsGroupIdSchedulesParams struct {
	Keyword *string `form:"keyword,omitempty" json:"keyword,omitempty"`
}

// PostGroupsGroupIdSchedulesJSONBody defines parameters for PostGroupsGroupIdSchedules.
type PostGroupsGroupIdSchedulesJSONBody struct {
	// Color Color (0 for red, 1 for blue)
	Color         *PostGroupsGroupIdSchedulesJSONBodyColor `json:"color,omitempty"`
	Description   *string                                  `json:"description,omitempty"`
	EndDatetime   *time.Time                               `json:"end_datetime,omitempty"`
	StartDatetime *time.Time                               `json:"start_datetime,omitempty"`
	Title         *string                                  `json:"title,omitempty"`
}

// PostGroupsGroupIdSchedulesJSONBodyColor defines parameters for PostGroupsGroupIdSchedules.
type PostGroupsGroupIdSchedulesJSONBodyColor int

// PutGroupsGroupIdSchedulesScheduleIdJSONBody defines parameters for PutGroupsGroupIdSchedulesScheduleId.
type PutGroupsGroupIdSchedulesScheduleIdJSONBody struct {
	Date        *time.Time `json:"date,omitempty"`
	Description *string    `json:"description,omitempty"`
	Title       *string    `json:"title,omitempty"`
}

// PostInvitationsInvitationCodeAcceptJSONBody defines parameters for PostInvitationsInvitationCodeAccept.
type PostInvitationsInvitationCodeAcceptJSONBody struct {
	Password *string `json:"password,omitempty"`
}

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	Group int `form:"group" json:"group"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody struct {
	// Email Email of the user
	Email *string `json:"email,omitempty"`

	// Image Image of the user
	Image *string `json:"image,omitempty"`

	// Name Name of the user
	Name *string `json:"name,omitempty"`

	// Password Password for the user
	Password *string `json:"password,omitempty"`

	// Sex Gender of the user (0 for male, 1 for female)
	Sex *PostUsersJSONBodySex `json:"sex,omitempty"`
}

// PostUsersJSONBodySex defines parameters for PostUsers.
type PostUsersJSONBodySex int

// PutUsersUserIdJSONBody defines parameters for PutUsersUserId.
type PutUsersUserIdJSONBody struct {
	// Email Email of the user
	Email *string `json:"email,omitempty"`

	// Image Image of the user
	Image *string `json:"image,omitempty"`

	// Name Name of the user
	Name *string `json:"name,omitempty"`

	// Password Password for the user
	Password *string `json:"password,omitempty"`

	// Sex Gender of the user (0 for male, 1 for female)
	Sex *PutUsersUserIdJSONBodySex `json:"sex,omitempty"`
}

// PutUsersUserIdJSONBodySex defines parameters for PutUsersUserId.
type PutUsersUserIdJSONBodySex int

// PostGroupsJSONRequestBody defines body for PostGroups for application/json ContentType.
type PostGroupsJSONRequestBody PostGroupsJSONBody

// PostGroupsGroupIdInvitationJSONRequestBody defines body for PostGroupsGroupIdInvitation for application/json ContentType.
type PostGroupsGroupIdInvitationJSONRequestBody PostGroupsGroupIdInvitationJSONBody

// PostGroupsGroupIdSchedulesJSONRequestBody defines body for PostGroupsGroupIdSchedules for application/json ContentType.
type PostGroupsGroupIdSchedulesJSONRequestBody PostGroupsGroupIdSchedulesJSONBody

// PutGroupsGroupIdSchedulesScheduleIdJSONRequestBody defines body for PutGroupsGroupIdSchedulesScheduleId for application/json ContentType.
type PutGroupsGroupIdSchedulesScheduleIdJSONRequestBody PutGroupsGroupIdSchedulesScheduleIdJSONBody

// PostInvitationsInvitationCodeAcceptJSONRequestBody defines body for PostInvitationsInvitationCodeAccept for application/json ContentType.
type PostInvitationsInvitationCodeAcceptJSONRequestBody PostInvitationsInvitationCodeAcceptJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody PutUsersUserIdJSONBody