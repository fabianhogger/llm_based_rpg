package prompt_model

import (
    "bytes"
    "encoding/json"
    "net/http"
    "log"
    "fmt"
     "io/ioutil"
 
)

type Response struct {
    Message string `json:"generated_text"`
 
}
func Ask(question string) string {
    prompt := map[string]string{"prompt": question}
    jsonData, err := json.Marshal(prompt)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("http://localhost:5000/generate", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return "..."
    }

    // Parse the JSON response
    var response Response
    err = json.Unmarshal(body, &response)
    if err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return "...."
    }

    // Print the parsed response
    fmt.Printf("Message: %s\n", response.Message)
   return string(response.Message)
}
 