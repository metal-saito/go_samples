# CLI Reservation Tool

コマンドラインから予約を登録・一覧表示・キャンセルするCLIツールです。

## 技術要素

- **`cobra`**: コマンドライン引数の解析とサブコマンド管理
- **JSON 永続化**: 予約データをJSONファイルで保存
- **単体テスト**: `testing` パッケージを使用したテスト

## セットアップ

```bash
cd samples/01_cli_reservation
go mod download
```

## 使い方

### 予約を追加

```bash
go run main.go add --name "Alice" --resource "Room-A" --starts-at "2025-01-02T09:00:00Z" --ends-at "2025-01-02T10:00:00Z"
```

### 予約一覧を表示

```bash
go run main.go list
```

### 予約をキャンセル

```bash
go run main.go cancel RES-0001
```

## ビルド

```bash
go build -o reserve main.go
./reserve add --name "Alice" --resource "Room-A" --starts-at "2025-01-02T09:00:00Z" --ends-at "2025-01-02T10:00:00Z"
```

## テスト

```bash
go test -v ./...
```

## 設計ポイント

- `cobra.Command` でサブコマンドを定義し、コマンド構造を明確化
- ドメインモデル（`Reservation`）と永続化層（`Store`）を分離
- エラーハンドリングを適切に行い、ユーザーフレンドリーなメッセージを表示

