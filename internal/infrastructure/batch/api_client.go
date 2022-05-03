package batch

type ApiClient interface {
	// APIのシークレットキーやアクセストークンをヘッダーや
	// リクエストパラメーターに埋め込む設定を行う。
	SetAuth()
}
