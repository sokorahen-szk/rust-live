## サイト名

Rust Live 　（仮）

## 要な機能

- 配信中一覧表示

Rust を配信している各サイトの情報を表示。表示したい情報として、

```
- 配信者名
- 視聴者数
- 配信時間　何分経過しているか、など。
- 配信サムネイル
- 配信URL
- 配信タイトル
```

件数が多い場合は、スクロールしたら自動で表示するようにする。

- アーカイブ機能

動画の配信が終わった時点で、リストとして配信履歴を見れるようにしたい。

配信履歴で表示したい情報として、

```
- 配信者名
- 配信開始時間
- 配信終了時間
- 配信サムネイル
- 配信URL
- 配信タイトル
```

検索の機能が必要。検索は　配信者名、配信開始時間〜配信終了時間、タイトル

どのサイトで配信しているか（Youtube,Twitch などのタグ検索）

## 採用する技術

- go + gRPC
- React
- Redis（配信中一覧管理）
- Postgres（配信履歴管理）
- CircleCI
- heroku

## 動画取得に関して

動画配信サイト側が公開している API があれば、それを使用。

公開 API がない場合、PHP 実装にてスクレイピングでデータを取得する。

1 分毎にバッチにて、動画配信サイトからデータを取得し、

Redis および、Postgres にデータを格納する。
