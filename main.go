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

type RouterObject struct {
	key   []rune
	value func(i chan rune, o chan rune)
}

type Router struct {
	filters    []RouterObject
	deque      lane.Deque
	currentStr []rune
}

func (r *Router) add_function(key []rune, f func(chan rune, chan rune)) {
	r.filters = append(r.filters, RouterObject{key, f})
}

func (r *Router) init() {
	r.deque = *lane.NewDeque()
}

func (r *Router) route() {
	//	filters := r.filters
	counter := 0
	tag := ""
	depth := 0
//	    out := make(chan rune)//This is a stream of characters to send to filter
//	    in := make(chan rune)//This is that return data, after it's been filtered.
	for i := range r.Iter() {
		counter++
		if i == rune(40) {
			depth++
		} else if i == rune(41) {
			depth--
		} else if depth == 0{
		    tag = tag + string(i)
        } else {
            fmt.Println(tag)
            tag = ""
        }
	}
}

func (r *Router) Append(l []rune) {
	r.deque.Prepend(r.currentStr)
	r.currentStr = []rune{}
	r.deque.Prepend(l)
}

func (r *Router) Iter() <-chan rune {
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
