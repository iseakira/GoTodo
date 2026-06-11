## Serverの起動

go run ./cmd/server/

### マイグレーション

初回のみ、マイグレーション履歴テーブルを作成する。

```bash
go run ./cmd/migrate init
```

マイグレーションを適用してテーブルを作成する。

```bash
go run ./cmd/migrate up
```

### サーバー起動

```bash
go run ./cmd/server
```

`http://localhost:8989` で起動する。

## マイグレーション

スキーマの変更は、テーブルを直接いじらずマイグレーションで管理する。

### コマンド一覧

| コマンド                      | 内容                                |
| ----------------------------- | ----------------------------------- |
| `go run ./cmd/migrate init`   | 履歴テーブルを作成（初回のみ）      |
| `go run ./cmd/migrate up`     | 未適用のマイグレーションを適用      |
| `go run ./cmd/migrate down`   | 直近のマイグレーションを1つ取り消す |
| `go run ./cmd/migrate status` | 適用状況を確認                      |

### 新しいマイグレーションの追加

`migrations/` にファイルを追加する。ファイル名の先頭をタイムスタンプにし、適用順を決める。
