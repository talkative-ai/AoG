package aog

import (
	"fmt"
)

type ResponseType string

const (
	ResponseTypePermission      ResponseType = "actions.intent.PERMISSIONS"
	ResponseTypeOption          ResponseType = "actions.intent.OPTION"
	ResponseTypeDateTime        ResponseType = "actions.intent.DATETIME"
	ResponseTypeSignIn          ResponseType = "actions.intent.SIGN_IN"
	ResponseTypePlace           ResponseType = "actions.intent.PLACE"
	ResponseTypeDeliveryAddress ResponseType = "actions.intent.DELIVERY_ADDRESS"
	ResponseTypeConfirmation    ResponseType = "actions.intent.CONFIRMATION"
	ResponseTypeLink            ResponseType = "actions.intent.LINK"
)

type Response struct {
	ConversationToken  string                 `json:"conversationToken"`
	ExpectUserResponse bool                   `json:"expectUserResponse"`
	ExpectedInputs     []ExpectedInput        `json:"expectedInputs"`
	ResponseMetadata   map[string]interface{} `json:"responseMetadata"`
}

func NewResponse(conversationToken, ssml, displayText string, expectUserResponse bool) (r Response) {
	r.ConversationToken = conversationToken
	r.ExpectUserResponse = expectUserResponse
	expectedInput := ExpectedInput{}
	expectedInput.AddSimpleResponse(ssml, displayText)
	expectedInput.PossibleIntents = append(expectedInput.PossibleIntents, map[string]string{
		"intent": "assistant.intent.action.TEXT",
	})
	expectedInput.SpeechBiasingHints = []string{}
	r.ExpectedInputs = append(r.ExpectedInputs, expectedInput)
	r.ResponseMetadata = map[string]interface{}{}
	r.ResponseMetadata["status"] = map[string]string{
		"message": "Success (200)",
	}
	return
}

type ExpectedInput struct {
	InputPrompt struct {
		RichInitialPrompt struct {
			Items []interface{} `json:"items"`
		} `json:"richInitialPrompt"`
	} `json:"inputPrompt"`
	PossibleIntents    []map[string]string `json:"possibleIntents"`
	SpeechBiasingHints []string            `json:"speechBiasingHints"`
}

var ErrTooManySimpleResponses error = fmt.Errorf("Too many simple responses. Maximum 2 allowed.")

func (i *ExpectedInput) AddSimpleResponse(ssml, displayText string) error {

	simpleResponseCount := 0

	for _, item := range i.InputPrompt.RichInitialPrompt.Items {
		for key := range item.(map[string]interface{}) {
			if key == "simpleResponse" {
				simpleResponseCount++
			}
			if simpleResponseCount > 2 {
				return ErrTooManySimpleResponses
			}
		}
	}

	item := SimpleResponse{
		SSML:        ssml,
		DisplayText: displayText,
	}
	i.InputPrompt.RichInitialPrompt.Items = append(i.InputPrompt.RichInitialPrompt.Items, map[string]SimpleResponse{
		"simpleResponse": item,
	})
	return nil
}
func (i *ExpectedInput) AddBasicCardResponse() {
	// TODO: Implement
}
func (i *ExpectedInput) AddStructuredResponse() {
	// TODO: Implement
}
func (i *ExpectedInput) AddMediaResponse() {
	// TODO: Implement
}

// Item is used to compose a list of UI elements in ExpectedInput which compose the response
// The items must meet the following requirements:
// 1. The first item must be a google.actions.v2.SimpleResponse
// 2. At most two google.actions.v2.SimpleResponse
// 3. At most one card (e.g. google.actions.v2.ui_elements.BasicCard or google.actions.v2.StructuredResponse or google.actions.v2.MediaResponse [google.actions.v2.ImmersiveResponse]
// 4. Cards may not be used if an actions.intent.OPTION intent is used ie google.actions.v2.ui_elements.ListSelect or google.actions.v2.ui_elements.CarouselSelect
// https://developers.google.com/actions/reference/rest/Shared.Types/AppResponse
type SimpleResponse struct {
	SSML        string `json:"ssml"`
	DisplayText string `json:"displayText"`
}
