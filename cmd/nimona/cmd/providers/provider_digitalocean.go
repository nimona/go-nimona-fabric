package providers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

const (
	cloudInit = `#cloud-config

  coreos:
    units:
      - name: nimona.service
        command: start
        enable: true
        content: |
          [Unit]
          Description=nimona
          After=docker.service
          Requires=docker.service

          [Service]
          TimeoutStartSec=0
          Restart=always
          ExecStartPre=-/usr/bin/mkdir /etc/nimona
          ExecStartPre=-/usr/bin/chown core:core /etc/nimona
          ExecStartPre=-/usr/bin/docker stop nimona
          ExecStartPre=-/usr/bin/docker rm nimona
          ExecStartPre=-/usr/bin/docker pull quay.io/nimona/nimona:{{ .docker.tag}}
          ExecStartPre=-/usr/bin/docker run --rm \
            --name nimona \
            --user 500:500 \
            -v /etc/nimona:/etc/nimona \
            quay.io/nimona/nimona:{{ .docker.tag }} \
            daemon init \
            --data-dir /etc/nimona
          ExecStart=/usr/bin/docker run --rm \
            --name nimona \
            --user 500:500 \
            -v /etc/nimona:/etc/nimona \
            -e LOG_LEVEL=INFO \
            -p 21013:21013 \
            -p 8080:8080 \
            quay.io/nimona/nimona:{{ .docker.tag }} \
            daemon start \
            --data-dir /etc/nimona \
            --port 21013 \
            --api-port 8080 \
            --announce-hostname {{ .hostname }}

          [Install]
          WantedBy=multi-user.target`
)

// DigitalOceanProvider prodides a DO operations
type DigitalOceanProvider struct {
	client *godo.Client
}

type tokenSource struct {
	AccessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

var (
	// ErrNewinstanceTimeout is returned when NewInstances times out
	// while waiting for the instance to start
	ErrNewinstanceTimeout = errors.New(
		"timeout while waiting for new instance, please check manually")
	// ErrInvalidName is returned when a domain cannot be split in two parts
	ErrInvalidName = errors.New("invalid name for domain")
)

// NewDigitalocean creates a new DigitalOcean Provider
func NewDigitalocean(token string) (Provider, error) {
	if token == "" {
		return nil, ErrNoToken
	}

	tokenSource := &tokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	return &DigitalOceanProvider{
		client: client,
	}, nil
}

// NewInstance creates a new DO Droplet
func (dp *DigitalOceanProvider) NewInstance(dockerTag, hostname,
	sshFingerprint, size, region string) (string, error) {
	if size == "" {
		size = "s-1vcpu-1gb"
	}

	if region == "" {
		region = "lon1"
	}

	if hostname == "" {
		hostname = fmt.Sprintf("nimona-%d", time.Now().Unix())
	}

	values := map[string]interface{}{
		"docker": map[string]interface{}{
			"tag": dockerTag,
		},
		"hostname": hostname,
	}

	t, err := template.New("user-data").Parse(cloudInit)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, values); err != nil {
		return "", err
	}

	userData := tpl.String()

	ctx := context.Background()
	createRequest := &godo.DropletCreateRequest{
		Name:   hostname,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: "coreos-stable",
		},
		SSHKeys: []godo.DropletCreateSSHKey{{
			Fingerprint: sshFingerprint,
		}},
		UserData: userData,
	}

	// Create server
	drop, _, err := dp.client.Droplets.Create(ctx, createRequest)
	if err != nil {
		return "", err
	}

	wn := 0

	// Wait for the API to return an IP
	for {
		d, _, err := dp.client.Droplets.Get(ctx, drop.ID)
		if err != nil {
			return "", err
		}

		ip, err := d.PublicIPv4()
		if err != nil {
			return "", err
		}
		if ip != "" {
			return ip, nil
		}

		if wn == 60 {
			break
		}

		wn++
		time.Sleep(2 * time.Second)
	}

	return "", ErrNewinstanceTimeout
}

func (dp *DigitalOceanProvider) UpdateDomain(ctx context.Context,
	name, ip string) error {

	ds := strings.SplitN(name, ".", 2)
	if len(ds) != 2 {
		return ErrInvalidName
	}

	userSubdomain := ds[0]
	userDomain := ds[1]

	list, _, err := dp.client.Domains.List(ctx, &godo.ListOptions{})
	if err != nil {
		return err
	}

	domainFound := false
	fullPathFound := false
	record := godo.DomainRecord{}

	for _, domain := range list {
		if domain.Name == userDomain {
			domainFound = true
			break
		}
	}

	if !domainFound {
		_, _, err := dp.client.Domains.Create(ctx,
			&godo.DomainCreateRequest{
				Name: userDomain,
			})
		if err != nil {
			return err
		}
	}

	if domainFound {
		recs, _, err := dp.client.Domains.Records(ctx, userDomain,
			&godo.ListOptions{})
		if err != nil {
			return err
		}

		for _, rec := range recs {
			if rec.Name == userSubdomain {
				fullPathFound = true
				record = rec
			}
		}
	}

	if !fullPathFound {
		_, _, err := dp.client.Domains.CreateRecord(ctx, userDomain,
			&godo.DomainRecordEditRequest{
				Name: userSubdomain,
				Data: ip,
				Type: "A",
			})
		if err != nil {
			return err
		}
	}

	if fullPathFound {
		_, _, err := dp.client.Domains.EditRecord(ctx, userDomain, record.ID,
			&godo.DomainRecordEditRequest{
				Name: userSubdomain,
				Data: ip,
				Type: "A",
			})
		if err != nil {
			return err
		}
	}

	return nil
}
