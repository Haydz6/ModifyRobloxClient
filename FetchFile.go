package main

import "net/http"

func FetchFile(URL string) (bool, *http.Response) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		println(err.Error())
		return false, nil
	}

	response, err := client.Do(request)

	if err != nil {
		println(err.Error())
		return false, nil
	}

	return true, response
}
