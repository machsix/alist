// Package speech_to_text code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkspeech_to_text

import (
	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type DepartmentId struct {
	DepartmentId *string `json:"department_id,omitempty"` //

	OpenDepartmentId *string `json:"open_department_id,omitempty"` //
}

type DepartmentIdBuilder struct {
	departmentId     string //
	departmentIdFlag bool

	openDepartmentId     string //
	openDepartmentIdFlag bool
}

func NewDepartmentIdBuilder() *DepartmentIdBuilder {
	builder := &DepartmentIdBuilder{}
	return builder
}

//
//
// 示例值：
func (builder *DepartmentIdBuilder) DepartmentId(departmentId string) *DepartmentIdBuilder {
	builder.departmentId = departmentId
	builder.departmentIdFlag = true
	return builder
}

//
//
// 示例值：
func (builder *DepartmentIdBuilder) OpenDepartmentId(openDepartmentId string) *DepartmentIdBuilder {
	builder.openDepartmentId = openDepartmentId
	builder.openDepartmentIdFlag = true
	return builder
}

func (builder *DepartmentIdBuilder) Build() *DepartmentId {
	req := &DepartmentId{}
	if builder.departmentIdFlag {
		req.DepartmentId = &builder.departmentId

	}
	if builder.openDepartmentIdFlag {
		req.OpenDepartmentId = &builder.openDepartmentId

	}
	return req
}

type FileConfig struct {
	FileId *string `json:"file_id,omitempty"` // 仅包含字母数字和下划线的 16 位字符串作为文件的标识，用户生成

	Format *string `json:"format,omitempty"` // 语音格式，目前仅支持：pcm

	EngineType *string `json:"engine_type,omitempty"` // 引擎类型，目前仅支持：16k_auto 中英混合
}

type FileConfigBuilder struct {
	fileId     string // 仅包含字母数字和下划线的 16 位字符串作为文件的标识，用户生成
	fileIdFlag bool

	format     string // 语音格式，目前仅支持：pcm
	formatFlag bool

	engineType     string // 引擎类型，目前仅支持：16k_auto 中英混合
	engineTypeFlag bool
}

func NewFileConfigBuilder() *FileConfigBuilder {
	builder := &FileConfigBuilder{}
	return builder
}

// 仅包含字母数字和下划线的 16 位字符串作为文件的标识，用户生成
//
// 示例值：qwe12dd34567890w
func (builder *FileConfigBuilder) FileId(fileId string) *FileConfigBuilder {
	builder.fileId = fileId
	builder.fileIdFlag = true
	return builder
}

// 语音格式，目前仅支持：pcm
//
// 示例值：pcm
func (builder *FileConfigBuilder) Format(format string) *FileConfigBuilder {
	builder.format = format
	builder.formatFlag = true
	return builder
}

// 引擎类型，目前仅支持：16k_auto 中英混合
//
// 示例值：16k_auto
func (builder *FileConfigBuilder) EngineType(engineType string) *FileConfigBuilder {
	builder.engineType = engineType
	builder.engineTypeFlag = true
	return builder
}

func (builder *FileConfigBuilder) Build() *FileConfig {
	req := &FileConfig{}
	if builder.fileIdFlag {
		req.FileId = &builder.fileId

	}
	if builder.formatFlag {
		req.Format = &builder.format

	}
	if builder.engineTypeFlag {
		req.EngineType = &builder.engineType

	}
	return req
}

type Speech struct {
	Speech *string `json:"speech,omitempty"` // pcm格式音频文件（文件识别）或音频分片（流式识别）经base64编码后的内容

	SpeechKey *string `json:"speech_key,omitempty"` // 上传到 drive 存储平台后获取到的 key （暂不支持）
}

type SpeechBuilder struct {
	speech     string // pcm格式音频文件（文件识别）或音频分片（流式识别）经base64编码后的内容
	speechFlag bool

	speechKey     string // 上传到 drive 存储平台后获取到的 key （暂不支持）
	speechKeyFlag bool
}

func NewSpeechBuilder() *SpeechBuilder {
	builder := &SpeechBuilder{}
	return builder
}

// pcm格式音频文件（文件识别）或音频分片（流式识别）经base64编码后的内容
//
// 示例值：PdmrfE267Cd/Z9KpmNFh71A2PSJZxSp7+8upCg==
func (builder *SpeechBuilder) Speech(speech string) *SpeechBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}

