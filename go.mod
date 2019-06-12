module siot-util-ali

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)

require (
	github.com/astaxie/beego v1.11.1
	github.com/souliot/siot-util v0.0.0-20190518164625-302008e3c0bc
)
