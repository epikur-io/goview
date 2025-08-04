# Functions for GoView

GoView 现在完全支持模板引擎的所有内置函数，提供了与 Hugo 类似的模板体验。

## 简介

GoView 集成了函数，涵盖以下功能类别：

- **字符串处理** (strings)
- **数学运算** (math) 
- **集合操作** (collections)
- **类型转换** (cast)
- **比较操作** (compare)
- **密码学哈希** (crypto)
- **编码解码** (encoding)
- **时间处理** (time)
- **路径操作** (path)
- **URL处理** (urls)
- **安全标记** (safe)
- **转换处理** (transform)
- **反射操作** (reflect)
- **操作系统接口** (os)
- **格式化输出** (fmt)

## 快速开始

```go
package main

import (
    "net/http"
    "github.com/foolin/goview"
)

func main() {
    // 创建默认的ViewEngine，已经包含所有函数
    engine := goview.Default()
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]interface{}{
            "title": "Hello Functions",
            "items": []string{"apple", "banana", "cherry"},
            "count": 42,
        }
        
        // 使用函数的模板
        engine.RenderWriter(w, "index", data)
    })
    
    http.ListenAndServe(":8080", nil)
}
```

## 函数使用示例

### 字符串处理函数

```html
<!-- 模板文件: views/index.html -->
<h1>{{ .title | upper }}</h1>
<p>{{ .title | lower }}</p>
<p>{{ .title | title }}</p>

<!-- 字符串操作 -->
<p>包含测试: {{ contains .title "Foobar" }}</p>
<p>前缀测试: {{ hasPrefix .title "Hello" }}</p>
<p>替换: {{ replace .title "Foobar" "GoView" }}</p>
<p>截取: {{ substr .title 0 5 }}</p>
<p>分割: {{ split .title " " }}</p>

<!-- 使用命名空间语法 -->
<p>{{ strings.ToUpper .title }}</p>
<p>{{ strings.Replace .title "Foobar" "GoView" -1 }}</p>
```

### 数学运算函数

```html
<!-- 数学运算 -->
<p>加法: {{ add 10 20 30 }}</p>
<p>减法: {{ sub 100 30 20 }}</p>
<p>乘法: {{ mul 5 6 7 }}</p>
<p>除法: {{ div 100 5 2 }}</p>
<p>最大值: {{ max 10 20 30 }}</p>
<p>最小值: {{ min 10 20 30 }}</p>

<!-- 使用命名空间语法 -->
<p>{{ math.Add 1 2 3 }}</p>
<p>{{ math.Pow 2 8 }}</p>
<p>{{ math.Sqrt 16 }}</p>
```

### 集合操作函数

```html
<!-- 集合操作 -->
<p>前3个元素: {{ first 3 .items }}</p>
<p>后2个元素: {{ last 2 .items }}</p>
<p>排序: {{ sort .items }}</p>
<p>反转: {{ reverse .items }}</p>
<p>包含检查: {{ in .items "apple" }}</p>

<!-- 使用命名空间语法 -->
<ul>
{{ range collections.First 2 .items }}
    <li>{{ . }}</li>
{{ end }}
</ul>

<!-- 高级集合操作 -->
{{ $list1 := slice "a" "b" "c" }}
{{ $list2 := slice "b" "c" "d" }}
<p>并集: {{ union $list1 $list2 }}</p>
<p>交集: {{ intersect $list1 $list2 }}</p>
<p>去重: {{ uniq (slice "a" "b" "b" "c") }}</p>
```

### 类型转换函数

```html
<!-- 类型转换 -->
<p>转为整数: {{ int "42" }}</p>
<p>转为字符串: {{ string 42 }}</p>
<p>转为浮点数: {{ cast.ToFloat "3.14" }}</p>

<!-- 默认值处理 -->
<p>{{ default "默认标题" .emptyTitle }}</p>
```

### 比较操作函数

```html
<!-- 条件判断 -->
{{ if eq .count 42 }}
    <p>计数等于42</p>
{{ end }}

{{ if gt .count 30 }}
    <p>计数大于30</p>
{{ end }}

{{ if and (ge .count 40) (le .count 50) }}
    <p>计数在40-50之间</p>
{{ end }}

<!-- 使用命名空间语法 -->
{{ if compare.Gt .count 30 }}
    <p>使用命名空间比较</p>
{{ end }}
```

### 时间处理函数

