package main

import ("fmt"
        "net/http"
        "encoding/json"
        "io/ioutil"
        "bufio"
        "os"
        "os/exec"
        "strconv")

func make_choice(arrayOfResults []string) (string) {
  reader := bufio.NewReader(os.Stdin)

  // chomp an input from a user via the command line
  fmt.Print("Please pick a number: ")
  choice, _, _ := reader.ReadLine()
  anInteger, _ := strconv.Atoi(string(choice))

  // Open up the url to the image in the browser
  exec.Command("open", "http://i.imgur.com/" + arrayOfResults[anInteger] + ".png").Start()
  return "Opening up your choice in a browser..."
}

func main() {
  // Make a request to the imgur.com API
  response, _ := http.Get("http://imgur.com/gallery.json")

  // Read the response body into a variable
  jsonBody, _ := ioutil.ReadAll(response.Body)

  // Close the request stream
  response.Body.Close()

  // Save the result to a map
  mappedResult := make(map[string]interface{},0)

  // Check for errors
  err := json.Unmarshal([]byte(jsonBody), &mappedResult)
  if err != nil {
    panic(err)
  }

  // Pull out the entire JSON 'data' object hash
  parents := mappedResult["data"].([]interface{})

  // Set up an array that stores the alphanumeric hashes of imgur image addresses.
  arrayOfResults := make([]string,0)

  // Get each instance in the array of JSON hashes
  // Print out the title
  // Append it to the array created on :38
  for k, v := range parents {
    instance := v.(map[string]interface{})
    fmt.Printf("%d - %s\n", k, instance["title"].(string))
    arrayOfResults = append(arrayOfResults, instance["hash"].(string))
  }

  // Ask user to make a choice on the command line
  make_choice(arrayOfResults)
}