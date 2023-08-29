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

// checks if the OSVarCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSVarCreate{}

// OSVarCreate struct for OSVarCreate
type OSVarCreate struct {
	Name string `json:"name"`
	Content *string `json:"content,omitempty"`
	Osusers []string `json:"osusers,omitempty"`
	Global *bool `json:"global,omitempty"`
}

// NewOSVarCreate instantiates a new OSVarCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSVarCreate(name string) *OSVarCreate {
	this := OSVarCreate{}
	this.Name = name
	return &this
}

// NewOSVarCreateWithDefaults instantiates a new OSVarCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSVarCreateWithDefaults() *OSVarCreate {
	this := OSVarCreate{}
	return &this
}

// GetName returns the Name field value
func (o *OSVarCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OSVarCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OSVarCreate) SetName(v string) {
	o.Name = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OSVarCreate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVarCreate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OSVarCreate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *OSVarCreate) SetContent(v string) {
	o.Content = &v
}

// GetOsusers returns the Osusers field value if set, zero value otherwise.
func (o *OSVarCreate) GetOsusers() []string {
	if o == nil || IsNil(o.Osusers) {
		var ret []string
		return ret
	}
	return o.Osusers
}

// GetOsusersOk returns a tuple with the Osusers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVarCreate) GetOsusersOk() ([]string, bool) {
	if o == nil || IsNil(o.Osusers) {
		return nil, false
	}
	return o.Osusers, true
}

// HasOsusers returns a boolean if a field has been set.
func (o *OSVarCreate) HasOsusers() bool {
	if o != nil && !IsNil(o.Osusers) {
		return true
	}

	return false
}

// SetOsusers gets a reference to the given []string and assigns it to the Osusers field.
func (o *OSVarCreate) SetOsusers(v []string) {
	o.Osusers = v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *OSVarCreate) GetGlobal() bool {
	if o == nil || IsNil(o.Global) {
		var ret bool
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVarCreate) GetGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.Global) {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *OSVarCreate) HasGlobal() bool {
	if o != nil && !IsNil(o.Global) {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given bool and assigns it to the Global field.
func (o *OSVarCreate) SetGlobal(v bool) {
	o.Global = &v
}

func (o OSVarCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSVarCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Osusers) {
		toSerialize["osusers"] = o.Osusers
	}
	if !IsNil(o.Global) {
		toSerialize["global"] = o.Global
	}
	return toSerialize, nil
}

type NullableOSVarCreate struct {
	value *OSVarCreate
	isSet bool
}

func (v NullableOSVarCreate) Get() *OSVarCreate {
	return v.value
}

func (v *NullableOSVarCreate) Set(val *OSVarCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOSVarCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOSVarCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSVarCreate(val *OSVarCreate) *NullableOSVarCreate {
	return &NullableOSVarCreate{value: val, isSet: true}
}

func (v NullableOSVarCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSVarCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