// 上传到 drive 存储平台后获取到的 key （暂不支持）
//
// 示例值：
func (builder *SpeechBuilder) SpeechKey(speechKey string) *SpeechBuilder {
	builder.speechKey = speechKey
	builder.speechKeyFlag = true
	return builder
}

func (builder *SpeechBuilder) Build() *Speech {
	req := &Speech{}
	if builder.speechFlag {
		req.Speech = &builder.speech

	}
	if builder.speechKeyFlag {
		req.SpeechKey = &builder.speechKey

	}
	return req
}

type StreamConfig struct {
	StreamId *string `json:"stream_id,omitempty"` // 仅包含字母数字和下划线的 16 位字符串作为同一数据流的标识，用户生成

	SequenceId *int `json:"sequence_id,omitempty"` // 数据流分片的序号，序号从 0 开始，每次请求递增 1

	Action *int `json:"action,omitempty"` // 数据流标记：1 首包，2 正常结束，等待结果返回，3 中断数据流不返回最终结果

	Format *string `json:"format,omitempty"` // 语音格式，目前仅支持：pcm

	EngineType *string `json:"engine_type,omitempty"` // 引擎类型，目前仅支持：16k_auto 中英混合
}

type StreamConfigBuilder struct {
	streamId     string // 仅包含字母数字和下划线的 16 位字符串作为同一数据流的标识，用户生成
	streamIdFlag bool

	sequenceId     int // 数据流分片的序号，序号从 0 开始，每次请求递增 1
	sequenceIdFlag bool

	action     int // 数据流标记：1 首包，2 正常结束，等待结果返回，3 中断数据流不返回最终结果
	actionFlag bool

	format     string // 语音格式，目前仅支持：pcm
	formatFlag bool

	engineType     string // 引擎类型，目前仅支持：16k_auto 中英混合
	engineTypeFlag bool
}

func NewStreamConfigBuilder() *StreamConfigBuilder {
	builder := &StreamConfigBuilder{}
	return builder
}

// 仅包含字母数字和下划线的 16 位字符串作为同一数据流的标识，用户生成
//
// 示例值：asd1234567890ddd
func (builder *StreamConfigBuilder) StreamId(streamId string) *StreamConfigBuilder {
	builder.streamId = streamId
	builder.streamIdFlag = true
	return builder
}

// 数据流分片的序号，序号从 0 开始，每次请求递增 1
//
// 示例值：1
func (builder *StreamConfigBuilder) SequenceId(sequenceId int) *StreamConfigBuilder {
	builder.sequenceId = sequenceId
	builder.sequenceIdFlag = true
	return builder
}

// 数据流标记：1 首包，2 正常结束，等待结果返回，3 中断数据流不返回最终结果
//
// 示例值：1
func (builder *StreamConfigBuilder) Action(action int) *StreamConfigBuilder {
	builder.action = action
	builder.actionFlag = true
	return builder
}

// 语音格式，目前仅支持：pcm
//
// 示例值：pcm
func (builder *StreamConfigBuilder) Format(format string) *StreamConfigBuilder {
	builder.format = format
	builder.formatFlag = true
	return builder
}

// 引擎类型，目前仅支持：16k_auto 中英混合
//
// 示例值：16k_auto
func (builder *StreamConfigBuilder) EngineType(engineType string) *StreamConfigBuilder {
	builder.engineType = engineType
	builder.engineTypeFlag = true
	return builder
}

