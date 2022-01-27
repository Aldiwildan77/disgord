// Code generated by generate/interfaces; DO NOT EDIT.

package disgord

func (a *Activity) copyOverTo(other interface{}) error {
	var dest *Activity
	var valid bool
	if dest, valid = other.(*Activity); !valid {
		return newErrorUnsupportedType("argument given is not a *Activity type")
	}
	dest.ApplicationID = a.ApplicationID
	dest.Assets = a.Assets
	dest.CreatedAt = a.CreatedAt
	dest.Details = a.Details
	dest.Emoji = a.Emoji
	dest.Flags = a.Flags
	dest.Instance = a.Instance
	dest.Name = a.Name
	dest.Party = a.Party
	dest.Secrets = a.Secrets
	dest.State = a.State
	dest.Timestamps = a.Timestamps
	dest.Type = a.Type
	dest.URL = a.URL

	return nil
}

func (a *ActivityAssets) copyOverTo(other interface{}) error {
	var dest *ActivityAssets
	var valid bool
	if dest, valid = other.(*ActivityAssets); !valid {
		return newErrorUnsupportedType("argument given is not a *ActivityAssets type")
	}
	dest.LargeImage = a.LargeImage
	dest.LargeText = a.LargeText
	dest.SmallImage = a.SmallImage
	dest.SmallText = a.SmallText

	return nil
}

func (a *ActivityEmoji) copyOverTo(other interface{}) error {
	var dest *ActivityEmoji
	var valid bool
	if dest, valid = other.(*ActivityEmoji); !valid {
		return newErrorUnsupportedType("argument given is not a *ActivityEmoji type")
	}
	dest.Animated = a.Animated
	dest.ID = a.ID
	dest.Name = a.Name

	return nil
}

func (a *ActivityParty) copyOverTo(other interface{}) error {
	var dest *ActivityParty
	var valid bool
	if dest, valid = other.(*ActivityParty); !valid {
		return newErrorUnsupportedType("argument given is not a *ActivityParty type")
	}
	dest.ID = a.ID
	dest.Size = make([]int, len(a.Size))
	copy(dest.Size, a.Size)

	return nil
}

func (a *ActivitySecrets) copyOverTo(other interface{}) error {
	var dest *ActivitySecrets
	var valid bool
	if dest, valid = other.(*ActivitySecrets); !valid {
		return newErrorUnsupportedType("argument given is not a *ActivitySecrets type")
	}
	dest.Join = a.Join
	dest.Match = a.Match
	dest.Spectate = a.Spectate

	return nil
}

func (a *ActivityTimestamp) copyOverTo(other interface{}) error {
	var dest *ActivityTimestamp
	var valid bool
	if dest, valid = other.(*ActivityTimestamp); !valid {
		return newErrorUnsupportedType("argument given is not a *ActivityTimestamp type")
	}
	dest.End = a.End
	dest.Start = a.Start

	return nil
}

func (a *Attachment) copyOverTo(other interface{}) error {
	var dest *Attachment
	var valid bool
	if dest, valid = other.(*Attachment); !valid {
		return newErrorUnsupportedType("argument given is not a *Attachment type")
	}
	dest.Filename = a.Filename
	dest.Height = a.Height
	dest.ID = a.ID
	dest.ProxyURL = a.ProxyURL
	dest.Size = a.Size
	dest.SpoilerTag = a.SpoilerTag
	dest.URL = a.URL
	dest.Width = a.Width

	return nil
}

func (a *AuditLog) copyOverTo(other interface{}) error {
	var dest *AuditLog
	var valid bool
	if dest, valid = other.(*AuditLog); !valid {
		return newErrorUnsupportedType("argument given is not a *AuditLog type")
	}
	dest.AuditLogEntries = make([]*AuditLogEntry, len(a.AuditLogEntries))
	for i := 0; i < len(a.AuditLogEntries); i++ {
		dest.AuditLogEntries[i] = DeepCopy(a.AuditLogEntries[i]).(*AuditLogEntry)
	}
	dest.Users = make([]*User, len(a.Users))
	for i := 0; i < len(a.Users); i++ {
		dest.Users[i] = DeepCopy(a.Users[i]).(*User)
	}
	dest.Webhooks = make([]*Webhook, len(a.Webhooks))
	for i := 0; i < len(a.Webhooks); i++ {
		dest.Webhooks[i] = DeepCopy(a.Webhooks[i]).(*Webhook)
	}

	return nil
}

