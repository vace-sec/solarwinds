package solarwinds

// Client describes a customer
type Client struct {
	ID                    int    `xml:"clientid"`
	Name                  string `xml:"name"`
	ViewDashboard         int    `xml:"view_dashboard"`
	ViewWorkstationAssets int    `xml:"view_wkstsn_assets"`
	DashboardUsername     string `xml:"dashboard_username"`
	Timezone              string `xml:"timezone"`
	CreationDate          string `xml:"creation_date"`
	ServerCount           int    `xml:"server_count"`
	WorkstationCount      int    `xml:"workstation_count"`
	MobileDeviceCount     int    `xml:"mobile_device_count"`
	DeviceCount           int    `xml:"device_count"`
}

// Site describes the different locations of a client
type Site struct {
	ID              int    `xml:"siteid"`
	Name            string `xml:"name"`
	PrimaryRouter   string `xml:"primary_router"`
	SecondaryRouter string `xml:"secondary_router"`
	ConnectionOK    string `xml:"connection_ok"`
	CreationDate    string `xml:"creation_date"`
}

// Server describes a computer used as a server
type Server struct {
	ID                   int    `xml:"serverid"`
	GUID                 string `xml:"guid"`
	Name                 string `xml:"name"`
	Description          string `xml:"description"`
	AgentVersion         string `xml:"agent_version"`
	InstallDate          string `xml:"install_date"`
	LastBootTime         int    `xml:"last_boot_time"`
	Online               int    `xml:"online"`
	Active24x7           int    `xml:"active_247"`
	CheckInterval24x7    int    `xml:"check_interval_247"`
	Status24x7           int    `xml:"status_247"`
	LocalDate24x7        string `xml:"local_date_247"`
	LocalTime24x7        string `xml:"local_time_247"`
	UTCTime24x7          string `xml:"utc_time_247"`
	DscActive            int    `xml:"dsc_active"`
	AgentMode            int    `xml:"agent_mode"`
	DscHour              int    `xml:"dsc_hour"`
	DscStatus            int    `xml:"dsc_status"`
	DscLocalDate         string `xml:"dsc_local_date"`
	DscLocalTime         string `xml:"dsc_local_time"`
	DscUTCTime           string `xml:"dsc_utc_time"`
	MissFactor24x7       string `xml:"miss_factor_247"`
	EMailOverdueAlert    int    `xml:"email_overdue_alert"`
	SMSOverdueAlert      int    `xml:"sms_overdue_alert"`
	Missed24x7           int    `xml:"missed_247"`
	RemoteConnectionType int    `xml:"remote_connection_type"`
	RemoteAddress        string `xml:"remote_address"`
	RemotePort           int    `xml:"remote_port"`
	RemoteUsername       string `xml:"remote_username"`
	RemoteDomain         string `xml:"remote_domain"`
	TzBias               int    `xml:"tz_bias"`
	TzDstBias            int    `xml:"tz_dst_bias"`
	TzStdBias            int    `xml:"tz_std_bias"`
	TzMode               int    `xml:"tz_mode"`
	ATzDstDate           string `xml:"atz_dst_date"`
	TzDstDate            string `xml:"tz_dst_date"`
	TzStdDate            string `xml:"tz_std_date"`
	UTCApt               string `xml:"utc_apt"`
	UTCOffset            int    `xml:"utc_offset"`
	AssetID              int    `xml:"assetid"`
	WINSName             string `xml:"wins_name"`
	User                 string `xml:"user"`
	Domain               string `xml:"domain"`
	Role                 string `xml:"role"`
	ChassisType          int    `xml:"chassis_type"`
	Manufacturer         string `xml:"manufacturer"`
	Model                string `xml:"model"`
	DeviceSerial         string `xml:"device_serial"`
	IP                   string `xml:"ip"`
	ExternalIP           string `xml:"external_ip"`
	Mac1                 string `xml:"mac1"`
	Mac2                 string `xml:"mac2"`
	Mac3                 string `xml:"mac3"`
	ProcessorCount       int    `xml:"processor_count"`
	TotalMemory          int    `xml:"total_memory"`
	OSType               int    `xml:"os_type"`
	OS                   string `xml:"os"`
	ServicePack          string `xml:"service_pack"`
	OSSerialNumber       string `xml:"os_serial_number"`
	OSPRoductKey         string `xml:"os_product_key"`
	LastScanTime         string `xml:"last_scan_time"`
}

