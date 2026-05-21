# **Building Images**

**Building Images** จะเกี่ยวข้องกับการสร้าง Docker Image ของคุณเองผ่านไฟล์ที่เรียกว่า `Dockerfile` และนี่คือคำสั่งในการ Build images

### 1. คำสั่งหลักในการสร้าง Image (CLI)

เมื่อคุณเขียน `Dockerfile` เสร็จแล้ว คำสั่งที่ใช้ใน Command Line เพื่อสร้าง Image คือ `docker build`

```bash
docker build -t [IMAGE_NAME]:[TAG] [PATH]

docker build -t my-web:v1 .

docker build -t my-web:v1 . --no-cache

docker buildx build -t my-web:dev1.0.0 . --no-cache

```

* **ตัวอย่าง:** `docker build -t my-app:1.0 .` (สร้าง Image ชื่อ my-app เวอร์ชัน 1.0 โดยใช้ Dockerfile ที่อยู่ในโฟลเดอร์ปัจจุบัน `.`)

---

### 2. โครงสร้างคำสั่งพื้นฐานใน Dockerfile

คำสั่งเหล่านี้จะถูกเขียนไว้ภายในไฟล์ `Dockerfile` เพื่อกำหนดขั้นตอนการสร้าง Image:

* **`FROM`**
ใช้กำหนด Base Image ที่จะนำมาเป็นพื้นฐาน (ต้องอยู่บรรทัดแรกเสมอ)
* *ตัวอย่าง:* `FROM node:18-alpine`


* **`RUN`**
ใช้รันคำสั่งต่างๆ (เช่น ติดตั้งโปรแกรมหรือไลบรารี) ในขณะที่กำลัง **สร้าง (Build)** Image
* *ตัวอย่าง:* `RUN npm install` หรือ `RUN apt-get update`


* **`COPY`**
ใช้คัดลอกไฟล์หรือโฟลเดอร์จากเครื่องโฮสต์ (เครื่องของคุณ) เข้าไปไว้ใน Image
* *ตัวอย่าง:* `COPY . /app` (คัดลอกไฟล์ทั้งหมดไปยังโฟลเดอร์ /app ใน Image)


* **`EXPOSE`**
ใช้ระบุว่าคอนเทนเนอร์นี้จะเปิดรับการเชื่อมต่อที่พอร์ตไหน (เป็นเพียงการบอกเอกสาร (Document) ไม่ได้เปิดพอร์ตจริงทันที)
* *ตัวอย่าง:* `EXPOSE 8080`



---

### 3. ความแตกต่างระหว่าง CMD และ ENTRYPOINT

ทั้งสองคำสั่งใช้กำหนดสิ่งที่จะทำงานตอนที่ **รัน (Start)** คอนเทนเนอร์ แต่มีพฤติกรรมต่างกัน:

* **`CMD`** (Command)
เป็นการกำหนดคำสั่งหรือพารามิเตอร์ "เริ่มต้น" (Default) ซึ่ง **สามารถถูกเขียนทับ (Override) ได้ง่าย** ผ่าน Command Line ตอนรัน `docker run`
* *ตัวอย่าง:* `CMD ["npm", "start"]`


* **`ENTRYPOINT`**
เป็นการกำหนดคำสั่งหลักที่จะต้องทำงานเสมอ **มักจะไม่ถูกเขียนทับ** (ยกเว้นจะจงใจใช้ flag `--entrypoint`) นิยมใช้ร่วมกับ `CMD` เพื่อให้ `CMD` ส่งค่าพารามิเตอร์มาให้
* *ตัวอย่าง:* `ENTRYPOINT ["nginx", "-g", "daemon off;"]`



---

### 4. Multi-stage builds (การลดขนาด Image)

เป็นเทคนิคขั้นสูงใน `Dockerfile` ที่อนุญาตให้เราใช้คำสั่ง `FROM` ได้หลายครั้งในไฟล์เดียว เพื่อแยกขั้นตอนการ "Build" โค้ด ออกจากขั้นตอนการ "Run" ทำให้ Image ตัวสุดท้าย (Production Image) มีขนาดเล็กและปลอดภัยขึ้น เพราะไม่มีเครื่องมือสำหรับการ Build หลงเหลืออยู่

**ตัวอย่างโครงสร้าง Multi-stage builds:**

```dockerfile
FROM public.ecr.aws/docker/library/golang:1.22.1-alpine3.19 AS build

WORKDIR /app

COPY . .

RUN go mod init app

RUN go build -o main .

FROM public.ecr.aws/docker/library/alpine:3.19.1

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]

```