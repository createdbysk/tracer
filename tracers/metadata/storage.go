package metadata

// Storage defines the interface for metadata storage
type Storage interface {
	// Store defines the function to store metadata to the specified path
	// path is expected to be / separated string.
	Store(path string, data interface{})
}
