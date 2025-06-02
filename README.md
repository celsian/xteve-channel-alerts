# xteve-channel-alerts

### Why?
Sometimes when I sit down to watch a specific channel I find it's missing from my xTeVe playlist.  
xTeVe refreshes the playlist nightly; this app is designed to run **after that update** and compare
today's channel list to yesterday's.  
If it detects any missing channels it sends an alert to a Discord webhook with the list of channels
that disappeared.

---

## Multi-Architecture Support ✅

The published Docker image `celsian/xteve-channel-alerts:latest` is a **multi-arch image**.  
It automatically pulls the correct build for your hardware:

| Architecture | Typical processors / hosts | Docker platform tag |
|--------------|---------------------------|---------------------|
| **AMD64**    | Intel & AMD 64-bit CPUs (most PCs/servers, Unraid) | `linux/amd64` |
| **ARM64**    | Apple Silicon (M1/M2/M3), Raspberry Pi 5 64-bit, ARM servers | `linux/arm64` |

No extra flags are required—Docker (and Unraid) will pick the right image
automatically.

---

### Setup
1. Copy the `.env.example` file to `.env`  
2. Populate the `.env` with your values  
3. Run the app  

