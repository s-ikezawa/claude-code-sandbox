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

**【重要】ソフトウェアエンジニアリングタスクを開始する前に、必ず [docs/development-workflow.md](./docs/development-workflow.md) を最初に読むこと**

### 開発開始時の必須チェックリスト
- [ ] **docs/development-workflow.md を精読**（最優先）
- [ ] 仕様理解漏れ防止チェックリストの実行
- [ ] TODOリスト作成（要件ベース）
- [ ] ブランチ作成とTDD実装開始

### development-workflow.md確認忘れ防止策
1. **タスク開始時**: いかなるコーディング作業も、development-workflow.mdの確認なしには開始しない
2. **TodoWrite実行時**: development-workflow.mdで定義された必須項目（ブランチ作成、仕様書の具体例テスト、プルリクエスト作成等）が含まれているかを確認
3. **実装前チェック**: "理解 → 計画 → 実装" の各段階で、development-workflow.mdの該当箇所を参照


