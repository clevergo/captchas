# Captchas Drivers

```shell
$ go get clevergo.tech/captchas/drivers
```

```go
import "go get clevergo.tech/captchas/drivers"
```

## Digit

```go
// all options are optional.
opts := []drivers.DigitOption{
	drivers.DigitHeight(50),
	drivers.DigitWidth(120),
	drivers.DigitLength(6),
	drivers.DigitMaxSkew(0.8),
	drivers.DigitDotCount(80),
}
driver := drivers.NewDigit(opts...)
```

## Audio

```go
// all options are optional.
opts := []drivers.AudioOption{
	drivers.AudioLangauge("en"),
	drivers.AudioLength(6),
}
driver := drivers.NewAudio(opts...)
```

## Math

```go
// all options are optional.
opts := []drivers.MathOption{
	drivers.MathHeight(50),
	drivers.MathWidth(120),
	drivers.MathNoiseCount(0),
	drivers.MathFonts([]string{}),
	drivers.MathBGColor(&color.RGBA{}),
}
driver := drivers.NewMath(opts...)
```

## String

```go
// all options are optional.
opts := []drivers.StringOption{
	drivers.StringHeight(50),
	drivers.StringWidth(120),
	drivers.StringLength(4),
	drivers.StringNoiseCount(0),
	drivers.StringFonts([]string{}),
	drivers.StringSource("abcdefghijklmnopqrstuvwxyz"),
	drivers.StringBGColor(&color.RGBA{}),
}
driver := drivers.NewString(opts...)
```

## Chinese

```go
// all options are optional.
opts := []drivers.ChineseOption{
	drivers.ChineseHeight(50),
	drivers.ChineseWidth(120),
	drivers.ChineseLength(4),
	drivers.ChineseNoiseCount(0),
	drivers.ChineseFonts([]string{"wqy-microhei.ttc"}),
	drivers.ChineseSource("零一二三四五六七八九十"),
	drivers.ChineseBGColor(&color.RGBA{}),
}
driver := drivers.NewChinese(opts...)
```
