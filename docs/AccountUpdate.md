# AccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PaymentProcessor** | Pointer to [**AccountUpdatePaymentProcessorEnum**](AccountUpdatePaymentProcessorEnum.md) |  | [optional] 

## Methods

### NewAccountUpdate

`func NewAccountUpdate(id string, ) *AccountUpdate`

NewAccountUpdate instantiates a new AccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateWithDefaults

`func NewAccountUpdateWithDefaults() *AccountUpdate`

NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentProcessor

`func (o *AccountUpdate) GetPaymentProcessor() AccountUpdatePaymentProcessorEnum`

GetPaymentProcessor returns the PaymentProcessor field if non-nil, zero value otherwise.

### GetPaymentProcessorOk

`func (o *AccountUpdate) GetPaymentProcessorOk() (*AccountUpdatePaymentProcessorEnum, bool)`

GetPaymentProcessorOk returns a tuple with the PaymentProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessor

`func (o *AccountUpdate) SetPaymentProcessor(v AccountUpdatePaymentProcessorEnum)`

SetPaymentProcessor sets PaymentProcessor field to given value.

### HasPaymentProcessor

`func (o *AccountUpdate) HasPaymentProcessor() bool`

HasPaymentProcessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


