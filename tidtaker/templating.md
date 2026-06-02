# Go `html/template` – rask referanse

Dette prosjektet bruker Go's innebygde `html/template`-pakke. Den er syntaktisk lik `text/template`, men HTML-escaper automatisk verdier for å hindre XSS.

**Dokumentasjon:** https://pkg.go.dev/html/template@go1.26.3

---

## Grunnleggende syntaks

| Hva | Syntaks |
|-----|---------|
| Skriv ut verdi | `{{.Felt}}` |
| Kall metode | `{{.Metode "arg"}}` |
| Kommentar | `{{/* kommentar */}}` |
| Rå streng (ikke escape) | `{{printf "%s" .Verdi}}` |

---

## Betingelser

```html
{{if .Aktiv}}
  <span>Aktiv</span>
{{else if .Pause}}
  <span>Pause</span>
{{else}}
  <span>Inaktiv</span>
{{end}}
```

## Løkker

```html
{{range .Liste}}
  <li>{{.}}</li>       <!-- . er nåværende element -->
{{end}}

{{range .Liste}}
  <li>{{$.GlobalFelt}} – {{.}}</li>  <!-- $ gir tilgang til rot-data -->
{{end}}
```

## Variabler

```html
{{$navn := .Bruker.Navn}}
<p>Hei, {{$navn}}!</p>
```

---

## Definere og bruke blokker / deler

**Definer en blokk (i en fil):**
```html
{{define "mitt-partial"}}
  <div>{{.Tittel}}</div>
{{end}}
```

**Bruk den i en annen template:**
```html
{{template "mitt-partial" .}}

<!-- Med eget data-objekt: -->
{{template "mitt-partial" dict "Tittel" "Hei"}}
```

**Overriderbar blokk (layout-mønster):**
```html
<!-- base.html -->
{{define "base"}}
<html>
  <body>{{block "content" .}}{{end}}</body>
</html>
{{end}}

<!-- side.html -->
{{define "content"}}
  <h1>Min side</h1>
{{end}}
```

---

## Template-funksjoner (tilgjengelig i dette prosjektet)

| Funksjon | Eksempel | Beskrivelse |
|----------|---------|-------------|
| `formatDuration` | `{{formatDuration .Start .Stop}}` | Viser varighet som `1t 23m 45s` |
| `formatTime` | `{{formatTime .StartTime}}` | Viser klokkeslett `HH:MM:SS` |
| `formatDate` | `{{formatDate .StartTime}}` | Viser dato `DD.MM.YYYY` |
| `localTime` | `{{localTime .StartTime}}` | ISO-format for `datetime-local` input |
| `add` / `sub` | `{{add .Index 1}}` | Enkel aritmetikk |
| `seq` | `{{range seq 5}}` | Lager liste `[0,1,2,3,4]` |
| `contains` | `{{if contains .Tags "viktig"}}` | Sjekker om slice inneholder element |
| `eq` | `{{if eq .Sort "-startTime"}}` | Sammenligner to strenger |
| `dict` | `{{template "x" dict "Key" .Val}}` | Lager map for å sende flere verdier til partial |

---

## Sende data til et partial med `dict`

Siden Go templates bare kan sende **ett** dataobjekt til `{{template}}`, brukes `dict` for å pakke flere verdier:

```html
{{template "timing_item" dict "Timing" . "CurrentUser" $.User}}
```

I mottaker:
```html
{{define "timing_item"}}
  {{.Timing.Id}} – {{.CurrentUser.Navn}}
{{end}}
```

---

## Kalle metoder på objekter fra PocketBase

PocketBase-poster (`*core.Record`) har ikke vanlige Go-felt – verdiene ligger i databasen og hentes via metoder:

| Metode | Brukes til |
|--------|-----------|
| `.GetString "feltNavn"` | Tekstverdi |
| `.GetBool "feltNavn"` | Boolean |
| `.GetDateTime "feltNavn"` | Dato/tid |
| `.GetStringSlice "feltNavn"` | Liste av strenger (f.eks. tags) |

**Eksempel – `{{if .Timing.GetBool "isActive"}}`:**

Verdien hentes i tre lag:

1. **Databaseskjema** (`pb_migrations/timings.go`) – feltet er definert som `BoolField`:
   ```go
   &core.BoolField{Name: "isActive"}
   ```

2. **Handler** (`handlers_timer.go`) – setter feltet på posten og sender den til templaten:
   ```go
   record.Set("isActive", true)
   // ...
   renderPartial(e, "timing_item", "timing_item", map[string]any{
       "Timing": record,
   })
   ```

3. **Template** (`timing_item.html`) – leser verdien via metode:
   ```html
   {{if .Timing.GetBool "isActive"}}
     <span class="text-green-500">● aktiv</span>
   {{end}}
   ```

`.Timing` er altså et `*core.Record`-objekt der verdiene er lagret i SQLite og leses dynamisk via metoder.

---

## Tilgang til rot-data inne i `range` med `$`

```html
{{range .Timings}}
  <div id="timing-{{.Id}}">
    Bruker: {{$.User.Navn}}   <!-- $ = rot-objektet -->
    Tid: {{.Duration}}
  </div>
{{end}}
```
