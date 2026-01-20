CREATE TABLE IF NOT EXISTS collections (
  name TEXT PRIMARY KEY,
  state TEXT,
  last_query_at DATETIME,
  snapshot_path TEXT
);
