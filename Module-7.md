#  **Docker Registry**

การใช้งาน Docker ร่วมกับ **Registry** (คลังเก็บ Image เช่น Docker Hub, AWS ECR หรือ GitLab Registry) เป็นหัวใจสำคัญในการนำแอปพลิเคชันไปใช้จริง (Deployment)

คำสั่งหลักๆ ที่คุณต้องใช้จะมีอยู่ 5 คำสั่ง: `login`, `pull`, `tag`, `push` และ `logout`

หากคุณต้องการดึง Image สาธารณะมาใช้ คุณสามารถใช้แค่คำสั่ง `docker pull <image_name>` ถ้าคุณต้องการ **อัปโหลด (Push) Image ของคุณเองขึ้นไปเก็บบน Registry** จะต้องใช้คำสั่งตามด้านล่าง

1. **เข้าสู่ระบบ (Login):**
ก่อนจะอัปโหลดได้ คุณต้องยืนยันตัวตนกับ Registry นั้นก่อน

```bash
# สำหรับ Docker Hub (ค่าเริ่มต้น)
docker login

# สำหรับ Private Registry อื่นๆ (ต้องระบุ URL)
docker login registry.example.com

```

ระบบจะถาม Username และ Password (หรือ Access Token)


2. **เตรียม Image และตั้งชื่อ (Tag):** ขั้นตอนนี้สำคัญมาก มักเป็นจุดที่มือใหม่พลาดบ่อยที่สุด.
Docker จะรู้ว่าต้องอัปโหลด Image ไปที่ไหน โดยดูจาก **ชื่อของ Image** ดังนั้นเราต้องใช้คำสั่ง `tag` เพื่อเปลี่ยนชื่อ Image ในเครื่องให้ตรงกับรูปแบบของ Registry ก่อน

รูปแบบคำสั่ง: `docker tag <ชื่อเดิม> <ชื่อใหม่>`

```bash
# สมมติคุณมี Image ในเครื่องชื่อ my-web:latest
# คุณต้องการอัปโหลดไปที่ Docker Hub บัญชีชื่อ "kitty"

docker tag my-web:latest kitty/my-web:v1.0

```

*ถ้าเป็น Private Registry ชื่อจะต้องนำหน้าด้วย URL เช่น `[registry.example.com/kitty/my-web:v1.0](https://registry.example.com/kitty/my-web:v1.0)*`


3. **อัปโหลดขึ้น Registry (Push):**
เมื่อตั้งชื่อ (Tag) ถูกต้องแล้ว ให้ใช้คำสั่ง `push` พร้อมระบุชื่อที่เราเพิ่ง Tag ไป

```bash
docker push kitty/my-web:v1.0

```

Docker จะทำการอัปโหลด Layer ต่างๆ ของ Image ขึ้นไปยัง Registry ของคุณ


4. **ออกจากระบบ (Logout):** เพื่อความปลอดภัยเมื่อใช้งานบนเครื่องส่วนรวม.
เมื่อจัดการทุกอย่างเสร็จแล้ว ควร Logout ออกเสมอเพื่อป้องกันไม่ให้คนอื่นมาใช้งานสิทธิ์ของเรา

```bash
docker logout
# หรือ docker logout registry.example.com

```


---

### สรุปคำสั่งที่ใช้งานบ่อย (Cheatsheet)

* **ดึง Image มาใช้:** `docker pull ubuntu:20.04`
* **ดู Image ทั้งหมดในเครื่อง:** `docker images`
* **เปลี่ยนชื่อ/ผูก Tag:** `docker tag <source> <target>`
* **อัปโหลด Image:** `docker push <target>`