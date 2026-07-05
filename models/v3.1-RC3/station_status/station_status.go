package station_status

import (
	station_status_v31_rc2 "github.com/phd-kerger/gbfs-go-adapter/models/v3.1-RC2/station_status"
)

// Describes the capacity and rental availability of the station
type StationStatus struct {
	station_status_v31_rc2.StationStatus
}
