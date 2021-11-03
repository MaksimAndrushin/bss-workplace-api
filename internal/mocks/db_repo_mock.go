package mocks

import (
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	"golang.org/x/net/context"
)

type WorkplaceDbRepoMock struct {
}

func NewWorkplaceDbRepoMock() repo.Repo{
	return &WorkplaceDbRepoMock{}
}

func (r *WorkplaceDbRepoMock) CreateWorkplace(ctx context.Context, foo string) (uint64, error) {
	return 0, nil
}

func (r *WorkplaceDbRepoMock) DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error) {
	return nil, nil
}

func (r *WorkplaceDbRepoMock) ListWorkplaces(ctx context.Context) ([]model.Workplace, error) {
	return nil, nil
}

func (r *WorkplaceDbRepoMock) RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error) {
	return false, nil
}