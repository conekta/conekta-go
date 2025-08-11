# CompanyResponseDocumentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileClassification** | Pointer to **string** | Classification of the document.  | Tipo de archivo              | Descripción                                               | | :--------------------------- | :-------------------------------------------------------- | | &#x60;id_legal_representative&#x60;      | identificación oficial frente                             | | &#x60;id_legal_representative_back&#x60; | identificación oficial atrás                              | | &#x60;cfdi&#x60;                         | Prueba de situación fiscal                                | | &#x60;constitutive_act_basic&#x60;       | Acta constitutiva                                         | | &#x60;proof_of_address&#x60;             | Comprobante de domicilio del negocio                      | | &#x60;power_of_attonery&#x60;            | Poderes de representación                                 | | &#x60;deposit_account_cover&#x60;        | Carátula de la cuenta de depósito                         | | &#x60;permit_casino&#x60;                | Permiso ante SEGOB                                        | | &#x60;license_sanitation&#x60;           | Licencia sanitaria de COFEPRIS                            | | &#x60;registration_tourism&#x60;         | Inscripción ante el Registro Nacional de Turismo (SECTUR) |  | [optional] 
**Status** | Pointer to **string** | The status of the document. | [optional] 
**FileName** | Pointer to **NullableString** | The name of the file. | [optional] 

## Methods

### NewCompanyResponseDocumentsInner

`func NewCompanyResponseDocumentsInner() *CompanyResponseDocumentsInner`

NewCompanyResponseDocumentsInner instantiates a new CompanyResponseDocumentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseDocumentsInnerWithDefaults

`func NewCompanyResponseDocumentsInnerWithDefaults() *CompanyResponseDocumentsInner`

NewCompanyResponseDocumentsInnerWithDefaults instantiates a new CompanyResponseDocumentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileClassification

`func (o *CompanyResponseDocumentsInner) GetFileClassification() string`

GetFileClassification returns the FileClassification field if non-nil, zero value otherwise.

### GetFileClassificationOk

`func (o *CompanyResponseDocumentsInner) GetFileClassificationOk() (*string, bool)`

GetFileClassificationOk returns a tuple with the FileClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileClassification

`func (o *CompanyResponseDocumentsInner) SetFileClassification(v string)`

SetFileClassification sets FileClassification field to given value.

### HasFileClassification

`func (o *CompanyResponseDocumentsInner) HasFileClassification() bool`

HasFileClassification returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyResponseDocumentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyResponseDocumentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyResponseDocumentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyResponseDocumentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFileName

`func (o *CompanyResponseDocumentsInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CompanyResponseDocumentsInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CompanyResponseDocumentsInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CompanyResponseDocumentsInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *CompanyResponseDocumentsInner) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *CompanyResponseDocumentsInner) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


