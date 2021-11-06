package api

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulCreateWorkplaceV1(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.Repo.EXPECT().CreateWorkplace(gomock.Any(), gomock.Any()).Return(uint64(1), nil).Times(1)

	req := &pb.CreateWorkplaceV1Request{Foo: "test value"}
	resp, err := apiFixture.Api.CreateWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")

	var expectedId uint64 = 1
	assert.Equal(t, expectedId, resp.GetWorkplaceId(), "WorkplaceId must be %d", expectedId)
}

func TestUnsuccessfulCreateWorkplaceV1_Internal(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	apiFixture.Repo.EXPECT().CreateWorkplace(gomock.Any(), gomock.Any()).Return(uint64(0), errors.New("error")).Times(1)

	req := &pb.CreateWorkplaceV1Request{Foo: "test value"}
	_, err := apiFixture.Api.CreateWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.Internal
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}
