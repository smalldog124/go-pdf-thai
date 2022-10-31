# GO-PDF
สร้าง api ที่ generate pdf เนื่อหามีภาษาไทย

1. run service ตัว service จะ run ที port `:3000`
```
go run main.go
```
call api ไปที่ part
```
curl --location --request GET 'http://localhost:3000/order/bill'
```
![My Remote Image](https://miro.medium.com/max/1400/1*GUQUWK-Bb7pgjiw9ggWeuw.png?dl=0)

## ตัวอย่างอื่นที่ lib เขียนไว้
https://github.com/johnfercher/maroto/tree/master/internal/examples