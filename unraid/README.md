# xTeVe Channel Alerts – Unraid Community Apps

Monitors your xTeVe playlist each day and notifies a Discord channel if any
channels that previously existed have disappeared.
The container:

* Downloads the latest M3U from xTeVe
* Compares it with yesterday’s copy that is stored on a volume
* Sends a rich-embed Discord alert listing the missing channels
* Runs automatically on a cron schedule that **you control with an
environment variable** (no editing files in the container).

---

## 1 • Required Environment Variables

| Variable | Example | Description |
|----------|---------|-------------|
| `XTEVE_URL` | `http://192.168.1.100:34400/m3u/xteve.m3u` | Public or LAN URL to the M3U produced by xTeVe. |
| `DISCORD_WEBHOOK_URL` | `https://discord.com/api/webhooks/…` | Discord webhook where alerts will be posted. |
| `CRON_SCHEDULE` | `0 4 * * *` | Standard cron expression that controls when the check runs. Default: “0 4 * * *” (every day at 04:00). |

---

## 2 • Volume Mappings

| Container path | Purpose | Recommended Unraid mapping |
|----------------|---------|----------------------------|
| `/app/file/tmp` | Stores the *current* and *previous* M3U files so they persist between runs. | `/mnt/user/appdata/xteve-channel-alerts/tmp` |
| `/app/log` | Application & cron logs. | `/mnt/user/appdata/xteve-channel-alerts/logs` |

Both volumes are small (a few kB) but must be **persistent**.
If they are not mapped, the container will treat every run as the “first run”.

---

## 3 • Unraid Template Example

Use these values when creating a new container in the *Community Apps* GUI.

| Field | Value |
|-------|-------|
| Repository | `celsian/xteve-channel-alerts:latest` |
| Network Type | `Bridge` (or whatever suits your setup) |
| Console shell command | `/bin/sh` |
| Env ‑ `XTEVE_URL` | `http://192.168.1.100:34400/m3u/xteve.m3u` |
| Env ‑ `DISCORD_WEBHOOK_URL` | Your Discord webhook |
| Env ‑ `CRON_SCHEDULE` | `0 4 * * *` |
| /app/file/tmp | `/mnt/user/appdata/xteve-channel-alerts/tmp` |
| /app/log | `/mnt/user/appdata/xteve-channel-alerts/logs` |

Save, then **start** the container.
A first-run message will be sent to Discord letting you know that no
“previous” M3U existed yet.

---

## 4 • Cron Schedule Configuration

`CRON_SCHEDULE` accepts any standard five-field cron expression:

* `*/30 * * * *` → every 30 minutes
* `15 2 * * *` → daily at 02:15
* `0 */6 * * *` → every 6 hours

The entrypoint script rewrites `/etc/cron.d/xteve-cron` on every start so
updating the variable and restarting the container is all that’s needed.

---

## 5 • Troubleshooting

| Symptom | Likely Cause / Fix |
|---------|--------------------|
| **No messages appear in Discord** | • Webhook URL incorrect → regenerate in Discord. <br>• Container has no outbound internet → check network settings / firewall. |
| **Container log shows “Previous m3u file not found” every run** | Volume `/app/file/tmp` is **not** mapped persistently. Add the volume and keep it mounted. |
| **Cron appears to do nothing** | • Mis-typed `CRON_SCHEDULE`. Validate with an online cron tester. |
| **“xTeVe: cannot GET …” errors** | xTeVe URL unreachable from the container. Use IP instead of hostname, or set `Network=Host` in Unraid. |
| **Discord message truncated** | Discord embeds are limited to ~7 k characters. The app appends “Too many channels to list…”. Check container log for full list. |

---

## 6 • Updating & Backup

* Pull a new image in Unraid, stop, then start the container – volumes
  keep your data.
* Back up `/mnt/user/appdata/xteve-channel-alerts/` to preserve log history
  and the previous M3U snapshot.

Happy monitoring! 🎉
