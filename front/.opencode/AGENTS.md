# Frontend
## Stack
- framework: hugo
- styles: bulma css

## Location
- ./front

## Where things live (KNOW before making file lookups!)
- Usual hugo site dir structure
- Templates and content are separated
- Content lives under /content
- Content is YAML foramtted not hardcoded text
- Templates located at /layouts
- Styles located at /static
- Styles per page take separate css file at static/css/
- Bulma docs can be found at https://bulma.io/documentation/components and https://bulma.io/documentation/elements/

## Instructions for agent (MUST HAVE)
1. Your changes are atomic. You don't change more than 2 files at a time.
2. You don't hardcode texts, all of them are places in yml header of corresponding page at content/
3. Do not search/explore the filesystem or templates unless the task explicitly requires it. When the task is a direct edit to known content (e.g. "change this value to that"), just make the edit using the data already in context — do not investigate how the value is used elsewhere first.
4. Do not webfetch (or otherwise look up) data the user already provided. If the user supplies a list, names, or content directly, use that as-is; never re-fetch the source to double-check unless the user asks.
5. Before changing layout/styles, prefer framework best practices over quick CSS overrides. For Bulma grid: keep `.column` as layout only, keep spacing via `.columns is-variable`, and put visual styles (background/border/radius) on inner wrapper like `.box`/`.card`.
6. In Hugo templates, avoid excessive temporary variables for straightforward paths. Prefer direct access like `.Params.sections.nav` unless it prevents context bugs.
7. Avoid speculative defensive code that has no matching markup or runtime path in current template (for example, element handlers without element itself). Keep JS/CSS minimal, purpose-driven, and tied to actual rendered elements.
8. When user asks to change smth it almost always about changing style and template