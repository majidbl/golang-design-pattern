
package main

import (
	"fmt"
	"log"

	usecase "design_pattern/pattern/templateMethod"
	)

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &usecase.SMS{}
	o := usecase.OTP{
		Otp: smsOTP,
	}
	err := o.GenAndSendOTP(4)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("")
	emailOTP := &usecase.Email{}
	o = usecase.OTP{
		Otp: emailOTP,
	}
	err = o.GenAndSendOTP(4)
	if err != nil{
		log.Println(err)
	}
}
