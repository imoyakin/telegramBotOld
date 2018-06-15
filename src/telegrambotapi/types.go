package telegrambotapi

import (
	"encoding/json"
)

type APIResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

type ResponseParameters struct{}

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

type GetUpdates struct {
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	Timeout       int    `json:"timeout"`
	AllowedUpdate string `json:"allowed_updates"`
}

type SetWebHook struct {
	Url            string `json:"url"`
	Certificate    []byte `json:"certificate"`
	MaxConnections int    `json:"max_connections"`
	AllowedUpdate  string `json:"allowed_updates"`
}

type WebHookInfo struct {
	Url                  string `json:"url"`
	HasCustomCertificate bool   `json:"has_custom_certificate"`
	PendingUpdateCount   int    `json:"pending_update_count"`
	LastErrorDate        int    `json:"last_error_date"`
	LastErrorMessage     string `json:"last_error_message"`
	MaxConnections       int    `json:"max_connections"`
	AllowedUpdates       string `json:"allowed_updates"`
}

type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID                          int64      `json:"id"`
	Type                        string   `json:"type"`
	Title                       string   `json:"title"`
	UserName                    string   `json:"username"`
	FirstName                   string   `json:"first_name"`
	LastName                    string   `json:"last_name"`
	AllMembersAreAdministrators bool     `json:"all_members_are_administrators"`
	photo                       []byte   `json:"photo"`
	Description                 string   `json:"description"`
	InviteLink                  string   `json:"invite_link"`
	PinnedMessage               *Message `json:"pinned_message"`
	StickerSetName              string   `json:"sticker_set_name"`
	CanSetStickerSet            bool     `json:"can_set_sticker_set"`
}

type Message struct {
	MessageID             int                `json:"message_id"`
	From                  *User              `json:"from"`
	Date                  int                `json:"date"`
	Chat                  *Chat              `json:"chat"`
	ForwardFrom           *User              `json:"forward_from"`
	ForwardFromChat       *Chat              `json:"forward_from_chat`
	ForwardFromMessageID  int                `json:"forward_from_message_id"`
	ForwardSignature      string             `json:"forward_signature"`
	ForwardDate           int                `json:"forward_date"`
	ReplyToMessage        *Message           `json:"reply_to_message"`
	EditDate              int                `json:"edit_date"`
	MediaGroupID          string             `json:"media_group_id"`
	AuthorSignature       string             `json:"author_signature"`
	Text                  string             `json:"text"`
	Entites               *[]MessageEntity   `json:"entities"`
	CaptionEntites        *MessageEntity     `json:"caption_entities"`
	Audio                 *Audio             `json:"audio"`
	Document              *Document          `json:"document"`
	Game                  *Game              `json:"game"`
	Photo                 *[]PhotoSize       `json:"photo"`
	Sticker               *Sticker           `json:"sticker"`
	Video                 *Video             `json:"video"`
	Voice                 *Voice             `json:"voice"`
	VideoNote             *VideoNote         `json:"video_note"`
	Caption               string             `json:"caption"`
	Contact               *Contact           `json:"contact"`
	Location              *Location          `json:"location"`
	Venue                 *Venue             `json:"venue"`
	NewChatMembers        *[]User            `json:"new_chat_members"`
	LeftChatMembers       *User              `json:"left_chat_member"`
	NewChatTitle          string             `json:"new_chat_title"`
	NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`
	GroupChatCreated      bool               `json:"group_chat_created"`
	SuperGroupChatCreated bool               `json:"supergroup_chat_created"`
	ChannelChatCreated    bool               `json:"channel_chat_created"`
	MigrateToChatID       int                `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                `json:"migrate_from_chat_id"`
	PinnedMessage         *Message           `json:"pinned_message"`
	Invoice               *Invoice           `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
	ConnectedWebsite      string             `json:"connected_website"`
}

type SendPhoto struct {
	ChatID              string       `josn:"string"`
	Photo               *InputFile   `josn:"photo"`
	Caption             string       `json:"caption"`
	ParseMode           string       `json:"parse_mode"`
	DisableNotification bool         `json:"disable_notification"`
	ReplyToMessageID    int          `json:"reply_to_message_id"`
	ReplyMarkup         *ReplyStruct `json:"reply_markup"`
}
type ReplyStruct struct {
	InlineKeyboard      *InlineKeyboard
	ReplyKeyBoardMarkUP *ReplyKeyBoardMarkUP
}

type ReplyKeyBoardMarkUP struct {
}
type InputFile struct {
	FileID string
	File   []byte
}
type MessageEntity struct{}
type Audio struct{}
type Document struct{}
type Game struct{}
type PhotoSize struct{}
type Sticker struct{}
type Video struct{}
type Voice struct{}
type VideoNote struct{}
type Invoice struct{}
type SuccessfulPayment struct{}
type Contact struct{}
type Location struct{}
type Venue struct{}

type InlineQuery struct{}

type ChosenInlineResult struct{}

type CallbackQuery struct{}

type ShippingQuery struct{}

type PreCheckoutQuery struct{}

type GetMe struct {
	ChatID                string         `json:"chat_id"`
	Text                  string         `json:"text"`
	ParseMode             string         `json:"parse_mode"`
	DisableWebPagePreview bool           `json:"disable_web_page_preview"`
	DisableNotification   bool           `json:"disable_notification"`
	ReplyToMessageID      int            `json:"reply_to_message_id"`
	ReplyMarkup           InlineKeyboard `json:"reply_markup"`
}

type InlineKeyboard struct{}

type SendMessage struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ReplyToMessageID      int    `json:"reply_to_message_id"`
	//ReplyMarkup 结果值待封装
	ReplyMarkup *InlineKeyboard `json:"reply_markup"`
}
