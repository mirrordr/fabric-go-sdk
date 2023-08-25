package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *Application) CompanyViewRecordRegister(args []string) (string, error) {
	var tempArgs [][]byte
	for i := 1; i < len(args); i++ {
		tempArgs = append(tempArgs, []byte(args[i]))
	}

	request := channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5])}}
	response, err := t.SdkEnvInfo.ChClient.Execute(request)
	if err != nil {
		return "", err
	}

	return string(response.TransactionID), nil
}
