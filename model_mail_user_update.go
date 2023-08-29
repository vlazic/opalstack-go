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

// checks if the MailUserUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailUserUpdate{}

// MailUserUpdate struct for MailUserUpdate
type MailUserUpdate struct {
	Id string `json:"id"`
	Password *string `json:"password,omitempty"`
	Procmailrc *string `json:"procmailrc,omitempty"`
	AutoresponderEnable *bool `json:"autoresponder_enable,omitempty"`
	AutoresponderSubject *string `json:"autoresponder_subject,omitempty"`
	AutoresponderMessage *string `json:"autoresponder_message,omitempty"`
	AutoresponderNoreply []string `json:"autoresponder_noreply,omitempty"`
}

// NewMailUserUpdate instantiates a new MailUserUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailUserUpdate(id string) *MailUserUpdate {
	this := MailUserUpdate{}
	this.Id = id
	return &this
}

// NewMailUserUpdateWithDefaults instantiates a new MailUserUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailUserUpdateWithDefaults() *MailUserUpdate {
	this := MailUserUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *MailUserUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MailUserUpdate) SetId(v string) {
	o.Id = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MailUserUpdate) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MailUserUpdate) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MailUserUpdate) SetPassword(v string) {
	o.Password = &v
}

// GetProcmailrc returns the Procmailrc field value if set, zero value otherwise.
func (o *MailUserUpdate) GetProcmailrc() string {
	if o == nil || IsNil(o.Procmailrc) {
		var ret string
		return ret
	}
	return *o.Procmailrc
}

// GetProcmailrcOk returns a tuple with the Procmailrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetProcmailrcOk() (*string, bool) {
	if o == nil || IsNil(o.Procmailrc) {
		return nil, false
	}
	return o.Procmailrc, true
}

// HasProcmailrc returns a boolean if a field has been set.
func (o *MailUserUpdate) HasProcmailrc() bool {
	if o != nil && !IsNil(o.Procmailrc) {
		return true
	}

	return false
}

// SetProcmailrc gets a reference to the given string and assigns it to the Procmailrc field.
func (o *MailUserUpdate) SetProcmailrc(v string) {
	o.Procmailrc = &v
}

// GetAutoresponderEnable returns the AutoresponderEnable field value if set, zero value otherwise.
func (o *MailUserUpdate) GetAutoresponderEnable() bool {
	if o == nil || IsNil(o.AutoresponderEnable) {
		var ret bool
		return ret
	}
	return *o.AutoresponderEnable
}

// GetAutoresponderEnableOk returns a tuple with the AutoresponderEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetAutoresponderEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoresponderEnable) {
		return nil, false
	}
	return o.AutoresponderEnable, true
}

// HasAutoresponderEnable returns a boolean if a field has been set.
func (o *MailUserUpdate) HasAutoresponderEnable() bool {
	if o != nil && !IsNil(o.AutoresponderEnable) {
		return true
	}

	return false
}

// SetAutoresponderEnable gets a reference to the given bool and assigns it to the AutoresponderEnable field.
func (o *MailUserUpdate) SetAutoresponderEnable(v bool) {
	o.AutoresponderEnable = &v
}

// GetAutoresponderSubject returns the AutoresponderSubject field value if set, zero value otherwise.
func (o *MailUserUpdate) GetAutoresponderSubject() string {
	if o == nil || IsNil(o.AutoresponderSubject) {
		var ret string
		return ret
	}
	return *o.AutoresponderSubject
}

// GetAutoresponderSubjectOk returns a tuple with the AutoresponderSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetAutoresponderSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.AutoresponderSubject) {
		return nil, false
	}
	return o.AutoresponderSubject, true
}

// HasAutoresponderSubject returns a boolean if a field has been set.
func (o *MailUserUpdate) HasAutoresponderSubject() bool {
	if o != nil && !IsNil(o.AutoresponderSubject) {
		return true
	}

	return false
}

// SetAutoresponderSubject gets a reference to the given string and assigns it to the AutoresponderSubject field.
func (o *MailUserUpdate) SetAutoresponderSubject(v string) {
	o.AutoresponderSubject = &v
}

// GetAutoresponderMessage returns the AutoresponderMessage field value if set, zero value otherwise.
func (o *MailUserUpdate) GetAutoresponderMessage() string {
	if o == nil || IsNil(o.AutoresponderMessage) {
		var ret string
		return ret
	}
	return *o.AutoresponderMessage
}

// GetAutoresponderMessageOk returns a tuple with the AutoresponderMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetAutoresponderMessageOk() (*string, bool) {
	if o == nil || IsNil(o.AutoresponderMessage) {
		return nil, false
	}
	return o.AutoresponderMessage, true
}

// HasAutoresponderMessage returns a boolean if a field has been set.
func (o *MailUserUpdate) HasAutoresponderMessage() bool {
	if o != nil && !IsNil(o.AutoresponderMessage) {
		return true
	}

	return false
}

// SetAutoresponderMessage gets a reference to the given string and assigns it to the AutoresponderMessage field.
func (o *MailUserUpdate) SetAutoresponderMessage(v string) {
	o.AutoresponderMessage = &v
}

// GetAutoresponderNoreply returns the AutoresponderNoreply field value if set, zero value otherwise.
func (o *MailUserUpdate) GetAutoresponderNoreply() []string {
	if o == nil || IsNil(o.AutoresponderNoreply) {
		var ret []string
		return ret
	}
	return o.AutoresponderNoreply
}

// GetAutoresponderNoreplyOk returns a tuple with the AutoresponderNoreply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailUserUpdate) GetAutoresponderNoreplyOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoresponderNoreply) {
		return nil, false
	}
	return o.AutoresponderNoreply, true
}

// HasAutoresponderNoreply returns a boolean if a field has been set.
func (o *MailUserUpdate) HasAutoresponderNoreply() bool {
	if o != nil && !IsNil(o.AutoresponderNoreply) {
		return true
	}

	return false
}

// SetAutoresponderNoreply gets a reference to the given []string and assigns it to the AutoresponderNoreply field.
func (o *MailUserUpdate) SetAutoresponderNoreply(v []string) {
	o.AutoresponderNoreply = v
}

func (o MailUserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailUserUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Procmailrc) {
		toSerialize["procmailrc"] = o.Procmailrc
	}
	if !IsNil(o.AutoresponderEnable) {
		toSerialize["autoresponder_enable"] = o.AutoresponderEnable
	}
	if !IsNil(o.AutoresponderSubject) {
		toSerialize["autoresponder_subject"] = o.AutoresponderSubject
	}
	if !IsNil(o.AutoresponderMessage) {
		toSerialize["autoresponder_message"] = o.AutoresponderMessage
	}
	if !IsNil(o.AutoresponderNoreply) {
		toSerialize["autoresponder_noreply"] = o.AutoresponderNoreply
	}
	return toSerialize, nil
}

type NullableMailUserUpdate struct {
	value *MailUserUpdate
	isSet bool
}

func (v NullableMailUserUpdate) Get() *MailUserUpdate {
	return v.value
}

func (v *NullableMailUserUpdate) Set(val *MailUserUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMailUserUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMailUserUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailUserUpdate(val *MailUserUpdate) *NullableMailUserUpdate {
	return &NullableMailUserUpdate{value: val, isSet: true}
}

func (v NullableMailUserUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailUserUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

