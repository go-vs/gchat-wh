package messages

import (
	"github.com/go-vs/gchat-wh/cards"
	"github.com/go-vs/gchat-wh/spaces"
	"github.com/go-vs/gchat-wh/types"
)

type Message struct {
	Name                    string                 `json:"name,omitempty"`
	Sender                  *types.User            `json:"sender,omitempty"`
	CreateTime              string                 `json:"createTime,omitempty"`
	LastUpdateTime          string                 `json:"lastUpdateTime,omitempty"`
	DeleteTime              string                 `json:"deleteTime,omitempty"`
	Text                    string                 `json:"text,omitempty"`
	FormattedText           string                 `json:"formattedText,omitempty"`
	CardsV2                 []CardWithID           `json:"cardsV2,omitempty"`
	Annotations             []Annotation           `json:"annotations,omitempty"`
	Thread                  *Thread                `json:"thread,omitempty"`
	Space                   *spaces.Space          `json:"space,omitempty"`
	FallbackText            string                 `json:"fallbackText,omitempty"`
	ActionResponse          *ActionResponse        `json:"actionResponse,omitempty"`
	ArgumentText            string                 `json:"argumentText,omitempty"`
	SlashCommand            *SlashCommand          `json:"slashCommand,omitempty"`
	Attachment              []Attachment           `json:"attachment,omitempty"`
	MatchedURL              *MatchedURL            `json:"matchedURL,omitempty"`
	ThreadReply             bool                   `json:"threadReply,omitempty"`
	ClientAssignedMessageID string                 `json:"clientAssignedMessageId,omitempty"`
	EmojiReactionSummaries  []EmojiReactionSummary `json:"emojiReactionSummaries,omitempty"`
	PrivateMessageViewer    *types.User            `json:"privateMessageViewer,omitempty"`
	DeletionMetadata        *DeletionMetadata      `json:"deletionMetadata,omitempty"`
	QuotedMessageMetadata   *QuotedMessageMetadata `json:"quotedMessageMetadata,omitempty"`
	AttachedGifs            []AttachedGif          `json:"attachedGifs,omitempty"`
	AccessoryWidgets        []AccessoryWidget      `json:"accessoryWidgets,omitempty"`
}

type CardWithID struct {
	CardID string      `json:"cardId,omitempty"`
	Card   *cards.Card `json:"card,omitempty"`
}

type Annotation struct {
	Type             AnnotationType        `json:"type,omitempty"`
	Length           int                   `json:"length,omitempty"`
	StartIndex       int                   `json:"startIndex,omitempty"`
	UserMention      *UserMentionMetadata  `json:"userMention,omitempty"`
	SlashCommand     *SlashCommandMetadata `json:"slashCommand,omitempty"`
	RichLinkMetadata *RichLinkMetadata     `json:"richLink,omitempty"`
}

type AnnotationType string

const (
	AnnotationUserMention  AnnotationType = "USER_MENTION"
	AnnotationSlashCommand AnnotationType = "SLASH_COMMAND"
	AnnotationRichLink     AnnotationType = "RICH_LINK"
)

type UserMentionMetadata struct {
	User *types.User `json:"user,omitempty"`
	Type Type        `json:"type,omitempty"`
}

type Type string

const (
	TypeAdd     Type = "ADD"
	TypeMention Type = "MENTION"
	TypeInvoke  Type = "INVOKE"
)

type SlashCommandMetadata struct {
	Bot            *types.User `json:"bot,omitempty"`
	Type           Type        `json:"type,omitempty"`
	CommandName    string      `json:"commandName,omitempty"`
	CommandID      string      `json:"commandId,omitempty"`
	TriggersDialog bool        `json:"triggersDialog,omitempty"`
}

type RichLinkMetadata struct {
	URI           string         `json:"uri,omitempty"`
	RichLinkType  RichLinkType   `json:"richLinkType,omitempty"`
	DriveLinkData *DriveLinkData `json:"driveLinkData,omitempty"`
}

type RichLinkType string

const (
	RichLinkDriveFile RichLinkType = "DRIVE_FILE"
)

