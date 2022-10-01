<h1 align="center">Lang LK</h1>

<p align="center">
    <img alt="badge-lang" src="https://badgen.net/badge/Lang/LK/cyan">
    <img alt="badge-lang" src="https://badgen.net/badge/Lk/0.1.1/blue">
    <img alt="badge-lang" src="https://badgen.net/badge/Go/1.19/purple">
</p>

## 🌳 生态
- Vscode插件：[高亮](https://git.lolli.tech/lollipopkit/vscode-lang-lk-highlight)

## ⌨️ 体验
可以前往 [Release](https://github.com/LollipopKit/lang-lk/releases) 下载 `LK CLI`，或使用`go build .`生成。
#### REPL交互解释器
`./lk`

#### 运行`.lk`
```bash
./lk <file>
# 如果修改了.lk文件导致无法运行，可以尝试添加-f参数
./lk -f <file>
```

## 📄 语法
### 详细语法
- [LANG.md](LANG.md)
- [测试集](test)

### 速览
```js
// 发送请求
shy resp, err = http.post(
    'http://httpbin.org/post', 
    {'accept': 'application/json'}, 
    '{"foo": "bar"}'
)
if err != nil {
    error(err)
}
print(resp.code, resp.body)

// json解析
if json.get(resp.body, 'json.foo') != 'bar' {
    error('mismatch result')
}

// 设置metatable
shy headers = {}
headers.__str = fn(a) {
    shy s = ''
    for k, v in a {
        shy ss = ''
        for _, vv in v {
            ss = ss .. vv .. ';'
        }
        s = s .. k .. ': ' .. ss .. '\n'
    }
    rt s
}


/*
处理监听事件
`req`包含属性`method`, `url`, `body`, `headers`
*/
shy fn handle(req) {
    setmetatable(req.headers, headers)
    rt 200, fmt('%s %s\n\n%s\n%s', req.method, req.url, req.headers, req.body)
}

// 监听
if http.listen(':8080', handle) != nil {
    error(err)
}
```

## 🔖 TODO
- 语法
  - [x] 注释：`//` `/* */`
  - [x] 去除 `repeat` `until` `goto`
  - [x] Raw String, 使用 ``` ` ``` 包裹字符
  - [x] 支持任意对象拼接( `concat` )，使用语法 `..`
  - [ ] Table
    - [x] key为StringExp，而不是NameExp
    - [x] 构造方式：`=` -> `:`, eg: `{a = 'a'}` -> `{a: 'a'}`
    - [x] 索引从 `0` 开始
    - [ ] 改变 `metatable` 设置方式
- 编译器
  - [x] 自动添加 `range` ( `paris` )
  - [x] 支持 `a++` `a+=b` 等
- CLI
  - [x] 利用HASH，文件无变化不编译
  - [x] 支持传入参数 ( `lk args.lk arg1` -> `.lk`内调用`os.args` )
  - [x] REPL
    - [x] 直接运行 `./lk` 即可进入
    - [x] 支持方向键
    - [ ] 识别代码块，并自动缩进
- [x] 资源
    - [x] 文档
      - [x] `CHANGELOG.md`
      - [x] `LANG.md` 
    - [x] 测试集，位于 `test` 文件夹
    - [x] IDE
      - [x] VSCode高亮  

## 💌 致谢
- Lua
- [luago](https://github.com/zxh0/luago-book)

## 📝 License
`LollipopKit 2022 LGPL-3.0`