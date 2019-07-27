/**
 * @Author: huangw1
 * @Date: 2019/6/18 17:31
 */

package email

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

type Attachment struct {
	Filename string
	Data     []byte
	Inline   bool
}

type Header struct {
	Key   string
	Value string
}

type Message struct {
	From        mail.Address
	To          []string
	Cc          []string // 抄送
	Bcc         []string // 密送
	ReplyTo     string
	Subject     string
	Body        string
	ContentType string
	Headers     []Header
	Attachments map[string]*Attachment
}

func (m *Message) attach(filename string, inline bool) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_, name := filepath.Split(filename)
	m.Attachments[name] = &Attachment{
		Filename: name,
		Data:     data,
		Inline:   inline,
	}
	return nil
}

func (m *Message) Attach(filename string) error {
	return m.attach(filename, false)
}

func (m *Message) AttachInline(filename string) error {
	return m.attach(filename, true)
}

func (m *Message) AttachBuffer(filename string, data []byte, inline bool) {
	m.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}
}

func (m *Message) AddTo(to mail.Address) []string {
	m.To = append(m.To, to.String())
	return m.To
}

func (m *Message) AddCc(to mail.Address) []string {
	m.To = append(m.To, to.String())
	return m.To
}

func (m *Message) AddBcc(bcc mail.Address) []string {
	m.Bcc = append(m.Bcc, bcc.String())
	return m.Bcc
}

func (m *Message) AddHeader(key, val string) Header {
	newHeader := Header{Key: key, Value: val}
	m.Headers = append(m.Headers, newHeader)
	return newHeader
}

func (m *Message) ToList() []string {
	rcpList := make([]string, 0)
	toList, _ := mail.ParseAddressList(strings.Join(m.To, ","))
	for _, to := range toList {
		rcpList = append(rcpList, to.Address)
	}
	ccList, _ := mail.ParseAddressList(strings.Join(m.Cc, ","))
	for _, cc := range ccList {
		rcpList = append(rcpList, cc.Address)
	}
	bccList, _ := mail.ParseAddressList(strings.Join(m.Bcc, ","))
	for _, bcc := range bccList {
		rcpList = append(rcpList, bcc.Address)
	}
	return rcpList
}

func format(key, val string) string {
	return fmt.Sprintf("%s: %s\r\n", key, val)
}

func (m *Message) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(format("From", m.From.String()))
	buf.WriteString(format("Date", time.Now().Format(time.RFC1123Z)))
	buf.WriteString("To: " + strings.Join(m.To, ",") + "\r\n")
	if len(m.Cc) > 0 {
		buf.WriteString(format("Cc", strings.Join(m.Cc, ",")))
	}

	buf.WriteString(format("Subject", "=?UTF-8?B?"+base64.StdEncoding.EncodeToString([]byte(m.Subject))+"?="))
	if m.ReplyTo != "" {
		buf.WriteString(format("Reply-To", m.ReplyTo))
	}
	buf.WriteString(format("MIME-Version", "1.0"))

	if len(m.Headers) > 0 {
		for _, header := range m.Headers {
			buf.WriteString(format(header.Key, header.Value))
		}
	}

	boundary := "f46d043c813270fc6b04c2d223da"
	if len(m.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n")
		buf.WriteString("\r\n--" + boundary + "\r\n")
	}

	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\r\n\r\n", m.ContentType))
	buf.WriteString(m.Body)
	buf.WriteString("\r\n")

	if len(m.Attachments) > 0 {
		for _, attachment := range m.Attachments {
			buf.WriteString("\r\n\r\n--" + boundary + "\r\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: message/rfc822\r\n")
				buf.WriteString("Content-Disposition: inline; filename=\"" + attachment.Filename + "\"\r\n\r\n")

				buf.Write(attachment.Data)
			} else {
				ext := filepath.Ext(attachment.Filename)
				mimetype := mime.TypeByExtension(ext)
				if mimetype != "" {
					buf.WriteString(fmt.Sprintf("Content-Type: %s\r\n", mimetype))
				} else {
					buf.WriteString("Content-Type: application/octet-stream\r\n")
				}
				buf.WriteString("Content-Transfer-Encoding: base64\r\n")

				buf.WriteString("Content-Disposition: attachment; filename=\"=?UTF-8?B?")
				buf.WriteString(base64.StdEncoding.EncodeToString([]byte(attachment.Filename)))
				buf.WriteString("?=\"\r\n\r\n")

				b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
				base64.StdEncoding.Encode(b, attachment.Data)

				// write base64 content in lines of up to 76 chars
				for i, l := 0, len(b); i < l; i++ {
					buf.WriteByte(b[i])
					if (i+1)%76 == 0 {
						buf.WriteString("\r\n")
					}
				}
			}

			buf.WriteString("\r\n--" + boundary)
		}

		buf.WriteString("--")
	}
	return buf.Bytes()
}

func newMessage(subject, body, contentType string) *Message {
	message := &Message{Subject: subject, Body: body, ContentType: contentType}
	message.Attachments = make(map[string]*Attachment)
	return message
}

func NewMessage(subject, body string) *Message {
	return newMessage(subject, body, "text/plain")
}

func NewHTMLMessage(subject, body string) *Message {
	return newMessage(subject, body, "text/html")
}

func Send(addr string, auth smtp.Auth, m *Message) error {
	return smtp.SendMail(addr, auth, m.From.Address, m.ToList(), m.Bytes())
}

func dial(addr string) (*smtp.Client, error) {
	host, _, _ := net.SplitHostPort(addr)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return nil, err
	}
	return smtp.NewClient(conn, host)
}

func SendUsingTLS(addr string, auth smtp.Auth, m *Message) error {
	c, err := dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err := c.Auth(auth); err != nil {
		return err
	}

	if err := c.Mail(m.From.Address); err != nil {
		return err
	}

	for _, to := range m.To {
		err := c.Rcpt(to)
		if err != nil {
			return err
		}
	}
	for _, cc := range m.Cc {
		err := c.Rcpt(cc)
		if err != nil {
			return err
		}
	}
	for _, bcc := range m.Bcc {
		err := c.Rcpt(bcc)
		if err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(m.Bytes())
	if err != nil {
		return err
	}

	c.Quit()
	return nil
}
