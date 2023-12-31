// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type ApplicationEvent struct {
	ID           string        `json:"id"`
	Name         string        `json:"Name"`
	WaymarkEvent *WaymarkEvent `json:"WaymarkEvent,omitempty"`
	Symbol       *Symbol       `json:"Symbol,omitempty"`
	Breakdown    *Breakdown    `json:"Breakdown,omitempty"`
	CreatedAt    time.Time     `json:"CreatedAt"`
	UpdatedAt    time.Time     `json:"UpdatedAt"`
}

type ApplicationEventInput struct {
	Name         string             `json:"Name"`
	WaymarkEvent *WaymarkEventInput `json:"WaymarkEvent,omitempty"`
	Symbol       *SymbolInput       `json:"Symbol,omitempty"`
	Breakdown    *BreakdownInput    `json:"Breakdown,omitempty"`
}

type ApplicationEvents struct {
	Count *int                `json:"count,omitempty"`
	List  []*ApplicationEvent `json:"list"`
}

type Breakdown struct {
	ID          string       `json:"id"`
	Name        string       `json:"Name"`
	Methodology *Methodology `json:"Methodology,omitempty"`
	ThreadGroup *ThreadGroup `json:"ThreadGroup,omitempty"`
	CreatedAt   time.Time    `json:"CreatedAt"`
	UpdatedAt   time.Time    `json:"UpdatedAt"`
}

type BreakdownInput struct {
	Name        string            `json:"Name"`
	Methodology *MethodologyInput `json:"Methodology,omitempty"`
	ThreadGroup *ThreadGroupInput `json:"ThreadGroup,omitempty"`
}

type Breakdowns struct {
	Count *int         `json:"count,omitempty"`
	List  []*Breakdown `json:"list"`
}

