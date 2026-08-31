package types

type PostgresPing struct {
	Connected bool   `json:"connected"`
	LatencyMs int64  `json:"latency_ms"`
	Message   string `json:"message,omitempty"`
	Address   string `json:"address,omitempty"`
	Version   string `json:"version,omitempty"`
}
type PostgresServerInfo struct {
	Version             string  `json:"version"`
	UptimeSeconds       int64   `json:"uptime_seconds"`
	Connections         int64   `json:"connections"`
	MaxConnections      int64   `json:"max_connections"`
	TPS                 float64 `json:"tps"`
	TotalTransactions   int64   `json:"total_transactions"`
	DatabaseCount       int     `json:"database_count"`
	UserDatabaseCount   int     `json:"user_database_count"`
	DataSizeBytes       int64   `json:"data_size_bytes"`
	CacheHitRatio       float64 `json:"cache_hit_ratio"`
	ActiveSessions      int64   `json:"active_sessions"`
	IdleSessions        int64   `json:"idle_sessions"`
	IdleInTxSessions    int64   `json:"idle_in_tx_sessions"`
	WaitingSessions     int64   `json:"waiting_sessions"`
	BlockingSessions    int64   `json:"blocking_sessions"`
	BlockedSessions     int64   `json:"blocked_sessions"`
	Deadlocks           int64   `json:"deadlocks"`
	TempFiles           int64   `json:"temp_files"`
	TempBytes           int64   `json:"temp_bytes"`
	SharedBuffersBytes  int64   `json:"shared_buffers_bytes"`
	EffectiveCacheBytes int64   `json:"effective_cache_bytes"`
	SlowQueries         int64   `json:"slow_queries"`
	PgStatStatements    bool    `json:"pg_stat_statements"`
	InRecovery          bool    `json:"in_recovery"`
}
type PostgresDatabase struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	SizeBytes int64  `json:"size_bytes"`
	Encoding  string `json:"encoding"`
	Collation string `json:"collation"`
}
type PostgresSchema struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
type PostgresTable struct {
	Schema    string `json:"schema"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Rows      int64  `json:"rows"`
	SizeBytes int64  `json:"size_bytes"`
}
type PostgresQueryRequest struct {
	Schema string `json:"schema"`
	SQL    string `json:"sql" binding:"required"`
	Limit  int64  `json:"limit,omitempty"`
}
type PostgresQueryResult struct {
	Columns   []string        `json:"columns,omitempty"`
	Rows      [][]interface{} `json:"rows,omitempty"`
	Affected  int64           `json:"affected_rows"`
	Duration  int64           `json:"duration_ms"`
	Truncated bool            `json:"truncated,omitempty"`
	Statement string          `json:"statement"`
}
type PostgresSession struct {
	PID             int64   `json:"pid"`
	User            string  `json:"user"`
	ClientAddr      string  `json:"client_addr"`
	Database        string  `json:"database"`
	Application     string  `json:"application"`
	State           string  `json:"state"`
	WaitEvent       string  `json:"wait_event"`
	QueryStart      string  `json:"query_start"`
	DurationSeconds float64 `json:"duration_seconds"`
	Query           string  `json:"query"`
}
