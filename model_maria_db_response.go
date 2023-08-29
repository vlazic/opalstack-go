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

// checks if the MariaDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MariaDBResponse{}

// MariaDBResponse struct for MariaDBResponse
type MariaDBResponse struct {
	Id string `json:"id"`
	State StateEnum `json:"state"`
	Ready bool `json:"ready"`
	Name string `json:"name"`
	Server string `json:"server"`
	Charset MariaCharset `json:"charset"`
	DbusersReadwrite []string `json:"dbusers_readwrite"`
	DbusersReadonly []string `json:"dbusers_readonly"`
}

// NewMariaDBResponse instantiates a new MariaDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMariaDBResponse(id string, state StateEnum, ready bool, name string, server string, charset MariaCharset, dbusersReadwrite []string, dbusersReadonly []string) *MariaDBResponse {
	this := MariaDBResponse{}
	this.Id = id
	this.State = state
	this.Ready = ready
	this.Name = name
	this.Server = server
	this.Charset = charset
	this.DbusersReadwrite = dbusersReadwrite
	this.DbusersReadonly = dbusersReadonly
	return &this
}

// NewMariaDBResponseWithDefaults instantiates a new MariaDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMariaDBResponseWithDefaults() *MariaDBResponse {
	this := MariaDBResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MariaDBResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MariaDBResponse) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *MariaDBResponse) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *MariaDBResponse) SetState(v StateEnum) {
	o.State = v
}

// GetReady returns the Ready field value
func (o *MariaDBResponse) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *MariaDBResponse) SetReady(v bool) {
	o.Ready = v
}

// GetName returns the Name field value
func (o *MariaDBResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MariaDBResponse) SetName(v string) {
	o.Name = v
}

// GetServer returns the Server field value
func (o *MariaDBResponse) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *MariaDBResponse) SetServer(v string) {
	o.Server = v
}

// GetCharset returns the Charset field value
func (o *MariaDBResponse) GetCharset() MariaCharset {
	if o == nil {
		var ret MariaCharset
		return ret
	}

	return o.Charset
}

// GetCharsetOk returns a tuple with the Charset field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetCharsetOk() (*MariaCharset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Charset, true
}

// SetCharset sets field value
func (o *MariaDBResponse) SetCharset(v MariaCharset) {
	o.Charset = v
}

// GetDbusersReadwrite returns the DbusersReadwrite field value
func (o *MariaDBResponse) GetDbusersReadwrite() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DbusersReadwrite
}

// GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetDbusersReadwriteOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbusersReadwrite, true
}

// SetDbusersReadwrite sets field value
func (o *MariaDBResponse) SetDbusersReadwrite(v []string) {
	o.DbusersReadwrite = v
}

// GetDbusersReadonly returns the DbusersReadonly field value
func (o *MariaDBResponse) GetDbusersReadonly() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DbusersReadonly
}

// GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field value
// and a boolean to check if the value has been set.
func (o *MariaDBResponse) GetDbusersReadonlyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbusersReadonly, true
}

// SetDbusersReadonly sets field value
func (o *MariaDBResponse) SetDbusersReadonly(v []string) {
	o.DbusersReadonly = v
}

func (o MariaDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MariaDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["ready"] = o.Ready
	toSerialize["name"] = o.Name
	toSerialize["server"] = o.Server
	toSerialize["charset"] = o.Charset
	toSerialize["dbusers_readwrite"] = o.DbusersReadwrite
	toSerialize["dbusers_readonly"] = o.DbusersReadonly
	return toSerialize, nil
}

type NullableMariaDBResponse struct {
	value *MariaDBResponse
	isSet bool
}

func (v NullableMariaDBResponse) Get() *MariaDBResponse {
	return v.value
}

func (v *NullableMariaDBResponse) Set(val *MariaDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMariaDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMariaDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMariaDBResponse(val *MariaDBResponse) *NullableMariaDBResponse {
	return &NullableMariaDBResponse{value: val, isSet: true}
}

func (v NullableMariaDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMariaDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