type Event struct {
	ID              string    `json:"id"`
	Title           *string   `json:"title,omitempty"`
	Description     *string   `json:"description,omitempty"`
	Date            *string   `json:"date,omitempty"`
	LiteralLocation *string   `json:"literal_location,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type EventInput struct {
	Title           *string `json:"title,omitempty"`
	Description     *string `json:"description,omitempty"`
	Date            *string `json:"date,omitempty"`
	LiteralLocation *string `json:"literal_location,omitempty"`
}

type Events struct {
	Count *int     `json:"count,omitempty"`
	List  []*Event `json:"list"`
}

type Methodologies struct {
	Count *int           `json:"count,omitempty"`
	List  []*Methodology `json:"list"`
}

type Methodology struct {
	ID          string    `json:"id"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type MethodologyInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type QueryFilter struct {
	Field         string             `json:"field"`
	LinkOperation *LinkOperationType `json:"linkOperation,omitempty"`
	Op            OperationType      `json:"op"`
	Value         interface{}        `json:"value,omitempty"`
	Values        []interface{}      `json:"values,omitempty"`
}

type Reformline struct {
	ID        string    `json:"id"`
	Title     *string   `json:"title,omitempty"`
	Type      *string   `json:"type,omitempty"`
	Date      *string   `json:"date,omitempty"`
	Name      *string   `json:"name,omitempty"`
	HeadEvent *Event    `json:"headEvent,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ReformlineInput struct {
	Title     *string     `json:"title,omitempty"`
	Type      *string     `json:"type,omitempty"`
	Date      *string     `json:"date,omitempty"`
	Head      *string     `json:"head,omitempty"`
	Name      *string     `json:"name,omitempty"`
	HeadEvent *EventInput `json:"headEvent,omitempty"`
}

type Reformlines struct {
	Count *int          `json:"count,omitempty"`
	List  []*Reformline `json:"list"`
}

type Symbol struct {
	ID          string     `json:"id"`
	Event       *Event     `json:"event,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type SymbolInput struct {
	Event       *EventInput `json:"event,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Type        *string     `json:"type,omitempty"`
	Description *string     `json:"description,omitempty"`
}

type Symbols struct {
	Count *int      `json:"count,omitempty"`
	List  []*Symbol `json:"list"`
}

type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"Name"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type TagInput struct {
	Name string `json:"Name"`
}

type Tags struct {
	Count *int   `json:"count,omitempty"`
	List  []*Tag `json:"list"`
}

type ThreadGroup struct {
	ID        string    `json:"id"`
	Name      string    `json:"Name"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type ThreadGroupInput struct {
	Name string `json:"Name"`
}

type ThreadGroups struct {
	Count *int           `json:"count,omitempty"`
	List  []*ThreadGroup `json:"list"`
}

type User struct {
	ID          string         `json:"id"`
	Email       string         `json:"email"`
	AvatarURL   *string        `json:"avatarURL,omitempty"`
	Name        *string        `json:"name,omitempty"`
	FirstName   *string        `json:"firstName,omitempty"`
	LastName    *string        `json:"lastName,omitempty"`
	NickName    *string        `json:"nickName,omitempty"`
	Description *string        `json:"description,omitempty"`
	Location    *string        `json:"location,omitempty"`
	APIkey      *string        `json:"APIkey,omitempty"`
	Profiles    []*UserProfile `json:"profiles"`
	CreatedBy   *User          `json:"createdBy,omitempty"`
	UpdatedBy   *User          `json:"updatedBy,omitempty"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
}

type UserInput struct {
	Email          *string   `json:"email,omitempty"`
	Password       *string   `json:"password,omitempty"`
	AvatarURL      *string   `json:"avatarURL,omitempty"`
	DisplayName    *string   `json:"displayName,omitempty"`
	Name           *string   `json:"name,omitempty"`
	FirstName      *string   `json:"firstName,omitempty"`
	LastName       *string   `json:"lastName,omitempty"`
	NickName       *string   `json:"nickName,omitempty"`
	Description    *string   `json:"description,omitempty"`
	Location       *string   `json:"location,omitempty"`
	AddRoles       []*string `json:"addRoles,omitempty"`
	RemRoles       []*string `json:"remRoles,omitempty"`
	AddPermissions []*string `json:"addPermissions,omitempty"`
	RemPermissions []*string `json:"remPermissions,omitempty"`
}

type UserProfile struct {
	ID             int        `json:"id"`
	Email          string     `json:"email"`
	ExternalUserID *string    `json:"externalUserId,omitempty"`
	AvatarURL      *string    `json:"avatarURL,omitempty"`
	Name           *string    `json:"name,omitempty"`
	FirstName      *string    `json:"firstName,omitempty"`
	LastName       *string    `json:"lastName,omitempty"`
	NickName       *string    `json:"nickName,omitempty"`
	Description    *string    `json:"description,omitempty"`
	Location       *string    `json:"location,omitempty"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	CreatedBy      *User      `json:"createdBy,omitempty"`
	UpdatedBy      *User      `json:"updatedBy,omitempty"`
}

type Users struct {
	Count *int    `json:"count,omitempty"`
	List  []*User `json:"list"`
}

type Waymark struct {
	ID          string      `json:"id"`
	Reformline  *Reformline `json:"reformline,omitempty"`
	Name        *string     `json:"name,omitempty"`
	NextWaymark *Waymark    `json:"next_waymark,omitempty"`
	PrevWaymark *Waymark    `json:"prev_waymark,omitempty"`
	Type        *string     `json:"type,omitempty"`
	Nickname    *string     `json:"nickname,omitempty"`
	Description *string     `json:"description,omitempty"`
	Topic       *string     `json:"topic,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type WaymarkEvent struct {
	ID          string    `json:"id"`
	Title       *string   `json:"title,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Event       *Event    `json:"event,omitempty"`
	Waymark     *Waymark  `json:"waymark,omitempty"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type WaymarkEventInput struct {
	Title       *string       `json:"title,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Event       *EventInput   `json:"event,omitempty"`
	Waymark     *WaymarkInput `json:"waymark,omitempty"`
	Description *string       `json:"description,omitempty"`
}

type WaymarkEvents struct {
	Count *int            `json:"count,omitempty"`
	List  []*WaymarkEvent `json:"list"`
}

type WaymarkInput struct {
	Reformline  *ReformlineInput `json:"reformline,omitempty"`
	Name        *string          `json:"name,omitempty"`
	NextWaymark *WaymarkInput    `json:"next_waymark,omitempty"`
	PrevWaymark *WaymarkInput    `json:"prev_waymark,omitempty"`
	Type        *string          `json:"type,omitempty"`
	Nickname    *string          `json:"nickname,omitempty"`
	Description *string          `json:"description,omitempty"`
	Topic       *string          `json:"topic,omitempty"`
}

type WaymarkTag struct {
	ID      string   `json:"id"`
	Tag     *Tag     `json:"tag,omitempty"`
	Waymark *Waymark `json:"waymark,omitempty"`
}

type WaymarkTagInput struct {
	Tag     *TagInput     `json:"tag,omitempty"`
	Waymark *WaymarkInput `json:"waymark,omitempty"`
}

type WaymarkTags struct {
	Count *int          `json:"count,omitempty"`
	List  []*WaymarkTag `json:"list"`
}

type Waymarks struct {
	Count *int       `json:"count,omitempty"`
	List  []*Waymark `json:"list"`
}

type LinkOperationType string

const (
	LinkOperationTypeAnd LinkOperationType = "AND"
	LinkOperationTypeOr  LinkOperationType = "OR"
)

var AllLinkOperationType = []LinkOperationType{
	LinkOperationTypeAnd,
	LinkOperationTypeOr,
}

func (e LinkOperationType) IsValid() bool {
	switch e {
	case LinkOperationTypeAnd, LinkOperationTypeOr:
		return true
	}
	return false
}

func (e LinkOperationType) String() string {
	return string(e)
}

func (e *LinkOperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LinkOperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LinkOperationType", str)
	}
	return nil
}

func (e LinkOperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OperationType string

const (
	OperationTypeEquals           OperationType = "Equals"
	OperationTypeNotEquals        OperationType = "NotEquals"
	OperationTypeLessThan         OperationType = "LessThan"
	OperationTypeLessThanEqual    OperationType = "LessThanEqual"
	OperationTypeGreaterThan      OperationType = "GreaterThan"
	OperationTypeGreaterThanEqual OperationType = "GreaterThanEqual"
	OperationTypeIs               OperationType = "Is"
	OperationTypeIsNull           OperationType = "IsNull"
	OperationTypeIsNotNull        OperationType = "IsNotNull"
	OperationTypeIn               OperationType = "In"
	OperationTypeNotIn            OperationType = "NotIn"
	OperationTypeLike             OperationType = "Like"
	OperationTypeILike            OperationType = "ILike"
	OperationTypeNotLike          OperationType = "NotLike"
	OperationTypeBetween          OperationType = "Between"
	OperationTypeMatch            OperationType = "Match"
)

var AllOperationType = []OperationType{
	OperationTypeEquals,
	OperationTypeNotEquals,
	OperationTypeLessThan,
	OperationTypeLessThanEqual,
	OperationTypeGreaterThan,
	OperationTypeGreaterThanEqual,
	OperationTypeIs,
	OperationTypeIsNull,
	OperationTypeIsNotNull,
	OperationTypeIn,
	OperationTypeNotIn,
	OperationTypeLike,
	OperationTypeILike,
	OperationTypeNotLike,
	OperationTypeBetween,
	OperationTypeMatch,
}

func (e OperationType) IsValid() bool {
	switch e {
	case OperationTypeEquals, OperationTypeNotEquals, OperationTypeLessThan, OperationTypeLessThanEqual, OperationTypeGreaterThan, OperationTypeGreaterThanEqual, OperationTypeIs, OperationTypeIsNull, OperationTypeIsNotNull, OperationTypeIn, OperationTypeNotIn, OperationTypeLike, OperationTypeILike, OperationTypeNotLike, OperationTypeBetween, OperationTypeMatch:
		return true
	}
	return false
}

func (e OperationType) String() string {
	return string(e)
}

func (e *OperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationType", str)
	}
	return nil
}

func (e OperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
