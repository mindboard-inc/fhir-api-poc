package hl73

import (
	"encoding/xml"
	"reflect"
	"fmt"
)

type ClinicalDocument struct {
	XMLName        xml.Name `xml:"ClinicalDocument"`
	Text           string   `xml:",chardata"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Cda            string   `xml:"cda,attr"`
	Sdtc           string   `xml:"sdtc,attr"`
	RealmCode      struct {
		Text string `xml:",chardata"`
		Code string `xml:"code,attr"`
	} `xml:"realmCode"`
	TypeId struct {
		Text      string `xml:",chardata"`
		Root      string `xml:"root,attr"`
		Extension string `xml:"extension,attr"`
	} `xml:"typeId"`
	TemplateId []struct {
		Text      string `xml:",chardata"`
		Root      string `xml:"root,attr"`
		Extension string `xml:"extension,attr"`
	} `xml:"templateId"`
	ID struct {
		Text string `xml:",chardata"`
		Root string `xml:"root,attr"`
	} `xml:"id"`
	Code struct {
		Text           string `xml:",chardata"`
		CodeSystem     string `xml:"codeSystem,attr"`
		CodeSystemName string `xml:"codeSystemName,attr"`
		Code           string `xml:"code,attr"`
		DisplayName    string `xml:"displayName,attr"`
	} `xml:"code"`
	Title         string `xml:"title"`
	EffectiveTime struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"effectiveTime"`
	ConfidentialityCode struct {
		Text       string `xml:",chardata"`
		Code       string `xml:"code,attr"`
		CodeSystem string `xml:"codeSystem,attr"`
	} `xml:"confidentialityCode"`
	LanguageCode struct {
		Text string `xml:",chardata"`
		Code string `xml:"code,attr"`
	} `xml:"languageCode"`
	SetId struct {
		Text string `xml:",chardata"`
		Root string `xml:"root,attr"`
	} `xml:"setId"`
	VersionNumber struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"versionNumber"`
	RecordTarget struct {
		Text        string `xml:",chardata"`
		PatientRole struct {
			Text string `xml:",chardata"`
			ID   []struct {
				Text      string `xml:",chardata"`
				Root      string `xml:"root,attr"`
				Extension string `xml:"extension,attr"`
			} `xml:"id"`
			Addr struct {
				Text              string `xml:",chardata"`
				Use               string `xml:"use,attr"`
				StreetAddressLine string `xml:"streetAddressLine"`
				City              string `xml:"city"`
				State             string `xml:"state"`
				PostalCode        string `xml:"postalCode"`
				Country           string `xml:"country"`
			} `xml:"addr"`
			Telecom struct {
				Text  string `xml:",chardata"`
				Use   string `xml:"use,attr"`
				Value string `xml:"value,attr"`
			} `xml:"telecom"`
			Patient struct {
				Text string `xml:",chardata"`
				Name struct {
					Text   string   `xml:",chardata"`
					Given  []string `xml:"given"`
					Family string   `xml:"family"`
					Suffix string   `xml:"suffix"`
				} `xml:"name"`
				AdministrativeGenderCode struct {
					Text        string `xml:",chardata"`
					Code        string `xml:"code,attr"`
					CodeSystem  string `xml:"codeSystem,attr"`
					DisplayName string `xml:"displayName,attr"`
				} `xml:"administrativeGenderCode"`
				BirthTime struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"birthTime"`
				MaritalStatusCode struct {
					Text           string `xml:",chardata"`
					Code           string `xml:"code,attr"`
					DisplayName    string `xml:"displayName,attr"`
					CodeSystem     string `xml:"codeSystem,attr"`
					CodeSystemName string `xml:"codeSystemName,attr"`
				} `xml:"maritalStatusCode"`
				RaceCode struct {
					Text           string `xml:",chardata"`
					Code           string `xml:"code,attr"`
					DisplayName    string `xml:"displayName,attr"`
					CodeSystem     string `xml:"codeSystem,attr"`
					CodeSystemName string `xml:"codeSystemName,attr"`
				} `xml:"raceCode"`
				EthnicGroupCode struct {
					Text           string `xml:",chardata"`
					Code           string `xml:"code,attr"`
					DisplayName    string `xml:"displayName,attr"`
					CodeSystem     string `xml:"codeSystem,attr"`
					CodeSystemName string `xml:"codeSystemName,attr"`
				} `xml:"ethnicGroupCode"`
				LanguageCommunication struct {
					Text         string `xml:",chardata"`
					LanguageCode struct {
						Text       string `xml:",chardata"`
						NullFlavor string `xml:"nullFlavor,attr"`
					} `xml:"languageCode"`
				} `xml:"languageCommunication"`
			} `xml:"patient"`
		} `xml:"patientRole"`
	} `xml:"recordTarget"`
	Author struct {
		Text string `xml:",chardata"`
		Time struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"time"`
		AssignedAuthor struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text      string `xml:",chardata"`
				Extension string `xml:"extension,attr"`
				Root      string `xml:"root,attr"`
			} `xml:"id"`
			Addr struct {
				Text              string `xml:",chardata"`
				StreetAddressLine string `xml:"streetAddressLine"`
				City              string `xml:"city"`
				State             string `xml:"state"`
				PostalCode        string `xml:"postalCode"`
				Country           string `xml:"country"`
			} `xml:"addr"`
			Telecom struct {
				Text  string `xml:",chardata"`
				Use   string `xml:"use,attr"`
				Value string `xml:"value,attr"`
			} `xml:"telecom"`
			AssignedAuthoringDevice struct {
				Text                  string `xml:",chardata"`
				ManufacturerModelName string `xml:"manufacturerModelName"`
				SoftwareName          string `xml:"softwareName"`
			} `xml:"assignedAuthoringDevice"`
		} `xml:"assignedAuthor"`
	} `xml:"author"`
	Custodian struct {
		Text              string `xml:",chardata"`
		AssignedCustodian struct {
			Text                             string `xml:",chardata"`
			RepresentedCustodianOrganization struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text      string `xml:",chardata"`
					Extension string `xml:"extension,attr"`
					Root      string `xml:"root,attr"`
				} `xml:"id"`
				Name    string `xml:"name"`
				Telecom struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
					Use   string `xml:"use,attr"`
				} `xml:"telecom"`
				Addr struct {
					Text              string `xml:",chardata"`
					Use               string `xml:"use,attr"`
					StreetAddressLine string `xml:"streetAddressLine"`
					City              string `xml:"city"`
					State             string `xml:"state"`
					PostalCode        string `xml:"postalCode"`
					Country           string `xml:"country"`
				} `xml:"addr"`
			} `xml:"representedCustodianOrganization"`
		} `xml:"assignedCustodian"`
	} `xml:"custodian"`
	DocumentationOf struct {
		Text         string `xml:",chardata"`
		TypeCode     string `xml:"typeCode,attr"`
		ServiceEvent struct {
			Text          string `xml:",chardata"`
			ClassCode     string `xml:"classCode,attr"`
			EffectiveTime struct {
				Text string `xml:",chardata"`
				Low  struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"low"`
			} `xml:"effectiveTime"`
			Performer []struct {
				Text     string `xml:",chardata"`
				TypeCode string `xml:"typeCode,attr"`
				Time     struct {
					Text string `xml:",chardata"`
					Low  struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"low"`
				} `xml:"time"`
				AssignedEntity struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text      string `xml:",chardata"`
						Root      string `xml:"root,attr"`
						Extension string `xml:"extension,attr"`
					} `xml:"id"`
					Code struct {
						Text           string `xml:",chardata"`
						Code           string `xml:"code,attr"`
						DisplayName    string `xml:"displayName,attr"`
						CodeSystem     string `xml:"codeSystem,attr"`
						CodeSystemName string `xml:"codeSystemName,attr"`
					} `xml:"code"`
				} `xml:"assignedEntity"`
				FunctionCode struct {
					Text           string `xml:",chardata"`
					Code           string `xml:"code,attr"`
					DisplayName    string `xml:"displayName,attr"`
					CodeSystem     string `xml:"codeSystem,attr"`
					CodeSystemName string `xml:"codeSystemName,attr"`
				} `xml:"functionCode"`
			} `xml:"performer"`
		} `xml:"serviceEvent"`
	} `xml:"documentationOf"`
	ComponentOf struct {
		Text                  string `xml:",chardata"`
		EncompassingEncounter struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
				Root string `xml:"root,attr"`
			} `xml:"id"`
			Code struct {
				Text           string `xml:",chardata"`
				Code           string `xml:"code,attr"`
				CodeSystem     string `xml:"codeSystem,attr"`
				DisplayName    string `xml:"displayName,attr"`
				CodeSystemName string `xml:"codeSystemName,attr"`
			} `xml:"code"`
			EffectiveTime struct {
				Text string `xml:",chardata"`
				Low  struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"low"`
				High struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value,attr"`
				} `xml:"high"`
			} `xml:"effectiveTime"`
			DischargeDispositionCode struct {
				Text           string `xml:",chardata"`
				Code           string `xml:"code,attr"`
				DisplayName    string `xml:"displayName,attr"`
				CodeSystem     string `xml:"codeSystem,attr"`
				CodeSystemName string `xml:"codeSystemName,attr"`
			} `xml:"dischargeDispositionCode"`
		} `xml:"encompassingEncounter"`
	} `xml:"componentOf"`
	Component struct {
		Text           string `xml:",chardata"`
		StructuredBody struct {
			Text      string `xml:",chardata"`
			Component []struct {
				Text    string `xml:",chardata"`
				Section struct {
					Chardata   string `xml:",chardata"`
					TemplateId []struct {
						Text      string `xml:",chardata"`
						Root      string `xml:"root,attr"`
						Extension string `xml:"extension,attr"`
					} `xml:"templateId"`
					Code struct {
						Text           string `xml:",chardata"`
						Code           string `xml:"code,attr"`
						CodeSystem     string `xml:"codeSystem,attr"`
						DisplayName    string `xml:"displayName,attr"`
						CodeSystemName string `xml:"codeSystemName,attr"`
					} `xml:"code"`
					Title string `xml:"title"`
					Text  struct {
						Text string `xml:",chardata"`
						List []struct {
							Text string `xml:",chardata"`
							Item []struct {
								Text string `xml:",chardata"`
								List struct {
									Text string   `xml:",chardata"`
									Item []string `xml:"item"`
								} `xml:"list"`
							} `xml:"item"`
						} `xml:"list"`
						Table []struct {
							Text   string `xml:",chardata"`
							Border string `xml:"border,attr"`
							Width  string `xml:"width,attr"`
							Thead  struct {
								Text string `xml:",chardata"`
								Tr   struct {
									Text string   `xml:",chardata"`
									Th   []string `xml:"th"`
								} `xml:"tr"`
							} `xml:"thead"`
							Tbody struct {
								Text string `xml:",chardata"`
								Tr   []struct {
									Text string   `xml:",chardata"`
									Td   []string `xml:"td"`
									Th   struct {
										Text  string `xml:",chardata"`
										Align string `xml:"align,attr"`
									} `xml:"th"`
								} `xml:"tr"`
							} `xml:"tbody"`
						} `xml:"table"`
						Paragraph string `xml:"paragraph"`
					} `xml:"text"`
					Entry []struct {
						Text        string `xml:",chardata"`
						TypeCode    string `xml:"typeCode,attr"`
						Observation struct {
							Text       string `xml:",chardata"`
							ClassCode  string `xml:"classCode,attr"`
							MoodCode   string `xml:"moodCode,attr"`
							TemplateId []struct {
								Text      string `xml:",chardata"`
								Root      string `xml:"root,attr"`
								Extension string `xml:"extension,attr"`
							} `xml:"templateId"`
							ID struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
							Code struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								DisplayName    string `xml:"displayName,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
								Translation    struct {
									Text           string `xml:",chardata"`
									Code           string `xml:"code,attr"`
									DisplayName    string `xml:"displayName,attr"`
									CodeSystemName string `xml:"codeSystemName,attr"`
									CodeSystem     string `xml:"codeSystem,attr"`
								} `xml:"translation"`
							} `xml:"code"`
							StatusCode struct {
								Text string `xml:",chardata"`
								Code string `xml:"code,attr"`
							} `xml:"statusCode"`
							Value struct {
								Text           string `xml:",chardata"`
								Type           string `xml:"type,attr"`
								Code           string `xml:"code,attr"`
								DisplayName    string `xml:"displayName,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
								NullFlavor     string `xml:"nullFlavor,attr"`
								Value          string `xml:"value,attr"`
							} `xml:"value"`
							EffectiveTime struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value,attr"`
								Low   struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"low"`
							} `xml:"effectiveTime"`
							EntryRelationship struct {
								Text        string `xml:",chardata"`
								TypeCode    string `xml:"typeCode,attr"`
								Observation struct {
									Text       string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									MoodCode   string `xml:"moodCode,attr"`
									TemplateId struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"templateId"`
									ID struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"id"`
									Code struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										DisplayName    string `xml:"displayName,attr"`
									} `xml:"code"`
									StatusCode struct {
										Text string `xml:",chardata"`
										Code string `xml:"code,attr"`
									} `xml:"statusCode"`
									Value struct {
										Text         string `xml:",chardata"`
										Type         string `xml:"type,attr"`
										NullFlavor   string `xml:"nullFlavor,attr"`
										OriginalText string `xml:"originalText"`
									} `xml:"value"`
								} `xml:"observation"`
							} `xml:"entryRelationship"`
						} `xml:"observation"`
						Encounter struct {
							Text       string `xml:",chardata"`
							ClassCode  string `xml:"classCode,attr"`
							MoodCode   string `xml:"moodCode,attr"`
							TemplateId []struct {
								Text      string `xml:",chardata"`
								Root      string `xml:"root,attr"`
								Extension string `xml:"extension,attr"`
							} `xml:"templateId"`
							ID struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
							Code struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								DisplayName    string `xml:"displayName,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
							} `xml:"code"`
							EffectiveTime struct {
								Text string `xml:",chardata"`
								Low  struct {
									Text       string `xml:",chardata"`
									Value      string `xml:"value,attr"`
									NullFlavor string `xml:"nullFlavor,attr"`
								} `xml:"low"`
								High struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"high"`
							} `xml:"effectiveTime"`
							EntryRelationship []struct {
								Text        string `xml:",chardata"`
								TypeCode    string `xml:"typeCode,attr"`
								Observation struct {
									Text       string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									MoodCode   string `xml:"moodCode,attr"`
									TemplateId struct {
										Text      string `xml:",chardata"`
										Root      string `xml:"root,attr"`
										Extension string `xml:"extension,attr"`
									} `xml:"templateId"`
									ID struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"id"`
									Code struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										DisplayName    string `xml:"displayName,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
									} `xml:"code"`
									StatusCode struct {
										Text string `xml:",chardata"`
										Code string `xml:"code,attr"`
									} `xml:"statusCode"`
									Value struct {
										Text           string `xml:",chardata"`
										Type           string `xml:"type,attr"`
										Code           string `xml:"code,attr"`
										DisplayName    string `xml:"displayName,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										Value          string `xml:"value,attr"`
									} `xml:"value"`
								} `xml:"observation"`
								Act struct {
									Text       string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									MoodCode   string `xml:"moodCode,attr"`
									TemplateId struct {
										Text      string `xml:",chardata"`
										Root      string `xml:"root,attr"`
										Extension string `xml:"extension,attr"`
									} `xml:"templateId"`
									ID struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"id"`
									Code struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										DisplayName    string `xml:"displayName,attr"`
									} `xml:"code"`
									StatusCode struct {
										Text string `xml:",chardata"`
										Code string `xml:"code,attr"`
									} `xml:"statusCode"`
									EffectiveTime struct {
										Text string `xml:",chardata"`
										Low  struct {
											Text  string `xml:",chardata"`
											Value string `xml:"value,attr"`
										} `xml:"low"`
									} `xml:"effectiveTime"`
									EntryRelationship struct {
										Text         string `xml:",chardata"`
										TypeCode     string `xml:"typeCode,attr"`
										InversionInd string `xml:"inversionInd,attr"`
										Observation  struct {
											Text       string `xml:",chardata"`
											ClassCode  string `xml:"classCode,attr"`
											MoodCode   string `xml:"moodCode,attr"`
											TemplateId struct {
												Text      string `xml:",chardata"`
												Root      string `xml:"root,attr"`
												Extension string `xml:"extension,attr"`
											} `xml:"templateId"`
											ID struct {
												Text string `xml:",chardata"`
												Root string `xml:"root,attr"`
											} `xml:"id"`
											Code struct {
												Text           string `xml:",chardata"`
												Code           string `xml:"code,attr"`
												CodeSystem     string `xml:"codeSystem,attr"`
												CodeSystemName string `xml:"codeSystemName,attr"`
												DisplayName    string `xml:"displayName,attr"`
											} `xml:"code"`
											StatusCode struct {
												Text string `xml:",chardata"`
												Code string `xml:"code,attr"`
											} `xml:"statusCode"`
											EffectiveTime struct {
												Text string `xml:",chardata"`
												Low  struct {
													Text  string `xml:",chardata"`
													Value string `xml:"value,attr"`
												} `xml:"low"`
											} `xml:"effectiveTime"`
											Value struct {
												Text           string `xml:",chardata"`
												Type           string `xml:"type,attr"`
												Code           string `xml:"code,attr"`
												CodeSystemName string `xml:"codeSystemName,attr"`
												CodeSystem     string `xml:"codeSystem,attr"`
												DisplayName    string `xml:"displayName,attr"`
											} `xml:"value"`
										} `xml:"observation"`
									} `xml:"entryRelationship"`
								} `xml:"act"`
							} `xml:"entryRelationship"`
							Reference []struct {
								Text             string `xml:",chardata"`
								TypeCode         string `xml:"typeCode,attr"`
								ExternalDocument struct {
									Chardata   string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									MoodCode   string `xml:"moodCode,attr"`
									TemplateId struct {
										Text      string `xml:",chardata"`
										Root      string `xml:"root,attr"`
										Extension string `xml:"extension,attr"`
									} `xml:"templateId"`
									ID struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"id"`
									Code struct {
										Text           string `xml:",chardata"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										Code           string `xml:"code,attr"`
										DisplayName    string `xml:"displayName,attr"`
									} `xml:"code"`
									Text struct {
										Text      string `xml:",chardata"`
										MediaType string `xml:"mediaType,attr"`
										Reference struct {
											Text  string `xml:",chardata"`
											Value string `xml:"value,attr"`
										} `xml:"reference"`
									} `xml:"text"`
									SetId struct {
										Text      string `xml:",chardata"`
										Extension string `xml:"extension,attr"`
										Root      string `xml:"root,attr"`
									} `xml:"setId"`
									VersionNumber struct {
										Text  string `xml:",chardata"`
										Value string `xml:"value,attr"`
									} `xml:"versionNumber"`
								} `xml:"externalDocument"`
							} `xml:"reference"`
							DischargeDispositionCode struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								DisplayName    string `xml:"displayName,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
							} `xml:"dischargeDispositionCode"`
							Participant []struct {
								Text            string `xml:",chardata"`
								TypeCode        string `xml:"typeCode,attr"`
								ParticipantRole struct {
									Text       string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									TemplateId struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"templateId"`
									Code struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										DisplayName    string `xml:"displayName,attr"`
									} `xml:"code"`
									ID struct {
										Text       string `xml:",chardata"`
										NullFlavor string `xml:"nullFlavor,attr"`
									} `xml:"id"`
								} `xml:"participantRole"`
							} `xml:"participant"`
						} `xml:"encounter"`
						Act struct {
							Text       string `xml:",chardata"`
							ClassCode  string `xml:"classCode,attr"`
							MoodCode   string `xml:"moodCode,attr"`
							TemplateId []struct {
								Text      string `xml:",chardata"`
								Root      string `xml:"root,attr"`
								Extension string `xml:"extension,attr"`
							} `xml:"templateId"`
							ID struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
							Code struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
								DisplayName    string `xml:"displayName,attr"`
								Translation    struct {
									Text           string `xml:",chardata"`
									Code           string `xml:"code,attr"`
									DisplayName    string `xml:"displayName,attr"`
									CodeSystemName string `xml:"codeSystemName,attr"`
									CodeSystem     string `xml:"codeSystem,attr"`
								} `xml:"translation"`
							} `xml:"code"`
							StatusCode struct {
								Text string `xml:",chardata"`
								Code string `xml:"code,attr"`
							} `xml:"statusCode"`
							EntryRelationship struct {
								Text     string `xml:",chardata"`
								TypeCode string `xml:"typeCode,attr"`
								Act      struct {
									Text      string `xml:",chardata"`
									ClassCode string `xml:"classCode,attr"`
									MoodCode  string `xml:"moodCode,attr"`
									ID        struct {
										Text       string `xml:",chardata"`
										NullFlavor string `xml:"nullFlavor,attr"`
									} `xml:"id"`
									Code struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										DisplayName    string `xml:"displayName,attr"`
									} `xml:"code"`
									StatusCode struct {
										Text string `xml:",chardata"`
										Code string `xml:"code,attr"`
									} `xml:"statusCode"`
								} `xml:"act"`
							} `xml:"entryRelationship"`
							EffectiveTime struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value,attr"`
							} `xml:"effectiveTime"`
						} `xml:"act"`
						SubstanceAdministration struct {
							Text       string `xml:",chardata"`
							ClassCode  string `xml:"classCode,attr"`
							MoodCode   string `xml:"moodCode,attr"`
							TemplateId struct {
								Text      string `xml:",chardata"`
								Root      string `xml:"root,attr"`
								Extension string `xml:"extension,attr"`
							} `xml:"templateId"`
							ID struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
							StatusCode struct {
								Text string `xml:",chardata"`
								Code string `xml:"code,attr"`
							} `xml:"statusCode"`
							EffectiveTime struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Low  struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"low"`
							} `xml:"effectiveTime"`
							ApproachSiteCode struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
								DisplayName    string `xml:"displayName,attr"`
								Qualifier      []struct {
									Text string `xml:",chardata"`
									Name struct {
										Text        string `xml:",chardata"`
										Code        string `xml:"code,attr"`
										CodeSystem  string `xml:"codeSystem,attr"`
										DisplayName string `xml:"displayName,attr"`
									} `xml:"name"`
									Value struct {
										Text        string `xml:",chardata"`
										Code        string `xml:"code,attr"`
										CodeSystem  string `xml:"codeSystem,attr"`
										DisplayName string `xml:"displayName,attr"`
									} `xml:"value"`
								} `xml:"qualifier"`
							} `xml:"approachSiteCode"`
							DoseQuantity struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value,attr"`
								Unit  string `xml:"unit,attr"`
							} `xml:"doseQuantity"`
							Consumable struct {
								Text                string `xml:",chardata"`
								ManufacturedProduct struct {
									Text       string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									TemplateId struct {
										Text      string `xml:",chardata"`
										Root      string `xml:"root,attr"`
										Extension string `xml:"extension,attr"`
									} `xml:"templateId"`
									ID struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"id"`
									ManufacturedMaterial struct {
										Text string `xml:",chardata"`
										Code struct {
											Text           string `xml:",chardata"`
											Code           string `xml:"code,attr"`
											DisplayName    string `xml:"displayName,attr"`
											CodeSystem     string `xml:"codeSystem,attr"`
											CodeSystemName string `xml:"codeSystemName,attr"`
										} `xml:"code"`
										LotNumberText struct {
											Text       string `xml:",chardata"`
											NullFlavor string `xml:"nullFlavor,attr"`
										} `xml:"lotNumberText"`
									} `xml:"manufacturedMaterial"`
								} `xml:"manufacturedProduct"`
							} `xml:"consumable"`
							RepeatNumber struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value,attr"`
							} `xml:"repeatNumber"`
						} `xml:"substanceAdministration"`
						Organizer struct {
							Text       string `xml:",chardata"`
							ClassCode  string `xml:"classCode,attr"`
							MoodCode   string `xml:"moodCode,attr"`
							TemplateId struct {
								Text      string `xml:",chardata"`
								Root      string `xml:"root,attr"`
								Extension string `xml:"extension,attr"`
							} `xml:"templateId"`
							ID struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
							Code struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
								DisplayName    string `xml:"displayName,attr"`
							} `xml:"code"`
							StatusCode struct {
								Text string `xml:",chardata"`
								Code string `xml:"code,attr"`
							} `xml:"statusCode"`
							Component []struct {
								Text        string `xml:",chardata"`
								Observation struct {
									Text       string `xml:",chardata"`
									ClassCode  string `xml:"classCode,attr"`
									MoodCode   string `xml:"moodCode,attr"`
									TemplateId struct {
										Text      string `xml:",chardata"`
										Root      string `xml:"root,attr"`
										Extension string `xml:"extension,attr"`
									} `xml:"templateId"`
									ID struct {
										Text string `xml:",chardata"`
										Root string `xml:"root,attr"`
									} `xml:"id"`
									Code struct {
										Text           string `xml:",chardata"`
										Code           string `xml:"code,attr"`
										CodeSystem     string `xml:"codeSystem,attr"`
										CodeSystemName string `xml:"codeSystemName,attr"`
										DisplayName    string `xml:"displayName,attr"`
									} `xml:"code"`
									StatusCode struct {
										Text string `xml:",chardata"`
										Code string `xml:"code,attr"`
									} `xml:"statusCode"`
									EffectiveTime struct {
										Text  string `xml:",chardata"`
										Value string `xml:"value,attr"`
									} `xml:"effectiveTime"`
									Value struct {
										Text  string `xml:",chardata"`
										Type  string `xml:"type,attr"`
										Value string `xml:"value,attr"`
										Unit  string `xml:"unit,attr"`
									} `xml:"value"`
								} `xml:"observation"`
							} `xml:"component"`
							EffectiveTime struct {
								Text string `xml:",chardata"`
								Low  struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"low"`
								High struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value,attr"`
								} `xml:"high"`
							} `xml:"effectiveTime"`
						} `xml:"organizer"`
						Procedure struct {
							Text       string `xml:",chardata"`
							ClassCode  string `xml:"classCode,attr"`
							MoodCode   string `xml:"moodCode,attr"`
							TemplateId []struct {
								Text      string `xml:",chardata"`
								Root      string `xml:"root,attr"`
								Extension string `xml:"extension,attr"`
							} `xml:"templateId"`
							ID struct {
								Text string `xml:",chardata"`
								Root string `xml:"root,attr"`
							} `xml:"id"`
							Code struct {
								Text           string `xml:",chardata"`
								Code           string `xml:"code,attr"`
								DisplayName    string `xml:"displayName,attr"`
								CodeSystem     string `xml:"codeSystem,attr"`
								CodeSystemName string `xml:"codeSystemName,attr"`
								Translation    struct {
									Text           string `xml:",chardata"`
									Code           string `xml:"code,attr"`
									DisplayName    string `xml:"displayName,attr"`
									CodeSystemName string `xml:"codeSystemName,attr"`
									CodeSystem     string `xml:"codeSystem,attr"`
								} `xml:"translation"`
							} `xml:"code"`
							StatusCode struct {
								Text string `xml:",chardata"`
								Code string `xml:"code,attr"`
							} `xml:"statusCode"`
						} `xml:"procedure"`
					} `xml:"entry"`
				} `xml:"section"`
			} `xml:"component"`
		} `xml:"structuredBody"`
	} `xml:"component"`
} 

func (f *ClinicalDocument) Reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		//valueField := val.Field(i)
		typeField := val.Type().Field(i)
		f := val.Field(i)
		fmt.Printf("Field Name: %s, %s Value: %v\n", typeField.Name, f.Type(), f.Interface())
	}


}
