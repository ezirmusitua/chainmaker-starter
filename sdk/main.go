package main

import (
	"fmt"
	"os"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"github.com/gin-gonic/gin"
)

const CONFIG_PATH = "/configs/sdk_config.yml"

func InitSdkClient() *sdk.ChainClient {
	client, err := sdk.NewChainClient(sdk.WithConfPath(CONFIG_PATH))

	if err != nil {
		fmt.Printf("Initialize sdk failed: %s", err)
		os.Exit(1)
		return nil
	}

	fmt.Printf("Client initialized\n")
	return client
}

func CallUserContract(
	client *sdk.ChainClient,
	action string,
	contractName string,
	method string,
	params map[string]string,
) (int, map[string]string) {

	kv_pairs := buildParamPairs(params)
	var resp *common.TxResponse
	var err error
	if action == "invoke" {
		resp, err = client.InvokeContract(
			contractName,
			method,
			"",
			kv_pairs,
			-1,
			true, // sync
		)
	} else {
		resp, err = client.QueryContract(
			contractName,
			method,
			kv_pairs,
			-1,
		)
	}

	if err != nil {
		fmt.Printf("[ERROR] Invoke contract failed: %s", err.Error())
		return 502, map[string]string{
			"message": err.Error(),
		}
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return 400, map[string]string{
			"message": resp.ContractResult.Message,
		}
	}

	return 200, map[string]string{
		"message": resp.Message,
		"data":    string(resp.ContractResult.Result),
	}
}

func buildParamPairs(params map[string]string) []*common.KeyValuePair {
	var kv_pairs []*common.KeyValuePair
	for k := range params {
		kv_pairs = append(kv_pairs, &common.KeyValuePair{
			Key:   k,
			Value: []byte(params[k]),
		})
	}
	return kv_pairs
}

type CallContractParamsDto struct {
	Action       string            `json:action`
	ContractName string            `json:contractName`
	Method       string            `json:method`
	Params       map[string]string `json:params`
}

func StartApi(client *sdk.ChainClient) {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	router.POST("/contract", func(c *gin.Context) {
		var dto CallContractParamsDto
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"message": "Invalid Body"})
			return
		}
		code, data := CallUserContract(client, dto.Action, dto.ContractName, dto.Method, dto.Params)
		c.JSON(code, data)
	})

	fmt.Println("Listening on port :8080")
	router.Run(":8080")
}

func main() {
	client := InitSdkClient()
	StartApi(client)
}
