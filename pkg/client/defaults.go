package client

func defaultPostMessageParameters() *PostMessageParameters {
	return &PostMessageParameters{
		Text:        "",
		Markdown:    true,
		LinkNames:   true,
		UnfurlLinks: true,
		UnfurlMedia: true,
	}
}