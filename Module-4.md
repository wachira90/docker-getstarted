#  **Docker Volumes**

**Docker Volumes** ซึ่งเป็นเรื่องสำคัญมากในการจัดการข้อมูล (Data Management) สำหรับคอนเทนเนอร์ นี่คือบทสรุปเนื้อหาและคำสั่งที่เกี่ยวข้องเพื่อใช้เป็นบทเรียนครับ:

### 1. ปัญหา: ข้อมูลชั่วคราว (Ephemeral Data)

โดยปกติแล้ว เมื่อเราสร้างคอนเทนเนอร์ขึ้นมา ข้อมูลทั้งหมดที่ถูกสร้างหรือแก้ไข (เช่น อัปโหลดรูปภาพ, เพิ่มข้อมูลลงฐานข้อมูล) จะถูกเก็บไว้ในชั้นของคอนเทนเนอร์ (Container Layer)

* **ปัญหาคือ:** เมื่อเราลบ (Remove) คอนเทนเนอร์นั้นทิ้งไป ข้อมูลทั้งหมดที่อยู่ข้างในจะหายไปอย่างถาวร

---

### 2. ทางแก้: ข้อมูลถาวร (Persistent Data) ด้วย Docker Volumes

เพื่อแก้ปัญหาข้อมูลสูญหาย Docker จึงมีสิ่งที่เรียกว่า **Volumes** ซึ่งทำหน้าที่เสมือน "ฮาร์ดดิสก์จำลอง" หรือพื้นที่จัดเก็บข้อมูลที่อยู่ **นอกเหนือ** จากวงจรชีวิตของคอนเทนเนอร์

* ข้อมูลจะถูกเก็บไว้บนเครื่องโฮสต์ (Host Machine) และให้ Docker เป็นคนบริหารจัดการให้โดยตรง
* ถึงแม้คอนเทนเนอร์จะถูก Stop หรือ Remove ไป ข้อมูลใน Volume ก็จะยังคงอยู่ปลอดภัย และสามารถนำคอนเทนเนอร์ตัวใหม่มาเชื่อมต่อเพื่อใช้งานข้อมูลเดิมต่อได้ทันที

---

### 3. คำสั่งพื้นฐานในการจัดการ Docker Volumes

* **สร้าง Volume ใหม่**
```bash
docker volume create [VOLUME_NAME]

```


* *ตัวอย่าง:* `docker volume create my-data`


* **ดูรายชื่อ Volume ทั้งหมดที่มีในเครื่อง**
```bash
docker volume ls

```


* **ดูรายละเอียดของ Volume (เช่น ถูกเก็บไว้ที่พาร์ทไหนของโฮสต์)**
```bash
docker volume inspect [VOLUME_NAME]

```


* **ลบ Volume (ลบได้เฉพาะเมื่อไม่มีคอนเทนเนอร์ไหนใช้งานอยู่)**
```bash
docker volume rm [VOLUME_NAME]

```



---

### 4. วิธีการนำ Volume ไปใช้งาน (Mount)

เมื่อเรามี Volume แล้ว เราสามารถนำไปเชื่อมต่อ (Mount) กับคอนเทนเนอร์ได้ตอนที่ใช้คำสั่ง `docker run` โดยใช้ flag `-v`

```bash
docker run -v [VOLUME_NAME]:[PATH_IN_CONTAINER] [IMAGE_NAME]

```

**ตัวอย่างการใช้งานจริง:**
สมมติว่าคุณต้องการรันฐานข้อมูล MySQL และไม่ต้องการให้ข้อมูลหายเมื่อคอนเทนเนอร์ถูกลบ

```bash
docker run -d --name my-db -v db-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=secret mysql

```

* **คำอธิบาย:** คำสั่งนี้จะสร้าง Volume ชื่อ `db-data` (ถ้ายังไม่มี) และนำไปเชื่อมกับแฟ้ม `/var/lib/mysql` (ซึ่งเป็นที่เก็บฐานข้อมูลของ MySQL) ภายในคอนเทนเนอร์ ข้อมูลทั้งหมดจึงถูกจัดเก็บอย่างปลอดภัยแบบ Persistent Data ครับ



