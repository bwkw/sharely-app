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

// Schedule is an object representing the database table.
type Schedule struct {
	ID            int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title         string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description   null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	StartDatetime time.Time   `boil:"start_datetime" json:"start_datetime" toml:"start_datetime" yaml:"start_datetime"`
	EndDatetime   time.Time   `boil:"end_datetime" json:"end_datetime" toml:"end_datetime" yaml:"end_datetime"`
	IsRepeated    bool        `boil:"is_repeated" json:"is_repeated" toml:"is_repeated" yaml:"is_repeated"`
	Color         int         `boil:"color" json:"color" toml:"color" yaml:"color"`
	UserGroupID   int         `boil:"user_group_id" json:"user_group_id" toml:"user_group_id" yaml:"user_group_id"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *scheduleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scheduleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ScheduleColumns = struct {
	ID            string
	Title         string
	Description   string
	StartDatetime string
	EndDatetime   string
	IsRepeated    string
	Color         string
	UserGroupID   string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	Title:         "title",
	Description:   "description",
	StartDatetime: "start_datetime",
	EndDatetime:   "end_datetime",
	IsRepeated:    "is_repeated",
	Color:         "color",
	UserGroupID:   "user_group_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var ScheduleTableColumns = struct {
	ID            string
	Title         string
	Description   string
	StartDatetime string
	EndDatetime   string
	IsRepeated    string
	Color         string
	UserGroupID   string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "schedules.id",
	Title:         "schedules.title",
	Description:   "schedules.description",
	StartDatetime: "schedules.start_datetime",
	EndDatetime:   "schedules.end_datetime",
	IsRepeated:    "schedules.is_repeated",
	Color:         "schedules.color",
	UserGroupID:   "schedules.user_group_id",
	CreatedAt:     "schedules.created_at",
	UpdatedAt:     "schedules.updated_at",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var ScheduleWhere = struct {
	ID            whereHelperint
	Title         whereHelperstring
	Description   whereHelpernull_String
	StartDatetime whereHelpertime_Time
	EndDatetime   whereHelpertime_Time
	IsRepeated    whereHelperbool
	Color         whereHelperint
	UserGroupID   whereHelperint
	CreatedAt     whereHelpernull_Time
	UpdatedAt     whereHelpernull_Time
}{
	ID:            whereHelperint{field: "`schedules`.`id`"},
	Title:         whereHelperstring{field: "`schedules`.`title`"},
	Description:   whereHelpernull_String{field: "`schedules`.`description`"},
	StartDatetime: whereHelpertime_Time{field: "`schedules`.`start_datetime`"},
	EndDatetime:   whereHelpertime_Time{field: "`schedules`.`end_datetime`"},
	IsRepeated:    whereHelperbool{field: "`schedules`.`is_repeated`"},
	Color:         whereHelperint{field: "`schedules`.`color`"},
	UserGroupID:   whereHelperint{field: "`schedules`.`user_group_id`"},
	CreatedAt:     whereHelpernull_Time{field: "`schedules`.`created_at`"},
	UpdatedAt:     whereHelpernull_Time{field: "`schedules`.`updated_at`"},
}

// ScheduleRels is where relationship names are stored.
var ScheduleRels = struct {
	UserGroup string
}{
	UserGroup: "UserGroup",
}

// scheduleR is where relationships are stored.
type scheduleR struct {
	UserGroup *UserGroup `boil:"UserGroup" json:"UserGroup" toml:"UserGroup" yaml:"UserGroup"`
}

// NewStruct creates a new relationship struct
func (*scheduleR) NewStruct() *scheduleR {
	return &scheduleR{}
}

// scheduleL is where Load methods for each relationship are stored.
type scheduleL struct{}

var (
	scheduleAllColumns            = []string{"id", "title", "description", "start_datetime", "end_datetime", "is_repeated", "color", "user_group_id", "created_at", "updated_at"}
	scheduleColumnsWithoutDefault = []string{"title", "description", "start_datetime", "end_datetime", "is_repeated", "color", "user_group_id", "created_at", "updated_at"}
	scheduleColumnsWithDefault    = []string{"id"}
	schedulePrimaryKeyColumns     = []string{"id"}
)

type (
	// ScheduleSlice is an alias for a slice of pointers to Schedule.
	// This should almost always be used instead of []Schedule.
	ScheduleSlice []*Schedule
	// ScheduleHook is the signature for custom Schedule hook methods
	ScheduleHook func(context.Context, boil.ContextExecutor, *Schedule) error

	scheduleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	scheduleType                 = reflect.TypeOf(&Schedule{})
	scheduleMapping              = queries.MakeStructMapping(scheduleType)
	schedulePrimaryKeyMapping, _ = queries.BindMapping(scheduleType, scheduleMapping, schedulePrimaryKeyColumns)
	scheduleInsertCacheMut       sync.RWMutex
	scheduleInsertCache          = make(map[string]insertCache)
	scheduleUpdateCacheMut       sync.RWMutex
	scheduleUpdateCache          = make(map[string]updateCache)
	scheduleUpsertCacheMut       sync.RWMutex
	scheduleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var scheduleBeforeInsertHooks []ScheduleHook
var scheduleBeforeUpdateHooks []ScheduleHook
var scheduleBeforeDeleteHooks []ScheduleHook
var scheduleBeforeUpsertHooks []ScheduleHook

var scheduleAfterInsertHooks []ScheduleHook
var scheduleAfterSelectHooks []ScheduleHook
var scheduleAfterUpdateHooks []ScheduleHook
var scheduleAfterDeleteHooks []ScheduleHook
var scheduleAfterUpsertHooks []ScheduleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Schedule) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Schedule) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Schedule) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Schedule) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Schedule) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Schedule) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Schedule) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Schedule) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Schedule) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scheduleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddScheduleHook registers your hook function for all future operations.
func AddScheduleHook(hookPoint boil.HookPoint, scheduleHook ScheduleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		scheduleBeforeInsertHooks = append(scheduleBeforeInsertHooks, scheduleHook)
	case boil.BeforeUpdateHook:
		scheduleBeforeUpdateHooks = append(scheduleBeforeUpdateHooks, scheduleHook)
	case boil.BeforeDeleteHook:
		scheduleBeforeDeleteHooks = append(scheduleBeforeDeleteHooks, scheduleHook)
	case boil.BeforeUpsertHook:
		scheduleBeforeUpsertHooks = append(scheduleBeforeUpsertHooks, scheduleHook)
	case boil.AfterInsertHook:
		scheduleAfterInsertHooks = append(scheduleAfterInsertHooks, scheduleHook)
	case boil.AfterSelectHook:
		scheduleAfterSelectHooks = append(scheduleAfterSelectHooks, scheduleHook)
	case boil.AfterUpdateHook:
		scheduleAfterUpdateHooks = append(scheduleAfterUpdateHooks, scheduleHook)
	case boil.AfterDeleteHook:
		scheduleAfterDeleteHooks = append(scheduleAfterDeleteHooks, scheduleHook)
	case boil.AfterUpsertHook:
		scheduleAfterUpsertHooks = append(scheduleAfterUpsertHooks, scheduleHook)
	}
}

// One returns a single schedule record from the query.
func (q scheduleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Schedule, error) {
	o := &Schedule{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for schedules")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Schedule records from the query.
func (q scheduleQuery) All(ctx context.Context, exec boil.ContextExecutor) (ScheduleSlice, error) {
	var o []*Schedule

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Schedule slice")
	}

	if len(scheduleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Schedule records in the query.
func (q scheduleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count schedules rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q scheduleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if schedules exists")
	}

	return count > 0, nil
}

// UserGroup pointed to by the foreign key.
func (o *Schedule) UserGroup(mods ...qm.QueryMod) userGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserGroupID),
	}

	queryMods = append(queryMods, mods...)

	query := UserGroups(queryMods...)
	queries.SetFrom(query.Query, "`user_group`")

	return query
}

// LoadUserGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (scheduleL) LoadUserGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSchedule interface{}, mods queries.Applicator) error {
	var slice []*Schedule
	var object *Schedule

	if singular {
		object = maybeSchedule.(*Schedule)
	} else {
		slice = *maybeSchedule.(*[]*Schedule)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &scheduleR{}
		}
		args = append(args, object.UserGroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &scheduleR{}
			}

			for _, a := range args {
				if a == obj.UserGroupID {
					continue Outer
				}
			}

			args = append(args, obj.UserGroupID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_group`),
		qm.WhereIn(`user_group.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserGroup")
	}

	var resultSlice []*UserGroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserGroup")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_group")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_group")
	}

	if len(scheduleAfterSelectHooks) != 0 {
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
		object.R.UserGroup = foreign
		if foreign.R == nil {
			foreign.R = &userGroupR{}
		}
		foreign.R.Schedules = append(foreign.R.Schedules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserGroupID == foreign.ID {
				local.R.UserGroup = foreign
				if foreign.R == nil {
					foreign.R = &userGroupR{}
				}
				foreign.R.Schedules = append(foreign.R.Schedules, local)
				break
			}
		}
	}

	return nil
}

// SetUserGroup of the schedule to the related item.
// Sets o.R.UserGroup to related.
// Adds o to related.R.Schedules.
func (o *Schedule) SetUserGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserGroup) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `schedules` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_group_id"}),
		strmangle.WhereClause("`", "`", 0, schedulePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserGroupID = related.ID
	if o.R == nil {
		o.R = &scheduleR{
			UserGroup: related,
		}
	} else {
		o.R.UserGroup = related
	}

	if related.R == nil {
		related.R = &userGroupR{
			Schedules: ScheduleSlice{o},
		}
	} else {
		related.R.Schedules = append(related.R.Schedules, o)
	}

	return nil
}

// Schedules retrieves all the records using an executor.
func Schedules(mods ...qm.QueryMod) scheduleQuery {
	mods = append(mods, qm.From("`schedules`"))
	return scheduleQuery{NewQuery(mods...)}
}

// FindSchedule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSchedule(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Schedule, error) {
	scheduleObj := &Schedule{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `schedules` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, scheduleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from schedules")
	}

	if err = scheduleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return scheduleObj, err
	}

	return scheduleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Schedule) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no schedules provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(scheduleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	scheduleInsertCacheMut.RLock()
	cache, cached := scheduleInsertCache[key]
	scheduleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			scheduleAllColumns,
			scheduleColumnsWithDefault,
			scheduleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(scheduleType, scheduleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(scheduleType, scheduleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `schedules` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `schedules` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `schedules` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, schedulePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into schedules")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == scheduleMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for schedules")
	}

CacheNoHooks:
	if !cached {
		scheduleInsertCacheMut.Lock()
		scheduleInsertCache[key] = cache
		scheduleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Schedule.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Schedule) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	scheduleUpdateCacheMut.RLock()
	cache, cached := scheduleUpdateCache[key]
	scheduleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			scheduleAllColumns,
			schedulePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update schedules, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `schedules` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, schedulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(scheduleType, scheduleMapping, append(wl, schedulePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update schedules row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for schedules")
	}

	if !cached {
		scheduleUpdateCacheMut.Lock()
		scheduleUpdateCache[key] = cache
		scheduleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q scheduleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for schedules")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ScheduleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `schedules` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, schedulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in schedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all schedule")
	}
	return rowsAff, nil
}

var mySQLScheduleUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Schedule) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no schedules provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(scheduleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLScheduleUniqueColumns, o)

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

	scheduleUpsertCacheMut.RLock()
	cache, cached := scheduleUpsertCache[key]
	scheduleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			scheduleAllColumns,
			scheduleColumnsWithDefault,
			scheduleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			scheduleAllColumns,
			schedulePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert schedules, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`schedules`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `schedules` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(scheduleType, scheduleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(scheduleType, scheduleMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for schedules")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == scheduleMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(scheduleType, scheduleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for schedules")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for schedules")
	}

CacheNoHooks:
	if !cached {
		scheduleUpsertCacheMut.Lock()
		scheduleUpsertCache[key] = cache
		scheduleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Schedule record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Schedule) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Schedule provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), schedulePrimaryKeyMapping)
	sql := "DELETE FROM `schedules` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for schedules")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q scheduleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no scheduleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from schedules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for schedules")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ScheduleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(scheduleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `schedules` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, schedulePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from schedule slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for schedules")
	}

	if len(scheduleAfterDeleteHooks) != 0 {
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
func (o *Schedule) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSchedule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ScheduleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ScheduleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schedulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `schedules`.* FROM `schedules` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, schedulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ScheduleSlice")
	}

	*o = slice

	return nil
}

// ScheduleExists checks if the Schedule row exists.
func ScheduleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `schedules` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if schedules exists")
	}

	return exists, nil
}
