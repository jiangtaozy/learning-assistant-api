/*
 * Maintained by jemo from 2020.2.15 to now
 * Created by jemo on 2020.2.15 10:14:36
 * Query
 */

package query

import (
  "github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(
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
