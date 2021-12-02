package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonmp/bss-workplace-api/internal/repo WorkplaceEventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/bss-workplace-api/internal/app/sender EventSender
//go:generate mockgen -destination=./mocks/db_repo_mock.go -package=mocks github.com/ozonmp/bss-workplace-api/internal/repo WorkplaceRepo
//go:generate mockgen -destination=./mocks/workplace_service_mock.go -package=mocks github.com/ozonmp/bss-workplace-api/internal/service WorkplaceService
