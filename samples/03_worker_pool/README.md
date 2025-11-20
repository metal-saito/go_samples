# Worker Pool Pattern

goroutine と channel を使用したワーカープールパターンの実装です。

## 技術要素

- **goroutine**: 並行処理の実現
- **channel**: goroutine間の通信
- **context**: キャンセレーションとタイムアウト管理
- **sync.WaitGroup**: goroutineの完了待機
- **エラーハンドリング**: エラー収集とリトライロジック

## セットアップ

```bash
cd samples/03_worker_pool
go mod download
```

## 使い方

### 実行

```bash
go run main.go
```

デフォルトで3つのワーカーが起動し、10個のジョブを処理します。

### カスタマイズ

```bash
go run main.go -workers 5 -jobs 20
```

## テスト

```bash
go test -v ./...
```

## 設計ポイント

- ワーカープールパターンで並行処理を効率的に管理
- context.Context を使用したキャンセレーションとタイムアウト
- channel による安全なデータ共有
- エラーハンドリングとリトライロジック
- メトリクス収集（処理時間、成功率など）

