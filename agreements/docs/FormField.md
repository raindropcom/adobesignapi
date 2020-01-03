# FormField

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | **string** | Alignment of the text. | [optional] [default to null]
**Assignee** | **string** | Who the field is assigned to.  Either a participant set id, null, NOBODY or PREFILL. | [optional] [default to null]
**BackgroundColor** | **string** | Background color of the form field in RGB or HEX format | [optional] [default to null]
**BorderColor** | **string** | Color of the border of the field in RGB or HEX format | [optional] [default to null]
**BorderStyle** | **string** | Style of the border of the field. | [optional] [default to null]
**BorderWidth** | **float64** | Width of the border of the field in pixels | [optional] [default to null]
**Calculated** | **bool** | true if this field&#39;s value is calculated from an expression, else false | [optional] [default to null]
**ConditionalAction** | [***FormFieldConditionalAction**](FormFieldConditionalAction.md) | A predicate (or set of predicates) that determines whether this field is visible and enabled. | [optional] [default to null]
**ContentType** | **string** | Content Type of the form field. | [optional] [default to null]
**DefaultValue** | **string** | Default value of the form field | [optional] [default to null]
**DisplayFormat** | **string** | Format of the value of the field to be displayed based on the displayFormatType property. | [optional] [default to null]
**DisplayFormatType** | **string** | Format type of the text field. | [optional] [default to null]
**DisplayLabel** | **string** | Display label attached to the field | [optional] [default to null]
**FontColor** | **string** | Font color of the form field in RGB or HEX format | [optional] [default to null]
**FontName** | **string** | Font name of the form field | [optional] [default to null]
**FontSize** | **float64** | Font size of the form field in points | [optional] [default to null]
**HiddenOptions** | **[]string** | Text values which are hidden in a drop down form field | [optional] [default to null]
**Hyperlink** | [***FormFieldHyperlink**](FormFieldHyperlink.md) | Hyperlink-specific data (e.g. as url, link type) | [optional] [default to null]
**InputType** | **string** | Input type of the form field | [optional] [default to null]
**Locations** | [**[]FormFieldLocation**](FormFieldLocation.md) | All locations in a document where the form field is placed | [default to null]
**Masked** | **bool** | true if the input entered by the signer has to be masked (like password), false if it shouldn&#39;t be | [optional] [default to null]
**MaskingText** | **string** | Text to mask the masked form field | [optional] [default to null]
**MaxLength** | **int32** | Maximum length of the input text field in terms of no. of characters | [optional] [default to null]
**MaxValue** | **float64** | Upper bound of the number that can be entered by the signer | [optional] [default to null]
**MinLength** | **int32** | Minimum length of the input text field in terms of no. of characters | [optional] [default to null]
**MinValue** | **float64** | Lower bound of the number that can be entered by the signer | [optional] [default to null]
**Name** | **string** | The name of the form field | [default to null]
**Origin** | **string** | Origin of Form Field | [optional] [default to null]
**RadioCheckType** | **string** | The type of radio button (if field is radio button, identified by inputType). | [optional] [default to null]
**ReadOnly** | **bool** | true if it is a read-only field, else false | [optional] [default to null]
**Required** | **bool** | true if it is a mandatory field to be filled by the signer, else false | [optional] [default to null]
**Tooltip** | **string** | Text that appears while hovering over the field | [optional] [default to null]
**Validation** | **string** | Rule for validating the field value. | [optional] [default to null]
**ValidationData** | **string** | Further data for validating input with regards to the field&#39;s specified format. The contents and interpretation of formatData depends on the value of validation. | [optional] [default to null]
**ValidationErrMsg** | **string** | Error message to be shown to the signer if filled value doesn&#39;t match the validations of the form field | [optional] [default to null]
**ValueExpression** | **string** | Expression to calculate value of the form field | [optional] [default to null]
**Visible** | **bool** | If set to false, then the form field is hidden.  Otherwise, it is visible. | [optional] [default to null]
**VisibleOptions** | **[]string** | Text values which are visible in a drop down form field | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


