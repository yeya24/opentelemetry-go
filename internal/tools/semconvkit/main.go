// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package semconvkit is used to generate opentelemetry-go specific semantic
// convention code. It is expected to be used in with the semconvgen utility
// (go.opentelemetry.io/build-tools/semconvgen) to completely generate
// versioned sub-packages of go.opentelemetry.io/otel/semconv.
package main

import (
	"embed"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	out = flag.String("output", "./", "output directory")
	tag = flag.String("tag", "", "OpenTelemetry tagged version")

	//go:embed templates/*.tmpl
	rootFS embed.FS
)

// SemanticConventions are information about the semantic conventions being
// generated.
type SemanticConventions struct {
	// TagVer is the tagged version (i.e. v1.7.0 and not 1.7.0).
	TagVer string
}

func (sc SemanticConventions) SemVer() string {
	return strings.TrimPrefix(*tag, "v")
}

// render renders all templates to the dest directory using the data.
func render(src, dest string, data *SemanticConventions) error {
	tmpls, err := template.ParseFS(rootFS, src)
	if err != nil {
		return err
	}
	for _, tmpl := range tmpls.Templates() {
		target := filepath.Join(dest, strings.TrimSuffix(tmpl.Name(), ".tmpl"))
		wr, err := os.Create(target)
		if err != nil {
			return err
		}

		err = tmpl.Execute(wr, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	if *tag == "" {
		log.Fatalf("invalid tag: %q", *tag)
	}

	sc := &SemanticConventions{TagVer: *tag}

	if err := render("templates/*.tmpl", *out, sc); err != nil {
		log.Fatal(err)
	}
}
