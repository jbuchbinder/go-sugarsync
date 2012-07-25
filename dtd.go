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

// https://www.sugarsync.com/dev/api/method/get-user-info.html
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

// https://www.sugarsync.com/dev/api/method/get-folders.html
type CollectionContents struct {
	Collections []Collection `xml:"collectionContents" json:"collections"`
	HasMore     bool         `xml:"hasMore,attr" json:"hasMore"`
	Start       int64        `xml:"start,attr" json:"start"`
	End         int64        `xml:"end,attr" json:"end"`
}

type Collection struct {
	Type        string `xml:"type,attr" json:"type"`
	DisplayName string `xml:"displayName" json:"displayName"`
	RefLink     string `xml:"ref" json:"ref"`
	IconId      int64  `xml:"iconId" json:"iconId"`
	Contents    string `xml:"contents" json:"contents"`
}

// https://www.sugarsync.com/dev/api/method/get-contacts.html
type Contact struct {
	PrimaryEmailAddress string `xml:"primaryEmailAddress" json:"primaryEmailAddress"`
	FirstName           string `xml:"firstName" json:"firstName"`
	LastName            string `xml:"lastName" json:"lastName"`
}

// https://www.sugarsync.com/dev/api/method/get-received-shares-list.html
type ReceviedShare struct {
	RefLink      string `xml:"ref,attr" json:"ref"`
	DisplayName  string `xml:"displayName" json:"displayName"`
	TimeReceived string `xml:"timeReceived" json:"timeReceived"`
	SharedFolder string `xml:"sharedFolder" json:"sharedFolder"`
	Owner        string `xml:"owner" json:"owner"`
}

// https://www.sugarsync.com/dev/api/method/get-version-history.html
type FileVersion struct {
	Size            string `xml:"size" json:"size"`
	LastModified    string `xml:"lastModified" json:"lastModified"`
	MediaType       string `xml:"mediaType" json:"mediaType"`
	PresentOnServer bool   `xml:"presentOnServer" json:"presentOnServer"`
	FileData        string `xml:"fileData" json:"fileData"`
	RefLink         string `xml:"ref" json:"ref"`
}

// https://www.sugarsync.com/dev/api/method/get-photo-album.html
type Album struct {
	DisplayName string `xml:"displayName" json:"displayName"`
	DsId        string `xml:"dsid" json:"dsid"`
	Files       string `xml:"files" json:"files"`
	Contents    string `xml:"contents" json:"contents"`
}
