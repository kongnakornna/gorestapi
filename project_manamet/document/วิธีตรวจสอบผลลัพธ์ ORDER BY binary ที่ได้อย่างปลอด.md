 
# วิธีตรวจสอบผลลัพธ์ ORDER BY binary ที่ได้อย่างปลอดภัย

**Understanding the User's Goal**

The user is asking for a "safe way to verify the result of an ORDER BY binary search" in the context of SQL Injection testing.

They likely mean: "After I use the binary search technique to guess the column count (e.g., I think it's 66), how do I definitively confirm this number is correct without breaking anything or causing false positives?" or perhaps "How do I do this testing safely (without damaging the database)?"

Given the context of previous questions (SQLi payloads, Union based), "safely" probably refers to **accuracy/verification** (making sure the number is right before moving to the `UNION SELECT` step) rather than just "not breaking the DB" (though that's part of it).

**Core Concept:**
Binary search gives you a *boundary* (e.g., 66 works, 67 fails). To "verify safely," you need to prove that exactly $N$ works and $N+1$ fails, and ideally, confirm it with a `UNION SELECT`.

**Plan:**

1. **Boundary Check (The "N vs N+1" Rule):** Explain that the final confirmation is testing `ORDER BY N` (Pass) vs `ORDER BY N+1` (Fail).
2. **The "NULL" Proof (The safest verification):** Before injecting real data, use `UNION SELECT NULL, NULL, ...` (N times) to confirm the page loads correctly without errors. This is the ultimate proof.
3. **Safety Precautions (Production vs Test):** Remind the user again that this is "safe" for the *tester's confidence* but still risky on *Production*.

**Structuring the Response:**

1. **Verification Step 1: The Boundary Test:** Explicitly check the winning number vs the next number.
2. **Verification Step 2: The UNION SELECT NULL Test:** The definitive proof.
3. **Why this is "Safe":** Explain that using `NULL` avoids data type mismatch errors (casting errors), which is a common pitfall if you try to inject strings into integer columns.

**Refining the content:**

- *Why verifying with UNION SELECT is better:* `ORDER BY` might pass even if the column count is wrong in some weird edge cases (like grouping), but `UNION` is strict.
- *The "Safe" aspect:* Emphasize `NULL` again. Using `'a'` or `1` can cause casting errors if the column types don't match, leading to false negatives. `NULL` fits everything.

Let's write the response.

การตรวจสอบผลลัพธ์จากการทำ Binary Search (เช่น มั่นใจว่าคือ 66 คอลัมน์แน่ๆ) ให้ **"ปลอดภัย" (Safe \& Accurate)** เพื่อไม่ให้หน้าแตกเวลาทำขั้นตอนต่อไป (`UNION SELECT`) มีขั้นตอนการยืนยันผลลัพธ์ที่แนะนำดังนี้ครับ:

### 1. ยืนยันด้วยกฎ "N vs N+1" (Boundary Check)

เมื่อ Binary Search บอกคุณว่าเลขจบที่ **N** (เช่น 66) อย่าเพิ่งเชื่อทันที ให้ทดสอบด้วยมืออีกครั้งแบบชัดๆ:

* **ยิง Payload:** `' ORDER BY 66 --`
    * **ผล:** ✅ **ต้องผ่าน** (หน้าเว็บโหลดปกติ ไม่ Error)
* **ยิง Payload:** `' ORDER BY 67 --`
    * **ผล:** ❌ **ต้องไม่ผ่าน** (Error หรือข้อมูลหาย)

ถ้าผลลัพธ์เป็นไปตามนี้ (ผ่านที่ 66, พังที่ 67) ความมั่นใจอยู่ที่ **90%** ครับ

***

### 2. ยืนยันด้วย "UNION SELECT NULL" (The Ultimate Proof) 🛡️

นี่คือวิธีที่ **"ปลอดภัยที่สุด"** ในการยืนยัน เพราะ `UNION` เคร่งครัดกว่า `ORDER BY` มาก และการใช้ `NULL` จะช่วยป้องกัน Error เรื่องชนิดข้อมูล (Data Type Mismatch)

**ทำไมต้อง NULL?**
ถ้าคุณทดสอบด้วยเลข `1` หรือตัวอักษร `'a'` อาจจะ Error ได้ถ้าคอลัมน์นั้นเป็นชนิดข้อมูลอื่น (เช่น ใส่ String ลงช่อง Integer) แต่ `NULL` เข้าได้กับทุกช่อง (Int, String, Date, Blob)

**วิธีการ:**
สร้าง Payload ที่มี `NULL` จำนวนเท่ากับ N (เช่น 66 ตัว)

```sql
' UNION SELECT NULL, NULL, NULL, ... (ครบ 66 ตัว) ... , NULL --
```

* **ถ้าหน้าเว็บโหลดปกติ:** ✅ **100% Confirmed!** (จำนวนคอลัมน์ถูกต้องแน่นอน)
* **ถ้าหน้าเว็บ Error:** แสดงว่าจำนวนคอลัมน์ผิด (อาจจะนับผิด หรือ WAF บล็อก) ต้องกลับไปเช็ค Binary Search ใหม่

***

### 3. ตรวจสอบจุดแสดงผล (Reflection Check)

หลังจาก `UNION SELECT NULL` ผ่านแล้ว ให้ค่อยๆ เปลี่ยน `NULL` ทีละตัวเป็น "ตัวเลขที่ไม่ซ้ำกัน" เพื่อดูว่าช่องไหนสะท้อนออกมาบนหน้าจออย่างปลอดภัย

**ตัวอย่าง (สมมติมี 3 คอลัมน์):**

```sql
' UNION SELECT 1111, 2222, 3333 --
```

* ถ้าหน้าเว็บโชว์เลข **2222** แทนชื่อสินค้า -> แปลว่าช่องที่ 2 คือช่องที่เราจะใช้ขโมยข้อมูล
* *ข้อควรระวัง:* อย่าเพิ่งใส่ Function (เช่น `user()`) ในขั้นตอนนี้ ให้ใช้เลขธรรมดาก่อนเพื่อให้แน่ใจว่า Injection Point ทำงานได้จริง

***

### 📝 สรุป Checklist การตรวจสอบความปลอดภัย

1. [ ] **Boundary Test:** N ผ่าน, N+1 พัง
2. [ ] **NULL Injection:** `UNION SELECT` ด้วย NULL จำนวน N ตัว ต้องไม่ Error
3. [ ] **No Data Type Error:** การใช้ `NULL` ช่วยเลี่ยง Error เรื่องชนิดตัวแปร ทำให้เราไม่หลงทางว่า "จำนวนคอลัมน์ผิด" ทั้งที่จริงๆ แค่ "ชนิดตัวแปรผิด"
