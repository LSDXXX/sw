package shenwu

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var sauthor string
var create_point time.Time

type Coord struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Word struct {
	Character  string
	Confidence float32
}
type Item struct {
	Itemstring string
	Itemcoord  Coord
	Words      []Word
}
type SData struct {
	Session_id string
	Items      []Item
}
type Response struct {
	Code    int
	Message string
	Data    SData
}

//post image to tencent ocr

func postToTencent(file_path string) (Response, error) {
	ct, _, err := createHttpContent(file_path)
	var out Response
	if err != nil {
		return out, err
	}
	req, err := http.NewRequest("POST", "http://recognition.image.myqcloud.com/ocr/general", ct)
	if err != nil {
		return out, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "recognition.image.myqcloud.com")
	req.Close = true
	if time.Since(create_point).Hours() >= 24.0 {
		generateAuthorization()
	}
	req.Header.Add("Authorization", sauthor)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return out, err
	}
	json_str, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	err = json.Unmarshal(json_str, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func createHttpContent(file_path string) (io.Reader, int, error) {
	file_content, err := ioutil.ReadFile(file_path)
	if err != nil {
		return bytes.NewReader([]byte("")), 0, err
	}
	base64_img := base64.StdEncoding.EncodeToString(file_content)
	json_map := map[string]string{
		"appid": "1257343604",
		"image": base64_img,
	}
	json_str, _ := json.Marshal(json_map)
	ioutil.WriteFile("2.txt", json_str, 0755)
	return bytes.NewReader(json_str), len(json_str), err
}

func generateAuthorization() {
	appid := "1257343604"
	secret_id := "AKIDzJYdaXnEszBC2mB6X1NYYQ6bu23unIrF"
	secret_key := "QhokVoz30oRvRnCUbeVAew31udBZpu06"
	buffer := bytes.NewBuffer([]byte(""))
	r := rand.Uint32()
	create_point = time.Now()
	fmt.Fprintf(buffer, "a=%s&k=%s&t=%d&e=%d&r=%d", appid, secret_id, int(create_point.Unix()), int(create_point.Unix()+int64(24*time.Hour.Seconds())), r)
	mac := hmac.New(sha1.New, []byte(secret_key))
	//tmp := bytes.NewBuffer([]byte(""))
	mac.Write([]byte(buffer.String()))
	//fmt.Fprintf(tmp, "%s", hex.EncodeToString(mac.Sum(nil)))
	sauthor = string(mac.Sum(nil)) + buffer.String()
	sauthor = base64.StdEncoding.EncodeToString([]byte(sauthor))
}
