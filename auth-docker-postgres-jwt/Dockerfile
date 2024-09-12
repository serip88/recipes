# Dựa trên một base image Go phù hợp với phiên bản Go của bạn
FROM golang:1.20-alpine

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy gói go.mod và go.sum để quản lý dependency
COPY go.mod go.sum ./

# Tải về các dependency
RUN go mod download

# Copy toàn bộ source code vào thư mục làm việc
COPY . .

# Xây dựng ứng dụng
RUN go build -o main .

# Thiết lập cổng để chạy ứng dụng
EXPOSE 3000

# Command để chạy ứng dụng khi container khởi động
CMD ["./main"]