package goview

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"html"
	"html/template"
	"math"
	"math/rand"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// ExtFunctions 返回所有模板函数的映射
// 这个函数提供了与模板引擎兼容的所有内置函数
func ExtFunctions() template.FuncMap {
	funcs := template.FuncMap{
		// cast 命名空间 - 类型转换函数
		"cast.ToFloat":  castToFloat,
		"cast.ToInt":    castToInt,
		"cast.ToString": castToString,

		// collections 命名空间 - 集合操作函数
		"collections.After":      collectionsAfter,
		"collections.Append":     collectionsAppend,
		"collections.Apply":      collectionsApply,
		"collections.Complement": collectionsComplement,
		"collections.Delimit":    collectionsDelimit,
		"collections.Dictionary": collectionsDictionary,
		"collections.First":      collectionsFirst,
		"collections.In":         collectionsIn,
		"collections.Index":      collectionsIndex,
		"collections.Intersect":  collectionsIntersect,
		"collections.IsSet":      collectionsIsSet,
		"collections.Last":       collectionsLast,
		"collections.Merge":      collectionsMerge,
		"collections.Querify":    collectionsQuerify,
		"collections.Reverse":    collectionsReverse,
		"collections.Seq":        collectionsSeq,
		"collections.Shuffle":    collectionsShuffle,
		"collections.Slice":      collectionsSlice,
		"collections.Sort":       collectionsSort,
		"collections.Union":      collectionsUnion,
		"collections.Uniq":       collectionsUniq,
		"collections.Where":      collectionsWhere,

		// compare 命名空间 - 比较函数
		"compare.Conditional": compareConditional,
		"compare.Default":     compareDefault,
		"compare.Eq":          compareEq,
		"compare.Ge":          compareGe,
		"compare.Gt":          compareGt,
		"compare.Le":          compareLe,
		"compare.Lt":          compareLt,
		"compare.Ne":          compareNe,

		// crypto 命名空间 - 密码学哈希函数
		"crypto.FNV32a": cryptoFNV32a,
		"crypto.MD5":    cryptoMD5,
		"crypto.SHA1":   cryptoSHA1,
		"crypto.SHA256": cryptoSHA256,

		// encoding 命名空间 - 编码解码函数
		"encoding.Base64Decode": encodingBase64Decode,
		"encoding.Base64Encode": encodingBase64Encode,
		"encoding.Jsonify":      encodingJsonify,

		// fmt 命名空间 - 格式化输出函数
		"fmt.Print":   fmtPrint,
		"fmt.Printf":  fmtPrintf,
		"fmt.Println": fmtPrintln,

		// hash 命名空间 - 非密码学哈希函数
		"hash.FNV32a": hashFNV32a,

		// math 命名空间 - 数学运算函数
		"math.Abs":   mathAbs,
		"math.Add":   mathAdd,
		"math.Ceil":  mathCeil,
		"math.Div":   mathDiv,
		"math.Floor": mathFloor,
		"math.Max":   mathMax,
		"math.Min":   mathMin,
		"math.Mod":   mathMod,
		"math.Mul":   mathMul,
		"math.Pi":    mathPi,
		"math.Pow":   mathPow,
		"math.Rand":  mathRand,
		"math.Round": mathRound,
		"math.Sqrt":  mathSqrt,
		"math.Sub":   mathSub,

		// os 命名空间 - 操作系统接口函数
		"os.FileExists": osFileExists,
		"os.Getenv":     osGetenv,

		// path 命名空间 - 路径操作函数
		"path.Base":     pathBase,
		"path.BaseName": pathBaseName,
		"path.Clean":    pathClean,
		"path.Dir":      pathDir,
		"path.Ext":      pathExt,
		"path.Join":     pathJoin,
		"path.Split":    pathSplit,

		// reflect 命名空间 - 反射函数
		"reflect.IsMap":   reflectIsMap,
		"reflect.IsSlice": reflectIsSlice,

		// safe 命名空间 - 安全标记函数
		"safe.CSS":      safeCSS,
		"safe.HTML":     safeHTML,
		"safe.HTMLAttr": safeHTMLAttr,
		"safe.JS":       safeJS,
		"safe.JSStr":    safeJSStr,
		"safe.URL":      safeURL,

		// strings 命名空间 - 字符串操作函数
		"strings.Chomp":            stringsChomp,
		"strings.Contains":         stringsContains,
		"strings.ContainsAny":      stringsContainsAny,
		"strings.ContainsNonSpace": stringsContainsNonSpace,
		"strings.Count":            stringsCount,
		"strings.CountRunes":       stringsCountRunes,
		"strings.CountWords":       stringsCountWords,
		"strings.FindRE":           stringsFindRE,
		"strings.FirstUpper":       stringsFirstUpper,
		"strings.HasPrefix":        stringsHasPrefix,
		"strings.HasSuffix":        stringsHasSuffix,
		"strings.Repeat":           stringsRepeat,
		"strings.Replace":          stringsReplace,
		"strings.ReplaceRE":        stringsReplaceRE,
		"strings.RuneCount":        stringsRuneCount,
		"strings.SliceString":      stringsSliceString,
		"strings.Split":            stringsSplit,
		"strings.Substr":           stringsSubstr,
		"strings.Title":            stringsTitle,
		"strings.ToLower":          stringsToLower,
		"strings.ToUpper":          stringsToUpper,
		"strings.Trim":             stringsTrim,
		"strings.TrimLeft":         stringsTrimLeft,
		"strings.TrimPrefix":       stringsTrimPrefix,
		"strings.TrimRight":        stringsTrimRight,
		"strings.TrimSpace":        stringsTrimSpace,
		"strings.TrimSuffix":       stringsTrimSuffix,
		"strings.Truncate":         stringsTruncate,

		// time 命名空间 - 时间处理函数
		"time.AsTime":        timeAsTime,
		"time.Format":        timeFormat,
		"time.Now":           timeNow,
		"time.ParseDuration": timeParseDuration,

		// transform 命名空间 - 转换函数
		"transform.HTMLEscape":   transformHTMLEscape,
		"transform.HTMLUnescape": transformHTMLUnescape,
		"transform.Markdownify":  transformMarkdownify,
		"transform.Plainify":     transformPlainify,

		// urls 命名空间 - URL处理函数
		"urls.AbsURL":    urlsAbsURL,
		"urls.Anchorize": urlsAnchorize,
		"urls.JoinPath":  urlsJoinPath,
		"urls.Parse":     urlsParse,
		"urls.RelURL":    urlsRelURL,
		"urls.URLize":    urlsURLize,

		// 兼容性别名 - 保持与模板的兼容性
		"add":          mathAdd,
		"sub":          mathSub,
		"mul":          mathMul,
		"div":          mathDiv,
		"mod":          mathMod,
		"abs":          mathAbs,
		"ceil":         mathCeil,
		"floor":        mathFloor,
		"round":        mathRound,
		"sqrt":         mathSqrt,
		"pow":          mathPow,
		"max":          mathMax,
		"min":          mathMin,
		"after":        collectionsAfter,
		"append":       collectionsAppend,
		"apply":        collectionsApply,
		"base64Decode": encodingBase64Decode,
		"base64Encode": encodingBase64Encode,
		"chomp":        stringsChomp,
		"contains":     stringsContains,
		"countRunes":   stringsCountRunes,
		"countWords":   stringsCountWords,
		"default":      compareDefault,
		"delimit":      collectionsDelimit,
		"dict":         collectionsDictionary,
		"eq":           compareEq,
		"first":        collectionsFirst,
		"ge":           compareGe,
		"gt":           compareGt,
		"hasPrefix":    stringsHasPrefix,
		"hasSuffix":    stringsHasSuffix,
		"htmlEscape":   transformHTMLEscape,
		"htmlUnescape": transformHTMLUnescape,
		"in":           collectionsIn,
		"index":        collectionsIndex,
		"int":          castToInt,
		"intersect":    collectionsIntersect,
		"isSet":        collectionsIsSet,
		"jsonify":      encodingJsonify,
		"last":         collectionsLast,
		"le":           compareLe,
		"lower":        stringsToLower,
		"lt":           compareLt,
		"markdownify":  transformMarkdownify,
		"md5":          cryptoMD5,
		"ne":           compareNe,
		"now":          timeNow,
		"plainify":     transformPlainify,
		"print":        fmtPrint,
		"printf":       fmtPrintf,
		"println":      fmtPrintln,
		"querify":      collectionsQuerify,
		"replace":      stringsReplace,
		"replaceRE":    stringsReplaceRE,
		"reverse":      collectionsReverse,
		"safeCSS":      safeCSS,
		"safeHTML":     safeHTML,
		"safeHTMLAttr": safeHTMLAttr,
		"safeJS":       safeJS,
		"safeURL":      safeURL,
		"seq":          collectionsSeq,
		"sha1":         cryptoSHA1,
		"sha256":       cryptoSHA256,
		"shuffle":      collectionsShuffle,
		"slice":        collectionsSlice,
		"sort":         collectionsSort,
		"split":        stringsSplit,
		"string":       castToString,
		"substr":       stringsSubstr,
		"title":        stringsTitle,
		"trim":         stringsTrim,
		"truncate":     stringsTruncate,
		"union":        collectionsUnion,
		"uniq":         collectionsUniq,
		"upper":        stringsToUpper,
		"urlize":       urlsURLize,
		"where":        collectionsWhere,
	}

	return funcs
}

