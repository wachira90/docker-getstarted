FROM nginx:alpine

# คัดลอกไฟล์คอนฟิกูเรชันของ Nginx
COPY nginx.conf /etc/nginx/default.conf

# คัดลอกไฟล์ HTML ไปยังโฟลเดอร์เริ่มต้นของ Nginx
COPY index.html /usr/share/nginx/html/index.html

# เปิดพอร์ต 80
EXPOSE 80

# docker build -t www:v1 .

# docker run -it --name www -p 8022:80 www:v1
# docker run -it --name www -p 8022:80 docker.io/library/www:v1