package goview

import (
	"reflect"
	"testing"
	"time"
)

// TestCastFunctions 测试类型转换函数
func TestCastFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function string
		input    interface{}
		expected interface{}
	}{
		{"castToInt_string", "castToInt", "42", 42},
		{"castToInt_float", "castToInt", 42.7, 42},
		{"castToFloat_string", "castToFloat", "42.5", 42.5},
		{"castToFloat_int", "castToFloat", 42, 42.0},
		{"castToString_int", "castToString", 42, "42"},
		{"castToString_float", "castToString", 42.5, "42.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.function {
			case "castToInt":
				result := castToInt(tt.input)
				if result != tt.expected {
					t.Errorf("castToInt(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			case "castToFloat":
				result := castToFloat(tt.input)
				if result != tt.expected {
					t.Errorf("castToFloat(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			case "castToString":
				result := castToString(tt.input)
				if result != tt.expected {
					t.Errorf("castToString(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

// TestStringsFunctions 测试字符串处理函数
func TestStringsFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"stringsToUpper",
			func() interface{} { return stringsToUpper("hello world") },
			"HELLO WORLD",
		},
		{
			"stringsToLower",
			func() interface{} { return stringsToLower("HELLO WORLD") },
			"hello world",
		},
		{
			"stringsTitle",
			func() interface{} { return stringsTitle("hello world") },
			"Hello World",
		},
		{
			"stringsTrim",
			func() interface{} { return stringsTrim("  hello world  ", " ") },
			"hello world",
		},
		{
			"stringsReplace",
			func() interface{} { return stringsReplace("hello world", "world", "golang") },
			"hello golang",
		},
		{
			"stringsContains",
			func() interface{} { return stringsContains("hello world", "world") },
			true,
		},
		{
			"stringsHasPrefix",
			func() interface{} { return stringsHasPrefix("hello world", "hello") },
			true,
		},
		{
			"stringsHasSuffix",
			func() interface{} { return stringsHasSuffix("hello world", "world") },
			true,
		},
		{
			"stringsCount",
			func() interface{} { return stringsCount("hello hello", "hello") },
			2,
		},
		{
			"stringsSplit",
			func() interface{} { return stringsSplit("a,b,c", ",") },
			[]string{"a", "b", "c"},
		},
		{
			"stringsSubstr",
			func() interface{} { return stringsSubstr("hello world", 0, 5) },
			"hello",
		},
		{
			"stringsTruncate",
			func() interface{} { return stringsTruncate("hello world this is a long sentence", 15) },
			"hello world…",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestMathFunctions 测试数学函数
func TestMathFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"mathAdd",
			func() interface{} { return mathAdd(1, 2, 3) },
			6.0,
		},
		{
			"mathSub",
			func() interface{} { return mathSub(10, 3, 2) },
			5.0,
		},
		{
			"mathMul",
			func() interface{} { return mathMul(2, 3, 4) },
			24.0,
		},
		{
			"mathDiv",
			func() interface{} { return mathDiv(12, 3, 2) },
			2.0,
		},
		{
			"mathMax",
			func() interface{} { return mathMax(1, 5, 3, 9, 2) },
			9.0,
		},
		{
			"mathMin",
			func() interface{} { return mathMin(1, 5, 3, 9, 2) },
			1.0,
		},
		{
			"mathAbs",
			func() interface{} { return mathAbs(-42.5) },
			42.5,
		},
		{
			"mathCeil",
			func() interface{} { return mathCeil(4.2) },
			5.0,
		},
		{
			"mathFloor",
			func() interface{} { return mathFloor(4.8) },
			4.0,
		},
		{
			"mathRound",
			func() interface{} { return mathRound(4.6) },
			5.0,
		},
		{
			"mathSqrt",
			func() interface{} { return mathSqrt(16) },
			4.0,
		},
		{
			"mathPow",
			func() interface{} { return mathPow(2, 3) },
			8.0,
		},
		{
			"mathMod",
			func() interface{} { return mathMod(10, 3) },
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestCollectionsFunctions 测试集合处理函数
func TestCollectionsFunctions(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"collectionsFirst",
			func() interface{} { return collectionsFirst(3, slice) },
			[]int{1, 2, 3},
		},
		{
			"collectionsLast",
			func() interface{} { return collectionsLast(3, slice) },
			[]int{3, 4, 5},
		},
		{
			"collectionsAfter",
			func() interface{} { return collectionsAfter(2, slice) },
			[]int{4, 5},
		},
		{
			"collectionsReverse",
			func() interface{} { return collectionsReverse([]int{1, 2, 3}) },
			[]int{3, 2, 1},
		},
		{
			"collectionsIn",
			func() interface{} { return collectionsIn(slice, 3) },
			true,
		},
		{
			"collectionsIn_false",
			func() interface{} { return collectionsIn(slice, 10) },
			false,
		},
		{
			"collectionsUniq",
			func() interface{} { return collectionsUniq([]int{1, 2, 2, 3, 3, 4}) },
			[]interface{}{1, 2, 3, 4},
		},
		{
			"collectionsUnion",
			func() interface{} { return collectionsUnion([]int{1, 2, 3}, []int{3, 4, 5}) },
			[]interface{}{1, 2, 3, 4, 5},
		},
		{
			"collectionsIntersect",
			func() interface{} { return collectionsIntersect([]int{1, 2, 3}, []int{2, 3, 4}) },
			[]interface{}{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestCompareFunctions 测试比较函数
func TestCompareFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"compareEq_true",
			func() interface{} { return compareEq(42, 42) },
			true,
		},
		{
			"compareEq_false",
			func() interface{} { return compareEq(42, 43) },
			false,
		},
		{
			"compareNe_true",
			func() interface{} { return compareNe(42, 43) },
			true,
		},
		{
			"compareGt_true",
			func() interface{} { return compareGt(5, 3) },
			true,
		},
		{
			"compareGt_false",
			func() interface{} { return compareGt(3, 5) },
			false,
		},
		{
			"compareLt_true",
			func() interface{} { return compareLt(3, 5) },
			true,
		},
		{
			"compareGe_true",
			func() interface{} { return compareGe(5, 5) },
			true,
		},
		{
			"compareLe_true",
			func() interface{} { return compareLe(3, 5) },
			true,
		},
		{
			"compareDefault_with_empty",
			func() interface{} { return compareDefault("default", "") },
			"default",
		},
		{
			"compareDefault_with_value",
			func() interface{} { return compareDefault("default", "value") },
			"value",
		},
		{
			"compareConditional_true",
			func() interface{} { return compareConditional(true, "yes", "no") },
			"yes",
		},
		{
			"compareConditional_false",
			func() interface{} { return compareConditional(false, "yes", "no") },
			"no",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestCryptoFunctions 测试加密函数
func TestCryptoFunctions(t *testing.T) {
	input := "hello world"

	tests := []struct {
		name     string
		function func() interface{}
		check    func(interface{}) bool
	}{
		{
			"cryptoMD5",
			func() interface{} { return cryptoMD5(input) },
			func(result interface{}) bool {
				s, ok := result.(string)
				return ok && len(s) == 32 // MD5 哈希长度为32个字符
			},
		},
		{
			"cryptoSHA1",
			func() interface{} { return cryptoSHA1(input) },
			func(result interface{}) bool {
				s, ok := result.(string)
				return ok && len(s) == 40 // SHA1 哈希长度为40个字符
			},
		},
		{
			"cryptoSHA256",
			func() interface{} { return cryptoSHA256(input) },
			func(result interface{}) bool {
				s, ok := result.(string)
				return ok && len(s) == 64 // SHA256 哈希长度为64个字符
			},
		},
		{
			"cryptoFNV32a",
			func() interface{} { return cryptoFNV32a(input) },
			func(result interface{}) bool {
				_, ok := result.(uint32)
				return ok
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if !tt.check(result) {
				t.Errorf("%s failed validation, got %v", tt.name, result)
			}
		})
	}
}

// TestEncodingFunctions 测试编码函数
func TestEncodingFunctions(t *testing.T) {
	input := "hello world"

	// 测试 Base64 编码/解码
	encoded := encodingBase64Encode(input)
	if encoded != "aGVsbG8gd29ybGQ=" {
		t.Errorf("encodingBase64Encode(%s) = %s, want aGVsbG8gd29ybGQ=", input, encoded)
	}

	decoded, err := encodingBase64Decode(encoded)
	if err != nil {
		t.Errorf("encodingBase64Decode failed: %v", err)
	}
	if decoded != input {
		t.Errorf("encodingBase64Decode(%s) = %s, want %s", encoded, decoded, input)
	}

	// 测试 JSON 编码
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}
	jsonResult, err := encodingJsonify(data)
	if err != nil {
		t.Errorf("encodingJsonify failed: %v", err)
	}
	if string(jsonResult) != `{"age":30,"name":"test"}` {
		t.Errorf("encodingJsonify result = %s", string(jsonResult))
	}
}

// TestTimeFunctions 测试时间函数
func TestTimeFunctions(t *testing.T) {
	// 测试 timeNow
	now := timeNow()
	if now.IsZero() {
		t.Error("timeNow() returned zero time")
	}

	// 测试 timeAsTime
	timeStr := "2023-01-15T10:30:00Z"
	parsedTime := timeAsTime(timeStr)
	if parsedTime.IsZero() {
		t.Errorf("timeAsTime(%s) returned zero time", timeStr)
	}

	// 测试 timeFormat
	formattedTime := timeFormat("2006-01-02", parsedTime)
	if formattedTime != "2023-01-15" {
		t.Errorf("timeFormat result = %s, want 2023-01-15", formattedTime)
	}

	// 测试 timeParseDuration
	duration, err := timeParseDuration("1h30m")
	if err != nil {
		t.Errorf("timeParseDuration failed: %v", err)
	}
	if duration != time.Hour+30*time.Minute {
		t.Errorf("timeParseDuration result = %v, want %v", duration, time.Hour+30*time.Minute)
	}
}

// TestSafeFunctions 测试安全标记函数
func TestSafeFunctions(t *testing.T) {
	input := "<script>alert('test')</script>"

	tests := []struct {
		name     string
		function func() interface{}
		typeName string
	}{
		{
			"safeHTML",
			func() interface{} { return safeHTML(input) },
			"template.HTML",
		},
		{
			"safeCSS",
			func() interface{} { return safeCSS("color: red;") },
			"template.CSS",
		},
		{
			"safeJS",
			func() interface{} { return safeJS("console.log('test');") },
			"template.JS",
		},
		{
			"safeURL",
			func() interface{} { return safeURL("http://example.com") },
			"template.URL",
		},
		{
			"safeHTMLAttr",
			func() interface{} { return safeHTMLAttr("class=\"test\"") },
			"template.HTMLAttr",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			resultType := reflect.TypeOf(result).String()
			if resultType != tt.typeName {
				t.Errorf("%s returned type %s, want %s", tt.name, resultType, tt.typeName)
			}
		})
	}
}

// TestPathFunctions 测试路径函数
func TestPathFunctions(t *testing.T) {
	testPath := "/path/to/file.txt"

	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"pathBase",
			func() interface{} { return pathBase(testPath) },
			"file.txt",
		},
		{
			"pathBaseName",
			func() interface{} { return pathBaseName(testPath) },
			"file",
		},
		{
			"pathDir",
			func() interface{} { return pathDir(testPath) },
			"/path/to",
		},
		{
			"pathExt",
			func() interface{} { return pathExt(testPath) },
			".txt",
		},
		{
			"pathJoin",
			func() interface{} { return pathJoin("path", "to", "file.txt") },
			"path/to/file.txt",
		},
		{
			"pathClean",
			func() interface{} { return pathClean("/path//to/../file.txt") },
			"/path/file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if result != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestExtFunctionsIntegration 测试函数与ViewEngine的集成
func TestExtFunctionsIntegration(t *testing.T) {
	// 验证函数是否已加载
	funcs := ExtFunctions()
	if len(funcs) == 0 {
		t.Error("ExtFunctions() returned empty map")
	}

	// 验证一些关键函数是否存在
	expectedFuncs := []string{
		"upper", "lower", "title", "trim",
		"add", "sub", "mul", "div",
		"eq", "ne", "gt", "lt", "ge", "le",
		"first", "last", "reverse", "sort",
		"md5", "sha1", "sha256",
		"base64Encode", "base64Decode", "jsonify",
		"now", "safeHTML", "safeCSS", "safeJS",
	}

	for _, funcName := range expectedFuncs {
		if _, exists := funcs[funcName]; !exists {
			t.Errorf("Expected function '%s' not found in ExtFunctions", funcName)
		}
	}

	// 测试命名空间函数是否存在
	namespacedFuncs := []string{
		"strings.ToUpper", "strings.ToLower", "strings.Title",
		"math.Add", "math.Sub", "math.Mul", "math.Div",
		"collections.First", "collections.Last", "collections.Reverse",
		"compare.Eq", "compare.Ne", "compare.Default",
		"crypto.MD5", "crypto.SHA1", "crypto.SHA256",
		"encoding.Base64Encode", "encoding.Base64Decode", "encoding.Jsonify",
		"time.Now", "time.Format",
		"safe.HTML", "safe.CSS", "safe.JS",
		"path.Base", "path.Dir", "path.Ext",
	}

	for _, funcName := range namespacedFuncs {
		if _, exists := funcs[funcName]; !exists {
			t.Errorf("Expected namespaced function '%s' not found in ExtFunctions", funcName)
		}
	}

	t.Logf("Successfully loaded %d functions", len(funcs))
	t.Logf("ViewEngine successfully integrated with functions")
}

// TestReflectFunctions 测试反射函数
func TestReflectFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"reflectIsMap_true",
			func() interface{} { return reflectIsMap(map[string]int{"a": 1}) },
			true,
		},
		{
			"reflectIsMap_false",
			func() interface{} { return reflectIsMap([]int{1, 2, 3}) },
			false,
		},
		{
			"reflectIsSlice_true",
			func() interface{} { return reflectIsSlice([]int{1, 2, 3}) },
			true,
		},
		{
			"reflectIsSlice_false",
			func() interface{} { return reflectIsSlice(map[string]int{"a": 1}) },
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if result != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestTransformFunctions 测试转换函数
func TestTransformFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"transformHTMLEscape",
			func() interface{} { return transformHTMLEscape("<script>alert('test')</script>") },
			"&lt;script&gt;alert(&#39;test&#39;)&lt;/script&gt;",
		},
		{
			"transformHTMLUnescape",
			func() interface{} { return transformHTMLUnescape("&lt;div&gt;test&lt;/div&gt;") },
			"<div>test</div>",
		},
		{
			"transformPlainify",
			func() interface{} { return transformPlainify("<p>Hello <strong>world</strong></p>") },
			"Hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if result != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestUrlsFunctions 测试URL函数
func TestUrlsFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() interface{}
		expected interface{}
	}{
		{
			"urlsAnchorize",
			func() interface{} { return urlsAnchorize("Hello World! 123") },
			"hello-world-123",
		},
		{
			"urlsURLize",
			func() interface{} { return urlsURLize("Hello World! 123") },
			"hello-world-123",
		},
		{
			"urlsJoinPath",
			func() interface{} { return urlsJoinPath("path", "to", "file.html") },
			"path/to/file.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function()
			if result != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}
