package engine

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

type Engine struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []Location
	ConcertDates dates
	Relations    relations
}

type list struct {
	Lists []Engine
}
type dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}
