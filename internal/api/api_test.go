package api

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-workplace-api/internal/mocks"
	bss_workplace_api "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"testing"
)

type DBMockFixture struct {
	Ctrl *gomock.Controller
	Repo *mocks.MockRepo
	Api  bss_workplace_api.BssWorkplaceApiServiceServer
}

func Setup(t *testing.T) DBMockFixture {

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)
	api := NewWorkplaceAPI(repo)

	return DBMockFixture{
		Ctrl: ctrl,
		Repo: repo,
		Api: api,
	}

}