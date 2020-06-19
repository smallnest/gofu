// Code generated by go2go; DO NOT EDIT.


//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer_test.go2:1
package gofu

//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer_test.go2:1
import "testing"

//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer_test.go2:5
type testConsumer struct {
	t     *testing.T
	after instantiate୦୦Consumer୦string
}

//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer_test.go2:9
func (c testConsumer) Accept(v string) {
	c.t.Log(v)
	if c.after != nil {
		c.after.Accept(v)
	}
}

func (c testConsumer) AndThen(after instantiate୦୦Consumer୦string,) instantiate୦୦Consumer୦string {
	c.after = after
	return c
}

func TestConsumer(t *testing.T) {
	tc := testConsumer{t: t}
	tc.Accept("hello")
	tc.AndThen(testConsumer{t: t}).Accept("world")
}

type testBiConsumer struct {
	t *testing.T
}

//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer_test.go2:30
func (c testBiConsumer) Accept(v string, i int) {
	c.t.Logf("%s,%d", v, i)
}

func (c testBiConsumer) AndThen(after instantiate୦୦BiConsumer୦string୦int,) instantiate୦୦BiConsumer୦string୦int {
	return after
}

func TestBiConsumer(t *testing.T) {
	tc := testBiConsumer{t}
	tc.Accept("hello", 100)
	tc.AndThen(tc).Accept("world", 200)
}

//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer_test.go2:42
type instantiate୦୦Consumer୦string interface {
//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer.go2:7
 Accept(string)

									AndThen(after instantiate୦୦Consumer୦string,) instantiate୦୦Consumer୦string
}
//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer.go2:10
type instantiate୦୦BiConsumer୦string୦int interface {
//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer.go2:17
 Accept(t string, u int)

	AndThen(after instantiate୦୦BiConsumer୦string୦int,) instantiate୦୦BiConsumer୦string୦int
}

//line /Users/yuepan/go/src/github.com/smallnest/gofu/consumer.go2:20
var _ = testing.AllocsPerRun