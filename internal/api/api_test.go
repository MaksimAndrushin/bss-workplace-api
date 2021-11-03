package api

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/mocks"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWorkplaceV1(t *testing.T) {
	t.Parallel()
	workplaceAPI := NewWorkplaceAPI(mocks.NewWorkplaceDbRepoMock())

	req := &pb.CreateWorkplaceV1Request{Foo: "test value"}
	resp, err := workplaceAPI.CreateWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")
}

func TestDescribeWorkplaceV1(t *testing.T) {
	t.Parallel()
	workplaceAPI := NewWorkplaceAPI(mocks.NewWorkplaceDbRepoMock())

	req := &pb.DescribeWorkplaceV1Request{WorkplaceId: 1}
	resp, err := workplaceAPI.DescribeWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")
}

func TestListWorkplaceV1(t *testing.T) {
	t.Parallel()
	workplaceAPI := NewWorkplaceAPI(mocks.NewWorkplaceDbRepoMock())

	var req = &pb.ListWorkplacesV1Request{}
	resp, err := workplaceAPI.ListWorkplacesV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")
}

func TestRemoveWorkplaceV1(t *testing.T) {
	t.Parallel()
	workplaceAPI := NewWorkplaceAPI(mocks.NewWorkplaceDbRepoMock())

	req := &pb.RemoveWorkplaceV1Request{WorkplaceId: 1}
	resp, err := workplaceAPI.RemoveWorkplaceV1(context.Background(), req)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "Response must not be nil")
}