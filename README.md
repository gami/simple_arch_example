# Simple Architecture Template

## What's?

MVCアーキテクチャで作成したシンプルなWEBアプリケーションのサンプルです。
ORM（SQLBoilerの構造体）をそのままEntityとして利用します。

### pros 
- レイヤーが少なく、実装量が少ない

### cons
- model層が肥大する
- RDBのレコードがコントローラーまで露出する

## ディレクトリ構成

```
├── model
│   ├── testdata
│   │   └── fixture ... DBテスト用
│   └── user.go ... モデル層。SQLBoiler経由でDBアクセスし、適切な処理をおこなう。データモデルはSQLBoilerが生成した構造体を使う
├── Makefile
├── README.md
├── client ... AWS/HTTP/Slack/Twitterなどの外部サービスのAPIを呼び出す
├── cmd
│   ├── server ... APIサーバエントリポイント
│   └── worker ... Workerエントリポイント
├── config
│   ├── config.go ... 設定
│   └── env 環境依存ファイル
│       └── *.toml
├── controller ... WEBサーバのHandler/Router
│   ├── controller.go ... Controllerのベース構造体（Handlerを単純にマージする）
│   ├── rooter.go ... Chi Rooter/ Middleware
│   └── user.go ... Handler
├── di ... DI
├── docker ... Docker関連
├── gen
│   ├── openapi ... oapi-codegenの生成済ファイル
│   └── schema　... SqlBoilerの生成済ファイル
├── go.mod
├── mysql
│   └── db.go ... DB接続
├── pkg　... ユーティリティ関数群
├── resources
│   ├── openapi ... OpenAPIスキーマ
│   └── sql ... DBスキーマ
└── scripts ... スクリプト
```

## 環境構築

```
make setup
# migration
docker compose -f docker/docker-compose.local.yaml up
```