func (a *AuditLogChanges) copyOverTo(other interface{}) error {
	var dest *AuditLogChanges
	var valid bool
	if dest, valid = other.(*AuditLogChanges); !valid {
		return newErrorUnsupportedType("argument given is not a *AuditLogChanges type")
	}
	dest.Key = a.Key
	dest.NewValue = a.NewValue
	dest.OldValue = a.OldValue

	return nil
}

func (a *AuditLogEntry) copyOverTo(other interface{}) error {
	var dest *AuditLogEntry
	var valid bool
	if dest, valid = other.(*AuditLogEntry); !valid {
		return newErrorUnsupportedType("argument given is not a *AuditLogEntry type")
	}
	dest.Changes = make([]*AuditLogChanges, len(a.Changes))
	for i := 0; i < len(a.Changes); i++ {
		dest.Changes[i] = DeepCopy(a.Changes[i]).(*AuditLogChanges)
	}
	dest.Event = a.Event
	dest.ID = a.ID
	dest.Options = a.Options
	dest.Reason = a.Reason
	dest.TargetID = a.TargetID
	dest.UserID = a.UserID

	return nil
}

func (a *AuditLogOption) copyOverTo(other interface{}) error {
	var dest *AuditLogOption
	var valid bool
	if dest, valid = other.(*AuditLogOption); !valid {
		return newErrorUnsupportedType("argument given is not a *AuditLogOption type")
	}
	dest.ChannelID = a.ChannelID
	dest.Count = a.Count
	dest.DeleteMemberDays = a.DeleteMemberDays
	dest.ID = a.ID
	dest.MembersRemoved = a.MembersRemoved
	dest.RoleName = a.RoleName
	dest.Type = a.Type

	return nil
}

func (b *Ban) copyOverTo(other interface{}) error {
	var dest *Ban
	var valid bool
	if dest, valid = other.(*Ban); !valid {
		return newErrorUnsupportedType("argument given is not a *Ban type")
	}
	dest.Reason = b.Reason
	dest.User = b.User

	return nil
}

func (c *Channel) copyOverTo(other interface{}) error {
	var dest *Channel
	var valid bool
	if dest, valid = other.(*Channel); !valid {
		return newErrorUnsupportedType("argument given is not a *Channel type")
	}
	dest.ApplicationID = c.ApplicationID
	dest.Bitrate = c.Bitrate
	dest.DefaultAutoArchiveDuration = c.DefaultAutoArchiveDuration
	dest.GuildID = c.GuildID
	dest.Icon = c.Icon
	dest.ID = c.ID
	dest.LastMessageID = c.LastMessageID
	dest.LastPinTimestamp = c.LastPinTimestamp
	dest.Member = c.Member
	dest.MemberCount = c.MemberCount
	dest.MessageCount = c.MessageCount
	dest.Name = c.Name
	dest.NSFW = c.NSFW
	dest.OwnerID = c.OwnerID
	dest.ParentID = c.ParentID
	dest.PermissionOverwrites = make([]PermissionOverwrite, len(c.PermissionOverwrites))
	copy(dest.PermissionOverwrites, c.PermissionOverwrites)
	dest.Position = c.Position
	dest.RateLimitPerUser = c.RateLimitPerUser
	dest.Recipients = make([]*User, len(c.Recipients))
	for i := 0; i < len(c.Recipients); i++ {
		dest.Recipients[i] = DeepCopy(c.Recipients[i]).(*User)
	}
	dest.ThreadMetadata = c.ThreadMetadata
	dest.Topic = c.Topic
	dest.Type = c.Type
	dest.UserLimit = c.UserLimit

	return nil
}

