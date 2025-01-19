package models

type Record struct {
	SiteID                int    `json:"site_id"`
	FixletID              int    `json:"fixlet_id"`
	Name                  string `json:"name"`
	Criticality           string `json:"criticality"`
	RelevantComputerCount int    `json:"relevant_computer_count"`
}
