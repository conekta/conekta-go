# CompanyDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileClassification** | **string** | Classification of the document.  | Tipo de archivo              | Descripción                                               | | :--------------------------- | :-------------------------------------------------------- | | &#x60;id_legal_representative&#x60;      | identificación oficial frente                             | | &#x60;id_legal_representative_back&#x60; | identificación oficial atrás                              | | &#x60;cfdi&#x60;                         | Prueba de situación fiscal                                | | &#x60;constitutive_act_basic&#x60;       | Acta constitutiva                                         | | &#x60;proof_of_address&#x60;             | Comprobante de domicilio del negocio                      | | &#x60;power_of_attonery&#x60;            | Poderes de representación                                 | | &#x60;deposit_account_cover&#x60;        | Carátula de la cuenta de depósito                         | | &#x60;permit_casino&#x60;                | Permiso ante SEGOB                                        | | &#x60;license_sanitation&#x60;           | Licencia sanitaria de COFEPRIS                            | | &#x60;registration_tourism&#x60;         | Inscripción ante el Registro Nacional de Turismo (SECTUR) |  | 
**ContentType** | **string** | MIME type of the file. Allowed values depend on the &#x60;file_classification&#x60;. - &#x60;image/jpeg&#x60; - &#x60;image/png&#x60; - &#x60;application/pdf&#x60;  | 
**International** | Pointer to **bool** | Indicates if the document is international. Defaults to false. | [optional] 
**FileName** | **string** | Name of the file being uploaded. | 
**FileData** | **string** | Base64 encoded content of the file. | 

## Methods

### NewCompanyDocumentRequest

`func NewCompanyDocumentRequest(fileClassification string, contentType string, fileName string, fileData string, ) *CompanyDocumentRequest`

NewCompanyDocumentRequest instantiates a new CompanyDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyDocumentRequestWithDefaults

`func NewCompanyDocumentRequestWithDefaults() *CompanyDocumentRequest`

NewCompanyDocumentRequestWithDefaults instantiates a new CompanyDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileClassification

`func (o *CompanyDocumentRequest) GetFileClassification() string`

GetFileClassification returns the FileClassification field if non-nil, zero value otherwise.

### GetFileClassificationOk

`func (o *CompanyDocumentRequest) GetFileClassificationOk() (*string, bool)`

GetFileClassificationOk returns a tuple with the FileClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileClassification

`func (o *CompanyDocumentRequest) SetFileClassification(v string)`

SetFileClassification sets FileClassification field to given value.


### GetContentType

`func (o *CompanyDocumentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CompanyDocumentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CompanyDocumentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetInternational

`func (o *CompanyDocumentRequest) GetInternational() bool`

GetInternational returns the International field if non-nil, zero value otherwise.

### GetInternationalOk

`func (o *CompanyDocumentRequest) GetInternationalOk() (*bool, bool)`

GetInternationalOk returns a tuple with the International field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternational

`func (o *CompanyDocumentRequest) SetInternational(v bool)`

SetInternational sets International field to given value.

### HasInternational

`func (o *CompanyDocumentRequest) HasInternational() bool`

HasInternational returns a boolean if a field has been set.

### GetFileName

`func (o *CompanyDocumentRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CompanyDocumentRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CompanyDocumentRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileData

`func (o *CompanyDocumentRequest) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *CompanyDocumentRequest) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *CompanyDocumentRequest) SetFileData(v string)`

SetFileData sets FileData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


