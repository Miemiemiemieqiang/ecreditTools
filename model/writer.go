package model

type Writer interface {
	Write(response *Response)
}

type ExcelWriter struct {}

func NewExcelWriter() *ExcelWriter {
	return &ExcelWriter{}
}


func (e ExcelWriter) Write(response *Response) {
	//f, _ := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	//writer := csv.NewWriter(f)
	//writer.Write()
	//_, err = f.Write(data)
	//if err1 := f.Close(); err == nil {
	//	err = err1
	//}
	//return err
	//csv.NewWriter()
	//response
	//panic("implement me")
}



