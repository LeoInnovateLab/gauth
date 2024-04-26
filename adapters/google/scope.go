package google

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
		UserOpenid,
		UserEmail,
		UserProfile,
		UserPhonenumbersRead,
		UserOrganizationRead,
		UserGenderRead,
		UserEmailsRead,
		UserBirthdayRead,
		UserAddressesRead,
		UserinfoProfile,
		UserinfoEmail,
		YtAnalyticsReadonly,
		YtAnalyticsMonetaryReadonly,
		YoutubepartnerChannelAudit,
		Youtubepartner,
		YoutubeUpload,
		YoutubeReadonly,
		YoutubeForceSsl,
		YoutubeChannelMembershipsCreator,
		Youtube,
		WebmastersReadonly,
		Webmasters,
		Verifiedaccess,
		TraceAppend,
		TasksReadonly,
		Tasks,
		TagmanagerReadonly,
		TagmanagerPublish,
		TagmanagerManageUsers,
		TagmanagerManageAccounts,
		TagmanagerEditContainerversions,
		TagmanagerEditContainers,
		TagmanagerDeleteContainers,
		Streetviewpublish,
		SqlserviceAdmin,
		SpreadsheetsReadonly,
		Spreadsheets,
		SpannerData,
		SpannerAdmin,
		SourceReadWrite,
		SourceReadOnly,
		SourceFullControl,
		SiteverificationVerifyOnly,
		Siteverification,
		Servicecontrol,
		ServiceManagementReadonly,
		ServiceManagement,
		ScriptProjectsReadonly,
		ScriptProjects,
		ScriptProcesses,
		ScriptMetrics,
		ScriptDeploymentsReadonly,
		ScriptDeployments,
		Pubsub,
		PresentationsReadonly,
		Presentations,
		PhotoslibrarySharing,
		PhotoslibraryReadonlyAppcreateddata,
		PhotoslibraryReadonly,
		PhotoslibraryAppendonly,
		Photoslibrary,
		NdevCloudmanReadonly,
		NdevCloudman,
		NdevClouddnsReadwrite,
		NdevClouddnsReadonly,
		MonitoringWrite,
		MonitoringRead,
		Monitoring,
		ManufacturerCenter,
		LoggingWrite,
		LoggingRead,
		LoggingAdmin,
		Jobs,
		Indexing,
		Groups,
		Gmail,
		GmailSettingsSharing,
		GmailSettingsBasic,
		GmailSend,
		GmailReadonly,
		GmailModify,
		GmailMetadata,
		GmailLabels,
		GmailInsert,
		GmailCompose,
		GmailAddonsCurrentMessageReadonly,
		GmailAddonsCurrentMessageMetadata,
		GmailAddonsCurrentMessageAction,
		GmailAddonsCurrentActionCompose,
		Genomics,
		Games,
		FormsCurrentonly,
		Forms,
		FitnessReproductiveHealthWrite,
		FitnessReproductiveHealthRead,
		FitnessOxygenSaturationWrite,
		FitnessOxygenSaturationRead,
		FitnessNutritionWrite,
		FitnessNutritionRead,
		FitnessLocationWrite,
		FitnessLocationRead,
		FitnessBodyTemperatureWrite,
		FitnessBodyTemperatureRead,
		FitnessBodyWrite,
		FitnessBodyRead,
		FitnessBloodPressureWrite,
		FitnessBloodPressureRead,
		FitnessBloodGlucoseWrite,
		FitnessBloodGlucoseRead,
		FitnessActivityWrite,
		FitnessActivityRead,
		FirebaseReadonly,
		Firebase,
		EdiscoveryReadonly,
		Ediscovery,
		DriveScripts,
		DriveReadonly,
		DrivePhotosReadonly,
		DriveMetadataReadonly,
		DriveMetadata,
		DriveFile,
		DriveAppdata,
		DriveActivityReadonly,
		DriveActivity,
		Drive,
		Activity,
		DoubleClickSearch,
		DoubleClickBidmanager,
		DocumentsReadonly,
		Documents,
		DisplayVideo,
		DirectoryReadonly,
		Dialogflow,
		DfaTrafficking,
		DfaReporting,
		DevstorageReadWrite,
		DevstorageReadOnly,
		DevstorageFullControl,
		Ddmconversions,
		Datastore,
		Content,
		Compute,
		CloudPlatform,
		ContactsReadonly,
		CloudSearch,
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
	UserOpenid = Scope{
		scope:       "openid",
		description: "Associate you with your personal info on Google",
		isDefault:   true,
	}
	UserEmail = Scope{
		scope:       "email",
		description: "View your email address",
		isDefault:   true,
	}
	UserProfile = Scope{
		scope:       "profile",
		description: "View your basic profile info",
		isDefault:   true,
	}
	UserPhonenumbersRead = Scope{
		scope:       "https://www.googleapis.com/auth/user.phonenumbers.read",
		description: "View your phone numbers",
		isDefault:   false,
	}
	UserOrganizationRead = Scope{
		scope:       "https://www.googleapis.com/auth/user.organization.read",
		description: "See your education, work history and org info",
		isDefault:   false,
	}
	UserGenderRead = Scope{
		scope:       "https://www.googleapis.com/auth/user.gender.read",
		description: "See your gender",
		isDefault:   false,
	}
	UserEmailsRead = Scope{
		scope:       "https://www.googleapis.com/auth/user.emails.read",
		description: "View your email addresses",
		isDefault:   false,
	}
	UserBirthdayRead = Scope{
		scope:       "https://www.googleapis.com/auth/user.birthday.read",
		description: "View your complete date of birth",
		isDefault:   false,
	}
	UserAddressesRead = Scope{
		scope:       "https://www.googleapis.com/auth/user.addresses.read",
		description: "View your street addresses",
		isDefault:   false,
	}
	UserinfoProfile = Scope{
		scope:       "https://www.googleapis.com/auth/userinfo.profile",
		description: "See your personal info, including any personal info you've made publicly available",
		isDefault:   false,
	}
	UserinfoEmail = Scope{
		scope:       "https://www.googleapis.com/auth/userinfo.email",
		description: "View your email address",
		isDefault:   false,
	}
	YtAnalyticsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/yt-analytics.readonly",
		description: "View YouTube Analytics reports for your YouTube content",
		isDefault:   false,
	}
	YtAnalyticsMonetaryReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
		description: "View monetary and non-monetary YouTube Analytics reports for your YouTube content",
		isDefault:   false,
	}
	YoutubepartnerChannelAudit = Scope{
		scope:       "https://www.googleapis.com/auth/youtubepartner-channel-audit",
		description: "View private information of your YouTube channel relevant during the audit process with a YouTube partner",
		isDefault:   false,
	}
	Youtubepartner = Scope{
		scope:       "https://www.googleapis.com/auth/youtubepartner",
		description: "View and manage your assets and associated content on YouTube",
		isDefault:   false,
	}
	YoutubeUpload = Scope{
		scope:       "https://www.googleapis.com/auth/youtube.upload",
		description: "Manage your YouTube videos",
		isDefault:   false,
	}
	YoutubeReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/youtube.readonly",
		description: "View your YouTube account",
		isDefault:   false,
	}
	YoutubeForceSsl = Scope{
		scope:       "https://www.googleapis.com/auth/youtube.force-ssl",
		description: "See, edit, and permanently delete your YouTube videos, ratings, comments and captions",
		isDefault:   false,
	}
	YoutubeChannelMembershipsCreator = Scope{
		scope:       "https://www.googleapis.com/auth/youtube.channel-memberships.creator",
		description: "See a list of your current active channel members, their current level, and when they became a member",
		isDefault:   false,
	}
	Youtube = Scope{
		scope:       "https://www.googleapis.com/auth/youtube",
		description: "Manage your YouTube account",
		isDefault:   false,
	}
	WebmastersReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/webmasters.readonly",
		description: "View Search Console data for your verified sites",
		isDefault:   false,
	}
	Webmasters = Scope{
		scope:       "https://www.googleapis.com/auth/webmasters",
		description: "View and manage Search Console data for your verified sites",
		isDefault:   false,
	}
	Verifiedaccess = Scope{
		scope:       "https://www.googleapis.com/auth/verifiedaccess",
		description: "Verify your enterprise credentials",
		isDefault:   false,
	}
	TraceAppend = Scope{
		scope:       "https://www.googleapis.com/auth/trace.append",
		description: "Write Trace data for a project or application",
		isDefault:   false,
	}
	TasksReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/tasks.readonly",
		description: "View your tasks",
		isDefault:   false,
	}
	Tasks = Scope{
		scope:       "https://www.googleapis.com/auth/tasks",
		description: "Create, edit, organize, and delete all your tasks",
		isDefault:   false,
	}
	TagmanagerReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.readonly",
		description: "View your Google Tag Manager container and its subcomponents",
		isDefault:   false,
	}
	TagmanagerPublish = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.publish",
		description: "Publish your Google Tag Manager container versions",
		isDefault:   false,
	}
	TagmanagerManageUsers = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.manage.users",
		description: "Manage user permissions of your Google Tag Manager account and container",
		isDefault:   false,
	}
	TagmanagerManageAccounts = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.manage.accounts",
		description: "View and manage your Google Tag Manager accounts",
		isDefault:   false,
	}
	TagmanagerEditContainerversions = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.edit.containerversions",
		description: "Manage your Google Tag Manager container versions",
		isDefault:   false,
	}
	TagmanagerEditContainers = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.edit.containers",
		description: "Manage your Google Tag Manager container and its subcomponents, excluding versioning and publishing",
		isDefault:   false,
	}
	TagmanagerDeleteContainers = Scope{
		scope:       "https://www.googleapis.com/auth/tagmanager.delete.containers",
		description: "Delete your Google Tag Manager containers",
		isDefault:   false,
	}
	Streetviewpublish = Scope{
		scope:       "https://www.googleapis.com/auth/streetviewpublish",
		description: "Publish and manage your 360 photos on Google Street View",
		isDefault:   false,
	}
	SqlserviceAdmin = Scope{
		scope:       "https://www.googleapis.com/auth/sqlservice.admin",
		description: "Manage your Google SQL Service instances",
		isDefault:   false,
	}
	SpreadsheetsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/spreadsheets.readonly",
		description: "View your Google Spreadsheets",
		isDefault:   false,
	}
	Spreadsheets = Scope{
		scope:       "https://www.googleapis.com/auth/spreadsheets",
		description: "See, edit, create, and delete your spreadsheets in Google Drive",
		isDefault:   false,
	}
	SpannerData = Scope{
		scope:       "https://www.googleapis.com/auth/spanner.data",
		description: "View and manage the contents of your Spanner databases",
		isDefault:   false,
	}
	SpannerAdmin = Scope{
		scope:       "https://www.googleapis.com/auth/spanner.admin",
		description: "Administer your Spanner databases",
		isDefault:   false,
	}
	SourceReadWrite = Scope{
		scope:       "https://www.googleapis.com/auth/source.read_write",
		description: "Manage the contents of your source code repositories",
		isDefault:   false,
	}
	SourceReadOnly = Scope{
		scope:       "https://www.googleapis.com/auth/source.read_only",
		description: "View the contents of your source code repositories",
		isDefault:   false,
	}
	SourceFullControl = Scope{
		scope:       "https://www.googleapis.com/auth/source.full_control",
		description: "Manage your source code repositories",
		isDefault:   false,
	}
	SiteverificationVerifyOnly = Scope{
		scope:       "https://www.googleapis.com/auth/siteverification.verify_only",
		description: "Manage your new site verifications with Google",
		isDefault:   false,
	}
	Siteverification = Scope{
		scope:       "https://www.googleapis.com/auth/siteverification",
		description: "Manage the list of sites and domains you control",
		isDefault:   false,
	}
	Servicecontrol = Scope{
		scope:       "https://www.googleapis.com/auth/servicecontrol",
		description: "Manage your Google Service Control data",
		isDefault:   false,
	}
	ServiceManagementReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/service.management.readonly",
		description: "View your Google API service configuration",
		isDefault:   false,
	}
	ServiceManagement = Scope{
		scope:       "https://www.googleapis.com/auth/service.management",
		description: "Manage your Google API service configuration",
		isDefault:   false,
	}
	ScriptProjectsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/script.projects.readonly",
		description: "View Google Apps Script projects",
		isDefault:   false,
	}
	ScriptProjects = Scope{
		scope:       "https://www.googleapis.com/auth/script.projects",
		description: "Create and update Google Apps Script projects",
		isDefault:   false,
	}
	ScriptProcesses = Scope{
		scope:       "https://www.googleapis.com/auth/script.processes",
		description: "View Google Apps Script processes",
		isDefault:   false,
	}
	ScriptMetrics = Scope{
		scope:       "https://www.googleapis.com/auth/script.metrics",
		description: "View Google Apps Script project's metrics",
		isDefault:   false,
	}
	ScriptDeploymentsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/script.deployments.readonly",
		description: "View Google Apps Script deployments",
		isDefault:   false,
	}
	ScriptDeployments = Scope{
		scope:       "https://www.googleapis.com/auth/script.deployments",
		description: "Create and update Google Apps Script deployments",
		isDefault:   false,
	}
	Pubsub = Scope{
		scope:       "https://www.googleapis.com/auth/pubsub",
		description: "View and manage Pub/Sub topics and subscriptions",
		isDefault:   false,
	}
	PresentationsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/presentations.readonly",
		description: "View your Google Slides presentations",
		isDefault:   false,
	}
	Presentations = Scope{
		scope:       "https://www.googleapis.com/auth/presentations",
		description: "View and manage your Google Slides presentations",
		isDefault:   false,
	}
	PhotoslibrarySharing = Scope{
		scope:       "https://www.googleapis.com/auth/photoslibrary.sharing",
		description: "Manage and add to shared albums on your behalf",
		isDefault:   false,
	}
	PhotoslibraryReadonlyAppcreateddata = Scope{
		scope:       "https://www.googleapis.com/auth/photoslibrary.readonly.appcreateddata",
		description: "Manage photos added by this app",
		isDefault:   false,
	}
	PhotoslibraryReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/photoslibrary.readonly",
		description: "View your Google Photos library",
		isDefault:   false,
	}
	PhotoslibraryAppendonly = Scope{
		scope:       "https://www.googleapis.com/auth/photoslibrary.appendonly",
		description: "Add to your Google Photos library",
		isDefault:   false,
	}
	Photoslibrary = Scope{
		scope:       "https://www.googleapis.com/auth/photoslibrary",
		description: "View and manage your Google Photos library",
		isDefault:   false,
	}
	NdevCloudmanReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/ndev.cloudman.readonly",
		description: "View your Google Cloud Platform management resources and deployment status information",
		isDefault:   false,
	}
	NdevCloudman = Scope{
		scope:       "https://www.googleapis.com/auth/ndev.cloudman",
		description: "View and manage your Google Cloud Platform management resources and deployment status information",
		isDefault:   false,
	}
	NdevClouddnsReadwrite = Scope{
		scope:       "https://www.googleapis.com/auth/ndev.clouddns.readwrite",
		description: "View and manage your DNS records hosted by Google Cloud DNS",
		isDefault:   false,
	}
	NdevClouddnsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/ndev.clouddns.readonly",
		description: "View your DNS records hosted by Google Cloud DNS",
		isDefault:   false,
	}
	MonitoringWrite = Scope{
		scope:       "https://www.googleapis.com/auth/monitoring.write",
		description: "Publish metric data to your Google Cloud projects",
		isDefault:   false,
	}
	MonitoringRead = Scope{
		scope:       "https://www.googleapis.com/auth/monitoring.read",
		description: "View monitoring data for all of your Google Cloud and third-party projects",
		isDefault:   false,
	}
	Monitoring = Scope{
		scope:       "https://www.googleapis.com/auth/monitoring",
		description: "View and write monitoring data for all of your Google and third-party Cloud and API projects",
		isDefault:   false,
	}
	ManufacturerCenter = Scope{
		scope:       "https://www.googleapis.com/auth/manufacturercenter",
		description: "Manage your product listings for Google Manufacturer Center",
		isDefault:   false,
	}
	LoggingWrite = Scope{
		scope:       "https://www.googleapis.com/auth/logging.write",
		description: "Submit log data for your projects",
		isDefault:   false,
	}
	LoggingRead = Scope{
		scope:       "https://www.googleapis.com/auth/logging.read",
		description: "View log data for your projects",
		isDefault:   false,
	}
	LoggingAdmin = Scope{
		scope:       "https://www.googleapis.com/auth/logging.admin",
		description: "Administrate log data for your projects",
		isDefault:   false,
	}
	Jobs = Scope{
		scope:       "https://www.googleapis.com/auth/jobs",
		description: "Manage job postings",
		isDefault:   false,
	}
	Indexing = Scope{
		scope:       "https://www.googleapis.com/auth/indexing",
		description: "Submit data to Google for indexing",
		isDefault:   false,
	}
	Groups = Scope{
		scope:       "https://www.googleapis.com/auth/groups",
		description: "View and manage your Google Groups",
		isDefault:   false,
	}
	Gmail = Scope{
		scope:       "https://mail.google.com/",
		description: "Read, compose, send, and permanently delete all your email from Gmail",
		isDefault:   false,
	}
	GmailSettingsSharing = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.settings.sharing",
		description: "Manage your sensitive mail settings, including who can manage your mail",
		isDefault:   false,
	}
	GmailSettingsBasic = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.settings.basic",
		description: "Manage your basic mail settings",
		isDefault:   false,
	}
	GmailSend = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.send",
		description: "Send email on your behalf",
		isDefault:   false,
	}
	GmailReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.readonly",
		description: "View your email messages and settings",
		isDefault:   false,
	}
	GmailModify = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.modify",
		description: "View and modify but not delete your email",
		isDefault:   false,
	}
	GmailMetadata = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.metadata",
		description: "View your email message metadata such as labels and headers, but not the email body",
		isDefault:   false,
	}
	GmailLabels = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.labels",
		description: "Manage mailbox labels",
		isDefault:   false,
	}
	GmailInsert = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.insert",
		description: "Insert mail into your mailbox",
		isDefault:   false,
	}
	GmailCompose = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.compose",
		description: "Manage drafts and send emails",
		isDefault:   false,
	}
	GmailAddonsCurrentMessageReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.addons.current.message.readonly",
		description: "View your email messages when the add-on is running",
		isDefault:   false,
	}
	GmailAddonsCurrentMessageMetadata = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.addons.current.message.metadata",
		description: "View your email message metadata when the add-on is running",
		isDefault:   false,
	}
	GmailAddonsCurrentMessageAction = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.addons.current.message.action",
		description: "View your email messages when you interact with the add-on",
		isDefault:   false,
	}
	GmailAddonsCurrentActionCompose = Scope{
		scope:       "https://www.googleapis.com/auth/gmail.addons.current.action.compose",
		description: "Manage drafts and send emails when you interact with the add-on",
		isDefault:   false,
	}
	Genomics = Scope{
		scope:       "https://www.googleapis.com/auth/genomics",
		description: "View and manage Genomics data",
		isDefault:   false,
	}
	Games = Scope{
		scope:       "https://www.googleapis.com/auth/games",
		description: "Create, edit, and delete your Google Play Games activity",
		isDefault:   false,
	}
	FormsCurrentonly = Scope{
		scope:       "https://www.googleapis.com/auth/forms.currentonly",
		description: "View and manage forms that this application has been installed in",
		isDefault:   false,
	}
	Forms = Scope{
		scope:       "https://www.googleapis.com/auth/forms",
		description: "View and manage your forms in Google Drive",
		isDefault:   false,
	}
	FitnessReproductiveHealthWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.reproductive_health.write",
		description: "See and add info about your reproductive health in Google Fit. I consent to Google sharing my reproductive health information with this app.",
		isDefault:   false,
	}
	FitnessReproductiveHealthRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.reproductive_health.read",
		description: "See info about your reproductive health in Google Fit. I consent to Google sharing my reproductive health information with this app.",
		isDefault:   false,
	}
	FitnessOxygenSaturationWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.oxygen_saturation.write",
		description: "See and add info about your oxygen saturation in Google Fit. I consent to Google sharing my oxygen saturation information with this app.",
		isDefault:   false,
	}
	FitnessOxygenSaturationRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.oxygen_saturation.read",
		description: "See info about your oxygen saturation in Google Fit. I consent to Google sharing my oxygen saturation information with this app.",
		isDefault:   false,
	}
	FitnessNutritionWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.nutrition.write",
		description: "See and add to info about your nutrition in Google Fit",
		isDefault:   false,
	}
	FitnessNutritionRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.nutrition.read",
		description: "See info about your nutrition in Google Fit",
		isDefault:   false,
	}
	FitnessLocationWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.location.write",
		description: "See and add to your Google Fit location data",
		isDefault:   false,
	}
	FitnessLocationRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.location.read",
		description: "See your Google Fit speed and distance data",
		isDefault:   false,
	}
	FitnessBodyTemperatureWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.body_temperature.write",
		description: "See and add to info about your body temperature in Google Fit. I consent to Google sharing my body temperature information with this app.",
		isDefault:   false,
	}
	FitnessBodyTemperatureRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.body_temperature.read",
		description: "See info about your body temperature in Google Fit. I consent to Google sharing my body temperature information with this app.",
		isDefault:   false,
	}
	FitnessBodyWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.body.write",
		description: "See and add info about your body measurements and heart rate to Google Fit",
		isDefault:   false,
	}
	FitnessBodyRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.body.read",
		description: "See info about your body measurements and heart rate in Google Fit",
		isDefault:   false,
	}
	FitnessBloodPressureWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.blood_pressure.write",
		description: "See and add info about your blood pressure in Google Fit. I consent to Google sharing my blood pressure information with this app.",
		isDefault:   false,
	}
	FitnessBloodPressureRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.blood_pressure.read",
		description: "See info about your blood pressure in Google Fit. I consent to Google sharing my blood pressure information with this app.",
		isDefault:   false,
	}
	FitnessBloodGlucoseWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.blood_glucose.write",
		description: "See and add info about your blood glucose to Google Fit. I consent to Google sharing my blood glucose information with this app.",
		isDefault:   false,
	}
	FitnessBloodGlucoseRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.blood_glucose.read",
		description: "See info about your blood glucose in Google Fit. I consent to Google sharing my blood glucose information with this app.",
		isDefault:   false,
	}
	FitnessActivityWrite = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.activity.write",
		description: "See and add to your Google Fit physical activity data",
		isDefault:   false,
	}
	FitnessActivityRead = Scope{
		scope:       "https://www.googleapis.com/auth/fitness.activity.read",
		description: "Use Google Fit to see and store your physical activity data",
		isDefault:   false,
	}
	FirebaseReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/firebase.readonly",
		description: "View all your Firebase data and settings",
		isDefault:   false,
	}
	Firebase = Scope{
		scope:       "https://www.googleapis.com/auth/firebase",
		description: "View and administer all your Firebase data and settings",
		isDefault:   false,
	}
	EdiscoveryReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/ediscovery.readonly",
		description: "View your eDiscovery data",
		isDefault:   false,
	}
	Ediscovery = Scope{
		scope:       "https://www.googleapis.com/auth/ediscovery",
		description: "Manage your eDiscovery data",
		isDefault:   false,
	}
	DriveScripts = Scope{
		scope:       "https://www.googleapis.com/auth/drive.scripts",
		description: "Modify your Google Apps Script scripts' behavior",
		isDefault:   false,
	}
	DriveReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/drive.readonly",
		description: "See and download all your Google Drive files",
		isDefault:   false,
	}
	DrivePhotosReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/drive.photos.readonly",
		description: "View the photos, videos and albums in your Google Photos",
		isDefault:   false,
	}
	DriveMetadataReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/drive.metadata.readonly",
		description: "View metadata for files in your Google Drive",
		isDefault:   false,
	}
	DriveMetadata = Scope{
		scope:       "https://www.googleapis.com/auth/drive.metadata",
		description: "View and manage metadata of files in your Google Drive",
		isDefault:   false,
	}
	DriveFile = Scope{
		scope:       "https://www.googleapis.com/auth/drive.file",
		description: "View and manage Google Drive files and folders that you have opened or created with this app",
		isDefault:   false,
	}
	DriveAppdata = Scope{
		scope:       "https://www.googleapis.com/auth/drive.appdata",
		description: "View and manage its own configuration data in your Google Drive",
		isDefault:   false,
	}
	DriveActivityReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/drive.activity.readonly",
		description: "View the activity record of files in your Google Drive",
		isDefault:   false,
	}
	DriveActivity = Scope{
		scope:       "https://www.googleapis.com/auth/drive.activity",
		description: "View and add to the activity record of files in your Google Drive",
		isDefault:   false,
	}
	Drive = Scope{
		scope:       "https://www.googleapis.com/auth/drive",
		description: "See, edit, create, and delete all of your Google Drive files",
		isDefault:   false,
	}
	Activity = Scope{
		scope:       "https://www.googleapis.com/auth/activity",
		description: "View the activity history of your Google apps",
		isDefault:   false,
	}
	DoubleClickSearch = Scope{
		scope:       "https://www.googleapis.com/auth/doubleclicksearch",
		description: "View and manage your advertising data in DoubleClick Search",
		isDefault:   false,
	}
	DoubleClickBidmanager = Scope{
		scope:       "https://www.googleapis.com/auth/doubleclickbidmanager",
		description: "View and manage your reports in DoubleClick Bid Manager",
		isDefault:   false,
	}
	DocumentsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/documents.readonly",
		description: "View your Google Docs documents",
		isDefault:   false,
	}
	Documents = Scope{
		scope:       "https://www.googleapis.com/auth/documents",
		description: "View and manage your Google Docs documents",
		isDefault:   false,
	}
	DisplayVideo = Scope{
		scope:       "https://www.googleapis.com/auth/display-video",
		description: "Create, see, edit, and permanently delete your Display & Video 360 entities and reports",
		isDefault:   false,
	}
	DirectoryReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/directory.readonly",
		description: "See and download your organization's GSuite directory",
		isDefault:   false,
	}
	Dialogflow = Scope{
		scope:       "https://www.googleapis.com/auth/dialogflow",
		description: "View, manage and query your Dialogflow agents",
		isDefault:   false,
	}
	DfaTrafficking = Scope{
		scope:       "https://www.googleapis.com/auth/dfatrafficking",
		description: "View and manage your DoubleClick Campaign Manager's (DCM) display ad campaigns",
		isDefault:   false,
	}
	DfaReporting = Scope{
		scope:       "https://www.googleapis.com/auth/dfareporting",
		description: "View and manage DoubleClick for Advertisers reports",
		isDefault:   false,
	}
	DevstorageReadWrite = Scope{
		scope:       "https://www.googleapis.com/auth/devstorage.read_write",
		description: "Manage your data in Google Cloud Storage",
		isDefault:   false,
	}
	DevstorageReadOnly = Scope{
		scope:       "https://www.googleapis.com/auth/devstorage.read_only",
		description: "View your data in Google Cloud Storage",
		isDefault:   false,
	}
	DevstorageFullControl = Scope{
		scope:       "https://www.googleapis.com/auth/devstorage.full_control",
		description: "Manage your data and permissions in Google Cloud Storage",
		isDefault:   false,
	}
	Ddmconversions = Scope{
		scope:       "https://www.googleapis.com/auth/ddmconversions",
		description: "Manage DoubleClick Digital Marketing conversions",
		isDefault:   false,
	}
	Datastore = Scope{
		scope:       "https://www.googleapis.com/auth/datastore",
		description: "View and manage your Google Cloud Datastore data",
		isDefault:   false,
	}
	Content = Scope{
		scope:       "https://www.googleapis.com/auth/content",
		description: "Manage your product listings and accounts for Google Shopping",
		isDefault:   false,
	}
	ContactsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/contacts.readonly",
		description: "See and download your contacts",
		isDefault:   false,
	}
	ContactsOtherReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/contacts.other.readonly",
		description: "See and download contact info automatically saved in your \"Other contacts\"",
		isDefault:   false,
	}
	Contacts = Scope{
		scope:       "https://www.googleapis.com/auth/contacts",
		description: "See, edit, download, and permanently delete your contacts",
		isDefault:   false,
	}
	ContactsFeeds = Scope{
		scope:       "https://www.google.com/m8/feeds",
		description: "See, edit, download, and permanently delete your contacts",
		isDefault:   false,
	}
	ComputeReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/compute.readonly",
		description: "View your Google Compute Engine resources",
		isDefault:   false,
	}
	Compute = Scope{
		scope:       "https://www.googleapis.com/auth/compute",
		description: "View and manage your Google Compute Engine resources",
		isDefault:   false,
	}
	Cloudruntimeconfig = Scope{
		scope:       "https://www.googleapis.com/auth/cloudruntimeconfig",
		description: "Manage your Google Cloud Platform services' runtime configuration",
		isDefault:   false,
	}
	Cloudkms = Scope{
		scope:       "https://www.googleapis.com/auth/cloudkms",
		description: "View and manage your keys and secrets stored in Cloud Key Management Service",
		isDefault:   false,
	}
	Cloudiot = Scope{
		scope:       "https://www.googleapis.com/auth/cloudiot",
		description: "Register and manage devices in the Google Cloud IoT service",
		isDefault:   false,
	}
	CloudSearchStatsIndexing = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.stats.indexing",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearchStats = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.stats",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearchSettingsQuery = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.settings.query",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearchSettingsIndexing = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.settings.indexing",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearchSettings = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.settings",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearchQuery = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.query",
		description: "Search your organization's data in the Cloud Search index",
		isDefault:   false,
	}
	CloudSearchIndexing = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.indexing",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearchDebug = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search.debug",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudSearch = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_search",
		description: "Index and serve your organization's data with Cloud Search",
		isDefault:   false,
	}
	CloudDebugger = Scope{
		scope:       "https://www.googleapis.com/auth/cloud_debugger",
		description: "Use Stackdriver Debugger",
		isDefault:   false,
	}
	CloudVision = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-vision",
		description: "Apply machine learning models to understand and label images",
		isDefault:   false,
	}
	CloudTranslation = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-translation",
		description: "Translate text from one language to another using Google Translate",
		isDefault:   false,
	}
	CloudPlatformReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-platform.read-only",
		description: "View your data across Google Cloud Platform services",
		isDefault:   false,
	}
	CloudPlatform = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-platform",
		description: "View and manage your data across Google Cloud Platform services",
		isDefault:   false,
	}
	CloudLanguage = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-language",
		description: "Apply machine learning models to reveal the structure and meaning of text",
		isDefault:   false,
	}
	CloudIdentityGroupsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-identity.groups.readonly",
		description: "See any Cloud Identity Groups that you can access, including group members and their emails",
		isDefault:   false,
	}
	CloudIdentityGroups = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-identity.groups",
		description: "See, change, create, and delete any of the Cloud Identity Groups that you can access, including the members of each group",
		isDefault:   false,
	}
	CloudBigtableAdminTable = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-bigtable.admin.table",
		description: "Administer your Cloud Bigtable tables",
		isDefault:   false,
	}
	CloudBigtableAdminCluster = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-bigtable.admin.cluster",
		description: "Administer your Cloud Bigtable clusters",
		isDefault:   false,
	}
	CloudBigtableAdmin = Scope{
		scope:       "https://www.googleapis.com/auth/cloud-bigtable.admin",
		description: "Administer your Cloud Bigtable tables and clusters",
		isDefault:   false,
	}
	ClassroomTopicsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.topics.readonly",
		description: "View topics in Google Classroom",
		isDefault:   false,
	}
	ClassroomTopics = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.topics",
		description: "See, create, and edit topics in Google Classroom",
		isDefault:   false,
	}
	ClassroomStudentSubmissionsStudentsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.student-submissions.students.readonly",
		description: "View course work and grades for students in the Google Classroom classes you teach or administer",
		isDefault:   false,
	}
	ClassroomStudentSubmissionsMeReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.student-submissions.me.readonly",
		description: "View your course work and grades in Google Classroom",
		isDefault:   false,
	}
	ClassroomRostersReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.rosters.readonly",
		description: "View your Google Classroom class rosters",
		isDefault:   false,
	}
	ClassroomRosters = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.rosters",
		description: "Manage your Google Classroom class rosters",
		isDefault:   false,
	}
	ClassroomPushNotifications = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.push-notifications",
		description: "Receive notifications about your Google Classroom data",
		isDefault:   false,
	}
	ClassroomProfilePhotos = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.profile.photos",
		description: "View the profile photos of people in your classes",
		isDefault:   false,
	}
	ClassroomProfileEmails = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.profile.emails",
		description: "View the email addresses of people in your classes",
		isDefault:   false,
	}
	ClassroomGuardianlinksStudentsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.guardianlinks.students.readonly",
		description: "View guardians for students in your Google Classroom classes",
		isDefault:   false,
	}
	ClassroomGuardianlinksStudents = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.guardianlinks.students",
		description: "View and manage guardians for students in your Google Classroom classes",
		isDefault:   false,
	}
	ClassroomGuardianlinksMeReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.guardianlinks.me.readonly",
		description: "View your Google Classroom guardians",
		isDefault:   false,
	}
	ClassroomCourseworkStudentsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.coursework.students.readonly",
		description: "View course work and grades for students in the Google Classroom classes you teach or administer",
		isDefault:   false,
	}
	ClassroomCourseworkStudents = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.coursework.students",
		description: "Manage course work and grades for students in the Google Classroom classes you teach and view the course work and grades for classes you administer",
		isDefault:   false,
	}
	ClassroomCourseworkMeReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.coursework.me.readonly",
		description: "View your course work and grades in Google Classroom",
		isDefault:   false,
	}
	ClassroomCourseworkMe = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.coursework.me",
		description: "Manage your course work and view your grades in Google Classroom",
		isDefault:   false,
	}
	ClassroomCoursesReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.courses.readonly",
		description: "View your Google Classroom classes",
		isDefault:   false,
	}
	ClassroomCourses = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.courses",
		description: "Manage your Google Classroom classes",
		isDefault:   false,
	}
	ClassroomAnnouncementsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.announcements.readonly",
		description: "View announcements in Google Classroom",
		isDefault:   false,
	}
	ClassroomAnnouncements = Scope{
		scope:       "https://www.googleapis.com/auth/classroom.announcements",
		description: "View and manage announcements in Google Classroom",
		isDefault:   false,
	}
	CalendarSettingsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/calendar.settings.readonly",
		description: "View your Calendar settings",
		isDefault:   false,
	}
	CalendarReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/calendar.readonly",
		description: "View your calendars",
		isDefault:   false,
	}
	CalendarEventsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/calendar.events.readonly",
		description: "View events on all your calendars",
		isDefault:   false,
	}
	CalendarEvents = Scope{
		scope:       "https://www.googleapis.com/auth/calendar.events",
		description: "View and edit events on all your calendars",
		isDefault:   false,
	}
	BigqueryReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/bigquery.readonly",
		description: "View your data in Google BigQuery",
		isDefault:   false,
	}
	Bigquery = Scope{
		scope:       "https://www.googleapis.com/auth/bigquery",
		description: "View and manage your data in Google BigQuery",
		isDefault:   false,
	}
	BigQueryInsertdata = Scope{
		scope:       "https://www.googleapis.com/auth/bigquery.insertdata",
		description: "Insert data into Google BigQuery",
		isDefault:   false,
	}
	BigQuery = Scope{
		scope:       "https://www.googleapis.com/auth/bigquery",
		description: "View and manage your data in Google BigQuery",
		isDefault:   false,
	}
	AppengineAdmin = Scope{
		scope:       "https://www.googleapis.com/auth/appengine.admin",
		description: "View and manage your applications deployed on Google App Engine",
		isDefault:   false,
	}
	Appengine = Scope{
		scope:       "https://www.googleapis.com/auth/appengine",
		description: "View your applications deployed on Google App Engine",
		isDefault:   false,
	}
	AndroidPublisher = Scope{
		scope:       "https://www.googleapis.com/auth/androidpublisher",
		description: "View and manage your Google Play Developer account",
		isDefault:   false,
	}
	AndroidManagement = Scope{
		scope:       "https://www.googleapis.com/auth/androidmanagement",
		description: "Manage Android devices and apps for your customers",
		isDefault:   false,
	}
	AndroidEnterprise = Scope{
		scope:       "https://www.googleapis.com/auth/androidenterprise",
		description: "Manage corporate Android devices",
		isDefault:   false,
	}
	AnalyticsUserDeletion = Scope{
		scope:       "https://www.googleapis.com/auth/analytics.user.deletion",
		description: "Manage Google Analytics user deletion requests",
		isDefault:   false,
	}
	AnalyticsReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/analytics.readonly",
		description: "View your Google Analytics data",
		isDefault:   false,
	}
	AnalyticsProvision = Scope{
		scope:       "https://www.googleapis.com/auth/analytics.provision",
		description: "Create a new Google Analytics account along with its default property and view",
		isDefault:   false,
	}
	AnalyticsManageUsersReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/analytics.manage.users.readonly",
		description: "View Google Analytics user permissions",
		isDefault:   false,
	}
	AnalyticsManageUsers = Scope{
		scope:       "https://www.googleapis.com/auth/analytics.manage.users",
		description: "Manage Google Analytics Account users by email address",
		isDefault:   false,
	}
	AnalyticsEdit = Scope{
		scope:       "https://www.googleapis.com/auth/analytics.edit",
		description: "Edit Google Analytics management entities",
		isDefault:   false,
	}
	Analytics = Scope{
		scope:       "https://www.googleapis.com/auth/analytics",
		description: "View and manage your Google Analytics data",
		isDefault:   false,
	}
	Adsensehost = Scope{
		scope:       "https://www.googleapis.com/auth/adsensehost",
		description: "View and manage your AdSense host data and associated accounts",
		isDefault:   false,
	}
	AdsenseReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/adsense.readonly",
		description: "View your AdSense data",
		isDefault:   false,
	}
	Adsense = Scope{
		scope:       "https://www.googleapis.com/auth/adsense",
		description: "View and manage your AdSense data",
		isDefault:   false,
	}
	AdminReportsUsageReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.reports.usage.readonly",
		description: "View usage reports for your G Suite domain",
		isDefault:   false,
	}
	AdminReportsAuditReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.reports.audit.readonly",
		description: "View audit reports for your G Suite domain",
		isDefault:   false,
	}
	AdminDirectoryUserschemaReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.userschema.readonly",
		description: "View user schemas on your domain",
		isDefault:   false,
	}
	AdminDirectoryUserschema = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.userschema",
		description: "View and manage the provisioning of user schemas on your domain",
		isDefault:   false,
	}
	AdminDirectoryUserSecurity = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.user.security",
		description: "Manage data access permissions for users on your domain",
		isDefault:   false,
	}
	AdminDirectoryUserReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.user.readonly",
		description: "View users on your domain",
		isDefault:   false,
	}
	AdminDirectoryUserAliasReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.user.alias.readonly",
		description: "View user aliases on your domain",
		isDefault:   false,
	}
	AdminDirectoryUserAlias = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.user.alias",
		description: "View and manage user aliases on your domain",
		isDefault:   false,
	}
	AdminDirectoryUser = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.user",
		description: "View and manage the provisioning of users on your domain",
		isDefault:   false,
	}
	AdminDirectoryRolemanagementReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.rolemanagement.readonly",
		description: "View delegated admin roles for your domain",
		isDefault:   false,
	}
	AdminDirectoryRolemanagement = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.rolemanagement",
		description: "Manage delegated admin roles for your domain",
		isDefault:   false,
	}
	AdminDirectoryResourceCalendarReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.resource.calendar.readonly",
		description: "View calendar resources on your domain",
		isDefault:   false,
	}
	AdminDirectoryResourceCalendar = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.resource.calendar",
		description: "View and manage the provisioning of calendar resources on your domain",
		isDefault:   false,
	}
	AdminDirectoryOrgunitReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.orgunit.readonly",
		description: "View organization units on your domain",
		isDefault:   false,
	}
	AdminDirectoryOrgunit = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.orgunit",
		description: "View and manage organization units on your domain",
		isDefault:   false,
	}
	AdminDirectoryNotifications = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.notifications",
		description: "View and manage notifications received on your domain",
		isDefault:   false,
	}
	AdminDirectoryGroupReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.group.readonly",
		description: "View groups on your domain",
		isDefault:   false,
	}
	AdminDirectoryGroupMemberReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.group.member.readonly",
		description: "View group subscriptions on your domain",
		isDefault:   false,
	}
	AdminDirectoryGroupMember = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.group.member",
		description: "View and manage group subscriptions on your domain",
		isDefault:   false,
	}
	AdminDirectoryGroup = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.group",
		description: "View and manage the provisioning of groups on your domain",
		isDefault:   false,
	}
	AdminDirectoryDomainReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.domain.readonly",
		description: "View domains related to your customers",
		isDefault:   false,
	}
	AdminDirectoryDomain = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.domain",
		description: "View and manage the provisioning of domains for your customers",
		isDefault:   false,
	}
	AdminDirectoryDeviceMobileReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.device.mobile.readonly",
		description: "View your mobile devices' metadata",
		isDefault:   false,
	}
	AdminDirectoryDeviceMobileAction = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.device.mobile.action",
		description: "Manage your mobile devices by performing administrative tasks",
		isDefault:   false,
	}
	AdminDirectoryDeviceMobile = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.device.mobile",
		description: "View and manage your mobile devices' metadata",
		isDefault:   false,
	}
	AdminDirectoryDeviceChromeosReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly",
		description: "View your Chrome OS devices' metadata",
		isDefault:   false,
	}
	AdminDirectoryDeviceChromeos = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.device.chromeos",
		description: "View and manage your Chrome OS devices' metadata",
		isDefault:   false,
	}
	AdminDirectoryCustomerReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.customer.readonly",
		description: "View customer related information",
		isDefault:   false,
	}
	AdminDirectoryCustomer = Scope{
		scope:       "https://www.googleapis.com/auth/admin.directory.customer",
		description: "View and manage customer related information",
		isDefault:   false,
	}
	AdminDatatransferReadonly = Scope{
		scope:       "https://www.googleapis.com/auth/admin.datatransfer.readonly",
		description: "View data transfers between users in your organization",
		isDefault:   false,
	}
	AdminDatatransfer = Scope{
		scope:       "https://www.googleapis.com/auth/admin.datatransfer",
		description: "View and manage data transfers between users in your organization",
		isDefault:   false,
	}
	AdexchangeBuyer = Scope{
		scope:       "https://www.googleapis.com/auth/adexchange.buyer",
		description: "Manage your Ad Exchange buyer account configuration",
		isDefault:   false,
	}
)