func (e *Embed) copyOverTo(other interface{}) error {
	var dest *Embed
	var valid bool
	if dest, valid = other.(*Embed); !valid {
		return newErrorUnsupportedType("argument given is not a *Embed type")
	}
	dest.Author = e.Author
	dest.Color = e.Color
	dest.Description = e.Description
	dest.Fields = make([]*EmbedField, len(e.Fields))
	for i := 0; i < len(e.Fields); i++ {
		dest.Fields[i] = DeepCopy(e.Fields[i]).(*EmbedField)
	}
	dest.Footer = e.Footer
	dest.Image = e.Image
	dest.Provider = e.Provider
	dest.Thumbnail = e.Thumbnail
	dest.Timestamp = e.Timestamp
	dest.Title = e.Title
	dest.Type = e.Type
	dest.URL = e.URL
	dest.Video = e.Video

	return nil
}

func (e *EmbedAuthor) copyOverTo(other interface{}) error {
	var dest *EmbedAuthor
	var valid bool
	if dest, valid = other.(*EmbedAuthor); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedAuthor type")
	}
	dest.IconURL = e.IconURL
	dest.Name = e.Name
	dest.ProxyIconURL = e.ProxyIconURL
	dest.URL = e.URL

	return nil
}

func (e *EmbedField) copyOverTo(other interface{}) error {
	var dest *EmbedField
	var valid bool
	if dest, valid = other.(*EmbedField); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedField type")
	}
	dest.Inline = e.Inline
	dest.Name = e.Name
	dest.Value = e.Value

	return nil
}

func (e *EmbedFooter) copyOverTo(other interface{}) error {
	var dest *EmbedFooter
	var valid bool
	if dest, valid = other.(*EmbedFooter); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedFooter type")
	}
	dest.IconURL = e.IconURL
	dest.ProxyIconURL = e.ProxyIconURL
	dest.Text = e.Text

	return nil
}

func (e *EmbedImage) copyOverTo(other interface{}) error {
	var dest *EmbedImage
	var valid bool
	if dest, valid = other.(*EmbedImage); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedImage type")
	}
	dest.Height = e.Height
	dest.ProxyURL = e.ProxyURL
	dest.URL = e.URL
	dest.Width = e.Width

	return nil
}

func (e *EmbedProvider) copyOverTo(other interface{}) error {
	var dest *EmbedProvider
	var valid bool
	if dest, valid = other.(*EmbedProvider); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedProvider type")
	}
	dest.Name = e.Name
	dest.URL = e.URL

	return nil
}

func (e *EmbedThumbnail) copyOverTo(other interface{}) error {
	var dest *EmbedThumbnail
	var valid bool
	if dest, valid = other.(*EmbedThumbnail); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedThumbnail type")
	}
	dest.Height = e.Height
	dest.ProxyURL = e.ProxyURL
	dest.URL = e.URL
	dest.Width = e.Width

	return nil
}

func (e *EmbedVideo) copyOverTo(other interface{}) error {
	var dest *EmbedVideo
	var valid bool
	if dest, valid = other.(*EmbedVideo); !valid {
		return newErrorUnsupportedType("argument given is not a *EmbedVideo type")
	}
	dest.Height = e.Height
	dest.URL = e.URL
	dest.Width = e.Width

	return nil
}

func (e *Emoji) copyOverTo(other interface{}) error {
	var dest *Emoji
	var valid bool
	if dest, valid = other.(*Emoji); !valid {
		return newErrorUnsupportedType("argument given is not a *Emoji type")
	}
	dest.Animated = e.Animated
	dest.Available = e.Available
	dest.ID = e.ID
	dest.Managed = e.Managed
	dest.Name = e.Name
	dest.RequireColons = e.RequireColons
	dest.Roles = make([]Snowflake, len(e.Roles))
	copy(dest.Roles, e.Roles)
	dest.User = e.User

	return nil
}

