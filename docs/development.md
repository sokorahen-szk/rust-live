## 開発資料

### テスト

- すべてテスト

```shell
cd /path/to/project
make test
```

- 特定のディレクトリのみテスト

```shell
cd /path/to/project
docker exec app gotestsum -- -count=1 ./target/unit/test/directory
```

- 特定のメソッド（頭に Test とついているメソッドをテスト実施する）

```shell
cd /path/to/project
docker exec app gotestsum -- -count=1 ./target/unit/test/directory -run {メソッド名}
```

### mock

```shell
mockgen -source=./internal/path/to/xxxx.go -destination=./tests/path/to/mock_xxxx.go
```

### grpc

- リスト

```shell
$ grpcurl -plaintext localhost:9000 list
```
