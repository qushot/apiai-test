# apiai-test
簡単な説明書き

## /model/apiai_req.go
- API.AI からのリクエストパース用構造体
- JOSN を golang の構造体に変換するサイトで作成したので整理されていない

## /model/apiai_res.go
### 項目の説明
|項目|説明|
|:--|:--|
|speech|Google Home や Google Assistant で読み上げられる音声文字列|
|displayText|Android などで Google Assistant を使用した際に画面上に表示される文字列|
|data|調査中|
|contextOut|調査中|
|source|調査中|
参考:
- https://api.ai/docs/fulfillment#response
- https://simple.bookbookbot.com/

## /app/hello.go
- 現状、疎通確認用の実装のみ
- リクエストをパースしてログ出力(見づらいので修正の必要あり)
- リクエストの内容に関わらず、Speech, DisplayTextに"I'm API.AI"と返す
- レスポンスヘッダーは"Content-type"を"application/json"に設定する必要あり

