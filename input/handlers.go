//Package input of Thunder
package input

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

/*
*	Public Functions
 */

// Show function that handles the presentation of an DataInput with the specified id.
// It handles the parameter "_id" and calls validations upon it.
// The id has to be 24 bytes long and exist in the mongo db.
// If all validations pass we send a 200 OK code and present the json.
func Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idObject := p.ByName("id")

	//if the id isn't exactly 24 bytes long, it's not valid and panics when connecting to mongo
	if !bson.IsObjectIdHex(idObject) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "The specified id has an invalid format.")
		return
	}

	//we fetch the element with the specified id and act on it
	element, err := FindOne(idObject)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "No element was found with the specified id.")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", element.String())
}

// Create function that handles the creation of a new DataInput from a body request
// This body has to follow the normal json rules and formatting to be valid
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	element := new(DataInput)
	defer r.Body.Close()

	//if the body is empty
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "The json body is empty.")
		return
	}

	//if the body has an invalid format
	ObjectID := bson.NewObjectId()
	element.ID = ObjectID
	err = json.Unmarshal(body, &element)

	if err != nil || !isJSONCorrect(element) {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			fmt.Fprintf(w, "The json format is invalid,: %v", err)
			return
		}
		fmt.Fprintf(w, "The json format is invalid as one or more values are incorrect.")
		return
	}

	Insert(element)

	fmt.Fprintf(w, "%s", element.String())
}

// Update function that updates the contents of an DataInput via a json object
// in the body and an id specified in the url
func Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idObject := p.ByName("id")

	//if the id isn't exactly 24 bytes long, it's not valid and panics when connecting to mongo
	if !bson.IsObjectIdHex(idObject) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "The specified id has an invalid format.")
		return
	}

	element := new(DataInput)
	defer r.Body.Close()

	//if the body is empty
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "The json body is empty.")
		return
	}

	//if the body has an invalid format
	err = json.Unmarshal(body, &element)

	if err != nil || !isJSONCorrect(element) {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			fmt.Fprintf(w, "The json format is invalid,: %v", err)
			return
		}
		fmt.Fprintf(w, "The json format is invalid as one or more values are incorrect.")
		return
	}

	//we update the element with the specified id and act on it
	elm, err := UpdateOne(idObject, element)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "No element was found with the specified id.")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", elm.String())
}

/*
*	Private Functions
 */

// isJSONCorrect function that checks the values of an DataInput and sees if any is empty.
// If any is empty it returns false as the body wasn't well formed.
// If no value is empty it returns true indicating the body was well formed.
func isJSONCorrect(e *DataInput) bool {

	if e.UserID == 0 {
		return false
	}
	if e.OfferID == 0 {
		return false
	}
	if e.DatePurchased == "" {
		return false
	}
	if e.DateEnded == 0 {
		return false
	}
	if e.SenderPhone == 0 {
		return false
	}
	if e.PromotionToolBough == "" {
		return false
	}
	if e.PromotionToolTimeframe == "" {
		return false
	}
	if e.Email == "" {
		return false
	}
	if e.Message == "" {
		return false
	}

	return true
}

// String function that receives a pointer to a DataInput and converts
// it to json format.
func (e *DataOutput) String() []byte {
	b, _ := json.Marshal(e)
	return b
}