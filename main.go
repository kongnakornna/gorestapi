// ไฟล์: main.go
// แก้ไขจากเดิมที่มีแค่ comment ให้ทำงานได้จริง

package main

// @title Go Rest API
// @version 1.0
// @description API for ICMON project
// @BasePath /api

// @securityDefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/auth/login
// @scope.read Grants read access
// @scope.write Grants write access

import (
	"icmongolang/cmd"
)

func main() {
	// เรียกใช้ root command ของ Cobra (กำหนดใน cmd/root.go)
	// Execute จะรันคำสั่งย่อยตาม argument เช่น serve, migrate, worker
	cmd.Execute()
}