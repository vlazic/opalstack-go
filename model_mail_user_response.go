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

// checks if the MailUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailUserResponse{}

// MailUserResponse struct for MailUserResponse
type MailUserResponse struct {
	Id string `json:"id"`
	State StateEnum `json:"state"`
	Ready bool `json:"ready"`
	Name string `json:"name"`
	ImapServer string `json:"imap_server"`
	Procmailrc string `json:"procmailrc"`
	AutoresponderEnable string `json:"autoresponder_enable"`
	AutoresponderSubject string `json:"autoresponder_subject"`
	AutoresponderMessage string `json:"autoresponder_message"`
	AutoresponderNoreply []string `json:"autoresponder_noreply"`
}

// NewMailUserResponse instantiates a new MailUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailUserResponse(id string, state StateEnum, ready bool, name string, imapServer string, procmailrc string, autoresponderEnable string, autoresponderSubject string, autoresponderMessage string, autoresponderNoreply []string) *MailUserResponse {
	this := MailUserResponse{}
	this.Id = id
	this.State = state
	this.Ready = ready
	this.Name = name
	this.ImapServer = imapServer
	this.Procmailrc = procmailrc
	this.AutoresponderEnable = autoresponderEnable
	this.AutoresponderSubject = autoresponderSubject
	this.AutoresponderMessage = autoresponderMessage
	this.AutoresponderNoreply = autoresponderNoreply
	return &this
}

// NewMailUserResponseWithDefaults instantiates a new MailUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailUserResponseWithDefaults() *MailUserResponse {
	this := MailUserResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MailUserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MailUserResponse) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *MailUserResponse) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *MailUserResponse) SetState(v StateEnum) {
	o.State = v
}

// GetReady returns the Ready field value
func (o *MailUserResponse) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *MailUserResponse) SetReady(v bool) {
	o.Ready = v
}

// GetName returns the Name field value
func (o *MailUserResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MailUserResponse) SetName(v string) {
	o.Name = v
}

// GetImapServer returns the ImapServer field value
func (o *MailUserResponse) GetImapServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImapServer
}

// GetImapServerOk returns a tuple with the ImapServer field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetImapServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImapServer, true
}

// SetImapServer sets field value
func (o *MailUserResponse) SetImapServer(v string) {
	o.ImapServer = v
}

// GetProcmailrc returns the Procmailrc field value
func (o *MailUserResponse) GetProcmailrc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Procmailrc
}

// GetProcmailrcOk returns a tuple with the Procmailrc field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetProcmailrcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Procmailrc, true
}

// SetProcmailrc sets field value
func (o *MailUserResponse) SetProcmailrc(v string) {
	o.Procmailrc = v
}

// GetAutoresponderEnable returns the AutoresponderEnable field value
func (o *MailUserResponse) GetAutoresponderEnable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoresponderEnable
}

// GetAutoresponderEnableOk returns a tuple with the AutoresponderEnable field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetAutoresponderEnableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoresponderEnable, true
}

// SetAutoresponderEnable sets field value
func (o *MailUserResponse) SetAutoresponderEnable(v string) {
	o.AutoresponderEnable = v
}

// GetAutoresponderSubject returns the AutoresponderSubject field value
func (o *MailUserResponse) GetAutoresponderSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoresponderSubject
}

// GetAutoresponderSubjectOk returns a tuple with the AutoresponderSubject field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetAutoresponderSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoresponderSubject, true
}

// SetAutoresponderSubject sets field value
func (o *MailUserResponse) SetAutoresponderSubject(v string) {
	o.AutoresponderSubject = v
}

// GetAutoresponderMessage returns the AutoresponderMessage field value
func (o *MailUserResponse) GetAutoresponderMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoresponderMessage
}

// GetAutoresponderMessageOk returns a tuple with the AutoresponderMessage field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetAutoresponderMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoresponderMessage, true
}

// SetAutoresponderMessage sets field value
func (o *MailUserResponse) SetAutoresponderMessage(v string) {
	o.AutoresponderMessage = v
}

// GetAutoresponderNoreply returns the AutoresponderNoreply field value
func (o *MailUserResponse) GetAutoresponderNoreply() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AutoresponderNoreply
}

// GetAutoresponderNoreplyOk returns a tuple with the AutoresponderNoreply field value
// and a boolean to check if the value has been set.
func (o *MailUserResponse) GetAutoresponderNoreplyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoresponderNoreply, true
}

// SetAutoresponderNoreply sets field value
func (o *MailUserResponse) SetAutoresponderNoreply(v []string) {
	o.AutoresponderNoreply = v
}

func (o MailUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	toSerialize["ready"] = o.Ready
	toSerialize["name"] = o.Name
	toSerialize["imap_server"] = o.ImapServer
	toSerialize["procmailrc"] = o.Procmailrc
	toSerialize["autoresponder_enable"] = o.AutoresponderEnable
	toSerialize["autoresponder_subject"] = o.AutoresponderSubject
	toSerialize["autoresponder_message"] = o.AutoresponderMessage
	toSerialize["autoresponder_noreply"] = o.AutoresponderNoreply
	return toSerialize, nil
}

type NullableMailUserResponse struct {
	value *MailUserResponse
	isSet bool
}

func (v NullableMailUserResponse) Get() *MailUserResponse {
	return v.value
}

func (v *NullableMailUserResponse) Set(val *MailUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMailUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMailUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailUserResponse(val *MailUserResponse) *NullableMailUserResponse {
	return &NullableMailUserResponse{value: val, isSet: true}
}

func (v NullableMailUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


