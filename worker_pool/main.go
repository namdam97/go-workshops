package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// worker pools

type Worker struct {
	id         int
	taskQueue  <-chan string
	resultChan chan<- Result
}

type Result struct {
	workerID int
	url      string
	data     string
	err      error
}

func (w *Worker) Start() {
	go func() {
		for url := range w.taskQueue {
			data, err := fetchAndProcess(url)
			w.resultChan <- Result{
				workerID: w.id,
				url:      url,
				data:     data,
				err:      err,
			}
		}

	}()
}

func fetchAndProcess(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch the URL")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	extractedData := string(body)
	return extractedData, nil
}

type WorkerPool struct {
	taskQueue   chan string
	resultChan  chan Result
	workerCount int
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan string),
		resultChan:  make(chan Result),
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		worker := Worker{
			id:         i,
			taskQueue:  wp.taskQueue,
			resultChan: wp.resultChan,
		}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(url string) {
	wp.taskQueue <- url
}

func (wp *WorkerPool) GetResult() Result {
	return <-wp.resultChan
}

func main() {
	urls := []string{
		"https://google.com",
		"https://bing.com",
		"https://apple.com",
	}

	workerPool := NewWorkerPool(3) // Create a worker pool with 3 workers
	workerPool.Start()

	// Submit the URLs to the worker pool for processing
	for _, url := range urls {
		workerPool.Submit(url)
	}

	// Collect the results and handle any errors
	for i := 0; i < len(urls); i++ {
		result := workerPool.GetResult()
		if result.err != nil {
			fmt.Printf("Worker ID: %d, URL: %s, Error: %v\n", result.workerID, result.url, result.err)
		} else {
			fmt.Printf("Worker ID: %d, URL: %s, Data: %s\n", result.workerID, result.url, result.data)
		}
	}
}
