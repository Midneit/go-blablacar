package blablacar

// Links struct.
type Links struct {
	Self  string `json:"_self"`
	Front string `json:"_front"`
}

// Place struct.
type Place struct {
	CityName       string  `json:"city_name"`
	Address        string  `json:"address"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	CountryCode    string  `json:"country_code"`
	DeparturePlace bool    `json:"departure_place,omitempty"`
	ArrivalPlace   bool    `json:"arrival_place,omitempty"`
}

// Price struct.
type Price struct {
	Value       float64 `json:"value"`
	Currency    string  `json:"currency"`
	Symbol      string  `json:"symbol"`
	StringValue string  `json:"string_value"`
	PriceColor  string  `json:"price_color"`
}

// Value struct.
type Value struct {
	Value int    `json:"value"`
	Unity string `json:"unity"`
}

// Car struct.
type Car struct {
	ID                      string `json:"id"`
	Model                   string `json:"model"`
	Make                    string `json:"make"`
	Color                   string `json:"color"`
	ColorHexa               string `json:"color_hexa"`
	Comfort                 string `json:"comfort"`
	ComfortNbStar           int    `json:"comfort_nb_star"`
	NumberOfSeat            int    `json:"number_of_seat"`
	Category                string `json:"category"`
	Picture                 string `json:"picture"`
	PictureModerationStatus string `json:"picture_moderation_status"`
}

// Passenger struct.
type Passenger struct {
	Links         Links  `json:"links"`
	HasPicture    bool   `json:"has_picture"`
	EncryptedID   string `json:"encrypted_id"`
	PhoneVerified bool   `json:"phone_verified"`
	PhoneHidden   bool   `json:"phone_hidden"`
	UUID          string `json:"uuid"`
	BookingStatus string `json:"booking_status"`
	Seats         []Seat `json:"seats"`
}

// Seat struct.
type Seat struct {
	DepartureCity string `json:"departure_city"`
	ArrivalCity   string `json:"arrival_city"`
	BookingStatus string `json:"booking_status"`
}

// SearchInfo struct.
type SearchInfo struct {
	Count         int `json:"count"`
	FullTripCount int `json:"full_trip_count"`
}

// WayPoint struct.
type WayPoint struct {
	DateTime string `json:"date_time"`
	Place    Place  `json:"place"`
}

// Vehicle struct.
type Vehicle struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}

// SearchTrip struct.
type SearchTrip struct {
	Link              string     `json:"link"`
	WayPoints         []WayPoint `json:"waypoints"`
	Price             Price      `json:"price"`
	Vehicle           Vehicle    `json:"vehicle"`
	DistanceInMeters  int        `json:"distance_in_meters"`
	DurationInSeconds int        `json:"duration_in_seconds"`
}
