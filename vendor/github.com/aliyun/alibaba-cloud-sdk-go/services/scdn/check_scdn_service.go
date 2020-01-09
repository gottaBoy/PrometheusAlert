package scdn

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

// CheckScdnService invokes the scdn.CheckScdnService API synchronously
// api document: https://help.aliyun.com/api/scdn/checkscdnservice.html
func (client *Client) CheckScdnService(request *CheckScdnServiceRequest) (response *CheckScdnServiceResponse, err error) {
	response = CreateCheckScdnServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CheckScdnServiceWithChan invokes the scdn.CheckScdnService API asynchronously
// api document: https://help.aliyun.com/api/scdn/checkscdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckScdnServiceWithChan(request *CheckScdnServiceRequest) (<-chan *CheckScdnServiceResponse, <-chan error) {
	responseChan := make(chan *CheckScdnServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckScdnService(request)
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

// CheckScdnServiceWithCallback invokes the scdn.CheckScdnService API asynchronously
// api document: https://help.aliyun.com/api/scdn/checkscdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckScdnServiceWithCallback(request *CheckScdnServiceRequest, callback func(response *CheckScdnServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckScdnServiceResponse
		var err error
		defer close(result)
		response, err = client.CheckScdnService(request)
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

// CheckScdnServiceRequest is the request struct for api CheckScdnService
type CheckScdnServiceRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// CheckScdnServiceResponse is the response struct for api CheckScdnService
type CheckScdnServiceResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Enabled       bool   `json:"Enabled" xml:"Enabled"`
	OnService     bool   `json:"OnService" xml:"OnService"`
	InDebt        bool   `json:"InDebt" xml:"InDebt"`
	InDebtOverdue bool   `json:"InDebtOverdue" xml:"InDebtOverdue"`
}

// CreateCheckScdnServiceRequest creates a request to invoke CheckScdnService API
func CreateCheckScdnServiceRequest() (request *CheckScdnServiceRequest) {
	request = &CheckScdnServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "CheckScdnService", "", "")
	return
}

// CreateCheckScdnServiceResponse creates a response to parse from CheckScdnService response
func CreateCheckScdnServiceResponse() (response *CheckScdnServiceResponse) {
	response = &CheckScdnServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
