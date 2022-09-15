// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build windows && npm
// +build windows,npm

package etw

const (
	// From https://github.com/repnz/etw-providers-docs/blob/master/Manifests-Win10-18990/Microsoft-Windows-HttpService.xml
	IdHTTP_DUMMY_START uint16 = iota
	IdHTTPRequestTraceTaskRecvReq
	IdHTTPRequestTraceTaskParse
	IdHTTPRequestTraceTaskDeliver
	IdHTTPRequestTraceTaskRecvResp
	IdHTTPRequestTraceTaskRecvRespLast
	IdHTTPRequestTraceTaskRecvBody
	IdHTTPRequestTraceTaskRecvBodyLast
	IdHTTPRequestTraceTaskFastResp
	IdHTTPRequestTraceTaskFastRespLast
	IdHTTPRequestTraceTaskSendComplete
	IdHTTPRequestTraceTaskCachedAndSend
	IdHTTPRequestTraceTaskFastSend
	IdHTTPRequestTraceTaskZeroSend
	IdHTTPRequestTraceTaskLastSndError
	IdHTTPRequestTraceTaskSndError
	IdHTTPRequestTraceTaskSrvdFrmCache
	IdHTTPRequestTraceTaskCachedNotModified
	IdHTTPSetupTraceTaskResvUrl
	IdHTTPSetupTraceTaskReadIpListEntry
	IdHTTPSetupTraceTaskCreatedSslCred
	IdHTTPConnectionTraceTaskConnConnect
	IdHTTPConnectionTraceTaskConnIdAssgn
	IdHTTPConnectionTraceTaskConnClose
	IdHTTPConnectionTraceTaskConnCleanup
	IdHTTPCacheTraceTaskAddedCacheEntry
	IdHTTPCacheTraceTaskAddCacheEntryFailed
	IdHTTPCacheTraceTaskFlushedCache
	IdHTTPConfigurationPropertyTraceTaskChgUrlGrpProp
	IdHTTPConfigurationPropertyTraceTaskChgSrvSesProp
	IdHTTPConfigurationPropertyTraceTaskChgReqQueueProp
	IdHTTPConfigurationPropertyTraceTaskAddUrl
	IdHTTPConfigurationPropertyTraceTaskRemUrl
	IdHTTPConfigurationPropertyTraceTaskRemAllUrls
	IdHTTPSSLTraceTaskSslConnEvent
	IdHTTPSSLTraceTaskSslInitiateHandshake
	IdHTTPSSLTraceTaskSslHandshakeComplete
	IdHTTPSSLTraceTaskSslInititateSslRcvClientCert
	IdHTTPSSLTraceTaskSslRcvClientCertFailed
	IdHTTPSSLTraceTaskSslRcvdRawData
	IdHTTPSSLTraceTaskSslDlvrdStreamData
	IdHTTPSSLTraceTaskSslAcceptStreamData
	IdHTTPAuthenticationTraceTaskSspiCall
	IdHTTPAuthenticationTraceTaskAuthCacheEntryAdded
	IdHTTPAuthenticationTraceTaskAuthCacheEntryFreed
	IdHTTPConnectionTraceTaskQosFlowSetReset
	IdHTTPLoggingTraceTaskLoggingConfigFailed
	IdHTTPLoggingTraceTaskLoggingConfig
	IdHTTPLoggingTraceTaskLogFileCreateFailed
	IdHTTPLoggingTraceTaskLogFileCreate
	IdHTTPLoggingTraceTaskLogFileWrite
	IdHTTPRequestTraceTaskParseRequestFailed
	IdHTTPTimeoutTraceTaskConnTimedOut
	IdHTTPSSLTraceTaskSslEndpointCreationFailed
	IdHTTPSSLTraceTaskSslDisconnEvent
	IdHTTPSSLTraceTaskSslDisconnReq
	IdHTTPSSLTraceTaskSslUnsealMsg
	IdHTTPSSLTraceTaskSslQueryConnInfoFailed
	IdHTTPSSLTraceTaskSslEndpointConfigNotFound
	IdHTTPSSLTraceTaskSslAsc
	IdHTTPSSLTraceTaskSslSealMsg
	IdHTTPRequestTraceTaskRequestRejected
	IdHTTPRequestTraceTaskRequestCancelled
	IdHTTPDriverGlobalSettingsTaskHotAddProcFailed
	IdHTTPDriverGlobalSettingsTaskHotAddProcSucceeded
	IdHTTPRequestTraceTaskUserResponseFlowInit
	IdHTTPRequestTraceTaskCachedResponseFlowInit
	IdHTTPRequestTraceTaskFlowInitFailed
	IdHTTPConnectionTraceTaskSetConnectionFlow
	IdHTTPConnectionTraceTaskRequestAssociatedToConfigurationFlow
	IdHTTPConnectionTraceTaskConnectionFlowFailed
	IdHTTPRequestTraceTaskResponseRangeProcessingOK
	IdHTTPCacheTraceTaskBeginBuildingSlices
	IdHTTPCacheTraceTaskSendSliceCacheContent
	IdHTTPCacheTraceTaskCachedSlicesMatchContent
	IdHTTPCacheTraceTaskMergeSlicesToCache
	IdHTTPCacheTraceTaskFlatCacheRangeSend
	IdHTTPAuthenticationTraceTaskChannelBindAscParams
	IdHTTPAuthenticationTraceTaskServiceBindCheckComplete
	IdHTTPAuthenticationTraceTaskChannelBindConfigCapture
	IdHTTPAuthenticationTraceTaskChannelBindPerResponseConfig
	IdHTTPConnectionTraceTaskUsePolicyBasedQoSFlow
	IdHTTPThreadPoolThreadPoolExtension
	IdHTTPThreadPoolThreadReady
	IdHTTPThreadPoolThreadPoolTrim
	IdHTTPThreadPoolThreadGone
	IdHTTPSSLTraceTaskSniParsed
	IdHTTPRequestTraceTaskInitiateOpaqueMode
	IdHTTPSSLTraceTaskEndpointAutoGenerated
	IdHTTPSSLTraceTaskAutoGeneratedEndpointDeleted
	IdHTTPSSLTraceTaskSslEndpointConfigFound
	IdHTTPSSLTraceTaskSslEndpointConfigRejected
	IdHTTPResponseTraceTaskParseRequestFailed
	IdHTTPSSLTraceTaskSslHandshakeFailure
	IdHTTPRequestTraceTaskHttpErrorResponseSent
	IdHTTPSSLTraceTaskSslRenegotiateTimedOut
	IdHTTPRequestTraceTaskHttp11Required
	IdHTTPConnectionTraceTaskQuicConnection
	IdHTTPConnectionTraceTaskQuicConnectionCallback
	IdHTTPConnectionTraceTaskQuicStream
	IdHTTPConnectionTraceTaskQuicStreamCallback
	IdHTTPDriverGlobalSettingsTaskQuicRegistration
	IdHTTP_DUMMY_MAX
)

