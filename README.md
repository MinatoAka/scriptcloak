# scriptcloak

Simple obfuscator for scripts. Wraps source code in an encoded payload + decoder stub.

> ⚠️ **WARNING**: This is **NOT encryption**. It only prevents casual copy-paste. Anyone with basic skills can decode the output (browser devtools, Python disassembler, etc.). Use it to make reverse engineering *annoying*, not *impossible*.

## Install

```bash
go install github.com/MinatoAka/scriptcloak@latest
