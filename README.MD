### wallpaper-mac

#### 介绍
- 调用Macos 原生Api修改桌面壁纸
- 依赖库
  - [macdriver](https://github.com/progrium/macdriver)

#### 安装
```yarn
go get -u github.com/boyyang-love/wallpaper-mac
```

#### 使用
```go
package main

import (
  "fmt"
  "wallpaper-mac/wallpaper"
)

func main() {
  w := wallpaper.NewWallpaper()

  // 设置壁纸
  err := w.SetWallpaper("/Users/boyyang/Downloads/anime-anime-girls-Hatsune-Miku-Vocaloid-blue-eyes-blue-hair-2185805-wallhere.com.jpg")
  if err != nil {
    return
  }
  // 获取当前桌面壁纸路径
  path := w.GetWallpaper()
  fmt.Println(path)
}

```


