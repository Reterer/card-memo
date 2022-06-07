package card

import (
	"time"
)

type Group struct {
	Header      string
	Description string
	id          int32
}

type Card struct {
	Header  string
	Body    string
	id      int32
	groupId int32

	creationTime time.Time
	knowledgeVal float32
}

func MakeCard(header string, groupId int32) Card {
	return Card{
		Header:       header,
		Body:         "",
		id:           -1,
		groupId:      groupId,
		creationTime: time.Now(),
		knowledgeVal: 0,
	}
}

func (c *Card) CreationTime() time.Time {
	return c.creationTime
}

func (c *Card) KnowledgeVal() float32 {
	return c.knowledgeVal
}

func (c *Card) UpdateKnowledge(ans float32) bool {
	if ans < 0 || ans > 1 {
		return false
	}
	c.knowledgeVal = c.knowledgeVal*0.6 + ans*0.4
	return true
}
