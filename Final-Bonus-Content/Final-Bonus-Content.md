# Final Bonus Content

- Command For Develop

- Secret

- Log Rotation


## Command For Develop


### FOR BUILD WITH INTERNET

docker run --rm -it -w /app --entrypoint bash public.ecr.aws/docker/library/python:3.12-slim-bullseye

docker run --rm -it -w /app --entrypoint bash public.ecr.aws/docker/library/python:3.12-slim-bullseye

docker run --rm -it -w /app -p 8022:8080 -v.\pydata\:/app --entrypoint bash public.ecr.aws/docker/library/python:3.12-slim-bullseye

pip install bottle

python main.py

### FOR RUN COMMAND

docker run --rm -it -p 8011:8080 --entrypoint bash public.ecr.aws/docker/library/python:3.12-slim-bullseye


### START WITH COMMAND

docker run -p 8099:8080 -w /var/log/ docker.io/library/python:3.12-slim python -m http.server 8080



docker exec mysql /bin/bash -c "/var/lib/mysql/backup/db-backup.sh"



## Secret

ตัวอย่าง:

หากต้องการสร้างทั้ง `dbuser` และ `dbpass` วิธีที่ถูกต้องและเป็นมาตรฐานที่สุดสำหรับ Docker Secrets คือการ **แยกไฟล์** สำหรับแต่ละค่าครับ

เนื่องจาก Docker Secrets จะทำงานโดยการนำไฟล์ที่เรากำหนดไปวางไว้ใน Container ที่ตำแหน่ง `/run/secrets/<ชื่อ_secret>` การแยกไฟล์จะทำให้แอปพลิเคชันดึงค่าไปใช้ได้ง่ายและปลอดภัย

### 1. คำสั่งสร้างไฟล์ Secret

แนะนำให้เพิ่ม flag `-n` ในคำสั่ง `echo` เพื่อป้องกันไม่ให้มีการขึ้นบรรทัดใหม่ (Newline) ติดเข้าไปในรหัสผ่าน ซึ่งมักจะทำให้แอปพลิเคชันเชื่อมต่อฐานข้อมูลไม่ผ่านครับ

```bash
# สร้างไฟล์สำหรับ Username
echo "sam01" > db_user.txt

# สร้างไฟล์สำหรับ Password
echo "example1234" > db_password.txt

```

### 2. ปรับแก้ไฟล์ `docker-compose.yml`

เพิ่ม Secret ตัวใหม่เข้าไปในทั้งหมวด `services` และหมวด `secrets` ด้านล่างสุดครับ

```yaml
services:

  web:
    image: docker.io/library/myapp:v2
    ports:
      - 8011:8080
    secrets:
      - db_user
      - db_password

secrets:
  db_user:
    file: ./db_user.txt
  db_password:
    file: ./db_password.txt

```

---

**ข้อควรรู้เพิ่มเติมสำหรับการนำไปใช้:**
เมื่อ Container รันขึ้นมา ตัวแอปพลิเคชันของคุณจะมองไม่เห็นค่าเหล่านี้เป็น Environment Variables โดยตรง (เช่น ไม่สามารถใช้ `getenv('dbuser')` ได้) แต่จะต้องเขียนโค้ดให้อ่านค่าที่เป็นข้อความออกมาจากไฟล์ตาม Path ด้านล่างนี้แทนครับ:

* **User:** `/run/secrets/db_user`
* **Password:** `/run/secrets/db_password`