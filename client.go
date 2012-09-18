// GO-SUGARSYNC
// https://github.com/jbuchbinder/go-sugarsync

package sugarsync

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

type SugarsyncClient struct {
	Username     string
	Password     string
	AuthToken    string
	RefreshToken string
	UserResource string
	Debug        bool
}

func (self *SugarsyncClient) GetAuthToken() (err error) {
	if self.RefreshToken == "" {
		err = errors.New("No valid refresh token.")
		return
	}
	client := http.Client{}
	payload := "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n" +
		"<tokenAuthRequest>\n" +
		"<accessKeyId>" + ACCESS_KEY_ID + "</accessKeyId>\n" +
		"<privateAccessKey>" + PRIVATE_ACCESS_KEY + "</privateAccessKey>\n" +
		"<refreshToken>" + self.RefreshToken + "</refreshToken>\n" +
		"</tokenAuthRequest>\n"
	req, e := http.NewRequest("POST", AUTHORIZATION_URL, strings.NewReader(string(payload)))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/xml; charset=UTF-8")

	if self.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()

	if self.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}

	// Extract user resource from body
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return e
	}
	var obj AuthorizationResponse
	e = xml.Unmarshal(body, &obj)
	if e != nil {
		return e
	}
	self.UserResource = obj.User
	self.AuthToken = res.Header.Get("Location")
	return
}

func (self *SugarsyncClient) GetRefreshToken() (err error) {
	if self.Username == "" || self.Password == "" {
		err = errors.New("Username and password must be set to get refresh token.")
		return
	}
	client := http.Client{}
	payload := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
		"<appAuthorization>\n" +
		"<username>" + self.Username + "</username>\n" +
		"<password>" + self.Password + "</password>\n" +
		"<application>" + APP_ID + "</application>\n" +
		"<accessKeyId>" + ACCESS_KEY_ID + "</accessKeyId>\n" +
		"<privateAccessKey>" + PRIVATE_ACCESS_KEY + "</privateAccessKey>\n" +
		"</appAuthorization>\n"
	if self.Debug {
		fmt.Println("Posting to " + REFRESH_URL + " with:\n" + payload)
	}
	req, e := http.NewRequest("POST", REFRESH_URL, strings.NewReader(string(payload)))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/xml; charset=UTF-8")
	req.SetBasicAuth(self.Username, self.Password)

	if self.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()

	if self.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}

	_, e = ioutil.ReadAll(res.Body)
	if e != nil {
		return e
	}
	self.RefreshToken = res.Header.Get("Location")
	return
}

func (self *SugarsyncClient) GetNewFileLocation(folder string, fileName string) (loc string, e error) {
	if self.AuthToken == "" {
		e = errors.New("Auth token must be retrieved first")
		return
	}
	client := http.Client{}
	payload := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
		"<file>\n" +
		"<displayName>" + fileName + "</displayName>\n" +
		"<mediaType>application/octet-stream</mediaType>\n" +
		"</file>\n"
	if self.Debug {
		fmt.Println("Posting to " + REFRESH_URL + " with:\n" + payload)
	}
	req, err := http.NewRequest("POST", folder, strings.NewReader(string(payload)))
	if err != nil {
		e = err
		return
	}
	req.Header.Set("Authorization", self.AuthToken)

	if self.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	res, err := client.Do(req)
	if err != nil {
		e = err
		return
	}
	defer res.Body.Close()

	if self.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}

	loc = res.Header.Get("Location")
	return
}

func (self *SugarsyncClient) GetUserInfo(authToken string, userResource string) (ui UserInfo, err error) {
	if self.AuthToken == "" {
		err = errors.New("Auth token must be retrieved first")
		return
	}
	client := http.Client{}
	req, _ := http.NewRequest("GET", userResource, nil)
	req.Header.Set("Authorization", authToken)

	if self.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	res, e := client.Do(req)
	if e != nil {
		err = e
		return
	}
	defer res.Body.Close()

	if self.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}

	// Extract user resource from body
	//io.Copy(os.Stderr, res.Body)
	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		err = e
		return
	}
	var obj UserInfo
	e = xml.Unmarshal(body, &obj)
	if e != nil {
		err = e
		return
	}

	ui = obj
	return
}

func (self *SugarsyncClient) UploadFile(fileLocation string, file string) (e error) {
	if self.AuthToken == "" {
		e = errors.New("Auth token must be retrieved first")
		return
	}
	client := http.Client{}
	fData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fileLocation+"/data", strings.NewReader(string(fData)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", self.AuthToken)
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Length", fmt.Sprint(len(fData)))

	if self.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if self.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}

	return
}

func (self *SugarsyncClient) CreateNewFolder(folder string, folderName string) (f string, e error) {
	if self.AuthToken == "" {
		e = errors.New("Auth token must be retrieved first")
		return
	}
	client := http.Client{}
	rObj := Folder{DisplayName: folderName}
	payload, err := xml.Marshal(rObj)
	if self.Debug {
		fmt.Println("Posting to " + folder + " with:\n" + string(payload))
	}
	req, err := http.NewRequest("POST", folder, strings.NewReader(string(payload)))
	req.Header.Set("Authorization", self.AuthToken)

	if self.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	res, err := client.Do(req)
	if err != nil {
		e = err
		return
	}
	defer res.Body.Close()

	if self.Debug {
		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))
	}

	f = res.Header.Get("Location")
	return
}
