package cloudapi

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

// DescribeApp invokes the cloudapi.DescribeApp API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapp.html
func (client *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
	response = CreateDescribeAppResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppWithChan invokes the cloudapi.DescribeApp API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppWithChan(request *DescribeAppRequest) (<-chan *DescribeAppResponse, <-chan error) {
	responseChan := make(chan *DescribeAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApp(request)
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

// DescribeAppWithCallback invokes the cloudapi.DescribeApp API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppWithCallback(request *DescribeAppRequest, callback func(response *DescribeAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppResponse
		var err error
		defer close(result)
		response, err = client.DescribeApp(request)
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

// DescribeAppRequest is the request struct for api DescribeApp
type DescribeAppRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
}

// DescribeAppResponse is the response struct for api DescribeApp
type DescribeAppResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	AppId        int64  `json:"AppId" xml:"AppId"`
	AppName      string `json:"AppName" xml:"AppName"`
	Description  string `json:"Description" xml:"Description"`
	CreatedTime  string `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime string `json:"ModifiedTime" xml:"ModifiedTime"`
}

// CreateDescribeAppRequest creates a request to invoke DescribeApp API
func CreateDescribeAppRequest() (request *DescribeAppRequest) {
	request = &DescribeAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApp", "apigateway", "openAPI")
	return
}

// CreateDescribeAppResponse creates a response to parse from DescribeApp response
func CreateDescribeAppResponse() (response *DescribeAppResponse) {
	response = &DescribeAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