# **Bind Mounts** 

**Bind Mounts** ซึ่งเป็นอีกหนึ่งรูปแบบของการจัดการข้อมูล (Data Management) ที่คล้ายกับ Volumes ครับ (ต่างกันตรงที่ Bind Mounts จะเป็นการเจาะจงเชื่อมโยงโฟลเดอร์จริงบนเครื่อง Host ของเราเข้ากับ Container โดยตรง)

### Bind Mounts คืออะไร?

Bind Mounts คือการเชื่อมต่อ (Mount) โฟลเดอร์หรือไฟล์ที่มีตำแหน่งอยู่แล้วบนเครื่องของนักพัฒนา (Host Machine) เข้าไปไว้ในตำแหน่งที่กำหนดภายใน Container

* ข้อมูลระหว่างสองฝั่งจะถูกซิงค์ตรงกันแบบ Real-time
* หากเราแก้ไขไฟล์โค้ดบนเครื่องของเรา การเปลี่ยนแปลงนั้นจะส่งผลเข้าไปใน Container ทันทีโดยไม่ต้องสั่ง Build Image ใหม่

---

### 1. รูปแบบคำสั่ง (Command)

การใช้งาน Bind Mounts จะใช้ flag `-v` ตอนรันคำสั่ง `docker run` คล้ายกับ Volumes แต่จะระบุ **Path เต็ม (Absolute Path)** ของเครื่อง Host แทนการตั้งชื่อ Volume ลอยๆ

```bash
docker run -v [HOST_PATH]:[CONTAINER_PATH] [IMAGE_NAME]

# windows
docker run -v .\data\:/usr/share/html/index.html -p 8011:80 nginx:1.29.5-alpine3.23-slim

# linux
docker run -v ./data/:/usr/share/html/index.html -p 8011:80 nginx:1.29.5-alpine3.23-slim

```

---

### 2. Workshop 1: การแก้ไข Code แบบ Real-time

สมมติว่าคุณกำลังพัฒนาเว็บไซต์ และต้องการให้แอปพลิเคชันที่รันอยู่ใน Container อัปเดตทันทีเมื่อคุณพิมพ์แก้โค้ดและกดเซฟไฟล์บนเครื่องตัวเอง

```bash
docker run -d --name my-web-app -v $(pwd)/src:/app/src -p 8080:80 my-web-image

```

* **คำอธิบาย:** คำสั่งนี้จะนำโฟลเดอร์ `src` ในตำแหน่งปัจจุบันบนเครื่องคุณ (`$(pwd)/src`) ไปวางทับโฟลเดอร์ `/app/src` ใน Container ทำให้คุณสามารถใช้ Text Editor ในเครื่องโฮสต์เขียนโค้ดได้อย่างสะดวกสบาย

---

### 3. Workshop 2: การรันระบบฐานข้อมูลอย่างปลอดภัย

นอกจากการเขียนโค้ดแล้ว Bind Mounts ยังนิยมใช้ร่วมกับระบบ Database เพื่อโยนไฟล์ตั้งค่า (Configuration) หรือไฟล์สคริปต์ตั้งต้น (Init scripts) เข้าไปตอนเริ่มรันระบบ

```bash
docker run -d --name secure-db -v /path/to/my/custom.cnf:/etc/mysql/conf.d/custom.cnf -e MYSQL_ROOT_PASSWORD=supersecret mysql

```

* **คำอธิบาย:** คำสั่งนี้เป็นการนำไฟล์ตั้งค่า `custom.cnf` จากเครื่องโฮสต์ เข้าไปเสียบแทนที่ไฟล์ตั้งค่าของ MySQL ภายใน Container เพื่อปรับแต่งความปลอดภัยหรือประสิทธิภาพของฐานข้อมูลตามที่นักพัฒนาต้องการครับ