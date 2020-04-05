package task

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/jgengo/slack_that/internal/utils"

	"github.com/slack-go/slack"
	"golang.org/x/time/rate"
)

var s = newSlackTask()

type SlackClient struct {
	Value *slack.Client
}

// Gateway is the Slack API client gateway
var Gateway = make(map[string]SlackClient)

// SlackRequest ...
type SlackRequest struct {
	Workspace   string               `json:"workspace"`
	Channel     string               `json:"channel"`
	UserEmail   string               `json:"userEmail"`
	Username    string               `json:"username"`
	Parse       string               `json:"parse"`
	IconURL     string               `json:"icon_url"`
	IconEmoji   string               `json:"icon_emoji"`
	Text        string               `json:"text"`
	LinkNames   int                  `json:"link_names"`
	AsUser      bool                 `json:"as_user"`
	UnfurlLinks bool                 `json:"unfurl_links"`
	UnfurlMedia bool                 `json:"unfurl_media"`
	Markdown    bool                 `json:"mrkdwn,omitempty"`
	Blocks      []slack.SectionBlock `json:"blocks"`
	Attachments []slack.Attachment   `json:"attachments"`
}

// SlackTask ...
type SlackTask struct {
	limit          *rate.Limiter
	activeTasks    uint
	maxActiveTasks uint
	mu             sync.Mutex
}

func newSlackTask() *SlackTask {
	return &SlackTask{
		limit: rate.NewLimiter(1, 1),
	}
}

func (s *SlackTask) doSlackTask(channel string, body *SlackRequest, options []slack.MsgOption) {
	s.limit.Wait(context.Background())

	s.mu.Lock()
	s.activeTasks--
	s.mu.Unlock()

	if body.UserEmail != "" {
		if imChannel, err := Gateway[body.Workspace].GetIM(body.UserEmail); err != nil {
			log.Printf("%sslack (error)%s while trying to get user IM ID. (%v)\n", utils.Red, utils.Reset, err)
		} else {
			channel = imChannel
		}
	}

	_, _, err := Gateway[body.Workspace].Value.PostMessage(channel, options...)

	if err != nil {
		log.Printf("%sslack (error)%s while trying to PostMessage(). (%v)\n", utils.Red, utils.Reset, err)
	} else {
		log.Printf("slack (info) successfully posted a message to '%s' on '%s'\n", channel, body.Workspace)
	}
}

// NewHealthResponse returns the size of the tasks queue
func NewHealthResponse() *utils.HealthResponse {
	s.mu.Lock()
	defer s.mu.Unlock()
	return &utils.HealthResponse{
		ActiveTasks:    s.activeTasks,
		MaxActiveTasks: s.maxActiveTasks,
	}
}

// ProcessCreate will threat the body and do the job!
func (b *SlackRequest) ProcessCreate() error {
	if err := b.checkParam(); err != nil {
		return err
	}

	s.mu.Lock()
	s.activeTasks++
	if s.activeTasks > s.maxActiveTasks {
		s.maxActiveTasks = s.activeTasks
	}
	s.mu.Unlock()

	myParam := b.buildParam()

	go s.doSlackTask(b.Channel, b, myParam)
	return nil
}

func (b *SlackRequest) buildParam() []slack.MsgOption {
	options := []slack.MsgOption{}

	postParameters := slack.PostMessageParameters{
		Username:    b.Username,
		AsUser:      b.AsUser,
		Parse:       b.Parse,
		LinkNames:   b.LinkNames,
		UnfurlLinks: b.UnfurlLinks,
		UnfurlMedia: b.UnfurlMedia,
		IconURL:     b.IconURL,
		IconEmoji:   b.IconEmoji,
		Markdown:    b.Markdown,
	}

	options = append(options, slack.MsgOptionPostMessageParameters(postParameters))
	options = append(options, slack.MsgOptionText(b.Text, true))

	if len(b.Blocks) != 0 {
		for _, block := range b.Blocks {
			options = append(options, slack.MsgOptionBlocks(block))
		}
	}

	if len(b.Attachments) != 0 {
		for _, attachment := range b.Attachments {
			options = append(options, slack.MsgOptionAttachments(attachment))
		}
	}

	return options
}

func (b *SlackRequest) checkParam() error {
	if b.Workspace == "" {
		return errors.New("workspace is required")
	}
	if _, ok := Gateway[b.Workspace]; !ok {
		return errors.New("workspace not found")
	}
	if len(b.Blocks) == 0 && len(b.Attachments) == 0 && b.Text == "" {
		return errors.New("text is required")
	}

	return nil
}
