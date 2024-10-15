package main

import (
    "bufio"
    "container/list"
    "fmt"
    "os"
    "strings"
)

// Defines a SlidingQueue struct with three fields: maxSize (maximum size of the queue), queue (a list to store words), and counts (a map to keep track of word frequencies).
type SlidingQueue struct {
    maxSize int
    queue   *list.List
    counts  map[string]int
}

//constructor
func NewSlidingQueue(size int) *SlidingQueue {
    return &SlidingQueue{
        maxSize: size,
        queue:   list.New(),
        counts:  make(map[string]int),
    }
}
//add/remove/update word count
func (sq *SlidingQueue) Add(word string) {
    if sq.queue.Len() == sq.maxSize {
        // Remove the oldest word
        oldest := sq.queue.Remove(sq.queue.Front()).(string)
        sq.counts[oldest]--
        if sq.counts[oldest] == 0 {
            delete(sq.counts, oldest)
        }
    }
    // Add the new word
    sq.queue.PushBack(word)
    sq.counts[word]++
}

//returns current word counts as a map
func (sq *SlidingQueue) WordCount() map[string]int {
    return sq.counts
}


func main() {
    sq := NewSlidingQueue(5)
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter words (press Ctrl+C to stop):")

    for {
        fmt.Print("> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        words := strings.Fields(input)

        for _, word := range words {
            sq.Add(word)
        }

        fmt.Println("Current queue:")
        for e := sq.queue.Front(); e != nil; e = e.Next() {
            fmt.Print(e.Value, " ")
        }
        fmt.Println()
        fmt.Println("Word counts:")
        for k, v := range sq.WordCount() {
            fmt.Printf("%s: %d\n", k, v)
        }
        fmt.Println("----------")
    }
}
