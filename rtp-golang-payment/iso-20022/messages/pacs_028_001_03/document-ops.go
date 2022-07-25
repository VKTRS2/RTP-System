// Package pacs_028_001_03
// Do not Edit. This stuff it's been automatically generated.
package pacs_028_001_03

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-rtp-cli/iso-20022/messages/xsdt"
	"reflect"
)

func (d *Document) Set(path string, src interface{}) error {

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(path)
	if err != nil {
		return err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, false)
	if err != nil {
		return err
	}

	return copy2Dest(path, values[0], src)
}

func (d *Document) Get(path string) (interface{}, error) {

	v := reflect.ValueOf(d)

	p, err := dotnotation.NewPath(path)
	if err != nil {
		return nil, err
	}
	paths := []dotnotation.DotPath{p}

	fields := d.mapper.TraversalsByName(v.Type(), paths)

	values := make([]interface{}, 1)
	err = fieldsByTraversal(v, fields, paths, values, true, true)
	if err != nil {
		return nil, err
	}

	/*
		rv := reflect.ValueOf(values[0])
		fmt.Println("Indirect type is:", reflect.Indirect(rv), reflect.Indirect(rv).Type(), rv.Kind(), rv.Elem(), rv.Elem().Type()) // prints main.CustomStruct

		if tv, ok := values[0].(*common.Max35Text); ok {
			return *tv, nil
		}
	*/

	return deref(path, values[0])
}

