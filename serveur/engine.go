package engine

// Je crée ma structure
type Engine struct {
	Artists   interface{} `json:"artists"`
	Locations interface{} `json:"locations"`
	Dates     interface{} `json:"dates"`
	Relation  interface{} `json:"relation"`
}
