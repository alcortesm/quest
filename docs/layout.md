 # Quest layout

```
~/quests/2026-10-29_api-hang/        # the root is configurable
  quest.toml        # manifest (data only, never executed)
  QUEST.md          # the brief: written by you, trusted
  context/          # inputs you add: logs, traces, Slack exports
  notes/            # your notes
  repos/<name>/     # TRUSTED worktrees or clones, branch quest/<id>
  results/<run-id>/ # vetted output of runs (agent claims)
  runs/<run-id>/    # UNTRUSTED zone, one per agent attempt
    in/             #   copy of context + TASK.md   (mounted read-only)
    repos/<name>/   #   fresh scratch clones         (mounted read-write)
    out/            #   agent output files           (mounted read-write)
    state/          #   agent session state          (mounted read-write)
    bundles/        #   git bundles made by the export phase
  .quest/           # log.jsonl, run records, lock
```

The host never runs anything inside `runs/`.
Only `quest run harvest` reads it, under the intake rules in section 12.
`runs/` is never archived, and `results/` is.

`runs/<run-id>/state/` is **proposed**: it holds the agent sessions so they come back as normal files (confirmed in session 1.5).
It is treated like `out/`: only `quest run harvest` reads it, and only regular files are copied to `results/<run-id>/state/`.

Config and state of the tool itself (not of a quest):

| Path | Content |
| --- | --- |
| `$XDG_CONFIG_HOME/quest/config.toml` (default `~/.config/quest/config.toml`) | Configuration. Same path on Linux and macOS. |
| `$XDG_STATE_HOME/quest/` (default `~/.local/state/quest/`) | Reserved. Empty in version 1. |
