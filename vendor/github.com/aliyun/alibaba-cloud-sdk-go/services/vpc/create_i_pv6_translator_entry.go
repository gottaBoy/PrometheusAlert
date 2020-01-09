package vpc

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

// CreateIPv6TranslatorEntry invokes the vpc.CreateIPv6TranslatorEntry API synchronously
// api document: https://help.aliyun.com/api/vpc/createipv6translatorentry.html
func (client *Client) CreateIPv6TranslatorEntry(request *CreateIPv6TranslatorEntryRequest) (response *CreateIPv6TranslatorEntryResponse, err error) {
	response = CreateCreateIPv6TranslatorEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIPv6TranslatorEntryWithChan invokes the vpc.CreateIPv6TranslatorEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/createipv6translatorentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIPv6TranslatorEntryWithChan(request *CreateIPv6TranslatorEntryRequest) (<-chan *CreateIPv6TranslatorEntryResponse, <-chan error) {
	responseChan := make(chan *CreateIPv6TranslatorEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIPv6TranslatorEntry(request)
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

// CreateIPv6TranslatorEntryWithCallback invokes the vpc.CreateIPv6TranslatorEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/createipv6translatorentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIPv6TranslatorEntryWithCallback(request *CreateIPv6TranslatorEntryRequest, callback func(response *CreateIPv6TranslatorEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIPv6TranslatorEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateIPv6TranslatorEntry(request)
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

// CreateIPv6TranslatorEntryRequest is the request struct for api CreateIPv6TranslatorEntry
type CreateIPv6TranslatorEntryRequest struct {
	*requests.RpcRequest
	BackendIpv4Port      requests.Integer `position:"Query" name:"BackendIpv4Port"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EntryName            string           `position:"Query" name:"EntryName"`
	AclStatus            string           `position:"Query" name:"AclStatus"`
	EntryBandwidth       requests.Integer `position:"Query" name:"EntryBandwidth"`
	AclType              string           `position:"Query" name:"AclType"`
	AllocateIpv6Port     requests.Integer `position:"Query" name:"AllocateIpv6Port"`
	EntryDescription     string           `position:"Query" name:"EntryDescription"`
	BackendIpv4Addr      string           `position:"Query" name:"BackendIpv4Addr"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TransProtocol        string           `position:"Query" name:"TransProtocol"`
	Ipv6TranslatorId     string           `position:"Query" name:"Ipv6TranslatorId"`
}

// CreateIPv6TranslatorEntryResponse is the response struct for api CreateIPv6TranslatorEntry
type CreateIPv6TranslatorEntryResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	Ipv6TranslatorEntryId string `json:"Ipv6TranslatorEntryId" xml:"Ipv6TranslatorEntryId"`
}

// CreateCreateIPv6TranslatorEntryRequest creates a request to invoke CreateIPv6TranslatorEntry API
func CreateCreateIPv6TranslatorEntryRequest() (request *CreateIPv6TranslatorEntryRequest) {
	request = &CreateIPv6TranslatorEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateIPv6TranslatorEntry", "vpc", "openAPI")
	return
}

// CreateCreateIPv6TranslatorEntryResponse creates a response to parse from CreateIPv6TranslatorEntry response
func CreateCreateIPv6TranslatorEntryResponse() (response *CreateIPv6TranslatorEntryResponse) {
	response = &CreateIPv6TranslatorEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
