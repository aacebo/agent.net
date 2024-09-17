CREATE TABLE IF NOT EXISTS agent_logs (
	id          	UUID,
	agent_id		UUID			NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
	level			TEXT			NOT NULL,
	text			TEXT,
	data			JSONB,
    created_at  	TIMESTAMPTZ    	NOT NULL
) PARTITION BY LIST (agent_id);
