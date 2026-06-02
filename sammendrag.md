
Her er et sammendrag av kurset:

Sammendrag: KI-assistert utvikling med GitHub Copilot
Kurset er en praktisk innføring i å bruke KI-agenter i daglig utviklingsarbeid, basert på eksempelapplikasjonen tidtaker (en enkel timeregistrering).

Modul 1 – Start
Oppsett av GitHub Codespaces, kobling til VS Code og verifisering av Copilot-tilgang (eller alternativt OpenRouter API-nøkkel).

Modul 2 – Chat
Bli kjent med chat som arbeidsverktøy: stille spørsmål om kodebasen, gjøre små kontekstavhengige endringer og utforske teknologier (PocketBase, htmx, Tailwind).

Modul 3 – Kontekst
Lære å gi KI-modellen riktig kontekst: legge ved filer, markere spesifikke linjer og hente oppdatert informasjon fra internett for å unngå utdaterte svar.

Modul 4 – Agenter
Bruke agentmodus til å gjøre mer enn én ting per instruks (TDD-syklus, dokumentasjon, sikkerhetsvurdering, commit-melding). Automatisere gjentakende instrukser via AGENTS.md.

Modul 5 – Feilsøking
Beskrive feil i naturlig språk og la agenten finne og rette dem ("Tiden teller ikke når en starter tidtaker...").

Modul 6 – Utforskning
Bruke skills (.agents/skills/) for å strukturere komplekse utforskningsoppgaver. Eksempel: bruke grill-me-skillen til å planlegge automatisk timeføring via nettleserplugin eller Playwright.

Modul 7 – Planer
Lage detaljerte planer som agenten kan jobbe med selvstendig. Bruke Playwright MCP-server og OpenCode CLI til å styre en nettleser med naturlig språk, dokumentere klikkflyter og implementere funksjonen.

Modul 8 – Veien videre
Tips for å få gode resultater mens agenten jobber AFK: verifiserbarhet, parallelt arbeid med git worktrees, og feilsøkingsstrategier når agenten gjetter feil, setter seg fast eller ikke kjører testene.

Gjennomgående tema: Jo bedre instrukser og kontekst du gir, jo bedre resultater. Bruk AGENTS.md og skills for å redusere gjentakelse, og sørg alltid for at agenten har en måte å verifisere sitt eget arbeid på.