func copy2Dest(docPath string, dest, src interface{}) error {

	var err error
	switch typedDest := dest.(type) {
	case *common.ActiveOrHistoricCurrencyCode:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ActiveOrHistoricCurrencyCode data for" + docPath)
		}

		*typedDest, err = common.ToActiveOrHistoricCurrencyCode(src)
		return err
	case *common.AddressType2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.AddressType2Code data for" + docPath)
		}

		*typedDest, err = common.ToAddressType2Code(src)
		return err
	case *common.AnyBICDec2014Identifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.AnyBICDec2014Identifier data for" + docPath)
		}

		*typedDest, err = common.ToAnyBICDec2014Identifier(src)
		return err
	case *common.BICFIDec2014Identifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.BICFIDec2014Identifier data for" + docPath)
		}

		*typedDest, err = common.ToBICFIDec2014Identifier(src)
		return err
	case *common.ClearingChannel2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ClearingChannel2Code data for" + docPath)
		}

		*typedDest, err = common.ToClearingChannel2Code(src)
		return err
	case *common.CountryCode:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.CountryCode data for" + docPath)
		}

		*typedDest, err = common.ToCountryCode(src)
		return err
	case *common.CreditDebitCode:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.CreditDebitCode data for" + docPath)
		}

		*typedDest, err = common.ToCreditDebitCode(src)
		return err
	case *common.DocumentType3Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType3Code data for" + docPath)
		}

		*typedDest, err = common.ToDocumentType3Code(src)
		return err
	case *common.DocumentType6Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.DocumentType6Code data for" + docPath)
		}

		*typedDest, err = common.ToDocumentType6Code(src)
		return err
	case *common.Exact2NumericText:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact2NumericText data for" + docPath)
		}

		*typedDest, err = common.ToExact2NumericText(src)
		return err
	case *common.Exact4AlphaNumericText:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Exact4AlphaNumericText data for" + docPath)
		}

		*typedDest, err = common.ToExact4AlphaNumericText(src)
		return err
	case *common.ExternalAccountIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalAccountIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalAccountIdentification1Code(src)
		return err
	case *common.ExternalCashAccountType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashAccountType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalCashAccountType1Code(src)
		return err
	case *common.ExternalCashClearingSystem1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCashClearingSystem1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalCashClearingSystem1Code(src)
		return err
	case *common.ExternalCategoryPurpose1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalCategoryPurpose1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalCategoryPurpose1Code(src)
		return err
	case *common.ExternalClearingSystemIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalClearingSystemIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalClearingSystemIdentification1Code(src)
		return err
	case *common.ExternalDiscountAmountType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDiscountAmountType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalDiscountAmountType1Code(src)
		return err
	case *common.ExternalDocumentLineType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalDocumentLineType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalDocumentLineType1Code(src)
		return err
	case *common.ExternalFinancialInstitutionIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalFinancialInstitutionIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalFinancialInstitutionIdentification1Code(src)
		return err
	case *common.ExternalGarnishmentType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalGarnishmentType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalGarnishmentType1Code(src)
		return err
	case *common.ExternalLocalInstrument1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalLocalInstrument1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalLocalInstrument1Code(src)
		return err
	case *common.ExternalMandateSetupReason1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalMandateSetupReason1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalMandateSetupReason1Code(src)
		return err
	case *common.ExternalOrganisationIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalOrganisationIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalOrganisationIdentification1Code(src)
		return err
	case *common.ExternalPersonIdentification1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPersonIdentification1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalPersonIdentification1Code(src)
		return err
	case *common.ExternalProxyAccountType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalProxyAccountType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalProxyAccountType1Code(src)
		return err
	case *common.ExternalPurpose1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalPurpose1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalPurpose1Code(src)
		return err
	case *common.ExternalServiceLevel1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalServiceLevel1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalServiceLevel1Code(src)
		return err
	case *common.ExternalTaxAmountType1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ExternalTaxAmountType1Code data for" + docPath)
		}

		*typedDest, err = common.ToExternalTaxAmountType1Code(src)
		return err
	case *common.Frequency6Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Frequency6Code data for" + docPath)
		}

		*typedDest, err = common.ToFrequency6Code(src)
		return err
	case *common.IBAN2007Identifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.IBAN2007Identifier data for" + docPath)
		}

		*typedDest, err = common.ToIBAN2007Identifier(src)
		return err
	case *common.ISODate:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODate data for" + docPath)
		}

		*typedDest, err = common.ToISODate(src)
		return err
	case *common.ISODateTime:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.ISODateTime data for" + docPath)
		}

		*typedDest, err = common.ToISODateTime(src)
		return err
	case *common.LEIIdentifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.LEIIdentifier data for" + docPath)
		}

		*typedDest, err = common.ToLEIIdentifier(src)
		return err
	case *common.Max1025Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max1025Text data for" + docPath)
		}

		*typedDest, err = common.ToMax1025Text(src)
		return err
	case *common.Max128Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max128Text data for" + docPath)
		}

		*typedDest, err = common.ToMax128Text(src)
		return err
	case *common.Max140Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max140Text data for" + docPath)
		}

		*typedDest, err = common.ToMax140Text(src)
		return err
	case *common.Max15NumericText:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max15NumericText data for" + docPath)
		}

		*typedDest, err = common.ToMax15NumericText(src)
		return err
	case *common.Max16Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max16Text data for" + docPath)
		}

		*typedDest, err = common.ToMax16Text(src)
		return err
	case *common.Max2048Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max2048Text data for" + docPath)
		}

		*typedDest, err = common.ToMax2048Text(src)
		return err
	case *common.Max34Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max34Text data for" + docPath)
		}

		*typedDest, err = common.ToMax34Text(src)
		return err
	case *common.Max350Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max350Text data for" + docPath)
		}

		*typedDest, err = common.ToMax350Text(src)
		return err
	case *common.Max35Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max35Text data for" + docPath)
		}

		*typedDest, err = common.ToMax35Text(src)
		return err
	case *common.Max4Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max4Text data for" + docPath)
		}

		*typedDest, err = common.ToMax4Text(src)
		return err
	case *common.Max70Text:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Max70Text data for" + docPath)
		}

		*typedDest, err = common.ToMax70Text(src)
		return err
	case *common.NamePrefix2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.NamePrefix2Code data for" + docPath)
		}

		*typedDest, err = common.ToNamePrefix2Code(src)
		return err
	case *common.PaymentMethod4Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.PaymentMethod4Code data for" + docPath)
		}

		*typedDest, err = common.ToPaymentMethod4Code(src)
		return err
	case *common.PhoneNumber:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.PhoneNumber data for" + docPath)
		}

		*typedDest, err = common.ToPhoneNumber(src)
		return err
	case *common.PreferredContactMethod1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.PreferredContactMethod1Code data for" + docPath)
		}

		*typedDest, err = common.ToPreferredContactMethod1Code(src)
		return err
	case *common.Priority2Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.Priority2Code data for" + docPath)
		}

		*typedDest, err = common.ToPriority2Code(src)
		return err
	case *common.SequenceType3Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.SequenceType3Code data for" + docPath)
		}

		*typedDest, err = common.ToSequenceType3Code(src)
		return err
	case *common.SettlementMethod1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.SettlementMethod1Code data for" + docPath)
		}

		*typedDest, err = common.ToSettlementMethod1Code(src)
		return err
	case *common.TaxRecordPeriod1Code:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.TaxRecordPeriod1Code data for" + docPath)
		}

		*typedDest, err = common.ToTaxRecordPeriod1Code(src)
		return err
	case *common.UUIDv4Identifier:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling common.UUIDv4Identifier data for" + docPath)
		}

		*typedDest, err = common.ToUUIDv4Identifier(src)
		return err
	case *xsdt.AnyType:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.AnyType data for" + docPath)
		}

		*typedDest, err = xsdt.ToAnyType(src)
		return err
	case *xsdt.Boolean:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Boolean data for" + docPath)
		}

		*typedDest, err = xsdt.ToBoolean(src)
		return err
	case *xsdt.Decimal:
		if typedDest == nil {
			return errors.New("nil pointer... in unmarshalling xsdt.Decimal data for" + docPath)
		}

		*typedDest, err = xsdt.ToDecimal(src)
		return err
	default:
		return fmt.Errorf("could not find the type to node %s of type %v", docPath, dest)
	}

}

