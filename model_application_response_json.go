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

// checks if the ApplicationResponseJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResponseJson{}

// ApplicationResponseJson struct for ApplicationResponseJson
type ApplicationResponseJson struct {
	DbName string `json:"db_name"`
	DbUser string `json:"db_user"`
	DbHost string `json:"db_host"`
	DbPort *int32 `json:"db_port,omitempty"`
	AppName string `json:"app_name"`
	AppPort *int32 `json:"app_port,omitempty"`
	AppPath string `json:"app_path"`
	AppVersion string `json:"app_version"`
	AppCommand string `json:"app_command"`
	AppLangVersion string `json:"app_lang_version"`
	SymLinkPath *string `json:"sym_link_path,omitempty"`
	AutoSiteUrl bool `json:"auto_site_url"`
	AppExec bool `json:"app_exec"`
	UrlFopen bool `json:"url_fopen"`
	FpmType string `json:"fpm_type"`
	FpmMaxChildren *int32 `json:"fpm_max_children,omitempty"`
	FpmMaxRequests *int32 `json:"fpm_max_requests,omitempty"`
	FpmStartServers *int32 `json:"fpm_start_servers,omitempty"`
	FpmMinSpareServers *int32 `json:"fpm_min_spare_servers,omitempty"`
	FpmMaxSpareServers *int32 `json:"fpm_max_spare_servers,omitempty"`
	PhpVersion *int32 `json:"php_version,omitempty"`
	Subroot string `json:"subroot"`
}

// NewApplicationResponseJson instantiates a new ApplicationResponseJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResponseJson(dbName string, dbUser string, dbHost string, appName string, appPath string, appVersion string, appCommand string, appLangVersion string, autoSiteUrl bool, appExec bool, urlFopen bool, fpmType string, subroot string) *ApplicationResponseJson {
	this := ApplicationResponseJson{}
	this.DbName = dbName
	this.DbUser = dbUser
	this.DbHost = dbHost
	var dbPort int32 = 3306
	this.DbPort = &dbPort
	this.AppName = appName
	var appPort int32 = 8000
	this.AppPort = &appPort
	this.AppPath = appPath
	this.AppVersion = appVersion
	this.AppCommand = appCommand
	this.AppLangVersion = appLangVersion
	var symLinkPath string = "/home/my_osuser/path/to/dir/"
	this.SymLinkPath = &symLinkPath
	this.AutoSiteUrl = autoSiteUrl
	this.AppExec = appExec
	this.UrlFopen = urlFopen
	this.FpmType = fpmType
	var fpmMaxChildren int32 = 5
	this.FpmMaxChildren = &fpmMaxChildren
	var fpmMaxRequests int32 = 500
	this.FpmMaxRequests = &fpmMaxRequests
	var fpmStartServers int32 = 5
	this.FpmStartServers = &fpmStartServers
	var fpmMinSpareServers int32 = 5
	this.FpmMinSpareServers = &fpmMinSpareServers
	var fpmMaxSpareServers int32 = 5
	this.FpmMaxSpareServers = &fpmMaxSpareServers
	var phpVersion int32 = 74
	this.PhpVersion = &phpVersion
	this.Subroot = subroot
	return &this
}

// NewApplicationResponseJsonWithDefaults instantiates a new ApplicationResponseJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponseJsonWithDefaults() *ApplicationResponseJson {
	this := ApplicationResponseJson{}
	var dbPort int32 = 3306
	this.DbPort = &dbPort
	var appPort int32 = 8000
	this.AppPort = &appPort
	var symLinkPath string = "/home/my_osuser/path/to/dir/"
	this.SymLinkPath = &symLinkPath
	var fpmMaxChildren int32 = 5
	this.FpmMaxChildren = &fpmMaxChildren
	var fpmMaxRequests int32 = 500
	this.FpmMaxRequests = &fpmMaxRequests
	var fpmStartServers int32 = 5
	this.FpmStartServers = &fpmStartServers
	var fpmMinSpareServers int32 = 5
	this.FpmMinSpareServers = &fpmMinSpareServers
	var fpmMaxSpareServers int32 = 5
	this.FpmMaxSpareServers = &fpmMaxSpareServers
	var phpVersion int32 = 74
	this.PhpVersion = &phpVersion
	return &this
}

// GetDbName returns the DbName field value
func (o *ApplicationResponseJson) GetDbName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *ApplicationResponseJson) SetDbName(v string) {
	o.DbName = v
}

// GetDbUser returns the DbUser field value
func (o *ApplicationResponseJson) GetDbUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbUser
}

// GetDbUserOk returns a tuple with the DbUser field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetDbUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbUser, true
}

