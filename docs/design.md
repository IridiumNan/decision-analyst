# Design Document

## Plugin

For the original event record data, there may be a lot of way to explain or rank them.

It's countless when I start this project.

So I use plugin architecture which the core just store objective data and summary.

The core just tell us what happened and when it happened.

For why this happened, what's its effect, how important it's, they should be handled by different plugins.

> [!NOTE]
> So if you check the `internal/store/init.sql`
> It's a minimal table for event record

---
