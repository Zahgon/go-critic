package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary Detects redundant fmt.Sprint calls
//doc:tags    style experimental
//doc:before  fmt.Sprint(x)
//doc:after   x.String()
func redundantSprint(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects deferred function literals that can be simplified
//doc:tags    style experimental
//doc:before  defer func() { f() }()
//doc:after   defer f()
func deferUnlambda(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious mutex lock/unlock operations
//doc:tags    diagnostic experimental
//doc:before  mu.Lock(); mu.Unlock()
//doc:after   mu.Lock(); defer mu.Unlock()
func badLock(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// `mu1` and `mu2` are added to make possible report a line where `m2` is used (with a defer)
	return
}

// no defer

// different lock operations

// double locks

//doc:summary Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative
//doc:tags    style experimental
//doc:before  http.NewRequest("GET", url, nil)
//doc:after   http.NewRequest("GET", url, http.NoBody)
func httpNoBody(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation
//doc:tags    performance experimental
//doc:before  r := []rune(s)[0]
//doc:after   r, _ := utf8.DecodeRuneInString(s)
//doc:note    See Go issue for details: https://github.com/golang/go/issues/45260
func preferDecodeRune(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects usage of `len` when result is obvious or doesn't make sense
//doc:tags    diagnostic
//doc:before  len(arr) <= 0
//doc:after   len(arr) == 0
func sloppyLen(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects value swapping code that are not using parallel assignment
//doc:tags    style
//doc:before  *tmp = *x; *x = *y; *y = *tmp
//doc:after   *x, *y = *y, *x
func valSwap(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects switch-over-bool statements that use explicit `true` tag value
//doc:tags    style
//doc:before  switch true {...}
//doc:after   switch {...}
func switchTrue(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects immediate dereferencing of `flag` package pointers
//doc:tags    diagnostic
//doc:before  b := *flag.Bool("b", false, "b docs")
//doc:after   var b bool; flag.BoolVar(&b, "b", false, "b docs")
func flagDeref(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects empty string checks that can be written more idiomatically
//doc:tags    style experimental
//doc:before  len(s) == 0
//doc:after   s == ""
func emptyStringTest(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects redundant conversions between string and []byte
//doc:tags    performance
//doc:before  copy(b, []byte(s))
//doc:after   copy(b, s)
func stringXbytes(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects strings.Index calls that may cause unwanted allocs
//doc:tags    performance
//doc:before  strings.Index(string(x), y)
//doc:after   bytes.Index(x, []byte(y))
//doc:note    See Go issue for details: https://github.com/golang/go/issues/25864
func indexAlloc(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects function calls that can be replaced with convenience wrappers
//doc:tags    style
//doc:before  wg.Add(-1)
//doc:after   wg.Done()
func wrapperFunc(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`
//doc:tags    style
//doc:before  re, _ := regexp.Compile("const pattern")
//doc:after   re := regexp.MustCompile("const pattern")
func regexpMust(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious function calls
//doc:tags    diagnostic
//doc:before  strings.Replace(s, from, to, 0)
//doc:after   strings.Replace(s, from, to, -1)
func badCall(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects assignments that can be simplified by using assignment operators
//doc:tags    style
//doc:before  x = x * 2
//doc:after   x *= 2
func assignOp(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects WriteRune calls with rune literal argument that is single byte and reports to use WriteByte instead
//doc:tags    performance experimental opinionated
//doc:before  w.WriteRune('\n')
//doc:after   w.WriteByte('\n')
func preferWriteByte(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// utf8.RuneSelf:
	// characters below RuneSelf are represented as themselves in a single byte.
	return
}

//doc:summary Detects fmt.Sprint(f/ln) calls which can be replaced with fmt.Fprint(f/ln)
//doc:tags    performance experimental
//doc:before  w.Write([]byte(fmt.Sprintf("%x", 10)))
//doc:after   fmt.Fprintf(w, "%x", 10)
func preferFprint(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious duplicated arguments
//doc:tags    diagnostic
//doc:before  copy(dst, dst)
//doc:after   copy(dst, src)
func dupArg(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious http.Error call without following return
//doc:tags    diagnostic experimental
//doc:before  if err != nil { http.Error(...); }
//doc:after   if err != nil { http.Error(...); return; }
func returnAfterHttpError(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects concatenation with os.PathSeparator which can be replaced with filepath.Join
//doc:tags    style experimental
//doc:before  x + string(os.PathSeparator) + y
//doc:after   filepath.Join(x, y)
func preferFilepathJoin(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects w.Write or io.WriteString calls which can be replaced with w.WriteString
//doc:tags    performance experimental
//doc:before  w.Write([]byte("foo"))
//doc:after   w.WriteString("foo")
func preferStringWriter(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects slice clear loops, suggests an idiom that is recognized by the Go compiler
//doc:tags    performance experimental
//doc:before  for i := 0; i < len(buf); i++ { buf[i] = 0 }
//doc:after   for i := range buf { buf[i] = 0 }
func sliceClear(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects sync.Map load+delete operations that can be replaced with LoadAndDelete
//doc:tags    diagnostic experimental
//doc:before  v, ok := m.Load(k); if ok { m.Delete($k); f(v); }
//doc:after   v, deleted := m.LoadAndDelete(k); if deleted { f(v) }
func syncMapLoadAndDelete(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects "%s" formatting directives that can be replaced with %q
//doc:tags    diagnostic experimental
//doc:before  fmt.Sprintf(`"%s"`, s)
//doc:after   fmt.Sprintf(`%q`, s)
func sprintfQuotedString(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects various off-by-one kind of errors
//doc:tags    diagnostic
//doc:before  xs[len(xs)]
//doc:after   xs[len(xs)-1]
func offBy1(m dsl.Matcher) { _ = "STUB: not implemented"; return }

// TODO: use $slicing[$i:$*_] form when we'll update go-ruleguard
// version so it includes https://github.com/quasilyte/go-ruleguard/pull/284

//doc:summary Detects slice expressions that can be simplified to sliced expression itself
//doc:tags    style
//doc:before  copy(b[:], values...)
//doc:after   copy(b, values...)
func unslice(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects Yoda style expressions and suggests to replace them
//doc:tags    style experimental
//doc:before  return nil != ptr
//doc:after   return ptr != nil
func yodaStyleExpr(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects unoptimal strings/bytes case-insensitive comparison
//doc:tags    performance experimental
//doc:before  strings.ToLower(x) == strings.ToLower(y)
//doc:after   strings.EqualFold(x, y)
func equalFold(m dsl.Matcher) {
	_ = "STUB: not implemented"
	// We specify so many patterns to avoid too generic
	// patterns that would match things like
	// `strings.ToLower(x) == strings.ToUpper(y)`
	// While it could be an EqualFold candidate,
	// it just looks wrong and should probably be
	// marked by some other checker.
	return
}

// string == patterns

// string != patterns

// bytes.Equal patterns

//doc:summary Detects suspicious arguments order
//doc:tags    diagnostic
//doc:before  strings.HasPrefix("#", userpass)
//doc:after   strings.HasPrefix(userpass, "#")
func argOrder(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects string concat operations that can be simplified
//doc:tags    style experimental
//doc:before  strings.Join([]string{x, y}, "_")
//doc:after   x + "_" + y
func stringConcatSimplify(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects manual conversion to milli- or microseconds
//doc:tags    style experimental
//doc:before  t.Unix() / 1000
//doc:after   t.UnixMilli()
func timeExprSimplify(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects exposed methods from sync.Mutex and sync.RWMutex
//doc:tags    style experimental
//doc:before  type Foo struct{ ...; sync.Mutex; ... }
//doc:after   type Foo struct{ ...; mu sync.Mutex; ... }
func exposedSyncMutex(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects bad usage of sort package
//doc:tags    diagnostic experimental
//doc:before  xs = sort.StringSlice(xs)
//doc:after   sort.Strings(xs)
func badSorting(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious reassignment of error from another package
//doc:tags    diagnostic experimental
//doc:before  io.EOF = nil
//doc:after   /* don't do it */
func externalErrorReassign(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious empty declarations blocks
//doc:tags    diagnostic experimental
//doc:before  var()
//doc:after   /* nothing */
func emptyDecl(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects suspicious formatting strings usage
//doc:tags    diagnostic experimental
//doc:before  fmt.Errorf(msg)
//doc:after   errors.New(msg) or fmt.Errorf("%s", msg)
func dynamicFmtString(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects strings.Compare usage
//doc:tags    style experimental
//doc:before  strings.Compare(x, y)
//doc:after   x < y
func stringsCompare(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects unchecked errors in if statements
//doc:tags    diagnostic experimental
//doc:before  if err := expr(); err2 != nil { /*...*/ }
//doc:after   if err := expr(); err != nil { /*...*/ }
func uncheckedInlineErr(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects bad usage of sync.OnceFunc
//doc:tags    diagnostic experimental
//doc:before  sync.OnceFunc(foo)()
//doc:after   fooOnce := sync.OnceFunc(foo); ...; fooOnce()
func badSyncOnceFunc(m dsl.Matcher) { _ = "STUB: not implemented"; return }

//doc:summary Detects bytes.Repeat with 0 value
//doc:tags    performance
//doc:before  bytes.Repeat([]byte{0}, x)
//doc:after   make([]byte, x)
func zeroByteRepeat(m dsl.Matcher) { _ = "STUB: not implemented"; return }

// Rule 2: const identifier with value 0
