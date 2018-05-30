package types

type AlexaRequest struct {
	Version string `json:"version"`
	Session `json:"session"`
	Context `json:"context"`
	Request `json:"request"`
}

type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        `json:"user"`
}

type Context struct {
	System      `json:"System"`
	AudioPlayer `json:"AudioPlayer"`
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type User struct {
	UserID      string `json:"userId"`
	Permissions `json:"permissions"`
	AccessToken string `json:"accessToken"`
}

type Request struct {
	Type        string  `json:"type"`
	RequestID   string  `json:"requestId"`
	Timestamp   string  `json:"timestamp"`
	DialogState string  `json:"dialogState"`
	Locale      string  `json:"locale"`
	Intent      *Intent `json:"intent,omitempty"`
}

type Intent struct {
	Name               string          `json:"name,omitempty"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitempty"`
}

type Slot struct {
	Name               string `json:"name,omitempty"`
	Value              string `json:"value,omitempty"`
	ConfirmationStatus string `json:"confirmationStatus,omitempty"`
	Resolutions        `json:"resolutions,omitempty"`
}

type Resolutions struct {
	ResolutionsPerAuthority []Resolution `json:"resolutionsPerAuthority,omitempty"`
}

type Resolution struct {
	Authority string `json:"authority,omitempty"`
	Status    `json:"status,omitempty"`
	Values    []ValueWrapper `json:"values,omitempty"`
}

type Status struct {
	Code string `json:"code,omitempty"`
}

type ValueWrapper struct {
	Value `json:"value,omitempty"`
}

type Value struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

type Permissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

type Device struct {
	DeviceID            string `json:"deviceId,omitempty"`
	SupportedInterfaces `json:"supportedInterfaces,omitempty"`
}

type SupportedInterfaces struct {
	AudioPlayer `json:"AudioPlayer,omitempty"`
}

type System struct {
	Application `json:"application,omitempty"`
	User        `json:"user,omitempty"`
	Device      `json:"device,omitempty"`
	APIEndpoint string `json:"apiEndpoint,omitempty"`
}

type AudioPlayer struct {
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
	PlayerActivity       string `json:"playerActivity,omitempty"`
}