```html
<!-- 时间处理 -->
<p>当前时间: {{ now.Format "2006-01-02 15:04:05" }}</p>
<p>格式化时间: {{ time.Format "January 2, 2006" now }}</p>

<!-- 使用命名空间语法 -->
{{ $time := time.Now }}
<p>{{ time.Format "2006年01月02日" $time }}</p>
```

### 编码和安全函数

```html
<!-- Base64编码 -->
<p>编码: {{ base64Encode "hello world" }}</p>
<p>解码: {{ base64Decode "aGVsbG8gd29ybGQ=" }}</p>

<!-- JSON处理 -->
<script>
var data = {{ jsonify .data }};
</script>

<!-- 安全标记 -->
<div>{{ safeHTML "<strong>这是HTML</strong>" }}</div>
<style>{{ safeCSS "color: red;" }}</style>
<script>{{ safeJS "console.log('safe');" }}</script>

<!-- 哈希函数 -->
<p>MD5: {{ md5 "hello" }}</p>
<p>SHA1: {{ sha1 "hello" }}</p>
<p>SHA256: {{ sha256 "hello" }}</p>
```

### URL和路径处理

```html
<!-- URL处理 -->
<a href="{{ urlize "Hello World Page" }}">链接</a>
<p>锚点ID: {{ anchorize "Section Title" }}</p>

<!-- 路径处理 -->
<p>文件名: {{ path.Base "/path/to/file.txt" }}</p>
<p>目录: {{ path.Dir "/path/to/file.txt" }}</p>
<p>扩展名: {{ path.Ext "/path/to/file.txt" }}</p>
<p>连接路径: {{ path.Join "static" "css" "style.css" }}</p>
```

## 命名空间 vs 简化别名

GoView 提供两种调用函数的方式：

### 1. 命名空间语法（推荐）
```html
{{ strings.ToUpper "hello" }}
{{ math.Add 1 2 3 }}
{{ collections.First 5 .items }}
{{ compare.Eq .value "test" }}
```

### 2. 简化别名
```html
{{ upper "hello" }}
{{ add 1 2 3 }}
{{ first 5 .items }}
{{ eq .value "test" }}
```

## 自定义函数优先级

用户自定义的函数会覆盖同名的函数：

```go
engine := goview.New(goview.Config{
    Root:      "views",
    Extension: ".html",
    Funcs: template.FuncMap{
        // 这个自定义函数会覆盖的upper函数
        "upper": func(s string) string {
            return "CUSTOM: " + strings.ToUpper(s)
        },
    },
})
```

## 完整函数列表

### cast 命名空间
- `cast.ToFloat` - 转换为浮点数
- `cast.ToInt` - 转换为整数  
- `cast.ToString` - 转换为字符串

### collections 命名空间
- `collections.After` - 获取第N个元素之后的元素
- `collections.Append` - 追加元素到切片
- `collections.Apply` - 对集合应用函数
- `collections.Complement` - 获取补集
- `collections.Delimit` - 用分隔符连接
- `collections.Dictionary` - 创建字典
- `collections.First` - 获取前N个元素
- `collections.In` - 检查元素是否在集合中
- `collections.Index` - 根据索引获取元素
- `collections.Intersect` - 获取交集
- `collections.IsSet` - 检查键是否存在
- `collections.Last` - 获取后N个元素
- `collections.Merge` - 合并映射
- `collections.Querify` - 转换为查询字符串
- `collections.Reverse` - 反转集合
- `collections.Seq` - 生成数字序列
- `collections.Shuffle` - 随机打乱
- `collections.Slice` - 创建切片
- `collections.Sort` - 排序
- `collections.Union` - 获取并集
- `collections.Uniq` - 去重
- `collections.Where` - 过滤集合

### compare 命名空间
- `compare.Conditional` - 条件选择
- `compare.Default` - 默认值
- `compare.Eq` - 等于比较
- `compare.Ge` - 大于等于比较
- `compare.Gt` - 大于比较
- `compare.Le` - 小于等于比较
- `compare.Lt` - 小于比较
- `compare.Ne` - 不等于比较

### crypto 命名空间
- `crypto.FNV32a` - FNV32a哈希
- `crypto.MD5` - MD5哈希
- `crypto.SHA1` - SHA1哈希
- `crypto.SHA256` - SHA256哈希

### encoding 命名空间
- `encoding.Base64Decode` - Base64解码
- `encoding.Base64Encode` - Base64编码
- `encoding.Jsonify` - JSON编码

