package dm

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

// ApproveReplyMailAddress invokes the dm.ApproveReplyMailAddress API synchronously
// api document: https://help.aliyun.com/api/dm/approvereplymailaddress.html
func (client *Client) ApproveReplyMailAddress(request *ApproveReplyMailAddressRequest) (response *ApproveReplyMailAddressResponse, err error) {
	response = CreateApproveReplyMailAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ApproveReplyMailAddressWithChan invokes the dm.ApproveReplyMailAddress API asynchronously
// api document: https://help.aliyun.com/api/dm/approvereplymailaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApproveReplyMailAddressWithChan(request *ApproveReplyMailAddressRequest) (<-chan *ApproveReplyMailAddressResponse, <-chan error) {
	responseChan := make(chan *ApproveReplyMailAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApproveReplyMailAddress(request)
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

// ApproveReplyMailAddressWithCallback invokes the dm.ApproveReplyMailAddress API asynchronously
// api document: https://help.aliyun.com/api/dm/approvereplymailaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApproveReplyMailAddressWithCallback(request *ApproveReplyMailAddressRequest, callback func(response *ApproveReplyMailAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApproveReplyMailAddressResponse
		var err error
		defer close(result)
		response, err = client.ApproveReplyMailAddress(request)
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

// ApproveReplyMailAddressRequest is the request struct for api ApproveReplyMailAddress
type ApproveReplyMailAddressRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Ticket               string           `position:"Query" name:"Ticket"`
}

// ApproveReplyMailAddressResponse is the response struct for api ApproveReplyMailAddress
type ApproveReplyMailAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateApproveReplyMailAddressRequest creates a request to invoke ApproveReplyMailAddress API
func CreateApproveReplyMailAddressRequest() (request *ApproveReplyMailAddressRequest) {
	request = &ApproveReplyMailAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "ApproveReplyMailAddress", "", "")
	return
}

// CreateApproveReplyMailAddressResponse creates a response to parse from ApproveReplyMailAddress response
func CreateApproveReplyMailAddressResponse() (response *ApproveReplyMailAddressResponse) {
	response = &ApproveReplyMailAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