// ====================
// cast 命名空间函数实现
// ====================

// castToFloat 将值转换为浮点数
// 支持包名cast.ToFloat函数
func castToFloat(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case string:
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f
		}
	}
	return 0
}

// castToInt 将值转换为整数
// 支持包名cast.ToInt函数
func castToInt(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return 0
}

// castToString 将值转换为字符串
// 支持包名cast.ToString函数
func castToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case template.HTML:
		return string(val)
	case template.CSS:
		return string(val)
	case template.JS:
		return string(val)
	case template.URL:
		return string(val)
	case template.HTMLAttr:
		return string(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// ====================
// collections 命名空间函数实现
// ====================

// collectionsAfter 返回数组中第N个元素之后的所有元素
// 支持包名collections.After函数
func collectionsAfter(index int, seq interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return nil
	}

	length := seqv.Len()
	if index >= length || index < 0 {
		return reflect.MakeSlice(seqv.Type(), 0, 0).Interface()
	}

	return seqv.Slice(index+1, length).Interface()
}

// collectionsAppend 将元素追加到切片中
// 支持包名collections.Append函数
func collectionsAppend(seq interface{}, values ...interface{}) interface{} {
	if seq == nil {
		return values
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return seq
	}

	// 创建新的切片
	result := reflect.MakeSlice(seqv.Type(), seqv.Len(), seqv.Len()+len(values))
	reflect.Copy(result, seqv)

	// 追加新值
	for _, v := range values {
		val := reflect.ValueOf(v)
		if val.Type().ConvertibleTo(result.Type().Elem()) {
			result = reflect.Append(result, val.Convert(result.Type().Elem()))
		}
	}

	return result.Interface()
}

// collectionsApply 对集合中的每个元素应用函数
// 支持包名collections.Apply函数
func collectionsApply(seq interface{}, fname string, params ...interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return nil
	}

	// 这里简化实现，实际应用中需要根据fname调用相应函数
	result := reflect.MakeSlice(seqv.Type(), seqv.Len(), seqv.Len())
	for i := 0; i < seqv.Len(); i++ {
		result.Index(i).Set(seqv.Index(i))
	}

	return result.Interface()
}

