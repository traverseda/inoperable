//A prototype terminal emulator
package main

import (
	"github.com/oleiade/lane"
	"fmt"
	"io/ioutil"
    //"gopkg.in/yaml.v2"
	//    "bufio"
	//    "io"
	//    "strings"
	//    "bytes"
	//    "reflect"
	//    "os"
	//    "container/list"
)

func passthrough(i chan rune, o chan rune) {
    depth:=1
    for foo := range i {
        //This chunk of code counts brackets, and stops the filter when it's done.
        //You could keep reading after you should be closed, but that's going to break everything.
        fmt.Println(string(foo))
        if foo == rune(40) {
            depth++
        } else if foo == rune(41){
            depth--
        }
        if depth == 0{
            close(o)
            break
        }
    }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Router struct {
	filters    map[string]func(i chan rune, o chan rune)
    unkownTag  func(i chan rune, o chan rune)
	deque      lane.Deque
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
	tag := []rune{}
    tags := []rune{}
	out := make(chan rune)//This is a stream of characters the filter returns
    bar := r.Iter()
	for i := range bar {
     //   fmt.Println(o)
		if i == rune(40) {
            filter := r.filters[string(tag)]
            tag = []rune{}
            if filter != nil {
                fmt.Println(filter)
                filter(r.Iter(), out)
            }
        } else {
            tag = append(tag, i)
            fmt.Println(string(i))
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
	r.add_function(strToRune("print"), passthrough)
	dat, err := ioutil.ReadFile("./test.tmp")
	check(err)
	r.deque.Prepend(strToRune(string(dat)))
	r.route()
}
