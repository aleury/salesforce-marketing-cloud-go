package soap

type APIProperty struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}
