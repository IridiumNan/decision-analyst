# Design SQL Table

> [!NOTE]  
> archive time: 2026-09-28 17:29:01.152082475 +0800 CST m=+1.123454479

---

## CONTEXT

Decision analyst project will use plugin architecture.

And the core is data storage and render interface.

For phrase one, the storage will be implemented by sqlite3.

The essentials will as below

- Title

- Summary

- Time

> Just mark the decision made time.
> for score or comment time, you should use plugins to append.

- Source

The source is the document that record the context and result and other details about this decision. And system should allow append or edit for this section.

Default format is markdown. (Html supported)

Must be plain text !!!

Remember that all source stored on the database should be lightweight text.

It's a summary for review and not contains all details.

The source itself is a long summary and review only

But wait, why the first version of summary and comment is the core of all these things ?

Why should we should 1-N relation on the event-source architecture ?

That's a essential question

> [!NOTE]
> So add a type or add a mark in the front of `Source` Field.

---

## Comment

The objective record and subjective things should be seperate.

So I use the minimal table for Event record finally.

See below table

```sql
-- This file is the system core schema
-- WARN: Don't touch it if you can add fields by plugins

CREATE TABLE Event (
    id INTERGER PRIMARY KEY,
    -- id is the distinct mark for each Event

    title TEXT,
    -- title contains a few words describe this event

    summary TEXT,
    -- summary is about 100 words about this event


    occur_time DATETIME,
    -- time record when this event happened
    
    create_time DATETIME,

    format TEXT CHECK (format in ('markdown', "html", "txt")),
    -- format of source data
    -- may support link later

    source TEXT,
    -- source contains all details about this evnet
    -- You can all some link on source
    -- It's format must be plain text or other parsable format
    -- default markdown, html, txt is supported
);
```

And check it on `internal/store/init.sql`

---

## Info

> id: 2a45a6184b2adefd
> create time: 2026-09-28 14:59:52.536799147 +0800 CST
> start time: 2026-09-28 17:28:28.583119658 +0800 CST
> duration: 20.674061076s
> context path: `/home/cai/WORK/project/decision_analyst/.project-todo/context/2a45a6184b2adefd-Design SQL Table.md`
