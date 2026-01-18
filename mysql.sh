#!/bin/bash

# Docker Compose MySQL 连接配置
MYSQL_USER="root"
MYSQL_PASS="123456"  # 从 docker-compose.yml 获取
MYSQL_HOST="localhost"
DB_NAME="go_app"
CONTAINER_NAME="mysql_service"

# 函数：等待 MySQL 服务可用
wait_for_mysql() {
    echo "等待 MySQL 服务启动..."
    
    while ! docker exec $CONTAINER_NAME mysqladmin -u"$MYSQL_USER" -p"$MYSQL_PASS" ping --silent; do
        echo "MySQL 尚未就绪，等待 5 秒..."
        sleep 5
    done
    
    echo "MySQL 服务已启动"
}

# 检查 Docker 容器是否运行
if ! docker ps | grep -q $CONTAINER_NAME; then
    echo "MySQL 容器未运行，正在启动..."
    
    # 如果使用 docker-compose
    if command -v docker-compose &> /dev/null; then
        docker-compose up -d mysql-service
    else
        docker compose up -d mysql-service
    fi
    
    # 等待 MySQL 启动
    wait_for_mysql
else
    echo "MySQL 容器已在运行"
fi

# 检查并创建数据库
if docker exec $CONTAINER_NAME mysql -u "$MYSQL_USER" -p"$MYSQL_PASS" -e "USE $DB_NAME" 2>/dev/null; then
    echo "数据库 '$DB_NAME' 已存在"
else
    echo "正在创建数据库 '$DB_NAME'..."
    
    # 创建数据库
    docker exec $CONTAINER_NAME mysql -u "$MYSQL_USER" -p"$MYSQL_PASS" -e "CREATE DATABASE IF NOT EXISTS $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
    
    if [ $? -eq 0 ]; then
        echo "数据库 '$DB_NAME' 已成功创建"
        
        # 可选：显示所有数据库
        echo "当前所有数据库:"
        docker exec $CONTAINER_NAME mysql -u "$MYSQL_USER" -p"$MYSQL_PASS" -e "SHOW DATABASES;"
    else
        echo "创建数据库失败"
        exit 1
    fi
fi