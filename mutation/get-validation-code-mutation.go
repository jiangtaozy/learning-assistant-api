/*
 * Maintained by jemo from 2020.2.15 to now
 * Created by jemo on 2020.2.15 9:31:36
 * Get validation code mutation
 */

package mutation

import (
  "log"
  "context"
  "github.com/graphql-go/graphql"
  "github.com/learning-assistant-api/db"
)

var GetValidationCodeMutation = &graphql.Field{
  Type: graphql.String,
  Args: graphql.FieldConfigArgument{
    "phone": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(
        graphql.String,
      ),
    },
  },
  Resolve: func(params graphql.ResolveParams) (interface{}, error) {
    phone, _ := params.Args["phone"].(string)
    log.Println("phone: ", phone)
    ctx, cancel := context.WithTimeout(
      context.Background(),
      10 * time.Second,
    )
    defer cancel()
    resp, err := db.Etcd.Get(ctx, phone)
    if err != nil {
      log.Printf("error: %v\n", err)
    }
    log.Printf("resp.Kvs: %v\n", resp.Kvs)
    if len(resp.Kvs) > 0 {
      log.Println("已经发送")
    }
    return "ok", nil
  },
}
