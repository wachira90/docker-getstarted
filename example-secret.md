# Docker Secrets 

ปกติแล้ว Image ของ PostgreSQL จะรองรับการอ่านค่าจากไฟล์ Secret โดยการเติมคำว่า `_FILE` ต่อท้ายตัวแปร Environment เดิม เช่น จาก `POSTGRES_PASSWORD` เปลี่ยนเป็น `POSTGRES_PASSWORD_FILE`

### 1. สร้างไฟล์ Secret สำหรับเก็บรหัสผ่าน

สร้างไฟล์ Text ธรรมดาเพื่อเก็บรหัสผ่าน (ไฟล์นี้ไม่ควรนำขึ้น Git repository)

เปิด Terminal แล้วรันคำสั่งนี้เพื่อสร้างไฟล์ชื่อ `db_password.txt`:

```bash
echo "MySuperSecretPassword123!" > db_password.txt

```

---

### 2. สร้างไฟล์ `docker-compose.yml`

สร้างไฟล์ `docker-compose.yml` ในโฟลเดอร์เดียวกัน และใส่โค้ดด้านล่างนี้ลงไป:

```yaml
services:
  postgres-db:
    image: postgres:15
    restart: always
    environment:
      # ตั้งค่า User และ Database (ส่วนนี้จะใส่ตรงๆ ก็ได้หากไม่ใช่ความลับ)
      POSTGRES_USER: my_app_user
      POSTGRES_DB: my_app_db
      
      # สำคัญ: ใช้ _FILE ต่อท้าย และชี้ไปที่ Path ของ Secret ใน Container
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
    
    # ประกาศให้ Service นี้ใช้งาน Secret ที่ชื่อ db_password
    secrets:
      - db_password
      
    ports:
      - "5432:5432"
    volumes:
      - pg_data:/var/lib/postgresql/data

# ประกาศ Secret ในระดับ Global ของ Compose
secrets:
  db_password:
    file: ./db_password.txt # ชี้ไปยังไฟล์ที่เราสร้างไว้ในข้อ 1

volumes:
  pg_data:

```

---

### 3. คำอธิบายการทำงาน

* **`POSTGRES_PASSWORD_FILE: /run/secrets/db_password`**: เราบอกให้ PostgreSQL ไปอ่านรหัสผ่านจากไฟล์ที่ระบุ (โดยปกติ Docker จะ Mount ไฟล์ Secrets ไว้ที่ `/run/secrets/` เสมอ)
* **`secrets: - db_password` (ในส่วนของ Service)**: เป็นการบอกว่า Service `postgres-db` มีสิทธิ์เข้าถึง Secret ตัวนี้
* **`secrets:` (ด้านล่างสุด)**: เป็นการกำหนดว่า Secret ที่ชื่อ `db_password` นั้น ให้อ่านข้อมูลมาจากไฟล์ `./db_password.txt` ที่อยู่ในเครื่องของเรา

---

### 4. วิธีการรันและทดสอบ

เมื่อเตรียมไฟล์เสร็จแล้ว คุณสามารถสั่งรันฐานข้อมูลได้ตามปกติด้วยคำสั่ง:

```bash
docker-compose up -d

```

หากต้องการทดสอบว่ารหัสผ่านถูกตั้งค่าอย่างถูกต้องหรือไม่ สามารถลองเชื่อมต่อเข้าฐานข้อมูล (หากคุณมี `psql` ติดตั้งอยู่ในเครื่อง) ด้วยคำสั่ง:

```bash
psql -h localhost -U my_app_user -d my_app_db

```

*(เมื่อระบบถามรหัสผ่าน ให้ใส่ `MySuperSecretPassword123!`)*

> **ข้อควรระวัง:**
> * อย่าลืมเพิ่ม `db_password.txt` ลงในไฟล์ `.gitignore` ของคุณ เพื่อป้องกันไม่ให้รหัสผ่านหลุดขึ้นไปยัง GitHub หรือระบบ Version Control อื่นๆ
> * หากต้องการซ่อน User หรือชื่อ Database ด้วย ก็สามารถสร้างไฟล์ Secret เพิ่มและใช้ตัวแปร `POSTGRES_USER_FILE` และ `POSTGRES_DB_FILE` ด้วยหลักการเดียวกันได้เลย
> 
>
