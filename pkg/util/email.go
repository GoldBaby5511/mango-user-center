package util

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var (
	Email     email
	emailConf struct {
		Host     string
		Port     int
		From     string
		Username string
		AuthCode string
	}
)

func init() {
	viper.UnmarshalKey("email", &emailConf)
}

type email struct{}

func (e *email) Send(to []string, title, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailConf.From)
	m.SetHeader("To", to...)
	// m.SetAddressHeader("Cc", "xxx@163.com", "Dan")  // 抄送
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(emailConf.Host, emailConf.Port, emailConf.Username, emailConf.AuthCode)

	err := d.DialAndSend(m)

	return err
}

// func ConnectAws() *session.Session {
// 	sess, err := session.NewSession(
// 		&aws.Config{
// 			Region: aws.String("ap-northeast-1"),
// 			Credentials: credentials.NewStaticCredentials(
// 				"AKIA55BRPJCRK46MJ7X7",
// 				"3oigv+JNG55hA6O0ywXJAf2H9x5aOmGpbAY1qI9y",
// 				"", // a token will be created when the session it's used.
// 			),
// 		})
// 	if err != nil {
// 		panic("Failed to create AWS session : " + err.Error())
// 	}
// 	return sess
// }

// func SendEmails(sess *session.Session, recipient string, subject string, htmlBody string) {
// 	// defer close(c)
// 	// Create an SES session.
// 	svc := ses.New(sess)
// 	// Assemble the email.
// 	input := &ses.SendEmailInput{
// 		Destination: &ses.Destination{
// 			CcAddresses: []*string{},
// 			ToAddresses: []*string{
// 				aws.String("z772532526@gmail.com"),
// 			},
// 		},
// 		Message: &ses.Message{
// 			Body: &ses.Body{
// 				Html: &ses.Content{
// 					Charset: aws.String("UTF-8"),
// 					Data:    aws.String(htmlBody),
// 				},
// 				// Text: &ses.Content{
// 				// 	Charset: aws.String(CharSet),
// 				// 	Data:    aws.String(TextBody),
// 				// },
// 			},
// 			Subject: &ses.Content{
// 				Charset: aws.String("UTF-8"),
// 				Data:    aws.String(subject),
// 			},
// 		},
// 		Source: aws.String("772532526@qq.com"),
// 		// Uncomment to use a configuration set
// 		//ConfigurationSetName: aws.String(ConfigurationSet),
// 	}
// 	// Attempt to send the email.
// 	result, err := svc.SendEmail(input)
// 	fmt.Println(result, err)
// }
