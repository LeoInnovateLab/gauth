package linkedin

type Scope struct {
	scope       string
	description string
	isDefault   bool
}

func (a *Scope) GetScope() string {
	return a.scope
}

func (a *Scope) IsDefault() bool {
	return a.isDefault
}

func (a *Scope) GetDefaultScopes() []string {
	scopes := []Scope{
		RLiteProfile,
		REmailAddress,
		WMemberSocial,
		RMemberSocial,
		RAdCampaigns,
		RAds,
		RAdsLeadgenAutomation,
		RAdsReporting,
		RBasicProfile,
		ROrganizationSocial,
		RWAdCampaigns,
		RWAds,
		RWCompanyAdmin,
		RWDmpSegments,
		RWOrganizationAdmin,
		RWOrganization,
		WOrganizationSocial,
		WShare,
	}
	var result []string
	for _, scope := range scopes {
		if scope.isDefault {
			result = append(result, scope.scope)
		}
	}

	return result
}

var (
	RLiteProfile = Scope{
		scope:       "r_liteprofile",
		description: "Use your name, headline, and photo",
		isDefault:   true,
	}
	REmailAddress = Scope{
		scope:       "r_emailaddress",
		description: "Use the primary email address associated with your LinkedIn account",
		isDefault:   true,
	}
	WMemberSocial = Scope{
		scope:       "w_member_social",
		description: "Post, comment and like posts on your behalf",
		isDefault:   true,
	}
	RMemberSocial = Scope{
		scope:       "r_member_social",
		description: "Retrieve your posts, comments, likes, and other engagement data",
		isDefault:   false,
	}
	RAdCampaigns = Scope{
		scope:       "r_ad_campaigns",
		description: "View advertising campaigns you manage",
		isDefault:   false,
	}
	RAds = Scope{
		scope:       "r_ads",
		description: "Retrieve your advertising accounts",
		isDefault:   false,
	}
	RAdsLeadgenAutomation = Scope{
		scope:       "r_ads_leadgen_automation",
		description: "Access your Lead Gen Forms and retrieve leads",
		isDefault:   false,
	}
	RAdsReporting = Scope{
		scope:       "r_ads_reporting",
		description: "Retrieve reporting for your advertising accounts",
		isDefault:   false,
	}
	RBasicProfile = Scope{
		scope:       "r_basicprofile",
		description: "Use your basic profile including your name, photo, headline, and current positions",
		isDefault:   false,
	}
	ROrganizationSocial = Scope{
		scope:       "r_organization_social",
		description: "Retrieve your organizations' posts, including any comments, likes and other engagement data",
		isDefault:   false,
	}
	RWAdCampaigns = Scope{
		scope:       "rw_ad_campaigns",
		description: "Manage your advertising campaigns",
		isDefault:   false,
	}
	RWAds = Scope{
		scope:       "rw_ads",
		description: "Manage your advertising accounts",
		isDefault:   false,
	}
	RWCompanyAdmin = Scope{
		scope:       "rw_company_admin",
		description: "For V1 calls: Manage your organization's page and post updates",
		isDefault:   false,
	}
	RWDmpSegments = Scope{
		scope:       "rw_dmp_segments",
		description: "Create and manage your matched audiences",
		isDefault:   false,
	}
	RWOrganizationAdmin = Scope{
		scope:       "rw_organization_admin",
		description: "Manage your organizations' pages and retrieve reporting data",
		isDefault:   false,
	}
	RWOrganization = Scope{
		scope:       "rw_organization",
		description: "For V2 calls: Manage your organization's page and post updates",
		isDefault:   false,
	}
	WOrganizationSocial = Scope{
		scope:       "w_organization_social",
		description: "Post, comment and like posts on your organization's behalf",
		isDefault:   false,
	}
	WShare = Scope{
		scope:       "w_share",
		description: "Post updates to LinkedIn as you",
		isDefault:   false,
	}
)
