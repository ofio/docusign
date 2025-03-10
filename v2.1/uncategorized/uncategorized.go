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
//       "github.com/ofio/esign/v2.1/uncategorized"
//       "github.com/ofio/esign/v2.1/model"
//   )
//   ...
//   uncategorizedService := uncategorized.New(esignCredential)
package uncategorized // import "github.com/ofio/esign/v2.1/uncategorized"

import (
	"context"
	"net/url"
	"strings"

	"github.com/ofio/esign"
	"github.com/ofio/esign/v2.1/model"
)

// Service implements DocuSign Uncategorized Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a uncategorized service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// CommentsCreateEnvelopeCommentsis uncategorized and subject to change
func (s *Service) CommentsCreateEnvelopeComments(envelopeID string, commentsPublish *model.CommentsPublish) *CommentsCreateEnvelopeCommentsOp {
	return &CommentsCreateEnvelopeCommentsOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"envelopes", envelopeID, "comments"}, "/"),
		Payload:    commentsPublish,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// CommentsCreateEnvelopeCommentsOp implements DocuSign API SDK Uncategorized::createEnvelopeComments
type CommentsCreateEnvelopeCommentsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CommentsCreateEnvelopeCommentsOp) Do(ctx context.Context) (*model.CommentHistoryResult, error) {
	var res *model.CommentHistoryResult
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// CommentsGetis uncategorized and subject to change
func (s *Service) CommentsGet(envelopeID string) *CommentsGetOp {
	return &CommentsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"envelopes", envelopeID, "comments", "transcript"}, "/"),
		Accept:     "application/pdf",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
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
		Version:    esign.VersionV21,
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
		Version:    esign.VersionV21,
	}
}

// EnvelopeDocumentHTMLDefinitionsGetOp implements DocuSign API SDK Uncategorized::getEnvelopeDocumentHtmlDefinitions
type EnvelopeDocumentHTMLDefinitionsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeDocumentHTMLDefinitionsGetOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitionOriginals, error) {
	var res *model.DocumentHTMLDefinitionOriginals
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeDocumentTabsCreateDocumentTabsis uncategorized and subject to change
func (s *Service) EnvelopeDocumentTabsCreateDocumentTabs(documentID string, envelopeID string, templateRecipientTabs *model.Tabs) *EnvelopeDocumentTabsCreateDocumentTabsOp {
	return &EnvelopeDocumentTabsCreateDocumentTabsOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"envelopes", envelopeID, "documents", documentID, "tabs"}, "/"),
		Payload:    templateRecipientTabs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeDocumentTabsCreateDocumentTabsOp implements DocuSign API SDK Uncategorized::createDocumentTabs
type EnvelopeDocumentTabsCreateDocumentTabsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeDocumentTabsCreateDocumentTabsOp) Do(ctx context.Context) (*model.Tabs, error) {
	var res *model.Tabs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeDocumentTabsDeleteDocumentTabsis uncategorized and subject to change
func (s *Service) EnvelopeDocumentTabsDeleteDocumentTabs(documentID string, envelopeID string, templateRecipientTabs *model.Tabs) *EnvelopeDocumentTabsDeleteDocumentTabsOp {
	return &EnvelopeDocumentTabsDeleteDocumentTabsOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"envelopes", envelopeID, "documents", documentID, "tabs"}, "/"),
		Payload:    templateRecipientTabs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeDocumentTabsDeleteDocumentTabsOp implements DocuSign API SDK Uncategorized::deleteDocumentTabs
type EnvelopeDocumentTabsDeleteDocumentTabsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeDocumentTabsDeleteDocumentTabsOp) Do(ctx context.Context) (*model.Tabs, error) {
	var res *model.Tabs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeDocumentTabsUpdateDocumentTabsis uncategorized and subject to change
func (s *Service) EnvelopeDocumentTabsUpdateDocumentTabs(documentID string, envelopeID string, templateRecipientTabs *model.Tabs) *EnvelopeDocumentTabsUpdateDocumentTabsOp {
	return &EnvelopeDocumentTabsUpdateDocumentTabsOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"envelopes", envelopeID, "documents", documentID, "tabs"}, "/"),
		Payload:    templateRecipientTabs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeDocumentTabsUpdateDocumentTabsOp implements DocuSign API SDK Uncategorized::updateDocumentTabs