### fmt 命名空间
- `fmt.Print` - 打印
- `fmt.Printf` - 格式化打印
- `fmt.Println` - 打印并换行

### math 命名空间
- `math.Abs` - 绝对值
- `math.Add` - 加法
- `math.Ceil` - 向上取整
- `math.Div` - 除法
- `math.Floor` - 向下取整
- `math.Max` - 最大值
- `math.Min` - 最小值
- `math.Mod` - 取模
- `math.Mul` - 乘法
- `math.Pi` - 圆周率
- `math.Pow` - 乘方
- `math.Rand` - 随机数
- `math.Round` - 四舍五入
- `math.Sqrt` - 平方根
- `math.Sub` - 减法

### os 命名空间
- `os.FileExists` - 检查文件是否存在
- `os.Getenv` - 获取环境变量

### path 命名空间
- `path.Base` - 获取路径的最后部分
- `path.BaseName` - 获取无扩展名的文件名
- `path.Clean` - 清理路径
- `path.Dir` - 获取目录部分
- `path.Ext` - 获取扩展名
- `path.Join` - 连接路径
- `path.Split` - 分割路径

### reflect 命名空间
- `reflect.IsMap` - 检查是否为映射
- `reflect.IsSlice` - 检查是否为切片

### safe 命名空间
- `safe.CSS` - 标记为安全CSS
- `safe.HTML` - 标记为安全HTML
- `safe.HTMLAttr` - 标记为安全HTML属性
- `safe.JS` - 标记为安全JavaScript
- `safe.JSStr` - 标记为安全JavaScript字符串
- `safe.URL` - 标记为安全URL

### strings 命名空间
- `strings.Chomp` - 移除尾随换行符
- `strings.Contains` - 检查是否包含子字符串
- `strings.ContainsAny` - 检查是否包含任意字符
- `strings.ContainsNonSpace` - 检查是否包含非空白字符
- `strings.Count` - 计算子字符串出现次数
- `strings.CountRunes` - 计算符文数量（不含空白）
- `strings.CountWords` - 计算单词数量
- `strings.FindRE` - 正则表达式查找
- `strings.FirstUpper` - 首字母大写
- `strings.HasPrefix` - 检查前缀
- `strings.HasSuffix` - 检查后缀
- `strings.Repeat` - 重复字符串
- `strings.Replace` - 替换字符串
- `strings.ReplaceRE` - 正则表达式替换
- `strings.RuneCount` - 计算符文数量
- `strings.SliceString` - 切割字符串
- `strings.Split` - 分割字符串
- `strings.Substr` - 子字符串
- `strings.Title` - 标题格式
- `strings.ToLower` - 转为小写
- `strings.ToUpper` - 转为大写
- `strings.Trim` - 修剪字符
- `strings.TrimLeft` - 修剪左侧字符
- `strings.TrimPrefix` - 移除前缀
- `strings.TrimRight` - 修剪右侧字符
- `strings.TrimSpace` - 修剪空白字符
- `strings.TrimSuffix` - 移除后缀
- `strings.Truncate` - 截断字符串

### time 命名空间
- `time.AsTime` - 转换为时间类型
- `time.Format` - 格式化时间
- `time.Now` - 当前时间
- `time.ParseDuration` - 解析时间间隔

### transform 命名空间
- `transform.HTMLEscape` - HTML转义
- `transform.HTMLUnescape` - HTML反转义
- `transform.Markdownify` - Markdown转HTML
- `transform.Plainify` - 移除HTML标签

### urls 命名空间
- `urls.AbsURL` - 绝对URL
- `urls.Anchorize` - 锚点化
- `urls.JoinPath` - 连接URL路径
- `urls.Parse` - 解析URL
- `urls.RelURL` - 相对URL
- `urls.URLize` - URL化

## 注意事项

1. **性能考虑**: 函数在ViewEngine初始化时加载，不会影响运行时性能
2. **类型安全**: 所有函数都进行了类型检查和错误处理
3. **兼容性**: 完全兼容Hugo模板语法
4. **扩展性**: 可以添加自定义函数并与函数混合使用

## 更多示例

查看 `_examples/` 目录中的完整示例，了解如何在实际项目中使用函数。

## 测试

运行测试确保所有函数正常工作：

```bash
go test -v -run TestExtFunctions
```

## 贡献

欢迎提交 Issue 和 Pull Request 来改进函数的实现。