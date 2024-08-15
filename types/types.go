package types

type User struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	DomainID    string `json:"domainId,omitempty"`
	Type        Type   `json:"type,omitempty"`
	IsAnonymous bool   `json:"isAnonymous,omitempty"`
}

type Type string

const (
	TypeHuman Type = "HUMAN"
	TypeBot   Type = "BOT"
)

type DriveDataRef struct {
	DriveFileId string `json:"driveFileId,omitempty"`
}

type Emoji struct {
	Unicode     string       `json:"unicode,omitempty"`
	CustomEmoji *CustomEmoji `json:"customEmoji,omitempty"`
}

type CustomEmoji struct {
	UID string `json:"uid,omitempty"`
}
