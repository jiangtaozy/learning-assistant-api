/*
 * Maintained by jemo from 2020.2.15 to now
 * Created by jemo on 2020.2.15 10:08:00
 * Mutation
 */

package mutation

import (
  "github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(
    graphql.ObjectConfig{
      Name: "mutation",
      Fields: graphql.Fields{
        "getValidationCode": GetValidationCodeMutation,
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
