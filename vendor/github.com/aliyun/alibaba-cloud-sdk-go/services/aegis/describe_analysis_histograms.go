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

// DescribeAnalysisHistograms invokes the aegis.DescribeAnalysisHistograms API synchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysishistograms.html
func (client *Client) DescribeAnalysisHistograms(request *DescribeAnalysisHistogramsRequest) (response *DescribeAnalysisHistogramsResponse, err error) {
	response = CreateDescribeAnalysisHistogramsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAnalysisHistogramsWithChan invokes the aegis.DescribeAnalysisHistograms API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysishistograms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAnalysisHistogramsWithChan(request *DescribeAnalysisHistogramsRequest) (<-chan *DescribeAnalysisHistogramsResponse, <-chan error) {
	responseChan := make(chan *DescribeAnalysisHistogramsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAnalysisHistograms(request)
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

// DescribeAnalysisHistogramsWithCallback invokes the aegis.DescribeAnalysisHistograms API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysishistograms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAnalysisHistogramsWithCallback(request *DescribeAnalysisHistogramsRequest, callback func(response *DescribeAnalysisHistogramsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAnalysisHistogramsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAnalysisHistograms(request)
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

// DescribeAnalysisHistogramsRequest is the request struct for api DescribeAnalysisHistograms
type DescribeAnalysisHistogramsRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	Query    string           `position:"Query" name:"Query"`
	From     requests.Integer `position:"Query" name:"From"`
	To       requests.Integer `position:"Query" name:"To"`
}

// DescribeAnalysisHistogramsResponse is the response struct for api DescribeAnalysisHistograms
type DescribeAnalysisHistogramsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Histograms Histograms `json:"Histograms" xml:"Histograms"`
}

// CreateDescribeAnalysisHistogramsRequest creates a request to invoke DescribeAnalysisHistograms API
func CreateDescribeAnalysisHistogramsRequest() (request *DescribeAnalysisHistogramsRequest) {
	request = &DescribeAnalysisHistogramsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeAnalysisHistograms", "vipaegis", "openAPI")
	return
}

// CreateDescribeAnalysisHistogramsResponse creates a response to parse from DescribeAnalysisHistograms response
func CreateDescribeAnalysisHistogramsResponse() (response *DescribeAnalysisHistogramsResponse) {
	response = &DescribeAnalysisHistogramsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
