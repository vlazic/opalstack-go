/*
Opalstack API

 ## The Opalstack JSON REST API  ### Authorization This API uses an **Authorization** header of the form: `\"Authorization: Token 1111111111111111111111111111111111111111\"`, where **1111111111111111111111111111111111111111** represents an API token created at https://my.opalstack.com/tokens/.  The typical format of an API request looks like the following: ``` GET request:     curl -s -H \"Authorization: Token 1111111111111111111111111111111111111111\" \"https://my.opalstack.com/api/v1/site/list/\" | jq .  POST request:     curl -s -H \"Content-Type: application/json\" -H \"Authorization: Token 1111111111111111111111111111111111111111\" \\             -X POST -d '[{\"id\": \"(site UUID)\", \"redirect\": true, ...}]' \"https://my.opalstack.com/api/v1/site/update/\" | jq . ``` (Further examples will omit **headers** and **jq** for the sake of clarity)  You can also authorize requests on our API Documentation page (https://my.opalstack.com/api/v1/doc/) in order to facilitate development. To do so, click the \"**Authorize**\" button on the right side of the page and enter \"**Token 1111111111111111111111111111111111111111**\" in the **Value** field within. Afterword, you will be able to perform requests directly from the documentation page. Be sure to logout when finished.  ### Embedding The Opalstack API supports _embedding_. This allows you to nest child API objects in a single GET request. For example, consider the following GET request performed with **curl**: ``` Request:     curl \"https://my.opalstack.com/api/v1/osuser/list/\"  Response:     [       {         \"id\": \"01010101-0202-0303-0404-050505050505\",         \"state\": \"READY\",         \"ready\": true,         \"name\": \"the_osuser_name\",         \"server\": \"11111111-1212-1313-1414-151515151515\"       }     ] ```  Suppose then that we would like additional information about the **server**. We _could_ proceed to query the **server** UUID (**11111111-1212-1313-1414-151515151515**) at the `/server/read/{uuid}` endpoint, like this: ``` Request:     curl \"https://my.opalstack.com/api/v1/server/read/11111111-1212-1313-1414-151515151515\"  Response:     {       \"id\": \"11111111-1212-1313-1414-151515151515\",       \"hostname\": \"vpsNNN.opalstack.com\"     } ```  However, we could have instead choosen to specify `?embed=server` as a query parameter to the original GET request. This will cause objects to be _embedded_ in the response directly: ``` Request:     curl \"https://my.opalstack.com/api/v1/osuser/list/?embed=server\"  Response:     [       {         \"id\": \"01010101-0202-0303-0404-050505050505\",         \"state\": \"READY\",         \"ready\": true,         \"name\": \"the_osuser_name\",         \"server\": {           \"id\": \"11111111-1212-1313-1414-151515151515\",           \"hostname\": \"vpsNNN.opalstack.com\"         }       }     ] ``` Here, the **server** field has been be populated with the full object instead of just a UUID.  The **embed** query parameter accepts multiple (comma-separated) fields to embed. For example: ``` curl \"https://my.opalstack.com/api/v1/account/info/?embed=web_servers,imap_servers,smtp_servers\" ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OpalStack

import (
	"encoding/json"
	"time"
)

// checks if the WebUsageMariaDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebUsageMariaDBResponse{}

// WebUsageMariaDBResponse struct for WebUsageMariaDBResponse
type WebUsageMariaDBResponse struct {
	Mariadb WebUsageMariaDB `json:"mariadb"`
	DiskUsageMb int32 `json:"disk_usage_mb"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewWebUsageMariaDBResponse instantiates a new WebUsageMariaDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebUsageMariaDBResponse(mariadb WebUsageMariaDB, diskUsageMb int32, updatedAt time.Time) *WebUsageMariaDBResponse {
	this := WebUsageMariaDBResponse{}
	this.Mariadb = mariadb
	this.DiskUsageMb = diskUsageMb
	this.UpdatedAt = updatedAt
	return &this
}

// NewWebUsageMariaDBResponseWithDefaults instantiates a new WebUsageMariaDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebUsageMariaDBResponseWithDefaults() *WebUsageMariaDBResponse {
	this := WebUsageMariaDBResponse{}
	return &this
}

// GetMariadb returns the Mariadb field value
func (o *WebUsageMariaDBResponse) GetMariadb() WebUsageMariaDB {
	if o == nil {
		var ret WebUsageMariaDB
		return ret
	}

	return o.Mariadb
}

// GetMariadbOk returns a tuple with the Mariadb field value
// and a boolean to check if the value has been set.
func (o *WebUsageMariaDBResponse) GetMariadbOk() (*WebUsageMariaDB, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mariadb, true
}

// SetMariadb sets field value
func (o *WebUsageMariaDBResponse) SetMariadb(v WebUsageMariaDB) {
	o.Mariadb = v
}

// GetDiskUsageMb returns the DiskUsageMb field value
func (o *WebUsageMariaDBResponse) GetDiskUsageMb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiskUsageMb
}

// GetDiskUsageMbOk returns a tuple with the DiskUsageMb field value
// and a boolean to check if the value has been set.
func (o *WebUsageMariaDBResponse) GetDiskUsageMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskUsageMb, true
}

// SetDiskUsageMb sets field value
func (o *WebUsageMariaDBResponse) SetDiskUsageMb(v int32) {
	o.DiskUsageMb = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WebUsageMariaDBResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WebUsageMariaDBResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WebUsageMariaDBResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o WebUsageMariaDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebUsageMariaDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mariadb"] = o.Mariadb
	toSerialize["disk_usage_mb"] = o.DiskUsageMb
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableWebUsageMariaDBResponse struct {
	value *WebUsageMariaDBResponse
	isSet bool
}

func (v NullableWebUsageMariaDBResponse) Get() *WebUsageMariaDBResponse {
	return v.value
}

func (v *NullableWebUsageMariaDBResponse) Set(val *WebUsageMariaDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebUsageMariaDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebUsageMariaDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebUsageMariaDBResponse(val *WebUsageMariaDBResponse) *NullableWebUsageMariaDBResponse {
	return &NullableWebUsageMariaDBResponse{value: val, isSet: true}
}

func (v NullableWebUsageMariaDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebUsageMariaDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

