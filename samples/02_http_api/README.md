# HTTP API Reservation Tool

Gin フレームワークを使用した RESTful API です。

## 技術要素

- **`gin`**: 高速なHTTPフレームワーク
- **レイヤードアーキテクチャ**: Handler / Service / Repository の分離
- **ミドルウェア**: ロギング、エラーハンドリング
- **JSON API**: 標準的なRESTful API設計

## セットアップ

```bash
cd samples/02_http_api
go mod download
```

## 使い方

### サーバー起動

```bash
go run main.go
```

デフォルトで `:8080` で起動します。

### API エンドポイント

#### 予約を作成

```bash
curl -X POST http://localhost:8080/reservations \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "Alice",
    "resource_name": "Room-A",
    "starts_at": "2025-01-02T09:00:00Z",
    "ends_at": "2025-01-02T10:00:00Z"
  }'
```

#### 予約一覧を取得

```bash
curl http://localhost:8080/reservations
```

#### 予約をキャンセル

```bash
curl -X DELETE http://localhost:8080/reservations/RES-0001
```

## テスト

```bash
go test -v ./...
```

## 設計ポイント

- Handler層でHTTPリクエスト/レスポンスを処理
- Service層でビジネスロジックを実装
- Repository層でデータ永続化を担当
- エラーハンドリングミドルウェアで統一的なエラー応答
- コンテキストを使用したリクエストのライフサイクル管理

