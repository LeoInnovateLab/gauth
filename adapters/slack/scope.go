package slack

type Scope struct {
	scope       string
	description string
	isDefault   bool
}

func (a *Scope) GetScope() string {
	return a.scope
}

func (a *Scope) IsDefault() bool {
	return a.isDefault
}

func (a *Scope) GetDefaultScopes() []string {
	scopes := []Scope{
		UsersProfileRead,
		UsersRead,
		UsersReadEmail,
		UsersProfileWrite,
		UsersProfileWriteUser,
		UsersWrite,
		Admin,
		AdminAnalyticsRead,
		AdminAppsRead,
		AdminAppsWrite,
		AdminBarriersRead,
		AdminBarriersWrite,
		AdminConversationsRead,
		AdminConversationsWrite,
		AdminInvitesRead,
		AdminInvitesWrite,
		AdminTeamsRead,
		AdminTeamsWrite,
		AdminUsergroupsRead,
		AdminUsergroupsWrite,
		AdminUsersRead,
		AdminUsersWrite,
		AppMentionsRead,
		AuditlogsRead,
		Bot,
		CallsRead,
		CallsWrite,
		ChannelsHistory,
		ChannelsJoin,
		ChannelsManage,
		ChannelsRead,
		ChannelsWrite,
		ChatWrite,
		ChatWriteCustomize,
		ChatWritePublic,
		ChatWriteBot,
		ChatWriteUser,
		Client,
		Commands,
		ConversationsHistory,
		ConversationsRead,
		ConversationsWrite,
		DNDRead,
		DNDWrite,
		DNDWriteUser,
		EmojiRead,
		FilesRead,
		FilesWrite,
		FilesWriteUser,
		GroupsHistory,
		GroupsRead,
		GroupsWrite,
		Identify,
		IdentityAvatar,
		IdentityAvatarReadUser,
		IdentityBasic,
		IdentityEmail,
		IdentityEmailReadUser,
		IdentityTeam,
		IdentityTeamReadUser,
		IdentityReadUser,
		ImHistory,
		ImRead,
		ImWrite,
		IncomingWebhook,
		LinksRead,
		LinksWrite,
		MPIMHistory,
		MPIMRead,
		MPIMWrite,
		None,
		PinsRead,
		PinsWrite,
		Post,
		ReactionsRead,
		ReactionsWrite,
		Read,
		RemindersRead,
		RemindersReadUser,
		RemindersWrite,
		RemindersWriteUser,
		RemoteFilesRead,
		RemoteFilesShare,
		RemoteFilesWrite,
		SearchRead,
		StarsRead,
		StarsWrite,
		TeamRead,
		TokensBasic,
		UsergroupsRead,
		UsergroupsWrite,
		WorkflowStepsExecute,
	}

	var result []string
	for _, scope := range scopes {
		if scope.isDefault {
			result = append(result, scope.scope)
		}
	}

	return result
}

