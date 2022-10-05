package main

import (
    "sort"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
)
//creating page structure
type Page struct {
    Url string `json:"url"`
    Views int `json:"views"`
    RelevanceScore float32 `json:"relevanceScore"`
}
//response to reference page
type Response struct {
    Data []Page `json:"data"`
}
//resulto reference page and count
type Result struct {
    Data []Page `json:"data"`
    Count int `json:"count"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	//create list to have all 3 URLs
    sources := [3]string{
        "https://raw.githubusercontent.com/assignment132/assignment/main/duckduckgo.json",
        "https://raw.githubusercontent.com/assignment132/assignment/main/google.json",
        "https://raw.githubusercontent.com/assignment132/assignment/main/wikipedia.json",
    }
    var pages = []Page{} // create empty struct pages of struct type Page
    for i:=0; i < len(sources); i++ { // iterate through the URL list
    //    fmt.Printf("\n"+sources[i]+"\n")
        http_resp, err_http_get := http.Get(sources[i]) //get http response from each URL
        if err_http_get != nil {
            fmt.Println("Err_http_getor in http.Get")
            log.Fatalln(err_http_get)
        }
        var response Response
        responseBytes, err_http_read := ioutil.ReadAll(http_resp.Body)// Read the http response and store in response byte variable
        if err_http_read != nil {
	    fmt.Println("Error in ioutil.ReadAll")
            log.Fatalln(err_http_read)
        }
        err_json_unmarshal := json.Unmarshal(responseBytes, &response) // make responseBytes a go object 
        if err_json_unmarshal != nil {
            fmt.Println("Error in json.Unmarshal")
            log.Fatalln(err_json_unmarshal)
        }
        for j:=0; j < len(response.Data); j++ {
            pages = append(pages, response.Data[j]) // add to the pages struct each page object           
        }
    }

  //  fmt.Printf("%+v\n", pages) // pages will have all the pages from all 3 URLs
    var limit int // create limit variable
    var sortKey string 
    var err_atoi error // create atoi error variable
    
    limit, err_atoi = strconv.Atoi(r.URL.Query().Get("limit")) // get the limit value from the URL string and convert string to int
    sortKey = r.URL.Query().Get("sortKey")
    
    if err_atoi != nil {
	fmt.Println("Error in strconv.Atoi")
	log.Fatalln(err_atoi)
    }
    fmt.Printf("sortKey: " + sortKey) // show sortKey value pulled in string
    fmt.Printf("Limit: " + string(limit)) // show limit value pulled in int
    
    if sortKey == "views"{
    	sort.SliceStable(pages, func(i, j int) bool {
        return pages[i].Views < pages[j].Views
    })
    //fmt.Println(pages)
    }
    if sortKey == "relevanceScore"{
	sort.SliceStable(pages, func(i, j int) bool {
        return pages[i].RelevanceScore < pages[j].RelevanceScore
    })
    //fmt.Println(pages)
    }

    result_object := Result{Data: pages[0:int(limit)], Count: limit} // Create the final result struct pulling from pages, adding all pages from 0 to the limit value specified in the URL, add count key and set value to limit

    var result, err_json_marshal = json.MarshalIndent(result_object, "", "  ")  // convert struct object to json string, no space, indentation
    if err_json_marshal != nil {
        fmt.Println("Error in json.Marshal")
        log.Fatalln(err_json_marshal)
    }
    
    if sortKey != "views" && sortKey != "relevanceScore"{
	io.WriteString(w, "Invalid sortKey value, please try using views or relevanceScore")
    }else{     
    io.WriteString(w, string(result)) // print result to web
    }
}


func main() {
    http.HandleFunc("/", getRoot)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
