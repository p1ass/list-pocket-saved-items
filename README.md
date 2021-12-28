# list-pocket-saved-item

2021年にPocketに保存した記事をマークダウン形式で出力するツール

## Usage

```shell
export POCKET_API_CONSUMER_KEY=xxxxx-xxxxxxxxxxxxxxxxxxxxxx
export POCKET_API_ACCESS_KEY=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxx

go run main.go

### {{TagName}}
- [{{Title}}]({{URL}})
- [{{Title}}]({{URL}})
- [{{Title}}]({{URL}})
...
```