package soap

type ClientID struct {
	// Deprecated.  Use ID.
	ClientID int32 `xml:"ClientID,omitempty" json:"ClientID,omitempty"`

	ID int32 `xml:"ID,omitempty" json:"ID,omitempty"`

	PartnerClientKey string `xml:"PartnerClientKey,omitempty" json:"PartnerClientKey,omitempty"`

	UserID int32 `xml:"UserID,omitempty" json:"UserID,omitempty"`

	PartnerUserKey string `xml:"PartnerUserKey,omitempty" json:"PartnerUserKey,omitempty"`

	CreatedBy int32 `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty"`

	ModifiedBy int32 `xml:"ModifiedBy,omitempty" json:"ModifiedBy,omitempty"`

	EnterpriseID int64 `xml:"EnterpriseID,omitempty" json:"EnterpriseID,omitempty"`

	CustomerKey string `xml:"CustomerKey,omitempty" json:"CustomerKey,omitempty"`

	CustomerID string `xml:"CustomerID,omitempty" json:"CustomerID,omitempty"`
}
