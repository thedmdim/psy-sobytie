# Frontend
## Stack
- hugo
- bulma css

## Location
- ./front

## Conventions
- content and templates are separated
- common styles for all pages live under ./static/base.css
- per page styles live inder separate files at ./static/<pagename>.css

## Instructions for agent (MUST!)
Your changes are atomic. You don't change more than 2 files at a time. You don't hardcode texts, all of them are places in yml header of corresponding page at ./front/content.

Do not search/explore the filesystem or templates unless the task explicitly requires it. When the task is a direct edit to known content (e.g. "change this value to that"), just make the edit using the data already in context — do not investigate how the value is used elsewhere first.

Do not webfetch (or otherwise look up) data the user already provided. If the user supplies a list, names, or content directly, use that as-is; never re-fetch the source to double-check unless the user asks.

Per page styles takes separate css file at ./front/static/css

Before changing layout/styles, prefer framework best practices over quick CSS overrides. For Bulma grid: keep `.column` as layout only, keep spacing via `.columns is-variable`, and put visual styles (background/border/radius) on inner wrapper like `.box`/`.card`.

In Hugo templates, avoid excessive temporary variables for straightforward paths. Prefer direct access like `.Params.sections.nav` unless variable materially reduces duplication or prevents context bugs.

Avoid speculative defensive code that has no matching markup or runtime path in current template (for example, element handlers without element itself). Keep JS/CSS minimal, purpose-driven, and tied to actual rendered elements.