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

// checks if the MailUsageMailUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailUsageMailUser{}

// MailUsageMailUser struct for MailUsageMailUser
type MailUsageMailUser struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImapServer MailUsageImapServer `json:"imap_server"`
}

// NewMailUsageMailUser instantiates a new MailUsageMailUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailUsageMailUser(id string, name string, imapServer MailUsageImapServer) *MailUsageMailUser {
	this := MailUsageMailUser{}
	this.Id = id
	this.Name = name
	this.ImapServer = imapServer
	return &this
}

// NewMailUsageMailUserWithDefaults instantiates a new MailUsageMailUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailUsageMailUserWithDefaults() *MailUsageMailUser {
	this := MailUsageMailUser{}
	return &this
}

// GetId returns the Id field value
func (o *MailUsageMailUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MailUsageMailUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MailUsageMailUser) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MailUsageMailUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MailUsageMailUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MailUsageMailUser) SetName(v string) {
	o.Name = v
}

// GetImapServer returns the ImapServer field value
func (o *MailUsageMailUser) GetImapServer() MailUsageImapServer {
	if o == nil {
		var ret MailUsageImapServer
		return ret
	}

	return o.ImapServer
}

// GetImapServerOk returns a tuple with the ImapServer field value
// and a boolean to check if the value has been set.
func (o *MailUsageMailUser) GetImapServerOk() (*MailUsageImapServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImapServer, true
}

// SetImapServer sets field value
func (o *MailUsageMailUser) SetImapServer(v MailUsageImapServer) {
	o.ImapServer = v
}

func (o MailUsageMailUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailUsageMailUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["imap_server"] = o.ImapServer
	return toSerialize, nil
}

type NullableMailUsageMailUser struct {
	value *MailUsageMailUser
	isSet bool
}

func (v NullableMailUsageMailUser) Get() *MailUsageMailUser {
	return v.value
}

func (v *NullableMailUsageMailUser) Set(val *MailUsageMailUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMailUsageMailUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMailUsageMailUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailUsageMailUser(val *MailUsageMailUser) *NullableMailUsageMailUser {
	return &NullableMailUsageMailUser{value: val, isSet: true}
}

func (v NullableMailUsageMailUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailUsageMailUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


