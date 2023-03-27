// Client code
package proxy

import "testing"


const maxCount = 2

func Test_proxy(t *testing.T) {
	nginxServer := newNginxServer(maxCount)
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	for i := 0; i < maxCount; i++ {
		httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")

		if httpCode != 200 || body != "Ok" {
			t.Error("Wrong GET API")
		}
	}

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	if httpCode != 403 || body != "Not Allowed" {
		t.Error("Wrong GET API")
	}

	for i := 0; i < maxCount; i++ {
		httpCode, body := nginxServer.handleRequest(createUserURL, "POST")

		if httpCode != 201 || body != "User Created" {
			t.Error("Wrong POST API")
		}
	}

	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	if httpCode != 403 || body != "Not Allowed" {
		t.Error("Wrong POST API")
	}
}