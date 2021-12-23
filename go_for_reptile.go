package main
//爬虫 
import (
	"crypto/tls"
	// "crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)
 
//客户端对服务器校验
func main() {
	//CertPool代表一个证书集合/证书池。
	//创建一个CertPool
	// pool := x509.NewCertPool()
	// caCertPath := "/Users/xuchangliang/object/sign/DigiCert Global Root CA.pem"
	// //调用ca.crt文件
	// caCrt, err := ioutil.ReadFile(caCertPath)
	// if err != nil {
	// 	fmt.Println("ReadFile err:", err)
	// 	return
	// }
	// //解析证书
	// pool.AppendCertsFromPEM(caCrt)
	// pool.AddCert(caCrt);
	tr := &http.Transport{
		//把从服务器传过来的非叶子证书，添加到中间证书的池中，使用设置的根证书和中间证书对叶子证书进行验证。
		// TLSClientConfig: &tls.Config{RootCAs: pool},		
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false}, //InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true,//		//则不会校验证书以及证书中的主机名和服务器主机名是否一致。		
	}
	client := &http.Client{Transport: tr}
    // url := "https://a01-g-naver-vod.pstatic.net/navertv/c/read/v2/VOD_ALPHA/navertv_2020_06_12_1974/hls/dbf0065b-ac6b-11ea-9071-246e963a41ed-000011.ts?__gda__=1639723479_6e7f8a9c7dff8ac2ce29fce6c00ff6b0"
	// url := "https://tv.naver.com/v/14139840"
	// url := "https://a01-g-naver-vod.pstatic.net/navertv/c/read/v2/VOD_ALPHA/navertv_2020_07_13_1445/hls/eec4c02d-c4ce-11ea-8c09-246e963a41ed-000000.ts?__gda__=1639749036_5ad700aa8a7e86cc75039f6ee07df4d8"
	// url := "https://5vgqukv.com:1443/hls/c9ecdb81e04b498d7cc33ba2ffbf1f18/200kb/hls/key.key"
	// url := "https://5vgquku.com:1443/hls/b92096a9cac455d71cd26f7a4c03554d/200kb/hls/T9SjYmyx.ts"
	    url:="https://5vgquku.com:1443/hls/b92096a9cac455d71cd26f7a4c03554d/200kb/hls/Jyk2ofRV.ts"
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("./Jyk2ofRV.ts", body, 0777)
	fmt.Println(string(body))
}
