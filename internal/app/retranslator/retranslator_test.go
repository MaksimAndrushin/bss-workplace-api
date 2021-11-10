package retranslator

import (
	"errors"
	"github.com/ozonmp/bss-workplace-api/internal/mocks/fixtures/events"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"golang.org/x/net/context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

var eventsData = []model.WorkplaceEvent{
	{ID: 1, Type: 0, Status: 1, Entity: &model.Workplace{ID: 1}},
	{ID: 2, Type: 0, Status: 1, Entity: &model.Workplace{ID: 2}},
	{ID: 3, Type: 0, Status: 1, Entity: &model.Workplace{ID: 3}},
	{ID: 4, Type: 0, Status: 1, Entity: &model.Workplace{ID: 1}},
}

func TestWithoutErrors(t *testing.T) {
	t.Parallel()

	fixture := events.Setup(t)
	defer fixture.TearDown()

	fixture.Repo.EXPECT().Lock(gomock.Any(), uint64(4), nil).Return(eventsData, nil).Times(1)
	fixture.Sender.EXPECT().Send(gomock.Any()).Return(nil).Times(4)
	fixture.Repo.EXPECT().Remove(gomock.Any(), gomock.Any(), nil).Return(nil).Times(4)

	startRetranslator(fixture)
}

func TestKafkaAndDBUpdErrors(t *testing.T) {
	t.Parallel()

	fixture := events.Setup(t)
	defer fixture.TearDown()

	fixture.Repo.EXPECT().Lock(gomock.Any(), uint64(4), nil).Return(eventsData, nil).Times(1)

	fixture.Sender.EXPECT().Send(gomock.Any()).Return(nil).Times(1)
	fixture.Repo.EXPECT().Remove(gomock.Any(), gomock.Any(), nil).Return(errors.New("Remove error")).Times(1)

	fixture.Sender.EXPECT().Send(gomock.Any()).Return(errors.New("Send error")).Times(1)
	fixture.Repo.EXPECT().Unlock(gomock.Any(), gomock.Any(), nil).Return(errors.New("Unlock error")).Times(1)

	fixture.Sender.EXPECT().Send(gomock.Any()).Return(nil).Times(1)
	fixture.Repo.EXPECT().Remove(gomock.Any(), gomock.Any(), nil).Return(nil).Times(1)

	fixture.Sender.EXPECT().Send(gomock.Any()).Return(nil).Times(1)
	fixture.Repo.EXPECT().Remove(gomock.Any(), gomock.Any(), nil).Return(nil).Times(1)

	startRetranslator(fixture)
}

func TestLockErrors(t *testing.T) {
	t.Parallel()

	fixture := events.Setup(t)
	defer fixture.TearDown()

	fixture.Repo.EXPECT().Lock(gomock.Any(), uint64(4), nil).Return(nil, errors.New("Lock error")).Times(1)

	startRetranslator(fixture)
}

func startRetranslator(fixture events.RetranslatorMockFixture) {
	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  1,
		ConsumeSize:    4,
		ConsumeTimeout: 3 * time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           fixture.Repo,
		Sender:         fixture.Sender,
	}

	retranslator := NewRetranslator(cfg)
	retranslator.Start(context.Background())

	time.Sleep(5 * time.Second)

	retranslator.Close()
}
