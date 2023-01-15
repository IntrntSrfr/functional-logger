package kvstore

import (
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
)

type DiscordMessage struct {
	Message     *discordgo.Message
	Attachments [][]byte
}

func NewDiscordMessage(msg *discordgo.Message, maxSize int) *DiscordMessage {
	m := &DiscordMessage{
		Message:     msg,
		Attachments: [][]byte{},
	}

	for _, a := range msg.Attachments {
		if a.Size > maxSize {
			continue
		}

		data, err := GetAttachment(a.URL)
		if err != nil {
			continue
		}

		m.Attachments = append(m.Attachments, data)
	}
	return m
}

func GetAttachment(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

type ByID []*DiscordMessage

func (m ByID) Len() int {
	return len(m)
}

func (m ByID) Less(i, j int) bool {
	return m[i].Message.ID < m[j].Message.ID
}

func (m ByID) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
