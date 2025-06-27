## Requirements
 - The service should support different log levels: INFO, WARNING, ERROR, and DEBUG.
 - Allow logging messages with timestamps and log levels to a file or the console.
 - Enable filtering of logs by level or specific keywords.
 - Should include an API to dynamically change log levels or configurations at runtime.
 - Support configurable log rotation based on file size or time intervals.
 - Provide thread-safe logging for concurrent applications.
 - Handle initialization gracefully, creating log directories/files if missing.
 - Implement extensibility for adding new output sinks (e.g., external systems like Elasticsearch).
