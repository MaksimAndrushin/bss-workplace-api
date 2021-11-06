package events

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-workplace-api/internal/mocks"
	"testing"
)

type RetranslatorMockFixture struct {
	Ctrl   *gomock.Controller
	Repo   *mocks.MockEventRepo
	Sender *mocks.MockEventSender
}

func Setup(t *testing.T) RetranslatorMockFixture {
	ctrl := gomock.NewController(t)

	return RetranslatorMockFixture{
		Ctrl:   ctrl,
		Repo:   mocks.NewMockEventRepo(ctrl),
		Sender: mocks.NewMockEventSender(ctrl),
	}
}

func (f RetranslatorMockFixture) TearDown() {
	f.Ctrl.Finish()
}
