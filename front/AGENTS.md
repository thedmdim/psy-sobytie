# Frontend
## Stack
- hugo
- bulma css

## Location
- ./front

## Instructions for agent
Your changes are atomic. You don't change more than 2 files at a time. You don't hardcode texts, all of them are places in yml header of corresponding page at ./front/content.

Per page styles takes separate css file at ./front/static/css

Before changing layout/styles, prefer framework best practices over quick CSS overrides. For Bulma grid: keep `.column` as layout only, keep spacing via `.columns is-variable`, and put visual styles (background/border/radius) on inner wrapper like `.box`/`.card`.
