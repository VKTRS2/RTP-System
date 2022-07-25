## tpm-rtp-cli

### What's that and why doing it
This project main objective is the generation of artifacts to manipulate IS20022 messages using Golang code. 
On the net there is stuff to generate go structs from schemas but all this stuff lacks some bells and whistles to easily use this structures. This generation is not general
if we mean a general generation of structs from schema files (it lacks some support of builtin types that are not used in IS20022 schemas for example) but tries to offer an approach useful for the task at hand.
For example, it doesn't generate a schema to a single file in a single package but tries to find common types across different messages and split them in different packages.
As such, you will have a package for each message and a common package where simple and complex types shared among messages are put together to simplify the creation of a message from a different one used as a source.
More over it provides some convenience methods for setting and getting values out of the structures with some protection for arrays and nil pointers. 
I hope you will find this stuff edible to get life easier down the road considering that:

* arrays are a matter of the schema where plurality gets in the way,
* pointers are a fact of life when generating empty structs is not an option
* those messages can have a number of nodes ranging from 3000 up to 8500 (considering leaves and intermediate)
* real life cases set a very few number of nodes: might be in the range of 30 or 40 leaves: more intermediate nodes than leaves as to speak. As such these trees are pretty sparse.
* the code to map some data to a message can be very much specific and the code can result very much hard-coded as to speak.

### Layout
A brief description of the layout of the project and what you can find where.

* [iso-20022](iso-20022). In here you can find the go generated structures of some iso20022 messages. This stuff is self-consistent in that it doesn't depend on external packages beside the standard ones.
* [cmds/tpm-rtp-cli](cmds/tpm-rtp-cli). It contains the driver of the command line utilities. At the moment the generation command is under way. It doesn't do much.
* [cmds/generate](cmds/generate). It is the generation stuff that is used to produce the iso-20022 folder content. The output has been produced by a test case [cmds/generate/golang/generate_test.go](cmds/generate/golang/generate_test.go).

### iso-20022 folder Layout
The folder contains artefacts produced by the generator and other stuff. More precisely.

* [iso-20022/schemas](iso-20022/schemas). Contains the xsd used for the generation.
* [iso-20022/messages](iso-20022/messages). In here a number of folders.
    - One folder for each message named after the message
    - [iso-20022/messages/common](iso-20022/messages/common) a common folder that contains stuff related to types that are shared among messages.
    - [iso-20022/messages/xsdt](iso-20022/messages/xsdt) a folder for definitions of go types related to schema builtin types.
* [iso-20022/examples](iso-20022/examples). A few test cases hand coded to give the idea of the structuring in some real case.

#### iso-20022 message folder
The folders are named after the message in a go-frienly way (i.e. `pain.013.001.07` gets to `pain_013_001_07`). 
Each folder contains a number of files. The size and the complexity of the messages and the size of the generated stuff has forced to distribute code on various files.

* document.go, document-paths.go, document-set-ops.go: they contain the document structure definition and the document level methods. They are considered below in a dedicated paragraph.
* complex-types.go: contains the declaration of structures not shared and specific to the message (at tleast in the context of the messages considered)
* complex-types-ops.go: contains the methods of the structures.
* <message-id>_test.go (where message-id stands for `pain_013_001_07` etc..): this file contains a literal sample of the document with sample values, it is meant to provide an overview of the message and to 
extract snippet that might be used somewhere else.

##### Typical layout of a structure
The snippet below provide an example of the type of stuff that is in there (Methods and structure might change over time but pretty much this is the idea). In the snippet below it is worth of note:

* Optional properties of complex types are pointers to structs.
* Mandatory structs are normal structs.
* Plurals are handled as array of items (not pointers to).

