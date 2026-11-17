"use client";

import { Button } from "@tawn/ui";
import { useEffect, useState, type FormEvent } from "react";

type Note = { id: string; body: string; created_at: string };
type Health = { status?: string; postgres?: string };

export function NotesPanel({ apiURL }: { apiURL: string }) {
  const [health, setHealth] = useState<Health | null>(null);
  const [notes, setNotes] = useState<Note[]>([]);
  const [body, setBody] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    const controller = new AbortController();
    fetch(`${apiURL}/ready`, { signal: controller.signal })
      .then(async (ready) => {
        setHealth((await ready.json()) as Health);
        const listed = await fetch(`${apiURL}/v1/notes`, { signal: controller.signal });
        if (!listed.ok) {
          setNotes([]);
          setError("Notes API is unavailable. Start Postgres, then the API.");
          return;
        }
        const payload = (await listed.json()) as { notes: Note[] };
        setNotes(payload.notes);
        setError("");
      })
      .catch((error: unknown) => {
        if (error instanceof DOMException && error.name === "AbortError") return;
        setHealth(null);
        setError("API is not reachable at " + apiURL);
      });
    return () => controller.abort();
  }, [apiURL]);

  async function onSubmit(event: FormEvent) {
    event.preventDefault();
    setError("");
    const response = await fetch(`${apiURL}/v1/notes`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ body }),
    });
    if (!response.ok) {
      const payload = (await response.json().catch(() => null)) as { error?: string } | null;
      setError(payload?.error ?? "Could not save the note");
      return;
    }
    setBody("");
    const listed = await fetch(`${apiURL}/v1/notes`);
    if (listed.ok) {
      const payload = (await listed.json()) as { notes: Note[] };
      setNotes(payload.notes);
    }
  }

  return (
    <section className="rounded-xl border border-[var(--line)] bg-[var(--surface)] p-5">
      <p className="text-sm text-[var(--muted)]">
        API {health ? `${health.status ?? "unknown"}${health.postgres ? ` · postgres ${health.postgres}` : ""}` : "offline"}
      </p>
      <form onSubmit={onSubmit} className="mt-4 flex gap-2">
        <input
          value={body}
          onChange={(event) => setBody(event.target.value)}
          placeholder="Write a note"
          className="min-w-0 flex-1 rounded-md border border-[var(--line)] bg-transparent px-3 py-2 text-sm outline-none"
        />
        <Button type="submit" disabled={body.trim() === ""}>
          Save
        </Button>
      </form>
      {error ? <p className="mt-3 text-sm text-red-700">{error}</p> : null}
      <ul className="mt-4 space-y-2">
        {notes.map((note) => (
          <li key={note.id} className="rounded-md bg-[var(--background)] px-3 py-2 text-sm">
            {note.body}
          </li>
        ))}
      </ul>
    </section>
  );
}
