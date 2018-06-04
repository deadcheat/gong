package values

// const value of request.type
const (
	// TypeLaunchRequest const value of request.type
	TypeLaunchRequest = "LaunchRequest"
	// TypeIntentRequest const value of request.type
	TypeIntentRequest = "IntentRequest"
	// TypeSessionEndedRequest const value of request.type
	TypeSessionEndedRequest = "SessionEndedRequest"
)

// const value of request.intent.name
const (
	// IntentGongRing const value of request.intent.name
	IntentGongRing           = "GongRing"
	IntentAMAZONCancelIntent = "AMAZON.CancelIntent"
	IntentAMAZONHelpIntent   = "AMAZON.HelpIntent"
	IntentAMAZONStopIntent   = "AMAZON.StopIntent"
)

const (
	TypePlainText = "PlainText"
	TypeSSML      = "SSML"
)

const (
	EnvKeyURL = "GONG_MP3_URL"
)
