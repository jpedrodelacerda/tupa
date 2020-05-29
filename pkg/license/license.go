package license

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// CreateLicense ...
func CreateLicense(license string, path string, author string) {
	fullPath := filepath.Join(path, "LICENSE")

	f, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var licenseString string
	switch license {
	case "gpl3", "gpl-3.0", "gpl-3":
		licenseString = initGPL3()
	case "lgpl3", "lgpl-3.0", "lgpl-3":
		licenseString = initLGPL3()
	case "apache", "apache2", "apache-2.0":
		licenseString = initApache2()
	case "wtfpl":
		licenseString = initWTFPL()
	case "unlicense":
		licenseString = initUnlicense()
	default:
		licenseString = initMIT(time.Now().Local().Year(), author)
	}

	f.WriteString(licenseString)
}