```
// CreditTransferTransaction35 type definition
type CreditTransferTransaction35 struct {
	PmtId           PaymentIdentification6                               `xml:"PmtId"`
	PmtTpInf        *common.PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty"`
	PmtCond         *common.PaymentCondition1                            `xml:"PmtCond,omitempty"`
	Amt             common.AmountType4Choice                             `xml:"Amt"`
	ChrgBr          common.ChargeBearerType1Code                         `xml:"ChrgBr"`
	ChqInstr        *Cheque11                                            `xml:"ChqInstr,omitempty"`
	UltmtDbtr       *common.PartyIdentification135                       `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt2      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt3      *common.BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty"`
	CdtrAgt         common.BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	Cdtr            common.PartyIdentification135                        `xml:"Cdtr"`
	CdtrAcct        *common.CashAccount38                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       *common.PartyIdentification135                       `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                       `xml:"InstrForCdtrAgt,omitempty"`
	Purp            *common.Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                               `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation8                                     `xml:"Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                                `xml:"RltdRmtInf,omitempty"`
	RmtInf          *common.RemittanceInformation16                      `xml:"RmtInf,omitempty"`
	NclsdFile       []common.Document12                                  `xml:"NclsdFile,omitempty"`
	SplmtryData     []common.SupplementaryData1                          `xml:"SplmtryData,omitempty"`
}

// IsValid checks if CreditTransferTransaction35 is valid
func (s CreditTransferTransaction35) IsValid(optional bool) bool {

	valid := true
	valid = valid && s.PmtId.IsValid(false)
	valid = valid && (s.PmtTpInf == nil || (s.PmtTpInf != nil && s.PmtTpInf.IsValid(true)))

	valid = valid && (s.PmtCond == nil || (s.PmtCond != nil && s.PmtCond.IsValid(true)))

	valid = valid && s.Amt.IsValid(false)
	valid = valid && s.ChrgBr.IsValid(false)
	valid = valid && (s.ChqInstr == nil || (s.ChqInstr != nil && s.ChqInstr.IsValid(true)))

	valid = valid && (s.UltmtDbtr == nil || (s.UltmtDbtr != nil && s.UltmtDbtr.IsValid(true)))

	valid = valid && (s.IntrmyAgt1 == nil || (s.IntrmyAgt1 != nil && s.IntrmyAgt1.IsValid(true)))

	valid = valid && (s.IntrmyAgt2 == nil || (s.IntrmyAgt2 != nil && s.IntrmyAgt2.IsValid(true)))

	valid = valid && (s.IntrmyAgt3 == nil || (s.IntrmyAgt3 != nil && s.IntrmyAgt3.IsValid(true)))

	valid = valid && s.CdtrAgt.IsValid(false)
	valid = valid && s.Cdtr.IsValid(false)
	valid = valid && (s.CdtrAcct == nil || (s.CdtrAcct != nil && s.CdtrAcct.IsValid(true)))

	valid = valid && (s.UltmtCdtr == nil || (s.UltmtCdtr != nil && s.UltmtCdtr.IsValid(true)))

	for j := 0; j < len(s.InstrForCdtrAgt); j++ {
		valid = valid && s.InstrForCdtrAgt[j].IsValid(true)
	}

	valid = valid && (s.Purp == nil || (s.Purp != nil && s.Purp.IsValid(true)))

	for j := 0; j < len(s.RgltryRptg); j++ {
		valid = valid && s.RgltryRptg[j].IsValid(true)
	}

	valid = valid && (s.Tax == nil || (s.Tax != nil && s.Tax.IsValid(true)))

	for j := 0; j < len(s.RltdRmtInf); j++ {
		valid = valid && s.RltdRmtInf[j].IsValid(true)
	}

	valid = valid && (s.RmtInf == nil || (s.RmtInf != nil && s.RmtInf.IsValid(true)))

	for j := 0; j < len(s.NclsdFile); j++ {
		valid = valid && s.NclsdFile[j].IsValid(true)
	}

	for j := 0; j < len(s.SplmtryData); j++ {
		valid = valid && s.SplmtryData[j].IsValid(true)
	}

	return valid
}
```

##### document files
In here a brief explanation of what is in the document* files that are found in the message folder.

* document.go: it is the structure of the document and a few methods. Previously put in the `complex-types.go` has been extracted since the Document is the root of the thing and is easier to find.
* document-paths.go: it is a set of constants with each one constant to each different path from the root to each node (not just the leaves).
* document-set-ops.goo: they contain a set function for setting each individual leaf node of the document by providing the path to the leaf by means of a string and a value. The consts in the document-paths.go can be used as the first parameter
of the document set function. The document set function use a bit of reflection for traversing the tree. The method has yet some limitations in handling arrays. In my current use case that's less than an issue but is clearly an open point.

```
	Path_CdtrPmtActvtnReq                                                                                            = "CdtrPmtActvtnReq"
	Path_CdtrPmtActvtnReq_GrpHdr                                                                                     = "CdtrPmtActvtnReq.GrpHdr"
	Path_CdtrPmtActvtnReq_GrpHdr_MsgId                                                                               = "CdtrPmtActvtnReq.GrpHdr.MsgId"
	Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm                                                                             = "CdtrPmtActvtnReq.GrpHdr.CreDtTm"
	Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs                                                                             = "CdtrPmtActvtnReq.GrpHdr.NbOfTxs"
	Path_CdtrPmtActvtnReq_GrpHdr_CtrlSum                                                                             = "CdtrPmtActvtnReq.GrpHdr.CtrlSum"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty                                                                            = "CdtrPmtActvtnReq.GrpHdr.InitgPty"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm                                                                         = "CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr                                                                    = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr"
	Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_PstlAdr_AdrTp                                                              = "CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.AdrTp"
```



#### iso-20022 common folder
The folder contains declaration of simple types and complex types shared between the messages. Files have been split in declarations and methods.

