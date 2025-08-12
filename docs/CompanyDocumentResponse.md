# CompanyDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileClassification** | **string** | Classification of the document.  | Tipo de archivo              | Descripción                                               | | :--------------------------- | :-------------------------------------------------------- | | &#x60;id_legal_representative&#x60;      | identificación oficial frente                             | | &#x60;id_legal_representative_back&#x60; | identificación oficial atrás                              | | &#x60;cfdi&#x60;                         | Prueba de situación fiscal                                | | &#x60;constitutive_act_basic&#x60;       | Acta constitutiva                                         | | &#x60;proof_of_address&#x60;             | Comprobante de domicilio del negocio                      | | &#x60;power_of_attonery&#x60;            | Poderes de representación                                 | | &#x60;deposit_account_cover&#x60;        | Carátula de la cuenta de depósito                         | | &#x60;permit_casino&#x60;                | Permiso ante SEGOB                                        | | &#x60;license_sanitation&#x60;           | Licencia sanitaria de COFEPRIS                            | | &#x60;registration_tourism&#x60;         | Inscripción ante el Registro Nacional de Turismo (SECTUR) |  | 
**FileName** | **string** | Name of the file as stored or processed. | 
**Status** | **string** | Current status of the document. | 

## Methods

### NewCompanyDocumentResponse

`func NewCompanyDocumentResponse(fileClassification string, fileName string, status string, ) *CompanyDocumentResponse`

NewCompanyDocumentResponse instantiates a new CompanyDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyDocumentResponseWithDefaults

`func NewCompanyDocumentResponseWithDefaults() *CompanyDocumentResponse`

NewCompanyDocumentResponseWithDefaults instantiates a new CompanyDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileClassification

`func (o *CompanyDocumentResponse) GetFileClassification() string`

GetFileClassification returns the FileClassification field if non-nil, zero value otherwise.

### GetFileClassificationOk

`func (o *CompanyDocumentResponse) GetFileClassificationOk() (*string, bool)`

GetFileClassificationOk returns a tuple with the FileClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileClassification

`func (o *CompanyDocumentResponse) SetFileClassification(v string)`

SetFileClassification sets FileClassification field to given value.


### GetFileName

`func (o *CompanyDocumentResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CompanyDocumentResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CompanyDocumentResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetStatus

`func (o *CompanyDocumentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyDocumentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyDocumentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


