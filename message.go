package dingtalk_group_robot

type MsgType string

const TEXT MsgType = "text"
const LINK MsgType = "link"
const MARKDOWN MsgType = "markdown"
const ACTIONCARD MsgType = "actionCard"
const FEEDCARD MsgType = "feedCard"

type TextMessage struct {
	MsgType MsgType `json:"msgtype"`
	At      *At     `json:"at"`
	Text    *Text   `json:"text"`
}

type LinkMessage struct {
	MsgType MsgType `json:"msgtype"`
	Link    *Link   `json:"link"`
}
type MarkDownMessage struct {
	MsgType  MsgType   `json:"msgtype"`
	At       *At       `json:"at"`
	MarkDown *MarkDown `json:"markdown"`
}
type ActionCardMessage struct {
	MsgType    MsgType     `json:"msgtype"`
	ActionCard *ActionCard `json:"actionCard"`
}
type FeedCardMessage struct {
	MsgType  MsgType   `json:"msgtype"`
	FeedCard *FeedCard `json:"feedCard"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Text struct {
	Content string `json:"content"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
}

type MarkDown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	Btns           []*Btn `json:"btns"`
}

type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type FeedCard struct {
	Links []*FeedCardLink `json:"links"`
}

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

/*** Text ***/
func NewText() *TextMessage {
	return &TextMessage{TEXT, &At{}, &Text{}}
}

func (t *TextMessage) SetContent(content string) *TextMessage {
	t.Text.Content = content
	return t
}

func (t *TextMessage) AtMobiles(mobiles []string) *TextMessage {
	t.At.AtMobiles = mobiles
	return t
}

func (t *TextMessage) AtAll() *TextMessage {
	t.At.IsAtAll = true
	return t
}

/*** Link ***/
func NewLink() *LinkMessage {
	return &LinkMessage{LINK, &Link{}}
}

func (l *LinkMessage) SetContent(title, text, messageUrl, picUrl string) *LinkMessage {
	l.Link.Title = title
	l.Link.Text = text
	l.Link.MessageUrl = messageUrl
	l.Link.PicUrl = picUrl
	return l
}

/*** markdown ***/
func NewMarkDown() *MarkDownMessage {
	return &MarkDownMessage{MARKDOWN, &At{}, &MarkDown{}}
}

func (m *MarkDownMessage) SetContent(title, text string) *MarkDownMessage {
	m.MarkDown.Title = title
	m.MarkDown.Text = text
	return m
}

func (m *MarkDownMessage) AtMobiles(mobiles []string) *MarkDownMessage {
	m.At.AtMobiles = mobiles
	return m
}

func (m *MarkDownMessage) AtAll() *MarkDownMessage {
	m.At.IsAtAll = true
	return m
}

/*** ActionCard ***/
func NewActionCard() *ActionCardMessage {
	return &ActionCardMessage{ACTIONCARD, &ActionCard{}}
}

func (a *ActionCardMessage) SetContent(title, text string) *ActionCardMessage {
	a.ActionCard.Title = title
	a.ActionCard.Text = text
	return a
}

func (a *ActionCardMessage) AddBtn(singleTitle, singleURL string) *ActionCardMessage {
	a.ActionCard.SingleTitle = singleTitle
	a.ActionCard.SingleURL = singleURL
	return a
}

func (a *ActionCardMessage) AddBtns(btns [][]string) *ActionCardMessage {
	for _, item := range btns {
		a.ActionCard.Btns = append(a.ActionCard.Btns, &Btn{item[0], item[1]})
	}
	return a
}

func (a *ActionCardMessage) HideAvatar() *ActionCardMessage {
	a.ActionCard.HideAvatar = "1"
	return a
}

func (a *ActionCardMessage) BtnOrientation() *ActionCardMessage {
	a.ActionCard.BtnOrientation = "0"
	return a
}

/*** FeedCard ***/
func NewFeedCard() *FeedCardMessage {
	return &FeedCardMessage{FEEDCARD, &FeedCard{}}
}

func (f *FeedCardMessage) AddCard(title, messageURL, picURL string) *FeedCardMessage {
	f.FeedCard.Links = append(f.FeedCard.Links, &FeedCardLink{title, messageURL, picURL})
	return f
}

func (f *FeedCardMessage) AddCards(cards [][]string) *FeedCardMessage {
	for _, item := range cards {
		f.FeedCard.Links = append(f.FeedCard.Links, &FeedCardLink{item[0], item[1], item[2]})
	}
	return f
}