var (
	UsersProfileRead = Scope{
		scope:       "users.profile:read",
		description: "View profile details about people in a workspace",
		isDefault:   true,
	}
	UsersRead = Scope{
		scope:       "users:read",
		description: "View people in a workspace",
		isDefault:   true,
	}
	UsersReadEmail = Scope{
		scope:       "users:read.email",
		description: "View email addresses of people in a workspace",
		isDefault:   true,
	}
	UsersProfileWrite = Scope{
		scope:       "users.profile:write",
		description: "Edit a user’s profile information and status",
		isDefault:   false,
	}
	UsersProfileWriteUser = Scope{
		scope:       "users.profile:write:user",
		description: "Change the user's profile fields",
		isDefault:   false,
	}
	UsersWrite = Scope{
		scope:       "users:write",
		description: "Set presence for your Slack app",
		isDefault:   false,
	}
	Admin = Scope{
		scope:       "admin",
		description: "Administer a workspace",
		isDefault:   false,
	}
	AdminAnalyticsRead = Scope{
		scope:       "admin.analytics:read",
		description: "Access analytics data about the organization",
		isDefault:   false,
	}
	AdminAppsRead = Scope{
		scope:       "admin.apps:read",
		description: "View apps and app requests in a workspace",
		isDefault:   false,
	}
	AdminAppsWrite = Scope{
		scope:       "admin.apps:write",
		description: "Manage apps in a workspace",
		isDefault:   false,
	}
	AdminBarriersRead = Scope{
		scope:       "admin.barriers:read",
		description: "Read information barriers in the organization",
		isDefault:   false,
	}
	AdminBarriersWrite = Scope{
		scope:       "admin.barriers:write",
		description: "Manage information barriers in the organization",
		isDefault:   false,
	}
	AdminConversationsRead = Scope{
		scope:       "admin.conversations:read",
		description: "View the channel’s member list, topic, purpose and channel name",
		isDefault:   false,
	}
	AdminConversationsWrite = Scope{
		scope:       "admin.conversations:write",
		description: "Start a new conversation, modify a conversation and modify channel details",
		isDefault:   false,
	}
	AdminInvitesRead = Scope{
		scope:       "admin.invites:read",
		description: "Gain information about invite requests in a Grid organization",
		isDefault:   false,
	}
	AdminInvitesWrite = Scope{
		scope:       "admin.invites:write",
		description: "Approve or deny invite requests in a Grid organization",
		isDefault:   false,
	}
	AdminTeamsRead = Scope{
		scope:       "admin.teams:read",
		description: "Access information about a workspace",
		isDefault:   false,
	}
	AdminTeamsWrite = Scope{
		scope:       "admin.teams:write",
		description: "Make changes to a workspace",
		isDefault:   false,
	}
	AdminUsergroupsRead = Scope{
		scope:       "admin.usergroups:read",
		description: "Access information about user groups",
		isDefault:   false,
	}
	AdminUsergroupsWrite = Scope{
		scope:       "admin.usergroups:write",
		description: "Make changes to your usergroups",
		isDefault:   false,
	}
	AdminUsersRead = Scope{
		scope:       "admin.users:read",
		description: "Access a workspace’s profile information",
		isDefault:   false,
	}
	AdminUsersWrite = Scope{
		scope:       "admin.users:write",
		description: "Modify account information",
		isDefault:   false,
	}
	AppMentionsRead = Scope{
		scope:       "app_mentions:read",
		description: "View messages that directly mention @your_slack_app in conversations that the app is in",
		isDefault:   false,
	}
	AuditlogsRead = Scope{
		scope:       "auditlogs:read",
		description: "View events from all workspaces, channels and users (Enterprise Grid only)",
		isDefault:   false,
	}
	Bot = Scope{
		scope:       "bot",
		description: "Add the ability for people to direct message or mention @your_slack_app",
		isDefault:   false,
	}
	CallsRead = Scope{
		scope:       "calls:read",
		description: "View information about ongoing and past calls",
		isDefault:   false,
	}
	CallsWrite = Scope{
		scope:       "calls:write",
		description: "Start and manage calls in a workspace",
		isDefault:   false,
	}
	ChannelsHistory = Scope{
		scope:       "channels:history",
		description: "View messages and other content in public channels that your slack app has been added to",
		isDefault:   false,
	}
	ChannelsJoin = Scope{
		scope:       "channels:join",
		description: "Join public channels in a workspace",
		isDefault:   false,
	}
	ChannelsManage = Scope{
		scope:       "channels:manage",
		description: "Manage public channels that your slack app has been added to and create new ones",
		isDefault:   false,
	}
	ChannelsRead = Scope{
		scope:       "channels:read",
		description: "View basic information about public channels in a workspace",
		isDefault:   false,
	}
	ChannelsWrite = Scope{
		scope:       "channels:write",
		description: "Manage a user’s public channels and create new ones on a user’s behalf",
		isDefault:   false,
	}
	ChatWrite = Scope{
		scope:       "chat:write",
		description: "Post messages in approved channels & conversations",
		isDefault:   false,
	}
	ChatWriteCustomize = Scope{
		scope:       "chat:write.customize",
		description: "Send messages as @your_slack_app with a customized username and avatar",
		isDefault:   false,
	}
	ChatWritePublic = Scope{
		scope:       "chat:write.public",
		description: "Send messages to channels @your_slack_app isn't a member of",
		isDefault:   false,
	}
	ChatWriteBot = Scope{
		scope:       "chat:write:bot",
		description: "Send messages as your slack app",
		isDefault:   false,
	}
	ChatWriteUser = Scope{
		scope:       "chat:write:user",
		description: "Send messages on a user’s behalf",
		isDefault:   false,
	}
	Client = Scope{
		scope:       "client",
		description: "Receive all events from a workspace in real time",
		isDefault:   false,
	}
	Commands = Scope{
		scope:       "commands",
		description: "Add shortcuts and/or slash commands that people can use",
		isDefault:   false,
	}
	ConversationsHistory = Scope{
		scope:       "conversations:history",
		description: "Deprecated: Retrieve conversation history for legacy workspace apps",
		isDefault:   false,
	}
	ConversationsRead = Scope{
		scope:       "conversations:read",
		description: "Deprecated: Retrieve information on conversations for legacy workspace apps",
		isDefault:   false,
	}
	ConversationsWrite = Scope{
		scope:       "conversations:write",
		description: "Deprecated: Edit conversation attributes for legacy workspace apps",
		isDefault:   false,
	}
	DNDRead = Scope{
		scope:       "dnd:read",
		description: "View Do Not Disturb settings for people in a workspace",
		isDefault:   false,
	}
	DNDWrite = Scope{
		scope:       "dnd:write",
		description: "Edit a user’s Do Not Disturb settings",
		isDefault:   false,
	}
	DNDWriteUser = Scope{
		scope:       "dnd:write:user",
		description: "Change the user's Do Not Disturb settings",
		isDefault:   false,
	}
	EmojiRead = Scope{
		scope:       "emoji:read",
		description: "View custom emoji in a workspace",
		isDefault:   false,
	}
	FilesRead = Scope{
		scope:       "files:read",
		description: "View files shared in channels and conversations that your slack app has been added to",
		isDefault:   false,
	}
	FilesWrite = Scope{
		scope:       "files:write",
		description: "Upload, edit, and delete files as your slack app",
		isDefault:   false,
	}
	FilesWriteUser = Scope{
		scope:       "files:write:user",
		description: "Upload, edit, and delete files as your slack app",
		isDefault:   false,
	}
	GroupsHistory = Scope{
		scope:       "groups:history",
		description: "View messages and other content in private channels that your slack app has been added to",
		isDefault:   false,
	}
	GroupsRead = Scope{
		scope:       "groups:read",
		description: "View basic information about private channels that your slack app has been added to",
		isDefault:   false,
	}
	GroupsWrite = Scope{
		scope:       "groups:write",
		description: "Manage private channels that your slack app has been added to and create new ones",
		isDefault:   false,
	}
	Identify = Scope{
		scope:       "identify",
		description: "View information about a user’s identity",
		isDefault:   false,
	}
	IdentityAvatar = Scope{
		scope:       "identity.avatar",
		description: "View a user’s Slack avatar",
		isDefault:   false,
	}
	IdentityAvatarReadUser = Scope{
		scope:       "identity.avatar:read:user",
		description: "View the user's profile picture",
		isDefault:   false,
	}
	IdentityBasic = Scope{
		scope:       "identity.basic",
		description: "View information about a user’s identity",
		isDefault:   false,
	}
	IdentityEmail = Scope{
		scope:       "identity.email",
		description: "View a user’s email address",
		isDefault:   false,
	}
	IdentityEmailReadUser = Scope{
		scope:       "identity.email:read:user",
		description: "This scope is not yet described.",
		isDefault:   false,
	}
	IdentityTeam = Scope{
		scope:       "identity.team",
		description: "View a user’s Slack workspace name",
		isDefault:   false,
	}
	IdentityTeamReadUser = Scope{
		scope:       "identity.team:read:user",
		description: "View the workspace's name, domain, and icon",
		isDefault:   false,
	}
	IdentityReadUser = Scope{
		scope:       "identity:read:user",
		description: "This scope is not yet described.",
		isDefault:   false,
	}
	ImHistory = Scope{
		scope:       "im:history",
		description: "View messages and other content in direct messages that your slack app has been added to",
		isDefault:   false,
	}
	ImRead = Scope{
		scope:       "im:read",
		description: "View basic information about direct messages that your slack app has been added to",
		isDefault:   false,
	}
	ImWrite = Scope{
		scope:       "im:write",
		description: "Start direct messages with people",
		isDefault:   false,
	}
	IncomingWebhook = Scope{
		scope:       "incoming-webhook",
		description: "Create one-way webhooks to post messages to a specific channel",
		isDefault:   false,
	}
	LinksRead = Scope{
		scope:       "links:read",
		description: "View  URLs in messages",
		isDefault:   false,
	}
	LinksWrite = Scope{
		scope:       "links:write",
		description: "Show previews of  URLs in messages",
		isDefault:   false,
	}
	MPIMHistory = Scope{
		scope:       "mpim:history",
		description: "View messages and other content in group direct messages that your slack app has been added to",
		isDefault:   false,
	}
	MPIMRead = Scope{
		scope:       "mpim:read",
		description: "View basic information about group direct messages that your slack app has been added to",
		isDefault:   false,
	}
	MPIMWrite = Scope{
		scope:       "mpim:write",
		description: "Start group direct messages with people",
		isDefault:   false,
	}
	None = Scope{
		scope:       "none",
		description: "Execute methods without needing a scope",
		isDefault:   false,
	}
	PinsRead = Scope{
		scope:       "pins:read",
		description: "View pinned content in channels and conversations that your slack app has been added to",
		isDefault:   false,
	}
	PinsWrite = Scope{
		scope:       "pins:write",
		description: "Add and remove pinned messages and files",
		isDefault:   false,
	}
	Post = Scope{
		scope:       "post",
		description: "Post messages to a workspace",
		isDefault:   false,
	}
	ReactionsRead = Scope{
		scope:       "reactions:read",
		description: "View emoji reactions and their associated content in channels and conversations that your slack app has been added to",
		isDefault:   false,
	}
	ReactionsWrite = Scope{
		scope:       "reactions:write",
		description: "Add and edit emoji reactions",
		isDefault:   false,
	}
	Read = Scope{
		scope:       "read",
		description: "View all content in a workspace",
		isDefault:   false,
	}
	RemindersRead = Scope{
		scope:       "reminders:read",
		description: "View reminders created by your slack app",
		isDefault:   false,
	}
	RemindersReadUser = Scope{
		scope:       "reminders:read:user",
		description: "Access reminders created by a user or for a user",
		isDefault:   false,
	}
	RemindersWrite = Scope{
		scope:       "reminders:write",
		description: "Add, remove, or mark reminders as complete",
		isDefault:   false,
	}
	RemindersWriteUser = Scope{
		scope:       "reminders:write:user",
		description: "Add, remove, or complete reminders for the user",
		isDefault:   false,
	}
	RemoteFilesRead = Scope{
		scope:       "remote_files:read",
		description: "View remote files added by the app in a workspace",
		isDefault:   false,
	}
	RemoteFilesShare = Scope{
		scope:       "remote_files:share",
		description: "Share remote files on a user’s behalf",
		isDefault:   false,
	}
	RemoteFilesWrite = Scope{
		scope:       "remote_files:write",
		description: "Add, edit, and delete remote files on a user’s behalf",
		isDefault:   false,
	}
	SearchRead = Scope{
		scope:       "search:read",
		description: "Search a workspace’s content",
		isDefault:   false,
	}
	StarsRead = Scope{
		scope:       "stars:read",
		description: "View messages and files that your slack app has starred",
		isDefault:   false,
	}
	StarsWrite = Scope{
		scope:       "stars:write",
		description: "Add or remove stars",
		isDefault:   false,
	}
	TeamRead = Scope{
		scope:       "team:read",
		description: "View the name, email domain, and icon for workspaces your slack app is connected to",
		isDefault:   false,
	}
	TokensBasic = Scope{
		scope:       "tokens.basic",
		description: "Execute methods without needing a scope",
		isDefault:   false,
	}
	UsergroupsRead = Scope{
		scope:       "usergroups:read",
		description: "View user groups in a workspace",
		isDefault:   false,
	}
	UsergroupsWrite = Scope{
		scope:       "usergroups:write",
		description: "Create and manage user groups",
		isDefault:   false,
	}
	WorkflowStepsExecute = Scope{
		scope:       "workflow.steps:execute",
		description: "Add steps that people can use in Workflow Builder",
		isDefault:   false,
	}
)
