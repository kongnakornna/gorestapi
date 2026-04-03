package cmd

import (
	"fmt"
	"log"
	"os"

	"icmongolang/config"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var (
	RootCmd = &cobra.Command{
		Use:   "go-base",
		Short: "A RESTful API dev",
		Long:  `A RESTful API dev with password less authentication.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-base.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//  กำหนดแฟลกและการตั้งค่า Config 
	// Cobra รองรับ persistent flags ซึ่งเมื่อกำหนดที่นี่แล้ว
	// จะเป็น global สำหรับแอปพลิเคชัน 
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "ไฟล์คอนฟิกูเรชัน (ค่าเริ่มต้นคือ $HOME/.go-base.yaml)")

	// Cobra ยังรองรับ local flags ซึ่งจะทำงานเฉพาะเมื่อ
	// เรียกใช้แอ็กชันนี้โดยตรงเท่านั้น
	// RootCmd.Flags().BoolP("toggle", "t", false, "ข้อความช่วยเหลือสำหรับ toggle")

	// Persistent flags - แฟลกที่ใช้ได้กับคำสั่งนี้และคำสั่งย่อยทั้งหมด
	// Local flags - แฟลกที่ใช้ได้กับคำสั่งนี้เท่านั้น
	// toggle -  เปิด/ปิด (true/false)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfgViper, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	_, err = config.ParseConfig(cfgViper)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
}