func (g *Guild) copyOverTo(other interface{}) error {
	var dest *Guild
	var valid bool
	if dest, valid = other.(*Guild); !valid {
		return newErrorUnsupportedType("argument given is not a *Guild type")
	}
	dest.AfkChannelID = g.AfkChannelID
	dest.AfkTimeout = g.AfkTimeout
	dest.ApplicationID = g.ApplicationID
	dest.Banner = g.Banner
	dest.Channels = make([]*Channel, len(g.Channels))
	for i := 0; i < len(g.Channels); i++ {
		dest.Channels[i] = DeepCopy(g.Channels[i]).(*Channel)
	}
	dest.DefaultMessageNotifications = g.DefaultMessageNotifications
	dest.Description = g.Description
	dest.DiscoverySplash = g.DiscoverySplash
	dest.Emojis = make([]*Emoji, len(g.Emojis))
	for i := 0; i < len(g.Emojis); i++ {
		dest.Emojis[i] = DeepCopy(g.Emojis[i]).(*Emoji)
	}
	dest.ExplicitContentFilter = g.ExplicitContentFilter
	dest.Features = make([]string, len(g.Features))
	copy(dest.Features, g.Features)
	dest.Icon = g.Icon
	dest.ID = g.ID
	dest.JoinedAt = g.JoinedAt
	dest.Large = g.Large
	dest.MemberCount = g.MemberCount
	dest.Members = make([]*Member, len(g.Members))
	for i := 0; i < len(g.Members); i++ {
		dest.Members[i] = DeepCopy(g.Members[i]).(*Member)
	}
	dest.MFALevel = g.MFALevel
	dest.Name = g.Name
	dest.Owner = g.Owner
	dest.OwnerID = g.OwnerID
	dest.Permissions = g.Permissions
	dest.PremiumSubscriptionCount = g.PremiumSubscriptionCount
	dest.PremiumTier = g.PremiumTier
	dest.Presences = make([]*UserPresence, len(g.Presences))
	for i := 0; i < len(g.Presences); i++ {
		dest.Presences[i] = DeepCopy(g.Presences[i]).(*UserPresence)
	}
	dest.Region = g.Region
	dest.Roles = make([]*Role, len(g.Roles))
	for i := 0; i < len(g.Roles); i++ {
		dest.Roles[i] = DeepCopy(g.Roles[i]).(*Role)
	}
	dest.Splash = g.Splash
	dest.SystemChannelID = g.SystemChannelID
	dest.Unavailable = g.Unavailable
	dest.VanityUrl = g.VanityUrl
	dest.VerificationLevel = g.VerificationLevel
	dest.VoiceStates = make([]*VoiceState, len(g.VoiceStates))
	for i := 0; i < len(g.VoiceStates); i++ {
		dest.VoiceStates[i] = DeepCopy(g.VoiceStates[i]).(*VoiceState)
	}
	dest.WidgetChannelID = g.WidgetChannelID
	dest.WidgetEnabled = g.WidgetEnabled

	return nil
}

func (g *GuildWidget) copyOverTo(other interface{}) error {
	var dest *GuildWidget
	var valid bool
	if dest, valid = other.(*GuildWidget); !valid {
		return newErrorUnsupportedType("argument given is not a *GuildWidget type")
	}
	dest.ChannelID = g.ChannelID
	dest.Enabled = g.Enabled

	return nil
}

func (i *Integration) copyOverTo(other interface{}) error {
	var dest *Integration
	var valid bool
	if dest, valid = other.(*Integration); !valid {
		return newErrorUnsupportedType("argument given is not a *Integration type")
	}
	dest.Account = i.Account
	dest.Enabled = i.Enabled
	dest.ExpireBehavior = i.ExpireBehavior
	dest.ExpireGracePeriod = i.ExpireGracePeriod
	dest.ID = i.ID
	dest.Name = i.Name
	dest.RoleID = i.RoleID
	dest.Syncing = i.Syncing
	dest.Type = i.Type
	dest.User = i.User

	return nil
}

func (i *IntegrationAccount) copyOverTo(other interface{}) error {
	var dest *IntegrationAccount
	var valid bool
	if dest, valid = other.(*IntegrationAccount); !valid {
		return newErrorUnsupportedType("argument given is not a *IntegrationAccount type")
	}
	dest.ID = i.ID
	dest.Name = i.Name

	return nil
}

