# goバージョン
FROM golang:1.16.3-alpine
# アップデートとgitのインストール！！
RUN apk add --update &&  apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

RUN go mod init quickstart
RUN go get google.golang.org/api/gmail/v1
RUN go get golang.org/x/oauth2/google
