package templateMethod

import domain "design_pattern/domain/templateMethod"


// type otp struct {
// }

// func (o *otp) GenAndSendOTP(Otp Otp, otpLength int) error {
//  otp := Otp.genRandomOTP(otpLength)
//  Otp.saveOTPCache(otp)
//  message := Otp.getMessage(otp)
//  err := Otp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  Otp.publishMetric()
//  return nil
// }


type OTP struct {
	Otp domain.Otp
}

func (o *OTP) GenAndSendOTP(otpLength int) error {
	otp := o.Otp.GenRandomOTP(otpLength)
	o.Otp.SaveOTPCache(otp)
	message := o.Otp.GetMessage(otp)
	err := o.Otp.SendNotification(message)
	if err != nil {
		return err
	}
	o.Otp.PublishMetric()
	return nil
}
