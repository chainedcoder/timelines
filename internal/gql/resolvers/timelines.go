package resolvers

import (
	"context"
	"errors"
	"fmt"

	"chainedcoder/timelines/pkg/utils/consts"

	"chainedcoder/timelines/internal/logger"
	"chainedcoder/timelines/pkg/utils"

	"chainedcoder/timelines/internal/gql/models"
	tf "chainedcoder/timelines/internal/gql/resolvers/transformations"
	"chainedcoder/timelines/internal/orm"
	dbm "chainedcoder/timelines/internal/orm/models"
)

// CreateReformline creates a record
func (r *mutationResolver) CreateReformline(ctx context.Context, input models.ReformlineInput) (*models.Reformline, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Reformlines); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Reformlines, err)
	}
	return reformlineCreateUpdate(r, input, false, cu)
}

// UpdateReformline updates a record
func (r *mutationResolver) UpdateReformline(ctx context.Context, id string, input models.ReformlineInput) (*models.Reformline, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Reformlines); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Reformlines, err)
	}
	return reformlineCreateUpdate(r, input, true, cu, id)
}

// DeleteReformline deletes a record
func (r *mutationResolver) DeleteReformline(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Reformlines); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.Reformlines, err)
	}
	return reformlineDelete(r, id)
}

// Reformlines lists records
func (r *queryResolver) Reformlines(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Reformlines, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Reformlines); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Reformlines, err)
	}
	return reformlineList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func reformlineCreateUpdate(r *mutationResolver, input models.ReformlineInput, update bool, cu *dbm.User, ids ...string) (*models.Reformline, error) {
	dbo, err := tf.TransformReformlineInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", dbo)
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx.Create(dbo.HeadEvent).First(dbo.HeadEvent)
		fmt.Printf("%#v\n", dbo)
		dbo.HeadEventID = &dbo.HeadEvent.ID
		tx = tx.Create(dbo).First(dbo) // Create the reformline
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformReformline(dbo), tx.Error
}

func reformlineDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func reformlineList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Reformlines, error) {
	whereID := "id = ?"
	record := &models.Reformlines{}
	dbRecords := []*dbm.Reformline{}
	tx := r.ORM.DB.Begin()

	if offset != nil {
		tx = tx.Offset(*offset)
	}
	if limit != nil {
		tx = tx.Limit(*limit)
	}
	if orderBy != nil && sortDirection != nil {
		tx = tx.Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	}

	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}

	tx = tx.Preload("HeadEvent").Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformReformline(dbRec))
	}
	return record, tx.Error
}

// CreateSymbol creates a record
func (r *mutationResolver) CreateSymbol(ctx context.Context, input models.SymbolInput) (*models.Symbol, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Symbols); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Symbols, err)
	}
	return symbolCreateUpdate(r, input, false, cu)
}

// UpdateSymbol updates a record
func (r *mutationResolver) UpdateSymbol(ctx context.Context, id string, input models.SymbolInput) (*models.Symbol, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Symbols); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Symbols, err)
	}
	return symbolCreateUpdate(r, input, true, cu, id)
}

// DeleteSymbol deletes a record
func (r *mutationResolver) DeleteSymbol(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Symbols); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.Symbols, err)
	}
	return symbolDelete(r, id)
}

// Symbols lists records
func (r *queryResolver) Symbols(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Symbols, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Symbols); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Symbols, err)
	}
	return symbolList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func symbolCreateUpdate(r *mutationResolver, input models.SymbolInput, update bool, cu *dbm.User, ids ...string) (*models.Symbol, error) {
	dbo, err := tf.TransformSymbolInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the symbol
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformSymbol(dbo), tx.Error
}

func symbolDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func symbolList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Symbols, error) {
	whereID := "id = ?"
	record := &models.Symbols{}
	dbRecords := []*dbm.Symbol{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformSymbol(dbRec))
	}
	return record, tx.Error
}

// CreateWaymarkEvent creates a record
func (r *mutationResolver) CreateWaymarkEvent(ctx context.Context, input models.WaymarkEventInput) (*models.WaymarkEvent, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.WaymarkEvents); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.WaymarkEvents, err)
	}
	return waymarkEventCreateUpdate(r, input, false, cu)
}

// UpdateWaymarkEvent updates a record
func (r *mutationResolver) UpdateWaymarkEvent(ctx context.Context, id string, input models.WaymarkEventInput) (*models.WaymarkEvent, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.WaymarkEvents); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.WaymarkEvents, err)
	}
	return waymarkEventCreateUpdate(r, input, true, cu, id)
}

// DeleteWaymarkEvent deletes a record
func (r *mutationResolver) DeleteWaymarkEvent(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.WaymarkEvents); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.WaymarkEvents, err)
	}
	return waymarkEventDelete(r, id)
}

