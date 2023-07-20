// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserGroup is an object representing the database table.
type UserGroup struct {
	UserID    int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GroupID   int       `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *userGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserGroupColumns = struct {
	UserID    string
	GroupID   string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_id",
	GroupID:   "group_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserGroupTableColumns = struct {
	UserID    string
	GroupID   string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_group.user_id",
	GroupID:   "user_group.group_id",
	CreatedAt: "user_group.created_at",
	UpdatedAt: "user_group.updated_at",
}

// Generated where

var UserGroupWhere = struct {
	UserID    whereHelperint
	GroupID   whereHelperint
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	UserID:    whereHelperint{field: "`user_group`.`user_id`"},
	GroupID:   whereHelperint{field: "`user_group`.`group_id`"},
	CreatedAt: whereHelpernull_Time{field: "`user_group`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`user_group`.`updated_at`"},
}

// UserGroupRels is where relationship names are stored.
var UserGroupRels = struct {
	User  string
	Group string
}{
	User:  "User",
	Group: "Group",
}

// userGroupR is where relationships are stored.
type userGroupR struct {
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
	Group *Group `boil:"Group" json:"Group" toml:"Group" yaml:"Group"`
}

// NewStruct creates a new relationship struct
func (*userGroupR) NewStruct() *userGroupR {
	return &userGroupR{}
}

// userGroupL is where Load methods for each relationship are stored.
type userGroupL struct{}

var (
	userGroupAllColumns            = []string{"user_id", "group_id", "created_at", "updated_at"}
	userGroupColumnsWithoutDefault = []string{"user_id", "group_id"}
	userGroupColumnsWithDefault    = []string{"created_at", "updated_at"}
	userGroupPrimaryKeyColumns     = []string{"user_id", "group_id"}
)

type (
	// UserGroupSlice is an alias for a slice of pointers to UserGroup.
	// This should almost always be used instead of []UserGroup.
	UserGroupSlice []*UserGroup
	// UserGroupHook is the signature for custom UserGroup hook methods
	UserGroupHook func(context.Context, boil.ContextExecutor, *UserGroup) error

	userGroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userGroupType                 = reflect.TypeOf(&UserGroup{})
	userGroupMapping              = queries.MakeStructMapping(userGroupType)
	userGroupPrimaryKeyMapping, _ = queries.BindMapping(userGroupType, userGroupMapping, userGroupPrimaryKeyColumns)
	userGroupInsertCacheMut       sync.RWMutex
	userGroupInsertCache          = make(map[string]insertCache)
	userGroupUpdateCacheMut       sync.RWMutex
	userGroupUpdateCache          = make(map[string]updateCache)
	userGroupUpsertCacheMut       sync.RWMutex
	userGroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userGroupBeforeInsertHooks []UserGroupHook
var userGroupBeforeUpdateHooks []UserGroupHook
var userGroupBeforeDeleteHooks []UserGroupHook
var userGroupBeforeUpsertHooks []UserGroupHook

var userGroupAfterInsertHooks []UserGroupHook
var userGroupAfterSelectHooks []UserGroupHook
var userGroupAfterUpdateHooks []UserGroupHook
var userGroupAfterDeleteHooks []UserGroupHook
var userGroupAfterUpsertHooks []UserGroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserGroup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserGroup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserGroup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserGroup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserGroup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserGroup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserGroup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserGroup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserGroup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGroupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserGroupHook registers your hook function for all future operations.
func AddUserGroupHook(hookPoint boil.HookPoint, userGroupHook UserGroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userGroupBeforeInsertHooks = append(userGroupBeforeInsertHooks, userGroupHook)
	case boil.BeforeUpdateHook:
		userGroupBeforeUpdateHooks = append(userGroupBeforeUpdateHooks, userGroupHook)
	case boil.BeforeDeleteHook:
		userGroupBeforeDeleteHooks = append(userGroupBeforeDeleteHooks, userGroupHook)
	case boil.BeforeUpsertHook:
		userGroupBeforeUpsertHooks = append(userGroupBeforeUpsertHooks, userGroupHook)
	case boil.AfterInsertHook:
		userGroupAfterInsertHooks = append(userGroupAfterInsertHooks, userGroupHook)
	case boil.AfterSelectHook:
		userGroupAfterSelectHooks = append(userGroupAfterSelectHooks, userGroupHook)
	case boil.AfterUpdateHook:
		userGroupAfterUpdateHooks = append(userGroupAfterUpdateHooks, userGroupHook)
	case boil.AfterDeleteHook:
		userGroupAfterDeleteHooks = append(userGroupAfterDeleteHooks, userGroupHook)
	case boil.AfterUpsertHook:
		userGroupAfterUpsertHooks = append(userGroupAfterUpsertHooks, userGroupHook)
	}
}

// One returns a single userGroup record from the query.
func (q userGroupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserGroup, error) {
	o := &UserGroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_group")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserGroup records from the query.
func (q userGroupQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserGroupSlice, error) {
	var o []*UserGroup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserGroup slice")
	}

	if len(userGroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserGroup records in the query.
func (q userGroupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_group rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userGroupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_group exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserGroup) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// Group pointed to by the foreign key.
func (o *UserGroup) Group(mods ...qm.QueryMod) groupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := Groups(queryMods...)
	queries.SetFrom(query.Query, "`groups`")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userGroupL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserGroup interface{}, mods queries.Applicator) error {
	var slice []*UserGroup
	var object *UserGroup

	if singular {
		object = maybeUserGroup.(*UserGroup)
	} else {
		slice = *maybeUserGroup.(*[]*UserGroup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userGroupR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userGroupR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userGroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserGroups = append(foreign.R.UserGroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserGroups = append(foreign.R.UserGroups, local)
				break
			}
		}
	}

	return nil
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userGroupL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserGroup interface{}, mods queries.Applicator) error {
	var slice []*UserGroup
	var object *UserGroup

	if singular {
		object = maybeUserGroup.(*UserGroup)
	} else {
		slice = *maybeUserGroup.(*[]*UserGroup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userGroupR{}
		}
		args = append(args, object.GroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userGroupR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`groups`),
		qm.WhereIn(`groups.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Group")
	}

	var resultSlice []*Group
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Group")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for groups")
	}

	if len(userGroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &groupR{}
		}
		foreign.R.UserGroups = append(foreign.R.UserGroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &groupR{}
				}
				foreign.R.UserGroups = append(foreign.R.UserGroups, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userGroup to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserGroups.
func (o *UserGroup) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, userGroupPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GroupID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userGroupR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserGroups: UserGroupSlice{o},
		}
	} else {
		related.R.UserGroups = append(related.R.UserGroups, o)
	}

	return nil
}

