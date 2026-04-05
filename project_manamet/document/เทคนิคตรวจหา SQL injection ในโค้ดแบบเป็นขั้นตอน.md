 

# เทคนิคตรวจหา SQL injection ในโค้ดแบบเป็นขั้นตอน

การตรวจหา **SQL Injection (SQLi)** เป็นทักษะสำคัญของ Reviewer เพราะถ้าหลุดไปคือความเสียหายใหญ่หลวง นี่คือเทคนิคการตรวจแบบเป็นขั้นตอน (Step-by-Step) ที่คุณสามารถทำตามได้ทันที โดยเน้นบริบทของ **Node.js / NestJS / TypeORM** ที่คุณใช้งานอยู่ครับ

***

### ขั้นตอนที่ 1: ค้นหาจุดที่มีการ "ประกอบร่าง" คำสั่ง SQL

Reviewer ต้องกวาดสายตาหา "จุดเสี่ยง" ที่มีการสร้าง SQL String ขึ้นมาเอง โดยค้นหา Keyword เหล่านี้ในโค้ด:

* **Keywords:** `query(`, `execute(`, `raw(`, `createQueryBuilder`
* **Symbols:** เครื่องหมายบวก `+` หรือ Template Literal `${...}` ที่อยู่ใน String SQL


### ขั้นตอนที่ 2: ตรวจสอบที่มาของตัวแปร (Trace the Input)

เมื่อเจอจุดเสี่ยงแล้ว ให้ไล่ดูว่าตัวแปรที่เอามาใส่ มาจากไหน?

* **มาจาก User โดยตรง?** (`req.body`, `req.query`, `req.params`, `args`) -> 🚨 **อันตรายสูงสุด**
* **มาจาก Internal Logic?** (ค่าคงที่, ตัวแปรที่ระบบสร้างเอง) -> ✅ ปลอดภัย


### ขั้นตอนที่ 3: เช็ครูปแบบการเขียน (Syntax Analysis)

ดูว่า Developer เขียน SQL แบบไหน?

#### ❌ แบบที่ 1: String Concatenation (อันตรายมาก!)

เอา String มาต่อกันดื้อๆ ถ้าเจอแบบนี้ **Reject ทันที**

```typescript
// ❌ DANGER: User ส่งค่า "admin' OR '1'='1" มาก็พังหมด
const sql = "SELECT * FROM users WHERE name = '" + req.body.name + "'";
await connection.query(sql);
```


#### ❌ แบบที่ 2: Template Literal (อันตรายเหมือนกัน!)

ถึงจะดูทันสมัย แต่ถ้าใส่ตัวแปรตรงๆ ก็ไม่รอด

```typescript
// ❌ DANGER: เหมือนข้างบน แค่เขียนสวยกว่า
const sql = `SELECT * FROM users WHERE name = '${req.body.name}'`;
await connection.query(sql);
```


#### ✅ แบบที่ 3: Parameterized Query (ปลอดภัย)

ใช้ตัวแทน (`?` หรือ `$1` หรือ `:name`) แล้วส่งค่าแยกต่างหาก

```typescript
// ✅ SAFE: Database จะมอง input เป็น "ข้อมูล" ไม่ใช่ "คำสั่ง"
await connection.query('SELECT * FROM users WHERE name = $1', [req.body.name]);
```


***

### ขั้นตอนที่ 4: ตรวจสอบการใช้ ORM (TypeORM Check)

แม้จะใช้ ORM ก็อาจพลาดได้ ให้เช็คดังนี้:

#### ❌ จุดตายของ TypeORM: `user input` ในฟังก์ชันดิบ

ถ้า Developer เผลอส่ง Object ที่คุมไม่ได้เข้าไปในฟังก์ชันบางตัว

```typescript
// ❌ DANGER: TypeORM เวอร์ชันเก่า (ก่อน 0.3.0) มีช่องโหว่ถ้าส่ง req.body เข้าไปตรงๆ
// User อาจส่ง { "id": 1, "password": { "$gt": "" } } เพื่อ bypass ได้
userRepository.findOne(req.body); 
```

**วิธีแก้:** ต้องระบุ field ชัดเจน

```typescript
// ✅ SAFE
userRepository.findOne({ where: { id: req.body.id } });
```


#### ❌ การใช้ `createQueryBuilder` แบบผิดๆ

```typescript
// ❌ DANGER: ต่อ String ใน .where()
createQueryBuilder("user")
  .where("user.name = '" + name + "'") 

// ✅ SAFE: ใช้ Parameters Object
createQueryBuilder("user")
  .where("user.name = :name", { name: name }) 
```


***

### 📝 สรุป Checklist สำหรับ Reviewer (Copy ไปใช้ได้เลย)

1. [ ] **No String Concatenation:** ห้ามใช้ `+` หรือ `${}` ต่อ SQL String เด็ดขาด
2. [ ] **Use Parameters:** ทุกค่าที่มาจาก User ต้องส่งผ่าน Parameter Array (`$1`, `?`, `:name`) เท่านั้น
3. [ ] **Validate Input:** ตัวแปรที่รับมาเป็น `string` หรือ `number` จริงไหม? (ไม่ใช่ Object แฝงตัวมา)
4. [ ] **ORM Usage:** ไม่ส่ง `req.body` เข้าไปใน ORM Method ตรงๆ ต้องแกะค่าออกมาทีละตัว (Explicit Mapping)

### 🛠 เครื่องมือช่วย (Automated Tools)

ถ้าไม่อยากตรวจตาเปล่า ให้ใช้ Tools ช่วยสแกน:

* **ESLint Plugin:** `eslint-plugin-security` (ช่วยเตือนเมื่อมีการต่อ String ใน SQL)
* **SonarQube:** แจ้งเตือนช่องโหว่ "Security Hotspot" อัตโนมัติใน CI Pipeline
 

