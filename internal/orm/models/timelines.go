package models

import (
	"github.com/gofrs/uuid"
)

type Reformline struct {
	MinBaseModelSeq
	Title      string
	Type       string
	Date       string
	Name       string
	HeadEventID *uuid.UUID `gorm:"type:uuid"`
	HeadEvent  *Event `gorm:"association_autoupdate:false;association_autocreate:false"`
  }
  

  type Event struct {
	MinBaseModelSeq
	Title            string
	Description      string
	Date             string
	LiteralLocation  string
  }
  
  type ReformlineEvent struct {
	ReformlineID uuid.UUID `gorm:"index"`
	EventID uuid.UUID        `gorm:"index"`
  }

  type Methodology struct {
	MinBaseModelSeq
	Name      string
	Description string

  }

  type Breakdown struct {
	MinBaseModelSeq
	Name      string
	Methodology *Methodology
	ThreadGroup *ThreadGroup 
  }
  
  type ApplicationEvent struct {
	MinBaseModelSeq
	Name      *string
	WaymarkEvent *WaymarkEvent
	Symbol *Symbol
	Breakdown *Breakdown
  }
  type ThreadGroup struct {
	MinBaseModelSeq
	Name      string
  }
  type Tag struct {
	MinBaseModelSeq
	Name      string
  }

  type WaymarkTag struct {
	MinBaseModelSeq
	Name      string
	Tag  *Tag
  	Waymark *Waymark
  }

  
  type Waymark struct {
	MinBaseModelSeq
	Reformline  *Reformline
	Name        string
	NextWaymark *Waymark
	PrevWaymark *Waymark
	Type        string
	Nickname    string
	Description string
	Topic       string
  }

  type WaymarkEvent struct {
	MinBaseModelSeq
	Title       string
	Name        string
	Event       *Event
	Waymark Waymark
	
	Description string
  }

  type Symbol struct {
	MinBaseModelSeq
	Event       *Event
	Name        string
	Type        string
	Description string
  }
	  

// // BeforeDelete hook for Waymark
// func (w *Waymark) BeforeDelete(scope *gorm.Scope) error {
// 	db := scope.DB()
// 	// return updateReformLineHead(w.UUID)
	
// 	headWaymark := &Waymark{}
// 	if err := db.Where("event_uuid = ?", w.UUID).First(&headWaymark).Error; err != nil {
// 		return err
// 	}
// 	headWaymark.PrevWaymark = ""
// 	// Save the reformline
// 	if err := db.Save(&headWaymark).Error; err != nil {
// 		return err
// 	}
	
// 	reformline := &Reformline{}
// 	if err := db.Where("Reformline = ?", w.Reformline).First(&reformline).Error; err != nil {
// 		return err
// 	}

// 	// Update the headEvent of the reformline
// 	reformline.HeadEvent = headWaymark.NextWaymark

// 	// Save the reformline
// 	if err := db.Save(&reformline).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
