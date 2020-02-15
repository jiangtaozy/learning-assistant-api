/*
 * Maintained by jemo from 2020.1.18 to now
 * Created by jemo on 2020.1.18 14:49:08
 * Main
 */

package main

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/graphql-go/graphql"
  "github.com/jiangtaozy/learning-assistant-api/db"
  "github.com/jiangtaozy/learning-assistant-api/query"
  "github.com/jiangtaozy/learning-assistant-api/mutation"
)

type PostData struct {
  Query string `json:"query"`
  Variables map[string]interface{} `json:"variables"`
}

var port = ":6000"
var schema graphql.Schema

func main() {
  db.Init()
  schema, _ = graphql.NewSchema(
    graphql.SchemaConfig{
      Query: query.Query,
      Mutation: mutation.Mutation,
    },
  )
  log.Println("listen at ", port)
  http.HandleFunc("/graphql", handle)
  log.Fatal(http.ListenAndServe(port, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var data PostData
  err := decoder.Decode(&data)
  if err != nil {
    log.Println("HandleDecodeError: ", err)
  }
  res := graphql.Do(
    graphql.Params{
      Schema: schema,
      RequestString: data.Query,
      VariableValues: data.Variables,
    },
  )
  if len(res.Errors) > 0 {
    log.Printf("HandleResError: %v\n", res.Errors)
  }
  json.NewEncoder(w).Encode(res)
}
