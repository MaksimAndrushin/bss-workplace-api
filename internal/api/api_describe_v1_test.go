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

func TestDescribeWorkplaceV1(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.WorkplaceService.EXPECT().DescribeWorkplace(gomock.Any(), gomock.Any()).Return(&model.Workplace{ID: 1, Name: "test", Size: 10}, nil).Times(1)

	req := &pb.DescribeWorkplaceV1Request{WorkplaceId: 1}
	resp, err := apiFixture.Api.DescribeWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")

	var expectedId uint64 = 1
	assert.Equal(t, expectedId, resp.GetValue().GetId(), "WorkplaceId must be %d", expectedId)

	var expectedName string = "test"
	assert.Equal(t, expectedName, resp.GetValue().GetName(), "Workplace name must be %s", expectedName)
}

func TestUnsuccessfulDescribeWorkplaceV1_Internal(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.WorkplaceService.EXPECT().DescribeWorkplace(gomock.Any(), gomock.Any()).Return(nil, errors.New("Error")).Times(1)

	req := &pb.DescribeWorkplaceV1Request{WorkplaceId: 1}
	_, err := apiFixture.Api.DescribeWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.Internal
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func TestUnsuccessfulDescribeWorkplaceV1_Argument(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	req := &pb.DescribeWorkplaceV1Request{WorkplaceId: 0}
	_, err := apiFixture.Api.DescribeWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.InvalidArgument
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}
