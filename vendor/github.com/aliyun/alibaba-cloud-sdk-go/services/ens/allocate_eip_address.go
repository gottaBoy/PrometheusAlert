package ens

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

// AllocateEipAddress invokes the ens.AllocateEipAddress API synchronously
// api document: https://help.aliyun.com/api/ens/allocateeipaddress.html
func (client *Client) AllocateEipAddress(request *AllocateEipAddressRequest) (response *AllocateEipAddressResponse, err error) {
	response = CreateAllocateEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateEipAddressWithChan invokes the ens.AllocateEipAddress API asynchronously
// api document: https://help.aliyun.com/api/ens/allocateeipaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateEipAddressWithChan(request *AllocateEipAddressRequest) (<-chan *AllocateEipAddressResponse, <-chan error) {
	responseChan := make(chan *AllocateEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateEipAddress(request)
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

// AllocateEipAddressWithCallback invokes the ens.AllocateEipAddress API asynchronously
// api document: https://help.aliyun.com/api/ens/allocateeipaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateEipAddressWithCallback(request *AllocateEipAddressRequest, callback func(response *AllocateEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateEipAddressResponse
		var err error
		defer close(result)
		response, err = client.AllocateEipAddress(request)
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

// AllocateEipAddressRequest is the request struct for api AllocateEipAddress
type AllocateEipAddressRequest struct {
	*requests.RpcRequest
	EnsRegionId string           `position:"Query" name:"EnsRegionId"`
	Count       requests.Integer `position:"Query" name:"Count"`
	Version     string           `position:"Query" name:"Version"`
}

// AllocateEipAddressResponse is the response struct for api AllocateEipAddress
type AllocateEipAddressResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	EipAddresses EipAddressesInAllocateEipAddress `json:"EipAddresses" xml:"EipAddresses"`
}

// CreateAllocateEipAddressRequest creates a request to invoke AllocateEipAddress API
func CreateAllocateEipAddressRequest() (request *AllocateEipAddressRequest) {
	request = &AllocateEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "AllocateEipAddress", "ens", "openAPI")
	return
}

// CreateAllocateEipAddressResponse creates a response to parse from AllocateEipAddress response
func CreateAllocateEipAddressResponse() (response *AllocateEipAddressResponse) {
	response = &AllocateEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
