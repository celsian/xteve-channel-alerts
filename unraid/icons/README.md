# Icon Conversion Guide

This directory contains the vector **SVG** source for the *xTeVe Channel Alerts*
icon (`xteve-channel-alerts.svg`).  
Unraid’s Community Apps store, however, expects a **PNG** image.  
Follow the steps below to generate the PNG that you will reference in the
template XML.

---

## Recommended Specs

| Property | Value |
|----------|-------|
| Size     | **256 × 256 px** (minimum 128 × 128 px) |
| Format   | PNG, 24-bit RGBA |
| File name| `xteve-channel-alerts.png` |

---

## 1 · Using Inkscape CLI  *(Linux / macOS / Windows WSL)*

```bash
# Install inkscape if you don’t have it
#   Debian/Ubuntu  : sudo apt install inkscape
#   macOS (brew)   : brew install inkscape

# Convert SVG → 256 px PNG
inkscape xteve-channel-alerts.svg \
        --export-type=png \
        --export-filename=xteve-channel-alerts.png \
        --export-width=256 \
        --export-height=256
```

---

## 2 · Using ImageMagick (convert)

ImageMagick will rasterise the SVG with the *librsvg* backend:

```bash
# Debian/Ubuntu: sudo apt install imagemagick librsvg2-bin
convert -background none \
        -resize 256x256 \
        xteve-channel-alerts.svg \
        xteve-channel-alerts.png
```

---

## 3 · Docker-based Conversion (no local installs)

```bash
docker run --rm -v "$PWD":/work -w /work \
           ghcr.io/julienr/inkscape-cli \
           inkscape xteve-channel-alerts.svg \
                    --export-type=png \
                    --export-filename=xteve-channel-alerts.png \
                    --export-width=256 \
                    --export-height=256
```

---

## 4 · Update the Unraid Template

1. Place `xteve-channel-alerts.png` in this `unraid/icons/` directory.
2. Open `unraid/template/xteve-channel-alerts.xml`.
3. Modify the `<Icon>` tag to point to the **raw** PNG file—e.g.

```xml
<Icon>https://raw.githubusercontent.com/celsian/xteve-channel-alerts/main/unraid/icons/xteve-channel-alerts.png</Icon>
```

4. Re-install / update the template in Community Apps.  
   The new PNG will appear as the application icon.

---

## 5 · Version Control

Keep **both** the SVG (source) and the generated PNG (binary) under version
control so that:

* Designers can tweak the vector file.
* Unraid users always have an up-to-date raster icon.

Happy hacking 🎉
