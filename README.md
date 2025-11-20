# Go Samples

このリポジトリは、予約管理システムを題材にした Go 言語の技術サンプル集です。  

## サンプル一覧

| ディレクトリ | テーマ | 主な技術要素 |
| --- | --- | --- |
| `samples/01_cli_reservation` | CLI での予約登録・一覧表示 | `cobra` / JSON 永続化 / 単体テスト (`testing`) |
| `samples/02_http_api` | Gin ベースの予約 API | レイヤードアーキテクチャ / `gin` / JSON API / ミドルウェア |
| `samples/03_worker_pool` | ワーカープールによる非同期処理 | `goroutine` / `channel` / コンテキスト管理 / エラーハンドリング |

## 推奨バージョン

- Go 1.21 以上
- 各サンプルは `go.mod` で依存関係を管理

## 使い方

1. 各サンプルディレクトリに移動します。
2. `go mod download` で依存関係を取得します。
3. `README.md` の手順に従ってアプリケーションまたはテストを実行します。




