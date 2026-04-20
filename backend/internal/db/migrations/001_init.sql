  -- Alerts table
  CREATE TABLE IF NOT EXISTS alerts (
      id              TEXT PRIMARY KEY,
      severity        TEXT NOT NULL,
      status          TEXT NOT NULL,
      description     TEXT NOT NULL,
      metric_name     TEXT NOT NULL,
      threshold       REAL NOT NULL,
      current_value   REAL NOT NULL,
      started_at      TIMESTAMPTZ,
      updated_at      TIMESTAMPTZ NOT NULL,
      acknowledged_at TIMESTAMPTZ,
      silenced_at     TIMESTAMPTZ,
      resolved_at     TIMESTAMPTZ
  );

  -- Actions table
  CREATE TABLE IF NOT EXISTS actions (
      id           TEXT PRIMARY KEY,
      type         TEXT NOT NULL,
      target       TEXT NOT NULL,
      parameters   TEXT NOT NULL,
      user         TEXT NOT NULL,
      request_time TIMESTAMPTZ NOT NULL,
      start_time   TIMESTAMPTZ,
      end_time     TIMESTAMPTZ,
      status       TEXT NOT NULL,
      output       TEXT,
      error        TEXT,
      risky        INTEGER NOT NULL
  );

  -- Metrics rollup table
  CREATE TABLE IF NOT EXISTS metrics_rollup (
      id         INTEGER PRIMARY KEY AUTOINCREMENT,
      timestamp  TIMESTAMPTZ NOT NULL,
      cpu_avg    REAL, cpu_max REAL,
      mem_avg    REAL, mem_max REAL,
      disk_avg   REAL, disk_max REAL,
      net_tx_avg REAL, net_tx_max REAL,
      net_rx_avg REAL, net_rx_max REAL,
      bucket     TEXT NOT NULL
  );

  -- History events table
  CREATE TABLE IF NOT EXISTS history_events (
      id              TEXT PRIMARY KEY,
      timestamp       TIMESTAMPTZ NOT NULL,
      source          TEXT NOT NULL,
      event_type      TEXT NOT NULL,
      alert_id        TEXT,
      action_id       TEXT,
      alert_severity  TEXT,
      alert_status    TEXT,
      action_type     TEXT,
      action_status   TEXT,
      latency_ms      INTEGER,
      details         TEXT
  );

  -- Schema migrations tracker
  CREATE TABLE IF NOT EXISTS schema_migrations (
      version   TEXT PRIMARY KEY,
      applied_at TIMESTAMPTZ NOT NULL
  );
