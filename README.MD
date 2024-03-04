### wallpaper-mac

#### 背景

使用go语言的跨平台开发框架wails开发桌面应用的时候，
需要修改桌面壁纸，但是wails并没有提供修改桌面壁纸的接口，
只能通过AppleScript脚本来修改桌面壁纸，正常情况下是没问题的。
但是如果想要上架苹果应用商店，由于系统权限问题，如若需要在沙盒环境下运行，
AppleScript脚本将无法正常运行，需要在沙盒模式下运行脚本，步骤相对比较复杂，
所以需要寻找其他方法来修改桌面壁纸。

#### 思路

通过调用Macos原生Api来修改桌面壁纸，具体实现方法可以参考[macdriver](https://github.com/progrium/macdriver)，
基于该库我将壁纸修改的功能简单的封装成了一个库，方便其他开发者调用。

#### 安装
```yarn
go get -u github.com/boyyang-love/wallpaper-mac
```

#### 使用
```go
package main

import (
  "fmt"
  "github.com/boyyang-love/wallpaper-mac/wallpaper"
)

func main() {
  w := wallpaper.NewWallpaper()

  // 设置壁纸
  err := w.SetWallpaper("/Users/boyyang/Downloads/test.jpg")
  if err != nil {
    return
  }
  // 获取当前桌面壁纸路径
  path := w.GetWallpaper()
  fmt.Println(path)
}

```


