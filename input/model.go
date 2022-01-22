//Package input of Thunder
package input

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"thunder/config"
)

// DataInput struct that holds the information received
type DataInput struct {
	ID                     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID                 int           `json:"userid"`
	OfferID                int           `json:"offerId"`
	DatePurchased          string        `json:"date_purchase"`
	DateEnded              int           `json:"date_end"`
	SenderPhone            int           `json:"senderPhone"`
	PromotionToolBough     string        `json:"promotion_tool_bough"`
	PromotionToolTimeframe string        `json:"promotion_tool_timeframe"`
	Email                  string        `json:"email"`
	Message                string        `json:"message"`
}

/*
*	Public Functions
 */

// Insert function that Inserts a new DataInput in the mongo db.
// This is stored but later converted into a DataOutput with
// the information kept in it.
func Insert(e *DataInput) error {
	collection := config.GetCollection("datainput")

	err := collection.Insert(e)
	if err != nil {
		log.Fatal(err)
	}

	showConsoleInfoInput("insert", e)
	return err
}

// FindOne function that finds an DataInput in the mongo db via its ObjectId.
// It then searches the DataOutput collection to see if there exists
// a converted value already. If not, it converts the DataInput into
// a DataOutput and returns that output to the user. If it already exists,
// then it just searches for it via its ObjectId and returns it.
func FindOne(id string) (*DataOutput, error) {
	objectID := bson.ObjectIdHex(id)

	collection := config.GetCollection("datainput")

	e := new(DataInput)
	query := bson.M{"_id": objectID}
	err := collection.Find(query).One(&e)
	if err != nil {
		return nil, err
	}

	collection = config.GetCollection("dataoutput")
	output := new(DataOutput)
	err = collection.Find(query).One(output)
	if err != nil {
		output = OutputData(e)
		InsertOutput(output)
	}

	showConsoleInfoOutput("get", output)
	return output, nil
}

// UpdateOne function that updates the contents of a DataInput with the specified ObjectId.
// It also searches for the converted value DataOutput and changes the values
// that correspond to them, if they need updating.
func UpdateOne(id string, e *DataInput) (*DataInput, error) {
	objectID := bson.ObjectIdHex(id)

	collection := config.GetCollection("datainput")

	query := bson.M{"_id": objectID}
	err := collection.Update(query, &e)
	if err != nil {
		return nil, err
	}

	collection = config.GetCollection("dataoutput")

	query = bson.M{"_id": objectID}
	err = collection.Update(query, OutputData(e))
	if err != nil {
		return nil, err
	}

	showConsoleInfoInput("update", e)
	return e, nil
}

// InsertOutput function that inserts a new DataOutput in the mongo db.
// This function runs inside Insert() after a conversion of
// DataInput was made into DataOutput.
// This functions also assigns the ObjectId of the Output the same as the Input
// so there is a relation between them.
func InsertOutput(e *DataOutput) {
	collection := config.GetCollection("dataoutput")

	err := collection.Insert(e)
	if err != nil {
		log.Fatal(err)
	}

}

/*
*	Private Functions
 */

// showConsoleInfoInput function that displays information about the DataInput on the console.
// It reads a string passed through as an argument and indicates the action
// requested by the user.
func showConsoleInfoInput(action string, e *DataInput) {
	if e != nil {
		switch action {
		case "insert":
			fmt.Println("\n>> Inserted DataInput <<")
			break
		case "update":
			fmt.Println("\n>> Updated DataInput <<")
			break
		default:
			break
		}
		fmt.Printf("userid: %d\n", e.UserID)
		fmt.Printf("offerId: %d\n", e.OfferID)
		fmt.Printf("date_purchase: %s\n", e.DatePurchased)
		fmt.Printf("date_end: %d\n", e.DateEnded)
		fmt.Printf("senderPhone: %d\n", e.SenderPhone)
		fmt.Printf("promotion_tool_bough: %s\n", e.PromotionToolBough)
		fmt.Printf("promotion_tool_timeframe: %s\n", e.PromotionToolTimeframe)
		fmt.Printf("email: %s\n", e.Email)
		fmt.Printf("message: %s\n", e.Message)
	} else {
		fmt.Printf("DataInput came empty or an error was found\n")
	}
}

// showConsoleInfoOutput function that displays information about the DataOutput on the console.
// It reads a string passed through as an argument and indicates the action
// requested by the user.
func showConsoleInfoOutput(action string, e *DataOutput) {
	if e != nil {
		switch action {
		case "get":
			fmt.Println("\n>> Found DataOutput <<")
			break
		default:
			break
		}
		fmt.Printf("addon: %s\n", e.Addon)
		fmt.Printf("offerId: %d\n", e.UserID)
		fmt.Printf("offerId: %d\n", e.OfferID)
		fmt.Printf("date_start: %d\n", e.DateStart)
		fmt.Printf("date_end: %d\n", e.DateEnd)
		fmt.Printf("timeframe: %d\n", e.Timeframe)
		fmt.Printf("timeframe_units: %s\n", e.TimeframeUnits)
	} else {
		fmt.Printf("DataOutput came empty or an error was found\n")
	}
}

// String function that receives a pointer to a DataInput and converts
// it to json format.
func (e *DataInput) String() []byte {
	b, _ := json.Marshal(e)
	return b
}
