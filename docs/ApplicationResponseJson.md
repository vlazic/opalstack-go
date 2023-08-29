# ApplicationResponseJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | **string** |  | 
**DbUser** | **string** |  | 
**DbHost** | **string** |  | 
**DbPort** | Pointer to **int32** |  | [optional] [default to 3306]
**AppName** | **string** |  | 
**AppPort** | Pointer to **int32** |  | [optional] [default to 8000]
**AppPath** | **string** |  | 
**AppVersion** | **string** |  | 
**AppCommand** | **string** |  | 
**AppLangVersion** | **string** |  | 
**SymLinkPath** | Pointer to **string** |  | [optional] [default to "/home/my_osuser/path/to/dir/"]
**AutoSiteUrl** | **bool** |  | 
**AppExec** | **bool** |  | 
**UrlFopen** | **bool** |  | 
**FpmType** | **string** |  | 
**FpmMaxChildren** | Pointer to **int32** |  | [optional] [default to 5]
**FpmMaxRequests** | Pointer to **int32** |  | [optional] [default to 500]
**FpmStartServers** | Pointer to **int32** |  | [optional] [default to 5]
**FpmMinSpareServers** | Pointer to **int32** |  | [optional] [default to 5]
**FpmMaxSpareServers** | Pointer to **int32** |  | [optional] [default to 5]
**PhpVersion** | Pointer to **int32** |  | [optional] [default to 74]
**Subroot** | **string** |  | 

## Methods

### NewApplicationResponseJson

`func NewApplicationResponseJson(dbName string, dbUser string, dbHost string, appName string, appPath string, appVersion string, appCommand string, appLangVersion string, autoSiteUrl bool, appExec bool, urlFopen bool, fpmType string, subroot string, ) *ApplicationResponseJson`

NewApplicationResponseJson instantiates a new ApplicationResponseJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseJsonWithDefaults

`func NewApplicationResponseJsonWithDefaults() *ApplicationResponseJson`

NewApplicationResponseJsonWithDefaults instantiates a new ApplicationResponseJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *ApplicationResponseJson) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *ApplicationResponseJson) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *ApplicationResponseJson) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbUser

`func (o *ApplicationResponseJson) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *ApplicationResponseJson) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *ApplicationResponseJson) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.


### GetDbHost

`func (o *ApplicationResponseJson) GetDbHost() string`

GetDbHost returns the DbHost field if non-nil, zero value otherwise.

### GetDbHostOk

`func (o *ApplicationResponseJson) GetDbHostOk() (*string, bool)`

GetDbHostOk returns a tuple with the DbHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbHost

`func (o *ApplicationResponseJson) SetDbHost(v string)`

SetDbHost sets DbHost field to given value.


### GetDbPort

`func (o *ApplicationResponseJson) GetDbPort() int32`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *ApplicationResponseJson) GetDbPortOk() (*int32, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *ApplicationResponseJson) SetDbPort(v int32)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *ApplicationResponseJson) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### GetAppName

`func (o *ApplicationResponseJson) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ApplicationResponseJson) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ApplicationResponseJson) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetAppPort

`func (o *ApplicationResponseJson) GetAppPort() int32`

GetAppPort returns the AppPort field if non-nil, zero value otherwise.

### GetAppPortOk

`func (o *ApplicationResponseJson) GetAppPortOk() (*int32, bool)`

GetAppPortOk returns a tuple with the AppPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPort

`func (o *ApplicationResponseJson) SetAppPort(v int32)`

SetAppPort sets AppPort field to given value.

### HasAppPort

`func (o *ApplicationResponseJson) HasAppPort() bool`

HasAppPort returns a boolean if a field has been set.

### GetAppPath

`func (o *ApplicationResponseJson) GetAppPath() string`

GetAppPath returns the AppPath field if non-nil, zero value otherwise.

### GetAppPathOk

`func (o *ApplicationResponseJson) GetAppPathOk() (*string, bool)`

GetAppPathOk returns a tuple with the AppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPath

`func (o *ApplicationResponseJson) SetAppPath(v string)`

SetAppPath sets AppPath field to given value.


### GetAppVersion