* complex-types.go: contains the declaration of shared complex types.
* complex-types-ops.go: contains the methods of shared complex types.
* simple-types.go: contains the declaration of shared simple types.
* complex-types-ops.go: contains the methods of shared simple types.
* util.go: a few functions referenced by other files and with bridge scope for some external custom utility projects. The idea is to limit external dependencies 
to be easily managed.

Below a snippet of the generation of a simple string type with enum restriction (Methods and structure might change over time but pretty much this is the idea). 

```
// ChargeBearerType1Code May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code xsdt.String

/*
 * ChargeBearerType1Code Ops
 */

const (
	ChargeBearerType1CodeDEBT = "DEBT"
	ChargeBearerType1CodeCRED = "CRED"
	ChargeBearerType1CodeSHAR = "SHAR"
	ChargeBearerType1CodeSLEV = "SLEV"
)

var ChargeBearerType1CodeEnumRestriction = []string{
     ChargeBearerType1CodeDEBT, ChargeBearerType1CodeCRED, ChargeBearerType1CodeSHAR, ChargeBearerType1CodeSLEV}

// IsValid checks if ChargeBearerType1Code of type String is valid
func (t ChargeBearerType1Code) IsValid(optional bool) bool {

	valid := xsdt.String(t).IsValid(optional)
	valid = valid && isEnumRestrictionValid(t.String(), ChargeBearerType1CodeEnumRestriction)

	return valid
}

// String method for easy conversion
func (t ChargeBearerType1Code) String() string {
	return string(t)
}

// ToChargeBearerType1Code method for easy conversion with application of restrictions
func ToChargeBearerType1Code(s string) (ChargeBearerType1Code, error) {
	if !isEnumRestrictionValid(s, ChargeBearerType1CodeEnumRestriction) {
		return "", fmt.Errorf("cannot satisfy enum restriction for %s of type ChargeBearerType1Code", s)
	}

	return ChargeBearerType1Code(s), nil
}

// MustToChargeBearerType1Code method for easy conversion with application of restrictions. Panics on error.
func MustToChargeBearerType1Code(s string) ChargeBearerType1Code {
	v, err := ToChargeBearerType1Code(s)
	if err != nil {
		panic(err)
	}

	return v
}

// ChargeBearerType1CodeExample method for generation of valid sample data
func ChargeBearerType1CodeExample() ChargeBearerType1Code {
	return ChargeBearerType1Code(generateSampleDataWithEnumRestriction(ChargeBearerType1CodeEnumRestriction))

}
```

#### iso-20022 examples
In the [iso-20022/examples](iso-20022/examples) are contained a few examples to show the use of structures in literals.
The use of the `MustTo` type of function is to allow the literal to be layed-out without dealing with error (The `MustTo`)
version panics in case of problems.


