# scriptcloak

Simple obfuscator for scripts. Wraps source code in an encoded payload + decoder stub.

> ⚠️ **WARNING**: This is **NOT encryption**. It only prevents casual copy-paste. Anyone with basic skills can decode the output (browser devtools, Python disassembler, etc.). Use it to make reverse engineering *annoying*, not *impossible*.

## Install

```bash
go install github.com/MinatoAka/scriptcloak@latest
```
Or build local:
```bash
git clone https://github.com/MinatoAka/scriptcloak
cd scriptcloak
go build -o scriptcloak
```
## Usage
```bash
# JavaScript with base64
./scriptcloak -in app.js -out app.obf.js -lang js -method base64

# Python with hex
./scriptcloak -in bot.py -out bot.obf.py -lang python -method hex

# Shell with base64
./scriptcloak -in deploy.sh -out deploy.obf.sh -lang shell -method base64
```

## Limitations
- Not encryption — trivially reversible
- Shell hex needs xxd
- Unicode only supports BMP (< U+10000)
- Larger output (base64: 1.33x, hex: 4x, unicode: 6x)
