package go_ernie

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	utils "github.com/anhao/go-ernie/internal"
)

var (
	headerData  = []byte("data: ")
	errorPrefix = []byte(`{"error_code":`)
)

type streamable interface {
	ErnieBotResponse |
		ErnieBotTurboResponse |
		Bloomz7b1Response |
		LlamaChatResponse |
		BaiduChatResponse |
		ErnieCustomPluginResponse |
		ErnieBot4Response |
		CompletionResponse |
		ErnieBot8KResponse |
		ErnieBotTurboAIResponse
}
type streamReader[T streamable] struct {
	emptyMessagesLimit uint
	isFinished         bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator utils.ErrorAccumulator
	unmarshaler    utils.Unmarshaler
}

func (stream *streamReader[T]) Recv() (response T, err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

func (stream *streamReader[T]) processLines() (T, error) {
	var (
		emptyMessagesCount uint
		hasErrorPrefix     bool
		apiError           APIError
	)

	for {
		rawLine, readErr := stream.reader.ReadBytes('\n')
		if readErr != nil || hasErrorPrefix {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return *new(T), fmt.Errorf("error, %s", respErr.Error())
			}
			noSpaceLine := bytes.TrimSpace(rawLine)
			if bytes.HasPrefix(noSpaceLine, errorPrefix) {
				unmarshaler := utils.JSONUnmarshaler{}
				err := unmarshaler.Unmarshal(noSpaceLine, &apiError)
				if err != nil {
					return *new(T), err
				}
				return *new(T), &apiError
			}
			return *new(T), readErr
		}

		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, errorPrefix) {
			hasErrorPrefix = true
		}
		if !bytes.HasPrefix(noSpaceLine, headerData) || hasErrorPrefix {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}
			writeErr := stream.errAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return *new(T), writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.emptyMessagesLimit {
				return *new(T), ErrTooManyEmptyStreamMessages
			}

			continue
		}

		noPrefixLine := bytes.TrimPrefix(noSpaceLine, headerData)
		if len(noPrefixLine) == 0 {
			stream.isFinished = true
			return *new(T), io.EOF
		}

		var response T
		unmarshalErr := stream.unmarshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return *new(T), unmarshalErr
		}

		return response, nil
	}
}

func (stream *streamReader[T]) unmarshalError() (errResp *APIError) {
	errBytes := stream.errAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.unmarshaler.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	return
}

func (stream *streamReader[T]) Close() {
	stream.response.Body.Close()
}
