package blablacar

import (
	"context"
	"fmt"
	"time"
)

type GetDetailsResponse struct {
	Links                  Links       `json:"links"`
	DepartureDate          string      `json:"departure_date"`
	DepartureDateIso8601   time.Time   `json:"departure_date_iso8601"`
	IsPassed               bool        `json:"is_passed"`
	DeparturePlace         Place       `json:"departure_place"`
	ArrivalPlace           Place       `json:"arrival_place"`
	Price                  Price       `json:"price"`
	PriceWithCommission    Price       `json:"price_with_commission"`
	PriceWithoutCommission Price       `json:"price_without_commission"`
	Commission             Price       `json:"commission"`
	SeatsLeft              int         `json:"seats_left"`
	Seats                  int         `json:"seats"`
	SeatsCountOrigin       int         `json:"seats_count_origin"`
	Duration               Value       `json:"duration"`
	Distance               Value       `json:"distance"`
	PermanentID            string      `json:"permanent_id"`
	MainPermanentID        string      `json:"main_permanent_id"`
	CorridoringID          string      `json:"corridoring_id"`
	TripOfferEncryptedID   string      `json:"trip_offer_encrypted_id"`
	Comment                string      `json:"comment"`
	Car                    Car         `json:"car"`
	ViaggioRosa            bool        `json:"viaggio_rosa"`
	IsComfort              bool        `json:"is_comfort"`
	Freeway                bool        `json:"freeway"`
	StopOvers              []Place     `json:"stop_overs"`
	BucketingEligible      bool        `json:"bucketing_eligible"`
	BookingMode            string      `json:"booking_mode"`
	BookingType            string      `json:"booking_type"`
	IsBookingAllowed       bool        `json:"is_booking_allowed"`
	ViewCount              int         `json:"view_count"`
	CrossBorderAlert       bool        `json:"cross_border_alert"`
	TripPlan               []Place     `json:"trip_plan"`
	MessagingStatus        string      `json:"messaging_status"`
	Passengers             []Passenger `json:"passengers"`
	DisplayContact         bool        `json:"display_contact"`
	//VehiclePictures             []interface{} `json:"vehicle_pictures"`
	CanReport                   bool `json:"can_report"`
	IsArchived                  bool `json:"is_archived"`
	IsTooLateToBook             bool `json:"is_too_late_to_book"`
	IsBookingRefusedByDriver    bool `json:"is_booking_refused_by_driver"`
	HasAlreadyBookedAnotherRide bool `json:"has_already_booked_another_ride"`
	IsWrongGender               bool `json:"is_wrong_gender"`
}

func (c *Client) GetDetails(ctx context.Context, tripID, locale string) (*GetDetailsResponse, error) {
	u := fmt.Sprintf("api/v2/trips/%s?locale=%s", tripID, locale)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	response := GetDetailsResponse{}

	_, err = c.Do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
