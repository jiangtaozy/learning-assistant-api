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
)

type PostData struct {
  Query string `json:"query"`
  Variables map[string]interface{} `json:"variables"`
}

var port = ":6000"
var schema graphql.Schema

func main() {
  query := graphql.NewObject(
    graphql.ObjectConfig{
      Name: "query",
      Fields: graphql.Fields{
        "hello": &graphql.Field{
          Type: graphql.String,
          Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            return "world", nil
          },
        },
      },
    },
  )
  mutation := graphql.NewObject(
    graphql.ObjectConfig{
      Name: "mutation",
      Fields: graphql.Fields{
        "create": &graphql.Field{
          Type: graphql.String,
          Args: graphql.FieldConfigArgument{
            "text": &graphql.ArgumentConfig{
              Type: graphql.NewNonNull(
                graphql.String,
              ),
            },
          },
          Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            text, _ := params.Args["text"].(string)
            return text, nil
          },
        },
      },
    },
  )
  schema, _ = graphql.NewSchema(
    graphql.SchemaConfig{
      Query: query,
      Mutation: mutation,
    },
  )
  log.Println("listen at ", port)
  http.HandleFunc("/", handle)
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
