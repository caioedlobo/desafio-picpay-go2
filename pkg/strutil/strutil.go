package strutil

type Envelope map[string]any

func ErrorEnvelope(obj any) Envelope {
	return Envelope{"error": obj}
}
