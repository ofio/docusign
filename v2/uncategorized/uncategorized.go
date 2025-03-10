// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package uncategorized implements the DocuSign SDK
// category Uncategorized.
//
//
// Uncategorized calls may change or move to other packages.
//
//
// Usage example:
//
//   import (
//       "github.com/ofio/esign"
//       "github.com/ofio/esign/v2/uncategorized"
//       "github.com/ofio/esign/v2/model"
//   )
//   ...
//   uncategorizedService := uncategorized.New(esignCredential)
package uncategorized // import "github.com/ofio/esign/v2/uncategorized"

import (
	"context"
	"net/url"
	"strings"

	"github.com/ofio/esign"
	"github.com/ofio/esign/v2/model"
)

// Service implements DocuSign Uncategorized Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a uncategorized service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// CommentsGetis uncategorized and subject to change
func (s *Service) CommentsGet(envelopeID string) *CommentsGetOp {
	return &CommentsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"envelopes", envelopeID, "comments", "transcript"}, "/"),
		Accept:     "application/pdf",
		QueryOpts:  make(url.Values),
	}
}

// CommentsGetOp implements DocuSign API SDK Uncategorized::getCommentsTranscript
type CommentsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CommentsGetOp) Do(ctx context.Context) (*esign.Download, error) {
	var res *esign.Download
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Encoding set the call query parameter encoding
func (op *CommentsGetOp) Encoding(val string) *CommentsGetOp {
	if op != nil {
		op.QueryOpts.Set("encoding", val)
	}
	return op
}

// DocumentResponsiveHTMLPreviewCreateis uncategorized and subject to change
func (s *Service) DocumentResponsiveHTMLPreviewCreate(documentID string, envelopeID string, documentHTMLDefinition *model.DocumentHTMLDefinition) *DocumentResponsiveHTMLPreviewCreateOp {
	return &DocumentResponsiveHTMLPreviewCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"envelopes", envelopeID, "documents", documentID, "responsive_html_preview"}, "/"),
		Payload:    documentHTMLDefinition,
		QueryOpts:  make(url.Values),
	}
}

// DocumentResponsiveHTMLPreviewCreateOp implements DocuSign API SDK Uncategorized::createDocumentResponsiveHtmlPreview
type DocumentResponsiveHTMLPreviewCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DocumentResponsiveHTMLPreviewCreateOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitions, error) {
	var res *model.DocumentHTMLDefinitions
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeDocumentHTMLDefinitionsGetis uncategorized and subject to change
func (s *Service) EnvelopeDocumentHTMLDefinitionsGet(documentID string, envelopeID string) *EnvelopeDocumentHTMLDefinitionsGetOp {
	return &EnvelopeDocumentHTMLDefinitionsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"envelopes", envelopeID, "documents", documentID, "html_definitions"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// EnvelopeDocumentHTMLDefinitionsGetOp implements DocuSign API SDK Uncategorized::getEnvelopeDocumentHtmlDefinitions
type EnvelopeDocumentHTMLDefinitionsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeDocumentHTMLDefinitionsGetOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitionOriginals, error) {
	var res *model.DocumentHTMLDefinitionOriginals
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeHTMLDefinitionsListis uncategorized and subject to change
func (s *Service) EnvelopeHTMLDefinitionsList(envelopeID string) *EnvelopeHTMLDefinitionsListOp {
	return &EnvelopeHTMLDefinitionsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"envelopes", envelopeID, "html_definitions"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// EnvelopeHTMLDefinitionsListOp implements DocuSign API SDK Uncategorized::getEnvelopeHtmlDefinitions
type EnvelopeHTMLDefinitionsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeHTMLDefinitionsListOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitionOriginals, error) {
	var res *model.DocumentHTMLDefinitionOriginals
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// ResponsiveHTMLPreviewCreateis uncategorized and subject to change
func (s *Service) ResponsiveHTMLPreviewCreate(envelopeID string, documentHTMLDefinition *model.DocumentHTMLDefinition) *ResponsiveHTMLPreviewCreateOp {
	return &ResponsiveHTMLPreviewCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"envelopes", envelopeID, "responsive_html_preview"}, "/"),
		Payload:    documentHTMLDefinition,
		QueryOpts:  make(url.Values),
	}
}

// ResponsiveHTMLPreviewCreateOp implements DocuSign API SDK Uncategorized::createResponsiveHtmlPreview
type ResponsiveHTMLPreviewCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ResponsiveHTMLPreviewCreateOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitions, error) {
	var res *model.DocumentHTMLDefinitions
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateDocumentHTMLDefinitionsListis uncategorized and subject to change
func (s *Service) TemplateDocumentHTMLDefinitionsList(documentID string, templateID string) *TemplateDocumentHTMLDefinitionsListOp {
	return &TemplateDocumentHTMLDefinitionsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"templates", templateID, "documents", documentID, "html_definitions"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// TemplateDocumentHTMLDefinitionsListOp implements DocuSign API SDK Uncategorized::getTemplateDocumentHtmlDefinitions
type TemplateDocumentHTMLDefinitionsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateDocumentHTMLDefinitionsListOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitionOriginals, error) {
	var res *model.DocumentHTMLDefinitionOriginals
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateDocumentResponsiveHTMLPreviewCreateis uncategorized and subject to change
func (s *Service) TemplateDocumentResponsiveHTMLPreviewCreate(documentID string, templateID string, documentHTMLDefinition *model.DocumentHTMLDefinition) *TemplateDocumentResponsiveHTMLPreviewCreateOp {
	return &TemplateDocumentResponsiveHTMLPreviewCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"templates", templateID, "documents", documentID, "responsive_html_preview"}, "/"),
		Payload:    documentHTMLDefinition,
		QueryOpts:  make(url.Values),
	}
}

// TemplateDocumentResponsiveHTMLPreviewCreateOp implements DocuSign API SDK Uncategorized::createTemplateDocumentResponsiveHtmlPreview
type TemplateDocumentResponsiveHTMLPreviewCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateDocumentResponsiveHTMLPreviewCreateOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitions, error) {
	var res *model.DocumentHTMLDefinitions
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateHTMLDefinitionsListis uncategorized and subject to change
func (s *Service) TemplateHTMLDefinitionsList(templateID string) *TemplateHTMLDefinitionsListOp {
	return &TemplateHTMLDefinitionsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"templates", templateID, "html_definitions"}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// TemplateHTMLDefinitionsListOp implements DocuSign API SDK Uncategorized::getTemplateHtmlDefinitions
type TemplateHTMLDefinitionsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateHTMLDefinitionsListOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitionOriginals, error) {
	var res *model.DocumentHTMLDefinitionOriginals
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateResponsiveHTMLPreviewCreateis uncategorized and subject to change
func (s *Service) TemplateResponsiveHTMLPreviewCreate(templateID string, documentHTMLDefinition *model.DocumentHTMLDefinition) *TemplateResponsiveHTMLPreviewCreateOp {
	return &TemplateResponsiveHTMLPreviewCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"templates", templateID, "responsive_html_preview"}, "/"),
		Payload:    documentHTMLDefinition,
		QueryOpts:  make(url.Values),
	}
}

// TemplateResponsiveHTMLPreviewCreateOp implements DocuSign API SDK Uncategorized::createTemplateResponsiveHtmlPreview
type TemplateResponsiveHTMLPreviewCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateResponsiveHTMLPreviewCreateOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitions, error) {
	var res *model.DocumentHTMLDefinitions
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
