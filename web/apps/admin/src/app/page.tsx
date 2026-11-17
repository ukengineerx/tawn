import { NotesPanel } from "@/components/notes-panel";

const apiURL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export default function HomePage() {
  return (
    <main className="mx-auto flex min-h-screen max-w-xl flex-col justify-center gap-8 px-6 py-16">
      <header>
        <p className="text-sm font-medium text-[var(--muted)]">Tawn</p>
        <h1 className="mt-2 text-3xl font-semibold tracking-tight">Admin</h1>
        <p className="mt-2 text-sm text-[var(--muted)]">
          One screen wired to the sample notes API. Replace this when the product has a real workflow.
        </p>
      </header>
      <NotesPanel apiURL={apiURL} />
    </main>
  );
}