```
d := pain_013_001_07.Document{
	CdtrPmtActvtnReq: pain_013_001_07.CreditorPaymentActivationRequestV07{
   
		GrpHdr: pain_013_001_07.GroupHeader78{
			MsgId:   common.MustToMax35Text("pain013-DS01-20220322"),
			CreDtTm: common.MustToISODateTime(time.Now()),
			NbOfTxs: common.MustToMax15NumericText("1"),
			InitgPty: common.PartyIdentification135{
				Nm: common.MustToMax140Text("Poste Italiane"),
				Id: &common.Party38Choice{
					OrgId: &common.OrganisationIdentification29{
						Othr: []common.GenericOrganisationIdentification1{
							{
								Id:   common.MustToMax35Text("01114601006"),
								Issr: common.MustToMax35Text("POSTE ITALIANE"),
							},
						},
					},
				},
			},
		},
		PmtInf: []pain_013_001_07.PaymentInstruction31{
			{
				PmtInfId: common.MustToMax35Text("Fatt2022-000001-2022-03-22-11:50:45"),
				PmtMtd:   common.MustToPaymentMethod7Code(common.PaymentMethod7CodeTRF),
				PmtTpInf: &common.PaymentTypeInformation26{
					SvcLvl: []common.ServiceLevel8Choice{
						{
							Cd: common.MustToExternalServiceLevel1Code("SRTP"),
						},
					},
					LclInstrm: &common.LocalInstrument2Choice{
						Prtry: common.MustToMax35Text("NOTPROVIDED"),
					},
					CtgyPurp: &common.CategoryPurpose1Choice{
						Cd: common.MustToExternalCategoryPurpose1Code("OTHR"),
					},
				},
				ReqdExctnDt: common.DateAndDateTime2Choice{
					Dt: common.MustToISODate(time.Now()),
				},
				XpryDt: &common.DateAndDateTime2Choice{
					Dt: common.MustToISODate(time.Now()),
				},
				Dbtr: common.PartyIdentification135{
					Nm: common.MustToMax140Text("LoremIpsumSPA"),
					PstlAdr: &common.PostalAddress24{
						AdrLine: []common.Max70Text{
							common.MustToMax70Text("ViaLoremIpsum 30 Roma"),
						},
					},
					Id: &common.Party38Choice{
						PrvtId: &common.PersonIdentification13{
							Othr: []common.GenericPersonIdentification1{
								{
									Id: common.MustToMax35Text("123456789"),
									SchmeNm: &common.PersonIdentificationSchemeName1Choice{
										Cd: common.MustToExternalPersonIdentification1Code("POID"),
									},
								},
							},
						},
					},
				},
				DbtrAcct: &common.CashAccount38{
					Id: common.AccountIdentification4Choice{
						IBAN: common.MustToIBAN2007Identifier("IT60X0760111101000000123456"),
					},
				},
				DbtrAgt: common.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: common.FinancialInstitutionIdentification18{
						BICFI: common.MustToBICFIDec2014Identifier("BPPIITRRXXX"),
					},
				},
				CdtTrfTx: []pain_013_001_07.CreditTransferTransaction35{
					{
						PmtId: pain_013_001_07.PaymentIdentification6{
							EndToEndId: common.MustToMax35Text("fatt2022-000001"),
						},
						PmtTpInf: &common.PaymentTypeInformation26{
							SvcLvl: []common.ServiceLevel8Choice{
								{
									Cd: common.MustToExternalServiceLevel1Code("SRTP"),
								},
							},
							LclInstrm: &common.LocalInstrument2Choice{
								Prtry: common.MustToMax35Text("NOTPROVIDED"),
							},
							CtgyPurp: &common.CategoryPurpose1Choice{
								Cd: common.MustToExternalCategoryPurpose1Code("OTHR"),
							},
						},
						PmtCond: &common.PaymentCondition1{
							AmtModAllwd:   false,
							EarlyPmtAllwd: false,
							GrntedPmtReqd: false,
						},
						Amt: common.AmountType4Choice{
							InstdAmt: &common.ActiveOrHistoricCurrencyAndAmount{
								Ccy:   common.MustToActiveOrHistoricCurrencyCode("EUR"),
								Value: xsdt.ToDecimal(535.35),
							},
						},
						ChrgBr: common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSLEV),
						CdtrAgt: common.BranchAndFinancialInstitutionIdentification6{
							FinInstnId: common.FinancialInstitutionIdentification18{
								BICFI: common.MustToBICFIDec2014Identifier("BPPIITRRXXX"),
							},
						},
						Cdtr: common.PartyIdentification135{
							Nm: common.MustToMax140Text("AFC Poste Italiane"),
							PstlAdr: &common.PostalAddress24{
								AdrLine: []common.Max70Text{
									common.MustToMax70Text("Via Del Creditore 75 Roma"),
								},
							},
							Id: &common.Party38Choice{
								OrgId: &common.OrganisationIdentification29{
									Othr: []common.GenericOrganisationIdentification1{
										{
											Id: common.MustToMax35Text("0468651441"),
											SchmeNm: &common.OrganisationIdentificationSchemeName1Choice{
												Cd: common.MustToExternalOrganisationIdentification1Code("BOID"),
											},
										},
									},
								},
							},
						},
						CdtrAcct: &common.CashAccount38{
							Id: common.AccountIdentification4Choice{
								IBAN: common.MustToIBAN2007Identifier("IT60X0760111101000004545561"),
							},
						},
						RmtInf: &common.RemittanceInformation16{
							Ustrd: []common.Max140Text{
								common.MustToMax140Text("AT41/fatt2022-000001/AT05/pagamento fattura"),
							},
						},
					},
				},
			},
		},
	},
}
```

This second example shows the approach of using document apis for setting the values and obtaining the exact same document. 
This second formulation is more convenient if you have a mapping issue and cannot layout the document as a whole single literal as in the example above.

