# Decision Analyst

## Why

We decision as who we are. For analysing myself more efficiently, a decision recorder and analyst is a great way.

However, most decision are plain and most people would do. It's not distinguished.

**I want to analysis decisions that make me different from other.**

So, I would score an decision with $Score = - log_2 p \cdot e^i$. If this decision rarely happened, it will take more information about who am I and how I learn about the world.

By the way, I hold the view that I'm really different with most of people.

Nixos instead of Windows, hyprland instead of gnome, left-hand mouse instead of smooth right-hand, Lazyvim instead of vscode, minimal configuration instead of fancy one, etc...

---

## Core

### Score

This project will feed the self-host LLM or cloud one with your plain documents (markdown, html, txt...).

Then the LLM extract main decision on this documents, evaluate it's rough possibility and impact (if you not offer this impact)

> If you have more efficiently method, write a issue for recommendation please

Based on the possibility of this decision, a score will be calculate by equation

$s(p, i) = - log_2 p \cdot e^i$

where

- $s$ is the **score** of this decision

- $p$ is **possibility** that this desicion happened on this cases (user is recommonded list all alternatives when decide)

- $i$ is the **impact** this decision makes. It's an integer range from [0, 3]
  
  - 0, rare impact
  
  - 1, you can get some benefit from it for a period of time
  
  - 2, benefits you for a long lasting peroid or makes you grow rapidly
  
  - 3, life long lasting benefits that impact massive decision after that.

This just a score for rarity. You can define more dimensions like effect time, experience score etc...

**This system support manually score and weight configuration about this decision.**

### Time Line

The system will load time from a single configuration file, if not provided, it will use loaded time.

It store the time for future visualization with time line.

### Visualization

Offer A time line visualization chart on browser. Then you can select and extract some of them so that provide your personal evidence when talking about the LLM or analysis them further.

### Append

For a decision you make before, you would have more thoughts about it. The system will store the raw context first, then you can append some comment documents or text then system will store them as well.

---

## Architecture

TODO

---

## Quick Start

---
