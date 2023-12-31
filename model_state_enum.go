/*
Opalstack API

 ## The Opalstack JSON REST API  ### Authorization This API uses an **Authorization** header of the form: `\"Authorization: Token 1111111111111111111111111111111111111111\"`, where **1111111111111111111111111111111111111111** represents an API token created at https://my.opalstack.com/tokens/.  The typical format of an API request looks like the following: ``` GET request:     curl -s -H \"Authorization: Token 1111111111111111111111111111111111111111\" \"https://my.opalstack.com/api/v1/site/list/\" | jq .  POST request:     curl -s -H \"Content-Type: application/json\" -H \"Authorization: Token 1111111111111111111111111111111111111111\" \\             -X POST -d '[{\"id\": \"(site UUID)\", \"redirect\": true, ...}]' \"https://my.opalstack.com/api/v1/site/update/\" | jq . ``` (Further examples will omit **headers** and **jq** for the sake of clarity)  You can also authorize requests on our API Documentation page (https://my.opalstack.com/api/v1/doc/) in order to facilitate development. To do so, click the \"**Authorize**\" button on the right side of the page and enter \"**Token 1111111111111111111111111111111111111111**\" in the **Value** field within. Afterword, you will be able to perform requests directly from the documentation page. Be sure to logout when finished.  ### Embedding The Opalstack API supports _embedding_. This allows you to nest child API objects in a single GET request. For example, consider the following GET request performed with **curl**: ``` Request:     curl \"https://my.opalstack.com/api/v1/osuser/list/\"  Response:     [       {         \"id\": \"01010101-0202-0303-0404-050505050505\",         \"state\": \"READY\",         \"ready\": true,         \"name\": \"the_osuser_name\",         \"server\": \"11111111-1212-1313-1414-151515151515\"       }     ] ```  Suppose then that we would like additional information about the **server**. We _could_ proceed to query the **server** UUID (**11111111-1212-1313-1414-151515151515**) at the `/server/read/{uuid}` endpoint, like this: ``` Request:     curl \"https://my.opalstack.com/api/v1/server/read/11111111-1212-1313-1414-151515151515\"  Response:     {       \"id\": \"11111111-1212-1313-1414-151515151515\",       \"hostname\": \"vpsNNN.opalstack.com\"     } ```  However, we could have instead choosen to specify `?embed=server` as a query parameter to the original GET request. This will cause objects to be _embedded_ in the response directly: ``` Request:     curl \"https://my.opalstack.com/api/v1/osuser/list/?embed=server\"  Response:     [       {         \"id\": \"01010101-0202-0303-0404-050505050505\",         \"state\": \"READY\",         \"ready\": true,         \"name\": \"the_osuser_name\",         \"server\": {           \"id\": \"11111111-1212-1313-1414-151515151515\",           \"hostname\": \"vpsNNN.opalstack.com\"         }       }     ] ``` Here, the **server** field has been be populated with the full object instead of just a UUID.  The **embed** query parameter accepts multiple (comma-separated) fields to embed. For example: ``` curl \"https://my.opalstack.com/api/v1/account/info/?embed=web_servers,imap_servers,smtp_servers\" ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OpalStack

import (
	"encoding/json"
	"fmt"
)

// StateEnum the model 'StateEnum'
type StateEnum string

// List of StateEnum
const (
	STATEENUM_READY StateEnum = "READY"
	STATEENUM_PENDING_CREATE StateEnum = "PENDING_CREATE"
	STATEENUM_PENDING_UPDATE StateEnum = "PENDING_UPDATE"
	STATEENUM_PENDING_DELETE StateEnum = "PENDING_DELETE"
)

// All allowed values of StateEnum enum
var AllowedStateEnumEnumValues = []StateEnum{
	"READY",
	"PENDING_CREATE",
	"PENDING_UPDATE",
	"PENDING_DELETE",
}

func (v *StateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StateEnum(value)
	for _, existing := range AllowedStateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StateEnum", value)
}

// NewStateEnumFromValue returns a pointer to a valid StateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStateEnumFromValue(v string) (*StateEnum, error) {
	ev := StateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StateEnum: valid values are %v", v, AllowedStateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StateEnum) IsValid() bool {
	for _, existing := range AllowedStateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateEnum value
func (v StateEnum) Ptr() *StateEnum {
	return &v
}

type NullableStateEnum struct {
	value *StateEnum
	isSet bool
}

func (v NullableStateEnum) Get() *StateEnum {
	return v.value
}

func (v *NullableStateEnum) Set(val *StateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateEnum(val *StateEnum) *NullableStateEnum {
	return &NullableStateEnum{value: val, isSet: true}
}

func (v NullableStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