// collectionsComplement 返回在最后一个集合中但不在其他集合中的元素
// 支持包名collections.Complement函数
func collectionsComplement(seqs ...interface{}) interface{} {
	if len(seqs) < 2 {
		return nil
	}

	last := seqs[len(seqs)-1]
	lastv := reflect.ValueOf(last)
	if lastv.Kind() != reflect.Slice && lastv.Kind() != reflect.Array {
		return nil
	}

	// 收集其他集合中的所有元素
	others := make(map[interface{}]bool)
	for i := 0; i < len(seqs)-1; i++ {
		seq := reflect.ValueOf(seqs[i])
		if seq.Kind() == reflect.Slice || seq.Kind() == reflect.Array {
			for j := 0; j < seq.Len(); j++ {
				others[seq.Index(j).Interface()] = true
			}
		}
	}

	// 找出complement
	var result []interface{}
	for i := 0; i < lastv.Len(); i++ {
		item := lastv.Index(i).Interface()
		if !others[item] {
			result = append(result, item)
		}
	}

	return result
}

// collectionsDelimit 用分隔符连接数组、切片或映射中的所有值
// 支持包名collections.Delimit函数
func collectionsDelimit(seq interface{}, delimiter string, last ...string) string {
	if seq == nil {
		return ""
	}

	seqv := reflect.ValueOf(seq)
	var strs []string

	switch seqv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < seqv.Len(); i++ {
			strs = append(strs, castToString(seqv.Index(i).Interface()))
		}
	case reflect.Map:
		for _, key := range seqv.MapKeys() {
			strs = append(strs, castToString(seqv.MapIndex(key).Interface()))
		}
	default:
		return castToString(seq)
	}

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// 处理最后一个元素的分隔符
	lastDelim := delimiter
	if len(last) > 0 {
		lastDelim = last[0]
	}

	if len(strs) == 2 {
		return strs[0] + lastDelim + strs[1]
	}

	// 超过两个元素时
	result := strings.Join(strs[:len(strs)-1], delimiter)
	return result + lastDelim + strs[len(strs)-1]
}

// collectionsDictionary 从键值对列表创建字典
// 支持包名collections.Dictionary和dict函数
func collectionsDictionary(values ...interface{}) map[string]interface{} {
	dict := make(map[string]interface{})
	for i := 0; i < len(values); i += 2 {
		if i+1 < len(values) {
			key := castToString(values[i])
			dict[key] = values[i+1]
		}
	}
	return dict
}

// collectionsFirst 返回集合的前N个元素
// 支持包名collections.First函数
func collectionsFirst(limit int, seq interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return nil
	}

	length := seqv.Len()
	if limit > length {
		limit = length
	}
	if limit < 0 {
		limit = 0
	}

	return seqv.Slice(0, limit).Interface()
}

// collectionsIn 检查值是否在集合中
// 支持包名collections.In函数
func collectionsIn(seq interface{}, value interface{}) bool {
	if seq == nil {
		return false
	}

	seqv := reflect.ValueOf(seq)
	switch seqv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < seqv.Len(); i++ {
			if reflect.DeepEqual(seqv.Index(i).Interface(), value) {
				return true
			}
		}
	case reflect.Map:
		return seqv.MapIndex(reflect.ValueOf(value)).IsValid()
	case reflect.String:
		return strings.Contains(seqv.String(), castToString(value))
	}

	return false
}

// collectionsIndex 返回与给定键或索引关联的对象、元素或值
// 支持包名collections.Index函数
func collectionsIndex(seq interface{}, indices ...interface{}) interface{} {
	if seq == nil || len(indices) == 0 {
		return nil
	}

	v := reflect.ValueOf(seq)
	for _, idx := range indices {
		switch v.Kind() {
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(idx))
			if !v.IsValid() {
				return nil
			}
		case reflect.Slice, reflect.Array:
			i := castToInt(idx)
			if i < 0 || i >= v.Len() {
				return nil
			}
			v = v.Index(i)
		default:
			return nil
		}
	}

	return v.Interface()
}

// collectionsIntersect 返回两个数组或切片的共同元素
// 支持包名collections.Intersect函数
func collectionsIntersect(seq1, seq2 interface{}) interface{} {
	if seq1 == nil || seq2 == nil {
		return nil
	}

	seq1v := reflect.ValueOf(seq1)
	seq2v := reflect.ValueOf(seq2)

	if seq1v.Kind() != reflect.Slice && seq1v.Kind() != reflect.Array {
		return nil
	}
	if seq2v.Kind() != reflect.Slice && seq2v.Kind() != reflect.Array {
		return nil
	}

	// 构建第二个序列的查找映射
	lookup := make(map[interface{}]bool)
	for i := 0; i < seq2v.Len(); i++ {
		lookup[seq2v.Index(i).Interface()] = true
	}

	// 找出交集
	var result []interface{}
	seen := make(map[interface{}]bool)
	for i := 0; i < seq1v.Len(); i++ {
		item := seq1v.Index(i).Interface()
		if lookup[item] && !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

// collectionsIsSet 检查键是否在集合中存在
// 支持包名collections.IsSet函数
func collectionsIsSet(seq interface{}, key interface{}) bool {
	if seq == nil {
		return false
	}

	v := reflect.ValueOf(seq)
	switch v.Kind() {
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(key)).IsValid()
	case reflect.Slice, reflect.Array:
		idx := castToInt(key)
		return idx >= 0 && idx < v.Len()
	}

	return false
}

// collectionsLast 返回集合的后N个元素
// 支持包名collections.Last函数
func collectionsLast(limit int, seq interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return nil
	}

	length := seqv.Len()
	if limit > length {
		limit = length
	}
	if limit < 0 {
		limit = 0
	}

	start := length - limit
	return seqv.Slice(start, length).Interface()
}

