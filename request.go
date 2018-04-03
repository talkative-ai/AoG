package aog

type Request struct {
	Inputs []struct {
		Arguments []struct {
			RawText   string `json:"rawText"`
			TextValue string `json:"textValue"`
			Name      string `json:"name"`
		} `json:"arguments"`
		Intent    string `json:"intent"`
		RawInputs []struct {
			Query     string      `json:"query"`
			InputType interface{} `json:"inputType"`
		} `json:"rawInputs"`
	} `json:"inputs"`
	User struct {
		UserID  string `json:"userId"`
		Profile struct {
			DisplayName string `json:"display_name"`
			GivenName   string `json:"given_name"`
			FamilyName  string `json:"family_name"`
		} `json:"profile"`
		AccessToken string `json:"accessToken"`
	} `json:"user"`
	Device struct {
		Location struct {
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"coordinates"`
			FormattedAddress string `json:"formattedAddress"`
			ZipCode          string `json:"zipCode"`
			City             string `json:"city"`
		} `json:"location"`
	} `json:"device"`
	Conversation struct {
		ConversationToken string      `json:"conversationToken"`
		ConversationID    string      `json:"conversationId"`
		Type              interface{} `json:"type"`
	} `json:"conversation"`
}
