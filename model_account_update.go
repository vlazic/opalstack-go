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

// checks if the AccountUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUpdate{}

// AccountUpdate struct for AccountUpdate
type AccountUpdate struct {
	Id string `json:"id"`
	PaymentProcessor *AccountUpdatePaymentProcessorEnum `json:"payment_processor,omitempty"`
}

// NewAccountUpdate instantiates a new AccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdate(id string) *AccountUpdate {
	this := AccountUpdate{}
	this.Id = id
	return &this
}

// NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateWithDefaults() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *AccountUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountUpdate) SetId(v string) {
	o.Id = v
}

// GetPaymentProcessor returns the PaymentProcessor field value if set, zero value otherwise.
func (o *AccountUpdate) GetPaymentProcessor() AccountUpdatePaymentProcessorEnum {
	if o == nil || IsNil(o.PaymentProcessor) {
		var ret AccountUpdatePaymentProcessorEnum
		return ret
	}
	return *o.PaymentProcessor
}

// GetPaymentProcessorOk returns a tuple with the PaymentProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetPaymentProcessorOk() (*AccountUpdatePaymentProcessorEnum, bool) {
	if o == nil || IsNil(o.PaymentProcessor) {
		return nil, false
	}
	return o.PaymentProcessor, true
}

// HasPaymentProcessor returns a boolean if a field has been set.
func (o *AccountUpdate) HasPaymentProcessor() bool {
	if o != nil && !IsNil(o.PaymentProcessor) {
		return true
	}

	return false
}

// SetPaymentProcessor gets a reference to the given AccountUpdatePaymentProcessorEnum and assigns it to the PaymentProcessor field.
func (o *AccountUpdate) SetPaymentProcessor(v AccountUpdatePaymentProcessorEnum) {
	o.PaymentProcessor = &v
}

func (o AccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.PaymentProcessor) {
		toSerialize["payment_processor"] = o.PaymentProcessor
	}
	return toSerialize, nil
}

type NullableAccountUpdate struct {
	value *AccountUpdate
	isSet bool
}

func (v NullableAccountUpdate) Get() *AccountUpdate {
	return v.value
}

func (v *NullableAccountUpdate) Set(val *AccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdate(val *AccountUpdate) *NullableAccountUpdate {
	return &NullableAccountUpdate{value: val, isSet: true}
}

func (v NullableAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