type EnvelopeDocumentTabsUpdateDocumentTabsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeDocumentTabsUpdateDocumentTabsOp) Do(ctx context.Context) (*model.Tabs, error) {
	var res *model.Tabs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeHTMLDefinitionsListis uncategorized and subject to change
func (s *Service) EnvelopeHTMLDefinitionsList(envelopeID string) *EnvelopeHTMLDefinitionsListOp {
	return &EnvelopeHTMLDefinitionsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"envelopes", envelopeID, "html_definitions"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeHTMLDefinitionsListOp implements DocuSign API SDK Uncategorized::getEnvelopeHtmlDefinitions
type EnvelopeHTMLDefinitionsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeHTMLDefinitionsListOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitionOriginals, error) {
	var res *model.DocumentHTMLDefinitionOriginals
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopePurgeConfigurationGetEnvelopePurgeConfigurationis uncategorized and subject to change
func (s *Service) EnvelopePurgeConfigurationGetEnvelopePurgeConfiguration() *EnvelopePurgeConfigurationGetEnvelopePurgeConfigurationOp {
	return &EnvelopePurgeConfigurationGetEnvelopePurgeConfigurationOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "settings/envelope_purge_configuration",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopePurgeConfigurationGetEnvelopePurgeConfigurationOp implements DocuSign API SDK Uncategorized::getEnvelopePurgeConfiguration
type EnvelopePurgeConfigurationGetEnvelopePurgeConfigurationOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopePurgeConfigurationGetEnvelopePurgeConfigurationOp) Do(ctx context.Context) (*model.EnvelopePurgeConfiguration, error) {
	var res *model.EnvelopePurgeConfiguration
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopePurgeConfigurationUpdateEnvelopePurgeConfigurationis uncategorized and subject to change
func (s *Service) EnvelopePurgeConfigurationUpdateEnvelopePurgeConfiguration(envelopePurgeConfiguration *model.EnvelopePurgeConfiguration) *EnvelopePurgeConfigurationUpdateEnvelopePurgeConfigurationOp {
	return &EnvelopePurgeConfigurationUpdateEnvelopePurgeConfigurationOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "settings/envelope_purge_configuration",
		Payload:    envelopePurgeConfiguration,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopePurgeConfigurationUpdateEnvelopePurgeConfigurationOp implements DocuSign API SDK Uncategorized::updateEnvelopePurgeConfiguration
type EnvelopePurgeConfigurationUpdateEnvelopePurgeConfigurationOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopePurgeConfigurationUpdateEnvelopePurgeConfigurationOp) Do(ctx context.Context) (*model.EnvelopePurgeConfiguration, error) {
	var res *model.EnvelopePurgeConfiguration
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeTransferRulesCreateEnvelopeTransferRulesis uncategorized and subject to change
func (s *Service) EnvelopeTransferRulesCreateEnvelopeTransferRules(envelopeTransferRuleRequest *model.EnvelopeTransferRuleRequest) *EnvelopeTransferRulesCreateEnvelopeTransferRulesOp {
	return &EnvelopeTransferRulesCreateEnvelopeTransferRulesOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "envelopes/transfer_rules",
		Payload:    envelopeTransferRuleRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeTransferRulesCreateEnvelopeTransferRulesOp implements DocuSign API SDK Uncategorized::createEnvelopeTransferRules
type EnvelopeTransferRulesCreateEnvelopeTransferRulesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeTransferRulesCreateEnvelopeTransferRulesOp) Do(ctx context.Context) (*model.EnvelopeTransferRuleInformation, error) {
	var res *model.EnvelopeTransferRuleInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeTransferRulesDeleteEnvelopeTransferRulesis uncategorized and subject to change
func (s *Service) EnvelopeTransferRulesDeleteEnvelopeTransferRules(envelopeTransferRuleID string) *EnvelopeTransferRulesDeleteEnvelopeTransferRulesOp {
	return &EnvelopeTransferRulesDeleteEnvelopeTransferRulesOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"envelopes", "transfer_rules", envelopeTransferRuleID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeTransferRulesDeleteEnvelopeTransferRulesOp implements DocuSign API SDK Uncategorized::deleteEnvelopeTransferRules
type EnvelopeTransferRulesDeleteEnvelopeTransferRulesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeTransferRulesDeleteEnvelopeTransferRulesOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// EnvelopeTransferRulesGetEnvelopeTransferRulesis uncategorized and subject to change
func (s *Service) EnvelopeTransferRulesGetEnvelopeTransferRules() *EnvelopeTransferRulesGetEnvelopeTransferRulesOp {
	return &EnvelopeTransferRulesGetEnvelopeTransferRulesOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "envelopes/transfer_rules",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeTransferRulesGetEnvelopeTransferRulesOp implements DocuSign API SDK Uncategorized::getEnvelopeTransferRules
type EnvelopeTransferRulesGetEnvelopeTransferRulesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeTransferRulesGetEnvelopeTransferRulesOp) Do(ctx context.Context) (*model.EnvelopeTransferRuleInformation, error) {
	var res *model.EnvelopeTransferRuleInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Count is the maximum number of results to return.
func (op *EnvelopeTransferRulesGetEnvelopeTransferRulesOp) Count(val string) *EnvelopeTransferRulesGetEnvelopeTransferRulesOp {
	if op != nil {
		op.QueryOpts.Set("count", val)
	}
	return op
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *EnvelopeTransferRulesGetEnvelopeTransferRulesOp) StartPosition(val string) *EnvelopeTransferRulesGetEnvelopeTransferRulesOp {
	if op != nil {
		op.QueryOpts.Set("start_position", val)
	}
	return op
}

// EnvelopeTransferRulesUpdateEnvelopeTransferRuleis uncategorized and subject to change
func (s *Service) EnvelopeTransferRulesUpdateEnvelopeTransferRule(envelopeTransferRuleID string, envelopeTransferRule *model.EnvelopeTransferRule) *EnvelopeTransferRulesUpdateEnvelopeTransferRuleOp {
	return &EnvelopeTransferRulesUpdateEnvelopeTransferRuleOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"envelopes", "transfer_rules", envelopeTransferRuleID}, "/"),
		Payload:    envelopeTransferRule,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeTransferRulesUpdateEnvelopeTransferRuleOp implements DocuSign API SDK Uncategorized::updateEnvelopeTransferRule
type EnvelopeTransferRulesUpdateEnvelopeTransferRuleOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeTransferRulesUpdateEnvelopeTransferRuleOp) Do(ctx context.Context) (*model.EnvelopeTransferRule, error) {
	var res *model.EnvelopeTransferRule
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeTransferRulesUpdateEnvelopeTransferRulesis uncategorized and subject to change
func (s *Service) EnvelopeTransferRulesUpdateEnvelopeTransferRules(envelopeTransferRules *model.EnvelopeTransferRuleInformation) *EnvelopeTransferRulesUpdateEnvelopeTransferRulesOp {
	return &EnvelopeTransferRulesUpdateEnvelopeTransferRulesOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "envelopes/transfer_rules",
		Payload:    envelopeTransferRules,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeTransferRulesUpdateEnvelopeTransferRulesOp implements DocuSign API SDK Uncategorized::updateEnvelopeTransferRules
type EnvelopeTransferRulesUpdateEnvelopeTransferRulesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeTransferRulesUpdateEnvelopeTransferRulesOp) Do(ctx context.Context) (*model.EnvelopeTransferRuleInformation, error) {
	var res *model.EnvelopeTransferRuleInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeViewsCreateEnvelopeRecipientPreviewis uncategorized and subject to change
func (s *Service) EnvelopeViewsCreateEnvelopeRecipientPreview(envelopeID string, recipientPreviewRequest *model.RecipientPreviewRequest) *EnvelopeViewsCreateEnvelopeRecipientPreviewOp {
	return &EnvelopeViewsCreateEnvelopeRecipientPreviewOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"envelopes", envelopeID, "views", "recipient_preview"}, "/"),
		Payload:    recipientPreviewRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeViewsCreateEnvelopeRecipientPreviewOp implements DocuSign API SDK Uncategorized::createEnvelopeRecipientPreview
type EnvelopeViewsCreateEnvelopeRecipientPreviewOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeViewsCreateEnvelopeRecipientPreviewOp) Do(ctx context.Context) (*model.ViewURL, error) {
	var res *model.ViewURL
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// EnvelopeViewsCreateTemplateRecipientPreviewis uncategorized and subject to change
func (s *Service) EnvelopeViewsCreateTemplateRecipientPreview(templateID string, recipientPreviewRequest *model.RecipientPreviewRequest) *EnvelopeViewsCreateTemplateRecipientPreviewOp {
	return &EnvelopeViewsCreateTemplateRecipientPreviewOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"templates", templateID, "views", "recipient_preview"}, "/"),
		Payload:    recipientPreviewRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// EnvelopeViewsCreateTemplateRecipientPreviewOp implements DocuSign API SDK Uncategorized::createTemplateRecipientPreview
type EnvelopeViewsCreateTemplateRecipientPreviewOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *EnvelopeViewsCreateTemplateRecipientPreviewOp) Do(ctx context.Context) (*model.ViewURL, error) {
	var res *model.ViewURL
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// NotificationDefaultsGetNotificationDefaultsis uncategorized and subject to change
func (s *Service) NotificationDefaultsGetNotificationDefaults() *NotificationDefaultsGetNotificationDefaultsOp {
	return &NotificationDefaultsGetNotificationDefaultsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "settings/notification_defaults",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// NotificationDefaultsGetNotificationDefaultsOp implements DocuSign API SDK Uncategorized::getNotificationDefaults
type NotificationDefaultsGetNotificationDefaultsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *NotificationDefaultsGetNotificationDefaultsOp) Do(ctx context.Context) (*model.NotificationDefaults, error) {
	var res *model.NotificationDefaults
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// NotificationDefaultsUpdateNotificationDefaultsis uncategorized and subject to change
func (s *Service) NotificationDefaultsUpdateNotificationDefaults(notificationDefaults *model.NotificationDefaults) *NotificationDefaultsUpdateNotificationDefaultsOp {
	return &NotificationDefaultsUpdateNotificationDefaultsOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "settings/notification_defaults",
		Payload:    notificationDefaults,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// NotificationDefaultsUpdateNotificationDefaultsOp implements DocuSign API SDK Uncategorized::updateNotificationDefaults
type NotificationDefaultsUpdateNotificationDefaultsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *NotificationDefaultsUpdateNotificationDefaultsOp) Do(ctx context.Context) (*model.NotificationDefaults, error) {
	var res *model.NotificationDefaults
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
		Version:    esign.VersionV21,
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
		Version:    esign.VersionV21,
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
		Version:    esign.VersionV21,
	}
}

// TemplateDocumentResponsiveHTMLPreviewCreateOp implements DocuSign API SDK Uncategorized::createTemplateDocumentResponsiveHtmlPreview
type TemplateDocumentResponsiveHTMLPreviewCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateDocumentResponsiveHTMLPreviewCreateOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitions, error) {
	var res *model.DocumentHTMLDefinitions
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateDocumentTabsCreateTemplateDocumentTabs create Template Document Tabs
// operation is uncategorized and subject to change.
func (s *Service) TemplateDocumentTabsCreateTemplateDocumentTabs(documentID string, templateID string, templateTabs *model.TemplateTabs) *TemplateDocumentTabsCreateTemplateDocumentTabsOp {
	return &TemplateDocumentTabsCreateTemplateDocumentTabsOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       strings.Join([]string{"templates", templateID, "documents", documentID, "tabs"}, "/"),
		Payload:    templateTabs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// TemplateDocumentTabsCreateTemplateDocumentTabsOp implements DocuSign API SDK Uncategorized::createTemplateDocumentTabs
type TemplateDocumentTabsCreateTemplateDocumentTabsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateDocumentTabsCreateTemplateDocumentTabsOp) Do(ctx context.Context) (*model.Tabs, error) {
	var res *model.Tabs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateDocumentTabsDeleteTemplateDocumentTabsis uncategorized and subject to change
func (s *Service) TemplateDocumentTabsDeleteTemplateDocumentTabs(documentID string, templateID string, templateTabs *model.TemplateTabs) *TemplateDocumentTabsDeleteTemplateDocumentTabsOp {
	return &TemplateDocumentTabsDeleteTemplateDocumentTabsOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"templates", templateID, "documents", documentID, "tabs"}, "/"),
		Payload:    templateTabs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// TemplateDocumentTabsDeleteTemplateDocumentTabsOp implements DocuSign API SDK Uncategorized::deleteTemplateDocumentTabs
type TemplateDocumentTabsDeleteTemplateDocumentTabsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateDocumentTabsDeleteTemplateDocumentTabsOp) Do(ctx context.Context) (*model.Tabs, error) {
	var res *model.Tabs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateDocumentTabsUpdateTemplateDocumentTabsis uncategorized and subject to change
func (s *Service) TemplateDocumentTabsUpdateTemplateDocumentTabs(documentID string, templateID string, templateTabs *model.TemplateTabs) *TemplateDocumentTabsUpdateTemplateDocumentTabsOp {
	return &TemplateDocumentTabsUpdateTemplateDocumentTabsOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"templates", templateID, "documents", documentID, "tabs"}, "/"),
		Payload:    templateTabs,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// TemplateDocumentTabsUpdateTemplateDocumentTabsOp implements DocuSign API SDK Uncategorized::updateTemplateDocumentTabs
type TemplateDocumentTabsUpdateTemplateDocumentTabsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateDocumentTabsUpdateTemplateDocumentTabsOp) Do(ctx context.Context) (*model.Tabs, error) {
	var res *model.Tabs
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// TemplateHTMLDefinitionsListis uncategorized and subject to change
func (s *Service) TemplateHTMLDefinitionsList(templateID string) *TemplateHTMLDefinitionsListOp {
	return &TemplateHTMLDefinitionsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"templates", templateID, "html_definitions"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
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
		Version:    esign.VersionV21,
	}
}

// TemplateResponsiveHTMLPreviewCreateOp implements DocuSign API SDK Uncategorized::createTemplateResponsiveHtmlPreview
type TemplateResponsiveHTMLPreviewCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *TemplateResponsiveHTMLPreviewCreateOp) Do(ctx context.Context) (*model.DocumentHTMLDefinitions, error) {
	var res *model.DocumentHTMLDefinitions
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
