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

// MariaCharset the model 'MariaCharset'
type MariaCharset string

// List of MariaCharset
const (
	MARIACHARSET_UTF8 MariaCharset = "utf8"
	MARIACHARSET_ARMSCII8 MariaCharset = "armscii8"
	MARIACHARSET_ASCII MariaCharset = "ascii"
	MARIACHARSET_BIG5 MariaCharset = "big5"
	MARIACHARSET_BINARY MariaCharset = "binary"
	MARIACHARSET_CP1250 MariaCharset = "cp1250"
	MARIACHARSET_CP1251 MariaCharset = "cp1251"
	MARIACHARSET_CP1256 MariaCharset = "cp1256"
	MARIACHARSET_CP1257 MariaCharset = "cp1257"
	MARIACHARSET_CP850 MariaCharset = "cp850"
	MARIACHARSET_CP852 MariaCharset = "cp852"
	MARIACHARSET_CP866 MariaCharset = "cp866"
	MARIACHARSET_CP932 MariaCharset = "cp932"
	MARIACHARSET_DEC8 MariaCharset = "dec8"
	MARIACHARSET_EUCJPMS MariaCharset = "eucjpms"
	MARIACHARSET_EUCKR MariaCharset = "euckr"
	MARIACHARSET_GB2312 MariaCharset = "gb2312"
	MARIACHARSET_GBK MariaCharset = "gbk"
	MARIACHARSET_GEOSTD8 MariaCharset = "geostd8"
	MARIACHARSET_GREEK MariaCharset = "greek"
	MARIACHARSET_HEBREW MariaCharset = "hebrew"
	MARIACHARSET_HP8 MariaCharset = "hp8"
	MARIACHARSET_KEYBCS2 MariaCharset = "keybcs2"
	MARIACHARSET_KOI8R MariaCharset = "koi8r"
	MARIACHARSET_KOI8U MariaCharset = "koi8u"
	MARIACHARSET_LATIN1 MariaCharset = "latin1"
	MARIACHARSET_LATIN2 MariaCharset = "latin2"
	MARIACHARSET_LATIN5 MariaCharset = "latin5"
	MARIACHARSET_LATIN7 MariaCharset = "latin7"
	MARIACHARSET_MACCE MariaCharset = "macce"
	MARIACHARSET_MACROMAN MariaCharset = "macroman"
	MARIACHARSET_SJIS MariaCharset = "sjis"
	MARIACHARSET_SWE7 MariaCharset = "swe7"
	MARIACHARSET_TIS620 MariaCharset = "tis620"
	MARIACHARSET_UCS2 MariaCharset = "ucs2"
	MARIACHARSET_UJIS MariaCharset = "ujis"
	MARIACHARSET_UTF16 MariaCharset = "utf16"
	MARIACHARSET_UTF16LE MariaCharset = "utf16le"
	MARIACHARSET_UTF32 MariaCharset = "utf32"
	MARIACHARSET_UTF8MB4 MariaCharset = "utf8mb4"
)

// All allowed values of MariaCharset enum
var AllowedMariaCharsetEnumValues = []MariaCharset{
	"utf8",
	"armscii8",
	"ascii",
	"big5",
	"binary",
	"cp1250",
	"cp1251",
	"cp1256",
	"cp1257",
	"cp850",
	"cp852",
	"cp866",
	"cp932",
	"dec8",
	"eucjpms",
	"euckr",
	"gb2312",
	"gbk",
	"geostd8",
	"greek",
	"hebrew",
	"hp8",
	"keybcs2",
	"koi8r",
	"koi8u",
	"latin1",
	"latin2",
	"latin5",
	"latin7",
	"macce",
	"macroman",
	"sjis",
	"swe7",
	"tis620",
	"ucs2",
	"ujis",
	"utf16",
	"utf16le",
	"utf32",
	"utf8mb4",
}

func (v *MariaCharset) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MariaCharset(value)
	for _, existing := range AllowedMariaCharsetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MariaCharset", value)
}

// NewMariaCharsetFromValue returns a pointer to a valid MariaCharset
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMariaCharsetFromValue(v string) (*MariaCharset, error) {
	ev := MariaCharset(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MariaCharset: valid values are %v", v, AllowedMariaCharsetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MariaCharset) IsValid() bool {
	for _, existing := range AllowedMariaCharsetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MariaCharset value
func (v MariaCharset) Ptr() *MariaCharset {
	return &v
}

type NullableMariaCharset struct {
	value *MariaCharset
	isSet bool
}

func (v NullableMariaCharset) Get() *MariaCharset {
	return v.value
}

func (v *NullableMariaCharset) Set(val *MariaCharset) {
	v.value = val
	v.isSet = true
}

func (v NullableMariaCharset) IsSet() bool {
	return v.isSet
}

func (v *NullableMariaCharset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMariaCharset(val *MariaCharset) *NullableMariaCharset {
	return &NullableMariaCharset{value: val, isSet: true}
}

func (v NullableMariaCharset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMariaCharset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