// Workstation describes a computer used as a workstation
type Workstation struct {
	ID                   int    `xml:"workstationid"`
	GUID                 string `xml:"guid"`
	Name                 string `xml:"name"`
	Description          string `xml:"description"`
	AgentVersion         string `xml:"agent_version"`
	InstallDate          string `xml:"install_date"`
	LastBootTime         int    `xml:"last_boot_time"`
	Online               int    `xml:"online"`
	Active24x7           int    `xml:"active_247"`
	CheckInterval24x7    int    `xml:"check_interval_247"`
	Status24x7           int    `xml:"status_247"`
	LocalDate24x7        string `xml:"local_date_247"`
	LocalTime24x7        string `xml:"local_time_247"`
	UTCTime24x7          string `xml:"utc_time_247"`
	DscActive            int    `xml:"dsc_active"`
	AgentMode            int    `xml:"agent_mode"`
	DscHour              int    `xml:"dsc_hour"`
	DscStatus            int    `xml:"dsc_status"`
	DscLocalDate         string `xml:"dsc_local_date"`
	DscLocalTime         string `xml:"dsc_local_time"`
	DscUTCTime           string `xml:"dsc_utc_time"`
	MissFactor24x7       string `xml:"miss_factor_247"`
	EMailOverdueAlert    int    `xml:"email_overdue_alert"`
	SMSOverdueAlert      int    `xml:"sms_overdue_alert"`
	Missed24x7           int    `xml:"missed_247"`
	RemoteConnectionType int    `xml:"remote_connection_type"`
	RemoteAddress        string `xml:"remote_address"`
	RemotePort           int    `xml:"remote_port"`
	RemoteUsername       string `xml:"remote_username"`
	RemoteDomain         string `xml:"remote_domain"`
	TzBias               int    `xml:"tz_bias"`
	TzDstBias            int    `xml:"tz_dst_bias"`
	TzStdBias            int    `xml:"tz_std_bias"`
	TzMode               int    `xml:"tz_mode"`
	ATzDstDate           string `xml:"atz_dst_date"`
	TzDstDate            string `xml:"tz_dst_date"`
	TzStdDate            string `xml:"tz_std_date"`
	UTCApt               string `xml:"utc_apt"`
	UTCOffset            int    `xml:"utc_offset"`
	AssetID              int    `xml:"assetid"`
	WINSName             string `xml:"wins_name"`
	User                 string `xml:"user"`
	Domain               string `xml:"domain"`
	Role                 string `xml:"role"`
	ChassisType          int    `xml:"chassis_type"`
	Manufacturer         string `xml:"manufacturer"`
	Model                string `xml:"model"`
	DeviceSerial         string `xml:"device_serial"`
	IP                   string `xml:"ip"`
	ExternalIP           string `xml:"external_ip"`
	Mac1                 string `xml:"mac1"`
	Mac2                 string `xml:"mac2"`
	Mac3                 string `xml:"mac3"`
	ProcessorCount       int    `xml:"processor_count"`
	TotalMemory          int    `xml:"total_memory"`
	OSType               int    `xml:"os_type"`
	OS                   string `xml:"os"`
	ServicePack          string `xml:"service_pack"`
	OSSerialNumber       string `xml:"os_serial_number"`
	OSPRoductKey         string `xml:"os_product_key"`
	LastScanTime         string `xml:"last_scan_time"`
}

// DeviceAssetDetails contains details about assets of a server or workstation
type DeviceAssetDetails struct {
	Software []Software `xml:"software>items"`
}

// Software describes installed software on a server or workstation
type Software struct {
	ID          int    `xml:"softwareid"`
	CatalogID   int    `xml:"catalogid"`
	Name        string `xml:"name"`
	Version     string `xml:"version"`
	InstallDate string `xml:"install_date"`
	Deleted     int    `xml:"deleted"`
	Modified    int    `xml:"modified"`
}

// Check contains details about a check that is performed on a server or workstation
type Check struct {
	ID               int    `xml:"checkid"`
	UID              int    `xml:"uid"`
	SyncStatus       string `xml:"sync_status"`
	CheckType        int    `xml:"check_type"`
	Description      string `xml:"description"`
	Dsc24x7          int    `xml:"dsc_247"`
	StatusID         int    `xml:"statusid"`
	ConsecutiveFails int    `xml:"consecutive_fails"`
	Date             string `xml:"date"`
	Time             string `xml:"time"`
	UTCRun           string `xml:"utx_run"`
	EMail            int    `xml:"email"`
	SMS              int    `xml:"sms"`
}