```
d := pain_013_001_07.NewDocument()
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId, common.MustToMax35Text("pain013-DS01-20220322"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_CreDtTm, common.MustToISODateTime(time.Now()))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_NbOfTxs, common.MustToMax15NumericText("1"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Nm, common.MustToMax140Text(NameOfInitiatingParty))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id, common.MustToMax35Text(InitiatingPartyOrgId))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Issr, common.MustToMax35Text(InitiatingPartyOrgIssr))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtInfId, common.MustToMax35Text(PaymentInstructionId))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtMtd, common.MustToPaymentMethod7Code(common.PaymentMethod7CodeTRF))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_SvcLvl_Cd, common.MustToExternalServiceLevel1Code("SRTP"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_LclInstrm_Prtry, common.MustToMax35Text("NOTPROVIDED"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_PmtTpInf_CtgyPurp_Cd, common.MustToExternalCategoryPurpose1Code("OTHR"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_ReqdExctnDt_Dt, common.MustToISODate(time.Now()))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_XpryDt_Dt, common.MustToISODate(time.Now()))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Nm, common.MustToMax140Text(DebitorName))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_PstlAdr_AdrLine, common.MustToMax70Text(DebitorAddress))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_Id, common.MustToMax35Text("123456789"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_Dbtr_Id_PrvtId_Othr_SchmeNm_Cd, common.MustToExternalPersonIdentification1Code("POID"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_DbtrAcct_Id_IBAN, common.MustToIBAN2007Identifier(DebitorIBAN))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_DbtrAgt_FinInstnId_BICFI, common.MustToBICFIDec2014Identifier(DebitorAgentBIC))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtId_EndToEndId, common.MustToMax35Text(InvoiceNumber))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_SvcLvl_Cd, common.MustToExternalServiceLevel1Code("SRTP"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_LclInstrm_Prtry, common.MustToMax35Text("NOTPROVIDED"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtTpInf_CtgyPurp_Cd, common.MustToExternalCategoryPurpose1Code("OTHR"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_AmtModAllwd, false)
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_EarlyPmtAllwd, false)
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_PmtCond_GrntedPmtReqd, false)
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Ccy, common.MustToActiveOrHistoricCurrencyCode("EUR"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Amt_InstdAmt_Value, xsdt.MustToDecimal(535.35))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_ChrgBr, common.MustToChargeBearerType1Code(common.ChargeBearerType1CodeSLEV))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAgt_FinInstnId_BICFI, common.MustToBICFIDec2014Identifier(CreditorAgentBIC))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Nm, common.MustToMax140Text(CreditorName))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_PstlAdr_AdrLine, common.MustToMax70Text(CreditorAddress))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_Id, common.MustToMax35Text(CreditorOrgId))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_Cdtr_Id_OrgId_Othr_SchmeNm_Cd, common.MustToExternalOrganisationIdentification1Code("BOID"))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_CdtrAcct_Id_IBAN, common.MustToIBAN2007Identifier(CreditorIBAN))
err = d.Set(pain_013_001_07.Path_CdtrPmtActvtnReq_PmtInf_CdtTrfTx_RmtInf_Ustrd, common.MustToMax140Text(RemittanceInfo))
```

