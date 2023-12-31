/*
Opalstack API

 ## The Opalstack JSON REST API  ### Authorization This API uses an **Authorization** header of the form: `\"Authorization: Token 1111111111111111111111111111111111111111\"`, where **1111111111111111111111111111111111111111** represents an API token created at https://my.opalstack.com/tokens/.  The typical format of an API request looks like the following: ``` GET request:     curl -s -H \"Authorization: Token 1111111111111111111111111111111111111111\" \"https://my.opalstack.com/api/v1/site/list/\" | jq .  POST request:     curl -s -H \"Content-Type: application/json\" -H \"Authorization: Token 1111111111111111111111111111111111111111\" \\             -X POST -d '[{\"id\": \"(site UUID)\", \"redirect\": true, ...}]' \"https://my.opalstack.com/api/v1/site/update/\" | jq . ``` (Further examples will omit **headers** and **jq** for the sake of clarity)  You can also authorize requests on our API Documentation page (https://my.opalstack.com/api/v1/doc/) in order to facilitate development. To do so, click the \"**Authorize**\" button on the right side of the page and enter \"**Token 1111111111111111111111111111111111111111**\" in the **Value** field within. Afterword, you will be able to perform requests directly from the documentation page. Be sure to logout when finished.  ### Embedding The Opalstack API supports _embedding_. This allows you to nest child API objects in a single GET request. For example, consider the following GET request performed with **curl**: ``` Request:     curl \"https://my.opalstack.com/api/v1/osuser/list/\"  Response:     [       {         \"id\": \"01010101-0202-0303-0404-050505050505\",         \"state\": \"READY\",         \"ready\": true,         \"name\": \"the_osuser_name\",         \"server\": \"11111111-1212-1313-1414-151515151515\"       }     ] ```  Suppose then that we would like additional information about the **server**. We _could_ proceed to query the **server** UUID (**11111111-1212-1313-1414-151515151515**) at the `/server/read/{uuid}` endpoint, like this: ``` Request:     curl \"https://my.opalstack.com/api/v1/server/read/11111111-1212-1313-1414-151515151515\"  Response:     {       \"id\": \"11111111-1212-1313-1414-151515151515\",       \"hostname\": \"vpsNNN.opalstack.com\"     } ```  However, we could have instead choosen to specify `?embed=server` as a query parameter to the original GET request. This will cause objects to be _embedded_ in the response directly: ``` Request:     curl \"https://my.opalstack.com/api/v1/osuser/list/?embed=server\"  Response:     [       {         \"id\": \"01010101-0202-0303-0404-050505050505\",         \"state\": \"READY\",         \"ready\": true,         \"name\": \"the_osuser_name\",         \"server\": {           \"id\": \"11111111-1212-1313-1414-151515151515\",           \"hostname\": \"vpsNNN.opalstack.com\"         }       }     ] ``` Here, the **server** field has been be populated with the full object instead of just a UUID.  The **embed** query parameter accepts multiple (comma-separated) fields to embed. For example: ``` curl \"https://my.opalstack.com/api/v1/account/info/?embed=web_servers,imap_servers,smtp_servers\" ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OpalStack

import (
	"encoding/json"
)

// checks if the UsageDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageDataResponse{}

// UsageDataResponse struct for UsageDataResponse
type UsageDataResponse struct {
	ImapTotal int32 `json:"imap_total"`
	ImapUsed int32 `json:"imap_used"`
	ImapAvailable int32 `json:"imap_available"`
	Webservers WebserverIdUsageResponse `json:"webservers"`
}

// NewUsageDataResponse instantiates a new UsageDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageDataResponse(imapTotal int32, imapUsed int32, imapAvailable int32, webservers WebserverIdUsageResponse) *UsageDataResponse {
	this := UsageDataResponse{}
	this.ImapTotal = imapTotal
	this.ImapUsed = imapUsed
	this.ImapAvailable = imapAvailable
	this.Webservers = webservers
	return &this
}

// NewUsageDataResponseWithDefaults instantiates a new UsageDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageDataResponseWithDefaults() *UsageDataResponse {
	this := UsageDataResponse{}
	return &this
}

// GetImapTotal returns the ImapTotal field value
func (o *UsageDataResponse) GetImapTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImapTotal
}

// GetImapTotalOk returns a tuple with the ImapTotal field value
// and a boolean to check if the value has been set.
func (o *UsageDataResponse) GetImapTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImapTotal, true
}

// SetImapTotal sets field value
func (o *UsageDataResponse) SetImapTotal(v int32) {
	o.ImapTotal = v
}

// GetImapUsed returns the ImapUsed field value
func (o *UsageDataResponse) GetImapUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImapUsed
}

// GetImapUsedOk returns a tuple with the ImapUsed field value
// and a boolean to check if the value has been set.
func (o *UsageDataResponse) GetImapUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImapUsed, true
}

// SetImapUsed sets field value
func (o *UsageDataResponse) SetImapUsed(v int32) {
	o.ImapUsed = v
}

// GetImapAvailable returns the ImapAvailable field value
func (o *UsageDataResponse) GetImapAvailable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImapAvailable
}

// GetImapAvailableOk returns a tuple with the ImapAvailable field value
// and a boolean to check if the value has been set.
func (o *UsageDataResponse) GetImapAvailableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImapAvailable, true
}

// SetImapAvailable sets field value
func (o *UsageDataResponse) SetImapAvailable(v int32) {
	o.ImapAvailable = v
}

// GetWebservers returns the Webservers field value
func (o *UsageDataResponse) GetWebservers() WebserverIdUsageResponse {
	if o == nil {
		var ret WebserverIdUsageResponse
		return ret
	}

	return o.Webservers
}

// GetWebserversOk returns a tuple with the Webservers field value
// and a boolean to check if the value has been set.
func (o *UsageDataResponse) GetWebserversOk() (*WebserverIdUsageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Webservers, true
}

// SetWebservers sets field value
func (o *UsageDataResponse) SetWebservers(v WebserverIdUsageResponse) {
	o.Webservers = v
}

func (o UsageDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imap_total"] = o.ImapTotal
	toSerialize["imap_used"] = o.ImapUsed
	toSerialize["imap_available"] = o.ImapAvailable
	toSerialize["webservers"] = o.Webservers
	return toSerialize, nil
}

type NullableUsageDataResponse struct {
	value *UsageDataResponse
	isSet bool
}

func (v NullableUsageDataResponse) Get() *UsageDataResponse {
	return v.value
}

func (v *NullableUsageDataResponse) Set(val *UsageDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageDataResponse(val *UsageDataResponse) *NullableUsageDataResponse {
	return &NullableUsageDataResponse{value: val, isSet: true}
}

func (v NullableUsageDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


