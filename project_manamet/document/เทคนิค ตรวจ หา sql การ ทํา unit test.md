
# เทคนิค ตรวจ หา sql การ ทํา unit test

สำหรับการ "ตรวจหา" หรือตรวจสอบ SQL ในการทำ Unit Test นั้นขึ้นอยู่กับว่าคุณต้องการตรวจสอบ **"Logic ของคำสั่ง SQL"** หรือ **"Logic ของโค้ด JavaScript/TypeScript"** ที่เรียกใช้ Database ครับ

เนื่องจากคุณใช้ **Next.js + Prisma + PostgreSQL** คำตอบจะแบ่งออกเป็น 3 ระดับตามความลึกของการทดสอบครับ:

### 1. ตรวจสอบ SQL ในระดับ Integration Test (แนะนำสูงสุด)

ในระดับนี้เท่านั้นที่ Prisma จะทำการ "สร้าง SQL จริงๆ" ออกมา หากคุณทำ Unit Test แบบ Mock (จำลอง) จะไม่มี SQL ถูกสร้างขึ้นครับ

**เทคนิค: ใช้ Prisma Event Logger เพื่อดู SQL**
วิธีนี้จะช่วยให้คุณ "เห็น" SQL ที่ถูกยิงออกไปจริงๆ และตรวจสอบ Performance ได้ (เช่น N+1 Problem)

```typescript
// ตั้งค่า PrismaClient ให้ส่ง Event เมื่อมี Query
const prisma = new PrismaClient({
  log: [
    {
      emit: 'event',
      level: 'query', // จับเฉพาะระดับ Query
    },
  ],
})

// เขียนฟังก์ชันดักจับ (Listener)
prisma.$on('query', (e) => {
  console.log('Query: ' + e.query)
  console.log('Params: ' + e.params)
  console.log('Duration: ' + e.duration + 'ms')
})

// ใน Test case ของคุณ
test('should find user', async () => {
  await prisma.user.findMany() // SQL จะถูกปรินต์ออกมาที่ Console
})
```


***

### 2. ตรวจสอบ Logic ในระดับ Unit Test (ใช้ Mock)

ในระดับนี้ เราจะ **ไม่เห็น SQL string** แต่เราจะตรวจสอบว่า **"Prisma ถูกสั่งงานด้วย parameters ที่ถูกต้องหรือไม่"** ซึ่งเป็นการตรวจ Logic แทนการตรวจ SQL ครับ

**เทคนิค: ใช้ `jest-mock-extended` หรือ `prisma-mock`**
ตรวจสอบว่าโค้ดของเราเรียกฟังก์ชัน `findMany`, `create` ด้วย `where` clause ที่ถูกต้องหรือไม่

```typescript
import { prisma } from './client' // client ที่ถูก mock แล้ว
import { mockDeep } from 'jest-mock-extended'

test('should call findMany with correct email', async () => {
  const email = 'test@example.com'
  
  // เรียกฟังก์ชันของคุณ
  await findUserByEmail(email)

  // ตรวจสอบว่า Prisma ถูกเรียกด้วย parameter ที่คาดหวังหรือไม่
  expect(prisma.user.findUnique).toHaveBeenCalledWith({
    where: { email: 'test@example.com' },
  })
})
```


***

### 3. ตรวจสอบความถูกต้องของ SQL จริง (Testcontainers)

หากคุณต้องการมั่นใจ 100% ว่า SQL นั้นทำงานได้จริง (เช่นพวก Raw Query ซับซ้อน หรือ Transaction) การใช้ Mock จะไม่เพียงพอ ต้องใช้ **Testcontainers** เพื่อรัน Database จริงใน Docker ขึ้นมาชั่วคราวขณะเทส

**เทคนิค: Testcontainers**
วิธีนี้จะรัน PostgreSQL จริงๆ ใน Docker ขึ้นมาให้เทส แล้วทำลายทิ้งเมื่อจบเทส ทำให้มั่นใจได้ว่า SQL ถูกต้องตาม Syntax ของ Postgres แน่นอน

* **ข้อดี:** เจอ Error จริงที่ Mock ไม่เจอ (เช่น constraint violation)
* **ข้อเสีย:** ช้ากว่า Unit Test ปกติ


### สรุป

* ถ้าอยาก **"เห็น"** SQL Text: ต้องใช้ **Integration Test + Prisma Log**
* ถ้าอยาก **"ตรวจ"** Logic การ query: ใช้ **Unit Test + Mock (`toHaveBeenCalledWith`)**
* ถ้าใช้ **Raw SQL (`$queryRaw`)**: ควรใช้ Integration Test กับ Database จริง (Testcontainers) เท่านั้น เพราะ Mock จะไม่สามารถบอกได้ว่า SQL string นั้นถูก Syntax หรือไม่
 

[^39]: https://www.prisma.io/docs/orm/prisma-client/client-extensions/middleware/session-data-middleware

[^40]: https://github.com/prisma/prisma/issues/16451

