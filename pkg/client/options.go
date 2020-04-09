package client

// PostMessageChannelOption changes the channel where to post the message.
func PostMessageChannelOption(channel string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Channel = channel
	}
}

// PostMessageWorkspaceOption changes the workspace where to post the message.
func PostMessageWorkspaceOption(workspace string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Workspace = workspace
	}
}

// PostMessageUserEmailsOption changes the users emails to whom post the message.
func PostMessageUserEmailsOption(userEmails... string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.UserEmails = userEmails
	}
}

// PostMessageAddUserEmailsOption adds users emails to whom post the message.
func PostMessageAddUserEmailsOption(userEmails... string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.UserEmails = append(parameters.UserEmails, userEmails...)
	}
}

// PostMessageAsUserOption sets the AsUser parameter.
func PostMessageAsUserOption(asUser bool) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.AsUser = asUser
	}
}

// PostMessageUsernameOption sets the Username parameter.
//
// Sets AsUser to `false` for Username not to be ignored.
func PostMessageUsernameOption(username string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.AsUser = false
		parameters.Username = username
	}
}

// PostMessageIconEmojiOption sets the IconEmoji parameter.
//
// Sets AsUser to `false` for IconEmoji not to be ignored.
func PostMessageIconEmojiOption(iconEmoji string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.AsUser = false
		parameters.IconEmoji = iconEmoji
	}
}

// PostMessageIconURLOption sets the IconURL parameter.
//
// Sets AsUser to `false` for IconURL not to be ignored.
func PostMessageIconURLOption(iconURL string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.AsUser = false
		parameters.IconURL = iconURL
	}
}

// PostMessageParseOption sets the Parse parameter.
func PostMessageParseOption(parse string) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Parse = parse
	}
}

// PostMessageMarkdownOption sets the Markdown parameter.
func PostMessageMarkdownOption(markdown bool) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Markdown = markdown
	}
}

// PostMessageLinkNamesOption sets the LinkNames parameter.
func PostMessageLinkNamesOption(linkNames bool) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.LinkNames = linkNames
	}
}

// PostMessageUnfurlLinksOption sets the UnfurlLinks parameter.
func PostMessageUnfurlLinksOption(unfurlLinks bool) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.UnfurlLinks = unfurlLinks
	}
}

// PostMessageUnfurlMediaOption sets the UnfurlMedia parameter.
func PostMessageUnfurlMediaOption(unfurlMedia bool) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.UnfurlMedia = unfurlMedia
	}
}

// PostMessageAttachmentsOption changes the users emails to whom post the message.
func PostMessageAttachmentsOption(attachments... map[string]interface{}) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Attachments = attachments
	}
}

// PostMessageAddAttachmentsOption adds users emails to whom post the message.
func PostMessageAddAttachmentsOption(attachments... map[string]interface{}) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Attachments = append(parameters.Attachments, attachments...)
	}
}

// PostMessageBlocksOption changes the users emails to whom post the message.
func PostMessageBlocksOption(blocks... map[string]interface{}) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Blocks = blocks
	}
}

// PostMessageAddBlocksOption adds users emails to whom post the message.
func PostMessageAddBlocksOption(blocks... map[string]interface{}) PostMessageOptions {
	return func (parameters *PostMessageParameters) {
		parameters.Blocks = append(parameters.Blocks, blocks...)
	}
}
