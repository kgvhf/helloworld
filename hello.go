package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

//@todo вынести таски и файндеры в отдельные файлы
type Task struct {
	url        string
	searchWord string
	countMatch int
	err        error
}

func NewTask(url, searchWord string) *Task {
	t := new(Task)
	t.url = url
	t.searchWord = searchWord
	return t
}

func (t *Task) run() {
	resp, err := http.Get(t.url)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	t.countMatch = t.findMatch(body)
}

func (t Task) findMatch(body []byte) int {
	return bytes.Count(body, []byte(t.searchWord))
}

func (t Task) view() {
	fmt.Println("Count for "+t.url+":", t.countMatch)
}

type Finder struct {
	size       int
	queue      chan string
	queueWG    sync.WaitGroup
	result     chan *Task
	searchWord string
}

func NewFinder(searchWord string, goroutineSize int) *Finder {
	p := new(Finder)
	p.queue = make(chan string, goroutineSize)
	p.result = make(chan *Task)
	p.searchWord = searchWord

	return p
}
func (f *Finder) start(url string) {
	f.queueWG.Add(1)
	go func(url string) {
		defer f.queueWG.Done()
		fmt.Println(url)
		task := NewTask(url, f.searchWord)
		task.run()
		f.result <- task
		time.Sleep(1 * time.Second)
		<-f.queue
	}(url)
}
func (f *Finder) render(done chan<- string) {
	var sum int
	for t := range f.result {
		t.view()
		sum += t.countMatch
	}
	fmt.Println("Total:", sum)
	done <- "I'm done"
}

func main() {
	done := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)
	finder := NewFinder("Go", 2)
	go finder.render(done)

	for scanner.Scan() {
		finder.queue <- "buffer"
		finder.start(scanner.Text())
	}

	finder.queueWG.Wait()
	close(finder.result)
	close(finder.queue)
	<-done
}
