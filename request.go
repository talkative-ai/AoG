package aog

type Request struct {
	Inputs       []Input      `json:"inputs"`
	User         User         `json:"user"`
	Device       Device       `json:"device"`
	Conversation Conversation `json:"conversation"`
}

type Input struct {
	Arguments []InputArgument `json:"arguments"`
	Intent    string          `json:"intent"`
	RawInputs []RawInput      `json:"rawInputs"`
}

type InputArgument struct {
	RawText   string `json:"rawText"`
	TextValue string `json:"textValue"`
	Name      string `json:"name"`
}

type RawInput struct {
	Query     string      `json:"query"`
	InputType interface{} `json:"inputType"`
}

type User struct {
	UserID      string      `json:"userId"`
	Profile     UserProfile `json:"profile"`
	AccessToken string      `json:"accessToken"`
}

type UserProfile struct {
	DisplayName string `json:"display_name"`
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
}

type Conversation struct {
	ConversationToken string      `json:"conversationToken"`
	ConversationID    string      `json:"conversationId"`
	Type              interface{} `json:"type"`
}

type Device struct {
	Location DeviceLocation `json:"location"`
}
type DeviceLocation struct {
	Coordinates      DeviceLocationCoordinates `json:"coordinates"`
	FormattedAddress string                    `json:"formattedAddress"`
	ZipCode          string                    `json:"zipCode"`
	City             string                    `json:"city"`
}

type DeviceLocationCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

const ConstIntentText = "actions.intent.TEXT"
const ConstInputTypeKeyboard = "KEYBOARD"
const ConstInputArgumentText = "text"
