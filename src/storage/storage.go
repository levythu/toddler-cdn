package storage

type StorageAPI interface {
    // generate one instance of this type, using info in context.
    // for any failure, return nil
    New(context map[string]interface{}) StorageAPI

    // issue a PUT to the storage and any error should be returned with this
    // operation failed. It must be thread safe and idempotent.
    Put(key string, content []byte) error

    // issue a GET. If the key does not exist or any error is present, return nil
    // it must be threadsafe and idempotent.
    Get(key string) []byte

    // abandon the storage media. Remove occupied resources and release all the
    // states required. The StorageAPI should never be called if the func returned
    // with nil error
    Abandon() error
}