// collectionsMerge 合并两个或多个映射
// 支持包名collections.Merge函数
func collectionsMerge(maps ...interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, m := range maps {
		if m == nil {
			continue
		}

		mv := reflect.ValueOf(m)
		if mv.Kind() != reflect.Map {
			continue
		}

		for _, key := range mv.MapKeys() {
			keyStr := castToString(key.Interface())
			result[keyStr] = mv.MapIndex(key).Interface()
		}
	}

	return result
}

// collectionsQuerify 将键值对转换为URL查询字符串
// 支持包名collections.Querify函数
func collectionsQuerify(params ...interface{}) string {
	values := url.Values{}

	for i := 0; i < len(params); i += 2 {
		if i+1 < len(params) {
			key := castToString(params[i])
			value := castToString(params[i+1])
			values.Add(key, value)
		}
	}

	return values.Encode()
}

// collectionsReverse 反转集合的顺序
// 支持包名collections.Reverse函数
func collectionsReverse(seq interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return seq
	}

	length := seqv.Len()
	result := reflect.MakeSlice(seqv.Type(), length, length)

	for i := 0; i < length; i++ {
		result.Index(i).Set(seqv.Index(length - 1 - i))
	}

	return result.Interface()
}

// collectionsSeq 创建整数序列
// 支持包名collections.Seq函数
func collectionsSeq(params ...int) []int {
	var start, stop, step int

	switch len(params) {
	case 1:
		start, stop, step = 1, params[0], 1
	case 2:
		start, stop, step = params[0], params[1], 1
	case 3:
		start, stop, step = params[0], params[1], params[2]
	default:
		return nil
	}

	if step == 0 {
		return nil
	}

	var result []int
	if step > 0 {
		for i := start; i <= stop; i += step {
			result = append(result, i)
		}
	} else {
		for i := start; i >= stop; i += step {
			result = append(result, i)
		}
	}

	return result
}

// collectionsShuffle 返回给定数组或切片的随机排列
// 支持包名collections.Shuffle函数
func collectionsShuffle(seq interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return seq
	}

	length := seqv.Len()
	result := reflect.MakeSlice(seqv.Type(), length, length)
	reflect.Copy(result, seqv)

	// Fisher-Yates 洗牌算法
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		temp := result.Index(i).Interface()
		result.Index(i).Set(result.Index(j))
		result.Index(j).Set(reflect.ValueOf(temp))
	}

	return result.Interface()
}

// collectionsSlice 从给定值创建切片
// 支持包名collections.Slice函数
func collectionsSlice(args ...interface{}) []interface{} {
	return args
}

// collectionsSort 对切片、映射和页面集合进行排序
// 支持包名collections.Sort函数
func collectionsSort(seq interface{}, key ...string) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return seq
	}

	length := seqv.Len()
	if length <= 1 {
		return seq
	}

	// 创建索引切片进行排序
	indices := make([]int, length)
	for i := range indices {
		indices[i] = i
	}

	// 简单排序（字符串比较）
	sort.Slice(indices, func(i, j int) bool {
		val1 := castToString(seqv.Index(indices[i]).Interface())
		val2 := castToString(seqv.Index(indices[j]).Interface())
		return val1 < val2
	})

	// 构建排序结果
	result := reflect.MakeSlice(seqv.Type(), length, length)
	for i, idx := range indices {
		result.Index(i).Set(seqv.Index(idx))
	}

	return result.Interface()
}

