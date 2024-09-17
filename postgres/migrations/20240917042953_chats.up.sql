CREATE TABLE IF NOT EXISTS chats (
    id          	UUID         	PRIMARY KEY,
	agent_id		UUID			NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    created_at  	TIMESTAMPTZ    	NOT NULL,
    updated_at  	TIMESTAMPTZ    	NOT NULL
);
