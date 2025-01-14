// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PageEntryFulfillment struct {
	/* The list of rich message responses to present to the user. */
	// +optional
	Messages []PageMessages `json:"messages,omitempty"`

	/* Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks. */
	// +optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	/* The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified. */
	// +optional
	Tag *string `json:"tag,omitempty"`

	/* The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>. */
	// +optional
	Webhook *string `json:"webhook,omitempty"`
}

type PageEventHandlers struct {
	/* The name of the event to handle. */
	// +optional
	Event *string `json:"event,omitempty"`

	/* The unique identifier of this event handler. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* The target flow to transition to.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>. */
	// +optional
	TargetFlow *string `json:"targetFlow,omitempty"`

	/* The target page to transition to.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>. */
	// +optional
	TargetPage *string `json:"targetPage,omitempty"`

	/* The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks. */
	// +optional
	TriggerFulfillment *PageTriggerFulfillment `json:"triggerFulfillment,omitempty"`
}

type PageFillBehavior struct {
	/* The fulfillment to provide the initial prompt that the agent can present to the user in order to fill the parameter. */
	// +optional
	InitialPromptFulfillment *PageInitialPromptFulfillment `json:"initialPromptFulfillment,omitempty"`
}

type PageForm struct {
	/* Parameters to collect from the user. */
	// +optional
	Parameters []PageParameters `json:"parameters,omitempty"`
}

type PageInitialPromptFulfillment struct {
	/* The list of rich message responses to present to the user. */
	// +optional
	Messages []PageMessages `json:"messages,omitempty"`

	/* Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks. */
	// +optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	/* The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified. */
	// +optional
	Tag *string `json:"tag,omitempty"`

	/* The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>. */
	// +optional
	Webhook *string `json:"webhook,omitempty"`
}

type PageMessages struct {
	/* The text response message. */
	// +optional
	Text *PageText `json:"text,omitempty"`
}

type PageParameters struct {
	/* The human-readable name of the parameter, unique within the form. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* The entity type of the parameter.
	Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity types. */
	// +optional
	EntityType *string `json:"entityType,omitempty"`

	/* Defines fill behavior for the parameter. */
	// +optional
	FillBehavior *PageFillBehavior `json:"fillBehavior,omitempty"`

	/* Indicates whether the parameter represents a list of values. */
	// +optional
	IsList *bool `json:"isList,omitempty"`

	/* Indicates whether the parameter content should be redacted in log.
	If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled. */
	// +optional
	Redact *bool `json:"redact,omitempty"`

	/* Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
	Required parameters must be filled before form filling concludes. */
	// +optional
	Required *bool `json:"required,omitempty"`
}

type PageText struct {
	/* Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request. */
	// +optional
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty"`

	/* A collection of text responses. */
	// +optional
	Text []PageText `json:"text,omitempty"`
}

type PageTransitionRoutes struct {
	/* The condition to evaluate against form parameters or session parameters.
	At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled. */
	// +optional
	Condition *string `json:"condition,omitempty"`

	/* The unique identifier of an Intent.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled. */
	// +optional
	Intent *string `json:"intent,omitempty"`

	/* The unique identifier of this transition route. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* The target flow to transition to.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>. */
	// +optional
	TargetFlow *string `json:"targetFlow,omitempty"`

	/* The target page to transition to.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>. */
	// +optional
	TargetPage *string `json:"targetPage,omitempty"`

	/* The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks. */
	// +optional
	TriggerFulfillment *PageTriggerFulfillment `json:"triggerFulfillment,omitempty"`
}

type PageTriggerFulfillment struct {
	/* The list of rich message responses to present to the user. */
	// +optional
	Messages []PageMessages `json:"messages,omitempty"`

	/* Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks. */
	// +optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty"`

	/* The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified. */
	// +optional
	Tag *string `json:"tag,omitempty"`

	/* The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>. */
	// +optional
	Webhook *string `json:"webhook,omitempty"`
}

type DialogflowCXPageSpec struct {
	/* The human-readable name of the page, unique within the agent. */
	DisplayName string `json:"displayName"`

	/* The fulfillment to call when the session is entering the page. */
	// +optional
	EntryFulfillment *PageEntryFulfillment `json:"entryFulfillment,omitempty"`

	/* Handlers associated with the page to handle events such as webhook errors, no match or no input. */
	// +optional
	EventHandlers []PageEventHandlers `json:"eventHandlers,omitempty"`

	/* The form associated with the page, used for collecting parameters relevant to the page. */
	// +optional
	Form *PageForm `json:"form,omitempty"`

	/* Immutable. The language of the following fields in page:

	Page.entry_fulfillment.messages
	Page.entry_fulfillment.conditional_cases
	Page.event_handlers.trigger_fulfillment.messages
	Page.event_handlers.trigger_fulfillment.conditional_cases
	Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	Page.transition_routes.trigger_fulfillment.messages
	Page.transition_routes.trigger_fulfillment.conditional_cases
	If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used. */
	// +optional
	LanguageCode *string `json:"languageCode,omitempty"`

	/* Immutable. The flow to create a page for.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>. */
	// +optional
	Parent *string `json:"parent,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes.
	If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>. */
	// +optional
	TransitionRouteGroups []string `json:"transitionRouteGroups,omitempty"`

	/* A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	TransitionRoutes defined in the page with intent specified.
	TransitionRoutes defined in the transition route groups with intent specified.
	TransitionRoutes defined in flow with intent specified.
	TransitionRoutes defined in the transition route groups with intent specified.
	TransitionRoutes defined in the page with only condition specified.
	TransitionRoutes defined in the transition route groups with only condition specified. */
	// +optional
	TransitionRoutes []PageTransitionRoutes `json:"transitionRoutes,omitempty"`
}

type DialogflowCXPageStatus struct {
	/* Conditions represent the latest available observations of the
	   DialogflowCXPage's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The unique identifier of the page.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DialogflowCXPage is the Schema for the dialogflowcx API
// +k8s:openapi-gen=true
type DialogflowCXPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DialogflowCXPageSpec   `json:"spec,omitempty"`
	Status DialogflowCXPageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DialogflowCXPageList contains a list of DialogflowCXPage
type DialogflowCXPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXPage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXPage{}, &DialogflowCXPageList{})
}
