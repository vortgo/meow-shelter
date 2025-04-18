package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<title>Мяу-Мятеж</title>
<style>
  body {
    background: linear-gradient(135deg, #7effc1, #a483ff);
    font-family: "Comic Sans MS", cursive, sans-serif;
    color: #2d2d2d;
    text-align: center;
    padding: 50px;
  }
  .cat {
    font-size: 120px;
    margin-bottom: 10px;
  }
  .title {
    font-size: 36px;
    font-weight: bold;
    margin-bottom: 20px;
  }
  .tagline {
    font-size: 20px;
    color: #333;
  }
</style>
</head>
<body>
<div class="cat">🐾😼🐾</div>
<div class="title">Слишком милая, чтобы слушаться</div>
<div class="tagline">Добро пожаловать в штаб 🐈‍⬛ Кошачьего Бунта</div>
</body>
</html>
`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
