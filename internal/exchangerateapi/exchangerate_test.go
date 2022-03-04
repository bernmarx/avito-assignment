package exchangerateapi

// type testdata struct {
// 	Info []byte
// }

// func (d *testdata) Read(p []byte) (n int, err error) {
// 	p = d.Info
// 	return len(p), nil
// }
// func (d *testdata) Close() error {
// 	return nil
// }

// func TestGetExchangeRate(t *testing.T) {
// 	b := []byte(`{"conversion_rate":0.5}`)
// 	data := testdata{Info: b}
// 	resp := &http.Response{Body: &data}

// 	ctrl := gomock.NewController(t)

// 	m := NewMockhttpClient(ctrl)

// 	m.EXPECT().Get("https://v6.exchangerate-api.com/v6/"+os.Getenv("ER_API_KEY")+"/pair/"+
// 		"RUB"+"/"+"USD").Return(resp, nil)

// 	e := ExchangeRate{m}

// 	eR, err := e.GetExchangeRate("RUB", "USD")

// 	assert.Nil(t, err)
// 	assert.Equal(t, float32(0.5), eR)
// }
