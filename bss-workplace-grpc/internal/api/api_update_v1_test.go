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
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestSuccessfulUpdateWorkplaceV1(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	modelWorkplace, pbWorkplace := generateTestWorkplace()

	apiFixture.WorkplaceService.EXPECT().UpdateWorkplace(gomock.Any(), modelWorkplace).Return(true, nil).Times(1)

	req := &pb.UpdateWorkplaceV1Request{Value: &pbWorkplace}

	resp, err := apiFixture.Api.UpdateWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")

	var expected bool = true
	assert.Equal(t, expected, resp.GetUpdated(), "WorkplaceId must be %d", expected)
}

func TestUnsuccessfulUpdateWorkplaceV1_Internal(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)
	modelWorkplace, pbWorkplace := generateTestWorkplace()

	apiFixture.WorkplaceService.EXPECT().UpdateWorkplace(gomock.Any(), modelWorkplace).Return(false, errors.New("Error")).Times(1)

	req := &pb.UpdateWorkplaceV1Request{Value: &pbWorkplace}
	_, err := apiFixture.Api.UpdateWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.Internal
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func TestUnsuccessfulUpdateWorkplaceV1_Notfound(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)
	modelWorkplace, pbWorkplace := generateTestWorkplace()

	apiFixture.WorkplaceService.EXPECT().UpdateWorkplace(gomock.Any(), modelWorkplace).Return(false, nil).Times(1)

	req := &pb.UpdateWorkplaceV1Request{Value: &pbWorkplace}
	_, err := apiFixture.Api.UpdateWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.NotFound
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func TestUnsuccessfulUpdateWorkplaceV1_Argument(t *testing.T) {
	t.Parallel()
	apiFixture := Setup(t)

	req := &pb.UpdateWorkplaceV1Request{Value: nil}
	_, err := apiFixture.Api.UpdateWorkplaceV1(context.Background(), req)

	assert.NotNil(t, err, "Error must not be nil")

	expectedStatus := codes.InvalidArgument
	actualStatus, _ := status.FromError(err)
	assert.Equal(t, expectedStatus, actualStatus.Code())
}

func generateTestWorkplace() (model.Workplace, pb.Workplace) {
	modelWorkplace := model.Workplace{
		ID:   1,
		Name: "test",
		Size: 10,
		Created: time.Now().UTC(),
	}

	pbWorkplace := pb.Workplace{
		Id:   modelWorkplace.ID,
		Name: modelWorkplace.Name,
		Size: modelWorkplace.Size,
		Created: timestamppb.New(modelWorkplace.Created),
	}

	return modelWorkplace, pbWorkplace
}