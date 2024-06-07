package mysql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RouteDetail struct {
	Id            uuid.UUID    `gorm:"primaryKey;not null" json:"id"`
	DestinationId uuid.UUID    `gorm:"type:varchar(191);index;not null" json:"destination_id"`
	Destination   Destination  `gorm:"foreignKey:DestinationId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"destination"`
	RouteId       uuid.UUID    `gorm:"type:varchar(191);index;not null" json:"route_id"`
	Route         Route        `gorm:"foreignKey:RouteId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"route"`
	Longitude     float64      `gorm:"type:decimal(9,6);not null" json:"longitude"`
	Latitude      float64      `gorm:"type:decimal(9,6);not null" json:"latitude"`
	Duration      int          `gorm:"type:int;not null" json:"duration"`
	Order         int          `gorm:"type:int(11);not null" json:"order"`
	VisitStart    sql.NullTime `gorm:"type:time;not null" json:"visit_start"`
	VisitEnd      sql.NullTime `gorm:"type:time;not null" json:"visit_end"`
	CreatedAt     time.Time    `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time    `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
	DeletedAt     time.Time    `json:"deleted_at"`
}

func ModifyColumnRoutesDetail(db *gorm.DB) {
	db.Exec("ALTER TABLE route_details MODIFY COLUMN visit_start TIME NOT NULL")
	db.Exec("ALTER TABLE route_details MODIFY COLUMN visit_end TIME NOT NULL")
}
