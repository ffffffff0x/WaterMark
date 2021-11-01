# 简介

给文件夹内图片添加水印的工具

源码来自 : https://blog.csdn.net/czy279470138/article/details/96993285

仅作部分修改

## 效果

```
main.exe -img 11.png -width 300 -hight 300
```

![](./demo.png)

## Usage

单文件添加水印
```
./main.exe -img 1.png
```

文件夹内图片添加水印
```
./main.exe -dir tmp
```

自定义水印 (默认logo.png)
```
./main.exe -logo logo2.png -img 1.png
```

水印区间 (默认皆为600)
```
./main.exe -width 500 -hight 500 -img 1.png
```

---

> create by ffffffff0x
