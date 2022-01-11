package proxylinenotifunction

import (
	"net/http"

	"github.com/ToniChawatphon/proxy-line-noti-function/app"
)

// ProxyLineNoti recieve an alert from Tradind View
// and forward the alert to line notify
func ProxyLineNoti(w http.ResponseWriter, r *http.Request) {
	app.InitSetting()
	app.Noti.SendNotification(r)
}
