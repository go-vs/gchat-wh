package spaces

type Space struct {
	Name                         string                       `json:"name,omitempty"`
	SpaceType                    Type                         `json:"spaceType,omitempty"`
	SingleUserBotDM              bool                         `json:"singleUserBotDm,omitempty"`
	DisplayName                  string                       `json:"displayName,omitempty"`
	ExternalUserAllowed          bool                         `json:"externalUserAllowed,omitempty"`
	SpaceThreadingState          SpaceThreadingState          `json:"spaceThreadingState,omitempty"`
	SpaceDetails                 *SpaceDetails                `json:"spaceDetails,omitempty"`
	SpaceHistoryState            HistoryState                 `json:"spaceHistoryState,omitempty"`
	ImportMode                   bool                         `json:"importMode,omitempty"`
	CreateTime                   string                       `json:"createTime,omitempty"`
	LastActiveTime               string                       `json:"lastActiveTime,omitempty"`
	AdminInstalled               bool                         `json:"adminInstalled,omitempty"`
	MembershipCount              *MembershipCount             `json:"membershipCount,omitempty"`
	AccessSettings               *AccessSettings              `json:"accessSettings,omitempty"`
	SpaceURI                     string                       `json:"spaceUri,omitempty"`
	PredefinedPermissionSettings PredefinedPermissionSettings `json:"predefinedPermissionSettings,omitempty"`
	PermissionSettings           *PermissionSettings          `json:"permissionSettings,omitempty"`
}

type Type string

const (
	SpaceSpace         Type = "SPACE"
	SpaceGroupChat     Type = "GROUP_CHAT"
	SpaceDirectMessage Type = "DIRECT_MESSAGE"
)

type SpaceThreadingState string

const (
	SpaceThreadingThreadedMessages   SpaceThreadingState = "THREADED_MESSAGES"
	SpaceThreadingGroupedMessages    SpaceThreadingState = "GROUPED_MESSAGES"
	SpaceThreadingUnthreadedMessages SpaceThreadingState = "UNTHREADED_MESSAGES"
)

type SpaceDetails struct {
	Description string `json:"description,omitempty"`
	Guidelines  string `json:"guidelines,omitempty"`
}

type HistoryState string

const (
	HistoryOff HistoryState = "HISTORY_OFF"
	HistoryOn  HistoryState = "HISTORY_ON"
)

type MembershipCount struct {
	JoinedDirectHumanUserCount int `json:"joinedDirectHumanUserCount,omitempty"`
	JoinedGroupCount           int `json:"joinedGroupCount,omitempty"`
}

type AccessSettings struct {
	AccessState AccessState `json:"accessState,omitempty"`
	Audience    string      `json:"audience,omitempty"`
}

type AccessState string

const (
	AccessPrivate      AccessState = "PRIVATE"
	AccessDiscoverable AccessState = "DISCOVERABLE"
)

type PredefinedPermissionSettings string

const (
	PredefinedPermissionCollaborationSpace PredefinedPermissionSettings = "COLLABORATION_SPACE"
	PredefinedPermissionAnnouncementSpace  PredefinedPermissionSettings = "ANNOUNCEMENT_SPACE"
)

type PermissionSettings struct {
	ManageMembersAndGroups *PermissionSetting `json:"manageMembersAndGroups,omitempty"`
	ModifySpaceDetails     *PermissionSetting `json:"modifySpaceDetails,omitempty"`
	ToggleHistory          *PermissionSetting `json:"toggleHistory,omitempty"`
	UseAtMentionAll        *PermissionSetting `json:"useAtMentionAll,omitempty"`
	ManageApps             *PermissionSetting `json:"manageApps,omitempty"`
	ManageWebhooks         *PermissionSetting `json:"manageWebhooks,omitempty"`
	PostMessages           *PermissionSetting `json:"postMessages,omitempty"`
	ReplyMessage           *PermissionSetting `json:"replyMessage,omitempty"`
}
type PermissionSetting struct {
	ManagersAllowed bool `json:"managersAllowed,omitempty"`
	MembersAllowed  bool `json:"membersAllowed,omitempty"`
}
