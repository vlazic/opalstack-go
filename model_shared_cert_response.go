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

// checks if the SharedCertResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedCertResponse{}

// SharedCertResponse struct for SharedCertResponse
type SharedCertResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Cert string `json:"cert"`
	Intermediates string `json:"intermediates"`
	ExpDate time.Time `json:"exp_date"`
	ListedDomains []string `json:"listed_domains"`
	IsOpalstackLetsencrypt *bool `json:"is_opalstack_letsencrypt,omitempty"`
	IsOpalstackSharedCert *bool `json:"is_opalstack_shared_cert,omitempty"`
}

// NewSharedCertResponse instantiates a new SharedCertResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedCertResponse(id string, name string, cert string, intermediates string, expDate time.Time, listedDomains []string) *SharedCertResponse {
	this := SharedCertResponse{}
	this.Id = id
	this.Name = name
	this.Cert = cert
	this.Intermediates = intermediates
	this.ExpDate = expDate
	this.ListedDomains = listedDomains
	var isOpalstackLetsencrypt bool = true
	this.IsOpalstackLetsencrypt = &isOpalstackLetsencrypt
	var isOpalstackSharedCert bool = true
	this.IsOpalstackSharedCert = &isOpalstackSharedCert
	return &this
}

// NewSharedCertResponseWithDefaults instantiates a new SharedCertResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedCertResponseWithDefaults() *SharedCertResponse {
	this := SharedCertResponse{}
	var isOpalstackLetsencrypt bool = true
	this.IsOpalstackLetsencrypt = &isOpalstackLetsencrypt
	var isOpalstackSharedCert bool = true
	this.IsOpalstackSharedCert = &isOpalstackSharedCert
	return &this
}

// GetId returns the Id field value
func (o *SharedCertResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SharedCertResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SharedCertResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SharedCertResponse) SetName(v string) {
	o.Name = v
}

// GetCert returns the Cert field value
func (o *SharedCertResponse) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *SharedCertResponse) SetCert(v string) {
	o.Cert = v
}

// GetIntermediates returns the Intermediates field value
func (o *SharedCertResponse) GetIntermediates() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Intermediates
}

// GetIntermediatesOk returns a tuple with the Intermediates field value
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetIntermediatesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intermediates, true
}

// SetIntermediates sets field value
func (o *SharedCertResponse) SetIntermediates(v string) {
	o.Intermediates = v
}

// GetExpDate returns the ExpDate field value
func (o *SharedCertResponse) GetExpDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpDate
}

// GetExpDateOk returns a tuple with the ExpDate field value
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetExpDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpDate, true
}

// SetExpDate sets field value
func (o *SharedCertResponse) SetExpDate(v time.Time) {
	o.ExpDate = v
}

// GetListedDomains returns the ListedDomains field value
func (o *SharedCertResponse) GetListedDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ListedDomains
}

// GetListedDomainsOk returns a tuple with the ListedDomains field value
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetListedDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ListedDomains, true
}

// SetListedDomains sets field value
func (o *SharedCertResponse) SetListedDomains(v []string) {
	o.ListedDomains = v
}

// GetIsOpalstackLetsencrypt returns the IsOpalstackLetsencrypt field value if set, zero value otherwise.
func (o *SharedCertResponse) GetIsOpalstackLetsencrypt() bool {
	if o == nil || IsNil(o.IsOpalstackLetsencrypt) {
		var ret bool
		return ret
	}
	return *o.IsOpalstackLetsencrypt
}

// GetIsOpalstackLetsencryptOk returns a tuple with the IsOpalstackLetsencrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetIsOpalstackLetsencryptOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpalstackLetsencrypt) {
		return nil, false
	}
	return o.IsOpalstackLetsencrypt, true
}

// HasIsOpalstackLetsencrypt returns a boolean if a field has been set.
func (o *SharedCertResponse) HasIsOpalstackLetsencrypt() bool {
	if o != nil && !IsNil(o.IsOpalstackLetsencrypt) {
		return true
	}

	return false
}

// SetIsOpalstackLetsencrypt gets a reference to the given bool and assigns it to the IsOpalstackLetsencrypt field.
func (o *SharedCertResponse) SetIsOpalstackLetsencrypt(v bool) {
	o.IsOpalstackLetsencrypt = &v
}

// GetIsOpalstackSharedCert returns the IsOpalstackSharedCert field value if set, zero value otherwise.
func (o *SharedCertResponse) GetIsOpalstackSharedCert() bool {
	if o == nil || IsNil(o.IsOpalstackSharedCert) {
		var ret bool
		return ret
	}
	return *o.IsOpalstackSharedCert
}

// GetIsOpalstackSharedCertOk returns a tuple with the IsOpalstackSharedCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedCertResponse) GetIsOpalstackSharedCertOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpalstackSharedCert) {
		return nil, false
	}
	return o.IsOpalstackSharedCert, true
}

// HasIsOpalstackSharedCert returns a boolean if a field has been set.
func (o *SharedCertResponse) HasIsOpalstackSharedCert() bool {
	if o != nil && !IsNil(o.IsOpalstackSharedCert) {
		return true
	}

	return false
}

// SetIsOpalstackSharedCert gets a reference to the given bool and assigns it to the IsOpalstackSharedCert field.
func (o *SharedCertResponse) SetIsOpalstackSharedCert(v bool) {
	o.IsOpalstackSharedCert = &v
}

func (o SharedCertResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedCertResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["cert"] = o.Cert
	toSerialize["intermediates"] = o.Intermediates
	toSerialize["exp_date"] = o.ExpDate
	toSerialize["listed_domains"] = o.ListedDomains
	if !IsNil(o.IsOpalstackLetsencrypt) {
		toSerialize["is_opalstack_letsencrypt"] = o.IsOpalstackLetsencrypt
	}
	if !IsNil(o.IsOpalstackSharedCert) {
		toSerialize["is_opalstack_shared_cert"] = o.IsOpalstackSharedCert
	}
	return toSerialize, nil
}

type NullableSharedCertResponse struct {
	value *SharedCertResponse
	isSet bool
}

func (v NullableSharedCertResponse) Get() *SharedCertResponse {
	return v.value
}

func (v *NullableSharedCertResponse) Set(val *SharedCertResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedCertResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedCertResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedCertResponse(val *SharedCertResponse) *NullableSharedCertResponse {
	return &NullableSharedCertResponse{value: val, isSet: true}
}

func (v NullableSharedCertResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedCertResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


