package vehicleavailability

import (
	vehicle_availability_v31_rc2 "github.com/phd-kerger/gbfs-go-adapter/models/v3.1-RC2/vehicle_availability"
	"github.com/phd-kerger/gbfs-go-adapter/common"
)

// Describes the capacity and rental availability of the station
type VehicleAvailability struct {
	vehicle_availability_v31_rc2.VehicleAvailability
	Data Data `json:"data"`
}


// Array that contains one object per station as defined below.
type Data struct {
	Vehicles []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	vehicle_availability_v31_rc2.Vehicle
// Unique identifier of a vehicle type as defined in vehicle_types.json.
	VehicleTypeID *common.ID `json:"vehicle_type_id,omitempty"`

}