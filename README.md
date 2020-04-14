# traQ (Project R)

## 環境構築

### 必要なソフト

- go 1.13.x
- git
- make
- docker
- docker-compose

### Dockerを使う

####　初回設定
`make update-frontend && make up`

アクセスができるようになってるので、
+ `http://localhost:3000` for traQ
    + アカウント名: `traq`
    + パスワード: `traq`
+ `http://localhost:3001` データベース閲覧用システム　traq_adminer_1
+ `http://localhost:6060` for traQ pprof web interface よくわからない
+ `3002/tcp` for traQ MariaDB（3001も同じ
    + username: `root`
    + password: `password`
    + database: `traq`

#### 再構築
`make up`

####　フロントエンド
`make update-frontend`

#### コンテナを初期化
`make down`

### Testing
1. Run mysql container for test by `make up-test-db`
2. `make test`

You can remove the container by `make rm-test-db`

### Code Lint
`make lint` (or individually `make golangci-lint`, `make swagger-lint`)

Installing below tools in advance is required:
+ [golangci-lint](https://github.com/golangci/golangci-lint) for go codes
+ [spectral](https://github.com/stoplightio/spectral) for swagger specs

### Generate DB Schema Docs
[tbls](https://github.com/k1LoW/tbls) is required.

`make db-gen-docs`

Test mysql container need to be running by `make up-test-db`.

#### DB Docs Lint
`make db-lint`

## License
Code licensed under [the MIT License](https://github.com/traPtitech/traQ/blob/master/LICENSE).

[twemoji](https://twemoji.twitter.com) (svg files in `/dev/data/twemoji`) by 2018 Twitter, Inc and other contributors is licensed under [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/). 

## api
### ログイン、ログアウト
/login,/logout
`{
  "name": "string",
  "pass": "string"
}`

メソッド：POST

## ユーザー登録
/users
`{
  "name": "string",
  "password": "string"
}`
重要:パスワードは10文字以上ないとはじかれる
（traPortalがないのでクライアントを作らんとな）
メソッド：PUT

