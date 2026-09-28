# Design SQL Table

> [!NOTE]
> energy requirement: Medium Energy

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
