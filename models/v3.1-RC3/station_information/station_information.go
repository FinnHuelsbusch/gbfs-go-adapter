package station_information

import (
	station_information_v31_rc2 "github.com/phd-kerger/gbfs-go-adapter/models/v3.1-RC2/station_information"
)

// List of all stations, their capacities and locations. REQUIRED of systems utilizing docks.
type StationInformation struct {
	station_information_v31_rc2.StationInformation
}