func (builder *StreamConfigBuilder) Build() *StreamConfig {
	req := &StreamConfig{}
	if builder.streamIdFlag {
		req.StreamId = &builder.streamId

	}
	if builder.sequenceIdFlag {
		req.SequenceId = &builder.sequenceId

	}
	if builder.actionFlag {
		req.Action = &builder.action

	}
	if builder.formatFlag {
		req.Format = &builder.format

	}
	if builder.engineTypeFlag {
		req.EngineType = &builder.engineType

	}
	return req
}

type FileRecognizeSpeechReqBodyBuilder struct {
	speech     *Speech // 语音资源
	speechFlag bool

	config     *FileConfig // 配置属性
	configFlag bool
}

func NewFileRecognizeSpeechReqBodyBuilder() *FileRecognizeSpeechReqBodyBuilder {
	builder := &FileRecognizeSpeechReqBodyBuilder{}
	return builder
}

// 语音资源
//
//示例值：
func (builder *FileRecognizeSpeechReqBodyBuilder) Speech(speech *Speech) *FileRecognizeSpeechReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}

// 配置属性
//
//示例值：
func (builder *FileRecognizeSpeechReqBodyBuilder) Config(config *FileConfig) *FileRecognizeSpeechReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

func (builder *FileRecognizeSpeechReqBodyBuilder) Build() *FileRecognizeSpeechReqBody {
	req := &FileRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech
	}
	if builder.configFlag {
		req.Config = builder.config
	}
	return req
}

type FileRecognizeSpeechPathReqBodyBuilder struct {
	speech     *Speech
	speechFlag bool
	config     *FileConfig
	configFlag bool
}

func NewFileRecognizeSpeechPathReqBodyBuilder() *FileRecognizeSpeechPathReqBodyBuilder {
	builder := &FileRecognizeSpeechPathReqBodyBuilder{}
	return builder
}

// 语音资源
//
// 示例值：
func (builder *FileRecognizeSpeechPathReqBodyBuilder) Speech(speech *Speech) *FileRecognizeSpeechPathReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}

// 配置属性
//
// 示例值：
func (builder *FileRecognizeSpeechPathReqBodyBuilder) Config(config *FileConfig) *FileRecognizeSpeechPathReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

func (builder *FileRecognizeSpeechPathReqBodyBuilder) Build() (*FileRecognizeSpeechReqBody, error) {
	req := &FileRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech
	}
	if builder.configFlag {
		req.Config = builder.config
	}
	return req, nil
}

type FileRecognizeSpeechReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *FileRecognizeSpeechReqBody
}

func NewFileRecognizeSpeechReqBuilder() *FileRecognizeSpeechReqBuilder {
	builder := &FileRecognizeSpeechReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 语音文件识别接口，上传整段语音文件进行一次性识别。接口适合 60 秒以内音频识别
func (builder *FileRecognizeSpeechReqBuilder) Body(body *FileRecognizeSpeechReqBody) *FileRecognizeSpeechReqBuilder {
	builder.body = body
	return builder
}

func (builder *FileRecognizeSpeechReqBuilder) Build() *FileRecognizeSpeechReq {
	req := &FileRecognizeSpeechReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type FileRecognizeSpeechReqBody struct {
	Speech *Speech `json:"speech,omitempty"` // 语音资源

	Config *FileConfig `json:"config,omitempty"` // 配置属性
}

type FileRecognizeSpeechReq struct {
	apiReq *larkcore.ApiReq
	Body   *FileRecognizeSpeechReqBody `body:""`
}

type FileRecognizeSpeechRespData struct {
	RecognitionText *string `json:"recognition_text,omitempty"` // 语音识别后的文本信息
}

type FileRecognizeSpeechResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *FileRecognizeSpeechRespData `json:"data"` // 业务数据
}

func (resp *FileRecognizeSpeechResp) Success() bool {
	return resp.Code == 0
}

type StreamRecognizeSpeechReqBodyBuilder struct {
	speech     *Speech // 语音资源
	speechFlag bool

	config     *StreamConfig // 配置属性
	configFlag bool
}

func NewStreamRecognizeSpeechReqBodyBuilder() *StreamRecognizeSpeechReqBodyBuilder {
	builder := &StreamRecognizeSpeechReqBodyBuilder{}
	return builder
}

// 语音资源
//
//示例值：
func (builder *StreamRecognizeSpeechReqBodyBuilder) Speech(speech *Speech) *StreamRecognizeSpeechReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}

