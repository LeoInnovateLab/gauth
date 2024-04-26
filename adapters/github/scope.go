package github

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
		RepoStatus,
		RepoDeployment,
		PublicRepo,
		RepoInvite,
		SecurityEvents,
		WriteOrg,
		Workflow,
		WritePackages,
		WriteGPGKey,
		WritePublicKey,
		ReadOrg,
		ReadPackages,
		ReadGPGKey,
		ReadDiscussion,
		ReadRepoHook,
		ReadPublicKey,
		WriteRepoHook,
		AdminOrg,
		AdminPublicKey,
		UserEmail,
		UserFollow,
		Gist,
		Notifications,
		User,
		DeleteRepo,
		DeletePackages,
		AdminGPGKey,
		WriteDiscussion,
		ReadUser,
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
	RepoStatus = Scope{
		scope:       "repo:status",
		description: "Grants read/write access to public and private repository commit statuses. This scope is only necessary to grant other users or services access to private repository commit statuses without granting access to the code.",
		isDefault:   false,
	}
	RepoDeployment = Scope{
		scope:       "repo_deployment",
		description: "Grants access to deployment statuses for public and private repositories. This scope is only necessary to grant other users or services access to deployment statuses, <em>without</em> granting access to the code.",
		isDefault:   false,
	}
	PublicRepo = Scope{
		scope:       "public_repo",
		description: "Limits access to public repositories. That includes read/write access to code, commit statuses, repository projects, collaborators, and deployment statuses for public repositories and organizations. Also required for starring public repositories.",
		isDefault:   false,
	}
	RepoInvite = Scope{
		scope:       "repo:invite",
		description: "Grants accept/decline abilities for invitations to collaborate on a repository. This scope is only necessary to grant other users or services access to invites <em>without</em> granting access to the code.",
		isDefault:   false,
	}
	SecurityEvents = Scope{
		scope:       "security_events",
		description: "Grants read and write access to security events in the code scanning API.",
		isDefault:   false,
	}
	WriteOrg = Scope{
		scope:       "write:org",
		description: "Read and write access to organization membership, organization projects, and team membership.",
		isDefault:   false,
	}
	WriteDiscussion = Scope{
		scope:       "write:discussion",
		description: "Allows read and write access for team discussions.",
		isDefault:   false,
	}
	Workflow = Scope{
		scope:       "workflow",
		description: "Grants the ability to add and update GitHub Actions workflow files.",
		isDefault:   false,
	}
	WritePackages = Scope{
		scope:       "write:packages",
		description: "Grants access to upload or publish a package in GitHub Packages.",
		isDefault:   false,
	}
	WriteGPGKey = Scope{
		scope:       "write:gpg_key",
		description: "Create, list, and view details for GPG keys.",
		isDefault:   false,
	}
	WritePublicKey = Scope{
		scope:       "write:public_key",
		description: "Create, list, and view details for public keys.",
		isDefault:   false,
	}
	ReadOrg = Scope{
		scope:       "read:org",
		description: "Read-only access to organization membership, organization projects, and team membership.",
		isDefault:   false,
	}
	ReadPackages = Scope{
		scope:       "read:packages",
		description: "Grants access to download or install packages from GitHub Packages.",
		isDefault:   false,
	}
	ReadGPGKey = Scope{
		scope:       "read:gpg_key",
		description: "List and view details for GPG keys.",
		isDefault:   false,
	}
	ReadDiscussion = Scope{
		scope:       "read:discussion",
		description: "Allows read access for team discussions.",
		isDefault:   false,
	}
	ReadUser = Scope{
		scope:       "read:user",
		description: "Grants access to read a user's profile data.",
		isDefault:   false,
	}
	ReadRepoHook = Scope{
		scope:       "read:repo_hook",
		description: "Grants read and ping access to hooks in public or private repositories.",
		isDefault:   false,
	}
	ReadPublicKey = Scope{
		scope:       "read:public_key",
		description: "List and view details for public keys.",
		isDefault:   false,
	}
	DeletePackages = Scope{
		scope:       "delete:packages",
		description: "Grants access to delete packages from GitHub Packages.",
		isDefault:   false,
	}
	WriteRepoHook = Scope{
		scope:       "write:repo_hook",
		description: "Grants read, write, and ping access to hooks in public or private repositories.",
		isDefault:   false,
	}
	AdminOrg = Scope{
		scope:       "admin:org",
		description: "Fully manage the organization and its teams, projects, and memberships.",
		isDefault:   false,
	}
	AdminPublicKey = Scope{
		scope:       "admin:public_key",
		description: "Fully manage public keys.",
		isDefault:   false,
	}
	UserEmail = Scope{
		scope:       "user:email",
		description: "Grants read access to a user's email addresses.",
		isDefault:   false,
	}
	UserFollow = Scope{
		scope:       "user:follow",
		description: "Grants access to follow or unfollow other users.",
		isDefault:   false,
	}
	Gist = Scope{
		scope:       "gist",
		description: "Grants write access to gists.",
		isDefault:   false,
	}
	Notifications = Scope{
		scope:       "notifications",
		description: "Grants: read access to a user's notifications, mark as read access to threads, watch and unwatch access to a repository, and read, write, and delete access to thread subscriptions.",
		isDefault:   false,
	}
	User = Scope{
		scope:       "user",
		description: "Grants read/write access to profile info only. Note that this scope includes user:email and user:follow.",
		isDefault:   false,
	}
	DeleteRepo = Scope{
		scope:       "delete_repo",
		description: "Grants access to delete adminable repositories.",
		isDefault:   false,
	}
	AdminGPGKey = Scope{
		scope:       "admin:gpg_key",
		description: "Fully manage GPG keys.",
		isDefault:   false,
	}
)
