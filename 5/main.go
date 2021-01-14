package main

import (
  "context"
  "crypto/tls"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"

  eventhub "github.com/Azure/azure-event-hubs-go"
  "github.com/go-redis/redis"
  "github.com/sirupsen/logrus"
)

type postRequest struct {
  Url      string `json:"url"`
  Strategy string `json:"strategy"`
}

func getJson(url string) []map[string]string {
  getClient := http.Client{
    Timeout: time.Second * 2, // Timeout after 2 seconds
  }

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    logrus.Fatal(err)
  }

  res, err := getClient.Do(req)
  if err != nil {
    logrus.Fatal(err)
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    logrus.Fatal(err)
  }

  var jsonArr []map[string]string

  err = json.Unmarshal(body, &jsonArr)
  if err != nil {
    logrus.Fatal(err)
  }

  return jsonArr
}

func operateRedis(url string) {
  dataJson := getJson(url)
  client := redis.NewClient(&redis.Options{
    Addr:      "lab5iot.redis.cache.windows.net:6380",
    Password:  "cRDGgYb6KqJCgJw0JTsmsD1jxzYU3tXIDOhiVe4VtIQ=",
    DB:        0,
    TLSConfig: &tls.Config{InsecureSkipVerify: true},
  })

  _, err := client.Ping().Result()
  if err != nil {
    logrus.Fatal(err)
  }
// send blocks of info
  for i, data := range dataJson {
    dataSingleJson, err := json.Marshal(data)
    if err != nil {
      logrus.Fatal(err)
    }
    err = client.Set(fmt.Sprintf("data_%d", i), dataSingleJson, 0).Err()
    if err != nil {
      logrus.Warn(err)
    } else {
      logrus.Infof("Document %d is written!", i)
    }
  }
}

func operateEventHub(url string) {
  // get Json
  dataJson := getJson(url)
  connStr := "Endpoint=sb://lab5iot.servicebus.windows.net/;SharedAccessKeyName=lab5iot;SharedAccessKey=kAYHFurz0T1fJ0MUfPdyhuFxEjP7AllmfKDvMZuhb18=;EntityPath=lab5iot"

  hub, err := eventhub.NewHubFromConnectionString(connStr)
  if err != nil {
    logrus.Info(err)
  }

  ctx := context.Background()

  // send a single message 
  for i, data := range dataJson {
    dataSingleJson, err := json.Marshal(data)
    if err != nil {
      logrus.Fatal(err)
    }
    event := eventhub.NewEvent(dataSingleJson)
    event.Set("content_type", "application/json")
    err = hub.Send(ctx, event)
    if err != nil {
      logrus.Warn(err)
    } else {
      logrus.Infof("Document %d was sent!", i)
    }
  }

  err = hub.Close(context.Background())
  if err != nil {
    logrus.Info(err)
  }

  logrus.Info("Json was sent to EventHub!")
}
//pars json 
func HelloServer(w http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var request postRequest
  err := decoder.Decode(&request)
  if err != nil {
    panic(err)
  }

  fmt.Println(request.Strategy)
//choose stratagy to send data
  switch {
  case request.Strategy == "redis":
    operateRedis(request.Url)
  case request.Strategy == "eventHub":
    operateEventHub(request.Url)
  default:
    logrus.Info("Wrong strategy choosed!")
  }
}
//start server 
func main() {
  http.HandleFunc("/url", HelloServer)
  logrus.Fatal(http.ListenAndServe(":20000", nil))
}