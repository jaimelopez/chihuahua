package storage

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/jaimelopez/chihuahua/executor"
	"gopkg.in/olivere/elastic.v6"
)

const (
	nsPerOpertationIndexName   string = "ns"
	allocationsNumberIndexName string = "mallocbytes"
	allocatedBytesIndexName    string = "mallocs"
	timestampField             string = "@timestamp"
)

// ElasticSearch todo
type ElasticSearch struct {
	URL    string
	Prefix string
}

// NewElasticSearchStorage todo
func NewElasticSearchStorage(url string, prefix string) *ElasticSearch {
	return &ElasticSearch{
		URL:    url,
		Prefix: prefix,
	}
}

// GetLatest todo
func (es *ElasticSearch) GetLatest() (*executor.Result, error) {
	nsPerOperation, err := es.getLatestFromIndex(nsPerOpertationIndexName)
	if err != nil {
		return nil, err
	}

	allocationsNumber, err := es.getLatestFromIndex(allocationsNumberIndexName)
	if err != nil {
		return nil, err
	}

	allocatedBytes, err := es.getLatestFromIndex(allocatedBytesIndexName)
	if err != nil {
		return nil, err
	}

	result := executor.Result{}

	for test, value := range nsPerOperation {
		if _, ok := result[test]; !ok {
			result[test] = &executor.TestResult{Name: test}
		}

		result[test].NsPerOp = value.(float64)
	}

	for test, value := range allocationsNumber {
		if _, ok := result[test]; !ok {
			result[test] = &executor.TestResult{Name: test}
		}

		result[test].AllocsPerOp = uint64(value.(float64))
	}

	for test, value := range allocatedBytes {
		if _, ok := result[test]; !ok {
			result[test] = &executor.TestResult{Name: test}
		}

		result[test].AllocedBytesPerOp = uint64(value.(float64))
	}

	return &result, nil
}

// Persist todo
func (es *ElasticSearch) Persist(r *executor.Result) error {
	nsPerOperation := map[string]interface{}{}
	allocationsNumber := map[string]interface{}{}
	allocatedBytes := map[string]interface{}{}

	for name, tr := range *r {
		nsPerOperation[name] = tr.NsPerOp
		allocationsNumber[name] = tr.AllocsPerOp
		allocatedBytes[name] = tr.AllocedBytesPerOp
	}

	err := es.index(nsPerOpertationIndexName, nsPerOperation)
	if err != nil {
		return err
	}

	err = es.index(allocationsNumberIndexName, allocationsNumber)
	if err != nil {
		return err
	}

	err = es.index(allocatedBytesIndexName, allocatedBytes)
	if err != nil {
		return err
	}

	return nil
}

func (es *ElasticSearch) index(idx string, doc map[string]interface{}) error {
	client, err := es.client()
	if err != nil {
		return err
	}

	idx = es.indexName(idx)
	doc[timestampField] = time.Now()

	_, err = client.Index().
		Index(idx).
		Type("doc").
		BodyJson(doc).
		Do(context.Background())

	return err
}

func (es *ElasticSearch) client() (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetURL(es.URL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
}

func (es *ElasticSearch) indexName(idx string) string {
	return fmt.Sprintf("%s-%s", es.Prefix, idx)
}

func (es *ElasticSearch) getLatestFromIndex(idx string) (map[string]interface{}, error) {
	client, err := es.client()
	if err != nil {
		return nil, err
	}

	sr, err := client.Search(es.indexName(idx)).
		Size(1).
		Sort("@timestamp", false).
		Do(context.Background())

	if err != nil {
		return nil, err
	} else if sr.TotalHits() == 0 {
		return map[string]interface{}{}, nil
	}

	var doc map[string]interface{}
	doc = sr.Each(reflect.TypeOf(doc))[0].(map[string]interface{})

	if _, ok := doc[timestampField]; ok {
		delete(doc, timestampField)
	}

	return doc, nil
}