package main
import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"github.com/gookit/color"
)

var Input_img = flag.String("img", "", "输入文件名")
var Input_dir = flag.String("dir", "", "输入文件夹名")
var Input_logo = flag.String("logo", "logo.png", "输入 logo 名(例如:logo.png)")
var Input_width = flag.Int("width", 600, "输入宽度")
var Input_hight = flag.Int("hight", 600, "输入高度")

func main() {
	flag.Parse()

	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	if *Input_img == ""{
		*Input_img="null"
	}else {
		imgwater()
	}

	if *Input_dir == ""{
		*Input_dir="null"
	}else {
		imgwater_dir()
	}

}

func imgwater() {

	color.Cyan.Println("目标图片:",*Input_img)

	srcImg := *Input_img             // 原始图片
	imgWaterMarkPath := *Input_logo // 水印图片

	// 原始图片
	originalImg, err := os.Open(srcImg)
	if err != nil {
		color.Red.Println("打开原始图片出错")
	}
	//
	defer originalImg.Close()

	// 水印图片
	waterMark, err := os.Open(imgWaterMarkPath)
	if err != nil {
		color.Red.Println("打开水印图片出错")
	}
	defer waterMark.Close()

	waterMarkImg, err := png.Decode(waterMark)
	if err != nil {
		color.Red.Println("把水印图片解码为结构体时出错")
	}

	buff := make([]byte, 512)

	file, err := os.Open(srcImg)
	if err != nil {
		log.Println(err)
	}

	_, err = file.Read(buff)
	if err != nil {
		log.Println(err)
	}

	imgType := http.DetectContentType(buff)

	if imgType == "image/jpeg" {
		//fmt.Println("这是JPG文件")
		imgJpeg, err := jpeg.Decode(originalImg)
		if err != nil {
			color.Red.Println("把jpeg图片解码为结构体时出错")
		}

		b := imgJpeg.Bounds()
		waterMarkWidth := b.Max.X
		waterMarkHeight := b.Max.Y
		//fmt.Println("jpeg原始图片宽", waterMarkWidth, "jpeg原始图片高", waterMarkHeight)

		m := image.NewRGBA(b)

		// 原始图片
		draw.Draw(m, b, imgJpeg, image.ZP, draw.Src)
		// 水印图片
		for i := 0; i < waterMarkWidth; i++ {
			offsetWidth := *Input_width * i
			if offsetWidth < waterMarkWidth {
				for j := 0; j < waterMarkHeight; j++ {
					offsetHeight := *Input_hight * j
					if offsetHeight < waterMarkHeight {
						offset := image.Pt(offsetWidth, offsetHeight)
						draw.Draw(m, waterMarkImg.Bounds().Add(offset), waterMarkImg, image.ZP, draw.Over)
						//fmt.Println("jpegOffset", offset)
					}
				}
			}
		}


		// 生成新图片new.jpg,并设置图片质量 100
		imgNew, err := os.Create(*Input_img) // 这里可以设置为上传图片 srcImg


		if err != nil {
			log.Println(err)
		}

		err = jpeg.Encode(imgNew, m, &jpeg.Options{100})
		if err != nil {
			log.Println(err)
		}
		// png.Encode(imgw, m)
		defer imgNew.Close()

		color.Green.Print("添加JPG水印图片结束请查看\n")
	}

	if imgType == "image/png" {
		//fmt.Println("这是PNG文件")
		imgPng, err := png.Decode(originalImg)
		if err != nil {
			color.Red.Println("把PNG图片解码为结构体时出错")
		}
		b := imgPng.Bounds()
		waterMarkWidth := b.Max.X
		waterMarkHeight := b.Max.Y
		//fmt.Println("png原始图片宽", waterMarkWidth, "png原始图片高", waterMarkHeight)

		m := image.NewRGBA(b)

		// 原始图片
		draw.Draw(m, b, imgPng, image.ZP, draw.Src)

		// 水印图片
		for i := 0; i < waterMarkWidth; i++ {
			// resWidth := 0
			offsetWidth := *Input_width * i
			// fmt.Println(offsetWidth)
			if offsetWidth < waterMarkWidth {
				for j := 0; j < waterMarkHeight; j++ {
					offsetHeight := *Input_hight * j
					if offsetHeight < waterMarkHeight {
						offset := image.Pt(offsetWidth, offsetHeight)
						draw.Draw(m, waterMarkImg.Bounds().Add(offset), waterMarkImg, image.ZP, draw.Over)
						//fmt.Println("pngOffset", offset)
					}
				}
			}
		}

		// 生成新图片new.png
		imgNew, err := os.Create(*Input_img) // 这里可以设置为上传图片 srcImg
		if err != nil {
			log.Println(err)
		}

		err = png.Encode(imgNew, m)
		if err != nil {
			log.Println(err)
		}
		defer imgNew.Close()
		color.Green.Println("添加PNG水印图片结束请查看")
	}
	if imgType == "image/gif" {
		color.Yellow.Println("暂不支持 gif 格式")
	}
}

func imgwater_dir() {
	color.Cyan.Println("目标文件夹:",*Input_dir)
	files, _ := ioutil.ReadDir(*Input_dir)
	for _, f := range files {
		*Input_img=*Input_dir + "/" +f.Name()
		imgwater()
	}
}
