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
docker exec app gotestsum ./target/unit/test/directory
```

- 特定のメソッド（頭に Test とついているメソッドをテスト実施する）

```shell
cd /path/to/project
docker exec app gotestsum ./target/unit/test/directory -run {メソッド名}
```
