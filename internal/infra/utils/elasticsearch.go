package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"io"
	"path/filepath"
	"reflect"
	"runtime"
)

func GetQueryMarshalledJson(query map[string]interface{}) []byte {
	funcName := runtime.FuncForPC(reflect.ValueOf(GetQueryMarshalledJson).Pointer()).Name()
	funcName = filepath.Base(funcName)

	queryJson, err := json.Marshal(query)

	HandleError(err, fmt.Sprintf("Error while marshalling query [%s]", funcName))

	return queryJson

}

func ExecuteElasticsearchQuery(queryJson []byte, client *opensearch.Client) *opensearchapi.Response {
	req := opensearchapi.SearchRequest{
		Index: []string{"seclab-events"},
		Body:  bytes.NewReader(queryJson),
	}

	fmt.Println(string(queryJson))

	res, err := req.Do(context.Background(), client)

	HandleError(err, "Error while executing Elasticsearch query")

	if res.IsError() {
		panic(fmt.Sprintf("Elasticsearch search error: %s", res.Status()))
	}

	return res
}

func DecodeElasticsearchResponse(res *opensearchapi.Response, response interface{}) error {

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if closer, ok := res.Body.(io.Closer); ok {
		err := closer.Close()
		if err != nil {
			return err
		}
	}

	return nil

}
