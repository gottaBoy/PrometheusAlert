package qualitycheck

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

// UploadAudioDataWithRules invokes the qualitycheck.UploadAudioDataWithRules API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadaudiodatawithrules.html
func (client *Client) UploadAudioDataWithRules(request *UploadAudioDataWithRulesRequest) (response *UploadAudioDataWithRulesResponse, err error) {
	response = CreateUploadAudioDataWithRulesResponse()
	err = client.DoAction(request, response)
	return
}

// UploadAudioDataWithRulesWithChan invokes the qualitycheck.UploadAudioDataWithRules API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadaudiodatawithrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadAudioDataWithRulesWithChan(request *UploadAudioDataWithRulesRequest) (<-chan *UploadAudioDataWithRulesResponse, <-chan error) {
	responseChan := make(chan *UploadAudioDataWithRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadAudioDataWithRules(request)
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

// UploadAudioDataWithRulesWithCallback invokes the qualitycheck.UploadAudioDataWithRules API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadaudiodatawithrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadAudioDataWithRulesWithCallback(request *UploadAudioDataWithRulesRequest, callback func(response *UploadAudioDataWithRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadAudioDataWithRulesResponse
		var err error
		defer close(result)
		response, err = client.UploadAudioDataWithRules(request)
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

// UploadAudioDataWithRulesRequest is the request struct for api UploadAudioDataWithRules
type UploadAudioDataWithRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UploadAudioDataWithRulesResponse is the response struct for api UploadAudioDataWithRules
type UploadAudioDataWithRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateUploadAudioDataWithRulesRequest creates a request to invoke UploadAudioDataWithRules API
func CreateUploadAudioDataWithRulesRequest() (request *UploadAudioDataWithRulesRequest) {
	request = &UploadAudioDataWithRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UploadAudioDataWithRules", "", "")
	return
}

// CreateUploadAudioDataWithRulesResponse creates a response to parse from UploadAudioDataWithRules response
func CreateUploadAudioDataWithRulesResponse() (response *UploadAudioDataWithRulesResponse) {
	response = &UploadAudioDataWithRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}