func deref(docPath string, val interface{}) (interface{}, error) {

	var err error
	switch tv := val.(type) {
	case *common.ActiveOrHistoricCurrencyCode:
		return *tv, nil
	case *common.AddressType2Code:
		return *tv, nil
	case *common.AnyBICDec2014Identifier:
		return *tv, nil
	case *common.BICFIDec2014Identifier:
		return *tv, nil
	case *common.ClearingChannel2Code:
		return *tv, nil
	case *common.CountryCode:
		return *tv, nil
	case *common.CreditDebitCode:
		return *tv, nil
	case *common.DocumentType3Code:
		return *tv, nil
	case *common.DocumentType6Code:
		return *tv, nil
	case *common.Exact2NumericText:
		return *tv, nil
	case *common.Exact4AlphaNumericText:
		return *tv, nil
	case *common.ExternalAccountIdentification1Code:
		return *tv, nil
	case *common.ExternalCashAccountType1Code:
		return *tv, nil
	case *common.ExternalCashClearingSystem1Code:
		return *tv, nil
	case *common.ExternalCategoryPurpose1Code:
		return *tv, nil
	case *common.ExternalClearingSystemIdentification1Code:
		return *tv, nil
	case *common.ExternalDiscountAmountType1Code:
		return *tv, nil
	case *common.ExternalDocumentLineType1Code:
		return *tv, nil
	case *common.ExternalFinancialInstitutionIdentification1Code:
		return *tv, nil
	case *common.ExternalGarnishmentType1Code:
		return *tv, nil
	case *common.ExternalLocalInstrument1Code:
		return *tv, nil
	case *common.ExternalMandateSetupReason1Code:
		return *tv, nil
	case *common.ExternalOrganisationIdentification1Code:
		return *tv, nil
	case *common.ExternalPersonIdentification1Code:
		return *tv, nil
	case *common.ExternalProxyAccountType1Code:
		return *tv, nil
	case *common.ExternalPurpose1Code:
		return *tv, nil
	case *common.ExternalServiceLevel1Code:
		return *tv, nil
	case *common.ExternalTaxAmountType1Code:
		return *tv, nil
	case *common.Frequency6Code:
		return *tv, nil
	case *common.IBAN2007Identifier:
		return *tv, nil
	case *common.ISODate:
		return *tv, nil
	case *common.ISODateTime:
		return *tv, nil
	case *common.LEIIdentifier:
		return *tv, nil
	case *common.Max1025Text:
		return *tv, nil
	case *common.Max128Text:
		return *tv, nil
	case *common.Max140Text:
		return *tv, nil
	case *common.Max15NumericText:
		return *tv, nil
	case *common.Max16Text:
		return *tv, nil
	case *common.Max2048Text:
		return *tv, nil
	case *common.Max34Text:
		return *tv, nil
	case *common.Max350Text:
		return *tv, nil
	case *common.Max35Text:
		return *tv, nil
	case *common.Max4Text:
		return *tv, nil
	case *common.Max70Text:
		return *tv, nil
	case *common.NamePrefix2Code:
		return *tv, nil
	case *common.PaymentMethod4Code:
		return *tv, nil
	case *common.PhoneNumber:
		return *tv, nil
	case *common.PreferredContactMethod1Code:
		return *tv, nil
	case *common.Priority2Code:
		return *tv, nil
	case *common.SequenceType3Code:
		return *tv, nil
	case *common.SettlementMethod1Code:
		return *tv, nil
	case *common.TaxRecordPeriod1Code:
		return *tv, nil
	case *common.UUIDv4Identifier:
		return *tv, nil
	case *xsdt.AnyType:
		return *tv, nil
	case *xsdt.Boolean:
		return *tv, nil
	case *xsdt.Decimal:
		return *tv, nil
	default:
		err = fmt.Errorf("could not find the type to node %s of type %v", docPath, val)
	}

	return val, err
}