// SetGroup of the userGroup to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.UserGroups.
func (o *UserGroup) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Group) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `user_group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
		strmangle.WhereClause("`", "`", 0, userGroupPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.GroupID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.ID
	if o.R == nil {
		o.R = &userGroupR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &groupR{
			UserGroups: UserGroupSlice{o},
		}
	} else {
		related.R.UserGroups = append(related.R.UserGroups, o)
	}

	return nil
}

// UserGroups retrieves all the records using an executor.
func UserGroups(mods ...qm.QueryMod) userGroupQuery {
	mods = append(mods, qm.From("`user_group`"))
	return userGroupQuery{NewQuery(mods...)}
}

// FindUserGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserGroup(ctx context.Context, exec boil.ContextExecutor, userID int, groupID int, selectCols ...string) (*UserGroup, error) {
	userGroupObj := &UserGroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_group` where `user_id`=? AND `group_id`=?", sel,
	)

	q := queries.Raw(query, userID, groupID)

	err := q.Bind(ctx, exec, userGroupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_group")
	}

	if err = userGroupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userGroupObj, err
	}

	return userGroupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserGroup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_group provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userGroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userGroupInsertCacheMut.RLock()
	cache, cached := userGroupInsertCache[key]
	userGroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userGroupAllColumns,
			userGroupColumnsWithDefault,
			userGroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userGroupType, userGroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userGroupType, userGroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_group` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_group` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_group` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userGroupPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_group")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UserID,
		o.GroupID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_group")
	}

CacheNoHooks:
	if !cached {
		userGroupInsertCacheMut.Lock()
		userGroupInsertCache[key] = cache
		userGroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserGroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserGroup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userGroupUpdateCacheMut.RLock()
	cache, cached := userGroupUpdateCache[key]
	userGroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userGroupAllColumns,
			userGroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_group, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_group` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userGroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userGroupType, userGroupMapping, append(wl, userGroupPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update user_group row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_group")
	}

	if !cached {
		userGroupUpdateCacheMut.Lock()
		userGroupUpdateCache[key] = cache
		userGroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userGroupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_group")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserGroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userGroupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userGroup")
	}
	return rowsAff, nil
}

var mySQLUserGroupUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserGroup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_group provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userGroupColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserGroupUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userGroupUpsertCacheMut.RLock()
	cache, cached := userGroupUpsertCache[key]
	userGroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userGroupAllColumns,
			userGroupColumnsWithDefault,
			userGroupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userGroupAllColumns,
			userGroupPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert user_group, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`user_group`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_group` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userGroupType, userGroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userGroupType, userGroupMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for user_group")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userGroupType, userGroupMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for user_group")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_group")
	}

CacheNoHooks:
	if !cached {
		userGroupUpsertCacheMut.Lock()
		userGroupUpsertCache[key] = cache
		userGroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserGroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserGroup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserGroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userGroupPrimaryKeyMapping)
	sql := "DELETE FROM `user_group` WHERE `user_id`=? AND `group_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_group")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userGroupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userGroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_group")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserGroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userGroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_group` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userGroupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_group")
	}

	if len(userGroupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserGroup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserGroup(ctx, exec, o.UserID, o.GroupID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserGroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserGroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_group`.* FROM `user_group` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userGroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserGroupSlice")
	}

	*o = slice

	return nil
}

// UserGroupExists checks if the UserGroup row exists.
func UserGroupExists(ctx context.Context, exec boil.ContextExecutor, userID int, groupID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_group` where `user_id`=? AND `group_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID, groupID)
	}
	row := exec.QueryRowContext(ctx, sql, userID, groupID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_group exists")
	}

	return exists, nil
}
