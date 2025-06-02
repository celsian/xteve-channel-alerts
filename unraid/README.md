# xTeVe Channel Alerts – Unraid Community Apps

Monitors your xTeVe playlist each day and notifies a Discord channel if any
channels that previously existed have disappeared.  
The container:

* Downloads the latest M3U from xTeVe  
* Compares it with yesterday’s copy that is stored on a **single `data/`
  volume** (separated into `logs/` and `m3us/` sub-directories)  
* Sends a rich-embed Discord alert listing the missing channels  
* Runs automatically on a cron schedule that **you control with an
  environment variable** (no editing files in the container).

---

## Multi-Architecture Support ✅

`celsian/xteve-channel-alerts` is published as a **multi-arch image** and
automatically pulls the correct build for your hardware:

| Architecture | Typical hardware                        | Docker platform tag |
|--------------|-----------------------------------------|---------------------|
| **AMD64**    | Intel / AMD 64-bit CPUs (Unraid, servers) | `linux/amd64`       |
| **ARM64**    | Apple Silicon (M1/M2/M3) & ARM servers   | `linux/arm64`       |

---

## 1 • Required Environment Variables

| Variable | Example | Description |
|----------|---------|-------------|
| `XTEVE_URL` | `http://192.168.1.100:34400/m3u/xteve.m3u` | Public or LAN URL produced by xTeVe |
| `DISCORD_WEBHOOK_URL` | `https://discord.com/api/webhooks/…` | Discord webhook that receives alerts |
| `CRON_SCHEDULE` | `0 4 * * *` | Cron expression controlling when the check runs |

---

## 2 • Volume Mapping

All persistent data lives in **one host directory** that the container
organises internally:

```
data/
├── logs/   ──▶  app.log , cron.log
└── m3us/   ──▶  current.m3u , previous.m3u
```

| Container path | Contains … | Must persist? | Recommended Unraid path |
|----------------|------------|---------------|-------------------------|
| `/app/data`    | `logs/` & `m3us/` sub-directories | **Yes** | `/mnt/user/appdata/xteve-channel-alerts/data` |

If this volume is **not** mapped persistently the container will treat every
run as the first run and logs will be lost.

---

## 3 • Unraid Template Example

| Field | Value |
|-------|-------|
| Repository | `celsian/xteve-channel-alerts:latest` |
| Network Type | `Bridge` (or whichever suits your setup) |
| Console shell command | `/bin/sh` |
| Env – `XTEVE_URL` | `http://192.168.1.100:34400/m3u/xteve.m3u` |
| Env – `DISCORD_WEBHOOK_URL` | *your Discord webhook* |
| Env – `CRON_SCHEDULE` | `0 4 * * *` |
| `/app/data` | `/mnt/user/appdata/xteve-channel-alerts/data` |

Save, then **start** the container. A first-run message will be sent to
Discord noting that no previous M3U existed yet.

---

## 4 • Cron Schedule

`CRON_SCHEDULE` accepts any standard five-field cron expression:

| Example | Meaning |
|---------|---------|
| `*/30 * * * *` | every 30 minutes |
| `15 2 * * *` | daily at 02:15 |
| `0 */6 * * *` | every 6 hours |

Change the variable and restart the container to apply.

---

## 5 • Troubleshooting

| Symptom | Likely cause / fix |
|---------|--------------------|
| No Discord messages | Incorrect webhook URL or outbound network blocked |
| “Previous m3u file not found” every run | `/app/data` not mapped persistently |
| Cron appears to do nothing | Mis-typed `CRON_SCHEDULE`; validate online |
| `xTeVe: cannot GET …` errors | xTeVe URL unreachable; use IP or `Network:Host` |
| Discord message truncated | Discord embeds limit ~7 kB; see full list in `data/logs/app.log` |

---

## 6 • Updating & Backup

* **Update** – pull the new image in Unraid, then restart; `/app/data`
  retains all state.
* **Backup** – copy
  `/mnt/user/appdata/xteve-channel-alerts/data/`  
  to preserve log history and the previous M3U snapshot.

Happy monitoring 🎉
