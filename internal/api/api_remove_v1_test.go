package api

import (
	"errors"
	"github.com/golang/mock/gomock"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestRemoveWorkplaceV1(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.Repo.EXPECT().RemoveWorkplace(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)

	req := &pb.RemoveWorkplaceV1Request{WorkplaceId: 1}
	resp, err := apiFixture.Api.RemoveWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")

	var expected bool = true
	assert.Equal(t, expected, resp.GetFound(), "WorkplaceId must be %d", expected)
}

func TestUnsuccessfulRemoveWorkplaceV1_Internal(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.Repo.EXPECT().RemoveWorkplace(gomock.Any(), gomock.Any()).Return(false, errors.New("Error")).Times(1)

	req := &pb.RemoveWorkplaceV1Request{WorkplaceId: 1}
	_, err := apiFixture.Api.RemoveWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.Internal
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func TestUnsuccessfulRemoveWorkplaceV1_Notfound(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.Repo.EXPECT().RemoveWorkplace(gomock.Any(), gomock.Any()).Return(false, nil).Times(1)

	req := &pb.RemoveWorkplaceV1Request{WorkplaceId: 1}
	_, err := apiFixture.Api.RemoveWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.NotFound
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func TestUnsuccessfulRemoveWorkplaceV1_Argument(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	req := &pb.RemoveWorkplaceV1Request{WorkplaceId: 0}
	_, err := apiFixture.Api.RemoveWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.InvalidArgument
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}
