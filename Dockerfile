# ベースとなるDockerイメージ指定
FROM golang:1.17-alpine

# コンテナ内に作業ディレクトリを作成
WORKDIR /app

# go.mod と go.sum を先にコピー
COPY app/go.mod app/go.sum ./

# 依存関係をダウンロード
RUN go mod download

# その他のファイルをコンテナの作業ディレクトリにコピー
COPY app/ .

# Goをビルド
RUN go build -o main .

# ポート9090を公開
EXPOSE 9090

# コンテナを起動する際に実行されるコマンド
CMD [ "./main" ]
