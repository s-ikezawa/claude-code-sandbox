# claude-code-sandbox

Claude Code の実験場

## ゴール

### その１ テスト駆動開発で実装していくことができるようにする

`Go` で `TDD` を実践しながら `Vibe Coding` していくことをゴールとする。

## 開発環境セットアップ

### 必要要件
- Go v1.24.4以上
- Git

### 推奨ツール（リファクタリング対象検出用）

以下のツールをインストールすることで、コード品質の自動チェックとリファクタリング対象の効率的な特定が可能になります：

#### コード重複検出ツール
```bash
# dupl - コード重複検出
go install github.com/mibk/dupl@latest
```

#### 複雑度測定ツール
```bash
# gocyclo - 循環的複雑度測定
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

#### 静的解析ツール
```bash
# golangci-lint - 包括的な静的解析
# macOS (Homebrew)
brew install golangci-lint

# Linux/Windows
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
```

### ツールの使用方法

#### 基本的な使用例
```bash
# プロジェクトルートで実行

# コード重複検出
dupl -plumbing -threshold 50 ./src/

# 複雑度チェック
gocyclo -over 10 ./src/

# 包括的コード品質チェック
golangci-lint run ./src/
```

#### リファクタリング対象の特定
```bash
# 変更されたファイルの確認
git diff --name-only main...HEAD

# 重複の可能性がある関数名の検出
find ./src/ -name "*.go" -exec grep -h "^func " {} \; | \
    sed 's/func \([^(]*\).*/\1/' | sort | uniq -c | sort -nr
```

### 開発ワークフロー

詳細な開発手順とリファクタリング指針については [CLAUDE.md](./CLAUDE.md) を参照してください。

