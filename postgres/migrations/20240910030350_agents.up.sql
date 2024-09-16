CREATE TABLE IF NOT EXISTS agents (
    id          	UUID         	PRIMARY KEY,
	parent_id		UUID			REFERENCES agents(id) ON DELETE CASCADE,
	container_id	TEXT,
	status			TEXT			NOT NULL,
	description		TEXT			NOT NULL,
	instructions	TEXT,
	address			TEXT,
	client_id		TEXT			NOT NULL,
	client_secret	TEXT			NOT NULL,
	settings		JSONB			NOT NULL,
    created_at  	TIMESTAMPTZ    	NOT NULL,
    updated_at  	TIMESTAMPTZ    	NOT NULL
);
