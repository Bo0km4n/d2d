## Charger
- Kawabe

## Setup

go, dockerのインストール
既に済んでいる場合はskip
- `brew install go`
- `brew cask install docker`

### 依存ライブラリインストール

- `go get -u github.com/golang/dep/cmd/dep`
- `cd infra &&  dep ensure`

## collectorサーバ起動

- `cd infra/docker && docker-compose up -d`
- `cd ../ && go run collector/main.go`