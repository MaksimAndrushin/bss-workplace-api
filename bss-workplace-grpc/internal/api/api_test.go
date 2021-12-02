package api

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-workplace-api/internal/mocks"
	bss_workplace_api "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"testing"
)

type DBMockFixture struct {
	Ctrl             *gomock.Controller
	WorkplaceService *mocks.MockWorkplaceService
	Api              bss_workplace_api.BssWorkplaceApiServiceServer
}

func Setup(t *testing.T) DBMockFixture {

	ctrl := gomock.NewController(t)
	workplaceService := mocks.NewMockWorkplaceService(ctrl)
	api := &workplaceAPI{
		WorkplaceService: workplaceService,
	}

	return DBMockFixture{
		Ctrl:             ctrl,
		WorkplaceService: workplaceService,
		Api:              api,
	}

}