func (i *Invite) copyOverTo(other interface{}) error {
	var dest *Invite
	var valid bool
	if dest, valid = other.(*Invite); !valid {
		return newErrorUnsupportedType("argument given is not a *Invite type")
	}
	dest.ApproximateMemberCount = i.ApproximateMemberCount
	dest.ApproximatePresenceCount = i.ApproximatePresenceCount
	dest.Channel = i.Channel
	dest.Code = i.Code
	dest.CreatedAt = i.CreatedAt
	dest.Guild = i.Guild
	dest.Inviter = i.Inviter
	dest.MaxAge = i.MaxAge
	dest.MaxUses = i.MaxUses
	dest.Revoked = i.Revoked
	dest.Temporary = i.Temporary
	dest.Unique = i.Unique
	dest.Uses = i.Uses

	return nil
}

func (i *InviteMetadata) copyOverTo(other interface{}) error {
	var dest *InviteMetadata
	var valid bool
	if dest, valid = other.(*InviteMetadata); !valid {
		return newErrorUnsupportedType("argument given is not a *InviteMetadata type")
	}
	dest.CreatedAt = i.CreatedAt
	dest.Inviter = i.Inviter
	dest.MaxAge = i.MaxAge
	dest.MaxUses = i.MaxUses
	dest.Revoked = i.Revoked
	dest.Temporary = i.Temporary
	dest.Uses = i.Uses

	return nil
}

func (m *Member) copyOverTo(other interface{}) error {
	var dest *Member
	var valid bool
	if dest, valid = other.(*Member); !valid {
		return newErrorUnsupportedType("argument given is not a *Member type")
	}
	dest.Deaf = m.Deaf
	dest.GuildID = m.GuildID
	dest.JoinedAt = m.JoinedAt
	dest.Mute = m.Mute
	dest.Nick = m.Nick
	dest.Pending = m.Pending
	dest.PremiumSince = m.PremiumSince
	dest.Roles = make([]Snowflake, len(m.Roles))
	copy(dest.Roles, m.Roles)
	dest.User = m.User
	dest.UserID = m.UserID

	return nil
}

func (m *MentionChannel) copyOverTo(other interface{}) error {
	var dest *MentionChannel
	var valid bool
	if dest, valid = other.(*MentionChannel); !valid {
		return newErrorUnsupportedType("argument given is not a *MentionChannel type")
	}
	dest.GuildID = m.GuildID
	dest.ID = m.ID
	dest.Name = m.Name
	dest.Type = m.Type

	return nil
}

func (m *Message) copyOverTo(other interface{}) error {
	var dest *Message
	var valid bool
	if dest, valid = other.(*Message); !valid {
		return newErrorUnsupportedType("argument given is not a *Message type")
	}
	dest.Activity = m.Activity
	dest.Application = m.Application
	dest.Attachments = make([]*Attachment, len(m.Attachments))
	for i := 0; i < len(m.Attachments); i++ {
		dest.Attachments[i] = DeepCopy(m.Attachments[i]).(*Attachment)
	}
	dest.Author = m.Author
	dest.ChannelID = m.ChannelID
	dest.Components = make([]*MessageComponent, len(m.Components))
	for i := 0; i < len(m.Components); i++ {
		dest.Components[i] = DeepCopy(m.Components[i]).(*MessageComponent)
	}
	dest.Content = m.Content
	dest.EditedTimestamp = m.EditedTimestamp
	dest.Embeds = make([]*Embed, len(m.Embeds))
	for i := 0; i < len(m.Embeds); i++ {
		dest.Embeds[i] = DeepCopy(m.Embeds[i]).(*Embed)
	}
	dest.Flags = m.Flags
	dest.GuildID = m.GuildID
	dest.HasSpoilerImage = m.HasSpoilerImage
	dest.ID = m.ID
	dest.Interaction = m.Interaction
	dest.Member = m.Member
	dest.MentionChannels = make([]*MentionChannel, len(m.MentionChannels))
	for i := 0; i < len(m.MentionChannels); i++ {
		dest.MentionChannels[i] = DeepCopy(m.MentionChannels[i]).(*MentionChannel)
	}
	dest.MentionEveryone = m.MentionEveryone
	dest.MentionRoles = make([]Snowflake, len(m.MentionRoles))
	copy(dest.MentionRoles, m.MentionRoles)
	dest.Mentions = make([]*User, len(m.Mentions))
	for i := 0; i < len(m.Mentions); i++ {
		dest.Mentions[i] = DeepCopy(m.Mentions[i]).(*User)
	}
	dest.MessageReference = m.MessageReference
	dest.Nonce = m.Nonce
	dest.Pinned = m.Pinned
	dest.Reactions = make([]*Reaction, len(m.Reactions))
	for i := 0; i < len(m.Reactions); i++ {
		dest.Reactions[i] = DeepCopy(m.Reactions[i]).(*Reaction)
	}
	dest.ReferencedMessage = m.ReferencedMessage
	dest.SpoilerTagAllAttachments = m.SpoilerTagAllAttachments
	dest.SpoilerTagContent = m.SpoilerTagContent
	dest.StickerItems = make([]*StickerItem, len(m.StickerItems))
	for i := 0; i < len(m.StickerItems); i++ {
		dest.StickerItems[i] = DeepCopy(m.StickerItems[i]).(*StickerItem)
	}
	dest.Timestamp = m.Timestamp
	dest.Tts = m.Tts
	dest.Type = m.Type
	dest.WebhookID = m.WebhookID

	return nil
}

