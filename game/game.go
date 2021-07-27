package game

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

func Run() {

	http.HandleFunc("/clans", getClans)
	http.HandleFunc("/", getClans)
	http.HandleFunc("/player", getPlayer)

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func ReqClansInfo() string {

	req, _ := http.NewRequest("GET", "https://api.clashofclans.com/v1/clans/%239C2JQ9PC", nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", " Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjhjYmQ4MDgxLTc0ZmUtNDEzNS1hODkyLTUzM2Y0NTllYTg2NCIsImlhdCI6MTYyNzAxOTg4MCwic3ViIjoiZGV2ZWxvcGVyLzQ1MzY1OGIwLTc0NmYtNTkyYS1hMmFiLTI0Njc3MDEyNGVjZiIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjExNi4yMzMuMjExLjkyIiwiMTA2LjE0LjE5OS43NSJdLCJ0eXBlIjoiY2xpZW50In1dfQ.eDocomhMrs5zjkf3ukk7s1lVCV0TCtldGpB01KKXEUJeeuin26wUqK6Gr0D_XY3b34dUzykbld61QKvUJAdiPw")

	resp, _ := (&http.Client{}).Do(req)

	buf := make([]byte, 1024*1024)

	var res string
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
		} else {
			res = string(buf[:n])
			break
		}
	}
	return res
}

func getClans(w http.ResponseWriter, r *http.Request) {
	clansInfo := ReqClansInfo()
	var clan Clan
	if err := json.Unmarshal([]byte(clansInfo), &clan); err != nil {
		fmt.Println(err)
	}
	res, err := json.Marshal(&clan)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(res)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	var result []Member

	clansInfo := ReqClansInfo()

	var clan Clan
	if err := json.Unmarshal([]byte(clansInfo), &clan); err != nil {
		fmt.Println(err)
	}

	var cm ClanMember
	for _, cm = range clan.MemberList {

		wg.Add(1)
		tag := cm.Tag
		tag = strings.Replace(tag, "#", "%23", -1)

		go func(tag string) {
			defer wg.Done()
			var member Member
			req, _ := http.NewRequest("GET", "https://api.clashofclans.com/v1/players/"+tag, nil)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("authorization", " Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjhjYmQ4MDgxLTc0ZmUtNDEzNS1hODkyLTUzM2Y0NTllYTg2NCIsImlhdCI6MTYyNzAxOTg4MCwic3ViIjoiZGV2ZWxvcGVyLzQ1MzY1OGIwLTc0NmYtNTkyYS1hMmFiLTI0Njc3MDEyNGVjZiIsInNjb3BlcyI6WyJjbGFzaCJdLCJsaW1pdHMiOlt7InRpZXIiOiJkZXZlbG9wZXIvc2lsdmVyIiwidHlwZSI6InRocm90dGxpbmcifSx7ImNpZHJzIjpbIjExNi4yMzMuMjExLjkyIiwiMTA2LjE0LjE5OS43NSJdLCJ0eXBlIjoiY2xpZW50In1dfQ.eDocomhMrs5zjkf3ukk7s1lVCV0TCtldGpB01KKXEUJeeuin26wUqK6Gr0D_XY3b34dUzykbld61QKvUJAdiPw")

			resp, _ := (&http.Client{}).Do(req)

			buf := make([]byte, 1024*1024)

			var res string
			for {
				// 接收服务端信息
				n, err := resp.Body.Read(buf)
				if err != nil && err != io.EOF {
				} else {
					res = string(buf[:n])
					break
				}
			}
			json.Unmarshal([]byte(res), &member)
			result = append(result, member)
		}(tag)
	}

	wg.Wait()
	res, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(res)
}