// WaymarkEvents lists records
func (r *queryResolver) WaymarkEvents(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.WaymarkEvents, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.WaymarkEvents); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.WaymarkEvents, err)
	}
	return waymarkEventList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func waymarkEventCreateUpdate(r *mutationResolver, input models.WaymarkEventInput, update bool, cu *dbm.User, ids ...string) (*models.WaymarkEvent, error) {
	dbo, err := tf.TransformWaymarkEventInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the waymarkEvent
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformWaymarkEvent(dbo), tx.Error
}

func waymarkEventDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func waymarkEventList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.WaymarkEvents, error) {
	whereID := "id = ?"
	record := &models.WaymarkEvents{}
	dbRecords := []*dbm.WaymarkEvent{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformWaymarkEvent(dbRec))
	}
	return record, tx.Error
}

// CreateWaymarkTag creates a record
func (r *mutationResolver) CreateWaymarkTag(ctx context.Context, input models.WaymarkTagInput) (*models.WaymarkTag, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.WaymarkTags); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.WaymarkTags, err)
	}
	return waymarkTagCreateUpdate(r, input, false, cu)
}

// UpdateWaymarkTag updates a record
func (r *mutationResolver) UpdateWaymarkTag(ctx context.Context, id string, input models.WaymarkTagInput) (*models.WaymarkTag, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.WaymarkTags); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.WaymarkTags, err)
	}
	return waymarkTagCreateUpdate(r, input, true, cu, id)
}

// DeleteWaymarkTag deletes a record
func (r *mutationResolver) DeleteWaymarkTag(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.WaymarkTags); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.WaymarkTags, err)
	}
	return waymarkTagDelete(r, id)
}

// WaymarkTags lists records
func (r *queryResolver) WaymarkTags(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.WaymarkTags, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.WaymarkTags); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.WaymarkTags, err)
	}
	return waymarkTagList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func waymarkTagCreateUpdate(r *mutationResolver, input models.WaymarkTagInput, update bool, cu *dbm.User, ids ...string) (*models.WaymarkTag, error) {
	dbo, err := tf.TransformWaymarkTagInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the waymarkTag
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformWaymarkTag(dbo), tx.Error
}

func waymarkTagDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func waymarkTagList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.WaymarkTags, error) {
	whereID := "id = ?"
	record := &models.WaymarkTags{}
	dbRecords := []*dbm.WaymarkTag{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformWaymarkTag(dbRec))
	}
	return record, tx.Error
}

// CreateWaymark creates a record
func (r *mutationResolver) CreateWaymark(ctx context.Context, input models.WaymarkInput) (*models.Waymark, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Waymarks); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Waymarks, err)
	}
	return waymarkCreateUpdate(r, input, false, cu)
}

// UpdateWaymark updates a record
func (r *mutationResolver) UpdateWaymark(ctx context.Context, id string, input models.WaymarkInput) (*models.Waymark, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Waymarks); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Waymarks, err)
	}
	return waymarkCreateUpdate(r, input, true, cu, id)
}

// DeleteWaymark deletes a record
func (r *mutationResolver) DeleteWaymark(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Waymarks); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.Waymarks, err)
	}
	return waymarkDelete(r, id)
}

// Waymarks lists records
func (r *queryResolver) Waymarks(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Waymarks, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Waymarks); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Waymarks, err)
	}
	return waymarkList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func waymarkCreateUpdate(r *mutationResolver, input models.WaymarkInput, update bool, cu *dbm.User, ids ...string) (*models.Waymark, error) {
	dbo, err := tf.TransformWaymarkInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the waymark
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformWaymark(dbo), tx.Error
}

func waymarkDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func waymarkList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Waymarks, error) {
	whereID := "id = ?"
	record := &models.Waymarks{}
	dbRecords := []*dbm.Waymark{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformWaymark(dbRec))
	}
	return record, tx.Error
}

// CreateTag creates a record
func (r *mutationResolver) CreateTag(ctx context.Context, input models.TagInput) (*models.Tag, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Tags); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Tags, err)
	}
	return tagCreateUpdate(r, input, false, cu)
}

// UpdateTag updates a record
func (r *mutationResolver) UpdateTag(ctx context.Context, id string, input models.TagInput) (*models.Tag, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Tags); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Tags, err)
	}
	return tagCreateUpdate(r, input, true, cu, id)
}

// DeleteTag deletes a record
func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Tags); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.Tags, err)
	}
	return tagDelete(r, id)
}

// Tags lists records
func (r *queryResolver) Tags(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Tags, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Tags); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Tags, err)
	}
	return tagList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func tagCreateUpdate(r *mutationResolver, input models.TagInput, update bool, cu *dbm.User, ids ...string) (*models.Tag, error) {
	dbo, err := tf.TransformTagInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the tag
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformTag(dbo), tx.Error
}

func tagDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func tagList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Tags, error) {
	whereID := "id = ?"
	record := &models.Tags{}
	dbRecords := []*dbm.Tag{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformTag(dbRec))
	}
	return record, tx.Error
}

// CreateThreadGroup creates a record
func (r *mutationResolver) CreateThreadGroup(ctx context.Context, input models.ThreadGroupInput) (*models.ThreadGroup, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.ThreadGroups); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.ThreadGroups, err)
	}
	return threadGroupCreateUpdate(r, input, false, cu)
}

// UpdateThreadGroup updates a record
func (r *mutationResolver) UpdateThreadGroup(ctx context.Context, id string, input models.ThreadGroupInput) (*models.ThreadGroup, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.ThreadGroups); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.ThreadGroups, err)
	}
	return threadGroupCreateUpdate(r, input, true, cu, id)
}

// DeleteThreadGroup deletes a record
func (r *mutationResolver) DeleteThreadGroup(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.ThreadGroups); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.ThreadGroups, err)
	}
	return threadGroupDelete(r, id)
}

// ThreadGroups lists records
func (r *queryResolver) ThreadGroups(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.ThreadGroups, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.ThreadGroups); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.ThreadGroups, err)
	}
	return threadGroupList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func threadGroupCreateUpdate(r *mutationResolver, input models.ThreadGroupInput, update bool, cu *dbm.User, ids ...string) (*models.ThreadGroup, error) {
	dbo, err := tf.TransformThreadGroupInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the threadGroup
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformThreadGroup(dbo), tx.Error
}

func threadGroupDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func threadGroupList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.ThreadGroups, error) {
	whereID := "id = ?"
	record := &models.ThreadGroups{}
	dbRecords := []*dbm.ThreadGroup{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformThreadGroup(dbRec))
	}
	return record, tx.Error
}

// CreateApplicationEvent creates a record
func (r *mutationResolver) CreateApplicationEvent(ctx context.Context, input models.ApplicationEventInput) (*models.ApplicationEvent, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.ApplicationEvents); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.ApplicationEvents, err)
	}
	return applicationEventCreateUpdate(r, input, false, cu)
}

// UpdateApplicationEvent updates a record
func (r *mutationResolver) UpdateApplicationEvent(ctx context.Context, id string, input models.ApplicationEventInput) (*models.ApplicationEvent, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.ApplicationEvents); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.ApplicationEvents, err)
	}
	return applicationEventCreateUpdate(r, input, true, cu, id)
}

// DeleteApplicationEvent deletes a record
func (r *mutationResolver) DeleteApplicationEvent(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.ApplicationEvents); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.ApplicationEvents, err)
	}
	return applicationEventDelete(r, id)
}

// ApplicationEvents lists records
func (r *queryResolver) ApplicationEvents(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.ApplicationEvents, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.ApplicationEvents); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.ApplicationEvents, err)
	}
	return applicationEventList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func applicationEventCreateUpdate(r *mutationResolver, input models.ApplicationEventInput, update bool, cu *dbm.User, ids ...string) (*models.ApplicationEvent, error) {
	dbo, err := tf.TransformApplicationEventInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the applicationEvent
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformApplicationEvent(dbo), tx.Error
}

func applicationEventDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func applicationEventList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.ApplicationEvents, error) {
	whereID := "id = ?"
	record := &models.ApplicationEvents{}
	dbRecords := []*dbm.ApplicationEvent{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformApplicationEvent(dbRec))
	}
	return record, tx.Error
}

// CreateBreakdown creates a record
func (r *mutationResolver) CreateBreakdown(ctx context.Context, input models.BreakdownInput) (*models.Breakdown, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Breakdowns); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Breakdowns, err)
	}
	return breakdownCreateUpdate(r, input, false, cu)
}

// UpdateBreakdown updates a record
func (r *mutationResolver) UpdateBreakdown(ctx context.Context, id string, input models.BreakdownInput) (*models.Breakdown, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Breakdowns); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Breakdowns, err)
	}
	return breakdownCreateUpdate(r, input, true, cu, id)
}

// DeleteBreakdown deletes a record
func (r *mutationResolver) DeleteBreakdown(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Breakdowns); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.Breakdowns, err)
	}
	return breakdownDelete(r, id)
}

// Breakdowns lists records
func (r *queryResolver) Breakdowns(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Breakdowns, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Breakdowns); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Breakdowns, err)
	}
	return breakdownList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func breakdownCreateUpdate(r *mutationResolver, input models.BreakdownInput, update bool, cu *dbm.User, ids ...string) (*models.Breakdown, error) {
	dbo, err := tf.TransformBreakdownInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the breakdown
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformBreakdown(dbo), tx.Error
}

func breakdownDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func breakdownList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Breakdowns, error) {
	whereID := "id = ?"
	record := &models.Breakdowns{}
	dbRecords := []*dbm.Breakdown{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformBreakdown(dbRec))
	}
	return record, tx.Error
}

// CreateMethodology creates a record
func (r *mutationResolver) CreateMethodology(ctx context.Context, input models.MethodologyInput) (*models.Methodology, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Methodologies); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Methodologies, err)
	}
	return methodologyCreateUpdate(r, input, false, cu)
}

// UpdateMethodology updates a record
func (r *mutationResolver) UpdateMethodology(ctx context.Context, id string, input models.MethodologyInput) (*models.Methodology, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Methodologies); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Methodologies, err)
	}
	return methodologyCreateUpdate(r, input, true, cu, id)
}

// DeleteMethodology deletes a record
func (r *mutationResolver) DeleteMethodology(ctx context.Context, id string) (bool, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Methodologies); !ok || err != nil {
		return false, logger.Errorfn(consts.EntityNames.Methodologies, err)
	}
	return methodologyDelete(r, id)
}

// Methodologies lists records
func (r *queryResolver) Methodologies(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Methodologies, error) {
	cu := getCurrentUser(ctx)
	if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Methodologies); !ok || err != nil {
		return nil, logger.Errorfn(consts.EntityNames.Methodologies, err)
	}
	return methodologyList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func methodologyCreateUpdate(r *mutationResolver, input models.MethodologyInput, update bool, cu *dbm.User, ids ...string) (*models.Methodology, error) {
	dbo, err := tf.TransformMethodologyInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the methodology
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformMethodology(dbo), tx.Error
}

func methodologyDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func methodologyList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Methodologies, error) {
	whereID := "id = ?"
	record := &models.Methodologies{}
	dbRecords := []*dbm.Methodology{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformMethodology(dbRec))
	}
	return record, tx.Error
}

// CreateEvent creates a record
func (r *mutationResolver) CreateEvent(ctx context.Context, input models.EventInput) (*models.Event, error) {
	cu := getCurrentUser(ctx)
	// if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Events); !ok || err != nil {
	// 	return nil, logger.Errorfn(consts.EntityNames.Events, err)
	// }
	return eventCreateUpdate(r, input, false, cu)
}

// UpdateEvent updates a record
func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, input models.EventInput) (*models.Event, error) {
	cu := getCurrentUser(ctx)
	// if ok, err := cu.HasPermission(consts.Permissions.Create, consts.EntityNames.Events); !ok || err != nil {
	// 	return nil, logger.Errorfn(consts.EntityNames.Events, err)
	// }
	return eventCreateUpdate(r, input, true, cu, id)
}

// DeleteEvent deletes a record
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	// cu := getCurrentUser(ctx)
	// if ok, err := cu.HasPermission(consts.Permissions.Delete, consts.EntityNames.Events); !ok || err != nil {
	// 	return false, logger.Errorfn(consts.EntityNames.Events, err)
	// }
	return eventDelete(r, id)
}

// Events lists records
func (r *queryResolver) Events(ctx context.Context, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Events, error) {
	// cu := getCurrentUser(ctx)
	// if ok, err := cu.HasPermission(consts.Permissions.List, consts.EntityNames.Events); !ok || err != nil {
	// 	return nil, logger.Errorfn(consts.EntityNames.Events, err)
	// }
	return eventList(r, id, filters, limit, offset, orderBy, sortDirection)
}

// ## Helper functions

func eventCreateUpdate(r *mutationResolver, input models.EventInput, update bool, cu *dbm.User, ids ...string) (*models.Event, error) {
	dbo, err := tf.TransformEventInput(&input, update, cu, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	tx := r.ORM.DB.Begin()
	defer tx.RollbackUnlessCommitted()
	if !update {
		tx = tx.Create(dbo).First(dbo) // Create the event
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	tx = tx.Commit()
	return tf.TransformEvent(dbo), tx.Error
}

func eventDelete(r *mutationResolver, id string) (bool, error) {
	return false, errors.New("not implemented")
}

func eventList(r *queryResolver, id *string, filters []*models.QueryFilter, limit *int, offset *int, orderBy *string, sortDirection *string) (*models.Events, error) {
	whereID := "id = ?"
	record := &models.Events{}
	dbRecords := []*dbm.Event{}
	tx := r.ORM.DB.Begin().
		Offset(*offset).Limit(*limit).Order(utils.ToSnakeCase(*orderBy) + " " + *sortDirection)
	if id != nil {
		tx = tx.Where(whereID, *id)
	}
	if filters != nil {
		if filtered, err := orm.ParseFilters(tx, filters); err == nil {
			tx = filtered
		} else {
			return nil, err
		}
	}
	tx = tx.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		record.List = append(record.List, tf.TransformEvent(dbRec))
	}
	return record, tx.Error
}