```go
func paymentInstruct(doc *Document) (paymentInstruct, error) {
	output := paymentInstruct{
		AdrLine:                doc.AdrLine.Max70Text,                                     // "<AdrLine>"
		Agt:                    doc.Agt.BranchAndFinancialInstitutionIdentification5,      // "<Agt>"
		Assgne:                 doc.Assgne.Party35Choice,                                  // "<Assgne>"
		Assgnr:                 doc.Assgnr.Party7Choice,                                   // "<Assgnr>"
		BICFI:                  doc.BICFI.BankIdentificationCode,                          // "<BICFI>"
		BizMsgIdr:              doc.BizMsgIdr.BusinessMessageIdentifier,                   // "<BizMsgIdr>"
		BldgNb:                 doc.BldgNb.Max16Text,                                      // "<BldgNb>"
		// CanonicalizationMethod https://docs.microsoft.com/en-us/dotnet/api/system.security.cryptography.xml.signedinfo.canonicalizationmethod?view=dotnet-plat-ext-6.0
		// CanonicalizationMethod https://docs.oracle.com/javase/8/docs/api/javax/xml/crypto/dsig/CanonicalizationMethod.html
		// CanonicalizationMethod https://docs.oracle.com/en/java/javase/13/docs/api/java.xml.crypto/javax/xml/crypto/dsig/CanonicalizationMethod.html
		// CanonicalizationMethod https://www.di-mgt.com.au/xmldsig-c14n.html
		CanonicalizationMethod: doc.CanonicalizationMethod.CanonicalizationMethod,         // "<CanonicalizationMethod>"
		CdtrAgt:                doc.CdtrAgt.BranchAndFinancialInstitutionIdentification4,  // "<CdtrAgt>"
		CdtTrfTxInf:            doc.CdtTrfTxInf.CreditTransferTransactionInformation11,    // "<CdtTrfTxInf>"
		ChrgBr:                 doc.ChrgBr.ChargeBearerType1Code,                          // "<ChrgBr>"
		ChrgsInf:               doc.ChrgsInf.ChargesInformation5,                          // "<ChrgsInf>"
		Conf:                   doc.Conf.ExternalInvestigationExecutionConfirmation1Co,    // "<Conf>"
		CreDt:                  doc.CreDt.ISONormalisedDateTime,                           // "<CreDt>"
		CreDtTm:                doc.CreDtTm.ISODateTime,                                   // "<CreDtTm>"
		Cretr:                  doc.Cretr.Party35Choice,                                   // "<Cretr>"
		Ctry:                   doc.Ctry.CountryCode,                                      // "<Ctry>"
		CxlRsnInf:              doc.CxlRsnInf.CancellationReasonInformation3,              // "<CxlRsnInf>"
		Dbtr:                   doc.Dbtr.PartyIdentification32,                            // "<Dbtr>"
		DbtrAgt:                doc.DbtrAgt.BranchAndFinancialInstitutionIdentification4,  // "<DbtrAgt>"
		// EndToEndId https://wiki.xmldation.com/Support/ISO20022/General_Rules/EndToEndId
		// EndToEndId https://www.jam-software.com/sepa-transfer/end-to-end-id.shtml
		// EndToEndId https://answers.sap.com/questions/12267089/element-endtoendid-not-filled-in-xml-payment-file.html
		// EndToEndId https://answers.sap.com/questions/10275743/dmee-%E2%80%93-endtoendid-with-paymantorder.html
		EndToEndId:             doc.EndToEndId.EndToEndId,                                 // "<EndToEndId>"
		Envlp:                  doc.Envlp.SupplementaryDataEnvelope1,                      // "<Envlp>"
		// FIId https://www.iso.org/iso-22000-food-safety-management.html
		// FIId https://www.qyriel.com/FullCatalogue/ISO_HEAD/out/ProtocolReport/xsd_head/head.001.001.01.xsd.html
		// FIId Financial Institution Identification
		// FIId AppHdr/Fr [Choice]
		FIId:                   doc.FIId.FinancialInstitutionIdentification,               // "<FIId>"
		// FinInstnId EPC limits the usage of Debtor Agent (DbtrAgt) and Creditor Agent CdtrAgt to allow only BIC and nothing else.
		// FinInstnId https://wiki.xmldation.com/Support/EPC/FinInstnId
		// FinInstnId https://wiki.xmldation.com/Support/RBS/CT_Rules/Global_Rules/CdtTrfTxInf%2F%2FCdtrAgt%2F%2FFinInstnId%2F%2FPstlAdr
		// FinInstnId CdtTrfTxInf/CdtrAgt/FinInstnId/PstlAdr Following fields from CreditorAgent / FinancialInstitutionIdentification / PostalAddress / Department '<CdtrAgt><FinInstnId><PstlAdr><Dept>'
		FinInstnId:             doc.FinInstnId.FinancialInstitutionIdentification7,        // "<FinInstnId>"
		FIToFICstmrCdtTrf:      doc.FIToFICstmrCdtTrf.FIToFICustomerCreditTransferV02,     // "<FIToFICstmrCdtTrf>"
		FIToFIPmtCxlReq:        doc.FIToFIPmtCxlReq.FIToFIPaymentCancellationRequestV01,   // "<FIToFIPmtCxlReq>"
		FIToFIPmtStsRpt:        doc.FIToFIPmtStsRpt.FIToFIPaymentStatusReportV03,          // "<FIToFIPmtStsRpt>"
		Fr:                     doc.Fr.Party9Choice,                                       // "<Fr>"
		GrpHdr:                 doc.GrpHdr.GroupHeader33,                                  // "<GrpHdr>"
		Id:                     doc.Id.Max35Text,                                          // "<Id>"
		InstdAgt:               doc.InstdAgt.BranchAndFinancialInstitutionIdentification4, // "<InstdAgt>"
		InstdAmt:               doc.InstdAmt.ActiveOrHistoricCurrencyAndAmount,            // "<InstdAmt>"
		// InstgAgt https://www.swift.com/swift-resource/248686/download
		// InstgAgt https://community.oracle.com/tech/developers/discussion/4327286/ora-00904-error-outer-join-19c
		// InstgAgt https://www.nacha.org/content/iso-20022-ach-mapping-guide
		// InstgAgt https://www.iso20022.org/sites/default/files/documents/D7/ISO20022_RTPG_pacs00800106_July_2017_v1_1.pdf
		InstgAgt:               doc.InstgAgt.BranchAndFinancialInstitutionIdentification4, // "<InstgAgt>"
		// InstrId https://wiki.xmldation.com/Support/ISO20022/General_Rules/InstrId
		// InstrId https://www.mathworks.com/help/instrument/instrid.html
		// InstrId https://wiki.xmldation.com/Support/Sampo/InstrId
		// InstrId https://docs.oracle.com/cd/E16582_01/doc.91/e15104/fields_sepa_pay_file_appx.htm#EOAEL01692
		InstrId:                doc.InstrId.InstructionIdentification,                     // "<InstrId>"
		// IntrBkSttlmAmt https://www.ecb.europa.eu/paym/groups/shared/docs/75299-tips-_cg_2017-09-28_presentation_udfs.pdf
		// IntrBkSttlmAmt https://wiki.xmldation.com/General_Information/ISO_20022/Difference_between_InstdAmt_and_EqvtAmt
		// IntrBkSttlmAmt https://www.iotafinance.com/en/SWIFT-ISO15022-Message-type-MT202-COV.html
		// IntrBkSttlmAmt https://www.bnymellon.com/content/dam/bnymellon/documents/pdf/iso-20022/Module%201_September%202020_Demystifying%20ISO20022.pdf
		IntrBkSttlmAmt:         doc.IntrBkSttlmAmt.ActiveOrHistoricCurrencyAndAmount,      // "<IntrBkSttlmAmt>"
		// IntrBkSttlmDt https://www.citibank.com/tts/sa/flippingbook/2021/ISO-20022-Citi-Mini-Series-and-Reference-Guide-Part-2/10/
		// IntrBkSttlmDt https://www.citibank.com/tts/sa/flippingbook/2021/ISO-20022-Citi-Mini-Series-and-Reference-Guide-Part-2/26/
		// IntrBkSttlmDt https://www.paymentstandards.ch/dam/mapping-rules_pacs008_esr.pdf
		// IntrBkSttlmDt https://www.payments.ca/sites/default/files/part_a_of_5_fitofi_customer_credit_transfers.pdf
		IntrBkSttlmDt:          doc.IntrBkSttlmDt.InterbankSettlementDate,                 // "<IntrBkSttlmDt>"
		Issr:                   doc.Issr.Issuer,                                           // "<Issr>"
		Justfn:                 doc.Justfn.CaseForwardingNotification3Code,                // "<Justfn>"
		KeyInfo:                doc.KeyInfo.KeyInfo,                                       // "<KeyInfo>"
		Mod:                    doc.Mod.RequestedModification7,                            // "<Mod>"
		MsgDefIdr:              doc.MsgDefIdr.MessageDefinitionIdentifier,                 // "<MsgDefIdr>"
		MsgId:                  doc.MsgId.MessageIdentification,                           // "<MsgId>"
		MssngOrIncrrctInf:      doc.MssngOrIncrrctInf.MissingOrIncorrectInformation3,      // "<MssngOrIncrrctInf>"
		// NbOfTxs https://wiki.xmldation.com/Support/RBS/DD_Rules/Global_Rules/NbOfTxs
		// NbOfTxs https://support.oracle.com/knowledge/Oracle%20E-Business%20Suite/1571592_1.html
		// NbOfTxs https://docs.oracle.com/cd/E16582_01/doc.91/e15104/fields_sepa_pay_file_appx.htm#EOAEL01692
		// NbOfTxs https://wiki.xmldation.com/Support/ISO20022/General_Rules/NbOfTxs
		NbOfTxs:                doc.NbOfTxs.Max15NumericText,                              // "<NbOfTxs>"
		NtfctnOfCaseAssgnmt:    doc.NtfctnOfCaseAssgnmt.NotificationOfCaseAssignmentV05,   // "<NtfctnOfCaseAssgnmt>"
		OrgnlCreDtTm:           doc.OrgnlCreDtTm.ISODateTime,                              // "<OrgnlCreDtTm>"
		// OrgnlEndToEndId https://wiki.xmldation.com/Support/ISO20022/General_Rules/EndToEndId
		// OrgnlEndToEndId https://paymentcomponents.atlassian.net/wiki/spaces/AH/pages/479428560/Sample+SEPA+messages+for+Testing
		// OrgnlEndToEndId https://answers.sap.com/questions/10275743/dmee-%E2%80%93-endtoendid-with-paymantorder.html
		// OrgnlEndToEndId https://blogs.sap.com/2021/07/30/pain.002-payment-rejections-processing-via-rfebka00/
		// OrgnlEndToEndId https://docs.crbcos.com/unicorncrb/docs/unicorn-output-files
		OrgnlEndToEndId:        doc.OrgnlEndToEndId.OriginalEndToEndIdentification,        // "<OrgnlEndToEndId>"
		// OrgnlGrpInf https://www.payments.ca/sites/default/files/part_c_of_5_payment_return.pdf
		// OrgnlGrpInf https://wiki.xmldation.com/Support/Nordea/CancellationRequest/Cancellation_Request_%2f%2f_CancellationReason2Code
		// OrgnlGrpInf https://www.iso20022.org/sites/default/files/documents/D7/Pacs004%20Real%20Time%20Payment%20Sep2018_v0.1.pdf
		// OrgnlGrpInf https://www.nacha.org/content/iso-20022-ach-mapping-guide
		// OrgnlGrpInf https://www.iso20022.org/sites/default/files/documents/D7/ISO20022_RTPG_pacs00200108_July_2017_v1_1.pdf
		OrgnlGrpInf:            doc.OrgnlGrpInf.OriginalGroupInformation3,                 // "<OrgnlGrpInf>"
		OrgnlGrpInfAndCxl:      doc.OrgnlGrpInfAndCxl.OriginalGroupInformation23,          // "<OrgnlGrpInfAndCxl>"
		OrgnlGrpInfAndSts:      doc.OrgnlGrpInfAndSts.OriginalGroupInformation20,          // "<OrgnlGrpInfAndSts>"
		OrgnlInstdAmt:          doc.OrgnlInstdAmt.OriginalInstructedAmount,                // "<OrgnlInstdAmt>"
		// OrgnlInstrId https://www.iso20022.org/sites/default/files/documents/D7/Pacs004%20Real%20Time%20Payment%20Sep2018_v0.1.pdf
		// OrgnlInstrId https://paymentcomponents.atlassian.net/wiki/spaces/AH/pages/479428560/Sample+SEPA+messages+for+Testing
		// OrgnlInstrId https://stackoverflow.com/questions/65199828/parsing-xml-in-c-sharp-with-xsd-file
		// OrgnlInstrId https://github.com/FasterXML/jackson-dataformat-xml/issues/217
		OrgnlInstrId:           doc.OrgnlInstrId.OriginalInstructionIdentification,        // "<OrgnlInstrId>"
		OrgnlIntrBkSttlmAmt:    doc.OrgnlIntrBkSttlmAmt.ActiveOrHistoricCurrencyAndAmount, // "<OrgnlIntrBkSttlmAmt>"
		OrgnlMsgId:             doc.OrgnlMsgId.OriginalMessageIdentification,              // "<OrgnlMsgId>"
		OrgnlMsgNmId:           doc.OrgnlMsgNmId.OriginalMessageNameIdentification,        // "<OrgnlMsgNmId>"
		OrgnlTxId:              doc.OrgnlTxId.OriginalTransactionIdentification,           // "<OrgnlTxId>"
		OrgnlTxRef:             doc.OrgnlTxRef.OriginalTransactionReference13,             // "<OrgnlTxRef>"
		Orgtr:                  doc.Orgtr.PartyIdentification32,                           // "<Orgtr>"
		PlcAndNm:               doc.PlcAndNm.PlcAndNm,                                     // "<PlcAndNm>"
		PmtTpInf:               doc.PmtTpInf.PaymentTypeInformation21,                     // "<PmtTpInf>"
		PstCd:                  doc.PstCd.Max16Text,                                       // "<PstCd>"
		PstlAdr:                doc.PstlAdr.PostalAddress6,                                // "<PstlAdr>"
		ReqToModfyPmt:          doc.ReqToModfyPmt.RequestToModifyPaymentV05,               // "<ReqToModfyPmt>"
		RsltnOfInvstgtn:        doc.RsltnOfInvstgtn.ResolutionOfInvestigationV08,          // "<RsltnOfInvstgtn>"
		RtrdInstdAmt:           doc.RtrdInstdAmt.ActiveOrHistoricCurrencyAndAmount,        // "<RtrdInstdAmt>"
		RtrdIntrBkSttlmAmt:     doc.RtrdIntrBkSttlmAmt.ActiveCurrencyAndAmount,            // "<RtrdIntrBkSttlmAmt>"
		RtrId:                  doc.RtrId.Max35Text,                                       // "<RtrId>"
		RtrRsnInf:              doc.RtrRsnInf.ReturnReasonInformation9,                    // "<RtrRsnInf>"
		Signature:              doc.Signature.Signature,                                   // "<Signature>"
		SignatureMethod:        doc.SignatureMethod.SignatureMethod,                       // "<SignatureMethod>"
		SplmtryData:            doc.SplmtryData.SupplementaryData1,                        // "<SplmtryData>"
		StrtNm:                 doc.StrtNm.Max70Text,                                      // "<StrtNm>"
		SttlmAcct:              doc.SttlmAcct.CashAccount16,                               // "<SttlmAcct>"
		SttlmInf:               doc.SttlmInf.SettlementInformation13,                      // "<SttlmInf>"
		SttlmMtd:               doc.SttlmMtd.SettlementMethod1Code,                        // "<SttlmMtd>"
		SvcLvl:                 doc.SvcLvl.ServiceLevel8Choice,                            // "<SvcLvl>"
		To:                     doc.To.Party9Choice,                                       // "<To>"
		TwnNm:                  doc.TwnNm.Max35Text,                                       // "<TwnNm>"
		TxId:                   doc.TxId.TransactionIdentification,                        // "<TxId>"
		TxInfAndSts:            doc.TxInfAndSts.PaymentTransactionInformation26,           // "<TxInfAndSts>"
		TxSts:                  doc.TxSts.TransactionIndividualStatus3Code,                // "<TxSts>"
		UblToApply:             doc.UblToApply.UnableToApplyV07,                           // "<UblToApply>"
		UltmtCdtr:              doc.UltmtCdtr.PartyIdentification32,                       // "<UltmtCdtr>"
		Undrlyg:                doc.Undrlyg.UnderlyingTransaction4Choice,                  // "<Undrlyg>"
		X509Data:               doc.X509Data.KeyInfoX509Data,                              // "<X509Data>"
		XchgRate:               doc.XchgRate.BaseOneRate,                                  // "<XchgRate>"
	}
```