package types

// SignagePointEvent is the data received for each signage point
type SignagePointEvent struct {
	Success    bool   `json:"success"`
	CC         string `json:"cc"`
	RC         string `json:"rc"`
	SPIndex    int    `json:"sp_index"`
	SPsSubSlot int    `json:"sps_sub_slot"`
}
