# การเพิ่ม (Join) เครื่อง หรือ Node 

การเพิ่ม (Join) เครื่องหรือ Node ใหม่เข้าไปในระบบคลัสเตอร์ของ Docker นั้น จะเป็นกระบวนการของ **Docker Swarm** (ส่วน Docker Stack คือคำสั่งที่ใช้ในการ Deploy แอปพลิเคชันลงบน Swarm อีกที)

เพื่อให้เห็นภาพรวมของระบบ Manager และ Worker ใน Docker Swarm ได้ชัดเจนขึ้น สามารถดูแผนภาพด้านล่างนี้ประกอบได้

ขั้นตอนและตัวอย่างคำสั่งในการนำ Node ใหม่เข้าร่วม (Join) Docker Swarm 

### ขั้นตอนที่ 1: ขอ Token จาก Manager Node

คุณต้องดึง "Join Token" จากเครื่องที่ทำหน้าที่เป็น **Manager Node** ก่อน เพื่อให้เครื่องใหม่ใช้เป็นกุญแจในการเข้าร่วม โดยให้คุณล็อกอินเข้าเครื่อง Manager แล้วรันคำสั่งใดคำสั่งหนึ่ง ดังนี้:

**หากต้องการให้เครื่องใหม่เป็น Worker Node (ใช้รันแอปพลิเคชันอย่างเดียว):**

```bash
docker swarm join-token worker

```

**หากต้องการให้เครื่องใหม่เป็น Manager Node (ใช้บริหารจัดการคลัสเตอร์ด้วย):**

```bash
docker swarm join-token manager

```

เมื่อรันคำสั่ง ระบบจะแสดงผลลัพธ์ (Output) กลับมาเป็นคำสั่งที่พร้อมใช้งาน ตัวอย่างเช่น:

> `To add a worker to this swarm, run the following command:`
> `docker swarm join --token SWMTKN-1-49nj1cmql0jkz5s954yi3oex3nedyz0fb0xx14ie39trti4wxv-8vxv8rssqyt6cw0dbjxcq03x2 192.168.99.100:2377`

---

### ขั้นตอนที่ 2: รันคำสั่ง Join บนเครื่องใหม่ (Node ที่ต้องการเพิ่ม)

ให้คุณคัดลอกคำสั่งที่ได้จากขั้นตอนแรก ไปวางและรันใน Terminal ของ **เครื่องใหม่** ที่ต้องการนำมาเข้าร่วมคลัสเตอร์

```bash
docker swarm join --token SWMTKN-1-49nj1cmql0jkz5s954yi3oex3nedyz0fb0xx14ie39trti4wxv-8vxv8rssqyt6cw0dbjxcq03x2 192.168.99.100:2377

```

หากสำเร็จ ระบบจะแสดงข้อความว่า:

> `This node joined a swarm as a worker.`

*(หมายเหตุ: `192.168.99.100` คือ IP ของ Manager Node และ `2377` คือ Port มาตรฐานที่ Docker Swarm ใช้สื่อสารกัน)*

---

### ขั้นตอนที่ 3: ตรวจสอบความสำเร็จ (ทำบน Manager Node)

กลับมาที่เครื่อง Manager Node เพื่อตรวจสอบว่า Node ใหม่ได้ถูกเพิ่มเข้ามาในระบบเรียบร้อยแล้วหรือไม่ โดยรันคำสั่ง:

```bash
docker node ls

```
คุณจะเห็นตารางแสดงรายชื่อ Node ทั้งหมดในระบบ พร้อมสถานะ (Status) เป็น `Ready` และ Availability เป็น `Active` แสดงว่า Node ใหม่พร้อมรับคำสั่งจาก Docker Stack เพื่อ Deploy แอปพลิเคชัน

## Port Require

Node ใหม่จะสามารถเชื่อมต่อและทำงานร่วมกับ Manager Node ใน Docker Swarm ได้อย่างสมบูรณ์ คุณจำเป็นต้องเปิด Port ที่เกี่ยวข้องใน Firewall (หรือ Security Group หากใช้ระบบ Cloud) ระหว่างเครื่องในคลัสเตอร์ดังต่อไปนี้:

* **Port 2377 (TCP):** ใช้สำหรับจัดการคลัสเตอร์ (Cluster Management Communications)
* *หน้าที่:* เป็น Port หลักที่เครื่อง Manager ใช้สื่อสาร แจกจ่ายงาน และรับข้อมูลจากเครื่อง Worker


* **Port 7946 (TCP และ UDP):** ใช้สำหรับการสื่อสารระหว่าง Node (Node-to-node Communication)
* *หน้าที่:* ใช้เพื่อให้แต่ละเครื่องในคลัสเตอร์มองเห็นและค้นหากันเองได้ รวมถึงการเช็คสถานะ (Health check/Gossip) ของเครื่องอื่นๆ


* **Port 4789 (UDP):** ใช้สำหรับระบบเครือข่ายจำลอง (Overlay Network Traffic)
* *หน้าที่:* สำคัญมากเพื่อให้ Container ที่รันอยู่คนละเครื่อง (Node) สามารถส่งข้อมูลสื่อสารหากันได้ผ่านเครือข่ายภายในของ Docker (Ingress Network)



> **ข้อควรระวัง:** หากเปิดเพียง Port 2377 เครื่องใหม่อาจจะพิมพ์คำสั่ง Join ได้สำเร็จและมองเห็นในระบบ แต่เมื่อคุณสั่ง Deploy แอปพลิเคชันด้วย Docker Stack แล้ว Container ที่อยู่คนละเครื่องจะไม่สามารถสื่อสารกันได้เลยหากลืมเปิด Port 7946 และ 4789