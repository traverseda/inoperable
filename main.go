//A prototype terminal emulator
package main

import (
	"fmt"
	"github.com/oleiade/lane"
	"io/ioutil"
	//    "bufio"
	//    "io"
	//    "strings"
	//    "bytes"
	//    "reflect"
	//    "os"
	//    "container/list"
)

func print(i chan rune, o chan rune) {
	//fmt.Println(i)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Router struct {
	filters    map[string]func(i chan rune, o chan rune)
	deque      lane.Deque
    warn       []rune
	currentStr []rune
}

func (r *Router) add_function(key []rune, f func(chan rune, chan rune)) {
	r.filters[string(key)] = f
}

func (r *Router) init() {
	r.deque = *lane.NewDeque()
    r.filters = map[string]func(i chan rune, o chan rune){}
}

func (r *Router) route() {
	//	filters := r.filters
	tag := []rune{}
    tags := []rune{}
	out := make(chan rune)//This is a stream of characters to send to filter
	for i := range r.Iter() {
		if i == rune(40) {
            filter := r.filters[string(tag)]
            fmt.Println(string(tag))
            tag = []rune{}
            if filter != nil {
                filter(r.Iter(), out)
            }
        } else {
            tag = append(tag, i)
     //       fmt.Println(string(i))
        }
	}
    fmt.Println(string(tags))
}

func (r *Router) Append(l []rune) {
	r.deque.Prepend(r.currentStr)
	r.currentStr = []rune{}
	r.deque.Prepend(l)
}

func (r *Router) Iter() chan rune {
	ch := make(chan rune)
	go func() {
		for {
			if r.deque.Empty() {
				close(ch)
				break
			}
			r.currentStr = r.deque.Pop().([]rune)
			for _, x := range r.currentStr {
				ch <- rune(x)
			}
		}
	}()
	return ch
}

func strToRune(input string) []rune {
	// Get Unicode code points.
	n := 0
	runeList := make([]rune, len(input))
	for _, r := range input {
		runeList[n] = r
		n++
	}
	runeList = runeList[0:n]
	return runeList
}

func main() {
	r := Router{}
	r.init()
	r.add_function(strToRune("print"), print)
	dat, err := ioutil.ReadFile("./test.tmp")
	check(err)
	r.deque.Prepend(strToRune(string(dat)))
	r.deque.Prepend(strToRune("print(hello world)"))
	r.route()
}