// SetDbUser sets field value
func (o *ApplicationResponseJson) SetDbUser(v string) {
	o.DbUser = v
}

// GetDbHost returns the DbHost field value
func (o *ApplicationResponseJson) GetDbHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbHost
}

// GetDbHostOk returns a tuple with the DbHost field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetDbHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbHost, true
}

// SetDbHost sets field value
func (o *ApplicationResponseJson) SetDbHost(v string) {
	o.DbHost = v
}

// GetDbPort returns the DbPort field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetDbPort() int32 {
	if o == nil || IsNil(o.DbPort) {
		var ret int32
		return ret
	}
	return *o.DbPort
}

// GetDbPortOk returns a tuple with the DbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetDbPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DbPort) {
		return nil, false
	}
	return o.DbPort, true
}

// HasDbPort returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasDbPort() bool {
	if o != nil && !IsNil(o.DbPort) {
		return true
	}

	return false
}

// SetDbPort gets a reference to the given int32 and assigns it to the DbPort field.
func (o *ApplicationResponseJson) SetDbPort(v int32) {
	o.DbPort = &v
}

// GetAppName returns the AppName field value
func (o *ApplicationResponseJson) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *ApplicationResponseJson) SetAppName(v string) {
	o.AppName = v
}

// GetAppPort returns the AppPort field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetAppPort() int32 {
	if o == nil || IsNil(o.AppPort) {
		var ret int32
		return ret
	}
	return *o.AppPort
}

// GetAppPortOk returns a tuple with the AppPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppPortOk() (*int32, bool) {
	if o == nil || IsNil(o.AppPort) {
		return nil, false
	}
	return o.AppPort, true
}

// HasAppPort returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasAppPort() bool {
	if o != nil && !IsNil(o.AppPort) {
		return true
	}

	return false
}

// SetAppPort gets a reference to the given int32 and assigns it to the AppPort field.
func (o *ApplicationResponseJson) SetAppPort(v int32) {
	o.AppPort = &v
}

// GetAppPath returns the AppPath field value
func (o *ApplicationResponseJson) GetAppPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppPath
}

// GetAppPathOk returns a tuple with the AppPath field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppPath, true
}

// SetAppPath sets field value
func (o *ApplicationResponseJson) SetAppPath(v string) {
	o.AppPath = v
}

// GetAppVersion returns the AppVersion field value
func (o *ApplicationResponseJson) GetAppVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppVersion, true
}

// SetAppVersion sets field value
func (o *ApplicationResponseJson) SetAppVersion(v string) {
	o.AppVersion = v
}

// GetAppCommand returns the AppCommand field value
func (o *ApplicationResponseJson) GetAppCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppCommand
}

// GetAppCommandOk returns a tuple with the AppCommand field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppCommand, true
}

// SetAppCommand sets field value
func (o *ApplicationResponseJson) SetAppCommand(v string) {
	o.AppCommand = v
}

// GetAppLangVersion returns the AppLangVersion field value
func (o *ApplicationResponseJson) GetAppLangVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLangVersion
}

// GetAppLangVersionOk returns a tuple with the AppLangVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppLangVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLangVersion, true
}

// SetAppLangVersion sets field value
func (o *ApplicationResponseJson) SetAppLangVersion(v string) {
	o.AppLangVersion = v
}

// GetSymLinkPath returns the SymLinkPath field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetSymLinkPath() string {
	if o == nil || IsNil(o.SymLinkPath) {
		var ret string
		return ret
	}
	return *o.SymLinkPath
}

// GetSymLinkPathOk returns a tuple with the SymLinkPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetSymLinkPathOk() (*string, bool) {
	if o == nil || IsNil(o.SymLinkPath) {
		return nil, false
	}
	return o.SymLinkPath, true
}

// HasSymLinkPath returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasSymLinkPath() bool {
	if o != nil && !IsNil(o.SymLinkPath) {
		return true
	}

	return false
}

// SetSymLinkPath gets a reference to the given string and assigns it to the SymLinkPath field.
func (o *ApplicationResponseJson) SetSymLinkPath(v string) {
	o.SymLinkPath = &v
}

// GetAutoSiteUrl returns the AutoSiteUrl field value
func (o *ApplicationResponseJson) GetAutoSiteUrl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoSiteUrl
}

// GetAutoSiteUrlOk returns a tuple with the AutoSiteUrl field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAutoSiteUrlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoSiteUrl, true
}

// SetAutoSiteUrl sets field value
func (o *ApplicationResponseJson) SetAutoSiteUrl(v bool) {
	o.AutoSiteUrl = v
}

// GetAppExec returns the AppExec field value
func (o *ApplicationResponseJson) GetAppExec() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AppExec
}

