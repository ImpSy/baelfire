package main

import (
	"errors"

	"github.com/anaskhan96/soup"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

type target struct {
	Name    string `json:"name" gorm:"PRIMARY_KEY"`
	Project string `json:"project" gorm:"NOT NULL"`
	Host    string `json:"host" gorm:"NOT NULL"`
}

func (t target) getVersion() (string, error) {
	switch t.Project {
	case "alertmanager":
		return getAlertmanagerVersion(t)
	case "baelfire":
		return getBaelfireVersion(t)
	case "grafana":
		return getGrafanaVersion(t)
	case "metabase":
		return getMetabaseVersion(t)
	case "prometheus":
		return getPrometheusVersion(t)
	default:
		return "", errors.New("target.project: '" + t.Project + "' is unknown")
	}
}

func getAlertmanagerVersion(t target) (string, error) {
	r, err := req.Get(t.Host + "/api/v1/status")
	if err != nil {
		return "", err
	}

	result := gjson.Get(r.String(), "data.versionInfo.version")
	if result.Exists() != true {
		return "", errors.New("version not found on Alertmanager api response")
	}
	return result.String(), nil
}

func getBaelfireVersion(t target) (string, error) {
	r, err := req.Get(t.Host + "/api/v1/version")
	if err != nil {
		return "", err
	}

	result := gjson.Get(r.String(), "version")
	if result.Exists() != true {
		return "", errors.New("version not found on Baelfire api response")
	}
	return result.String(), nil
}

func getGrafanaVersion(t target) (string, error) {
	r, err := req.Get(t.Host + "/api/health")
	if err != nil {
		return "", err
	}

	result := gjson.Get(r.String(), "version")
	if result.Exists() != true {
		return "", errors.New("version not found on Grafana api response")
	}
	return result.String(), nil
}

func getMetabaseVersion(t target) (string, error) {
	r, err := req.Get(t.Host + "/public")
	if err != nil {
		return "", err
	}

	doc := soup.HTMLParse(r.String())
	json := doc.Find("script", "id", "_metabaseBootstrap").Text()
	result := gjson.Get(json, "version.tag")
	if result.Exists() != true {
		return "", errors.New("version not found on Metabase api response")
	}
	return result.String(), nil
}

func getPrometheusVersion(t target) (string, error) {
	r, err := req.Get(t.Host + "/version")
	if err != nil {
		return "", err
	}

	result := gjson.Get(r.String(), "version")
	if result.Exists() != true {
		return "", errors.New("version not found on Prometheus api response")
	}
	return result.String(), nil
}
