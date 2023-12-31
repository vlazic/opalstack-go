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

// checks if the PsqlUserCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PsqlUserCreate{}

// PsqlUserCreate struct for PsqlUserCreate
type PsqlUserCreate struct {
	Server string `json:"server"`
	Name string `json:"name"`
	Password *string `json:"password,omitempty"`
	External *bool `json:"external,omitempty"`
}

// NewPsqlUserCreate instantiates a new PsqlUserCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPsqlUserCreate(server string, name string) *PsqlUserCreate {
	this := PsqlUserCreate{}
	this.Server = server
	this.Name = name
	return &this
}

// NewPsqlUserCreateWithDefaults instantiates a new PsqlUserCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPsqlUserCreateWithDefaults() *PsqlUserCreate {
	this := PsqlUserCreate{}
	return &this
}

// GetServer returns the Server field value
func (o *PsqlUserCreate) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *PsqlUserCreate) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *PsqlUserCreate) SetServer(v string) {
	o.Server = v
}

// GetName returns the Name field value
func (o *PsqlUserCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PsqlUserCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PsqlUserCreate) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PsqlUserCreate) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsqlUserCreate) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PsqlUserCreate) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PsqlUserCreate) SetPassword(v string) {
	o.Password = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *PsqlUserCreate) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PsqlUserCreate) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *PsqlUserCreate) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *PsqlUserCreate) SetExternal(v bool) {
	o.External = &v
}

func (o PsqlUserCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PsqlUserCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["server"] = o.Server
	toSerialize["name"] = o.Name
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	return toSerialize, nil
}

type NullablePsqlUserCreate struct {
	value *PsqlUserCreate
	isSet bool
}

func (v NullablePsqlUserCreate) Get() *PsqlUserCreate {
	return v.value
}

func (v *NullablePsqlUserCreate) Set(val *PsqlUserCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePsqlUserCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePsqlUserCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsqlUserCreate(val *PsqlUserCreate) *NullablePsqlUserCreate {
	return &NullablePsqlUserCreate{value: val, isSet: true}
}

func (v NullablePsqlUserCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsqlUserCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


