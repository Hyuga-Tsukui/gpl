#!/bin/bash

# サーバーの起動
echo "Starting clock servers..."

TZ=Asia/Tokyo go run ch8/clock2 -p :8000 &
TZ=Europe/London go run ch8/clock2 -p :8030 &
TZ=America/New_York go run ch8/clock2 -p :8060 &

# サーバーが起動するのを待つ
sleep 2

# クライアントの起動
echo "Starting clock client..."

go run ch8/ex1 -clocks "Tokyo=:8000" -clocks "London=:8030" -clocks "NewYork=:8060"

# 終了処理
echo "Shutting down..."

# サーバーのプロセスを終了
kill $(jobs -p)

echo "All processes terminated."