type DriveLinkData struct {
	DriveDataRef *types.DriveDataRef `json:"driveDataRef,omitempty"`
	MimeType     string              `json:"mimeType,omitempty"`
}

type Thread struct {
	Name      string `json:"name,omitempty"`
	ThreadKey string `json:"threadKey,omitempty"`
}

type ActionResponse struct {
	Type          ResponseType   `json:"type,omitempty"`
	URL           string         `json:"url,omitempty"`
	DialogAction  *DialogAction  `json:"dialogAction,omitempty"`
	UpdatedWidget *UpdatedWidget `json:"updatedWidget,omitempty"`
}

type ResponseType string

const (
	ResponseNewMessage             ResponseType = "NEW_MESSAGE"
	ResponseUpdateMessage          ResponseType = "UPDATE_MESSAGE"
	ResponseUpdateUserMessageCards ResponseType = "UPDATE_USER_MESSAGE_CARDS"
	ResponseRequestConfig          ResponseType = "REQUEST_CONFIG"
	ResponseDialog                 ResponseType = "DIALOG"
	ResponseUpdateWidget           ResponseType = "UPDATE_WIDGET"
)

type DialogAction struct {
	ActionStatus *ActionStatus `json:"actionStatus,omitempty"`
	Dialog       *Dialog       `json:"dialog,omitempty"`
}

type Dialog struct {
	Body *cards.Card `json:"body,omitempty"`
}

type ActionStatus struct {
	StatusCode        Code   `json:"statusCode,omitempty"`
	UserFacingMessage string `json:"userFacingMessage,omitempty"`
}

type Code string

const (
	CodeOk                 Code = "OK"
	CodeCancelled          Code = "CANCELLED"
	CodeUnknown            Code = "UNKNOWN"
	CodeInValidArgument    Code = "INVALID_ARGUMENT"
	CodeDeadlineExceeded   Code = "DEADLINE_EXCEEDED"
	CodeNotFound           Code = "NOT_FOUND"
	CodeAlreadyExists      Code = "ALREADY_EXISTS"
	CodePermissionDenied   Code = "PERMISSION_DENIED"
	CodeUnauthenticated    Code = "UNAUTHENTICATED"
	CodeResourceExhausted  Code = "RESOURCE_EXHAUSTED"
	CodeFailedPrecondition Code = "FAILED_PRECONDITION"
	CodeAborted            Code = "ABORTED"
	CodeOutOfRange         Code = "OUT_OF_RANGE"
	CodeUnimplemented      Code = "UNIMPLEMENTED"
	CodeInternal           Code = "INTERNAL"
	CodeUnavailable        Code = "UNAVAILABLE"
	CodeDataLoss           Code = "DATA_LOSS"
)

type UpdatedWidget struct {
	Widget      string          `json:"widget,omitempty"`
	Suggestions *SelectionItems `json:"suggestions,omitempty"`
}

type SelectionItems struct {
	Items []cards.SelectionItem `json:"items,omitempty"`
}

type SlashCommand struct {
	CommandID string `json:"commandId,omitempty"`
}

type MatchedURL struct {
	URL string `json:"url,omitempty"`
}

type EmojiReactionSummary struct {
	Emoji         *types.Emoji `json:"emoji,omitempty"`
	ReactionCount int          `json:"reactionCount,omitempty"`
}

type DeletionMetadata struct {
	DeletionType DeletionType `json:"deletionType,omitempty"`
}

type DeletionType string

const (
	DeletionCreator          DeletionType = "CREATOR"
	DeletionSpaceOwner       DeletionType = "SPACE_OWNER"
	DeletionAdmin            DeletionType = "ADMIN"
	DeletionAppMessageExpiry DeletionType = "APP_MESSAGE_EXPIRY"
	DeletionCreatorViaApp    DeletionType = "CREATOR_VIA_APP"
	DeletionSpaceOwnerViaApp DeletionType = "SPACE_OWNER_VIA_APP"
)

type QuotedMessageMetadata struct {
	Name           string `json:"name,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
}

type AttachedGif struct {
	URI string `json:"uri,omitempty"`
}

type AccessoryWidget struct {
	ButtonList *cards.ButtonList `json:"buttonList,omitempty"`
}