`func (o *ApplicationResponseJson) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *ApplicationResponseJson) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *ApplicationResponseJson) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.


### GetAppCommand

`func (o *ApplicationResponseJson) GetAppCommand() string`

GetAppCommand returns the AppCommand field if non-nil, zero value otherwise.

### GetAppCommandOk

`func (o *ApplicationResponseJson) GetAppCommandOk() (*string, bool)`

GetAppCommandOk returns a tuple with the AppCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCommand

`func (o *ApplicationResponseJson) SetAppCommand(v string)`

SetAppCommand sets AppCommand field to given value.


### GetAppLangVersion

`func (o *ApplicationResponseJson) GetAppLangVersion() string`

GetAppLangVersion returns the AppLangVersion field if non-nil, zero value otherwise.

### GetAppLangVersionOk

`func (o *ApplicationResponseJson) GetAppLangVersionOk() (*string, bool)`

GetAppLangVersionOk returns a tuple with the AppLangVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLangVersion

`func (o *ApplicationResponseJson) SetAppLangVersion(v string)`

SetAppLangVersion sets AppLangVersion field to given value.


### GetSymLinkPath

`func (o *ApplicationResponseJson) GetSymLinkPath() string`

GetSymLinkPath returns the SymLinkPath field if non-nil, zero value otherwise.

### GetSymLinkPathOk

`func (o *ApplicationResponseJson) GetSymLinkPathOk() (*string, bool)`

GetSymLinkPathOk returns a tuple with the SymLinkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymLinkPath

`func (o *ApplicationResponseJson) SetSymLinkPath(v string)`

SetSymLinkPath sets SymLinkPath field to given value.

### HasSymLinkPath

`func (o *ApplicationResponseJson) HasSymLinkPath() bool`

HasSymLinkPath returns a boolean if a field has been set.

### GetAutoSiteUrl

`func (o *ApplicationResponseJson) GetAutoSiteUrl() bool`

GetAutoSiteUrl returns the AutoSiteUrl field if non-nil, zero value otherwise.

### GetAutoSiteUrlOk

`func (o *ApplicationResponseJson) GetAutoSiteUrlOk() (*bool, bool)`

GetAutoSiteUrlOk returns a tuple with the AutoSiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSiteUrl

`func (o *ApplicationResponseJson) SetAutoSiteUrl(v bool)`

SetAutoSiteUrl sets AutoSiteUrl field to given value.


### GetAppExec

`func (o *ApplicationResponseJson) GetAppExec() bool`

GetAppExec returns the AppExec field if non-nil, zero value otherwise.

### GetAppExecOk

`func (o *ApplicationResponseJson) GetAppExecOk() (*bool, bool)`

GetAppExecOk returns a tuple with the AppExec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppExec

`func (o *ApplicationResponseJson) SetAppExec(v bool)`

SetAppExec sets AppExec field to given value.


### GetUrlFopen

`func (o *ApplicationResponseJson) GetUrlFopen() bool`

GetUrlFopen returns the UrlFopen field if non-nil, zero value otherwise.

### GetUrlFopenOk

`func (o *ApplicationResponseJson) GetUrlFopenOk() (*bool, bool)`

GetUrlFopenOk returns a tuple with the UrlFopen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFopen

`func (o *ApplicationResponseJson) SetUrlFopen(v bool)`

SetUrlFopen sets UrlFopen field to given value.


### GetFpmType

`func (o *ApplicationResponseJson) GetFpmType() string`

GetFpmType returns the FpmType field if non-nil, zero value otherwise.

### GetFpmTypeOk

`func (o *ApplicationResponseJson) GetFpmTypeOk() (*string, bool)`

GetFpmTypeOk returns a tuple with the FpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpmType

`func (o *ApplicationResponseJson) SetFpmType(v string)`

SetFpmType sets FpmType field to given value.


### GetFpmMaxChildren

`func (o *ApplicationResponseJson) GetFpmMaxChildren() int32`

GetFpmMaxChildren returns the FpmMaxChildren field if non-nil, zero value otherwise.

### GetFpmMaxChildrenOk

`func (o *ApplicationResponseJson) GetFpmMaxChildrenOk() (*int32, bool)`

GetFpmMaxChildrenOk returns a tuple with the FpmMaxChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpmMaxChildren

`func (o *ApplicationResponseJson) SetFpmMaxChildren(v int32)`

SetFpmMaxChildren sets FpmMaxChildren field to given value.

### HasFpmMaxChildren

`func (o *ApplicationResponseJson) HasFpmMaxChildren() bool`

HasFpmMaxChildren returns a boolean if a field has been set.

### GetFpmMaxRequests

`func (o *ApplicationResponseJson) GetFpmMaxRequests() int32`

GetFpmMaxRequests returns the FpmMaxRequests field if non-nil, zero value otherwise.

### GetFpmMaxRequestsOk

`func (o *ApplicationResponseJson) GetFpmMaxRequestsOk() (*int32, bool)`

GetFpmMaxRequestsOk returns a tuple with the FpmMaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpmMaxRequests

`func (o *ApplicationResponseJson) SetFpmMaxRequests(v int32)`

SetFpmMaxRequests sets FpmMaxRequests field to given value.

### HasFpmMaxRequests

`func (o *ApplicationResponseJson) HasFpmMaxRequests() bool`

HasFpmMaxRequests returns a boolean if a field has been set.

### GetFpmStartServers

`func (o *ApplicationResponseJson) GetFpmStartServers() int32`

GetFpmStartServers returns the FpmStartServers field if non-nil, zero value otherwise.

### GetFpmStartServersOk

`func (o *ApplicationResponseJson) GetFpmStartServersOk() (*int32, bool)`

GetFpmStartServersOk returns a tuple with the FpmStartServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpmStartServers

`func (o *ApplicationResponseJson) SetFpmStartServers(v int32)`

SetFpmStartServers sets FpmStartServers field to given value.

### HasFpmStartServers

`func (o *ApplicationResponseJson) HasFpmStartServers() bool`

HasFpmStartServers returns a boolean if a field has been set.

### GetFpmMinSpareServers

`func (o *ApplicationResponseJson) GetFpmMinSpareServers() int32`

GetFpmMinSpareServers returns the FpmMinSpareServers field if non-nil, zero value otherwise.

### GetFpmMinSpareServersOk

`func (o *ApplicationResponseJson) GetFpmMinSpareServersOk() (*int32, bool)`

GetFpmMinSpareServersOk returns a tuple with the FpmMinSpareServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpmMinSpareServers

`func (o *ApplicationResponseJson) SetFpmMinSpareServers(v int32)`

SetFpmMinSpareServers sets FpmMinSpareServers field to given value.

### HasFpmMinSpareServers

`func (o *ApplicationResponseJson) HasFpmMinSpareServers() bool`

HasFpmMinSpareServers returns a boolean if a field has been set.

### GetFpmMaxSpareServers

`func (o *ApplicationResponseJson) GetFpmMaxSpareServers() int32`

GetFpmMaxSpareServers returns the FpmMaxSpareServers field if non-nil, zero value otherwise.

### GetFpmMaxSpareServersOk

`func (o *ApplicationResponseJson) GetFpmMaxSpareServersOk() (*int32, bool)`

GetFpmMaxSpareServersOk returns a tuple with the FpmMaxSpareServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpmMaxSpareServers

`func (o *ApplicationResponseJson) SetFpmMaxSpareServers(v int32)`

SetFpmMaxSpareServers sets FpmMaxSpareServers field to given value.

### HasFpmMaxSpareServers

`func (o *ApplicationResponseJson) HasFpmMaxSpareServers() bool`

HasFpmMaxSpareServers returns a boolean if a field has been set.

### GetPhpVersion

`func (o *ApplicationResponseJson) GetPhpVersion() int32`

GetPhpVersion returns the PhpVersion field if non-nil, zero value otherwise.

### GetPhpVersionOk

`func (o *ApplicationResponseJson) GetPhpVersionOk() (*int32, bool)`

GetPhpVersionOk returns a tuple with the PhpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpVersion

`func (o *ApplicationResponseJson) SetPhpVersion(v int32)`

SetPhpVersion sets PhpVersion field to given value.

### HasPhpVersion

`func (o *ApplicationResponseJson) HasPhpVersion() bool`

HasPhpVersion returns a boolean if a field has been set.

### GetSubroot

`func (o *ApplicationResponseJson) GetSubroot() string`

GetSubroot returns the Subroot field if non-nil, zero value otherwise.

### GetSubrootOk

`func (o *ApplicationResponseJson) GetSubrootOk() (*string, bool)`

GetSubrootOk returns a tuple with the Subroot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubroot

`func (o *ApplicationResponseJson) SetSubroot(v string)`

SetSubroot sets Subroot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