var (
	httpServiceEventId2Name = []string{
		"HTTP_DUMMY_START",
		"HTTPRequestTraceTaskRecvReq",
		"HTTPRequestTraceTaskParse",
		"HTTPRequestTraceTaskDeliver",
		"HTTPRequestTraceTaskRecvResp",
		"HTTPRequestTraceTaskRecvRespLast",
		"HTTPRequestTraceTaskRecvBody",
		"HTTPRequestTraceTaskRecvBodyLast",
		"HTTPRequestTraceTaskFastResp",
		"HTTPRequestTraceTaskFastRespLast",
		"HTTPRequestTraceTaskSendComplete",
		"HTTPRequestTraceTaskCachedAndSend",
		"HTTPRequestTraceTaskFastSend",
		"HTTPRequestTraceTaskZeroSend",
		"HTTPRequestTraceTaskLastSndError",
		"HTTPRequestTraceTaskSndError",
		"HTTPRequestTraceTaskSrvdFrmCache",
		"HTTPRequestTraceTaskCachedNotModified",
		"HTTPSetupTraceTaskResvUrl",
		"HTTPSetupTraceTaskReadIpListEntry",
		"HTTPSetupTraceTaskCreatedSslCred",
		"HTTPConnectionTraceTaskConnConnect",
		"HTTPConnectionTraceTaskConnIdAssgn",
		"HTTPConnectionTraceTaskConnClose",
		"HTTPConnectionTraceTaskConnCleanup",
		"HTTPCacheTraceTaskAddedCacheEntry",
		"HTTPCacheTraceTaskAddCacheEntryFailed",
		"HTTPCacheTraceTaskFlushedCache",
		"HTTPConfigurationPropertyTraceTaskChgUrlGrpProp",
		"HTTPConfigurationPropertyTraceTaskChgSrvSesProp",
		"HTTPConfigurationPropertyTraceTaskChgReqQueueProp",
		"HTTPConfigurationPropertyTraceTaskAddUrl",
		"HTTPConfigurationPropertyTraceTaskRemUrl",
		"HTTPConfigurationPropertyTraceTaskRemAllUrls",
		"HTTPSSLTraceTaskSslConnEvent",
		"HTTPSSLTraceTaskSslInitiateHandshake",
		"HTTPSSLTraceTaskSslHandshakeComplete",
		"HTTPSSLTraceTaskSslInititateSslRcvClientCert",
		"HTTPSSLTraceTaskSslRcvClientCertFailed",
		"HTTPSSLTraceTaskSslRcvdRawData",
		"HTTPSSLTraceTaskSslDlvrdStreamData",
		"HTTPSSLTraceTaskSslAcceptStreamData",
		"HTTPAuthenticationTraceTaskSspiCall",
		"HTTPAuthenticationTraceTaskAuthCacheEntryAdded",
		"HTTPAuthenticationTraceTaskAuthCacheEntryFreed",
		"HTTPConnectionTraceTaskQosFlowSetReset",
		"HTTPLoggingTraceTaskLoggingConfigFailed",
		"HTTPLoggingTraceTaskLoggingConfig",
		"HTTPLoggingTraceTaskLogFileCreateFailed",
		"HTTPLoggingTraceTaskLogFileCreate",
		"HTTPLoggingTraceTaskLogFileWrite",
		"HTTPRequestTraceTaskParseRequestFailed",
		"HTTPTimeoutTraceTaskConnTimedOut",
		"HTTPSSLTraceTaskSslEndpointCreationFailed",
		"HTTPSSLTraceTaskSslDisconnEvent",
		"HTTPSSLTraceTaskSslDisconnReq",
		"HTTPSSLTraceTaskSslUnsealMsg",
		"HTTPSSLTraceTaskSslQueryConnInfoFailed",
		"HTTPSSLTraceTaskSslEndpointConfigNotFound",
		"HTTPSSLTraceTaskSslAsc",
		"HTTPSSLTraceTaskSslSealMsg",
		"HTTPRequestTraceTaskRequestRejected",
		"HTTPRequestTraceTaskRequestCancelled",
		"HTTPDriverGlobalSettingsTaskHotAddProcFailed",
		"HTTPDriverGlobalSettingsTaskHotAddProcSucceeded",
		"HTTPRequestTraceTaskUserResponseFlowInit",
		"HTTPRequestTraceTaskCachedResponseFlowInit",
		"HTTPRequestTraceTaskFlowInitFailed",
		"HTTPConnectionTraceTaskSetConnectionFlow",
		"HTTPConnectionTraceTaskRequestAssociatedToConfigurationFlow",
		"HTTPConnectionTraceTaskConnectionFlowFailed",
		"HTTPRequestTraceTaskResponseRangeProcessingOK",
		"HTTPCacheTraceTaskBeginBuildingSlices",
		"HTTPCacheTraceTaskSendSliceCacheContent",
		"HTTPCacheTraceTaskCachedSlicesMatchContent",
		"HTTPCacheTraceTaskMergeSlicesToCache",
		"HTTPCacheTraceTaskFlatCacheRangeSend",
		"HTTPAuthenticationTraceTaskChannelBindAscParams",
		"HTTPAuthenticationTraceTaskServiceBindCheckComplete",
		"HTTPAuthenticationTraceTaskChannelBindConfigCapture",
		"HTTPAuthenticationTraceTaskChannelBindPerResponseConfig",
		"HTTPConnectionTraceTaskUsePolicyBasedQoSFlow",
		"HTTPThreadPoolThreadPoolExtension",
		"HTTPThreadPoolThreadReady",
		"HTTPThreadPoolThreadPoolTrim",
		"HTTPThreadPoolThreadGone",
		"HTTPSSLTraceTaskSniParsed",
		"HTTPRequestTraceTaskInitiateOpaqueMode",
		"HTTPSSLTraceTaskEndpointAutoGenerated",
		"HTTPSSLTraceTaskAutoGeneratedEndpointDeleted",
		"HTTPSSLTraceTaskSslEndpointConfigFound",
		"HTTPSSLTraceTaskSslEndpointConfigRejected",
		"HTTPResponseTraceTaskParseRequestFailed",
		"HTTPSSLTraceTaskSslHandshakeFailure",
		"HTTPRequestTraceTaskHttpErrorResponseSent",
		"HTTPSSLTraceTaskSslRenegotiateTimedOut",
		"HTTPRequestTraceTaskHttp11Required",
		"HTTPConnectionTraceTaskQuicConnection",
		"HTTPConnectionTraceTaskQuicConnectionCallback",
		"HTTPConnectionTraceTaskQuicStream",
		"HTTPConnectionTraceTaskQuicStreamCallback",
		"HTTPDriverGlobalSettingsTaskQuicRegistration",
	}
)

func formatHttpServiceEventId(eventId uint16) string {
	if eventId == IdHTTP_DUMMY_START || eventId >= IdHTTP_DUMMY_MAX {
		return "<UNKNOWN_HTTP_EVENT_ID>"
	}

	return httpServiceEventId2Name[eventId]
}