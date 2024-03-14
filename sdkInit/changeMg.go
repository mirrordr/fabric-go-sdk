package sdkInit

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *Application) ChangeMg(args []string) (string, error) {
	var tempArgs [][]byte
	for i := 1; i < len(args); i++ {
		tempArgs = append(tempArgs, []byte(args[i]))
	}

	request := channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := t.SdkEnvInfo.ChClient.Execute(request)
	fmt.Println(err)
	if err != nil {
		return "", err
	}

	return string(response.TransactionID), nil
}
