package common

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
	simpleMessage "github.com/jeronimobarea/simple-ethereum-client/message"
)

type (
	variablesInput struct {
		Mainnet   string `json:"mainnet"`
		Owner     string `json:"owner"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	jwtVariables struct {
		Input *variablesInput `json:"input"`
	}

	jwtPayload struct {
		OperationName string        `json:"operationName"`
		Variables     *jwtVariables `json:"variables"`
		Query         string        `json:"query"`
	}
)

func CreateRandomMessage() (message string, err error) {
	payload := jwtPayload{
		OperationName: "CreateRandomMessage",
		Query:         "mutation CreateRandomMessage{createRandomMessage}",
	}
	parsedPayload, err := json.Marshal(&payload)
	if err != nil {
		return
	}

	requestBody := bytes.NewBuffer(parsedPayload)

	resp, err := http.Post(
		RoninGraphqlUri,
		"application/json",
		requestBody,
	)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data map[string]map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	var ok bool
	if message, ok = data["data"]["createRandomMessage"]; !ok {
		err = errors.New("error checking dict keys")
		return
	}
	return
}

func GetJWT(privateKey *ecdsa.PrivateKey, message string) (token string, err error) {
	signedMessage, err := simpleMessage.SignMessage(privateKey, message)
	if err != nil {
		return
	}

	ownerAddress, err := simpleCommon.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return
	}

	payload := jwtPayload{
		OperationName: "CreateAccessTokenWithSignature",
		Variables: &jwtVariables{
			Input: &variablesInput{
				Mainnet:   "ronin",
				Owner:     ownerAddress.Hex(),
				Message:   message,
				Signature: hexutil.Encode(signedMessage),
			},
		},
		Query: "mutation CreateAccessTokenWithSignature($input: SignatureInput!)" +
			"{createAccessTokenWithSignature(input: $input) " +
			"{newAccount result accessToken __typename}}",
	}

	parsedPayload, err := json.Marshal(&payload)
	if err != nil {
		return
	}

	requestBody := bytes.NewBuffer(parsedPayload)

	resp, err := http.Post(
		RoninGraphqlUri,
		"application/json",
		requestBody,
	)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data map[string]map[string]map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	var ok bool
	if token, ok = data["data"]["createAccessTokenWithSignature"]["accessToken"]; !ok {
		err = errors.New("error checking dict keys")
		return
	}
	return
}