func (m *MessageComponent) copyOverTo(other interface{}) error {
	var dest *MessageComponent
	var valid bool
	if dest, valid = other.(*MessageComponent); !valid {
		return newErrorUnsupportedType("argument given is not a *MessageComponent type")
	}
	dest.Components = make([]*MessageComponent, len(m.Components))
	for i := 0; i < len(m.Components); i++ {
		dest.Components[i] = DeepCopy(m.Components[i]).(*MessageComponent)
	}
	dest.CustomID = m.CustomID
	dest.Disabled = m.Disabled
	dest.Emoji = m.Emoji
	dest.Label = m.Label
	dest.MaxValues = m.MaxValues
	dest.MinValues = m.MinValues
	dest.Options = make([]*SelectMenuOption, len(m.Options))
	for i := 0; i < len(m.Options); i++ {
		dest.Options[i] = DeepCopy(m.Options[i]).(*SelectMenuOption)
	}
	dest.Placeholder = m.Placeholder
	dest.Style = m.Style
	dest.Type = m.Type
	dest.Url = m.Url

	return nil
}

func (m *MessageSticker) copyOverTo(other interface{}) error {
	var dest *MessageSticker
	var valid bool
	if dest, valid = other.(*MessageSticker); !valid {
		return newErrorUnsupportedType("argument given is not a *MessageSticker type")
	}
	dest.Asset = m.Asset
	dest.Description = m.Description
	dest.FormatType = m.FormatType
	dest.ID = m.ID
	dest.Name = m.Name
	dest.PackID = m.PackID
	dest.PreviewAsset = m.PreviewAsset
	dest.Tags = m.Tags

	return nil
}

func (r *Reaction) copyOverTo(other interface{}) error {
	var dest *Reaction
	var valid bool
	if dest, valid = other.(*Reaction); !valid {
		return newErrorUnsupportedType("argument given is not a *Reaction type")
	}
	dest.Count = r.Count
	dest.Emoji = r.Emoji
	dest.Me = r.Me

	return nil
}

func (r *Role) copyOverTo(other interface{}) error {
	var dest *Role
	var valid bool
	if dest, valid = other.(*Role); !valid {
		return newErrorUnsupportedType("argument given is not a *Role type")
	}
	dest.Color = r.Color
	dest.guildID = r.guildID
	dest.Hoist = r.Hoist
	dest.ID = r.ID
	dest.Managed = r.Managed
	dest.Mentionable = r.Mentionable
	dest.Name = r.Name
	dest.Permissions = r.Permissions
	dest.Position = r.Position

	return nil
}

func (s *SelectMenuOption) copyOverTo(other interface{}) error {
	var dest *SelectMenuOption
	var valid bool
	if dest, valid = other.(*SelectMenuOption); !valid {
		return newErrorUnsupportedType("argument given is not a *SelectMenuOption type")
	}
	dest.Default = s.Default
	dest.Description = s.Description
	dest.Emoji = s.Emoji
	dest.Label = s.Label
	dest.Value = s.Value

	return nil
}

