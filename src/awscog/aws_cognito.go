package awscog

import (
	"fmt"

	// read config
	"encoding/json"
	"io/ioutil"

	// aws
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Upi string `json:"userPoolId"`
	CI  string `json:"clientId"`
	Rg  string `json:"region"`
}

func SignIn(user string, pwd string, config string) string {

	// cognito設定ファイル読み込み
	username := user
	password := pwd
	fmt.Println(username)
	c, err := ReadConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	svc := cognitoidentityprovider.New(session.New(), &aws.Config{Region: aws.String(c.Rg)})
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String("ADMIN_NO_SRP_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(c.CI),
		UserPoolId: aws.String(c.Upi),
	}

	res, err := svc.AdminInitiateAuth(params)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return aws.StringValue(res.AuthenticationResult.AccessToken)
}

func SignOut(actoken string, config string) bool {
	c, err := ReadConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	svc := cognitoidentityprovider.New(session.New(), &aws.Config{Region: aws.String(c.Rg)})
	params := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(actoken),
	}
	res, err := svc.GlobalSignOut(params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(res)
	return true
}

// 設定ファイル読み込み
func ReadConfig(filename string) (*Config, error) {
	c := new(Config)

	jsonS, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(jsonS, c)
	if err != nil {
		return c, err
	}

	return c, nil
}
