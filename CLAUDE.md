# CLAUDE.md

## 重要な制約事項

- **言語**: ユーザーとのやりとり、プログラム中のコメント、コミットメッセージ、Pull Requestの内容は日本語で記述する
- **CLIコマンド**: フルパスを利用して実行する

## プロジェクト構成

### フォルダ構造
```text
root/
├── src/                      # プログラミングコードを実装するフォルダ
├── docs/                     # ドキュメントが格納されているフォルダ
│   ├── specification.md      # 仕様書
│   ├── test-guidelines.md    # テストガイドライン（必読）
│   └── development-workflow.md # 開発ワークフロー
├── CLAUDE.md                 # ClaudeCodeへの指示が書かれているファイル
└── README.md                 # このプロジェクトのREADME
```

### 技術構成
- **プログラミング言語**: Go v1.24.4
- **静的解析**: golangci-lint (`golangci-lint run ./src/`)

## 開発ワークフロー

詳細な開発ワークフローについては [docs/development-workflow.md](./docs/development-workflow.md) を参照してください。


