package p

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ToniChawatphon/proxy-line-noti-function/app"
)

// ProxyLineNoti recieve an alert from Tradind View
// and for the alert to line notify
func ProxyLineNoti(w http.ResponseWriter, r *http.Request) {
	app.InitSetting()
	body := app.WriteRequest(r)
	res := app.Noti.SendNotification(string(body))
	status := fmt.Sprintf("OK %s", strconv.Itoa((res.StatusCode)))
	log.Println(status)
}
