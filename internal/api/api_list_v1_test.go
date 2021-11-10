package api

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var respList = []model.Workplace {
	{ID: 1, Name: "elem 1", Size: 10},
	{ID: 2, Name: "elem 2", Size: 10},
	{ID: 3, Name: "elem 3", Size: 10},
}

func TestListWorkplacesV1(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.WorkplaceService.EXPECT().ListWorkplaces(gomock.Any()).Return(respList, nil).Times(1)

	req := &pb.ListWorkplacesV1Request{}
	resp, err := apiFixture.Api.ListWorkplacesV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")

	var expected = []*pb.Workplace {
		{Id: 1, Name: "elem 1", Size: 10},
		{Id: 2, Name: "elem 2", Size: 10},
		{Id: 3, Name: "elem 3", Size: 10},
	}

	assert.Equal(t, expected, resp.GetItems(), "WorkplaceId must be %v", expected)
}

func TestUnsuccessfulListWorkplacesV1_Internal(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.WorkplaceService.EXPECT().ListWorkplaces(gomock.Any()).Return(nil, errors.New("Error")).Times(1)

	req := &pb.ListWorkplacesV1Request{}
	_, err := apiFixture.Api.ListWorkplacesV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.Internal
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func TestUnsuccessfulListWorkplacesV1_Notfound(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.WorkplaceService.EXPECT().ListWorkplaces(gomock.Any()).Return(nil, nil).Times(1)

	req := &pb.ListWorkplacesV1Request{}
	_, err := apiFixture.Api.ListWorkplacesV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.NotFound
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

