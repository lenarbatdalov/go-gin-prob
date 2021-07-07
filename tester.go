package main

// func main() {
// 	url := "http://localhost:8080/view/videos"

// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJhZ21hdGljIiwiYWRtaW4iOnRydWUsImV4cCI6MTYyNTg5NTAyMSwiaWF0IjoxNjI1NjM1ODIxLCJpc3MiOiJwcmFnbWF0aWNyZXZpZXdzLmNvbSJ9.erZJa6UWS0tlZ2N8tagXnLz502wzv8Wz1qhrW4uMldU")
// 	req.Header.Add("cache-contol", "no-cache")

// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	fmt.Println(res)
// 	fmt.Println(string(body))
// }