// collectionsUnion 返回两个数组或切片的并集
// 支持包名collections.Union函数
func collectionsUnion(seq1, seq2 interface{}) interface{} {
	if seq1 == nil {
		return seq2
	}
	if seq2 == nil {
		return seq1
	}

	seq1v := reflect.ValueOf(seq1)
	seq2v := reflect.ValueOf(seq2)

	if seq1v.Kind() != reflect.Slice && seq1v.Kind() != reflect.Array {
		return seq2
	}
	if seq2v.Kind() != reflect.Slice && seq2v.Kind() != reflect.Array {
		return seq1
	}

	// 收集所有唯一元素
	seen := make(map[interface{}]bool)
	var result []interface{}

	// 添加第一个序列的元素
	for i := 0; i < seq1v.Len(); i++ {
		item := seq1v.Index(i).Interface()
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	// 添加第二个序列的元素
	for i := 0; i < seq2v.Len(); i++ {
		item := seq2v.Index(i).Interface()
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

// collectionsUniq 返回给定集合，移除重复元素
// 支持包名collections.Uniq函数
func collectionsUniq(seq interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return seq
	}

	seen := make(map[interface{}]bool)
	var result []interface{}

	for i := 0; i < seqv.Len(); i++ {
		item := seqv.Index(i).Interface()
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

// collectionsWhere 过滤集合，只保留满足比较条件的元素
// 支持包名collections.Where函数
func collectionsWhere(seq interface{}, key, operator string, value interface{}) interface{} {
	if seq == nil {
		return nil
	}

	seqv := reflect.ValueOf(seq)
	if seqv.Kind() != reflect.Slice && seqv.Kind() != reflect.Array {
		return nil
	}

	var result []interface{}

	for i := 0; i < seqv.Len(); i++ {
		item := seqv.Index(i).Interface()

		// 获取字段值
		var fieldValue interface{}
		if key == "." {
			fieldValue = item
		} else {
			itemv := reflect.ValueOf(item)
			if itemv.Kind() == reflect.Map {
				fieldValue = itemv.MapIndex(reflect.ValueOf(key)).Interface()
			} else if itemv.Kind() == reflect.Struct {
				field := itemv.FieldByName(key)
				if field.IsValid() {
					fieldValue = field.Interface()
				}
			}
		}

		// 执行比较
		match := false
		switch operator {
		case "==", "eq":
			match = reflect.DeepEqual(fieldValue, value)
		case "!=", "ne":
			match = !reflect.DeepEqual(fieldValue, value)
		case "<", "lt":
			match = compareValues(fieldValue, value) < 0
		case "<=", "le":
			match = compareValues(fieldValue, value) <= 0
		case ">", "gt":
			match = compareValues(fieldValue, value) > 0
		case ">=", "ge":
			match = compareValues(fieldValue, value) >= 0
		case "in":
			match = collectionsIn(value, fieldValue)
		case "not in":
			match = !collectionsIn(value, fieldValue)
		}

		if match {
			result = append(result, item)
		}
	}

	return result
}

// compareValues 比较两个值，返回-1、0或1
func compareValues(a, b interface{}) int {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	// 尝试数值比较
	if av.Kind() >= reflect.Int && av.Kind() <= reflect.Float64 &&
		bv.Kind() >= reflect.Int && bv.Kind() <= reflect.Float64 {
		af := castToFloat(a)
		bf := castToFloat(b)
		if af < bf {
			return -1
		} else if af > bf {
			return 1
		}
		return 0
	}

	// 字符串比较
	as := castToString(a)
	bs := castToString(b)
	if as < bs {
		return -1
	} else if as > bs {
		return 1
	}
	return 0
}

// ====================
// compare 命名空间函数实现
// ====================

// compareConditional 根据控制参数的值返回两个参数中的一个
// 支持包名compare.Conditional函数
func compareConditional(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// compareDefault 如果第一个参数未设置，则返回第二个参数
// 支持包名compare.Default函数
func compareDefault(def, given interface{}) interface{} {
	if given == nil {
		return def
	}

	// 检查零值
	v := reflect.ValueOf(given)
	switch v.Kind() {
	case reflect.String:
		if v.String() == "" {
			return def
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if v.Len() == 0 {
			return def
		}
	case reflect.Bool:
		if !v.Bool() {
			return def
		}
	}

	return given
}

// compareEq 返回 arg1 == arg2 的布尔真值
// 支持包名compare.Eq函数
func compareEq(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// compareGe 返回 arg1 >= arg2 的布尔真值
// 支持包名compare.Ge函数
func compareGe(a, b interface{}) bool {
	return compareValues(a, b) >= 0
}

// compareGt 返回 arg1 > arg2 的布尔真值
// 支持包名compare.Gt函数
func compareGt(a, b interface{}) bool {
	return compareValues(a, b) > 0
}

// compareLe 返回 arg1 <= arg2 的布尔真值
// 支持包名compare.Le函数
func compareLe(a, b interface{}) bool {
	return compareValues(a, b) <= 0
}

// compareLt 返回 arg1 < arg2 的布尔真值
// 支持包名compare.Lt函数
func compareLt(a, b interface{}) bool {
	return compareValues(a, b) < 0
}

// compareNe 返回 arg1 != arg2 的布尔真值
// 支持包名compare.Ne函数
func compareNe(a, b interface{}) bool {
	return !reflect.DeepEqual(a, b)
}

// ====================
// crypto 命名空间函数实现
// ====================

// cryptoFNV32a 返回给定字符串的32位FNV非密码学哈希
// 支持包名crypto.FNV32a函数
func cryptoFNV32a(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// cryptoMD5 计算输入的MD5哈希值并返回十六进制字符串
// 支持包名crypto.MD5函数
func cryptoMD5(input interface{}) string {
	h := md5.New()
	h.Write([]byte(castToString(input)))
	return hex.EncodeToString(h.Sum(nil))
}

// cryptoSHA1 计算输入的SHA1哈希值并返回十六进制字符串
// 支持包名crypto.SHA1函数
func cryptoSHA1(input interface{}) string {
	h := sha1.New()
	h.Write([]byte(castToString(input)))
	return hex.EncodeToString(h.Sum(nil))
}

// cryptoSHA256 计算输入的SHA256哈希值并返回十六进制字符串
// 支持包名crypto.SHA256函数
func cryptoSHA256(input interface{}) string {
	h := sha256.New()
	h.Write([]byte(castToString(input)))
	return hex.EncodeToString(h.Sum(nil))
}

// ====================
// encoding 命名空间函数实现
// ====================

// encodingBase64Decode 返回给定内容的base64解码
// 支持包名encoding.Base64Decode函数
func encodingBase64Decode(input interface{}) (string, error) {
	data, err := base64.StdEncoding.DecodeString(castToString(input))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// encodingBase64Encode 返回给定内容的base64编码
// 支持包名encoding.Base64Encode函数
func encodingBase64Encode(input interface{}) string {
	return base64.StdEncoding.EncodeToString([]byte(castToString(input)))
}

// encodingJsonify 将给定对象编码为JSON
// 支持包名encoding.Jsonify函数
func encodingJsonify(v interface{}) (template.JS, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.JS(b), nil
}

// ====================
// fmt 命名空间函数实现
// ====================

// fmtPrint 使用标准fmt.Print函数打印参数的默认表示
// 支持包名fmt.Print函数
func fmtPrint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

// fmtPrintf 使用标准fmt.Sprintf函数格式化字符串
// 支持包名fmt.Printf函数
func fmtPrintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// fmtPrintln 使用标准fmt.Print函数打印参数的默认表示并强制换行
// 支持包名fmt.Println函数
func fmtPrintln(args ...interface{}) string {
	return fmt.Sprintln(args...)
}

// ====================
// hash 命名空间函数实现
// ====================

// hashFNV32a 返回给定字符串的32位FNV非密码学哈希
// 支持包名hash.FNV32a函数
func hashFNV32a(s string) uint32 {
	return cryptoFNV32a(s)
}

// ====================
// math 命名空间函数实现
// ====================

// mathAbs 返回给定数字的绝对值
// 支持包名math.Abs函数
func mathAbs(n interface{}) float64 {
	return math.Abs(castToFloat(n))
}

// mathAdd 将两个或多个数字相加
// 支持包名math.Add函数
func mathAdd(args ...interface{}) float64 {
	var result float64
	for _, arg := range args {
		result += castToFloat(arg)
	}
	return result
}

// mathCeil 返回大于或等于给定数字的最小整数值
// 支持包名math.Ceil函数
func mathCeil(n interface{}) float64 {
	return math.Ceil(castToFloat(n))
}

// mathDiv 将第一个数字除以一个或多个数字
// 支持包名math.Div函数
func mathDiv(args ...interface{}) float64 {
	if len(args) == 0 {
		return 0
	}

	result := castToFloat(args[0])
	for i := 1; i < len(args); i++ {
		divisor := castToFloat(args[i])
		if divisor != 0 {
			result /= divisor
		}
	}
	return result
}

// mathFloor 返回小于或等于给定数字的最大整数值
// 支持包名math.Floor函数
func mathFloor(n interface{}) float64 {
	return math.Floor(castToFloat(n))
}

// mathMax 返回所有数字中的最大值
// 支持包名math.Max函数
func mathMax(args ...interface{}) float64 {
	if len(args) == 0 {
		return 0
	}

	max := castToFloat(args[0])
	for i := 1; i < len(args); i++ {
		val := castToFloat(args[i])
		if val > max {
			max = val
		}
	}
	return max
}

// mathMin 返回所有数字中的最小值
// 支持包名math.Min函数
func mathMin(args ...interface{}) float64 {
	if len(args) == 0 {
		return 0
	}

	min := castToFloat(args[0])
	for i := 1; i < len(args); i++ {
		val := castToFloat(args[i])
		if val < min {
			min = val
		}
	}
	return min
}

// mathMod 返回两个整数的模数
// 支持包名math.Mod函数
func mathMod(a, b interface{}) int {
	ai := castToInt(a)
	bi := castToInt(b)
	if bi == 0 {
		return 0
	}
	return ai % bi
}

// mathMul 将两个或多个数字相乘
// 支持包名math.Mul函数
func mathMul(args ...interface{}) float64 {
	if len(args) == 0 {
		return 0
	}

	result := castToFloat(args[0])
	for i := 1; i < len(args); i++ {
		result *= castToFloat(args[i])
	}
	return result
}

// mathPi 返回数学常数π
// 支持包名math.Pi函数
func mathPi() float64 {
	return math.Pi
}

// mathPow 返回第一个数字的第二个数字次方
// 支持包名math.Pow函数
func mathPow(x, y interface{}) float64 {
	return math.Pow(castToFloat(x), castToFloat(y))
}

// mathRand 返回半开区间[0.0, 1.0)中的伪随机数
// 支持包名math.Rand函数
func mathRand() float64 {
	return rand.Float64()
}

// mathRound 返回最接近的整数，零值向远离零的方向舍入
// 支持包名math.Round函数
func mathRound(n interface{}) float64 {
	return math.Round(castToFloat(n))
}

// mathSqrt 返回给定数字的平方根
// 支持包名math.Sqrt函数
func mathSqrt(n interface{}) float64 {
	return math.Sqrt(castToFloat(n))
}

// mathSub 从第一个数字中减去一个或多个数字
// 支持包名math.Sub函数
func mathSub(args ...interface{}) float64 {
	if len(args) == 0 {
		return 0
	}

	result := castToFloat(args[0])
	for i := 1; i < len(args); i++ {
		result -= castToFloat(args[i])
	}
	return result
}

// ====================
// os 命名空间函数实现
// ====================

// osFileExists 报告文件或目录是否存在
// 支持包名os.FileExists函数
func osFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// osGetenv 返回环境变量的值，如果未设置则返回空字符串
// 支持包名os.Getenv函数
func osGetenv(key string) string {
	return os.Getenv(key)
}

// ====================
// path 命名空间函数实现
// ====================

// pathBase 将路径分隔符替换为斜杠并返回给定路径的最后一个元素
// 支持包名path.Base函数
func pathBase(p string) string {
	return path.Base(filepath.ToSlash(p))
}

// pathBaseName 将路径分隔符替换为斜杠并返回给定路径的最后一个元素，如果存在则移除扩展名
// 支持包名path.BaseName函数
func pathBaseName(p string) string {
	base := path.Base(filepath.ToSlash(p))
	ext := path.Ext(base)
	return strings.TrimSuffix(base, ext)
}

// pathClean 将路径分隔符替换为斜杠并返回与给定路径等效的最短路径名
// 支持包名path.Clean函数
func pathClean(p string) string {
	return path.Clean(filepath.ToSlash(p))
}

// pathDir 将路径分隔符替换为斜杠并返回给定路径除最后一个元素外的所有元素
// 支持包名path.Dir函数
func pathDir(p string) string {
	return path.Dir(filepath.ToSlash(p))
}

// pathExt 将路径分隔符替换为斜杠并返回给定路径的文件扩展名
// 支持包名path.Ext函数
func pathExt(p string) string {
	return path.Ext(filepath.ToSlash(p))
}

// pathJoin 将路径分隔符替换为斜杠，将给定路径元素连接成单个路径
// 支持包名path.Join函数
func pathJoin(elements ...string) string {
	return path.Join(elements...)
}

// pathSplit 将路径分隔符替换为斜杠并在最后一个斜杠之后立即分割
// 支持包名path.Split函数
func pathSplit(p string) (dir, file string) {
	return path.Split(filepath.ToSlash(p))
}

// ====================
// reflect 命名空间函数实现
// ====================

// reflectIsMap 报告给定值是否为映射
// 支持包名reflect.IsMap函数
func reflectIsMap(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.ValueOf(v).Kind() == reflect.Map
}

// reflectIsSlice 报告给定值是否为切片
// 支持包名reflect.IsSlice函数
func reflectIsSlice(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.ValueOf(v).Kind() == reflect.Slice
}

// ====================
// safe 命名空间函数实现
// ====================

// safeCSS 将给定字符串声明为已知的"安全"CSS字符串
// 支持包名safe.CSS函数
func safeCSS(s interface{}) template.CSS {
	return template.CSS(castToString(s))
}

// safeHTML 将给定字符串声明为"安全"HTML文档
// 支持包名safe.HTML函数
func safeHTML(s interface{}) template.HTML {
	return template.HTML(castToString(s))
}

// safeHTMLAttr 将给定字符串声明为安全HTML属性
// 支持包名safe.HTMLAttr函数
func safeHTMLAttr(s interface{}) template.HTMLAttr {
	return template.HTMLAttr(castToString(s))
}

// safeJS 将给定字符串声明为已知安全的JavaScript字符串
// 支持包名safe.JS函数
func safeJS(s interface{}) template.JS {
	return template.JS(castToString(s))
}

// safeJSStr 将给定字符串声明为安全的JavaScript字符串
// 支持包名safe.JSStr函数
func safeJSStr(s interface{}) template.JSStr {
	return template.JSStr(castToString(s))
}

// safeURL 将给定字符串声明为安全URL或URL子字符串
// 支持包名safe.URL函数
func safeURL(s interface{}) template.URL {
	return template.URL(castToString(s))
}

// ====================
// strings 命名空间函数实现
// ====================

// stringsChomp 返回给定字符串，移除所有尾随换行符和回车符
// 支持包名strings.Chomp函数
func stringsChomp(s string) string {
	return strings.TrimRight(s, "\r\n")
}

// stringsContains 报告给定字符串是否包含给定子字符串
// 支持包名strings.Contains函数
func stringsContains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// stringsContainsAny 报告给定字符串是否包含给定集合中的任何字符
// 支持包名strings.ContainsAny函数
func stringsContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// stringsContainsNonSpace 报告给定字符串是否包含Unicode定义的任何非空格字符
// 支持包名strings.ContainsNonSpace函数
func stringsContainsNonSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

// stringsCount 返回给定字符串中给定子字符串的非重叠实例数
// 支持包名strings.Count函数
func stringsCount(s, substr string) int {
	return strings.Count(s, substr)
}

// stringsCountRunes 返回给定字符串中的符文数（不包括空白字符）
// 支持包名strings.CountRunes函数
func stringsCountRunes(s string) int {
	count := 0
	for _, r := range s {
		if !unicode.IsSpace(r) {
			count++
		}
	}
	return count
}

// stringsCountWords 返回给定字符串中的单词数
// 支持包名strings.CountWords函数
func stringsCountWords(s string) int {
	return len(strings.Fields(s))
}

// stringsFindRE 返回匹配正则表达式的字符串切片
// 支持包名strings.FindRE函数
func stringsFindRE(pattern, s string, limit ...int) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}

	n := -1
	if len(limit) > 0 {
		n = limit[0]
	}

	return re.FindAllString(s, n)
}

// stringsFirstUpper 返回给定字符串，将第一个字符大写
// 支持包名strings.FirstUpper函数
func stringsFirstUpper(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

// stringsHasPrefix 报告给定字符串是否以给定前缀开始
// 支持包名strings.HasPrefix函数
func stringsHasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// stringsHasSuffix 报告给定字符串是否以给定后缀结束
// 支持包名strings.HasSuffix函数
func stringsHasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// stringsRepeat 返回由另一个字符串的零个或多个副本组成的新字符串
// 支持包名strings.Repeat函数
func stringsRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// stringsReplace 返回INPUT的副本，将所有OLD替换为NEW
// 支持包名strings.Replace函数
func stringsReplace(s, old, new string, n ...int) string {
	limit := -1
	if len(n) > 0 {
		limit = n[0]
	}
	return strings.Replace(s, old, new, limit)
}

// stringsReplaceRE 返回INPUT的副本，使用替换模式替换正则表达式的所有匹配
// 支持包名strings.ReplaceRE函数
func stringsReplaceRE(pattern, repl, s string) string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return s
	}
	return re.ReplaceAllString(s, repl)
}

// stringsRuneCount 返回给定字符串中的符文数
// 支持包名strings.RuneCount函数
func stringsRuneCount(s string) int {
	return utf8.RuneCountInString(s)
}

// stringsSliceString 返回给定字符串的子字符串，从开始位置开始，在结束位置之前结束
// 支持包名strings.SliceString函数
func stringsSliceString(s string, start, end int) string {
	runes := []rune(s)
	length := len(runes)

	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start > end {
		return ""
	}

	return string(runes[start:end])
}

// stringsSplit 通过分隔符将给定字符串分割为字符串切片
// 支持包名strings.Split函数
func stringsSplit(s, sep string) []string {
	return strings.Split(s, sep)
}

// stringsSubstr 返回给定字符串的子字符串，从开始位置开始，长度为给定长度
// 支持包名strings.Substr函数
func stringsSubstr(s string, start int, length ...int) string {
	runes := []rune(s)
	runeLength := len(runes)

	if start < 0 {
		start = runeLength + start
	}
	if start < 0 {
		start = 0
	}
	if start >= runeLength {
		return ""
	}

	end := runeLength
	if len(length) > 0 && length[0] >= 0 {
		end = start + length[0]
		if end > runeLength {
			end = runeLength
		}
	}

	return string(runes[start:end])
}

// stringsTitle 返回给定字符串，将其转换为标题大小写
// 支持包名strings.Title函数
func stringsTitle(s string) string {
	return strings.Title(s)
}

// stringsToLower 返回给定字符串，将所有字符转换为小写
// 支持包名strings.ToLower函数
func stringsToLower(s string) string {
	return strings.ToLower(s)
}

// stringsToUpper 返回给定字符串，将所有字符转换为大写
// 支持包名strings.ToUpper函数
func stringsToUpper(s string) string {
	return strings.ToUpper(s)
}

// stringsTrim 返回给定字符串，移除前导和尾随字符集中指定的字符
// 支持包名strings.Trim函数
func stringsTrim(s, cutset string) string {
	return strings.Trim(s, cutset)
}

// stringsTrimLeft 返回给定字符串，移除字符集中指定的前导字符
// 支持包名strings.TrimLeft函数
func stringsTrimLeft(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// stringsTrimPrefix 返回给定字符串，从字符串开头移除前缀
// 支持包名strings.TrimPrefix函数
func stringsTrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// stringsTrimRight 返回给定字符串，移除字符集中指定的尾随字符
// 支持包名strings.TrimRight函数
func stringsTrimRight(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// stringsTrimSpace 返回给定字符串，移除Unicode定义的前导和尾随空白字符
// 支持包名strings.TrimSpace函数
func stringsTrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// stringsTrimSuffix 返回给定字符串，从字符串末尾移除后缀
// 支持包名strings.TrimSuffix函数
func stringsTrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// stringsTruncate 返回给定字符串，截断到最大长度而不切断单词或留下未关闭的HTML标签
// 支持包名strings.Truncate函数
func stringsTruncate(s string, max int, suffix ...string) string {
	if len(s) <= max {
		return s
	}

	suf := "…"
	if len(suffix) > 0 {
		suf = suffix[0]
	}

	// 找到最后一个空格
	truncated := s[:max-len(suf)]
	if idx := strings.LastIndex(truncated, " "); idx > 0 {
		truncated = truncated[:idx]
	}

	return truncated + suf
}

// ====================
// time 命名空间函数实现
// ====================

// timeAsTime 将给定字符串表示的日期/时间值作为time.Time值返回
// 支持包名time.AsTime函数
func timeAsTime(v interface{}) time.Time {
	switch val := v.(type) {
	case time.Time:
		return val
	case string:
		// 尝试解析常见的时间格式
		formats := []string{
			time.RFC3339,
			time.RFC3339Nano,
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
			"2006-01-02",
			"01/02/2006",
		}
		for _, format := range formats {
			if t, err := time.Parse(format, val); err == nil {
				return t
			}
		}
	}
	return time.Time{}
}

// timeFormat 将给定日期/时间格式化并本地化为字符串
// 支持包名time.Format函数
func timeFormat(format string, t interface{}) string {
	tm := timeAsTime(t)
	if tm.IsZero() {
		return ""
	}
	return tm.Format(format)
}

// timeNow 返回当前本地时间
// 支持包名time.Now函数
func timeNow() time.Time {
	return time.Now()
}

// timeParseDuration 通过解析给定持续时间字符串返回time.Duration值
// 支持包名time.ParseDuration函数
func timeParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

// ====================
// transform 命名空间函数实现
// ====================

// transformHTMLEscape 返回给定字符串，通过用HTML实体替换特殊字符来转义
// 支持包名transform.HTMLEscape函数
func transformHTMLEscape(s string) string {
	return html.EscapeString(s)
}

// transformHTMLUnescape 返回给定字符串，将每个HTML实体替换为其对应字符
// 支持包名transform.HTMLUnescape函数
func transformHTMLUnescape(s string) string {
	return html.UnescapeString(s)
}

// transformMarkdownify 将Markdown渲染为HTML
// 支持包名transform.Markdownify函数
func transformMarkdownify(s string) template.HTML {
	// 这是一个简化实现，实际应用中应该使用完整的Markdown解析器
	// 这里只做基本的替换
	result := s
	result = strings.ReplaceAll(result, "\n\n", "</p><p>")
	result = "<p>" + result + "</p>"
	return template.HTML(result)
}

// transformPlainify 返回删除所有HTML标记的字符串的纯文本版本
// 支持包名transform.Plainify函数
func transformPlainify(s string) string {
	// 简单的HTML标签移除
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

// ====================
// urls 命名空间函数实现
// ====================

// urlsAbsURL 返回绝对URL
// 支持包名urls.AbsURL函数
func urlsAbsURL(s string) string {
	// 这是一个简化实现，实际应用中需要配置baseURL
	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
		return s
	}
	return "http://localhost" + s
}

// urlsAnchorize 返回给定字符串，清理后用于HTML id属性
// 支持包名urls.Anchorize函数
func urlsAnchorize(s string) string {
	// 转换为小写，替换空格和特殊字符为连字符
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// urlsJoinPath 将提供的元素连接成URL字符串并清理结果
// 支持包名urls.JoinPath函数
func urlsJoinPath(elements ...string) string {
	return path.Join(elements...)
}

// urlsParse 将URL解析为URL结构
// 支持包名urls.Parse函数
func urlsParse(s string) (*url.URL, error) {
	return url.Parse(s)
}

// urlsRelURL 返回相对URL
// 支持包名urls.RelURL函数
func urlsRelURL(s string) string {
	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
		u, err := url.Parse(s)
		if err != nil {
			return s
		}
		return u.Path
	}
	if !strings.HasPrefix(s, "/") {
		return "/" + s
	}
	return s
}

// urlsURLize 返回给定字符串，清理后用于URL
// 支持包名urls.URLize函数
func urlsURLize(s string) string {
	// 转换为小写，替换空格和特殊字符为连字符
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
