# ベースイメージとしてGoを使用
FROM golang:1.20-alpine

# ワーキングディレクトリを設定
WORKDIR /app

# Goのモジュールファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o myapp

# アプリケーションを実行
CMD ["./myapp"]
