package stationxml

/*
 *  WARNING: CODE GENERATED AUTOMATICALLY.
 *
 *  Use "go generate" to update these files.
 *
 *  THIS FILE SHOULD NOT BE EDITED BY HAND.
 *
 */

type OperatorType struct {
	Agency string `xml:"Agency"`

	Contact []PersonType `xml:"Contact,omitempty"`
}
