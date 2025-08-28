#!/bin/bash

set -e

log() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [INFO] $1"
}

error() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [ERROR] $1" >&2
}

log "[1] 停止服务"
systemctl stop game_engine

log "[2] 检测服务是否停止"
# 使用 systemctl is-active 检查服务是否真的停止（仅关注状态，不输出详细日志）
if systemctl is-active --quiet game_engine; then
  error "服务未成功停止，无法继续操作"
  exit 1
else
  log "服务已成功停止"
fi

log "[3] 拷贝可执行文件到/usr/share/game_engine"
/bin/cp -f ./bin/game_engine /usr/share/game_engine/
/bin/cp -f ./config.yaml /usr/share/game_engine/

log "[4] 重启game_engine服务"
systemctl restart game_engine

log "[5] 查看game_engine服务状态"
systemctl status game_engine
