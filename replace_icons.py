#!/usr/bin/env python3
"""Replace icons2 references with appropriate icons3 icons in program files."""

import os

BASE = "/Users/maksimstupcenko/psy-sobytie/front/content/programs"

# Mapping: file -> [rocket_replacement, point_replacements_in_order]
mapping = {
    "adult-therapy-group/index.md": {
        "rocket": "iconadult.png",
        "points": ["iconbell.png", "iconnetwork.png"],
    },
    "association/index.md": {
        "rocket": "iconhandshake.png",
        "points": ["iconpuzzle.png", "iconbalance.png"],
    },
    "beautiful-childhood/index.md": {
        "rocket": "iconbaby.png",
        "points": ["iconbooks.png"],
    },
    "beautiful-growing-up/index.md": {
        "rocket": "iconteen.png",
        "points": ["iconlook.png"],
    },
    "become-a-partner/index.md": {
        "rocket": "iconhandshake.png",
        "points": ["iconbell.png", "icongear.png", "iconnetwork.png", "iconbalance.png", "iconfire.png"],
    },
    "career-guidance/index.md": {
        "rocket": "iconowl.png",
        "points": ["iconchess.png", "iconbooks.png"],
    },
    "child-psychologist/index.md": {
        "rocket": "icontoy.png",
        "points": ["iconlook.png", "iconowl.png"],
    },
    "corporate-training/index.md": {
        "rocket": "icongear.png",
        "points": ["iconlook.png", "iconbrush.png", "iconnetwork.png", "iconowl.png"],
    },
    "education-and-training/index.md": {
        "rocket": "iconbooks.png",
        "points": ["iconpuzzle.png", "iconbalance.png"],
    },
    "education/index.md": {
        "rocket": "iconbooks.png",
        "points": ["iconbell.png", "iconowl.png"],
    },
    "family-consultation/index.md": {
        "rocket": "iconhandshake.png",
        "points": ["iconnetwork.png", "iconcheir.png", "iconlook.png", "iconchess.png"],
    },
    "for-employees/index.md": {
        "rocket": "iconcheir.png",
        "points": ["iconbell.png", "iconbrush.png", "iconhandshake.png", "iconbalance.png", "icongear.png"],
    },
    "for-leaders/index.md": {
        "rocket": "iconowl.png",
        "points": ["iconlook.png", "iconpuzzle.png", "iconchess.png", "iconnetwork.png", "iconfire.png"],
    },
    "individual-therapy/index.md": {
        "rocket": "iconadult.png",
        "points": ["iconlook.png", "iconbell.png", "iconhandshake.png", "iconfire.png"],
    },
    "invite-us-to-school/index.md": {
        "rocket": "iconbooks.png",
        "points": ["iconhandshake.png", "iconnetwork.png", "iconlook.png", "iconpuzzle.png"],
    },
    "mediation/index.md": {
        "rocket": "iconbalance.png",
        "points": ["iconnetwork.png", "iconlook.png", "iconhandshake.png", "iconpuzzle.png", "iconowl.png"],
    },
    "organizational-integration/index.md": {
        "rocket": "iconhandshake.png",
        "points": ["iconbrush.png", "iconbalance.png", "icongear.png", "iconlook.png", "iconpuzzle.png"],
    },
    "parents/index.md": {
        "rocket": "iconadult.png",
        "points": ["iconbell.png", "iconnetwork.png", "iconbrush.png", "iconfire.png", "iconbalance.png"],
    },
    "psychology-integration/index.md": {
        "rocket": "icongear.png",
        "points": ["iconowl.png", "iconhandshake.png", "iconfire.png", "iconlook.png", "iconpuzzle.png"],
    },
    "psychotherapy-for-psychologists/index.md": {
        "rocket": "iconadult.png",
        "points": ["iconlook.png", "iconbell.png", "iconowl.png", "iconcheir.png", "iconfire.png"],
    },
    "summer-camp/index.md": {
        "rocket": "iconfire.png",
        "points": ["iconbrush.png", "iconbell.png", "iconbalance.png", "icontoy.png"],
    },
    "supervision/index.md": {
        "rocket": "iconlook.png",
        "points": ["iconadult.png", "iconnetwork.png", "iconpuzzle.png", "iconbell.png", "iconcheir.png"],
    },
    "team-engagement/index.md": {
        "rocket": "iconnetwork.png",
        "points": ["iconlook.png", "iconbrush.png", "iconfire.png", "iconhandshake.png", "iconpuzzle.png"],
    },
    "teen-psychologist/index.md": {
        "rocket": "iconteen.png",
        "points": ["iconhandshake.png", "iconbell.png", "iconlook.png", "iconnetwork.png", "iconfire.png"],
    },
    "temp-klim/index.md": {
        "rocket": "iconfire.png",
        "points": ["iconlook.png", "iconnetwork.png", "iconbooks.png", "iconowl.png", "icongear.png"],
    },
}

for rel_path, data in mapping.items():
    filepath = os.path.join(BASE, rel_path)
    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    # Replace rocket (first occurrence)
    if "rocket" in data:
        content = content.replace(
            'image: "/icons2/rocket.png"',
            f'image: "/icons3/{data["rocket"]}"',
            1,
        )

    # Replace points in order
    for point_icon in data["points"]:
        content = content.replace(
            'image: "/icons2/point.png"',
            f'image: "/icons3/{point_icon}"',
            1,
        )

    with open(filepath, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"OK: {rel_path}")

print("\nDone!")