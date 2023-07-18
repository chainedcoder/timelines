package transformations

import (
	graphql "chainedcoder/timelines/internal/gql/models"
	models "chainedcoder/timelines/internal/orm/models"
)

func TransformReformline(reformline *models.Reformline) *graphql.Reformline {
    return &graphql.Reformline{
        ID:          reformline.ID.String(),
        Title:     &reformline.Title,
        Type:      &reformline.Type,
        Date:      &reformline.Date,
        Name:      &reformline.Name,
		CreatedAt:   *reformline.CreatedAt,
		UpdatedAt:   *reformline.UpdatedAt,
        HeadEvent: TransformEvent(reformline.HeadEvent),
    }
}

func TransformWaymark(waymark *models.Waymark) *graphql.Waymark {
    return &graphql.Waymark{
        ID:          waymark.ID.String(),
        Reformline:   TransformReformline(waymark.Reformline),
        Name:         &waymark.Name,
        NextWaymark:  TransformWaymark(waymark.NextWaymark),
        PrevWaymark:  TransformWaymark(waymark.PrevWaymark),
        Type:         &waymark.Type,
        Nickname:     &waymark.Nickname,
        Description:  &waymark.Description,
        Topic:        &waymark.Topic,
    }
}
func TransformEvent(event *models.Event) *graphql.Event {
    return &graphql.Event{
        ID:          event.ID.String(),
        Date:            &event.Date,
        Title:           &event.Title,
        Description:     &event.Description,
        LiteralLocation: &event.LiteralLocation,
		CreatedAt:   *event.CreatedAt,
		UpdatedAt:   *event.UpdatedAt,
    }
}

func TransformWaymarkEvent(waymarkEvent *models.WaymarkEvent) *graphql.WaymarkEvent {
    // var uuid string = waymarkEvent.UUID.String()
	return &graphql.WaymarkEvent{
        ID:          waymarkEvent.ID.String(),
        Title:     &waymarkEvent.Title,
        Event:     TransformEvent(waymarkEvent.Event),
        Waymark:   TransformWaymark(&waymarkEvent.Waymark),
		CreatedAt:   *waymarkEvent.CreatedAt,
		UpdatedAt:   *waymarkEvent.UpdatedAt,
    }
}

func TransformSymbol(symbol *models.Symbol) *graphql.Symbol {
    
	return &graphql.Symbol{
        ID:          symbol.ID.String(),
        Event:       TransformEvent(symbol.Event),
        Description: &symbol.Description,
        Name:        &symbol.Name,
		CreatedAt:   symbol.CreatedAt,
		UpdatedAt:   symbol.UpdatedAt,
    }
}

func TransformMethodology(methodology *models.Methodology) *graphql.Methodology {
    return &graphql.Methodology{
		ID:          methodology.ID.String(),
        Description: &methodology.Description,
        Name:        &methodology.Name,
		CreatedAt:   *methodology.CreatedAt,
		UpdatedAt:   *methodology.UpdatedAt,
    }
}

func TransformBreakdown(Breakdown *models.Breakdown) *graphql.Breakdown {
    return &graphql.Breakdown{
		ID:          Breakdown.ID.String(),
        Methodology:   TransformMethodology(Breakdown.Methodology),
        ThreadGroup:   TransformThreadGroup(Breakdown.ThreadGroup),
		CreatedAt:   *Breakdown.CreatedAt,
		UpdatedAt:   *Breakdown.UpdatedAt,
    }
}

func TransformApplicationEvent(applicationEvent *models.ApplicationEvent) *graphql.ApplicationEvent {
    return &graphql.ApplicationEvent{
        ID:          applicationEvent.ID.String(),
        WaymarkEvent:   TransformWaymarkEvent(applicationEvent.WaymarkEvent),
        Symbol:         TransformSymbol(applicationEvent.Symbol),
        Breakdown:      TransformBreakdown(applicationEvent.Breakdown),
    }
}

func TransformThreadGroup(threadGroup *models.ThreadGroup) *graphql.ThreadGroup {
    return &graphql.ThreadGroup{
        ID:          threadGroup.ID.String(),
        Name:   threadGroup.Name,
    }
}

func TransformTag(tag *models.Tag) *graphql.Tag {
    return &graphql.Tag{
        ID:          tag.ID.String(),
        Name:   tag.Name,
    }
}
func TransformWaymarkTag(tag *models.WaymarkTag) *graphql.WaymarkTag {
    return &graphql.WaymarkTag{
        ID:          tag.ID.String(),
        Tag:   TransformTag(tag.Tag),
		Waymark: TransformWaymark(tag.Waymark),
    }
}

