https://stackoverflow.com/questions/63588254/how-to-set-up-an-https-server-with-a-self-signed-certificate-in-golang


mkdir golangssl
cd golangssl
Generate a selfsigned key and cert (install openssl if not yet installed)

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650


// Run the server (needs root rights, because of binding of port 443), not needed ?

sudo go run main.go or go run main.go + https://localhost:10443

// Execute a request (-k for ignoring the self signed certificate)
// curl -k https://localhost, not needed ?