// GetAppExecOk returns a tuple with the AppExec field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetAppExecOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppExec, true
}

// SetAppExec sets field value
func (o *ApplicationResponseJson) SetAppExec(v bool) {
	o.AppExec = v
}

// GetUrlFopen returns the UrlFopen field value
func (o *ApplicationResponseJson) GetUrlFopen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UrlFopen
}

// GetUrlFopenOk returns a tuple with the UrlFopen field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetUrlFopenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlFopen, true
}

// SetUrlFopen sets field value
func (o *ApplicationResponseJson) SetUrlFopen(v bool) {
	o.UrlFopen = v
}

// GetFpmType returns the FpmType field value
func (o *ApplicationResponseJson) GetFpmType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FpmType
}

// GetFpmTypeOk returns a tuple with the FpmType field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetFpmTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FpmType, true
}

// SetFpmType sets field value
func (o *ApplicationResponseJson) SetFpmType(v string) {
	o.FpmType = v
}

// GetFpmMaxChildren returns the FpmMaxChildren field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetFpmMaxChildren() int32 {
	if o == nil || IsNil(o.FpmMaxChildren) {
		var ret int32
		return ret
	}
	return *o.FpmMaxChildren
}

// GetFpmMaxChildrenOk returns a tuple with the FpmMaxChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetFpmMaxChildrenOk() (*int32, bool) {
	if o == nil || IsNil(o.FpmMaxChildren) {
		return nil, false
	}
	return o.FpmMaxChildren, true
}

// HasFpmMaxChildren returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasFpmMaxChildren() bool {
	if o != nil && !IsNil(o.FpmMaxChildren) {
		return true
	}

	return false
}

// SetFpmMaxChildren gets a reference to the given int32 and assigns it to the FpmMaxChildren field.
func (o *ApplicationResponseJson) SetFpmMaxChildren(v int32) {
	o.FpmMaxChildren = &v
}

// GetFpmMaxRequests returns the FpmMaxRequests field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetFpmMaxRequests() int32 {
	if o == nil || IsNil(o.FpmMaxRequests) {
		var ret int32
		return ret
	}
	return *o.FpmMaxRequests
}

// GetFpmMaxRequestsOk returns a tuple with the FpmMaxRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetFpmMaxRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.FpmMaxRequests) {
		return nil, false
	}
	return o.FpmMaxRequests, true
}

// HasFpmMaxRequests returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasFpmMaxRequests() bool {
	if o != nil && !IsNil(o.FpmMaxRequests) {
		return true
	}

	return false
}

// SetFpmMaxRequests gets a reference to the given int32 and assigns it to the FpmMaxRequests field.
func (o *ApplicationResponseJson) SetFpmMaxRequests(v int32) {
	o.FpmMaxRequests = &v
}

// GetFpmStartServers returns the FpmStartServers field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetFpmStartServers() int32 {
	if o == nil || IsNil(o.FpmStartServers) {
		var ret int32
		return ret
	}
	return *o.FpmStartServers
}

// GetFpmStartServersOk returns a tuple with the FpmStartServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetFpmStartServersOk() (*int32, bool) {
	if o == nil || IsNil(o.FpmStartServers) {
		return nil, false
	}
	return o.FpmStartServers, true
}

// HasFpmStartServers returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasFpmStartServers() bool {
	if o != nil && !IsNil(o.FpmStartServers) {
		return true
	}

	return false
}

// SetFpmStartServers gets a reference to the given int32 and assigns it to the FpmStartServers field.
func (o *ApplicationResponseJson) SetFpmStartServers(v int32) {
	o.FpmStartServers = &v
}

// GetFpmMinSpareServers returns the FpmMinSpareServers field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetFpmMinSpareServers() int32 {
	if o == nil || IsNil(o.FpmMinSpareServers) {
		var ret int32
		return ret
	}
	return *o.FpmMinSpareServers
}

// GetFpmMinSpareServersOk returns a tuple with the FpmMinSpareServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetFpmMinSpareServersOk() (*int32, bool) {
	if o == nil || IsNil(o.FpmMinSpareServers) {
		return nil, false
	}
	return o.FpmMinSpareServers, true
}

// HasFpmMinSpareServers returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasFpmMinSpareServers() bool {
	if o != nil && !IsNil(o.FpmMinSpareServers) {
		return true
	}

	return false
}

// SetFpmMinSpareServers gets a reference to the given int32 and assigns it to the FpmMinSpareServers field.
func (o *ApplicationResponseJson) SetFpmMinSpareServers(v int32) {
	o.FpmMinSpareServers = &v
}

