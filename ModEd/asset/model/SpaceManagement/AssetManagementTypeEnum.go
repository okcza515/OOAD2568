// MEP-1013

package spacemanagement

type AssetManagementTypeEnum string

const(
	INSTRUMENT_MANAGEMENT AssetManagementTypeEnum = "Instrument Management"
	SUPPLY_MANAGEMENT AssetManagementTypeEnum = "Supply Management"
)

func (r AssetManagementTypeEnum) TypeAssetManagementString() string {
	switch r {
	case INSTRUMENT_MANAGEMENT:
		return "Instrument"
	case SUPPLY_MANAGEMENT:
		return "Supply"
	default:
		return string(r)
	}
}