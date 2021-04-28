package model

type MachineReview struct {
	Sign                   string              `json:"sign,omitempty"`
	ProductID              string              `json:"productId,omitempty"`
	ProductName            string              `json:"productName,omitempty"`
	Source                 interface{}         `json:"source,omitempty"`
	OrderID                string              `json:"orderId,omitempty"`
	IDNo                   string              `json:"idNo,omitempty"`
	Phone                  string              `json:"phone,omitempty"`
	SubmitTime             string              `json:"submitTime,omitempty"`
	CallbackURL            interface{}         `json:"callbackUrl,omitempty"`
	Callback               interface{}         `json:"callback,omitempty"`
	LivingImg              interface{}         `json:"livingImg,omitempty"`
	FaceCompareScore       string              `json:"faceCompareScore,omitempty"`
	ChannelSource          string              `json:"channelSource,omitempty"`
	MerchantID             string              `json:"merchantId,omitempty"`
	Marital                string              `json:"marital,omitempty"`
	RealName               string              `json:"realName,omitempty"`
	Profession             string              `json:"profession,omitempty"`
	UserLoans              []*UserLoans        `json:"userLoans,omitempty"`
	LoanTerm               string              `json:"loanTerm,omitempty"`
	LoanAmount             string              `json:"loanAmount,omitempty"`
	LoanTermUnit           string              `json:"loanTermUnit,omitempty"`
	InterestAmount         interface{}         `json:"interestAmount,omitempty"`
	InterestRate           interface{}         `json:"interestRate,omitempty"`
	ServiceFee             interface{}         `json:"serviceFee,omitempty"`
	IsDelaySupported       bool                `json:"isDelaySupported,omitempty"`
	IsReloan               bool                `json:"isReloan,omitempty"`
	DebitBankCard          string              `json:"debitBankCard,omitempty"`
	DebitOpenBankID        string              `json:"debitOpenBankId,omitempty"`
	DebitBankName          string              `json:"debitBankName,omitempty"`
	DeviceInfo             *DeviceInfo         `json:"deviceInfo,omitempty"`
	UserEmerContacts       []*UserEmerContacts `json:"userEmerContacts,omitempty"`
	Status                 string              `json:"status,omitempty"`
	RepaymentedTime        interface{}         `json:"repaymentedTime,omitempty"`
	RepaymentTime          string              `json:"repaymentTime,omitempty"`
	ExpiryTime             interface{}         `json:"expiryTime,omitempty"`
	LoanTime               string              `json:"loanTime,omitempty"`
	ActualAmount           interface{}         `json:"actualAmount,omitempty"`
	RepaymentAmount        interface{}         `json:"repaymentAmount,omitempty"`
	ExpiryAmount           interface{}         `json:"expiryAmount,omitempty"`
	CurrentRepaymentAmount interface{}         `json:"currentRepaymentAmount,omitempty"`
}
type RepaymentPlans struct {
	Status            interface{} `json:"status,omitempty"`
	Term              interface{} `json:"term,omitempty"`
	RepaymentPlanTime string      `json:"repaymentPlanTime,omitempty"`
	RealRepaymentTime interface{} `json:"realRepaymentTime,omitempty"`
}
type UserLoans struct {
	UserIdentity   interface{}      `json:"userIdentity,omitempty"`
	LoanIdentity   interface{}      `json:"loanIdentity,omitempty"`
	Status         string           `json:"status,omitempty"`
	CreateTime     string           `json:"createTime,omitempty"`
	MerchantNo     interface{}      `json:"merchantNo,omitempty"`
	RegisterDate   interface{}      `json:"registerDate,omitempty"`
	RepaymentPlans []RepaymentPlans `json:"repaymentPlans,omitempty"`
}
type PublicIP struct {
	SecondIP interface{} `json:"second_ip,omitempty"`
	FirstIP  string      `json:"first_ip,omitempty"`
}
type OtherData struct {
	RootJailbreak string      `json:"root_jailbreak,omitempty"`
	LastBootTime  interface{} `json:"last_boot_time,omitempty"`
	Keyboard      interface{} `json:"keyboard,omitempty"`
	Simulator     string      `json:"simulator,omitempty"`
	Dbm           interface{} `json:"dbm,omitempty"`
}
type Storage struct {
	RAMTotalSize          string      `json:"ram_total_size,omitempty"`
	RAMUsableSize         string      `json:"ram_usable_size,omitempty"`
	MemoryCardSize        string      `json:"memory_card_size,omitempty"`
	MemoryCardSizeUse     string      `json:"memory_card_size_use,omitempty"`
	InternalStorageUsable interface{} `json:"internal_storage_usable,omitempty"`
	InternalStorageTotal  string      `json:"internal_storage_total,omitempty"`
}
type BatteryStatus struct {
	BatteryPct  string `json:"battery_pct,omitempty"`
	IsUsbCharge bool   `json:"is_usb_charge,omitempty"`
	IsCharging  bool   `json:"is_charging,omitempty"`
	IsAcCharge  bool   `json:"is_ac_charge,omitempty"`
}
type CurrentWifi struct {
	Bssid string      `json:"bssid,omitempty"`
	Name  string      `json:"name,omitempty"`
	Ssid  interface{} `json:"ssid,omitempty"`
	Mac   interface{} `json:"mac,omitempty"`
}
type ConfiguredWifi struct {
	Bssid string `json:"bssid,omitempty"`
	Name  string `json:"name,omitempty"`
	Ssid  string `json:"ssid,omitempty"`
	Mac   string `json:"mac,omitempty"`
}
type Network struct {
	CurrentWifi    CurrentWifi      `json:"current_wifi,omitempty"`
	IP             string           `json:"IP,omitempty"`
	ConfiguredWifi []ConfiguredWifi `json:"configured_wifi,omitempty"`
}
type Application struct {
	AppName     string `json:"app_name,omitempty"`
	Package     string `json:"package,omitempty"`
	InTime      int64  `json:"in_time,omitempty"`
	VersionName string `json:"version_name,omitempty"`
	AppType     int    `json:"app_type,omitempty"`
	VersionCode int    `json:"version_code,omitempty"`
	Flags       int    `json:"flags,omitempty"`
	UpTime      int64  `json:"up_time,omitempty"`
}
type Contact struct {
	LastTimeContacted  string `json:"last_time_contacted,omitempty"`
	Number             string `json:"number,omitempty"`
	ContactDisplayName string `json:"contact_display_name,omitempty"`
	UpTime             string `json:"up_time,omitempty"`
	TimesContacted     string `json:"times_contacted,omitempty"`
}
type GeneralData struct {
	AndID                 interface{} `json:"and_id,omitempty"`
	PhoneType             interface{} `json:"phone_type,omitempty"`
	Gaid                  string      `json:"gaid,omitempty"`
	Language              interface{} `json:"language,omitempty"`
	Mac                   string      `json:"mac,omitempty"`
	LocaleIso3Language    interface{} `json:"locale_iso_3_language,omitempty"`
	LocaleDisplayLanguage interface{} `json:"locale_display_language,omitempty"`
	Imei                  string      `json:"imei,omitempty"`
	PhoneNumber           interface{} `json:"phone_number,omitempty"`
	NetworkOperatorName   interface{} `json:"network_operator_name,omitempty"`
	NetworkType           string      `json:"network_type,omitempty"`
	TimeZoneID            interface{} `json:"time_zone_id,omitempty"`
	LocaleIso3Country     interface{} `json:"locale_iso_3_country,omitempty"`
}
type Gps struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
type Location struct {
	Gps Gps `json:"gps,omitempty"`
}
type Hardware struct {
	DeviceName   string      `json:"device_name,omitempty"`
	Release      interface{} `json:"release,omitempty"`
	SdkVersion   interface{} `json:"sdk_version,omitempty"`
	Model        interface{} `json:"model,omitempty"`
	SerialNumber interface{} `json:"serial_number,omitempty"`
	PhysicalSize interface{} `json:"physical_size,omitempty"`
	Brand        interface{} `json:"brand,omitempty"`
}
type DeviceInfo struct {
	VideoExternal  string        `json:"video_external,omitempty"`
	PublicIP       *PublicIP      `json:"public_ip,omitempty"`
	OtherData      *OtherData     `json:"other_data,omitempty"`
	AudioExternal  string        `json:"audio_external,omitempty"`
	Storage        *Storage       `json:"storage,omitempty"`
	ImagesInternal string        `json:"images_internal,omitempty"`
	BatteryStatus  *BatteryStatus `json:"battery_status,omitempty"`
	Network        *Network       `json:"network,omitempty"`
	ContactGroup   string        `json:"contact_group,omitempty"`
	ImagesExternal string        `json:"images_external,omitempty"`
	Application    []*Application `json:"application,omitempty"`
	Call           interface{}   `json:"call,omitempty"`
	VideoInternal  string        `json:"video_internal,omitempty"`
	Contact        []*Contact     `json:"contact,omitempty"`
	GeneralData    *GeneralData   `json:"general_data,omitempty"`
	From           string        `json:"from,omitempty"`
	Location       *Location      `json:"location,omitempty"`
	AudioInternal  string        `json:"audio_internal,omitempty"`
	DownloadFiles  string        `json:"download_files,omitempty"`
	Hardware       *Hardware      `json:"hardware,omitempty"`
}
type UserEmerContacts struct {
	EmergencyName     string `json:"emergencyName,omitempty"`
	EmergencyRelation int    `json:"emergencyRelation,omitempty"`
	EmergencyPhone    string `json:"emergencyPhone,omitempty"`
}
