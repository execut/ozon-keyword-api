package producer

import (
    "github.com/execut/omp-ozon-api/internal/mocks"
    "github.com/execut/omp-ozon-api/internal/model"
    "github.com/gammazero/workerpool"
    "go.uber.org/mock/gomock"
    "gotest.tools/v3/assert"
    "testing"
    "time"
)

func Test_producer_Start(t *testing.T) {
    t.Parallel()
    t.Run("Start send event via sender", func(t *testing.T) {
        t.Parallel()
        eventCh := make(chan *model.KeywordEvent)
        defer close(eventCh)
        eventLink := newEventLink()
        ctrl := gomock.NewController(t)
        sender := mocks.NewMockEventSender(ctrl)
        sender.EXPECT().
            Send(gomock.Eq(eventLink)).Times(1)
        sut := NewProducer(eventCh, sender, 1, workerpool.New(1))
        defer sut.Close()

        sut.Start()
        eventCh <- eventLink
        time.Sleep(time.Microsecond * 25)
    })
    t.Run("Start send event via sender multiple times", func(t *testing.T) {
        t.Parallel()
        eventCh := make(chan *model.KeywordEvent)
        defer close(eventCh)
        eventLink := newEventLink()
        ctrl := gomock.NewController(t)
        sender := mocks.NewMockEventSender(ctrl)
        sender.EXPECT().
            Send(gomock.Eq(eventLink)).Times(2)
        sut := NewProducer(eventCh, sender, 2, workerpool.New(1))
        defer sut.Close()

        sut.Start()
        eventCh <- eventLink
        eventCh <- eventLink
        time.Sleep(time.Microsecond * 25)
    })
    t.Run("Close stopped all following sends", func(t *testing.T) {
        t.Parallel()
        eventCh := make(chan *model.KeywordEvent, 1)
        defer close(eventCh)
        eventLink := newEventLink()
        ctrl := gomock.NewController(t)
        sender := mocks.NewMockEventSender(ctrl)
        sender.EXPECT().Send(gomock.Any()).Times(0)
        sut := NewProducer(eventCh, sender, 1, workerpool.New(1))
        defer sut.Close()

        sut.Start()
        sut.Close()

        eventCh <- eventLink
        time.Sleep(time.Microsecond * 25)
        assert.Equal(t, 1, len(eventCh))
    })
}

func newEventLink() *model.KeywordEvent {
    keyword := model.NewTestKeyword(2)
    event := model.NewTestKeywordEvent(1, &keyword)
    eventLink := &event
    return eventLink
}
