/*
 * AdobeSign Aggreements API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 6.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package agreements

// A predicate used to determine whether the condtion succeeds
type FormFieldConditionPredicate struct {
	// Index of the location of the form field used in the predicate
	FieldLocationIndex int32 `json:"fieldLocationIndex,omitempty"`
	// Name of the field whose value is the basis of predicate
	FieldName string `json:"fieldName,omitempty"`
	// Operator to be applied on the value of the predicate field.
	Operator string `json:"operator,omitempty"`
	// Value to compare against the value of the predicate's form field, using the specified operator
	Value string `json:"value,omitempty"`
}