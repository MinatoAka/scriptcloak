package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MinatoAka/mylog_go"
	"github.com/MinatoAka/scriptcloak/encoders"
	"github.com/MinatoAka/scriptcloak/templates"
)

func main() {
	in := flag.String("in", "", "Input script path")
	out := flag.String("out", "", "Output path (default: <in>.obf.<ext>)")
	lang := flag.String("lang", "", "Target language: js|python|shell|php")
	method := flag.String("method", "base64", "Encoder: base64|hex|unicode")
	verbose := flag.Bool("v", false, "Verbose output (debug mode)")
	flag.Parse()

	// Setup mylog — default hanya WARN/ERROR, verbose aktifkan DEBUG
	mylog.SetLevel(mylog.WARN)
	mylog.SetTimestamp(false)
	mylog.SetBold(true)
	if *verbose {
		mylog.SetLevel(mylog.DEBUG)
	}

	if *in == "" || *lang == "" {
		fmt.Println("Usage: scriptcloak -in script.js -out obf.js -lang js -method base64")
		flag.PrintDefaults()
		os.Exit(1)
	}

	*lang = strings.ToLower(*lang)
	*method = strings.ToLower(*method)

	if !supportedCombos(*lang, *method) {
		mylog.Error("Kombinasi lang=%q + method=%q tidak didukung", *lang, *method)
		os.Exit(1)
	}

	// Auto-generate out path kalau kosong
	if *out == "" {
		ext := filepath.Ext(*in)
		base := strings.TrimSuffix(*in, ext)
		*out = base + ".obf" + ext
	}

	// Safety: jangan sampai out == in
	if filepath.Clean(*in) == filepath.Clean(*out) {
		mylog.Error("Out path sama dengan input: %s", *out)
		os.Exit(1)
	}

	mylog.Debug("Reading input: %s", *in)
	source, err := os.ReadFile(*in)
	if err != nil {
		mylog.Error("Read input: %v", err)
		os.Exit(1)
	}
	mylog.Debug("Source size: %d bytes", len(source))

	// Encode
	mylog.Debug("Encoding with method: %s", *method)
	var payload string
	switch *method {
	case "base64":
		payload = encoders.Base64(string(source))
	case "hex":
		payload = encoders.Hex(string(source))
	case "unicode":
		payload = encoders.Unicode(string(source))
	}
	mylog.Debug("Payload size: %d bytes", len(payload))

	// Wrap sesuai bahasa
	mylog.Debug("Wrapping for language: %s", *lang)
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
	mylog.Debug("Output size: %d bytes", len(output))

	if err := os.WriteFile(*out, []byte(output), 0644); err != nil {
		mylog.Error("Write output: %v", err)
		os.Exit(1)
	}

	// Summary — tetap fmt.Println (laporan ke user)
	fmt.Println("Obfuscated successfully")
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
