package dto

type SubmitCodePayload struct {
	LanguageId string `json:"language_id"`
	ProblemId  string `json:"problem_id"`
	SourceCode string `json:"source_code"`
}

type HandshakePayload struct {
	JwtToken string `json:"jwt_token"`
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
