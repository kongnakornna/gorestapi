ทกอยางเหมือนเดิมหมดแค่ เปลียนมาใช้ table sd_user
Module : KPI PMS
# URL Local http://localhost:8088 

# โครงสร้าง Folder  พร้อม คอมเมน้นอธิบาย  ภาษาไทย และ ภาษาอังกถษ
เช่น

## 1. โครงสร้างโฟลเดอร์ (Folder Structure)

```text
internal/user/                                 # root module ของ user
│
├── users/                                      
│   ├── delivery/                               
│   │   ├── http/
│   │   │   ├── handlers.go                    
│   │   │   └── routes.go                                      
│   ├── distributor/                          
│   │   └── distributor.go                     
│   ├── presenter/                              
│   │   └── presenters.go        
│   ├── processor/                              
│   │   └── processor.go                   
│   ├── repository/                            
│   │   ├── pg_repository.go                    
│   │   └── redis_repository.go                 
│   └── usecase/                                
│       └── usecase.go                                           
│
├── handler.go                                  
├── pg_repository.go                           
├── redis_repository.go        
├── usecase.go                       
└── worker.go                                  
```

**คำอธิบายแต่ละโฟลเดอร์/ไฟล์** (ไทย/อังกฤษ):
	1.วัตุประสงค์
	2.กลุ่มเป้าหมาย
	3.ความรู้พื้นฐาน
	4.เนื้อหา โดยย่อ กระชับ เน้น วัตถุประสงค์  ประโยชน์ของการใช้
	5.สร้างบทนำ
        - หลักการ (Concept)
        - คืออะไร?
        - มีกี่แบบ?  
        - ใช้อย่างไร / นำไปใช้กรณีไหน
        - ประโยชน์ที่ได้รับ
        - ข้อควรระวัง
        - ข้อดี
        - ข้อเสีย
        - ข้อห้าม
	6.สร้างบทนิยาม
	7.สร้างบทหัวข้อ
	8.ออกแบบคู่มือ
	9.ออกแบบ workflow
	10.TASK LIST Template
	11.CHECKLIST Template
	12.สรุป


# พร้อมสร้างเอกสาร
## หลักการ (Concept)
### คืออะไร?
### มีกี่แบบ?  
### ใช้อย่างไร / นำไปใช้กรณีไหน
### ประโยชน์ที่ได้รับ
### ข้อควรระวัง
### ข้อดี
### ข้อเสีย
### ข้อห้าม
**ข้อห้ามสำคัญ**

## คอมเม้น CODE ไทย อังกถษ คนละบรรทัด
## การออกแบบ Workflow และ Dataflow ระวัง อักขระ พิเศษ จำทำให้รูปแสดงไม่ได้ ระวังให้ดี
## คู่มือการทดสอบ
## คู่มือการการใช้งาน
## คู่มือการบำรุงรักษา
## คู่มือการขยาย หรือแก้ไข หรือ เพิมเติม ในอนาคต
## CHECK List Test Module