func (s *StickerItem) copyOverTo(other interface{}) error {
	var dest *StickerItem
	var valid bool
	if dest, valid = other.(*StickerItem); !valid {
		return newErrorUnsupportedType("argument given is not a *StickerItem type")
	}
	dest.FormatType = s.FormatType
	dest.ID = s.ID
	dest.Name = s.Name

	return nil
}

func (u *User) copyOverTo(other interface{}) error {
	var dest *User
	var valid bool
	if dest, valid = other.(*User); !valid {
		return newErrorUnsupportedType("argument given is not a *User type")
	}
	dest.Avatar = u.Avatar
	dest.Bot = u.Bot
	dest.Discriminator = u.Discriminator
	dest.Email = u.Email
	dest.Flags = u.Flags
	dest.ID = u.ID
	dest.Locale = u.Locale
	dest.MFAEnabled = u.MFAEnabled
	dest.PartialMember = u.PartialMember
	dest.PremiumType = u.PremiumType
	dest.PublicFlags = u.PublicFlags
	dest.System = u.System
	dest.Username = u.Username
	dest.Verified = u.Verified

	return nil
}

func (u *UserConnection) copyOverTo(other interface{}) error {
	var dest *UserConnection
	var valid bool
	if dest, valid = other.(*UserConnection); !valid {
		return newErrorUnsupportedType("argument given is not a *UserConnection type")
	}
	dest.ID = u.ID
	dest.Integrations = make([]*IntegrationAccount, len(u.Integrations))
	for i := 0; i < len(u.Integrations); i++ {
		dest.Integrations[i] = DeepCopy(u.Integrations[i]).(*IntegrationAccount)
	}
	dest.Name = u.Name
	dest.Revoked = u.Revoked
	dest.Type = u.Type

	return nil
}

func (u *UserPresence) copyOverTo(other interface{}) error {
	var dest *UserPresence
	var valid bool
	if dest, valid = other.(*UserPresence); !valid {
		return newErrorUnsupportedType("argument given is not a *UserPresence type")
	}
	dest.Game = u.Game
	dest.GuildID = u.GuildID
	dest.Nick = u.Nick
	dest.Roles = make([]Snowflake, len(u.Roles))
	copy(dest.Roles, u.Roles)
	dest.Status = u.Status
	dest.User = u.User

	return nil
}

func (v *VoiceRegion) copyOverTo(other interface{}) error {
	var dest *VoiceRegion
	var valid bool
	if dest, valid = other.(*VoiceRegion); !valid {
		return newErrorUnsupportedType("argument given is not a *VoiceRegion type")
	}
	dest.Custom = v.Custom
	dest.Deprecated = v.Deprecated
	dest.ID = v.ID
	dest.Name = v.Name
	dest.Optimal = v.Optimal
	dest.SampleHostname = v.SampleHostname
	dest.SamplePort = v.SamplePort
	dest.VIP = v.VIP

	return nil
}

func (v *VoiceState) copyOverTo(other interface{}) error {
	var dest *VoiceState
	var valid bool
	if dest, valid = other.(*VoiceState); !valid {
		return newErrorUnsupportedType("argument given is not a *VoiceState type")
	}
	dest.ChannelID = v.ChannelID
	dest.Deaf = v.Deaf
	dest.GuildID = v.GuildID
	dest.Member = v.Member
	dest.Mute = v.Mute
	dest.RequestToSpeakTimestamp = v.RequestToSpeakTimestamp
	dest.SelfDeaf = v.SelfDeaf
	dest.SelfMute = v.SelfMute
	dest.SelfStream = v.SelfStream
	dest.SelfVideo = v.SelfVideo
	dest.SessionID = v.SessionID
	dest.Suppress = v.Suppress
	dest.UserID = v.UserID

	return nil
}

func (w *Webhook) copyOverTo(other interface{}) error {
	var dest *Webhook
	var valid bool
	if dest, valid = other.(*Webhook); !valid {
		return newErrorUnsupportedType("argument given is not a *Webhook type")
	}
	dest.Avatar = w.Avatar
	dest.ChannelID = w.ChannelID
	dest.GuildID = w.GuildID
	dest.ID = w.ID
	dest.Name = w.Name
	dest.Token = w.Token
	dest.User = w.User

	return nil
}