// GetFpmMaxSpareServers returns the FpmMaxSpareServers field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetFpmMaxSpareServers() int32 {
	if o == nil || IsNil(o.FpmMaxSpareServers) {
		var ret int32
		return ret
	}
	return *o.FpmMaxSpareServers
}

// GetFpmMaxSpareServersOk returns a tuple with the FpmMaxSpareServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetFpmMaxSpareServersOk() (*int32, bool) {
	if o == nil || IsNil(o.FpmMaxSpareServers) {
		return nil, false
	}
	return o.FpmMaxSpareServers, true
}

// HasFpmMaxSpareServers returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasFpmMaxSpareServers() bool {
	if o != nil && !IsNil(o.FpmMaxSpareServers) {
		return true
	}

	return false
}

// SetFpmMaxSpareServers gets a reference to the given int32 and assigns it to the FpmMaxSpareServers field.
func (o *ApplicationResponseJson) SetFpmMaxSpareServers(v int32) {
	o.FpmMaxSpareServers = &v
}

// GetPhpVersion returns the PhpVersion field value if set, zero value otherwise.
func (o *ApplicationResponseJson) GetPhpVersion() int32 {
	if o == nil || IsNil(o.PhpVersion) {
		var ret int32
		return ret
	}
	return *o.PhpVersion
}

// GetPhpVersionOk returns a tuple with the PhpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetPhpVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.PhpVersion) {
		return nil, false
	}
	return o.PhpVersion, true
}

// HasPhpVersion returns a boolean if a field has been set.
func (o *ApplicationResponseJson) HasPhpVersion() bool {
	if o != nil && !IsNil(o.PhpVersion) {
		return true
	}

	return false
}

// SetPhpVersion gets a reference to the given int32 and assigns it to the PhpVersion field.
func (o *ApplicationResponseJson) SetPhpVersion(v int32) {
	o.PhpVersion = &v
}

// GetSubroot returns the Subroot field value
func (o *ApplicationResponseJson) GetSubroot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subroot
}

// GetSubrootOk returns a tuple with the Subroot field value
// and a boolean to check if the value has been set.
func (o *ApplicationResponseJson) GetSubrootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subroot, true
}

// SetSubroot sets field value
func (o *ApplicationResponseJson) SetSubroot(v string) {
	o.Subroot = v
}

func (o ApplicationResponseJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResponseJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["db_name"] = o.DbName
	toSerialize["db_user"] = o.DbUser
	toSerialize["db_host"] = o.DbHost
	if !IsNil(o.DbPort) {
		toSerialize["db_port"] = o.DbPort
	}
	toSerialize["app_name"] = o.AppName
	if !IsNil(o.AppPort) {
		toSerialize["app_port"] = o.AppPort
	}
	toSerialize["app_path"] = o.AppPath
	toSerialize["app_version"] = o.AppVersion
	toSerialize["app_command"] = o.AppCommand
	toSerialize["app_lang_version"] = o.AppLangVersion
	if !IsNil(o.SymLinkPath) {
		toSerialize["sym_link_path"] = o.SymLinkPath
	}
	toSerialize["auto_site_url"] = o.AutoSiteUrl
	toSerialize["app_exec"] = o.AppExec
	toSerialize["url_fopen"] = o.UrlFopen
	toSerialize["fpm_type"] = o.FpmType
	if !IsNil(o.FpmMaxChildren) {
		toSerialize["fpm_max_children"] = o.FpmMaxChildren
	}
	if !IsNil(o.FpmMaxRequests) {
		toSerialize["fpm_max_requests"] = o.FpmMaxRequests
	}
	if !IsNil(o.FpmStartServers) {
		toSerialize["fpm_start_servers"] = o.FpmStartServers
	}
	if !IsNil(o.FpmMinSpareServers) {
		toSerialize["fpm_min_spare_servers"] = o.FpmMinSpareServers
	}
	if !IsNil(o.FpmMaxSpareServers) {
		toSerialize["fpm_max_spare_servers"] = o.FpmMaxSpareServers
	}
	if !IsNil(o.PhpVersion) {
		toSerialize["php_version"] = o.PhpVersion
	}
	toSerialize["subroot"] = o.Subroot
	return toSerialize, nil
}

type NullableApplicationResponseJson struct {
	value *ApplicationResponseJson
	isSet bool
}

func (v NullableApplicationResponseJson) Get() *ApplicationResponseJson {
	return v.value
}

func (v *NullableApplicationResponseJson) Set(val *ApplicationResponseJson) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponseJson) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponseJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponseJson(val *ApplicationResponseJson) *NullableApplicationResponseJson {
	return &NullableApplicationResponseJson{value: val, isSet: true}
}

func (v NullableApplicationResponseJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponseJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


