package rpc

/*

TODO: codegen this

This bridges the jsonpb with ffjson.

*/

import "github.com/golang/protobuf/jsonpb"

func (j *QueueId) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *QueueId) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *ListQueuesRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *ListQueuesRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *ListQueuesResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *ListQueuesResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *GetQueueRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *GetQueueRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *GetQueueResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *GetQueueResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *CreateQueueRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *CreateQueueRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *CreateQueueResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *CreateQueueResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *DeleteQueueRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *DeleteQueueRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *ModifyQueueRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *ModifyQueueRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *PurgeQueueRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error)       { return j.MarshalJSON() }
func (j *PurgeQueueRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error { return j.UnmarshalJSON(b) }

func (j *MessageReceipt) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *MessageReceipt) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *FailedPublish) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *FailedPublish) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *Message) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *Message) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}


func (j *AckMessagesRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *AckMessagesRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *AckMessagesResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *AckMessagesResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *GetMessagesRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *GetMessagesRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *GetMessagesResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *GetMessagesResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *PublishMessagesRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *PublishMessagesRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *PublishMessagesResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *PublishMessagesResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *ModifyAckDeadlineRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *ModifyAckDeadlineRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *ModifyAckDeadlineResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *ModifyAckDeadlineResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *HealthcheckRequest) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *HealthcheckRequest) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

func (j *HealthcheckResponse) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) { return j.MarshalJSON() }
func (j *HealthcheckResponse) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	return j.UnmarshalJSON(b)
}

