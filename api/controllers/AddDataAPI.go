package controllers

import (
	"fmt"
	"github.com/m7shapan/njson"
	"io/ioutil"
	"net/http"
)

/*type APIBooks struct {
	Title string
	Author string
	Abstract string
}

type Book struct {
	Title string `json:"title"`
	Author []struct{
		Author string `json:"authors"`
	} `json:"authors"`
	Abstract string `json:"description"`
}

type TmpBook struct {
	items []struct {
		volumeInfo struct {
			title string `json:"title"`
		} `json:"volumeInfo"`
	} `json:"items"`
} */

type APILibrary struct {
	Items []string `njson:"items"`
}

type APIBooks struct {
	Title         string  `njson:"volumeInfo.title"`
	Author        string  `njson:"volumeInfo.authors.0"`
	Abstract      string  `njson:"volumeInfo.description"`
}

func AddDataAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//keyword := r.URL.Query().Get("keyword")
	//apiKey := "&key=AIzaSyCT_Pt_MCIoLBd4BHVWfA1lyd7R1YOWnPw"
	//req := "https://www.googleapis.com/books/v1/volumes?maxResults=40" + keyword + apiKey
	// https://www.googleapis.com/books/v1/volumes?q=test&maxResults=40&key=AIzaSyCT_Pt_MCIoLBd4BHVWfA1lyd7R1YOWnPw

	res, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=test&maxResults=2&key=AIzaSyCT_Pt_MCIoLBd4BHVWfA1lyd7R1YOWnPw")
	if err != nil {
		fmt.Println("Cannot get request!")
	}

	responseData, err := ioutil.ReadAll(res.Body)
	// fmt.Println(string(responseData) + "\n---------------------------------------------\n")
	/* var tmpbook APIBooks
	errorMessage := njson.Unmarshal([]byte(string(responseData)), &tmpbook)
	if err != nil {
		panic(errorMessage)
	} */
	// fmt.Printf("%+v\n\n", tmpbook)
	//fmt.Println(string(responseData))

	var resData APILibrary
	errormessage := njson.Unmarshal([]byte(string(responseData)), &resData)
	if err != nil {
		fmt.Println(errormessage)
	}

	var dataTab []APIBooks

	for _, data := range resData.Items {
		var finalData APIBooks
		errormessage = njson.Unmarshal([]byte(string(data)), &finalData)
		if errormessage != nil {
			fmt.Println(errormessage)
		}
		dataTab = append(dataTab, finalData)

	}

	AddBulkBooks(dataTab)
}
