// GO-SUGARSYNC
// https://github.com/jbuchbinder/go-sugarsync

package sugarsync

const (
	ACCESS_KEY_ID      = "MzI1MzA4MjEzNDI4MTA0NzczNjI"
	APP_ID             = "/sc/3253082/350_101900713"
	PRIVATE_ACCESS_KEY = "OWFlMmI0MWI0MjljNGNkMGJiNzFlNjM1NjUwNTZlODU"

	AUTHORIZATION_URL = "https://api.sugarsync.com/authorization"
	REFRESH_URL       = "https://api.sugarsync.com/app-authorization"
)

type AuthorizationResponse struct {
	Expiration string `xml:"expiration"`
	User       string `xml:"user"`
}

type UserInfo struct {
	Username         string `xml:"username" json:"username"`
	Nickname         string `xml:"nickname" json:"nickname"`
	Workspaces       string `xml:"workspaces" json:"workspaces"`
	SyncFolders      string `xml:"syncfolders" json:"syncfolders"`
	Deleted          string `xml:"deleted" json:"deleted"`
	MagicBriefcase   string `xml:"magicBriefcase" json:"magicBriefcase"`
	WebArchive       string `xml:"webArchive" json:"webArchive"`
	MobilePhotos     string `xml:"mobilePhotos" json:"mobilePhotos"`
	ReceivedShares   string `xml:"receivedShares" json:"receivedShares"`
	Contacts         string `xml:"contacts" json:"contacts"`
	Albums           string `xml:"albums" json:"albums"`
	RecentActivities string `xml:"recentActivities" json:"recentActivities"`
	PublicLinks      string `xml:"publicLinks" json:"publicLinks"`
}
