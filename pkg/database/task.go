package database

import (
	"fmt"

	"github.com/jovandeginste/workout-tracker/v2/pkg/background"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"gorm.io/gorm"
)

const (
	taskTypeUpdateMapDataAddress      background.TaskType = "updateMapDataAddress"
	taskTypeUpdateRouteSegmentAddress background.TaskType = "updateRouteSegmentAddress"
)

type updateMapDataAddressTask struct {
	mapDataID uint64
}

func NewUpdateMapDataAddressTask(mapDataID uint64) background.Task {
	return &updateMapDataAddressTask{mapDataID: mapDataID}
}

func (t *updateMapDataAddressTask) Run(db *gorm.DB) error {
	var m MapData
	if err := db.First(&m, t.mapDataID).Error; err != nil {
		return err
	}
	if m.Address != nil || m.Center.IsZero() {
		return nil
	}
	addr, err := geocoder.Reverse(geocoder.Query{
		Lat:    m.Center.Lat,
		Lon:    m.Center.Lng,
		Format: "json",
	})
	if err != nil {
		return fmt.Errorf("geocoder search error: %w", err)
	}

	if addr == nil || m.HasAddressString() {
		return nil
	}

	m.AddressString = GetAddressString(addr)

	return m.Save(db)
}

func (t *updateMapDataAddressTask) TaskType() background.TaskType {
	return taskTypeUpdateMapDataAddress
}

type updateRouteSegmentAddressTask struct {
	id uint64
}

func NewUpdateRouteSegmentAddressTask(id uint64) background.Task {
	return &updateRouteSegmentAddressTask{id: id}
}

func (t *updateRouteSegmentAddressTask) Run(db *gorm.DB) error {
	var rs RouteSegment
	if err := db.First(&rs, t.id).Error; err != nil {
		return err
	}
	if rs.Address != nil || rs.Center.IsZero() {
		return nil
	}
	addr, err := geocoder.Reverse(geocoder.Query{
		Lat:    rs.Center.Lat,
		Lon:    rs.Center.Lng,
		Format: "json",
	})
	if err != nil {
		return fmt.Errorf("geocoder search error: %w", err)
	}

	if addr == nil {
		return nil
	}

	rs.GeoAddress = addr
	rs.AddressString = GetAddressString(addr)

	return rs.Save(db)
}

func (t *updateRouteSegmentAddressTask) TaskType() background.TaskType {
	return taskTypeUpdateRouteSegmentAddress
}
