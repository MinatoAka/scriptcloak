package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/MinatoAka/scriptcloak/encoders"
	"github.com/MinatoAka/scriptcloak/templates"
)

func main() {
	in := flag.String("in", "", "Input script path")
	out := flag.String("out", "", "Output path (default: <in>.obf.<ext>)")
	lang := flag.String("lang", "", "Target language: js|python|shell|php")
	method := flag.String("method", "base64", "Encoder: base64|hex|unicode")
	flag.Parse()

	if *in == "" || *lang == "" {
		fmt.Println("Usage: scriptcloak -in script.js -out obf.js -lang js -method base64")
		flag.PrintDefaults()
		os.Exit(1)
	}

	*lang = strings.ToLower(*lang)
	*method = strings.ToLower(*method)

	// Validasi combo lang + method
	if !supportedCombos(*lang, *method) {
		log.Fatalf("[×] Kombinasi lang=%q + method=%q tidak didukung.\n   Coba: base64 untuk semua bahasa, hex untuk semua kecuali shell-perlu-xxd, unicode hanya JS/Python.", *lang, *method)
	}

	// Auto-generate out path kalau kosong
	if *out == "" {
		ext := filepath.Ext(*in)
		base := strings.TrimSuffix(*in, ext)
		*out = base + ".obf" + ext
	}

	source, err := os.ReadFile(*in)
	if err != nil {
		log.Fatalf("Read input: %v", err)
	}

	// Encode
	var payload string
	switch *method {
	case "base64":
		payload = encoders.Base64(string(source))
	case "hex":
		payload = encoders.Hex(string(source))
	case "unicode":
		payload = encoders.Unicode(string(source))
	}

	// Wrap sesuai bahasa
	var output string
	switch *lang {
	case "js":
		output = templates.JS(payload, *method)
	case "python":
		output = templates.Python(payload, *method)
	case "shell":
		output = templates.Shell(payload, *method)
	case "php":
		output = templates.PHP(payload, *method)
	}

	if err := os.WriteFile(*out, []byte(output), 0644); err != nil {
		log.Fatalf("Write output: %v", err)
	}

	fmt.Println("[✓] Obfuscated successfully")
	fmt.Printf("   In     : %s (%d bytes)\n", *in, len(source))
	fmt.Printf("   Out    : %s (%d bytes)\n", *out, len(output))
	fmt.Printf("   Lang   : %s\n", *lang)
	fmt.Printf("   Method : %s\n", *method)
}

func supportedCombos(lang, method string) bool {
	switch lang {
	case "js", "python":
		return method == "base64" || method == "hex" || method == "unicode"
	case "shell":
		return method == "base64" || method == "hex"
	case "php":
		return method == "base64" || method == "hex"
	}
	return false
}
