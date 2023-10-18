package dto

type SubmitCodePayload struct {
	LanguageId string `json:"languageId"`
	ProblemId  string `json:"problemId"`
	SourceCode string `json:"sourceCode"`
}

type HandshakePayload struct {
	JwtToken string `json:"jwtToken"`
}

type PayloadData struct {
	SubmitCodePayload
	HandshakePayload
}

type ClientWebsocketMessage struct {
	Type    string      `json:"type"`
	Payload PayloadData `json:"payload"`
}

const (
	HandshakeMessageType string = "handshake"
	SubmitMessageType    string = "submit"
)
