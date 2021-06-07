// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package messaging

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Messenger interface {
	SendSMS(to []string, message string) error
	SendEmail(to []string, message, subject string) error
}

type Options struct {
	URL           string
	SMSReceiver   string
	EmailReceiver string
}

type client struct {
	Options
}

func NewClient(opts Options) (Messenger, error) {
	if opts.URL == "" {
		err := errors.New("client url is required.")
		return nil, err
	}

	if opts.SMSReceiver == "" {
		opts.SMSReceiver = "sms"
	}

	if opts.EmailReceiver == "" {
		opts.EmailReceiver = "email"
	}

	return &client{opts}, nil
}

func (c *client) send(t string, to []string, message, subject string) error {
	url := c.URL + "/notify"

	d := map[string]interface{}{
		"to":      to,
		"message": message,
	}
	switch t {
	case "sms":
		d["receiver"] = c.SMSReceiver
	case "email":
		d["receiver"] = c.EmailReceiver
		d["subj"] = subject
	}
	payload, _ := json.Marshal(d)

	// TODO(william): check response data
	_, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	return err
}

func (c *client) SendSMS(to []string, message string) error {
	return c.send("sms", to, message, "")
}

func (c *client) SendEmail(to []string, message, subject string) error {
	return c.send("email", to, message, subject)
}
