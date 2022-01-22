//Package input of Thunder
package input

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

/*
*	Public Functions
 */

// DataOutput to be converted into from DataInput, saved and later presented
type DataOutput struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Addon          string        `json:"addon"`
	UserID         int           `json:"userid"`
	OfferID        int           `json:"offerId"`
	DateStart      int           `json:"date_start"`
	DateEnd        int           `json:"date_end"`
	Timeframe      int           `json:"timeframe"`
	TimeframeUnits string        `json:"timeframe_units"`
}

// OutputData function that converts an element of the type DataInput
// into an element of the type DataOutput to reply to
// user GET requests.
func OutputData(element *DataInput) *DataOutput {

	dataOutput := new(DataOutput)

	dataOutput.ID = element.ID //equals the ObjectId to that of the DataInput version
	dataOutput.Addon = "thunderstorm_bump_reference"
	dataOutput.UserID = element.UserID
	dataOutput.OfferID = element.OfferID
	dataOutput.DateStart = convertDateToTimestamp(element.DatePurchased)
	dataOutput.DateEnd = element.DateEnded
	dataOutput.Timeframe, dataOutput.TimeframeUnits = convertTimeFrame(element.PromotionToolTimeframe)

	return dataOutput
}

/*
*	Private Functions
 */

// convertDateToTimestamp function that receives a date and parses it's format
// to return a unix timestamp of the date.
func convertDateToTimestamp(date string) int {
	format := "2006-01-02 15:04:05"

	t, err := time.Parse(format, date)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t.Unix())
		return int(t.Unix())
	}
	return 0
}

// conertTimeFrame function that receives the field 'promotion_tool_timeframe' and
// extracts the days of the promotion and the time unit at the end.
// It then returns two strings, one with the numbered value and the
// other with the unit.
func convertTimeFrame(promotionToolTimeframe string) (int, string) {

	stringTimeframe := strings.Split(promotionToolTimeframe, "_")

	timeframe := stringTimeframe[0]      // the value number
	timeframeUnits := stringTimeframe[2] // the unit e.g: days, etc

	convTimeframe, err := strconv.Atoi(timeframe)
	if err != nil {
		fmt.Println(err)
	}

	return convTimeframe, timeframeUnits
}
