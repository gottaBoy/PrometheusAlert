package aegis

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

// DescribeMachineConfig invokes the aegis.DescribeMachineConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/describemachineconfig.html
func (client *Client) DescribeMachineConfig(request *DescribeMachineConfigRequest) (response *DescribeMachineConfigResponse, err error) {
	response = CreateDescribeMachineConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMachineConfigWithChan invokes the aegis.DescribeMachineConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describemachineconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMachineConfigWithChan(request *DescribeMachineConfigRequest) (<-chan *DescribeMachineConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeMachineConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMachineConfig(request)
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

// DescribeMachineConfigWithCallback invokes the aegis.DescribeMachineConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describemachineconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMachineConfigWithCallback(request *DescribeMachineConfigRequest, callback func(response *DescribeMachineConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMachineConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeMachineConfig(request)
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

// DescribeMachineConfigRequest is the request struct for api DescribeMachineConfig
type DescribeMachineConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Types           string           `position:"Query" name:"Types"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	Type            string           `position:"Query" name:"Type"`
	Lang            string           `position:"Query" name:"Lang"`
	Config          string           `position:"Query" name:"Config"`
	Target          string           `position:"Query" name:"Target"`
}

// DescribeMachineConfigResponse is the response struct for api DescribeMachineConfig
type DescribeMachineConfigResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	Count         int            `json:"Count" xml:"Count"`
	ConfigTargets []ConfigTarget `json:"ConfigTargets" xml:"ConfigTargets"`
}

// CreateDescribeMachineConfigRequest creates a request to invoke DescribeMachineConfig API
func CreateDescribeMachineConfigRequest() (request *DescribeMachineConfigRequest) {
	request = &DescribeMachineConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeMachineConfig", "vipaegis", "openAPI")
	return
}

// CreateDescribeMachineConfigResponse creates a response to parse from DescribeMachineConfig response
func CreateDescribeMachineConfigResponse() (response *DescribeMachineConfigResponse) {
	response = &DescribeMachineConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
