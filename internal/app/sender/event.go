package sender

import (
    "github.com/execut/ozon-keyword-api/internal/model"
)

type EventSender interface {
    Send(keywordEvent *model.KeywordEvent) error
}

func NewStubEventSender() StubEventSender {
    return StubEventSender{}
}

type StubEventSender struct{}

func (StubEventSender) Send(keywordEvent *model.KeywordEvent) error {
    return nil
}
