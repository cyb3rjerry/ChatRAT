package telegram

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lithammer/shortuuid/v4"
)

type TelegramConfig struct {
	ApiKey    string
	ChannelID int64
}

type TelegramBot struct {
	ChannelID int64
	Bot       *tgbotapi.BotAPI
}

func NewConn(config TelegramConfig) (*TelegramBot, error) {
	bot := new(TelegramBot)
	botInstance, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, err
	}

	bot.Bot = botInstance
	bot.ChannelID = config.ChannelID

	return bot, nil
}

// TODO: remove testing stuff
func (tb TelegramBot) SendMessage(content string) error {

	msg := tgbotapi.NewMessage(tb.ChannelID, content)
	// TODO: evaluate if resp contains valuable info
	_, err := tb.Bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (tb TelegramBot) GetMessages() ([]string, error) {

	u := tgbotapi.NewUpdate(0)

	updates, err := tb.Bot.GetUpdates(u)
	if err != nil {
		return nil, err
	}

	var msgs []string

	for _, update := range updates {
		msgs = append(msgs, update.Message.Text)
	}

	return msgs, err

}

func (tb TelegramBot) SendMedia(text, fPath string) error {
	file := tgbotapi.FilePath(fPath)
	doc := tgbotapi.NewDocument(tb.ChannelID, file)
	doc.Caption = "Test caption"
	_, err := tb.Bot.Send(doc)
	if err != nil {
		return err
	}

	return nil
}

func (tb TelegramBot) GetLatestMedia(outPath string) error {
	u := tgbotapi.NewUpdate(0)

	updates, err := tb.Bot.GetUpdates(u)
	if err != nil {
		return err
	}

	var latestFileID string

	// This sucks, but I don't know how to do it better
	// Telegram has way too many types of files
	switch m := updates[len(updates)-1].Message; {
	case m.Animation != nil:
		latestFileID = updates[len(updates)-1].Message.Animation.FileID
	case m.Document != nil:
		latestFileID = updates[len(updates)-1].Message.Document.FileID
	case m.Audio != nil:
		latestFileID = updates[len(updates)-1].Message.Audio.FileID
	case m.Photo != nil:
		latestFileID = updates[len(updates)-1].Message.Photo[len(m.Photo)-1].FileID
	case m.Sticker != nil:
		latestFileID = updates[len(updates)-1].Message.Sticker.FileID
	case m.Video != nil:
		latestFileID = updates[len(updates)-1].Message.Video.FileID
	case m.VideoNote != nil:
		latestFileID = updates[len(updates)-1].Message.VideoNote.FileID
	case m.Voice != nil:
		latestFileID = updates[len(updates)-1].Message.Voice.FileID
	default:
		return fmt.Errorf("no file found")
	}

	downloadUrl, err := tb.Bot.GetFileDirectURL(latestFileID)
	if err != nil {
		return err
	}

	filename := path.Base(downloadUrl)

	out, err := os.Create(outPath + "/" + shortuuid.New() + "_" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(downloadUrl)
	log.Println(resp.Request.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status on download: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
