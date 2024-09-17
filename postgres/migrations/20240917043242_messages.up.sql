CREATE TABLE IF NOT EXISTS messages (
    id          	UUID         	PRIMARY KEY,
	chat_id			UUID			NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
	from_id			UUID			NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
	text			text			NOT NULL,
    created_at  	TIMESTAMPTZ    	NOT NULL,
    updated_at  	TIMESTAMPTZ    	NOT NULL
);
