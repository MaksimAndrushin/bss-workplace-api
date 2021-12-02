module github.com/ozonmp/omp-bot

go 1.17

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/joho/godotenv v1.4.0
	github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api v0.0.0-00010101000000-000000000000
	github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.26.0
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210816183151-1e6c022a8912 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211129164237-f09f9a12af12 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api => ./pkg/bss-workplace-api

replace github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade => ./pkg/bss-workplace-facade