// 配置属性
//
//示例值：
func (builder *StreamRecognizeSpeechReqBodyBuilder) Config(config *StreamConfig) *StreamRecognizeSpeechReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

func (builder *StreamRecognizeSpeechReqBodyBuilder) Build() *StreamRecognizeSpeechReqBody {
	req := &StreamRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech
	}
	if builder.configFlag {
		req.Config = builder.config
	}
	return req
}

type StreamRecognizeSpeechPathReqBodyBuilder struct {
	speech     *Speech
	speechFlag bool
	config     *StreamConfig
	configFlag bool
}

func NewStreamRecognizeSpeechPathReqBodyBuilder() *StreamRecognizeSpeechPathReqBodyBuilder {
	builder := &StreamRecognizeSpeechPathReqBodyBuilder{}
	return builder
}

// 语音资源
//
// 示例值：
func (builder *StreamRecognizeSpeechPathReqBodyBuilder) Speech(speech *Speech) *StreamRecognizeSpeechPathReqBodyBuilder {
	builder.speech = speech
	builder.speechFlag = true
	return builder
}

// 配置属性
//
// 示例值：
func (builder *StreamRecognizeSpeechPathReqBodyBuilder) Config(config *StreamConfig) *StreamRecognizeSpeechPathReqBodyBuilder {
	builder.config = config
	builder.configFlag = true
	return builder
}

func (builder *StreamRecognizeSpeechPathReqBodyBuilder) Build() (*StreamRecognizeSpeechReqBody, error) {
	req := &StreamRecognizeSpeechReqBody{}
	if builder.speechFlag {
		req.Speech = builder.speech
	}
	if builder.configFlag {
		req.Config = builder.config
	}
	return req, nil
}

type StreamRecognizeSpeechReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *StreamRecognizeSpeechReqBody
}

func NewStreamRecognizeSpeechReqBuilder() *StreamRecognizeSpeechReqBuilder {
	builder := &StreamRecognizeSpeechReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 语音流式接口，将整个音频文件分片进行传入模型。能够实时返回数据。建议每个音频分片的大小为 100-200ms
func (builder *StreamRecognizeSpeechReqBuilder) Body(body *StreamRecognizeSpeechReqBody) *StreamRecognizeSpeechReqBuilder {
	builder.body = body
	return builder
}

func (builder *StreamRecognizeSpeechReqBuilder) Build() *StreamRecognizeSpeechReq {
	req := &StreamRecognizeSpeechReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type StreamRecognizeSpeechReqBody struct {
	Speech *Speech `json:"speech,omitempty"` // 语音资源

	Config *StreamConfig `json:"config,omitempty"` // 配置属性
}

type StreamRecognizeSpeechReq struct {
	apiReq *larkcore.ApiReq
	Body   *StreamRecognizeSpeechReqBody `body:""`
}

type StreamRecognizeSpeechRespData struct {
	StreamId *string `json:"stream_id,omitempty"` // 16 位 String 随机串作为同一数据流的标识

	SequenceId *int `json:"sequence_id,omitempty"` // 数据流分片的序号，序号从 0 开始，每次请求递增 1

	RecognitionText *string `json:"recognition_text,omitempty"` // 语音流识别后的文本信息
}

type StreamRecognizeSpeechResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *StreamRecognizeSpeechRespData `json:"data"` // 业务数据
}

func (resp *StreamRecognizeSpeechResp) Success() bool {
	return resp.Code == 0
}
