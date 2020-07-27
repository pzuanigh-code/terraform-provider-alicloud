package smartag

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifySagWanSnat invokes the smartag.ModifySagWanSnat API synchronously
// api document: https://help.aliyun.com/api/smartag/modifysagwansnat.html
func (client *Client) ModifySagWanSnat(request *ModifySagWanSnatRequest) (response *ModifySagWanSnatResponse, err error) {
	response = CreateModifySagWanSnatResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySagWanSnatWithChan invokes the smartag.ModifySagWanSnat API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifysagwansnat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySagWanSnatWithChan(request *ModifySagWanSnatRequest) (<-chan *ModifySagWanSnatResponse, <-chan error) {
	responseChan := make(chan *ModifySagWanSnatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySagWanSnat(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifySagWanSnatWithCallback invokes the smartag.ModifySagWanSnat API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifysagwansnat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySagWanSnatWithCallback(request *ModifySagWanSnatRequest, callback func(response *ModifySagWanSnatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySagWanSnatResponse
		var err error
		defer close(result)
		response, err = client.ModifySagWanSnat(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifySagWanSnatRequest is the request struct for api ModifySagWanSnat
type ModifySagWanSnatRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
	Snat                 string           `position:"Query" name:"Snat"`
}

// ModifySagWanSnatResponse is the response struct for api ModifySagWanSnat
type ModifySagWanSnatResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySagWanSnatRequest creates a request to invoke ModifySagWanSnat API
func CreateModifySagWanSnatRequest() (request *ModifySagWanSnatRequest) {
	request = &ModifySagWanSnatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySagWanSnat", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySagWanSnatResponse creates a response to parse from ModifySagWanSnat response
func CreateModifySagWanSnatResponse() (response *ModifySagWanSnatResponse) {
	response = &ModifySagWanSnatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}