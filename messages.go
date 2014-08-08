package twilio

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

// Messages is an intermediate class that will allow the creation of lists of
// recipients/messages
type Messages struct {
	client *TwilioRestClient
}

// MessagesResponse is the struct where we will Unmarshal the API response
type MessagesResponse struct {
	AccountSid  string `json:"account_sid,string"`
	APIVersion  string `json:"api_version,string"`
	Body        string
	NumSegments string `json:"num_segments,string"`
	NumMedia    string `json:"num_media,string"`
	DateCreated string `json:"date_created,string"`
	DateSent    string `json:"date_sent,string"`
	DateUpdated string `json:"date_updated,string"`
	Direciton   string
	From        string
	Prices      string
	Sid         string
	Status      string
	To          string
	Uri         string
}

// getResponseStruct will try to unmarshal the json body into a response
// struct, if it's not possible it will create an ErrorResponse struct
func (m *Messages) getResponseStruct(body io.Reader) (*MessagesResponse, error) {
	messagesResponse := MessagesResponse{}
	bodyContent, _ := ioutil.ReadAll(body)
	if err := json.Unmarshal(bodyContent, &messagesResponse); err != nil {
		return nil, err
	}
	return &messagesResponse, nil
}

// Create a new SMS
func (m *Messages) Create(from string, to string, body string) (*MessagesResponse, *ErrorResponse) {
	apiURL := fmt.Sprintf("%s/Messages.json", m.client.getAPIBaseURL())
	values := url.Values{"Body": {body}, "From": {from}, "To": {to}}

	response, err := m.client.post(apiURL, values)
	if err != nil {
		log.Print(err)
	}

	if response.StatusCode == 200 {
		messagesResponse, err := m.getResponseStruct(response.Body)
		if err != nil {
			log.Print(err)
		}
		return messagesResponse, nil
	}

	return nil, NewErrorResponse(response.Body)
}
