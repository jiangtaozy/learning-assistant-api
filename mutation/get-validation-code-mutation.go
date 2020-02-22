/*
 * Maintained by jemo from 2020.2.15 to now
 * Created by jemo on 2020.2.15 9:31:36
 * Get validation code mutation
 */

package mutation

import (
  "log"
  "time"
  "strconv"
  "math/rand"
  "github.com/satori/go.uuid"
  "github.com/graphql-go/graphql"
  "github.com/GiterLab/aliyun-sms-go-sdk/dysms"
  "github.com/jiangtaozy/learning-assistant-api/db"
  "github.com/jiangtaozy/learning-assistant-api/config"
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
    validationCode, _ := db.RedisClient.Get(phone).Result()
    if validationCode != "" {
      return "已经发送过了", nil
    }
    dysms.HTTPDebugEnable = true
    dysms.SetACLClient(config.AlismsAccessKeyId, config.AlismsAccessKeySecret)
    uid := uuid.NewV4()
    randomNumber := rand.Intn(10000)
    // send sms
    respSendSms, err := dysms.SendSms(uid.String(), phone, "小符问题", "SMS_145594497", `{"code":"` + strconv.Itoa(randomNumber) + `"}`).DoActionWithException()
    if err != nil {
      log.Println("send sms failed", err, respSendSms.Error())
      return "发送失败", nil
    }
    // save sms
    _ = db.RedisClient.Set(phone, strconv.Itoa(randomNumber), 3 * time.Minute).Err()
    return "发送成功", nil
  },
}
