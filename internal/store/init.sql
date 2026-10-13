-- This file is the system core schema
-- WARN: Don't touch it if you can add fields by plugins

CREATE TABLE Event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- id is the distinct mark for each Event

    title TEXT,
    -- title contains a few words describe this event

    summary TEXT,
    -- summary is about 100 words about this event


    occur_time DATETIME,
    -- time record when this event happened
    
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,

    format TEXT CHECK (format in ('markdown', "html", "txt")),
    -- format of source data
    -- may support link later

    source TEXT,
    -- source contains all details about this evnet
    -- You can all some link on source
    -- It's format must be plain text or other parsable format
    -- default markdown, html, txt is supported


    is_delete INTEGER DEFAULT 0
    -- mark if this event is deleted
    -- 0 is false
    -- 1 is true (means this event is deleted)
);
