package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	var host = "http://192.168.56.101:9200"

	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")

	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("ES ping code %d, version %s\n", code, info.Version.Number)

	// 添加
	//p1 := Person{Name: "wangzi", Age: 42, Married: false}
	//put1, err := client.Index().
	//	Index("user").
	//	Id("666").
	//	BodyJson(p1).
	//	Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Get tweet with specified ID
	//get1, err := client.Get().
	//	Index("user").
	//	Id("666").
	//	Do(context.Background())
	//if err != nil {
	//	switch {
	//	case elastic.IsNotFound(err):
	//		panic(fmt.Sprintf("Document not found: %v", err))
	//	case elastic.IsTimeout(err):
	//		panic(fmt.Sprintf("Timeout retrieving document: %v", err))
	//	case elastic.IsConnErr(err):
	//		panic(fmt.Sprintf("Connection problem: %v", err))
	//	default:
	//		// Some other kind of error
	//		panic(err)
	//	}
	//}
	//fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	//fmt.Printf("result: %s", get1.Source)

	//var res *elastic.SearchResult
	//取所有
	//res, err = client.Search("user").Do(context.Background())
	//printUser(res, err)

	// Search with a term query
	//termQuery := elastic.NewQueryStringQuery("name:wangzi")
	//searchResult, err := client.Search().
	//	Index("user"). // search in index "twitter"
	//	Query(termQuery). // specify the query
	//	//Sort("name", true).      // sort by "user" field, ascending
	//	From(0).Size(10). // take documents 0-9
	//	Pretty(true). // pretty print request and response JSON
	//	Do(context.Background()) // execute
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//printUser(searchResult, err)

	//boolQ := elastic.NewBoolQuery()
	//boolQ.Must(elastic.NewMatchQuery("name", "zhaonan"))
	//boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	//res, err = client.Search("user").Query(boolQ).Do(context.Background())
	//printUser(res, err)
}

func printUser(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Person
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Person)
		fmt.Printf("%#v\n", t)
	}
}
