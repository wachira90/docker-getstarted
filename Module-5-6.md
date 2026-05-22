# Docker networking

การเขียน docker-compose.yml สำหรับ Network แต่ละแบบจะมีความแตกต่างกัน บางแบบใช้การประกาศในส่วน networks: และบางแบบใช้คำสั่ง network_mode: โดยตรง

ตัวอย่างการตั้งค่าของทั้ง 5 แบบ:

### 1. Bridge Network (แบบ Custom)

โดยปกติถ้าเราไม่ระบุอะไรเลย Compose จะสร้าง Default Bridge ให้อยู่แล้ว แต่แนวทางปฏิบัติที่ดี (Best Practice) คือการสร้าง Custom Bridge เพื่อให้จัดการชื่อและแยกส่วนแอปพลิเคชันได้ชัดเจน

```yaml
services:
  web:
    image: nginx:alpine
    ports:
      - "8080:80"
    networks:
      - my_app_net

  db:
    image: redis:alpine
    networks:
      - my_app_net

networks:
  my_app_net:
    driver: bridge

```

> **จุดสังเกต:** web และ db สามารถคุยกันเองผ่านชื่อ Service (เช่น ping db) ได้เลย เพราะอยู่ในวง my_app_net เดียวกัน

### 2. Host Network

การใช้ Host จะไม่ประกาศใน block networks ด้านล่าง แต่จะใช้คำสั่ง network_mode: "host" ที่ตัว Service เลย และ**ไม่ต้องใช้คำสั่ง ports** เพราะ Container จะใช้ Port ของเครื่อง Host โดยตรง

```yaml
services:
  web:
    image: nginx:alpine
    network_mode: "host"

```

> **จุดสังเกต:** Nginx ตัวนี้จะรันบน Port 80 ของเครื่อง Host โดยตรงทันที (ถ้าเครื่อง Host มี Service อื่นใช้ Port 80 อยู่ก่อนแล้ว Container นี้จะ Error รันไม่ขึ้น)

### 3. None Network (ปิดการเชื่อมต่อ)

ใช้ network_mode: "none" คล้ายกับแบบ Host

```yaml
services:
  secure_worker:
    image: alpine
    command: ping localhost
    network_mode: "none"

```

> **จุดสังเกต:** Container นี้จะไม่สามารถออกอินเทอร์เน็ตได้ และ Container อื่นก็ไม่สามารถเข้าถึงมันได้ผ่าน Network

### 4. Macvlan Network

การทำ Macvlan จะต้องระบุรายละเอียดให้ตรงกับ Physical Network ของเครื่อง Host ของจริง เช่น Subnet, Gateway และชื่อ Network Interface (parent) ของเครื่อง Host

```yaml
services:
  web:
    image: nginx:alpine
    networks:
      my_macvlan_net:
        ipv4_address: 192.168.1.150 # ฟิกซ์ IP ให้ Container (ต้องอยู่ในวง Subnet)

networks:
  my_macvlan_net:
    driver: macvlan
    driver_opts:
      parent: eth0 # ชื่อ LAN Interface ของเครื่อง Host (เช่น eth0, eno1)
    ipam:
      config:
        - subnet: "192.168.1.0/24" # วง LAN ของออฟฟิศ/บ้าน
          gateway: "192.168.1.1" # Gateway ของเร้าเตอร์

```

### 5. Overlay Network

การใช้ Overlay มักจะใช้ร่วมกับ **Docker Swarm** (การใช้คำสั่ง docker-compose up ธรรมดาอาจจะ Error หรือแจ้งเตือนได้ ต้องใช้คำสั่ง docker stack deploy แทน)

```yaml
services:
  web:
    image: nginx:alpine
    networks:
      - my_overlay_net
    deploy:
      replicas: 3 # รัน 3 ตัว กระจายไปตามเครื่องต่างๆ ใน Swarm Cluster

networks:
  my_overlay_net:
    driver: overlay
    attachable: true # ใส่ไว้เพื่อให้ Container ธรรมดา (ที่ไม่ได้รันแบบ swarm service) สามารถมาเกาะวงนี้ได้ด้วย

```