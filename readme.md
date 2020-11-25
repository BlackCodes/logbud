## Logbud

Logbud在每个日志打印的地方，都会自动加上文件名和行号，所加的文件名与行号，对于源代码都没有任何入侵；

Logbud现在能解决的问题，每类日志包不都可以加上文件与所在行的信息吗？为什么还需要通过logbud多此一举呢？
的确颇有道理，但是我们在追求一些高性能时，常见的日志包就不能满足我们的需求了；

## 测试

我们现在用zerolog，logrus和logbud做性能参照测试

```json
BenchmarkLogbud-12               4215670               282 ns/op              64 B/op          1 allocs/op
BenchmarkLogrus-12                220213              5511 ns/op            1506 B/op         29 allocs/op
BenchmarkZeroUnCaller-12         3565722               340 ns/op              64 B/op          1 allocs/op
BenchmarkZeroCaller-12            901479              1355 ns/op             288 B/op          5 allocs/op
```

- BenchmarkLogbud是经过logbud编辑后的压测，基本接近原生；

- BenchmarkZeroUnCaller 使用了zerolog并不打印文件与行号信息，发现与原生非常提接近了；
- BenchmarkZeroCaller zerolog调用了运行时的栈，收集了行号与文件信息，发现与logbud相差接近5倍的性能；
- BenchmarkLogrus logrus调用了运行时的栈，收集了行号与文件信息，与zerologCaller相差接近4倍，与logbud相差接近20倍性能；

通过以上的性能测试数据，如果对性能要求足够高，logbudg有相当的优势；

## 使用

```shell
#编译logbud
git clone xx
cd logbud
go build
cp logud /usr/local/bin
# 运行项目
cd go_project
# logbud -h 查看使用参数
logbud -bagrs="-o go_binary" -pos=tail -pmod=file -cp=true
```

参数说明:
- -bargs: go编译时所用的参数；
- -pos: 文件行信息插入的位置，默认在头部，参数有header,tail
- -pmod: 所涉及的文件路径选项，full:全路径，relative:相对路径（默认值），file:仅路文件
- -cp: 是否对所编译的Go二进制文件进行压缩，默认否，当前仅支持Mac版本的进行压缩；
