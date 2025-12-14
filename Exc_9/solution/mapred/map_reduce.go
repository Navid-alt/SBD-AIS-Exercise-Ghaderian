package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
func (mr *MapReduce) Run(input []string) map[string]int {
	mapResults := make(chan []KeyValue)
	var mapWg sync.WaitGroup

	for _, chunk := range input {
		mapWg.Add(1)
		go func(text string) {
			defer mapWg.Done()
			results := mr.wordCountMapper(text)
			mapResults <- results
		}(chunk)
	}

	go func() {
		mapWg.Wait()
		close(mapResults)
	}()

	intermediate := make(map[string][]int)
	for kvs := range mapResults {
		for _, kv := range kvs {
			intermediate[kv.Key] = append(intermediate[kv.Key], kv.Value)
		}
	}

	finalResult := make(map[string]int)
	var reduceWg sync.WaitGroup
	var mu sync.Mutex

	for key, values := range intermediate {
		reduceWg.Add(1)
		go func(k string, v []int) {
			defer reduceWg.Done()
			reducedKV := mr.wordCountReducer(k, v)
			mu.Lock()
			finalResult[reducedKV.Key] = reducedKV.Value
			mu.Unlock()
		}(key, values)
	}
	reduceWg.Wait()

	return finalResult
}

func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	reg := regexp.MustCompile(`[^a-zA-Z]+`)

	words := reg.Split(text, -1)
	var results []KeyValue

	for _, word := range words {
		if word != "" {
			results = append(results, KeyValue{
				Key:   strings.ToLower(word),
				Value: 1,
			})
		}
	}
	return results
}

func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	count := 0
	for _, v := range values {
		count += v
	}
	return KeyValue{
		Key:   key,
		Value: count,
	}
}
