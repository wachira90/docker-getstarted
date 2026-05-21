package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// ฟังก์ชันสำหรับอ่านค่าจาก Docker Secret
func readSecret(secretName string) (string, error) {
	// กำหนด Path มาตรฐานของ Docker Secrets
	path := fmt.Sprintf("/run/secrets/%s", secretName)

	// อ่านไฟล์
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// ลบช่องว่างหรือบรรทัดใหม่ที่อาจติดมาด้วย
	return strings.TrimSpace(string(data)), nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Api!!")
	fmt.Println("Endpoint Hit: Api")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api", api)
	fmt.Println("REST STARTING .... ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// ดึงค่า Database Username
	dbUser, err := readSecret("db_user")
	if err != nil {
		log.Printf("Warning: Could not read db_user: %v\n", err)
	} else {
		fmt.Printf("Loaded DB User: %s\n", dbUser)
	}

	// ดึงค่า Database Password
	dbPassword, err := readSecret("db_password")
	if err != nil {
		log.Printf("Warning: Could not read db_password: %v\n", err)
	} else {
		// พิมพ์แค่ให้รู้ว่าโหลดสำเร็จ (ไม่ควร Print รหัสผ่านจริงออก Log ในระบบ Production)
		fmt.Println("Loaded DB Password: [REDACTED]")
	}

	// ตรงนี้สามารถนำ dbUser และ dbPassword ไปสร้าง Connection String เพื่อต่อ Database ได้เลย
	// ตัวอย่างเช่น: dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/dbname", dbUser, dbPassword)
	fmt.Sprintf("%s:%s", dbUser, dbPassword)

	handleRequests()
}
