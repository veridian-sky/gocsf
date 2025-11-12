package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Remediation The Remediation object describes the recommended remediation steps to address
// identified issue(s).
type Remediation struct {
	// An array of Center for Internet Security (CIS) Controls that can be
	// optionally mapped to provide additional remediation details.
	CisControls []*CisControl `json:"cis_controls,omitempty"`

	// The description of the remediation strategy.
	Desc string `json:"desc,omitempty"`

	// A list of KB articles or patches related to an endpoint. A KB Article
	// contains metadata that describes the patch or an update.
	KbArticleList []*KbArticle `json:"kb_article_list,omitempty"`

	// The KB article/s related to the entity. A KB Article contains metadata that
	// describes the patch or an update.
	KbArticles []string `json:"kb_articles,omitempty"`

	// A list of supporting URL/s, references that help describe the remediation
	// strategy.
	References []string `json:"references,omitempty"`
}
