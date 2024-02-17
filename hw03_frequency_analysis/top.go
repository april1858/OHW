package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Frequency struct {
	w string
	n int
}

func Top10(str string) []string {
	if len(str) == 0 {
		return nil
	}
	counts := countsWords(str)
	countsSort := countsToSort(counts)
	allFr := frAnalysis10(countsSort)
	answer := make([]string, 0, 10)
	for _, frequency := range allFr {
		answer = append(answer, frequency.w)
	}
	return answer
}

// countsWords() возвращает количество вхождлений слова в тексте

func countsWords(str string) map[string]int {
	counts := make(map[string]int)
	for _, v := range strings.Fields(str) {
		counts[v]++
	}
	return counts
}

// countsToSort() возвращает срез структур Frequecy отсортированных по количеству вхождений

func countsToSort(counts map[string]int) []Frequency {
	countsSort := make([]Frequency, 0, len(counts))
	for w, n := range counts {
		countsSort = append(countsSort, Frequency{w, n})
	}
	sort.Slice(countsSort, func(i, j int) bool {
		return countsSort[i].n > countsSort[j].n
	})
	return countsSort
}

// frAnalysis10() возвращает достаточное количество сегментов отсортированных по частоте и лексикографически
// для получения среза с 10-ю наиболее часто встречаемыми в тексте словами

func frAnalysis10(cs []Frequency) []Frequency {
	stop := len(cs)
	buckets := make([]Frequency, 0, 1)
	for {
		bucket := getBucket(cs[len(buckets):])
		buckets = append(buckets, bucket...)
		if len(bucket) >= 10 || len(buckets) >= stop {
			break
		}
	}
	if len(buckets) >= 10 {
		return buckets[:10]
	}
	return buckets
}

// getBucket() возвращает отсортированный сегмент с определенной частотой

func getBucket(cs []Frequency) []Frequency {
	m := cs[0].n // частота сегмента
	var bucket []Frequency
	for _, v := range cs {
		if v.n == m {
			bucket = append(bucket, v)
		}
	}
	bucketSort := bucketSort(bucket)
	return bucketSort
}

// bucketSort() возвращает лексикографически отсортированный сегмент

func bucketSort(bucket []Frequency) []Frequency {
	sort.Slice(bucket, func(i, j int) bool {
		return bucket[i].w < bucket[j].w
	})
	return bucket
}