func TransformApplicationEventInput(i *graphql.ApplicationEventInput, update bool, u *models.User, ids ...string) (o *models.ApplicationEvent, errors error) {

	if i.Name == "" && !update {
		return nil, errors
	}
	o = &models.ApplicationEvent{
		Name:  &i.Name,

	}
	if i.WaymarkEvent != nil {
		e, err := TransformWaymarkEventInput(i.WaymarkEvent, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.WaymarkEvent = e
	}
	if i.Symbol != nil {
		e, err := TransformSymbolInput(i.Symbol, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Symbol = e
	}
	if i.Breakdown != nil {
		e, err := TransformBreakdownInput(i.Breakdown, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Breakdown = e
	}
	return o, errors
}
func TransformBreakdownInput(i *graphql.BreakdownInput, update bool, u *models.User, ids ...string) (o *models.Breakdown, errors error) {

	if i.Name == "" && !update {
		return nil, errors
	}
	o = &models.Breakdown{
		Name:  i.Name,
	}
	if i.ThreadGroup != nil {
		e, err := TransformThreadGroupInput(i.ThreadGroup, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.ThreadGroup = e
	}

	if i.Methodology != nil {
		e, err := TransformMethodologyInput(i.Methodology, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Methodology = e
	}
	return o, errors
}

func TransformMethodologyInput(i *graphql.MethodologyInput, update bool, u *models.User, ids ...string) (o *models.Methodology, errors error) {

	if *i.Name == "" && !update {
		return nil, errors
	}
	o = &models.Methodology{
		Name:  *i.Name,
		Description: *i.Description,
	}
	return o, errors
}

func TransformThreadGroupInput(i *graphql.ThreadGroupInput, update bool, u *models.User, ids ...string) (o *models.ThreadGroup, errors error) {

	if i.Name == "" && !update {
		return nil, errors
	}
	o = &models.ThreadGroup{
		Name:  i.Name,
	}
	return o, errors
}

func TransformTagInput(i *graphql.TagInput, update bool, u *models.User, ids ...string) (o *models.Tag, errors error) {

	if i.Name == "" && !update {
		return nil, errors
	}
	o = &models.Tag{
		Name:  i.Name,
	}
	return o, errors
}

func TransformSymbolInput(i *graphql.SymbolInput, update bool, u *models.User, ids ...string) (o *models.Symbol, errors error) {

	if *i.Description == "" && !update {
		return nil, errors
	}
	o = &models.Symbol{
		Name:  *i.Name,
		Description: *i.Description,
	}
	if i.Event != nil {
		e, err := TransformEventInput(i.Event, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Event = e
	}
	return o, errors
}

func TransformWaymarkTagInput(i *graphql.WaymarkTagInput, update bool, u *models.User, ids ...string) (o *models.WaymarkTag, errors error) {

	if i.Tag == nil && !update {
		return nil, errors
	}
	if i.Waymark == nil && !update {
		return nil, errors
	}
	o = &models.WaymarkTag{
		// Tag: TransformTag(i.Tag),
	}
	if i.Tag != nil {
		n, err := TransformTagInput(i.Tag, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Tag = n
	}
	if i.Waymark != nil {
		n, err := TransformWaymarkInput(i.Waymark, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Waymark = n
	}
	return o, errors
}

// GQLInputWaymarkEventToDBWaymarkEvent transforms [waymarkEvent] gql input to db model
func TransformWaymarkEventInput(i *graphql.WaymarkEventInput, update bool, u *models.User, ids ...string) (o *models.WaymarkEvent, errors error) {
	if i.Title == nil && !update {
		return nil, errors
	}
	if i.Name == nil && !update {
		return nil, errors
	}
	o = &models.WaymarkEvent{
		Title: *i.Title,
		Name:  *i.Name,
	}
	if i.Event != nil {
		e, err := TransformEventInput(i.Event, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Event = e
	}
	if i.Waymark != nil {
		w, err := TransformWaymarkInput(i.Waymark, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Waymark = *w
	}
	if i.Description != nil {
		o.Description = *i.Description
	}
	return o, errors
}

func TransformEventInput(i *graphql.EventInput, update bool, u *models.User, ids ...string) (o *models.Event, errors error) {
	if i.Title == nil && !update {
		return nil, errors
	}
	if i.Description == nil && !update {
		return nil, errors
	}
	o = &models.Event{
		Title: *i.Title,
		Description: *i.Description,
	}
	if i.Date != nil {
		o.Date = *i.Date
	}
	if i.LiteralLocation != nil {
		o.LiteralLocation = *i.LiteralLocation
	}
	return o, errors
}

func TransformWaymarkInput(i *graphql.WaymarkInput, update bool, u *models.User, ids ...string) (o *models.Waymark, errors error) {

	if i.Name == nil && !update {
		return nil, errors
	}
	o = &models.Waymark{
		Name:  *i.Name,
	}
	if i.Reformline != nil {
		r, err := TransformReformlineInput(i.Reformline, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.Reformline = r
	}
	if i.NextWaymark != nil {
		n, err := TransformWaymarkInput(i.NextWaymark, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.NextWaymark = n
	}
	if i.PrevWaymark != nil {
		p, err := TransformWaymarkInput(i.PrevWaymark, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.PrevWaymark = p
	}
	if i.Type != nil {
		o.Type = *i.Type
	}
	if i.Nickname != nil {
		o.Nickname = *i.Nickname
	}
	if i.Description != nil {
		o.Description = *i.Description
	}
	if i.Topic != nil {
		o.Topic = *i.Topic
	}
	return o, errors
}

func TransformReformlineInput(i *graphql.ReformlineInput, update bool, u *models.User, ids ...string) (o *models.Reformline, errors error) {
	if i.Title == nil && !update {
		return nil, errors
	}
	if i.Type == nil && !update {
		return nil, errors
	}
	o = &models.Reformline{
		Title: *i.Title,
		Type:  *i.Type,
	}
	if i.Date != nil {
		o.Date = *i.Date
	}
	if i.Name != nil {
		o.Name = *i.Name
	}
	if i.HeadEvent != nil {
		e, err := TransformEventInput(i.HeadEvent, update, u, ids...)
		if err != nil {
			return nil, err
		}
		o.HeadEvent = e
	}
	return o, errors
}

